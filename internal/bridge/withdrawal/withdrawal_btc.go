package withdrawal

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/hyle-team/tss-svc/internal/bridge/clients/bitcoin"
	"github.com/hyle-team/tss-svc/internal/db"
	"github.com/hyle-team/tss-svc/internal/p2p"
	"github.com/hyle-team/tss-svc/internal/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/anypb"
)

var _ DepositSigningData = BitcoinWithdrawalData{}

type BitcoinWithdrawalData struct {
	ProposalData *p2p.BitcoinProposalData
	SignedInputs [][]byte
}

func (e BitcoinWithdrawalData) DepositIdentifier() db.DepositIdentifier {
	identifier := db.DepositIdentifier{}

	if e.ProposalData == nil || e.ProposalData.DepositId == nil {
		return identifier
	}

	identifier.ChainId = e.ProposalData.DepositId.ChainId
	identifier.TxHash = e.ProposalData.DepositId.TxHash
	identifier.TxNonce = int(e.ProposalData.DepositId.TxNonce)

	return identifier
}

func (e BitcoinWithdrawalData) ToPayload() *anypb.Any {
	pb, _ := anypb.New(e.ProposalData)

	return pb
}

type BitcoinWithdrawalConstructor struct {
	client *bitcoin.Client
	tssPkh *btcutil.AddressPubKeyHash
}

func NewBitcoinConstructor(client *bitcoin.Client, tssPub *ecdsa.PublicKey) *BitcoinWithdrawalConstructor {
	tssPkh, err := bitcoin.PubKeyToPkhCompressed(tssPub, client.ChainParams())
	if err != nil {
		panic(fmt.Sprintf("failed to create TSS public key hash: %v", err))
	}

	return &BitcoinWithdrawalConstructor{client: client, tssPkh: tssPkh}
}

func (c *BitcoinWithdrawalConstructor) FromPayload(payload *anypb.Any) (BitcoinWithdrawalData, error) {
	proposalData := &p2p.BitcoinProposalData{}
	if err := payload.UnmarshalTo(proposalData); err != nil {
		return BitcoinWithdrawalData{}, errors.Wrap(err, "failed to unmarshal proposal data")
	}

	return BitcoinWithdrawalData{ProposalData: proposalData}, nil
}

func (c *BitcoinWithdrawalConstructor) FormSigningData(deposit db.Deposit) (BitcoinWithdrawalData, error) {
	tx, sigHashes, err := c.client.CreateUnsignedWithdrawalTx(deposit, c.tssPkh.EncodeAddress())
	if err != nil {
		return BitcoinWithdrawalData{}, errors.Wrap(err, "failed to create unsigned transaction")
	}

	var buf bytes.Buffer
	if err = tx.Serialize(&buf); err != nil {
		return BitcoinWithdrawalData{}, errors.Wrap(err, "failed to serialize transaction")
	}

	return BitcoinWithdrawalData{
		ProposalData: &p2p.BitcoinProposalData{
			DepositId: &types.DepositIdentifier{
				ChainId: deposit.ChainId,
				TxNonce: uint32(deposit.TxNonce),
				TxHash:  deposit.TxHash,
			},
			SerializedTx: buf.Bytes(),
			SigData:      sigHashes,
		},
	}, nil
}

func (c *BitcoinWithdrawalConstructor) IsValid(data BitcoinWithdrawalData, deposit db.Deposit) (bool, error) {
	tx := wire.MsgTx{}
	if err := tx.Deserialize(bytes.NewReader(data.ProposalData.SerializedTx)); err != nil {
		return false, errors.Wrap(err, "failed to deserialize transaction")
	}

	outputsSum, err := c.validateOutputs(&tx, deposit)
	if err != nil {
		return false, errors.Wrap(err, "failed to validate outputs")
	}

	usedInputs, err := c.findUsedInputs(&tx)
	if err != nil {
		return false, errors.Wrap(err, "failed to find used inputs")
	}

	inputsSum, err := c.validateInputs(&tx, usedInputs, data.ProposalData.SigData)
	if err != nil {
		return false, errors.Wrap(err, "failed to validate inputs")
	}

	if err = c.validateChange(&tx, usedInputs, inputsSum, outputsSum); err != nil {
		return false, errors.Wrap(err, "failed to validate change")
	}

	return true, nil
}

func (c *BitcoinWithdrawalConstructor) findUsedInputs(tx *wire.MsgTx) (map[OutPoint]btcjson.ListUnspentResult, error) {
	unspent, err := c.client.ListUnspent()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get available UTXOs")
	}

	usedInputs := make(map[OutPoint]btcjson.ListUnspentResult, len(tx.TxIn))
	for _, inp := range tx.TxIn {
		if inp == nil {
			return nil, errors.New("nil input in transaction")
		}

		for _, u := range unspent {
			if u.TxID != inp.PreviousOutPoint.Hash.String() ||
				u.Vout != inp.PreviousOutPoint.Index {
				continue
			}

			outPoint := OutPoint{TxID: u.TxID, Index: u.Vout}
			if _, found := usedInputs[outPoint]; found {
				return nil, errors.New(fmt.Sprintf("double spending detected for %s:%d", u.TxID, u.Vout))
			}

			usedInputs[outPoint] = u
			break
		}
	}

	if len(usedInputs) != len(tx.TxIn) {
		return nil, errors.New("not all inputs were found")
	}

	return usedInputs, nil
}

func (c *BitcoinWithdrawalConstructor) validateOutputs(tx *wire.MsgTx, deposit db.Deposit) (int64, error) {
	outputsSum, receiverIdx := int64(0), 0
	switch len(tx.TxOut) {
	case 2:
		// 1st output is for the change, 2nd is for the receiver
		receiverIdx = 1
		changeOutput := tx.TxOut[0]
		outputsSum += changeOutput.Value

		outScript, err := txscript.PayToAddrScript(c.tssPkh)
		if err != nil {
			return 0, errors.Wrap(err, "failed to create change output script")
		}
		if !bytes.Equal(changeOutput.PkScript, outScript) {
			return 0, errors.New("invalid change output script")
		}

		fallthrough
	case 1:
		receiverOutput := tx.TxOut[receiverIdx]
		withdrawalAmount, ok := new(big.Int).SetString(*deposit.WithdrawalAmount, 10)
		if !ok || receiverOutput.Value != withdrawalAmount.Int64() {
			return 0, errors.New("invalid withdrawal amount")
		}
		outputsSum += receiverOutput.Value

		outAddr, err := btcutil.DecodeAddress(*deposit.Receiver, c.client.ChainParams())
		if err != nil {
			return 0, errors.Wrap(err, "failed to decode receiver address")
		}
		outScript, err := txscript.PayToAddrScript(outAddr)
		if err != nil {
			return 0, errors.Wrap(err, "failed to create change output script")
		}
		if !bytes.Equal(receiverOutput.PkScript, outScript) {
			return 0, errors.New("invalid receiver output script")
		}
	default:
		return 0, errors.New("invalid number of transaction outputs")
	}

	return outputsSum, nil
}

func (c *BitcoinWithdrawalConstructor) validateInputs(
	tx *wire.MsgTx,
	inputs map[OutPoint]btcjson.ListUnspentResult,
	sigHashes [][]byte,
) (int64, error) {
	if sigHashes == nil || len(sigHashes) != len(tx.TxIn) {
		return 0, errors.New("invalid signature hashes")
	}

	inputsSum := int64(0)
	for idx, inp := range tx.TxIn {
		unspent := inputs[OutPoint{TxID: inp.PreviousOutPoint.Hash.String(), Index: inp.PreviousOutPoint.Index}]

		scriptDecoded, err := hex.DecodeString(unspent.ScriptPubKey)
		if err != nil {
			return 0, errors.Wrap(err, fmt.Sprintf("failed to decode script for input %d", idx))
		}
		sigHash, err := txscript.CalcSignatureHash(scriptDecoded, bitcoin.SigHashType, tx, idx)
		if err != nil {
			return 0, errors.Wrap(err, fmt.Sprintf("failed to calculate signature hash for input %d", idx))
		}
		if !bytes.Equal(sigHashes[idx], sigHash) {
			return 0, errors.New(fmt.Sprintf("invalid signature hash for input %d", idx))
		}

		inputsSum += bitcoin.ToAmount(unspent.Amount, bitcoin.Decimals).Int64()
	}

	return inputsSum, nil
}

func (c *BitcoinWithdrawalConstructor) validateChange(tx *wire.MsgTx, inputs map[OutPoint]btcjson.ListUnspentResult, inputsSum, outputsSum int64) error {
	actualFee := inputsSum - outputsSum
	if actualFee <= 0 {
		return errors.New("invalid change amount")
	}

	mockedTx, err := c.mockTransaction(tx, inputs)
	if err != nil {
		return errors.Wrap(err, "failed to mock transaction")
	}

	var (
		targetFeeRate = bitcoin.DefaultFeeRateBtcPerKvb * 1e5 // btc/kB -> sat/byte
		feeTolerance  = 0.1 * targetFeeRate                   // 10%
		estimatedSize = mockedTx.SerializeSize()
		actualFeeRate = float64(actualFee) / float64(estimatedSize)
	)

	if math.Abs(actualFeeRate-targetFeeRate) > feeTolerance {
		return errors.New(fmt.Sprintf("provided fee rate %f is not within %f of target %f", actualFeeRate, feeTolerance, targetFeeRate))
	}

	return nil
}

type OutPoint struct {
	TxID  string
	Index uint32
}

func (c *BitcoinWithdrawalConstructor) mockTransaction(tx *wire.MsgTx, inputs map[OutPoint]btcjson.ListUnspentResult) (*wire.MsgTx, error) {
	mockKey, err := btcec.NewPrivateKey()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create mock private key")
	}

	mockedTx := tx.Copy()

	for i, inp := range mockedTx.TxIn {
		unspent := inputs[OutPoint{TxID: inp.PreviousOutPoint.Hash.String(), Index: inp.PreviousOutPoint.Index}]
		scriptDecoded, err := hex.DecodeString(unspent.ScriptPubKey)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to decode script for input %d", i))
		}

		sig, err := txscript.SignatureScript(mockedTx, i, scriptDecoded, bitcoin.SigHashType, mockKey, true)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to sign input %d", i))
		}

		mockedTx.TxIn[i].SignatureScript = sig
	}

	return mockedTx, nil
}

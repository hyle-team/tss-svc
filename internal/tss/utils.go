package tss

import (
	"github.com/hyle-team/tss-svc/internal/core"
)

const (
	OutChannelSize = 1000
	EndChannelSize = 1
	MsgsCapacity   = 100
)

type partyMsg struct {
	Sender      core.Address
	WireMsg     []byte
	IsBroadcast bool
}
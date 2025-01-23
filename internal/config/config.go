package config

import (
	"github.com/hyle-team/tss-svc/internal/bridge/chain"
	connector "github.com/hyle-team/tss-svc/internal/core/connector/config"
	p2p "github.com/hyle-team/tss-svc/internal/p2p/config"
	vaulter "github.com/hyle-team/tss-svc/internal/secrets/vault/config"
	tss "github.com/hyle-team/tss-svc/internal/tss/config"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	vaulter.Vaulter
	Listenerer
	p2p.PartiesConfigurator
	tss.ParamsConfigurator
	chain.Chainer
	connector.ConnectorConfigurer
}

type config struct {
	getter kv.Getter

	comfig.Logger
	pgdb.Databaser
	vaulter.Vaulter
	Listenerer
	p2p.PartiesConfigurator
	tss.ParamsConfigurator
	chain.Chainer
	connector.ConnectorConfigurer
}

func New(getter kv.Getter) Config {
	return &config{
		getter:              getter,
		Vaulter:             vaulter.NewVaulter(),
		Logger:              comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Databaser:           pgdb.NewDatabaser(getter),
		Listenerer:          NewListenerer(getter),
		PartiesConfigurator: p2p.NewPartiesConfigurator(getter),
		ParamsConfigurator:  tss.NewParamsConfigurator(getter),
		Chainer:             chain.NewChainer(getter),
		ConnectorConfigurer: connector.NewConnectorConfigurer(getter),
	}
}

package config

import (
	"crypto/tls"
	core "github.com/hyle-team/tss-svc/internal/core"
	connector "github.com/hyle-team/tss-svc/internal/core/connector"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"reflect"
	"time"
)

type ConnectorConfigurer interface {
	CoreConnectorConfig() ConnectorConfig
}

type ConnectorConfig struct {
	Settings   connector.ConnectorSettings
	Connection *grpc.ClientConn
}

type Connection struct {
	Addr      string `fig:"addr,required"`
	EnableTLS bool   `fig:"enable_tls"`
}

type configurer struct {
	once   comfig.Once
	getter kv.Getter
}

func NewConnectorConfigurer(getter kv.Getter) ConnectorConfigurer {
	return &configurer{
		getter: getter,
	}
}

func (c *configurer) CoreConnectorConfig() ConnectorConfig {
	return c.once.Do(func() interface{} {
		const yamlKey = "core_connector"
		var cfg struct {
			Settings   connector.ConnectorSettings `fig:"settings,required"`
			Connection Connection                  `fig:"connection,required"`
		}

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks, accountHook).
			From(kv.MustGetStringMap(c.getter, yamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to configure core connector"))
		}

		connectSecurityOptions := grpc.WithInsecure()
		if cfg.Connection.EnableTLS {
			tlsConfig := &tls.Config{
				MinVersion: tls.VersionTLS13,
			}
			connectSecurityOptions = grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig))
		}
		keepaliveOptions := grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                20 * time.Second, // wait time before ping if no activity
			Timeout:             5 * time.Second,  // ping timeout
			PermitWithoutStream: true,
		})

		client, err := grpc.NewClient(cfg.Connection.Addr, connectSecurityOptions, keepaliveOptions)
		if err != nil {
			panic(errors.Wrap(err, "failed to connect to core via gRPC"))
		}

		return ConnectorConfig{
			Settings:   cfg.Settings,
			Connection: client,
		}
	}).(ConnectorConfig)
}

var accountHook = figure.Hooks{
	"core.Account": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			acc, err := core.NewAccount(v)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to create account")
			}

			return reflect.ValueOf(*acc), nil
		default:
			return reflect.Value{}, errors.Errorf("unsupported conversion from %T", value)
		}
	},
}

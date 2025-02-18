package ctx

import (
	"context"

	"github.com/hyle-team/tss-svc/internal/bridge"
	bridgeTypes "github.com/hyle-team/tss-svc/internal/bridge/clients"
	"github.com/hyle-team/tss-svc/internal/core"
	"github.com/hyle-team/tss-svc/internal/db"
	"github.com/hyle-team/tss-svc/internal/p2p"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	dbKey          ctxKey = iota
	loggerKey      ctxKey = iota
	clientsKey     ctxKey = iota
	processorKey   ctxKey = iota
	broadcasterKey ctxKey = iota
	selfKey        ctxKey = iota
)

func DBProvider(q db.DepositsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {

		return context.WithValue(ctx, dbKey, q)
	}
}

// DB always returns unique connection
func DB(ctx context.Context) db.DepositsQ {
	return ctx.Value(dbKey).(db.DepositsQ).New()
}

func LoggerProvider(l *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {

		return context.WithValue(ctx, loggerKey, l)
	}
}
func Logger(ctx context.Context) *logan.Entry {

	return ctx.Value(loggerKey).(*logan.Entry)
}

func ClientsProvider(cr bridgeTypes.Repository) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, clientsKey, cr)
	}
}

func Clients(ctx context.Context) bridgeTypes.Repository {
	return ctx.Value(clientsKey).(bridgeTypes.Repository)
}

func FetcherProvider(processor *bridge.DepositFetcher) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, processorKey, processor)
	}
}

func Fetcher(ctx context.Context) *bridge.DepositFetcher {
	return ctx.Value(processorKey).(*bridge.DepositFetcher)
}

func BroadcasterProvider(b *p2p.Broadcaster) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, broadcasterKey, b)
	}
}

func Broadcaster(ctx context.Context) *p2p.Broadcaster {
	return ctx.Value(broadcasterKey).(*p2p.Broadcaster)
}

func SelfProvider(self core.Address) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, selfKey, self)
	}
}

func Self(ctx context.Context) core.Address {
	return ctx.Value(selfKey).(core.Address)
}

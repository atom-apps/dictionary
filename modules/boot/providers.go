package boot

import (
	"context"
	"time"

	"github.com/atom-providers/log"
	microGoMicro "github.com/atom-providers/micro-gomicro"
	"github.com/go-micro/plugins/v4/logger/zap"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
	"github.com/rogeecn/gomicro-plugins/registry/etcd"
	"github.com/samber/lo"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: provideGoMicroOptions},
	}
}

func provideGoMicroOptions(opts ...opt.Option) error {
	options := []micro.Option{
		micro.Name(atom.AppName),
		micro.Version(atom.AppVersion),
		micro.RegisterTTL(time.Second * 30),
		micro.RegisterInterval(time.Second * 15),
		micro.Broker(broker.DefaultBroker),
	}

	lo.ForEach(options, func(opt micro.Option, i int) {
		_ = container.Container.Provide(func() micro.Option {
			return opt
		}, microGoMicro.GroupGoMicroOptions)
	})

	_ = container.Container.Provide(func(ctx context.Context, log *log.Logger, client *clientv3.Client) micro.Option {
		logger, _ := zap.NewLogger(
			zap.WithLogger(log.RawLogger),
		)

		return micro.Registry(etcd.NewRegistry(
			registry.Logger(logger),
			registry.Timeout(time.Second*5),
			etcd.Client(client),
		))
	}, microGoMicro.GroupGoMicroOptions)
	return nil
}

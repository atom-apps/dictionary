package boot

import (
	"context"
	"time"

	"github.com/atom-providers/log"
	"github.com/go-micro/plugins/v4/logger/zap"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
	"github.com/rogeecn/gomicro-plugins/registry/etcd"
	"go-micro.dev/v4/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: provideGoMicroOptions},
	}
}

func provideGoMicroOptions(opts ...opt.Option) error {
	_ = container.Container.Provide(func(ctx context.Context, log *log.Logger, client *clientv3.Client) registry.Registry {
		logger, _ := zap.NewLogger(
			zap.WithLogger(log.Logger),
		)

		return etcd.NewRegistry(
			registry.Logger(logger),
			registry.Timeout(time.Second*5),
			etcd.Client(client),
		)
	})

	return nil
}

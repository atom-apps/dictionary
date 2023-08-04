package boot

import (
	"time"

	microGoMicro "github.com/atom-providers/micro-gomicro"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
	"github.com/samber/lo"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/registry"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: provideName},
	}
}

func provideName(opts ...opt.Option) error {
	options := []micro.Option{
		micro.Name(atom.AppName),
		micro.Version(atom.AppVersion),
		micro.RegisterTTL(time.Second * 30),
		micro.RegisterInterval(time.Second * 15),
		micro.Registry(registry.DefaultRegistry),
		micro.Broker(broker.DefaultBroker),
	}
	lo.ForEach(options, func(opt micro.Option, i int) {
		_ = container.Container.Provide(func() micro.Option {
			return opt
		}, microGoMicro.GroupGoMicroOptions)
	})

	return nil
}

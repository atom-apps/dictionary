//go:generate atomctl gen routes
//go:generate swag fmt
//go:generate swag init -ot json
package main

import (
	"context"
	"fmt"

	"github.com/atom-apps/dictionary/modules/boot"
	v1 "github.com/atom-apps/dictionary/proto/v1"
	"github.com/atom-providers/etcd"
	"github.com/atom-providers/log"
	serviceGoMicro "github.com/atom-providers/service-gomicro"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/spf13/cobra"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/selector"
	"go.uber.org/dig"
)

type GoMicroService struct {
	dig.In
	Server   contracts.MicroService
	Initials []contracts.Initial `group:"initials"`
}

func main() {
	providers := serviceGoMicro.
		Default(
			etcd.DefaultProvider(),
		).
		With(boot.Providers())

	opts := []atom.Option{
		atom.Name("dictionary"),
		atom.Version("1.0.2"),
		atom.RunE(func(cmd *cobra.Command, args []string) error {
			return container.Container.Invoke(func(svc GoMicroService) error {
				service := svc.Server.GetEngine().(micro.Service)
				service.Init()

				return Run(service)
			})
		}),
	}

	if err := atom.Serve(providers, opts...); err != nil {
		log.Fatal(err)
	}
}

func Run(m micro.Service) error {
	svc := v1.NewDictionaryService(atom.AppName, m.Client())
	for i := 0; i < 5; i++ {
		out, err := svc.GetItems(context.Background(), &v1.GetItemsRequest{
			Language: fmt.Sprintf("%d", i),
		}, client.WithSelectOption(
			selector.WithStrategy(selector.RoundRobin),
		))
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info(out)
	}

	return nil
}

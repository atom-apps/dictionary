//go:generate atomctl gen routes
//go:generate swag fmt
//go:generate swag init -ot json
package main

import (
	"github.com/atom-apps/dictionary/handlers"
	"github.com/atom-apps/dictionary/modules/boot"
	"github.com/atom-providers/etcd"
	"github.com/atom-providers/log"
	serviceGoMicro "github.com/atom-providers/service-gomicro"
	"github.com/rogeecn/atom"
)

func main() {
	providers := serviceGoMicro.
		Default(
			etcd.DefaultProvider(),
		).
		With(boot.Providers()).
		With(handlers.Providers())

	opts := []atom.Option{
		atom.Name("dictionary"),
		atom.Version("1.0.2"),
		atom.RunE(serviceGoMicro.ServeE),
	}

	if err := atom.Serve(providers, opts...); err != nil {
		log.Fatal(err)
	}
}

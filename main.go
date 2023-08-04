//go:generate atomctl gen routes
//go:generate swag fmt
//go:generate swag init -ot json
package main

import (
	"log"

	"github.com/atom-apps/dictionary/handlers"
	"github.com/atom-apps/dictionary/modules/boot"
	serviceGoMicro "github.com/atom-providers/service-gomicro"
	"github.com/rogeecn/atom"
	// serviceGrpc "github.com/atom-providers/service-grpc"
)

func main() {
	providers := serviceGoMicro.
		Default().
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

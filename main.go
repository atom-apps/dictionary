//go:generate atomctl gen routes
//go:generate swag fmt
//go:generate swag init -ot json
package main

import (
	"github.com/atom-apps/dictionary/database/query"
	"github.com/atom-apps/dictionary/handlers"
	"github.com/atom-apps/dictionary/modules/boot"
	modulesDictionary "github.com/atom-apps/dictionary/modules/dictionary"
	"github.com/atom-providers/casdoor"
	databasePostgres "github.com/atom-providers/database-postgres"
	"github.com/atom-providers/log"

	// serviceGoMicro "github.com/atom-providers/service-gomicro"

	serviceHttp "github.com/atom-providers/service-http"
	"github.com/rogeecn/atom"
)

func main() {
	providers := serviceHttp.
		Default(
			casdoor.DefaultProvider(),
			query.DefaultProvider(),
			databasePostgres.DefaultProvider(),
		).
		With(boot.Providers()).
		With(handlers.Providers()).
		With(modulesDictionary.Providers())

	opts := []atom.Option{
		atom.Name("dictionary"),
		atom.Version("1.0.3"),
		atom.RunE(serviceHttp.ServeE),
	}

	if err := atom.Serve(providers, opts...); err != nil {
		log.Fatal(err)
	}
}

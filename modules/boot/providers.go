package boot

import (
	"github.com/atom-providers/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/utils/opt"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: provideHttpMiddleware},
	}
}

func provideHttpMiddleware(opts ...opt.Option) error {
	return container.Container.Provide(func(
		httpsvc contracts.HttpService,
		jwt *jwt.JWT,
	) contracts.Initial {
		engine := httpsvc.GetEngine().(*fiber.App)

		engine.Use(httpMiddlewareJWT(jwt))
		return nil
	}, atom.GroupInitial)
}

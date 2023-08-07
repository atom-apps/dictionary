package boot

import (
	"github.com/atom-providers/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/utils/opt"
	// clientv3 "go.etcd.io/etcd/client/v3"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: provideHttpMiddleware},
		// {Provider: provideGoMicroOptions},
		// {Provider: provideSwagger},
	}
}

// func provideSwagger(opts ...opt.Option) error {
// 	return container.Container.Provide(func(swagger *swagger.Swagger) contracts.Initial {
// 		lo.Must0(swagger.Load(docs.SwaggerSpec))
// 		return nil
// 	}, atom.GroupInitial)
// }

// func provideGoMicroOptions(opts ...opt.Option) error {
// 	_ = container.Container.Provide(func(ctx context.Context, log *log.Logger, client *clientv3.Client) registry.Registry {
// 		logger, _ := zap.NewLogger(
// 			zap.WithLogger(log.Logger),
// 		)

// 		return etcd.NewRegistry(
// 			registry.Logger(logger),
// 			registry.Timeout(time.Second*5),
// 			etcd.Client(client),
// 		)
// 	})

// 	return nil
// }

func provideHttpMiddleware(opts ...opt.Option) error {
	return container.Container.Provide(
		func(httpsvc contracts.HttpService, jwt *jwt.JWT) contracts.Initial {
			engine := httpsvc.GetEngine().(*fiber.App)

			engine.Use(httpMiddlewareJWT(jwt))
			return nil
		}, atom.GroupInitial)
}

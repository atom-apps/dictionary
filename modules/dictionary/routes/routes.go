package routes

import (
	"github.com/atom-apps/dictionary/modules/dictionary/controller"
	"github.com/atom-providers/log"
	"github.com/gofiber/fiber/v2"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	return container.Container.Provide(newRoute, atom.GroupRoutes)
}

func newRoute(svc contracts.HttpService, dictionaryGroupController *controller.DictionaryGroupController, dictionaryGroupItemController *controller.DictionaryGroupItemController) contracts.HttpRoute {
	engine := svc.GetEngine().(*fiber.App)
	group := engine.Group("dictionaries")
	log.Infof("register route group: %s", group.(*fiber.Group).Prefix)

	routeDictionaryGroupItemController(group, dictionaryGroupItemController)
	routeDictionaryGroupController(group, dictionaryGroupController)
	return nil
}

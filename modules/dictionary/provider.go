package dictionary

import (
	"github.com/atom-apps/dictionary/modules/dictionary/controller"
	"github.com/atom-apps/dictionary/modules/dictionary/dao"
	"github.com/atom-apps/dictionary/modules/dictionary/routes"
	"github.com/atom-apps/dictionary/modules/dictionary/service"

	"github.com/rogeecn/atom/container"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: dao.Provide},
		{Provider: service.Provide},
		{Provider: controller.Provide},
		{Provider: routes.Provide},
	}
}

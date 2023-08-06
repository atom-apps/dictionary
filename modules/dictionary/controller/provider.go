package controller

import (
	"github.com/atom-apps/dictionary/modules/dictionary/service"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		dictionaryGroupItemSvc *service.DictionaryGroupItemService,
	) (*DictionaryGroupItemController, error) {
		obj := &DictionaryGroupItemController{
			dictionaryGroupItemSvc: dictionaryGroupItemSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		dictionaryGroupSvc *service.DictionaryGroupService,
	) (*DictionaryGroupController, error) {
		obj := &DictionaryGroupController{
			dictionaryGroupSvc: dictionaryGroupSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}

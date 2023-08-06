package service

import (
	"github.com/atom-apps/dictionary/modules/dictionary/dao"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		dictionaryGroupItemDao *dao.DictionaryGroupItemDao,
	) (*DictionaryGroupItemService, error) {
		obj := &DictionaryGroupItemService{
			dictionaryGroupItemDao: dictionaryGroupItemDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		dictionaryGroupDao *dao.DictionaryGroupDao,
		dictionaryGroupItemDao *dao.DictionaryGroupItemDao,
	) (*DictionaryGroupService, error) {
		obj := &DictionaryGroupService{
			dictionaryGroupDao:     dictionaryGroupDao,
			dictionaryGroupItemDao: dictionaryGroupItemDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}

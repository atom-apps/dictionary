package dao

import (
	"github.com/atom-apps/dictionary/database/query"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		query *query.Query,
	) (*DictionaryGroupItemDao, error) {
		obj := &DictionaryGroupItemDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		query *query.Query,
	) (*DictionaryGroupDao, error) {
		obj := &DictionaryGroupDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}

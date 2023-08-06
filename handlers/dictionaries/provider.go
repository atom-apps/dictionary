package dictionaries

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		micro contracts.MicroService,
	) (*DictionaryService, error) {
		obj := &DictionaryService{
			micro: micro,
		}
		if err := obj.Prepare(); err != nil {
			return nil, err
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}

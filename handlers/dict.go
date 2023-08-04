package handlers

import (
	"context"
	"log"

	v1 "github.com/atom-apps/dictionary/proto/v1"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/utils/opt"
	"github.com/samber/lo"
	"go-micro.dev/v4"
)

func dictProvider(opts ...opt.Option) error {
	return container.Container.Provide(func(svc contracts.MicroService) contracts.Initial {
		log.Println("register: DictionaryService")
		lo.Must0(v1.RegisterDictionaryServiceHandler(svc.GetEngine().(micro.Service).Server(), &DictionaryService{}))
		return nil
	}, atom.GroupInitial)
}

var _ v1.DictionaryServiceHandler = (*DictionaryService)(nil)

type DictionaryService struct{}

// GetItems implements v1.DictionaryServiceHandler.
func (*DictionaryService) GetItems(ctx context.Context, in *v1.GetItemsRequest, out *v1.GetItemsResponse) error {
	out.Items = []*v1.Item{
		{Key: "Key", Value: "Value"},
	}
	return nil
}

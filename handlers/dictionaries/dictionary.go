package dictionaries

import (
	"context"

	v1 "github.com/atom-apps/dictionary/proto/v1"
	"github.com/atom-providers/log"
	"github.com/rogeecn/atom/contracts"
	"go-micro.dev/v4"
)

var _ v1.DictionaryServiceHandler = (*DictionaryService)(nil)

// @provider
type DictionaryService struct {
	micro contracts.MicroService
	Name  string `inject:"false"`
}

func (svc *DictionaryService) Prepare() error {
	log.Info("register: DictionaryService")
	return v1.RegisterDictionaryServiceHandler(svc.micro.GetEngine().(micro.Service).Server(), svc)
}

// GetDictionaries implements v1.DictionaryServiceHandler.
func (*DictionaryService) GetDictionaries(context.Context, *v1.GetDictionariesRequest, *v1.GetDictionariesResponse) error {
	panic("unimplemented")
}

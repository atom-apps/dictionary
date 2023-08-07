package dictionaries

import (
	"context"

	"github.com/atom-apps/dictionary/modules/dictionary/dto"
	"github.com/atom-apps/dictionary/modules/dictionary/service"
	v1 "github.com/atom-apps/dictionary/proto/v1"
	"github.com/atom-providers/jwt"
	"github.com/atom-providers/log"
	"github.com/rogeecn/atom/contracts"
	"go-micro.dev/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ v1.DictionaryItemServiceHandler = (*DictionaryItemService)(nil)

//	@provider
type DictionaryItemService struct {
	*Common
	micro                  contracts.MicroService
	jwt                    *jwt.JWT
	dictionaryGroupSvc     *service.DictionaryGroupService
	dictionaryGroupItemSvc *service.DictionaryGroupItemService
	Name                   string `inject:"false"`
}

func (svc *DictionaryItemService) Prepare() error {
	log.Info("register: DictionaryItemService")
	return v1.RegisterDictionaryItemServiceHandler(svc.micro.GetEngine().(micro.Service).Server(), svc)
}

// Create implements v1.DictionaryItemServiceHandler.
func (svc *DictionaryItemService) Create(ctx context.Context, in *v1.DictionaryItemCreateRequest, out *emptypb.Empty) error {
	claim, err := svc.Claim(ctx, svc.jwt)
	if err != nil {
		return err
	}

	if claim.IsAdmin() {
		_, err = svc.dictionaryGroupSvc.GetByID(ctx, in.GroupId)
	} else if claim.IsTenantAdmin() {
		_, err = svc.dictionaryGroupSvc.GetByTenantID(ctx, claim.TenantID, in.GroupId)
	} else {
		_, err = svc.dictionaryGroupSvc.GetByUserID(ctx, claim.TenantID, claim.UserID, in.GroupId)
	}
	if err != nil {
		return err
	}

	return svc.dictionaryGroupItemSvc.Create(ctx, in.GroupId, &dto.DictionaryGroupItemForm{
		Value: in.Value,
		Order: &in.Order,
	})
}

// Delete implements v1.DictionaryItemServiceHandler.
func (svc *DictionaryItemService) Delete(ctx context.Context, in *v1.DictionaryItemDeleteRequest, out *emptypb.Empty) error {
	claim, err := svc.Claim(ctx, svc.jwt)
	if err != nil {
		return err
	}

	if claim.IsAdmin() {
		_, err = svc.dictionaryGroupSvc.GetByID(ctx, in.GroupId)
	} else if claim.IsTenantAdmin() {
		_, err = svc.dictionaryGroupSvc.GetByTenantID(ctx, claim.TenantID, in.GroupId)
	} else {
		_, err = svc.dictionaryGroupSvc.GetByUserID(ctx, claim.TenantID, claim.UserID, in.GroupId)
	}
	if err != nil {
		return err
	}

	return svc.dictionaryGroupItemSvc.Delete(ctx, in.Id)
}

// Show implements v1.DictionaryItemServiceHandler.
func (svc *DictionaryItemService) Show(ctx context.Context, in *v1.DictionaryItemShowRequest, out *v1.DictionaryItemShowResponse) error {
	claim, err := svc.Claim(ctx, svc.jwt)
	if err != nil {
		return err
	}

	if claim.IsAdmin() {
		_, err = svc.dictionaryGroupSvc.GetByID(ctx, in.GroupId)
	} else if claim.IsTenantAdmin() {
		_, err = svc.dictionaryGroupSvc.GetByTenantID(ctx, claim.TenantID, in.GroupId)
	} else {
		_, err = svc.dictionaryGroupSvc.GetByUserID(ctx, claim.TenantID, claim.UserID, in.GroupId)
	}
	if err != nil {
		return err
	}

	item, err := svc.dictionaryGroupItemSvc.GetByID(ctx, in.Id)
	if err != nil {
		return err
	}

	data := svc.dictionaryGroupItemSvc.DecorateItem(item, 0)

	out.GroupId = in.GroupId
	out.Id = data.ID
	out.Order = data.Order
	out.Value = data.Value

	return nil
}

// Update implements v1.DictionaryItemServiceHandler.
func (svc *DictionaryItemService) Update(ctx context.Context, in *v1.DictionaryItemUpdateRequest, out *emptypb.Empty) error {
	claim, err := svc.Claim(ctx, svc.jwt)
	if err != nil {
		return err
	}

	if claim.IsAdmin() {
		_, err = svc.dictionaryGroupSvc.GetByID(ctx, in.GroupId)
	} else if claim.IsTenantAdmin() {
		_, err = svc.dictionaryGroupSvc.GetByTenantID(ctx, claim.TenantID, in.GroupId)
	} else {
		_, err = svc.dictionaryGroupSvc.GetByUserID(ctx, claim.TenantID, claim.UserID, in.GroupId)
	}
	if err != nil {
		return err
	}

	body := &dto.DictionaryGroupItemForm{
		Value: in.Value,
		Order: &in.Order,
	}
	return svc.dictionaryGroupItemSvc.Update(ctx, in.GroupId, in.Id, body)
}

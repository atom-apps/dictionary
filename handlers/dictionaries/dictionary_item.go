package dictionaries

import (
	"context"

	authv1 "github.com/atom-apps/auth/proto/v1"
	"github.com/atom-apps/common/consts"
	"github.com/atom-apps/dictionary/modules/dictionary/dto"
	"github.com/atom-apps/dictionary/modules/dictionary/service"
	v1 "github.com/atom-apps/dictionary/proto/v1"
	"github.com/atom-providers/log"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/rogeecn/atom/contracts"
	"go-micro.dev/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ v1.DictionaryItemServiceHandler = (*DictionaryItemService)(nil)

// @provider
type DictionaryItemService struct {
	*Common

	micro                  contracts.MicroService
	door                   *casdoorsdk.Client
	dictionaryGroupSvc     *service.DictionaryGroupService
	dictionaryGroupItemSvc *service.DictionaryGroupItemService

	authSvc authv1.UserService `inject:"false"`
}

func (svc *DictionaryItemService) Prepare() error {
	log.Info("register: DictionaryItemService")

	microService := svc.micro.GetEngine().(micro.Service)

	svc.authSvc = authv1.NewUserService(consts.AppAuth.String(), microService.Client())

	return v1.RegisterDictionaryItemServiceHandler(microService.Server(), svc)
}

// Create implements v1.DictionaryItemServiceHandler.
func (svc *DictionaryItemService) Create(ctx context.Context, in *v1.DictionaryItemCreateRequest, out *emptypb.Empty) error {
	claim, err := svc.Claim(ctx, svc.door)
	if err != nil {
		return err
	}

	userMapping, err := svc.authSvc.GetMapping(ctx, &authv1.GetMappingRequest{
		Name: claim.Name,
	})
	if err != nil {
		return err
	}

	if claim.IsAdmin {
		_, err = svc.dictionaryGroupSvc.GetByID(ctx, in.GroupId)
	} else if userMapping.IsTenantAdmin {
		_, err = svc.dictionaryGroupSvc.GetByTenantID(ctx, userMapping.TenantId, in.GroupId)
	} else {
		_, err = svc.dictionaryGroupSvc.GetByUserID(ctx, userMapping.TenantId, userMapping.Id, in.GroupId)
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
	claim, err := svc.Claim(ctx, svc.door)
	if err != nil {
		return err
	}

	userMapping, err := svc.authSvc.GetMapping(ctx, &authv1.GetMappingRequest{
		Name: claim.Name,
	})
	if err != nil {
		return err
	}

	if claim.IsAdmin {
		_, err = svc.dictionaryGroupSvc.GetByID(ctx, in.GroupId)
	} else if userMapping.IsTenantAdmin {
		_, err = svc.dictionaryGroupSvc.GetByTenantID(ctx, userMapping.TenantId, in.GroupId)
	} else {
		_, err = svc.dictionaryGroupSvc.GetByUserID(ctx, userMapping.TenantId, userMapping.Id, in.GroupId)
	}
	if err != nil {
		return err
	}

	return svc.dictionaryGroupItemSvc.Delete(ctx, in.Id)
}

// Show implements v1.DictionaryItemServiceHandler.
func (svc *DictionaryItemService) Show(ctx context.Context, in *v1.DictionaryItemShowRequest, out *v1.DictionaryItemShowResponse) error {
	claim, err := svc.Claim(ctx, svc.door)
	if err != nil {
		return err
	}

	userMapping, err := svc.authSvc.GetMapping(ctx, &authv1.GetMappingRequest{
		Name: claim.Name,
	})
	if err != nil {
		return err
	}

	if claim.IsAdmin {
		_, err = svc.dictionaryGroupSvc.GetByID(ctx, in.GroupId)
	} else if userMapping.IsTenantAdmin {
		_, err = svc.dictionaryGroupSvc.GetByTenantID(ctx, userMapping.TenantId, in.GroupId)
	} else {
		_, err = svc.dictionaryGroupSvc.GetByUserID(ctx, userMapping.TenantId, userMapping.Id, in.GroupId)
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
	claim, err := svc.Claim(ctx, svc.door)
	if err != nil {
		return err
	}

	userMapping, err := svc.authSvc.GetMapping(ctx, &authv1.GetMappingRequest{
		Name: claim.Name,
	})
	if err != nil {
		return err
	}

	if claim.IsAdmin {
		_, err = svc.dictionaryGroupSvc.GetByID(ctx, in.GroupId)
	} else if userMapping.IsTenantAdmin {
		_, err = svc.dictionaryGroupSvc.GetByTenantID(ctx, userMapping.TenantId, in.GroupId)
	} else {
		_, err = svc.dictionaryGroupSvc.GetByUserID(ctx, userMapping.TenantId, userMapping.Id, in.GroupId)
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

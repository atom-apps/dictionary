package dictionaries

import (
	"context"

	"github.com/atom-apps/dictionary/common"
	"github.com/atom-apps/dictionary/database/models"
	"github.com/atom-apps/dictionary/modules/dictionary/dto"
	"github.com/atom-apps/dictionary/modules/dictionary/service"
	v1 "github.com/atom-apps/dictionary/proto/v1"
	"github.com/atom-providers/jwt"
	"github.com/atom-providers/log"
	"github.com/rogeecn/atom/contracts"
	"github.com/samber/lo"
	"go-micro.dev/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ v1.DictionaryServiceHandler = (*DictionaryService)(nil)

//	@provider
type DictionaryService struct {
	*Common
	micro              contracts.MicroService
	jwt                *jwt.JWT
	dictionaryGroupSvc *service.DictionaryGroupService
	Name               string `inject:"false"`
}

func (svc *DictionaryService) Prepare() error {
	log.Info("register: DictionaryService")
	return v1.RegisterDictionaryServiceHandler(svc.micro.GetEngine().(micro.Service).Server(), svc)
}

// Create implements v1.DictionaryServiceHandler.
func (svc *DictionaryService) Create(ctx context.Context, in *v1.DictionaryCreateRequest, out *emptypb.Empty) error {
	claim, err := svc.Claim(ctx, svc.jwt)
	if err != nil {
		return err
	}

	body := &dto.DictionaryGroupForm{}
	if err := body.FromProto(in); err != nil {
		return err
	}

	body.TenantID = claim.TenantID
	body.UserID = claim.UserID

	return svc.dictionaryGroupSvc.Create(ctx, body)
}

// Delete implements v1.DictionaryServiceHandler.
func (svc *DictionaryService) Delete(ctx context.Context, in *v1.IDRequest, out *emptypb.Empty) error {
	claim, err := svc.Claim(ctx, svc.jwt)
	if err != nil {
		return err
	}
	if claim.IsSuperAdmin() {
		return svc.dictionaryGroupSvc.Delete(ctx, claim.TenantID, in.Id)
	}

	return svc.dictionaryGroupSvc.DeleteByUserID(ctx, claim.TenantID, claim.UserID, in.Id)
}

// Paginate implements v1.DictionaryServiceHandler.
func (svc *DictionaryService) Paginate(ctx context.Context, in *v1.DictionaryPaginateRequest, out *v1.DictionaryPaginateResponse) error {
	claim, err := svc.Claim(ctx, svc.jwt)
	if err != nil {
		return err
	}

	queryFilter := &dto.DictionaryGroupListQueryFilter{}
	err = queryFilter.FromProto(in)
	if err != nil {
		return err
	}

	queryFilter.TenantID = claim.TenantID
	queryFilter.UserID = claim.UserID

	pageFilter := common.NewPageQueryFilter(in.Page, in.Limit)
	sortFilter := common.NewSortQueryFilter()

	items, total, err := svc.dictionaryGroupSvc.PageByQueryFilter(ctx, queryFilter, pageFilter, sortFilter)
	if err != nil {
		return err
	}

	outItems := lo.Map(items, svc.dictionaryGroupSvc.DecorateItem)

	out.Total = total
	out.Items = lo.Map(outItems, func(item *dto.DictionaryGroupItem, _ int) *v1.DictionaryItem {
		return item.ToProto()
	})

	return nil
}

// Share implements v1.DictionaryServiceHandler.
func (svc *DictionaryService) Share(ctx context.Context, in *v1.IDRequest, out *emptypb.Empty) error {
	claim, err := svc.Claim(ctx, svc.jwt)
	if err != nil {
		return err
	}

	return svc.dictionaryGroupSvc.ShareByID(ctx, claim.TenantID, claim.UserID, in.Id)
}

// Show implements v1.DictionaryServiceHandler.
func (svc *DictionaryService) Show(ctx context.Context, in *v1.IDRequest, out *v1.DictionaryItem) error {
	claim, err := svc.Claim(ctx, svc.jwt)
	if err != nil {
		return err
	}

	var item *models.DictionaryGroup
	if claim.IsSuperAdmin() {
		item, err = svc.dictionaryGroupSvc.GetByID(ctx, in.Id)
	} else if claim.IsTenantAdmin() {
		item, err = svc.dictionaryGroupSvc.GetByTenantID(ctx, claim.TenantID, in.Id)
	} else {
		item, err = svc.dictionaryGroupSvc.GetByUserID(ctx, claim.TenantID, claim.UserID, in.Id)
	}

	if err != nil {
		return err
	}

	out = svc.dictionaryGroupSvc.DecorateItem(item, 0).ToProto()
	return nil
}

// ShowBySlug implements v1.DictionaryServiceHandler.
func (svc *DictionaryService) ShowBySlug(ctx context.Context, in *v1.ShowBySlugRequest, out *v1.DictionaryItem) error {
	claim, err := svc.Claim(ctx, svc.jwt)
	if err != nil {
		return err
	}

	var item *models.DictionaryGroup
	if claim.IsSuperAdmin() {
		item, err = svc.dictionaryGroupSvc.GetFromSlug(ctx, in.Slug)
	} else if claim.IsTenantAdmin() {
		item, err = svc.dictionaryGroupSvc.GetFromSlugByTenantID(ctx, claim.TenantID, in.Slug)
	} else {
		item, err = svc.dictionaryGroupSvc.GetFromSlugByUserID(ctx, claim.TenantID, claim.UserID, in.Slug)
	}

	if err != nil {
		return err
	}

	out = svc.dictionaryGroupSvc.DecorateItem(item, 0).ToProto()
	return nil
}

// Update implements v1.DictionaryServiceHandler.
func (svc *DictionaryService) Update(ctx context.Context, in *v1.DictionaryUpdateRequest, out *emptypb.Empty) error {
	claim, err := svc.Claim(ctx, svc.jwt)
	if err != nil {
		return err
	}

	body := &dto.DictionaryGroupForm{
		Slug:        in.Slug,
		Name:        in.Name,
		Description: in.Description,
	}

	if claim.IsAdmin() {
		return svc.dictionaryGroupSvc.Update(ctx, in.Id, body)
	}
	return svc.dictionaryGroupSvc.UpdateByUserID(ctx, claim.TenantID, claim.UserID, in.Id, body)
}

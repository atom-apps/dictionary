package controller

import (
	authv1 "github.com/atom-apps/auth/proto/v1"
	"github.com/atom-apps/common/consts"
	"github.com/atom-apps/common/errorx"
	"github.com/atom-apps/dictionary/common"
	"github.com/atom-apps/dictionary/database/models"
	"github.com/atom-apps/dictionary/modules/dictionary/dto"
	"github.com/atom-apps/dictionary/modules/dictionary/service"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/rogeecn/atom/contracts"
	"go-micro.dev/v4"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type DictionaryGroupController struct {
	micro contracts.MicroService

	authSvc authv1.UserService `inject:"false"`

	dictionaryGroupSvc *service.DictionaryGroupService
}

func (c *DictionaryGroupController) Prepare() error {
	microService := c.micro.GetEngine().(micro.Service)

	c.authSvc = authv1.NewUserService(consts.AppAuth.String(), microService.Client())
	return nil
}

// Show get single item info
//
//	@Summary		get by id
//	@Description	get info by id
//	@Tags			Dictionary
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"DictionaryGroupID"
//	@Success		200	{object}	dto.DictionaryGroupItem
//	@Router			/dictionaries/{id} [get]
func (c *DictionaryGroupController) Show(ctx *fiber.Ctx, id int64) (*dto.DictionaryGroupItem, error) {
	claim, ok := ctx.Locals(consts.JwtCtx).(*casdoorsdk.Claims)
	if !ok {
		return nil, errorx.ErrTokenInvalid
	}

	userMapping, err := c.authSvc.GetMapping(ctx.Context(), &authv1.GetMappingRequest{
		Name: claim.Name,
	})
	if err != nil {
		return nil, err
	}

	var item *models.DictionaryGroup
	if claim.IsAdmin {
		item, err = c.dictionaryGroupSvc.GetByID(ctx.Context(), id)
	} else if userMapping.IsTenantAdmin {
		item, err = c.dictionaryGroupSvc.GetByTenantID(ctx.Context(), userMapping.TenantId, id)
	} else {
		item, err = c.dictionaryGroupSvc.GetByUserID(ctx.Context(), userMapping.TenantId, userMapping.Id, id)
	}

	if err != nil {
		return nil, err
	}

	return c.dictionaryGroupSvc.DecorateItem(item, 0), nil
}

// Show by group slug
//
//	@Summary		get by slug
//	@Description	get info by slug
//	@Tags			Dictionary
//	@Accept			json
//	@Produce		json
//	@Param			slug	path		string	true	"DictionaryGroupSlug"
//	@Success		200		{object}	dto.DictionaryGroupItem
//	@Router			/dictionaries/slug/{slug} [get]
func (c *DictionaryGroupController) ShowBySlug(ctx *fiber.Ctx, slug string) (*dto.DictionaryGroupItem, error) {
	claim, ok := ctx.Locals(consts.JwtCtx).(*casdoorsdk.Claims)
	if !ok {
		return nil, errorx.ErrTokenInvalid
	}

	userMapping, err := c.authSvc.GetMapping(ctx.Context(), &authv1.GetMappingRequest{
		Name: claim.Name,
	})
	if err != nil {
		return nil, err
	}

	var item *models.DictionaryGroup
	if claim.IsAdmin {
		item, err = c.dictionaryGroupSvc.GetFromSlug(ctx.Context(), slug)
	} else if userMapping.IsTenantAdmin {
		item, err = c.dictionaryGroupSvc.GetFromSlugByTenantID(ctx.Context(), userMapping.TenantId, slug)
	} else {
		item, err = c.dictionaryGroupSvc.GetFromSlugByUserID(ctx.Context(), userMapping.TenantId, userMapping.Id, slug)
	}

	if err != nil {
		return nil, err
	}
	return c.dictionaryGroupSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		list by query filter
//	@Description	list by query filter
//	@Tags			Dictionary
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.DictionaryGroupListQueryFilter	true	"DictionaryGroupListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter				true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter				true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.DictionaryGroupItem}
//	@Router			/dictionaries [get]
func (c *DictionaryGroupController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.DictionaryGroupListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	claim, ok := ctx.Locals(consts.JwtCtx).(*casdoorsdk.Claims)
	if !ok {
		return nil, errorx.ErrTokenInvalid
	}

	userMapping, err := c.authSvc.GetMapping(ctx.Context(), &authv1.GetMappingRequest{
		Name: claim.Name,
	})
	if err != nil {
		return nil, err
	}

	queryFilter.TenantID = userMapping.TenantId
	queryFilter.UserID = userMapping.Id

	items, total, err := c.dictionaryGroupSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.dictionaryGroupSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		create new item
//	@Description	create new item
//	@Tags			Dictionary
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.DictionaryGroupForm	true	"DictionaryGroupForm"
//	@Success		200		{string}	DictionaryGroupID
//	@Router			/dictionaries [post]
func (c *DictionaryGroupController) Create(ctx *fiber.Ctx, body *dto.DictionaryGroupForm) error {
	claim, ok := ctx.Locals(consts.JwtCtx).(*casdoorsdk.Claims)
	if !ok {
		return errorx.ErrTokenInvalid
	}

	userMapping, err := c.authSvc.GetMapping(ctx.Context(), &authv1.GetMappingRequest{
		Name: claim.Name,
	})
	if err != nil {
		return err
	}

	body.TenantID = userMapping.TenantId
	body.UserID = userMapping.Id

	return c.dictionaryGroupSvc.Create(ctx.Context(), body)
}

// Update update by id
//
//	@Summary		update by id
//	@Description	update by id
//	@Tags			Dictionary
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"DictionaryGroupID"
//	@Param			body	body		dto.DictionaryGroupForm	true	"DictionaryGroupForm"
//	@Success		200		{string}	DictionaryGroupID
//	@Failure		500		{string}	DictionaryGroupID
//	@Router			/dictionaries/{id} [put]
func (c *DictionaryGroupController) Update(ctx *fiber.Ctx, id int64, body *dto.DictionaryGroupForm) error {
	claim, ok := ctx.Locals(consts.JwtCtx).(*casdoorsdk.Claims)
	if !ok {
		return errorx.ErrTokenInvalid
	}

	userMapping, err := c.authSvc.GetMapping(ctx.Context(), &authv1.GetMappingRequest{
		Name: claim.Name,
	})
	if err != nil {
		return err
	}

	if claim.IsAdmin {
		return c.dictionaryGroupSvc.Update(ctx.Context(), id, body)
	}
	return c.dictionaryGroupSvc.UpdateByUserID(ctx.Context(), userMapping.TenantId, userMapping.Id, id, body)
}

// Delete delete by id
//
//	@Summary		delete by id
//	@Description	delete by id
//	@Tags			Dictionary
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"DictionaryGroupID"
//	@Success		200	{string}	DictionaryGroupID
//	@Failure		500	{string}	DictionaryGroupID
//	@Router			/dictionaries/{id} [delete]
func (c *DictionaryGroupController) Delete(ctx *fiber.Ctx, id int64) error {
	claim, ok := ctx.Locals(consts.JwtCtx).(*casdoorsdk.Claims)
	if !ok {
		return errorx.ErrTokenInvalid
	}

	userMapping, err := c.authSvc.GetMapping(ctx.Context(), &authv1.GetMappingRequest{
		Name: claim.Name,
	})
	if err != nil {
		return err
	}

	if claim.IsAdmin {
		return c.dictionaryGroupSvc.Delete(ctx.Context(), userMapping.TenantId, id)
	}

	return c.dictionaryGroupSvc.DeleteByUserID(ctx.Context(), userMapping.TenantId, userMapping.Id, id)
}

// Share share by id
//
//	@Summary		share by user id
//	@Description	share by user id
//	@Tags			Dictionary
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"DictionaryGroupID"
//	@Success		200	{string}	DictionaryGroupID
//	@Failure		500	{string}	DictionaryGroupID
//	@Router			/dictionaries/{id}/share [put]
func (c *DictionaryGroupController) Share(ctx *fiber.Ctx, id int64) error {
	claim, ok := ctx.Locals(consts.JwtCtx).(*casdoorsdk.Claims)
	if !ok {
		return errorx.ErrTokenInvalid
	}

	userMapping, err := c.authSvc.GetMapping(ctx.Context(), &authv1.GetMappingRequest{
		Name: claim.Name,
	})
	if err != nil {
		return err
	}

	return c.dictionaryGroupSvc.ShareByID(ctx.Context(), userMapping.TenantId, userMapping.Id, id)
}

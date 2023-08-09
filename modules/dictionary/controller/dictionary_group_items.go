package controller

import (
	authv1 "github.com/atom-apps/auth/proto/v1"
	"github.com/atom-apps/common/consts"
	"github.com/atom-apps/common/errorx"
	"github.com/atom-apps/dictionary/modules/dictionary/dto"
	"github.com/atom-apps/dictionary/modules/dictionary/service"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/rogeecn/atom/contracts"
	"go-micro.dev/v4"

	"github.com/gofiber/fiber/v2"
)

// @provider
type DictionaryGroupItemController struct {
	micro contracts.MicroService

	authSvc authv1.UserService `inject:"false"`

	dictionaryGroupSvc     *service.DictionaryGroupService
	dictionaryGroupItemSvc *service.DictionaryGroupItemService
}

func (c *DictionaryGroupItemController) Prepare() error {
	microService := c.micro.GetEngine().(micro.Service)

	c.authSvc = authv1.NewUserService(consts.AppAuth.String(), microService.Client())
	return nil
}

// Show get single item info
//
//	@Summary		get by id
//	@Description	get info by id
//	@Tags			DictionaryGroupItem
//	@Accept			json
//	@Produce		json
//	@Param			dictionaryId	path		int	true	"DictionaryId"
//	@Param			id				path		int	true	"DictionaryGroupItemID"
//	@Success		200				{object}	dto.DictionaryGroupItemItem
//	@Router			/dictionaries/{dictionary_id}/items/{id} [get]
func (c *DictionaryGroupItemController) Show(ctx *fiber.Ctx, dictionaryId, id int64) (*dto.DictionaryGroupItemItem, error) {
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

	if claim.IsAdmin {
		_, err = c.dictionaryGroupSvc.GetByID(ctx.Context(), dictionaryId)
	} else if userMapping.IsTenantAdmin {
		_, err = c.dictionaryGroupSvc.GetByTenantID(ctx.Context(), userMapping.TenantId, dictionaryId)
	} else {
		_, err = c.dictionaryGroupSvc.GetByUserID(ctx.Context(), userMapping.TenantId, userMapping.Id, dictionaryId)
	}
	if err != nil {
		return nil, err
	}

	item, err := c.dictionaryGroupItemSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.dictionaryGroupItemSvc.DecorateItem(item, 0), nil
}

// Create a new item
//
//	@Summary		create new item
//	@Description	create new item
//	@Tags			DictionaryGroupItem
//	@Accept			json
//	@Produce		json
//	@Param			dictionaryId	path		int							true	"DictionaryId"
//	@Param			body			body		dto.DictionaryGroupItemForm	true	"DictionaryGroupItemForm"
//	@Success		200				{string}	DictionaryGroupItemID
//	@Router			/dictionaries/{dictionary_id}/items [post]
func (c *DictionaryGroupItemController) Create(ctx *fiber.Ctx, dictionaryId int64, body *dto.DictionaryGroupItemForm) error {
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
		_, err = c.dictionaryGroupSvc.GetByID(ctx.Context(), dictionaryId)
	} else if userMapping.IsTenantAdmin {
		_, err = c.dictionaryGroupSvc.GetByTenantID(ctx.Context(), userMapping.TenantId, dictionaryId)
	} else {
		_, err = c.dictionaryGroupSvc.GetByUserID(ctx.Context(), userMapping.TenantId, userMapping.Id, dictionaryId)
	}
	if err != nil {
		return err
	}

	return c.dictionaryGroupItemSvc.Create(ctx.Context(), dictionaryId, body)
}

// Update update by id
//
//	@Summary		update by id
//	@Description	update by id
//	@Tags			DictionaryGroupItem
//	@Accept			json
//	@Produce		json
//	@Param			dictionaryId	path		int							true	"DictionaryId"
//	@Param			id				path		int							true	"DictionaryGroupItemID"
//	@Param			body			body		dto.DictionaryGroupItemForm	true	"DictionaryGroupItemForm"
//	@Success		200				{string}	DictionaryGroupItemID
//	@Failure		500				{string}	DictionaryGroupItemID
//	@Router			/dictionaries/{dictionary_id}/items/{id} [put]
func (c *DictionaryGroupItemController) Update(ctx *fiber.Ctx, dictionaryId, id int64, body *dto.DictionaryGroupItemForm) error {
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
		_, err = c.dictionaryGroupSvc.GetByID(ctx.Context(), dictionaryId)
	} else if userMapping.IsTenantAdmin {
		_, err = c.dictionaryGroupSvc.GetByTenantID(ctx.Context(), userMapping.TenantId, dictionaryId)
	} else {
		_, err = c.dictionaryGroupSvc.GetByUserID(ctx.Context(), userMapping.TenantId, userMapping.Id, dictionaryId)
	}
	if err != nil {
		return err
	}

	return c.dictionaryGroupItemSvc.Update(ctx.Context(), dictionaryId, id, body)
}

// Delete delete by id
//
//	@Summary		delete by id
//	@Description	delete by id
//	@Tags			DictionaryGroupItem
//	@Accept			json
//	@Produce		json
//	@Param			dictionaryId	path		int	true	"DictionaryId"
//	@Param			id				path		int	true	"DictionaryGroupItemID"
//	@Success		200				{string}	DictionaryGroupItemID
//	@Failure		500				{string}	DictionaryGroupItemID
//	@Router			/dictionaries/{dictionary_id}/items/{id} [delete]
func (c *DictionaryGroupItemController) Delete(ctx *fiber.Ctx, dictionaryId, id int64) error {
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
		_, err = c.dictionaryGroupSvc.GetByID(ctx.Context(), dictionaryId)
	} else if userMapping.IsTenantAdmin {
		_, err = c.dictionaryGroupSvc.GetByTenantID(ctx.Context(), userMapping.TenantId, dictionaryId)
	} else {
		_, err = c.dictionaryGroupSvc.GetByUserID(ctx.Context(), userMapping.TenantId, userMapping.Id, dictionaryId)
	}
	if err != nil {
		return err
	}

	return c.dictionaryGroupItemSvc.Delete(ctx.Context(), id)
}

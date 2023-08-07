package controller

import (
	"github.com/atom-apps/dictionary/common"
	"github.com/atom-apps/dictionary/modules/dictionary/dto"
	"github.com/atom-apps/dictionary/modules/dictionary/service"
	"github.com/atom-providers/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type DictionaryGroupController struct {
	dictionaryGroupSvc *service.DictionaryGroupService
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
	claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
	if !ok {
		return nil, jwt.TokenInvalid
	}

	item, err := c.dictionaryGroupSvc.GetByTenantID(ctx.Context(), claim.TenantID, id)
	if err != nil {
		return nil, err
	}
	return c.dictionaryGroupSvc.DecorateItem(item, 0), nil
}

// Show by user id
//
//	@Summary		get by user id
//	@Description	get info by user id
//	@Tags			Dictionary
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"DictionaryGroupID"
//	@Success		200	{object}	dto.DictionaryGroupItem
//	@Router			/dictionaries/{id}/user [get]
func (c *DictionaryGroupController) ShowByUserID(ctx *fiber.Ctx, id int64) (*dto.DictionaryGroupItem, error) {
	claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
	if !ok {
		return nil, jwt.TokenInvalid
	}

	item, err := c.dictionaryGroupSvc.GetByUserID(ctx.Context(), claim.TenantID, claim.UserID, id)
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
	claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
	if !ok {
		return nil, jwt.TokenInvalid
	}
	queryFilter.TenantID = claim.TenantID
	queryFilter.UserID = claim.UserID

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
	claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
	if !ok {
		return jwt.TokenInvalid
	}
	body.TenantID = claim.TenantID
	body.UserID = claim.UserID
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
	claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
	if !ok {
		return jwt.TokenInvalid
	}
	return c.dictionaryGroupSvc.Update(ctx.Context(), claim.TenantID, id, body)
}

// Update update by user id
//
//	@Summary		update by user id
//	@Description	update by user id
//	@Tags			Dictionary
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"DictionaryGroupID"
//	@Param			body	body		dto.DictionaryGroupForm	true	"DictionaryGroupForm"
//	@Success		200		{string}	DictionaryGroupID
//	@Failure		500		{string}	DictionaryGroupID
//	@Router			/dictionaries/{id}/user [put]
func (c *DictionaryGroupController) UpdateByUserID(ctx *fiber.Ctx, id int64, body *dto.DictionaryGroupForm) error {
	claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
	if !ok {
		return jwt.TokenInvalid
	}
	return c.dictionaryGroupSvc.UpdateByUserID(ctx.Context(), claim.TenantID, claim.UserID, id, body)
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
	claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
	if !ok {
		return jwt.TokenInvalid
	}
	return c.dictionaryGroupSvc.Delete(ctx.Context(), claim.TenantID, id)
}

// Delete delete by user id
//
//	@Summary		delete by user id
//	@Description	delete by user id
//	@Tags			Dictionary
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"DictionaryGroupID"
//	@Success		200	{string}	DictionaryGroupID
//	@Failure		500	{string}	DictionaryGroupID
//	@Router			/dictionaries/{id}/user [delete]
func (c *DictionaryGroupController) DeleteByUserID(ctx *fiber.Ctx, id int64) error {
	claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
	if !ok {
		return jwt.TokenInvalid
	}
	return c.dictionaryGroupSvc.DeleteByUserID(ctx.Context(), claim.TenantID, claim.UserID, id)
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
	claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
	if !ok {
		return jwt.TokenInvalid
	}
	return c.dictionaryGroupSvc.ShareByID(ctx.Context(), claim.TenantID, claim.UserID, id)
}

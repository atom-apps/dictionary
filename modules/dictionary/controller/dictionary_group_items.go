package controller

import (
	"github.com/atom-apps/dictionary/modules/dictionary/dto"
	"github.com/atom-apps/dictionary/modules/dictionary/service"
	"github.com/atom-providers/jwt"

	"github.com/gofiber/fiber/v2"
)

// @provider
type DictionaryGroupItemController struct {
	dictionaryGroupSvc     *service.DictionaryGroupService
	dictionaryGroupItemSvc *service.DictionaryGroupItemService
}

// Show get single item info
//
//	@Summary		get by id
//	@Description	get info by id
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			dictionaryId	path		int	true	"DictionaryId"
//	@Param			id	path		int	true	"DictionaryGroupItemID"
//	@Success		200	{object}	dto.DictionaryGroupItemItem
//	@Router			/dictionaries/{dictionary_id}/items/{id} [get]
func (c *DictionaryGroupItemController) Show(ctx *fiber.Ctx, dictionaryId, id int64) (*dto.DictionaryGroupItemItem, error) {
	claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
	if !ok {
		return nil, jwt.TokenInvalid
	}
	_, err := c.dictionaryGroupSvc.GetByUserID(ctx.Context(), claim.TenantID, claim.UserID, dictionaryId)
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
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			dictionaryId	path		int	true	"DictionaryId"
//	@Param			body	body		dto.DictionaryGroupItemForm	true	"DictionaryGroupItemForm"
//	@Success		200		{string}	DictionaryGroupItemID
//	@Router			/dictionaries/{dictionary_id}/items [post]
func (c *DictionaryGroupItemController) Create(ctx *fiber.Ctx, dictionaryId int64, body *dto.DictionaryGroupItemForm) error {
	claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
	if !ok {
		return jwt.TokenInvalid
	}
	_, err := c.dictionaryGroupSvc.GetByUserID(ctx.Context(), claim.TenantID, claim.UserID, dictionaryId)
	if err != nil {
		return err
	}

	return c.dictionaryGroupItemSvc.Create(ctx.Context(), dictionaryId, body)
}

// Update update by id
//
//	@Summary		update by id
//	@Description	update by id
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			dictionaryId	path		int	true	"DictionaryId"
//	@Param			id		path		int				true	"DictionaryGroupItemID"
//	@Param			body	body		dto.DictionaryGroupItemForm	true	"DictionaryGroupItemForm"
//	@Success		200		{string}	DictionaryGroupItemID
//	@Failure		500		{string}	DictionaryGroupItemID
//	@Router			/dictionaries/{dictionary_id}/items/{id} [put]
func (c *DictionaryGroupItemController) Update(ctx *fiber.Ctx, dictionaryId, id int64, body *dto.DictionaryGroupItemForm) error {
	claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
	if !ok {
		return jwt.TokenInvalid
	}
	_, err := c.dictionaryGroupSvc.GetByUserID(ctx.Context(), claim.TenantID, claim.UserID, dictionaryId)
	if err != nil {
		return err
	}

	return c.dictionaryGroupItemSvc.Update(ctx.Context(), dictionaryId, id, body)
}

// Delete delete by id
//
//	@Summary		delete by id
//	@Description	delete by id
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			dictionaryId	path		int	true	"DictionaryId"
//	@Param			id	path		int	true	"DictionaryGroupItemID"
//	@Success		200	{string}	DictionaryGroupItemID
//	@Failure		500	{string}	DictionaryGroupItemID
//	@Router			/dictionaries/{dictionary_id}/items/{id} [delete]
func (c *DictionaryGroupItemController) Delete(ctx *fiber.Ctx, dictionaryId, id int64) error {
	claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
	if !ok {
		return jwt.TokenInvalid
	}
	_, err := c.dictionaryGroupSvc.GetByUserID(ctx.Context(), claim.TenantID, claim.UserID, dictionaryId)
	if err != nil {
		return err
	}

	return c.dictionaryGroupItemSvc.Delete(ctx.Context(), id)
}

package service

import (
	"context"

	"github.com/atom-apps/common/errorx"
	"github.com/atom-apps/dictionary/common"
	"github.com/atom-apps/dictionary/database/models"
	"github.com/atom-apps/dictionary/modules/dictionary/dao"
	"github.com/atom-apps/dictionary/modules/dictionary/dto"

	"github.com/jinzhu/copier"
)

// @provider
type DictionaryGroupItemService struct {
	dictionaryGroupItemDao *dao.DictionaryGroupItemDao
}

func (svc *DictionaryGroupItemService) DecorateItem(model *models.DictionaryGroupItem, id int) *dto.DictionaryGroupItemItem {
	var dtoItem *dto.DictionaryGroupItemItem
	_ = copier.Copy(dtoItem, model)

	return dtoItem
}

func (svc *DictionaryGroupItemService) GetByID(ctx context.Context, id int64) (*models.DictionaryGroupItem, error) {
	return svc.dictionaryGroupItemDao.GetByID(ctx, id)
}

func (svc *DictionaryGroupItemService) GetByValue(ctx context.Context, dictionaryId int64, value string) (*models.DictionaryGroupItem, error) {
	return svc.dictionaryGroupItemDao.GetByValue(ctx, dictionaryId, value)
}

func (svc *DictionaryGroupItemService) FindByQueryFilter(
	ctx context.Context,
	dictionaryId int,
	queryFilter *dto.DictionaryGroupItemListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.DictionaryGroupItem, error) {
	return svc.dictionaryGroupItemDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

// CreateFromModel
func (svc *DictionaryGroupItemService) CreateFromModel(ctx context.Context, model *models.DictionaryGroupItem) error {
	return svc.dictionaryGroupItemDao.Create(ctx, model)
}

// Create
func (svc *DictionaryGroupItemService) Create(ctx context.Context, dictionaryId int64, body *dto.DictionaryGroupItemForm) error {
	m, _ := svc.GetByValue(ctx, dictionaryId, body.Value)
	if m != nil {
		return errorx.ErrRecordAlreadyExists
	}

	model := &models.DictionaryGroupItem{
		DictionaryGroupID: dictionaryId,
		Value:             body.Value,
		Order:             0,
	}

	if body.Order != nil {
		model.Order = *body.Order
	}
	return svc.dictionaryGroupItemDao.Create(ctx, model)
}

// Update
func (svc *DictionaryGroupItemService) Update(ctx context.Context, dictionaryId, id int64, body *dto.DictionaryGroupItemForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.dictionaryGroupItemDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *DictionaryGroupItemService) UpdateFromModel(ctx context.Context, model *models.DictionaryGroupItem) error {
	return svc.dictionaryGroupItemDao.Update(ctx, model)
}

// Delete
func (svc *DictionaryGroupItemService) Delete(ctx context.Context, id int64) error {
	return svc.dictionaryGroupItemDao.Delete(ctx, id)
}

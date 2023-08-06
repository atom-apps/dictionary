package service

import (
	"context"

	"github.com/atom-apps/dictionary/common"
	"github.com/atom-apps/dictionary/database/models"
	"github.com/atom-apps/dictionary/modules/dictionary/dao"
	"github.com/atom-apps/dictionary/modules/dictionary/dto"
	"github.com/samber/lo"

	"github.com/jinzhu/copier"
)

// @provider
type DictionaryGroupService struct {
	dictionaryGroupDao     *dao.DictionaryGroupDao
	dictionaryGroupItemDao *dao.DictionaryGroupItemDao
}

func (svc *DictionaryGroupService) DecorateItem(model *models.DictionaryGroup, id int) *dto.DictionaryGroupItem {
	dtoItem := &dto.DictionaryGroupItem{
		ID:          model.ID,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		UserID:      model.UserID,
		TenantID:    model.TenantID,
		Name:        model.Name,
		Slug:        model.Slug,
		Description: model.Description,
		Items:       []dto.KeyValue{},
	}

	items, err := svc.dictionaryGroupItemDao.FindByGroupID(context.Background(), model.ID)
	if err == nil {
		dtoItem.Items = lo.Map(items, func(item *models.DictionaryGroupItem, _ int) dto.KeyValue {
			return dto.KeyValue{
				Key:   item.Key,
				Value: item.Value,
			}
		})
	}

	return dtoItem
}

func (svc *DictionaryGroupService) GetByTenantID(ctx context.Context, tenantID, id int64) (*models.DictionaryGroup, error) {
	return svc.dictionaryGroupDao.GetByTenantID(ctx, tenantID, id)
}

func (svc *DictionaryGroupService) GetByUserID(ctx context.Context, tenantID, userID, id int64) (*models.DictionaryGroup, error) {
	return svc.dictionaryGroupDao.GetByUserID(ctx, tenantID, userID, id)
}

func (svc *DictionaryGroupService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DictionaryGroupListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.DictionaryGroup, error) {
	return svc.dictionaryGroupDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *DictionaryGroupService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DictionaryGroupListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.DictionaryGroup, int64, error) {
	return svc.dictionaryGroupDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *DictionaryGroupService) CreateFromModel(ctx context.Context, model *models.DictionaryGroup) error {
	return svc.dictionaryGroupDao.Create(ctx, model)
}

// Create
func (svc *DictionaryGroupService) Create(ctx context.Context, body *dto.DictionaryGroupForm) error {
	model := &models.DictionaryGroup{}
	_ = copier.Copy(model, body)
	return svc.dictionaryGroupDao.Create(ctx, model)
}

// Update
func (svc *DictionaryGroupService) Update(ctx context.Context, tenantID, id int64, body *dto.DictionaryGroupForm) error {
	model, err := svc.GetByTenantID(ctx, tenantID, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.dictionaryGroupDao.Update(ctx, model)
}

func (svc *DictionaryGroupService) UpdateByUserID(ctx context.Context, tenantID, userID, id int64, body *dto.DictionaryGroupForm) error {
	model, err := svc.GetByUserID(ctx, tenantID, userID, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.dictionaryGroupDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *DictionaryGroupService) UpdateFromModel(ctx context.Context, model *models.DictionaryGroup) error {
	return svc.dictionaryGroupDao.Update(ctx, model)
}

// Delete
func (svc *DictionaryGroupService) Delete(ctx context.Context, tenantID, id int64) error {
	return svc.dictionaryGroupDao.Delete(ctx, tenantID, id)
}

// Delete
func (svc *DictionaryGroupService) DeleteByUserID(ctx context.Context, tenantID, userID, id int64) error {
	return svc.dictionaryGroupDao.DeleteByUserID(ctx, tenantID, userID, id)
}

// Delete
func (svc *DictionaryGroupService) ShareByID(ctx context.Context, tenantID, userID, id int64) error {
	model, err := svc.GetByUserID(ctx, tenantID, userID, id)
	if err != nil {
		return err
	}

	model.UserID = 0
	return svc.dictionaryGroupDao.Update(ctx, model)
}

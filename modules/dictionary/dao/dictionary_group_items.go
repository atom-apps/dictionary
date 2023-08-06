package dao

import (
	"context"

	"github.com/atom-apps/dictionary/common"
	"github.com/atom-apps/dictionary/database/models"
	"github.com/atom-apps/dictionary/database/query"
	"github.com/atom-apps/dictionary/modules/dictionary/dto"

	"gorm.io/gen/field"
)

// @provider
type DictionaryGroupItemDao struct {
	query *query.Query
}

func (dao *DictionaryGroupItemDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *DictionaryGroupItemDao) Context(ctx context.Context) query.IDictionaryGroupItemDo {
	return dao.query.DictionaryGroupItem.WithContext(ctx)
}

func (dao *DictionaryGroupItemDao) decorateSortQueryFilter(query query.IDictionaryGroupItemDo, sortFilter *common.SortQueryFilter) query.IDictionaryGroupItemDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.DictionaryGroupItem.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.DictionaryGroupItem.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *DictionaryGroupItemDao) decorateQueryFilter(query query.IDictionaryGroupItemDo, queryFilter *dto.DictionaryGroupItemListQueryFilter) query.IDictionaryGroupItemDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.DictionaryGroupID != nil {
		query = query.Where(dao.query.DictionaryGroupItem.DictionaryGroupID.Eq(*queryFilter.DictionaryGroupID))
	}
	if queryFilter.Value != nil {
		query = query.Where(dao.query.DictionaryGroupItem.Value.Eq(*queryFilter.Value))
	}
	if queryFilter.Order != nil {
		query = query.Where(dao.query.DictionaryGroupItem.Order.Eq(*queryFilter.Order))
	}

	return query
}

func (dao *DictionaryGroupItemDao) UpdateColumn(ctx context.Context, id int64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.DictionaryGroupItem.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *DictionaryGroupItemDao) Update(ctx context.Context, model *models.DictionaryGroupItem) error {
	_, err := dao.Context(ctx).Where(dao.query.DictionaryGroupItem.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *DictionaryGroupItemDao) Delete(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Where(dao.query.DictionaryGroupItem.ID.Eq(id)).Delete()
	return err
}

func (dao *DictionaryGroupItemDao) DeletePermanently(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.DictionaryGroupItem.ID.Eq(id)).Delete()
	return err
}

func (dao *DictionaryGroupItemDao) Restore(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.DictionaryGroupItem.ID.Eq(id)).UpdateSimple(dao.query.DictionaryGroupItem.DeletedAt.Null())
	return err
}

func (dao *DictionaryGroupItemDao) Create(ctx context.Context, model *models.DictionaryGroupItem) error {
	return dao.Context(ctx).Create(model)
}

func (dao *DictionaryGroupItemDao) GetByID(ctx context.Context, id int64) (*models.DictionaryGroupItem, error) {
	return dao.Context(ctx).Where(dao.query.DictionaryGroupItem.ID.Eq(id)).First()
}

func (dao *DictionaryGroupItemDao) GetByValue(ctx context.Context, dictionaryID int64, value string) (*models.DictionaryGroupItem, error) {
	return dao.Context(ctx).Where(
		dao.query.DictionaryGroupItem.DictionaryGroupID.Eq(dictionaryID),
		dao.query.DictionaryGroupItem.Value.Eq(value),
	).First()
}

func (dao *DictionaryGroupItemDao) GetByIDs(ctx context.Context, ids []int64) ([]*models.DictionaryGroupItem, error) {
	return dao.Context(ctx).Where(dao.query.DictionaryGroupItem.ID.In(ids...)).Find()
}

func (dao *DictionaryGroupItemDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DictionaryGroupItemListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.DictionaryGroupItem, int64, error) {
	query := dao.query.DictionaryGroupItem
	dictionaryGroupItemQuery := query.WithContext(ctx)
	dictionaryGroupItemQuery = dao.decorateQueryFilter(dictionaryGroupItemQuery, queryFilter)
	dictionaryGroupItemQuery = dao.decorateSortQueryFilter(dictionaryGroupItemQuery, sortFilter)
	return dictionaryGroupItemQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *DictionaryGroupItemDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DictionaryGroupItemListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.DictionaryGroupItem, error) {
	query := dao.query.DictionaryGroupItem
	dictionaryGroupItemQuery := query.WithContext(ctx)
	dictionaryGroupItemQuery = dao.decorateQueryFilter(dictionaryGroupItemQuery, queryFilter)
	dictionaryGroupItemQuery = dao.decorateSortQueryFilter(dictionaryGroupItemQuery, sortFilter)
	return dictionaryGroupItemQuery.Find()
}

func (dao *DictionaryGroupItemDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DictionaryGroupItemListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.DictionaryGroupItem, error) {
	query := dao.query.DictionaryGroupItem
	dictionaryGroupItemQuery := query.WithContext(ctx)
	dictionaryGroupItemQuery = dao.decorateQueryFilter(dictionaryGroupItemQuery, queryFilter)
	dictionaryGroupItemQuery = dao.decorateSortQueryFilter(dictionaryGroupItemQuery, sortFilter)
	return dictionaryGroupItemQuery.First()
}

func (dao *DictionaryGroupItemDao) FindByGroupID(ctx context.Context, groupID int64) ([]*models.DictionaryGroupItem, error) {
	query := dao.query.DictionaryGroupItem
	dictionaryGroupItemQuery := query.WithContext(ctx)
	return dictionaryGroupItemQuery.Where(query.DictionaryGroupID.Eq(groupID)).Order(query.Order.Desc()).Find()
}

package dao

import (
	"context"

	"github.com/atom-apps/dictionary/common"
	"github.com/atom-apps/dictionary/database/models"
	"github.com/atom-apps/dictionary/database/query"
	"github.com/atom-apps/dictionary/modules/dictionary/dto"

	"gorm.io/gen/field"
)

//	@provider
type DictionaryGroupDao struct {
	query *query.Query
}

func (dao *DictionaryGroupDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *DictionaryGroupDao) Context(ctx context.Context) query.IDictionaryGroupDo {
	return dao.query.DictionaryGroup.WithContext(ctx)
}

func (dao *DictionaryGroupDao) decorateSortQueryFilter(query query.IDictionaryGroupDo, sortFilter *common.SortQueryFilter) query.IDictionaryGroupDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.DictionaryGroup.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.DictionaryGroup.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *DictionaryGroupDao) decorateQueryFilter(query query.IDictionaryGroupDo, queryFilter *dto.DictionaryGroupListQueryFilter) query.IDictionaryGroupDo {
	query = query.Where(
		dao.query.DictionaryGroup.TenantID.Eq(queryFilter.TenantID),
		dao.query.DictionaryGroup.UserID.In(0, queryFilter.UserID),
	)

	if queryFilter == nil {
		return query
	}
	if queryFilter.Name != nil {
		query = query.Where(dao.query.DictionaryGroup.Name.Eq(*queryFilter.Name))
	}
	if queryFilter.Slug != nil {
		query = query.Where(dao.query.DictionaryGroup.Slug.Eq(*queryFilter.Slug))
	}
	if queryFilter.Description != nil {
		query = query.Where(dao.query.DictionaryGroup.Description.Eq(*queryFilter.Description))
	}

	return query
}

func (dao *DictionaryGroupDao) UpdateColumn(ctx context.Context, id int64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.DictionaryGroup.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *DictionaryGroupDao) Update(ctx context.Context, model *models.DictionaryGroup) error {
	_, err := dao.Context(ctx).Where(dao.query.DictionaryGroup.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *DictionaryGroupDao) Delete(ctx context.Context, tenantID, id int64) error {
	_, err := dao.Context(ctx).Where(
		dao.query.DictionaryGroup.ID.Eq(id),
		dao.query.DictionaryGroup.TenantID.Eq(tenantID),
	).Delete()
	return err
}

func (dao *DictionaryGroupDao) DeleteByUserID(ctx context.Context, tenantID, userID, id int64) error {
	_, err := dao.Context(ctx).Where(
		dao.query.DictionaryGroup.ID.Eq(id),
		dao.query.DictionaryGroup.TenantID.Eq(tenantID),
		dao.query.DictionaryGroup.UserID.Eq(userID),
	).Delete()
	return err
}

func (dao *DictionaryGroupDao) DeletePermanently(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.DictionaryGroup.ID.Eq(id)).Delete()
	return err
}

func (dao *DictionaryGroupDao) Restore(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.DictionaryGroup.ID.Eq(id)).UpdateSimple(dao.query.DictionaryGroup.DeletedAt.Null())
	return err
}

func (dao *DictionaryGroupDao) Create(ctx context.Context, model *models.DictionaryGroup) error {
	return dao.Context(ctx).Create(model)
}

func (dao *DictionaryGroupDao) IsIDBelongToTenant(ctx context.Context, tenantID, id int64) bool {
	count, err := dao.Context(ctx).Where(
		dao.query.DictionaryGroup.ID.Eq(id),
		dao.query.DictionaryGroup.TenantID.Eq(tenantID),
	).Count()
	if err != nil {
		return false
	}

	return count > 0
}

func (dao *DictionaryGroupDao) GetByTenantID(ctx context.Context, tenantID, id int64) (*models.DictionaryGroup, error) {
	return dao.Context(ctx).Where(
		dao.query.DictionaryGroup.ID.Eq(id),
		dao.query.DictionaryGroup.TenantID.Eq(tenantID),
	).First()
}

func (dao *DictionaryGroupDao) GetByUserID(ctx context.Context, tenantID, userID, id int64) (*models.DictionaryGroup, error) {
	return dao.Context(ctx).Where(
		dao.query.DictionaryGroup.ID.Eq(id),
		dao.query.DictionaryGroup.TenantID.Eq(tenantID),
		dao.query.DictionaryGroup.UserID.Eq(userID),
	).First()
}

func (dao *DictionaryGroupDao) GetByIDs(ctx context.Context, ids []int64) ([]*models.DictionaryGroup, error) {
	return dao.Context(ctx).Where(dao.query.DictionaryGroup.ID.In(ids...)).Find()
}

func (dao *DictionaryGroupDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DictionaryGroupListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.DictionaryGroup, int64, error) {
	query := dao.query.DictionaryGroup
	dictionaryGroupQuery := query.WithContext(ctx)
	dictionaryGroupQuery = dao.decorateQueryFilter(dictionaryGroupQuery, queryFilter)
	dictionaryGroupQuery = dao.decorateSortQueryFilter(dictionaryGroupQuery, sortFilter)
	return dictionaryGroupQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *DictionaryGroupDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DictionaryGroupListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.DictionaryGroup, error) {
	query := dao.query.DictionaryGroup
	dictionaryGroupQuery := query.WithContext(ctx)
	dictionaryGroupQuery = dao.decorateQueryFilter(dictionaryGroupQuery, queryFilter)
	dictionaryGroupQuery = dao.decorateSortQueryFilter(dictionaryGroupQuery, sortFilter)
	return dictionaryGroupQuery.Find()
}

func (dao *DictionaryGroupDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DictionaryGroupListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.DictionaryGroup, error) {
	query := dao.query.DictionaryGroup
	dictionaryGroupQuery := query.WithContext(ctx)
	dictionaryGroupQuery = dao.decorateQueryFilter(dictionaryGroupQuery, queryFilter)
	dictionaryGroupQuery = dao.decorateSortQueryFilter(dictionaryGroupQuery, sortFilter)
	return dictionaryGroupQuery.First()
}

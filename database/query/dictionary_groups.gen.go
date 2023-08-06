// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/atom-apps/dictionary/database/models"
)

func newDictionaryGroup(db *gorm.DB, opts ...gen.DOOption) dictionaryGroup {
	_dictionaryGroup := dictionaryGroup{}

	_dictionaryGroup.dictionaryGroupDo.UseDB(db, opts...)
	_dictionaryGroup.dictionaryGroupDo.UseModel(&models.DictionaryGroup{})

	tableName := _dictionaryGroup.dictionaryGroupDo.TableName()
	_dictionaryGroup.ALL = field.NewAsterisk(tableName)
	_dictionaryGroup.ID = field.NewInt64(tableName, "id")
	_dictionaryGroup.CreatedAt = field.NewTime(tableName, "created_at")
	_dictionaryGroup.UpdatedAt = field.NewTime(tableName, "updated_at")
	_dictionaryGroup.DeletedAt = field.NewField(tableName, "deleted_at")
	_dictionaryGroup.UserID = field.NewInt64(tableName, "user_id")
	_dictionaryGroup.TenantID = field.NewInt64(tableName, "tenant_id")
	_dictionaryGroup.Name = field.NewString(tableName, "name")
	_dictionaryGroup.Slug = field.NewString(tableName, "slug")
	_dictionaryGroup.Description = field.NewString(tableName, "description")

	_dictionaryGroup.fillFieldMap()

	return _dictionaryGroup
}

type dictionaryGroup struct {
	dictionaryGroupDo dictionaryGroupDo

	ALL         field.Asterisk
	ID          field.Int64
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	UserID      field.Int64  // 用户ID
	TenantID    field.Int64  // 租户ID
	Name        field.String // 字典组名称
	Slug        field.String
	Description field.String // 字典组描述

	fieldMap map[string]field.Expr
}

func (d dictionaryGroup) Table(newTableName string) *dictionaryGroup {
	d.dictionaryGroupDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dictionaryGroup) As(alias string) *dictionaryGroup {
	d.dictionaryGroupDo.DO = *(d.dictionaryGroupDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dictionaryGroup) updateTableName(table string) *dictionaryGroup {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewInt64(table, "id")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewField(table, "deleted_at")
	d.UserID = field.NewInt64(table, "user_id")
	d.TenantID = field.NewInt64(table, "tenant_id")
	d.Name = field.NewString(table, "name")
	d.Slug = field.NewString(table, "slug")
	d.Description = field.NewString(table, "description")

	d.fillFieldMap()

	return d
}

func (d *dictionaryGroup) WithContext(ctx context.Context) IDictionaryGroupDo {
	return d.dictionaryGroupDo.WithContext(ctx)
}

func (d dictionaryGroup) TableName() string { return d.dictionaryGroupDo.TableName() }

func (d dictionaryGroup) Alias() string { return d.dictionaryGroupDo.Alias() }

func (d dictionaryGroup) Columns(cols ...field.Expr) gen.Columns {
	return d.dictionaryGroupDo.Columns(cols...)
}

func (d *dictionaryGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dictionaryGroup) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 9)
	d.fieldMap["id"] = d.ID
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
	d.fieldMap["user_id"] = d.UserID
	d.fieldMap["tenant_id"] = d.TenantID
	d.fieldMap["name"] = d.Name
	d.fieldMap["slug"] = d.Slug
	d.fieldMap["description"] = d.Description
}

func (d dictionaryGroup) clone(db *gorm.DB) dictionaryGroup {
	d.dictionaryGroupDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dictionaryGroup) replaceDB(db *gorm.DB) dictionaryGroup {
	d.dictionaryGroupDo.ReplaceDB(db)
	return d
}

type dictionaryGroupDo struct{ gen.DO }

type IDictionaryGroupDo interface {
	gen.SubQuery
	Debug() IDictionaryGroupDo
	WithContext(ctx context.Context) IDictionaryGroupDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDictionaryGroupDo
	WriteDB() IDictionaryGroupDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDictionaryGroupDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDictionaryGroupDo
	Not(conds ...gen.Condition) IDictionaryGroupDo
	Or(conds ...gen.Condition) IDictionaryGroupDo
	Select(conds ...field.Expr) IDictionaryGroupDo
	Where(conds ...gen.Condition) IDictionaryGroupDo
	Order(conds ...field.Expr) IDictionaryGroupDo
	Distinct(cols ...field.Expr) IDictionaryGroupDo
	Omit(cols ...field.Expr) IDictionaryGroupDo
	Join(table schema.Tabler, on ...field.Expr) IDictionaryGroupDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDictionaryGroupDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDictionaryGroupDo
	Group(cols ...field.Expr) IDictionaryGroupDo
	Having(conds ...gen.Condition) IDictionaryGroupDo
	Limit(limit int) IDictionaryGroupDo
	Offset(offset int) IDictionaryGroupDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDictionaryGroupDo
	Unscoped() IDictionaryGroupDo
	Create(values ...*models.DictionaryGroup) error
	CreateInBatches(values []*models.DictionaryGroup, batchSize int) error
	Save(values ...*models.DictionaryGroup) error
	First() (*models.DictionaryGroup, error)
	Take() (*models.DictionaryGroup, error)
	Last() (*models.DictionaryGroup, error)
	Find() ([]*models.DictionaryGroup, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DictionaryGroup, err error)
	FindInBatches(result *[]*models.DictionaryGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.DictionaryGroup) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDictionaryGroupDo
	Assign(attrs ...field.AssignExpr) IDictionaryGroupDo
	Joins(fields ...field.RelationField) IDictionaryGroupDo
	Preload(fields ...field.RelationField) IDictionaryGroupDo
	FirstOrInit() (*models.DictionaryGroup, error)
	FirstOrCreate() (*models.DictionaryGroup, error)
	FindByPage(offset int, limit int) (result []*models.DictionaryGroup, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDictionaryGroupDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dictionaryGroupDo) Debug() IDictionaryGroupDo {
	return d.withDO(d.DO.Debug())
}

func (d dictionaryGroupDo) WithContext(ctx context.Context) IDictionaryGroupDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dictionaryGroupDo) ReadDB() IDictionaryGroupDo {
	return d.Clauses(dbresolver.Read)
}

func (d dictionaryGroupDo) WriteDB() IDictionaryGroupDo {
	return d.Clauses(dbresolver.Write)
}

func (d dictionaryGroupDo) Session(config *gorm.Session) IDictionaryGroupDo {
	return d.withDO(d.DO.Session(config))
}

func (d dictionaryGroupDo) Clauses(conds ...clause.Expression) IDictionaryGroupDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dictionaryGroupDo) Returning(value interface{}, columns ...string) IDictionaryGroupDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dictionaryGroupDo) Not(conds ...gen.Condition) IDictionaryGroupDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dictionaryGroupDo) Or(conds ...gen.Condition) IDictionaryGroupDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dictionaryGroupDo) Select(conds ...field.Expr) IDictionaryGroupDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dictionaryGroupDo) Where(conds ...gen.Condition) IDictionaryGroupDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dictionaryGroupDo) Order(conds ...field.Expr) IDictionaryGroupDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dictionaryGroupDo) Distinct(cols ...field.Expr) IDictionaryGroupDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dictionaryGroupDo) Omit(cols ...field.Expr) IDictionaryGroupDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dictionaryGroupDo) Join(table schema.Tabler, on ...field.Expr) IDictionaryGroupDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dictionaryGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDictionaryGroupDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dictionaryGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) IDictionaryGroupDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dictionaryGroupDo) Group(cols ...field.Expr) IDictionaryGroupDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dictionaryGroupDo) Having(conds ...gen.Condition) IDictionaryGroupDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dictionaryGroupDo) Limit(limit int) IDictionaryGroupDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dictionaryGroupDo) Offset(offset int) IDictionaryGroupDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dictionaryGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDictionaryGroupDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dictionaryGroupDo) Unscoped() IDictionaryGroupDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dictionaryGroupDo) Create(values ...*models.DictionaryGroup) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dictionaryGroupDo) CreateInBatches(values []*models.DictionaryGroup, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dictionaryGroupDo) Save(values ...*models.DictionaryGroup) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dictionaryGroupDo) First() (*models.DictionaryGroup, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.DictionaryGroup), nil
	}
}

func (d dictionaryGroupDo) Take() (*models.DictionaryGroup, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.DictionaryGroup), nil
	}
}

func (d dictionaryGroupDo) Last() (*models.DictionaryGroup, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.DictionaryGroup), nil
	}
}

func (d dictionaryGroupDo) Find() ([]*models.DictionaryGroup, error) {
	result, err := d.DO.Find()
	return result.([]*models.DictionaryGroup), err
}

func (d dictionaryGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DictionaryGroup, err error) {
	buf := make([]*models.DictionaryGroup, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dictionaryGroupDo) FindInBatches(result *[]*models.DictionaryGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dictionaryGroupDo) Attrs(attrs ...field.AssignExpr) IDictionaryGroupDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dictionaryGroupDo) Assign(attrs ...field.AssignExpr) IDictionaryGroupDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dictionaryGroupDo) Joins(fields ...field.RelationField) IDictionaryGroupDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dictionaryGroupDo) Preload(fields ...field.RelationField) IDictionaryGroupDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dictionaryGroupDo) FirstOrInit() (*models.DictionaryGroup, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.DictionaryGroup), nil
	}
}

func (d dictionaryGroupDo) FirstOrCreate() (*models.DictionaryGroup, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.DictionaryGroup), nil
	}
}

func (d dictionaryGroupDo) FindByPage(offset int, limit int) (result []*models.DictionaryGroup, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dictionaryGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dictionaryGroupDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dictionaryGroupDo) Delete(models ...*models.DictionaryGroup) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dictionaryGroupDo) withDO(do gen.Dao) *dictionaryGroupDo {
	d.DO = *do.(*gen.DO)
	return d
}

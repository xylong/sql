// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"sql/model"
)

func newAddress(db *gorm.DB, opts ...gen.DOOption) address {
	_address := address{}

	_address.addressDo.UseDB(db, opts...)
	_address.addressDo.UseModel(&model.Address{})

	tableName := _address.addressDo.TableName()
	_address.ALL = field.NewAsterisk(tableName)
	_address.ID = field.NewInt64(tableName, "id")
	_address.UserID = field.NewInt64(tableName, "user_id")
	_address.Province = field.NewString(tableName, "province")
	_address.City = field.NewString(tableName, "city")
	_address.County = field.NewString(tableName, "county")
	_address.Address = field.NewString(tableName, "address")
	_address.CreatedAt = field.NewTime(tableName, "created_at")
	_address.UpdatedAt = field.NewTime(tableName, "updated_at")
	_address.DeletedAt = field.NewField(tableName, "deleted_at")
	_address.User = addressBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "model.User"),
		Address: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("User.Address", "model.Address"),
		},
	}

	_address.fillFieldMap()

	return _address
}

type address struct {
	addressDo

	ALL       field.Asterisk
	ID        field.Int64  // id
	UserID    field.Int64  // 用户id
	Province  field.String // 省
	City      field.String // 市
	County    field.String // 县/区
	Address   field.String // 详细地址
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 更新时间
	DeletedAt field.Field  // 删除时间
	User      addressBelongsToUser

	fieldMap map[string]field.Expr
}

func (a address) Table(newTableName string) *address {
	a.addressDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a address) As(alias string) *address {
	a.addressDo.DO = *(a.addressDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *address) updateTableName(table string) *address {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.UserID = field.NewInt64(table, "user_id")
	a.Province = field.NewString(table, "province")
	a.City = field.NewString(table, "city")
	a.County = field.NewString(table, "county")
	a.Address = field.NewString(table, "address")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")

	a.fillFieldMap()

	return a
}

func (a *address) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *address) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["id"] = a.ID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["province"] = a.Province
	a.fieldMap["city"] = a.City
	a.fieldMap["county"] = a.County
	a.fieldMap["address"] = a.Address
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt

}

func (a address) clone(db *gorm.DB) address {
	a.addressDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a address) replaceDB(db *gorm.DB) address {
	a.addressDo.ReplaceDB(db)
	return a
}

type addressBelongsToUser struct {
	db *gorm.DB

	field.RelationField

	Address struct {
		field.RelationField
	}
}

func (a addressBelongsToUser) Where(conds ...field.Expr) *addressBelongsToUser {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a addressBelongsToUser) WithContext(ctx context.Context) *addressBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a addressBelongsToUser) Session(session *gorm.Session) *addressBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a addressBelongsToUser) Model(m *model.Address) *addressBelongsToUserTx {
	return &addressBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

type addressBelongsToUserTx struct{ tx *gorm.Association }

func (a addressBelongsToUserTx) Find() (result *model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a addressBelongsToUserTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a addressBelongsToUserTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a addressBelongsToUserTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a addressBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a addressBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

type addressDo struct{ gen.DO }

type IAddressDo interface {
	gen.SubQuery
	Debug() IAddressDo
	WithContext(ctx context.Context) IAddressDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAddressDo
	WriteDB() IAddressDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAddressDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAddressDo
	Not(conds ...gen.Condition) IAddressDo
	Or(conds ...gen.Condition) IAddressDo
	Select(conds ...field.Expr) IAddressDo
	Where(conds ...gen.Condition) IAddressDo
	Order(conds ...field.Expr) IAddressDo
	Distinct(cols ...field.Expr) IAddressDo
	Omit(cols ...field.Expr) IAddressDo
	Join(table schema.Tabler, on ...field.Expr) IAddressDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAddressDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAddressDo
	Group(cols ...field.Expr) IAddressDo
	Having(conds ...gen.Condition) IAddressDo
	Limit(limit int) IAddressDo
	Offset(offset int) IAddressDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAddressDo
	Unscoped() IAddressDo
	Create(values ...*model.Address) error
	CreateInBatches(values []*model.Address, batchSize int) error
	Save(values ...*model.Address) error
	First() (*model.Address, error)
	Take() (*model.Address, error)
	Last() (*model.Address, error)
	Find() ([]*model.Address, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Address, err error)
	FindInBatches(result *[]*model.Address, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Address) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAddressDo
	Assign(attrs ...field.AssignExpr) IAddressDo
	Joins(fields ...field.RelationField) IAddressDo
	Preload(fields ...field.RelationField) IAddressDo
	FirstOrInit() (*model.Address, error)
	FirstOrCreate() (*model.Address, error)
	FindByPage(offset int, limit int) (result []*model.Address, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAddressDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a addressDo) Debug() IAddressDo {
	return a.withDO(a.DO.Debug())
}

func (a addressDo) WithContext(ctx context.Context) IAddressDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a addressDo) ReadDB() IAddressDo {
	return a.Clauses(dbresolver.Read)
}

func (a addressDo) WriteDB() IAddressDo {
	return a.Clauses(dbresolver.Write)
}

func (a addressDo) Session(config *gorm.Session) IAddressDo {
	return a.withDO(a.DO.Session(config))
}

func (a addressDo) Clauses(conds ...clause.Expression) IAddressDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a addressDo) Returning(value interface{}, columns ...string) IAddressDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a addressDo) Not(conds ...gen.Condition) IAddressDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a addressDo) Or(conds ...gen.Condition) IAddressDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a addressDo) Select(conds ...field.Expr) IAddressDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a addressDo) Where(conds ...gen.Condition) IAddressDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a addressDo) Order(conds ...field.Expr) IAddressDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a addressDo) Distinct(cols ...field.Expr) IAddressDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a addressDo) Omit(cols ...field.Expr) IAddressDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a addressDo) Join(table schema.Tabler, on ...field.Expr) IAddressDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a addressDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAddressDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a addressDo) RightJoin(table schema.Tabler, on ...field.Expr) IAddressDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a addressDo) Group(cols ...field.Expr) IAddressDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a addressDo) Having(conds ...gen.Condition) IAddressDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a addressDo) Limit(limit int) IAddressDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a addressDo) Offset(offset int) IAddressDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a addressDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAddressDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a addressDo) Unscoped() IAddressDo {
	return a.withDO(a.DO.Unscoped())
}

func (a addressDo) Create(values ...*model.Address) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a addressDo) CreateInBatches(values []*model.Address, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a addressDo) Save(values ...*model.Address) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a addressDo) First() (*model.Address, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Address), nil
	}
}

func (a addressDo) Take() (*model.Address, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Address), nil
	}
}

func (a addressDo) Last() (*model.Address, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Address), nil
	}
}

func (a addressDo) Find() ([]*model.Address, error) {
	result, err := a.DO.Find()
	return result.([]*model.Address), err
}

func (a addressDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Address, err error) {
	buf := make([]*model.Address, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a addressDo) FindInBatches(result *[]*model.Address, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a addressDo) Attrs(attrs ...field.AssignExpr) IAddressDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a addressDo) Assign(attrs ...field.AssignExpr) IAddressDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a addressDo) Joins(fields ...field.RelationField) IAddressDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a addressDo) Preload(fields ...field.RelationField) IAddressDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a addressDo) FirstOrInit() (*model.Address, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Address), nil
	}
}

func (a addressDo) FirstOrCreate() (*model.Address, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Address), nil
	}
}

func (a addressDo) FindByPage(offset int, limit int) (result []*model.Address, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a addressDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a addressDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a addressDo) Delete(models ...*model.Address) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *addressDo) withDO(do gen.Dao) *addressDo {
	a.DO = *do.(*gen.DO)
	return a
}

package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"godbfexam/adf/entity"
	"godbfexam/adf/cb"
	"godbfexam/adf/meta"
)

type ProductBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *ProductBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.ProductDbm
	return &meta
}
func (l *ProductBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *ProductBhv) Update(product *entity.Product, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = product
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *ProductBhv) Insert(product *entity.Product, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = product
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *ProductBhv) Delete(product *entity.Product, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = product
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *ProductBhv) QueryUpdate(ent *entity.Product, cb *cb.ProductCB, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = ent
	return l.BaseBehavior.DoQueryUpdate(&entity, cb, nil, tx, ctx)
}
func (l *ProductBhv) QueryDelete(cb *cb.ProductCB, tx *sql.Tx) (int64, error) {
	return l.BaseBehavior.DoQueryDelete(cb, "Product", nil, tx)
}
func (l *ProductBhv) SelectList(cb *cb.ProductCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "Product", tx)
}
func (l *ProductBhv) SelectCount(cb *cb.ProductCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *ProductBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *ProductBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var ProductBhv_I *ProductBhv

package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"godbfexam/adf/entity"
	"godbfexam/adf/cb"
	"godbfexam/adf/meta"
)

type CustomerBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *CustomerBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.CustomerDbm
	return &meta
}
func (l *CustomerBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *CustomerBhv) Update(customer *entity.Customer, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = customer
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *CustomerBhv) Insert(customer *entity.Customer, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = customer
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *CustomerBhv) Delete(customer *entity.Customer, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = customer
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *CustomerBhv) QueryUpdate(ent *entity.Customer, cb *cb.CustomerCB, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = ent
	return l.BaseBehavior.DoQueryUpdate(&entity, cb, nil, tx, ctx)
}
func (l *CustomerBhv) QueryDelete(cb *cb.CustomerCB, tx *sql.Tx) (int64, error) {
	return l.BaseBehavior.DoQueryDelete(cb, "Customer", nil, tx)
}
func (l *CustomerBhv) SelectList(cb *cb.CustomerCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "Customer", tx)
}
func (l *CustomerBhv) SelectCount(cb *cb.CustomerCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *CustomerBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *CustomerBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var CustomerBhv_I *CustomerBhv

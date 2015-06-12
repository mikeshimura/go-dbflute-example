package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"godbfexam/adf/entity"
	"godbfexam/adf/cb"
	"godbfexam/adf/meta"
)

type SalesSlipBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *SalesSlipBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.SalesSlipDbm
	return &meta
}
func (l *SalesSlipBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *SalesSlipBhv) Update(salesSlip *entity.SalesSlip, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = salesSlip
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *SalesSlipBhv) Insert(salesSlip *entity.SalesSlip, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = salesSlip
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *SalesSlipBhv) Delete(salesSlip *entity.SalesSlip, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = salesSlip
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *SalesSlipBhv) QueryUpdate(ent *entity.SalesSlip, cb *cb.SalesSlipCB, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = ent
	return l.BaseBehavior.DoQueryUpdate(&entity, cb, nil, tx, ctx)
}
func (l *SalesSlipBhv) QueryDelete(cb *cb.SalesSlipCB, tx *sql.Tx) (int64, error) {
	return l.BaseBehavior.DoQueryDelete(cb, "SalesSlip", nil, tx)
}
func (l *SalesSlipBhv) SelectList(cb *cb.SalesSlipCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "SalesSlip", tx)
}
func (l *SalesSlipBhv) SelectCount(cb *cb.SalesSlipCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *SalesSlipBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *SalesSlipBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var SalesSlipBhv_I *SalesSlipBhv

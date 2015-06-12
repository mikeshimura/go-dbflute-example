package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"godbfexam/adf/entity"
	"godbfexam/adf/cb"
	"godbfexam/adf/meta"
)

type StockBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *StockBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.StockDbm
	return &meta
}
func (l *StockBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *StockBhv) Update(stock *entity.Stock, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = stock
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *StockBhv) Insert(stock *entity.Stock, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = stock
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *StockBhv) Delete(stock *entity.Stock, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = stock
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *StockBhv) QueryUpdate(ent *entity.Stock, cb *cb.StockCB, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = ent
	return l.BaseBehavior.DoQueryUpdate(&entity, cb, nil, tx, ctx)
}
func (l *StockBhv) QueryDelete(cb *cb.StockCB, tx *sql.Tx) (int64, error) {
	return l.BaseBehavior.DoQueryDelete(cb, "Stock", nil, tx)
}
func (l *StockBhv) SelectList(cb *cb.StockCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "Stock", tx)
}
func (l *StockBhv) SelectCount(cb *cb.StockCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *StockBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *StockBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var StockBhv_I *StockBhv

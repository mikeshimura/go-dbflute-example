package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"godbfexam/adf/entity"
	"godbfexam/adf/cb"
	"godbfexam/adf/meta"
)

type TestTableBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *TestTableBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.TestTableDbm
	return &meta
}
func (l *TestTableBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *TestTableBhv) Update(testTable *entity.TestTable, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = testTable
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *TestTableBhv) Insert(testTable *entity.TestTable, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = testTable
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *TestTableBhv) Delete(testTable *entity.TestTable, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = testTable
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *TestTableBhv) QueryUpdate(ent *entity.TestTable, cb *cb.TestTableCB, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = ent
	return l.BaseBehavior.DoQueryUpdate(&entity, cb, nil, tx, ctx)
}
func (l *TestTableBhv) QueryDelete(cb *cb.TestTableCB, tx *sql.Tx) (int64, error) {
	return l.BaseBehavior.DoQueryDelete(cb, "TestTable", nil, tx)
}
func (l *TestTableBhv) SelectList(cb *cb.TestTableCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "TestTable", tx)
}
func (l *TestTableBhv) SelectCount(cb *cb.TestTableCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *TestTableBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *TestTableBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var TestTableBhv_I *TestTableBhv

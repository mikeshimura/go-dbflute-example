package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"mikegou/adf/entity"
	"mikegou/adf/cb"
	"mikegou/adf/meta"
)

type DbfluteBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *DbfluteBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.DbfluteDbm
	return &meta
}
func (l *DbfluteBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *DbfluteBhv) Update(dbflute *entity.Dbflute, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = dbflute
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *DbfluteBhv) Insert(dbflute *entity.Dbflute, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = dbflute
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *DbfluteBhv) Delete(dbflute *entity.Dbflute, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = dbflute
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *DbfluteBhv) QueryUpdate(ent *entity.Dbflute, cb *cb.DbfluteCB, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = ent
	return l.BaseBehavior.DoQueryUpdate(&entity, cb, nil, tx, ctx)
}
func (l *DbfluteBhv) QueryDelete(cb *cb.DbfluteCB, tx *sql.Tx) (int64, error) {
	return l.BaseBehavior.DoQueryDelete(cb, "Dbflute", nil, tx)
}
func (l *DbfluteBhv) SelectList(cb *cb.DbfluteCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "Dbflute", tx)
}
func (l *DbfluteBhv) SelectCount(cb *cb.DbfluteCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *DbfluteBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *DbfluteBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var DbfluteBhv_I *DbfluteBhv

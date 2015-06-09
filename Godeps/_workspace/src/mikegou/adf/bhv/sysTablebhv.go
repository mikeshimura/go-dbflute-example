package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"mikegou/adf/entity"
	"mikegou/adf/cb"
	"mikegou/adf/meta"
)

type SysTableBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *SysTableBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.SysTableDbm
	return &meta
}
func (l *SysTableBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *SysTableBhv) Update(sysTable *entity.SysTable, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = sysTable
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *SysTableBhv) Insert(sysTable *entity.SysTable, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = sysTable
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *SysTableBhv) Delete(sysTable *entity.SysTable, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = sysTable
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *SysTableBhv) QueryUpdate(ent *entity.SysTable, cb *cb.SysTableCB, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = ent
	return l.BaseBehavior.DoQueryUpdate(&entity, cb, nil, tx, ctx)
}
func (l *SysTableBhv) QueryDelete(cb *cb.SysTableCB, tx *sql.Tx) (int64, error) {
	return l.BaseBehavior.DoQueryDelete(cb, "SysTable", nil, tx)
}
func (l *SysTableBhv) SelectList(cb *cb.SysTableCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "SysTable", tx)
}
func (l *SysTableBhv) SelectCount(cb *cb.SysTableCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *SysTableBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *SysTableBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var SysTableBhv_I *SysTableBhv

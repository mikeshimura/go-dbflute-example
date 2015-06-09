package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"mikegou/adf/entity"
	"mikegou/adf/cb"
	"mikegou/adf/meta"
)

type UserTableBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *UserTableBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.UserTableDbm
	return &meta
}
func (l *UserTableBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *UserTableBhv) Update(userTable *entity.UserTable, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = userTable
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *UserTableBhv) Insert(userTable *entity.UserTable, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = userTable
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *UserTableBhv) Delete(userTable *entity.UserTable, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = userTable
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *UserTableBhv) QueryUpdate(ent *entity.UserTable, cb *cb.UserTableCB, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = ent
	return l.BaseBehavior.DoQueryUpdate(&entity, cb, nil, tx, ctx)
}
func (l *UserTableBhv) QueryDelete(cb *cb.UserTableCB, tx *sql.Tx) (int64, error) {
	return l.BaseBehavior.DoQueryDelete(cb, "UserTable", nil, tx)
}
func (l *UserTableBhv) SelectList(cb *cb.UserTableCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "UserTable", tx)
}
func (l *UserTableBhv) SelectCount(cb *cb.UserTableCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *UserTableBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *UserTableBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var UserTableBhv_I *UserTableBhv

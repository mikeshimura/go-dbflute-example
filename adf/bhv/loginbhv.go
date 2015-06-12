package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"godbfexam/adf/entity"
	"godbfexam/adf/cb"
	"godbfexam/adf/meta"
)

type LoginBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *LoginBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.LoginDbm
	return &meta
}
func (l *LoginBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *LoginBhv) Update(login *entity.Login, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = login
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *LoginBhv) Insert(login *entity.Login, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = login
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *LoginBhv) Delete(login *entity.Login, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = login
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *LoginBhv) QueryUpdate(ent *entity.Login, cb *cb.LoginCB, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = ent
	return l.BaseBehavior.DoQueryUpdate(&entity, cb, nil, tx, ctx)
}
func (l *LoginBhv) QueryDelete(cb *cb.LoginCB, tx *sql.Tx) (int64, error) {
	return l.BaseBehavior.DoQueryDelete(cb, "Login", nil, tx)
}
func (l *LoginBhv) SelectList(cb *cb.LoginCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "Login", tx)
}
func (l *LoginBhv) SelectCount(cb *cb.LoginCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *LoginBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *LoginBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var LoginBhv_I *LoginBhv

package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"mikegou/adf/entity"
	"mikegou/adf/cb"
	"mikegou/adf/meta"
)

type SessionBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *SessionBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.SessionDbm
	return &meta
}
func (l *SessionBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *SessionBhv) Update(session *entity.Session, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = session
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *SessionBhv) Insert(session *entity.Session, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = session
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *SessionBhv) Delete(session *entity.Session, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = session
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *SessionBhv) QueryUpdate(ent *entity.Session, cb *cb.SessionCB, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = ent
	return l.BaseBehavior.DoQueryUpdate(&entity, cb, nil, tx, ctx)
}
func (l *SessionBhv) QueryDelete(cb *cb.SessionCB, tx *sql.Tx) (int64, error) {
	return l.BaseBehavior.DoQueryDelete(cb, "Session", nil, tx)
}
func (l *SessionBhv) SelectList(cb *cb.SessionCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "Session", tx)
}
func (l *SessionBhv) SelectCount(cb *cb.SessionCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *SessionBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *SessionBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var SessionBhv_I *SessionBhv

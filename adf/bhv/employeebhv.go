package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"godbfexam/adf/entity"
	"godbfexam/adf/cb"
	"godbfexam/adf/meta"
)

type EmployeeBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *EmployeeBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.EmployeeDbm
	return &meta
}
func (l *EmployeeBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *EmployeeBhv) Update(employee *entity.Employee, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = employee
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *EmployeeBhv) Insert(employee *entity.Employee, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = employee
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *EmployeeBhv) Delete(employee *entity.Employee, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = employee
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *EmployeeBhv) QueryUpdate(ent *entity.Employee, cb *cb.EmployeeCB, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = ent
	return l.BaseBehavior.DoQueryUpdate(&entity, cb, nil, tx, ctx)
}
func (l *EmployeeBhv) QueryDelete(cb *cb.EmployeeCB, tx *sql.Tx) (int64, error) {
	return l.BaseBehavior.DoQueryDelete(cb, "Employee", nil, tx)
}
func (l *EmployeeBhv) SelectList(cb *cb.EmployeeCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "Employee", tx)
}
func (l *EmployeeBhv) SelectCount(cb *cb.EmployeeCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *EmployeeBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *EmployeeBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var EmployeeBhv_I *EmployeeBhv

package pmb

import (
    "github.com/mikeshimura/dbflute/df"
    "database/sql"
	"mikegou/adf/bhv"
	_ "mikegou/adf/cmeta"
	_ "mikegou/adf/centity"
)

type C_DbTableSelectPmb struct {
    df.BasePmb
    LoginId interface{}
    Db interface{}
    DbName interface{}
}
func (o *C_DbTableSelectPmb) GetLoginId() interface{} {
	return o.LoginId
}
func (o *C_DbTableSelectPmb) SetLoginId(value int64) {
	o.LoginId = value	
}
func (o *C_DbTableSelectPmb) GetDb() interface{} {
	return o.Db
}
func (o *C_DbTableSelectPmb) SetDb(value string) {
	o.Db = o.CheckAndComvertEmptyToNull(value)
}
func (o *C_DbTableSelectPmb) GetDbName() interface{} {
	return o.DbName
}
func (o *C_DbTableSelectPmb) SetDbName(value string) {
	o.DbName = o.CheckAndComvertEmptyToNull(value)
}

func (o *C_DbTableSelectPmb) GetEntityType() string {
	return "C_DbTableSelect"
}
func (o *C_DbTableSelectPmb) GetOutsideSqlPath() string {
	return "mikegou/adf/bhv/sql/DbfluteBhv_DbTableSelect.sql"
}

func (o *C_DbTableSelectPmb) SelectList(tx *sql.Tx) (*df.ListResultBean, error) {
	return bhv.DbfluteBhv_I.OutsideSql().SelectList(o,tx)

}

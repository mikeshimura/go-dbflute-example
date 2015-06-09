package pmb

import (
    "github.com/mikeshimura/dbflute/df"
    "database/sql"
	"mikegou/adf/bhv"
	_ "mikegou/adf/cmeta"
	_ "mikegou/adf/centity"
)

type C_DatabaseSelectPmb struct {
    df.BasePmb
    LoginId interface{}
    Db interface{}
}
func (o *C_DatabaseSelectPmb) GetLoginId() interface{} {
	return o.LoginId
}
func (o *C_DatabaseSelectPmb) SetLoginId(value int64) {
	o.LoginId = value	
}
func (o *C_DatabaseSelectPmb) GetDb() interface{} {
	return o.Db
}
func (o *C_DatabaseSelectPmb) SetDb(value string) {
	o.Db = o.CheckAndComvertEmptyToNull(value)
}

func (o *C_DatabaseSelectPmb) GetEntityType() string {
	return "C_DatabaseSelect"
}
func (o *C_DatabaseSelectPmb) GetOutsideSqlPath() string {
	return "mikegou/adf/bhv/sql/DbfluteBhv_DatabaseSelect.sql"
}

func (o *C_DatabaseSelectPmb) SelectList(tx *sql.Tx) (*df.ListResultBean, error) {
	return bhv.DbfluteBhv_I.OutsideSql().SelectList(o,tx)

}

package pmb

import (
    "github.com/mikeshimura/dbflute/df"
    "database/sql"
	"mikegou/adf/bhv"
	_ "mikegou/adf/cmeta"
	_ "mikegou/adf/centity"
)

type C_DatabaseDeletePmb struct {
    df.BasePmb
    LoginId interface{}
    Database interface{}
}
func (o *C_DatabaseDeletePmb) GetLoginId() interface{} {
	return o.LoginId
}
func (o *C_DatabaseDeletePmb) SetLoginId(value int64) {
	o.LoginId = value	
}
func (o *C_DatabaseDeletePmb) GetDatabase() interface{} {
	return o.Database
}
func (o *C_DatabaseDeletePmb) SetDatabase(value string) {
	o.Database = o.CheckAndComvertEmptyToNull(value)
}

func (o *C_DatabaseDeletePmb) GetEntityType() string {
	return "C_DatabaseDelete"
}
func (o *C_DatabaseDeletePmb) GetOutsideSqlPath() string {
	return "mikegou/adf/bhv/sql/DbfluteBhv_DatabaseDelete.sql"
}

func (o *C_DatabaseDeletePmb) Execute(tx *sql.Tx) (int64, error) {
	return bhv.DbfluteBhv_I.OutsideSql().Execute(o,tx)

}

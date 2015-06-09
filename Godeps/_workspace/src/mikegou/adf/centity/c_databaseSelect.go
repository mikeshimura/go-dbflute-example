package centity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type C_DatabaseSelect  struct {
	DbName sql.NullString
	df.BaseEntity
}
func (l *C_DatabaseSelect) GetDbName () sql.NullString {
	return l.DbName
}


func (t *C_DatabaseSelect) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 1)
	i[0] = &(t.DbName)
	return i
}

func (t *C_DatabaseSelect) AsTableDbName() string {
	return "c_DatabaseSelect"
}

func (t *C_DatabaseSelect) HasPrimaryKeyValue() bool{
        return false;
}
func (t *C_DatabaseSelect) SetDbName(dbName sql.NullString) {
	t.AddPropertyName("dbName")
	t.DbName = dbName
}


func (t *C_DatabaseSelect) SetUp(){
}


func (t *C_DatabaseSelect)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}

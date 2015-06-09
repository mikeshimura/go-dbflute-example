package centity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type C_DbTableSelect  struct {
	DbTable sql.NullString
	df.BaseEntity
}
func (l *C_DbTableSelect) GetDbTable () sql.NullString {
	return l.DbTable
}


func (t *C_DbTableSelect) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 1)
	i[0] = &(t.DbTable)
	return i
}

func (t *C_DbTableSelect) AsTableDbName() string {
	return "c_DbTableSelect"
}

func (t *C_DbTableSelect) HasPrimaryKeyValue() bool{
        return false;
}
func (t *C_DbTableSelect) SetDbTable(dbTable sql.NullString) {
	t.AddPropertyName("dbTable")
	t.DbTable = dbTable
}


func (t *C_DbTableSelect) SetUp(){
}


func (t *C_DbTableSelect)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}

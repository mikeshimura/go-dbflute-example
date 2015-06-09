package cmeta

import (
	"github.com/mikeshimura/dbflute/df"
)

type C_DatabaseSelectDbm_T struct {
	df.BaseDBMeta
	ColumnDbName *df.ColumnInfo
}

func (b *C_DatabaseSelectDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *C_DatabaseSelectDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var C_DatabaseSelectDbm *C_DatabaseSelectDbm_T

func Create_C_DatabaseSelectDbm() {
	C_DatabaseSelectDbm = new(C_DatabaseSelectDbm_T)
	C_DatabaseSelectDbm.TableDbName = "c_DatabaseSelect"
	C_DatabaseSelectDbm.TableDispName = "c_DatabaseSelect"
	C_DatabaseSelectDbm.TablePropertyName = "c_DatabaseSelect"
	C_DatabaseSelectDbm.TableSqlName = new(df.TableSqlName)
	C_DatabaseSelectDbm.TableSqlName.TableSqlName = "c_DatabaseSelect"
	C_DatabaseSelectDbm.TableSqlName.CorrespondingDbName = C_DatabaseSelectDbm.TableDbName


	var databaseSelect df.DBMeta
	databaseSelect = C_DatabaseSelectDbm
	C_DatabaseSelectDbm.DBMeta=&databaseSelect
	dbNameSqlName := new(df.ColumnSqlName)
	dbNameSqlName.ColumnSqlName = "db_name"
	dbNameSqlName.IrregularChar = false
	C_DatabaseSelectDbm.ColumnDbName = df.CCI(&databaseSelect, "db_name", dbNameSqlName, "", "", "String.class", "dbName", "", false, false,false, "varchar", 40, 0, "",false,"","", "","","",false,"sql.NullString")

	C_DatabaseSelectDbm.ColumnInfoList = new(df.List)
	C_DatabaseSelectDbm.ColumnInfoList.Add(C_DatabaseSelectDbm.ColumnDbName)


	C_DatabaseSelectDbm.ColumnInfoMap=make(map[string]int)
	C_DatabaseSelectDbm.ColumnInfoMap["dbName"]=0
	    C_DatabaseSelectDbm.PrimaryKey = false
    C_DatabaseSelectDbm.CompoundPrimaryKey = false

	var databaseSelectMeta df.DBMeta = C_DatabaseSelectDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["c_DatabaseSelect"] = &databaseSelectMeta
}

package cmeta

import (
	"github.com/mikeshimura/dbflute/df"
)

type C_DbTableSelectDbm_T struct {
	df.BaseDBMeta
	ColumnDbTable *df.ColumnInfo
}

func (b *C_DbTableSelectDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *C_DbTableSelectDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var C_DbTableSelectDbm *C_DbTableSelectDbm_T

func Create_C_DbTableSelectDbm() {
	C_DbTableSelectDbm = new(C_DbTableSelectDbm_T)
	C_DbTableSelectDbm.TableDbName = "c_DbTableSelect"
	C_DbTableSelectDbm.TableDispName = "c_DbTableSelect"
	C_DbTableSelectDbm.TablePropertyName = "c_DbTableSelect"
	C_DbTableSelectDbm.TableSqlName = new(df.TableSqlName)
	C_DbTableSelectDbm.TableSqlName.TableSqlName = "c_DbTableSelect"
	C_DbTableSelectDbm.TableSqlName.CorrespondingDbName = C_DbTableSelectDbm.TableDbName


	var dbTableSelect df.DBMeta
	dbTableSelect = C_DbTableSelectDbm
	C_DbTableSelectDbm.DBMeta=&dbTableSelect
	dbTableSqlName := new(df.ColumnSqlName)
	dbTableSqlName.ColumnSqlName = "db_table"
	dbTableSqlName.IrregularChar = false
	C_DbTableSelectDbm.ColumnDbTable = df.CCI(&dbTableSelect, "db_table", dbTableSqlName, "", "", "String.class", "dbTable", "", false, false,false, "varchar", 40, 0, "",false,"","", "","","",false,"sql.NullString")

	C_DbTableSelectDbm.ColumnInfoList = new(df.List)
	C_DbTableSelectDbm.ColumnInfoList.Add(C_DbTableSelectDbm.ColumnDbTable)


	C_DbTableSelectDbm.ColumnInfoMap=make(map[string]int)
	C_DbTableSelectDbm.ColumnInfoMap["dbTable"]=0
	    C_DbTableSelectDbm.PrimaryKey = false
    C_DbTableSelectDbm.CompoundPrimaryKey = false

	var dbTableSelectMeta df.DBMeta = C_DbTableSelectDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["c_DbTableSelect"] = &dbTableSelectMeta
}

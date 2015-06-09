package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type DbfluteDbm_T struct {
	df.BaseDBMeta
	ColumnId *df.ColumnInfo
	ColumnLoginId *df.ColumnInfo
	ColumnDb *df.ColumnInfo
	ColumnDbName *df.ColumnInfo
	ColumnDbTable *df.ColumnInfo
	ColumnName *df.ColumnInfo
	ColumnDbType *df.ColumnInfo
	ColumnRequired *df.ColumnInfo
	ColumnJavaType *df.ColumnInfo
	ColumnGoType *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
	ColumnDelFlag *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
}

func (b *DbfluteDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *DbfluteDbm_T) foreignLogin() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		DbfluteDbm.GetColumnInfoByPropertyName("loginId"),
		LoginDbm.GetColumnInfoByPropertyName("id"),
	}

	return b.BaseDBMeta.Cfi("dbflute_fkey", "Login",
		columns, 0, false, false, false, false,
		"", nil, false, "dbfluteList")
}	
func (b *DbfluteDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
	b.ForeignInfoMap["Login"] = b.foreignLogin()
}

func (b *DbfluteDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var DbfluteDbm *DbfluteDbm_T

func Create_DbfluteDbm() {
	DbfluteDbm = new(DbfluteDbm_T)
	DbfluteDbm.TableDbName = "dbflute"
	DbfluteDbm.TableDispName = "dbflute"
	DbfluteDbm.TablePropertyName = "dbflute"
	DbfluteDbm.TableSqlName = new(df.TableSqlName)
	DbfluteDbm.TableSqlName.TableSqlName = "dbflute"
	DbfluteDbm.TableSqlName.CorrespondingDbName = DbfluteDbm.TableDbName

	var dbflute df.DBMeta
	dbflute = DbfluteDbm
	DbfluteDbm.DBMeta=&dbflute
	idSqlName := new(df.ColumnSqlName)
	idSqlName.ColumnSqlName = "id"
	idSqlName.IrregularChar = false
	DbfluteDbm.ColumnId = df.CCI(&dbflute, "id", idSqlName, "", "",
               "Integer.class", "id", "", true, true,true, "serial", 10, 0,
               "nextval('dbflute_id_seq'::regclass)",false,"","", "","","",false,"int64")
	loginIdSqlName := new(df.ColumnSqlName)
	loginIdSqlName.ColumnSqlName = "login_id"
	loginIdSqlName.IrregularChar = false
	DbfluteDbm.ColumnLoginId = df.CCI(&dbflute, "login_id", loginIdSqlName, "", "",
               "Integer.class", "loginId", "", false, false,true, "int4", 10, 0,
               "",false,"","", "login","","",false,"int64")
	dbSqlName := new(df.ColumnSqlName)
	dbSqlName.ColumnSqlName = "db"
	dbSqlName.IrregularChar = false
	DbfluteDbm.ColumnDb = df.CCI(&dbflute, "db", dbSqlName, "", "",
               "String.class", "db", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	dbNameSqlName := new(df.ColumnSqlName)
	dbNameSqlName.ColumnSqlName = "db_name"
	dbNameSqlName.IrregularChar = false
	DbfluteDbm.ColumnDbName = df.CCI(&dbflute, "db_name", dbNameSqlName, "", "",
               "String.class", "dbName", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	dbTableSqlName := new(df.ColumnSqlName)
	dbTableSqlName.ColumnSqlName = "db_table"
	dbTableSqlName.IrregularChar = false
	DbfluteDbm.ColumnDbTable = df.CCI(&dbflute, "db_table", dbTableSqlName, "", "",
               "String.class", "dbTable", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	nameSqlName := new(df.ColumnSqlName)
	nameSqlName.ColumnSqlName = "name"
	nameSqlName.IrregularChar = false
	DbfluteDbm.ColumnName = df.CCI(&dbflute, "name", nameSqlName, "", "",
               "String.class", "name", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	dbTypeSqlName := new(df.ColumnSqlName)
	dbTypeSqlName.ColumnSqlName = "db_type"
	dbTypeSqlName.IrregularChar = false
	DbfluteDbm.ColumnDbType = df.CCI(&dbflute, "db_type", dbTypeSqlName, "", "",
               "String.class", "dbType", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	requiredSqlName := new(df.ColumnSqlName)
	requiredSqlName.ColumnSqlName = "required"
	requiredSqlName.IrregularChar = false
	DbfluteDbm.ColumnRequired = df.CCI(&dbflute, "required", requiredSqlName, "", "",
               "Boolean.class", "required", "", false, false,true, "bool", 1, 0,
               "",false,"","", "","","",false,"bool")
	javaTypeSqlName := new(df.ColumnSqlName)
	javaTypeSqlName.ColumnSqlName = "java_type"
	javaTypeSqlName.IrregularChar = false
	DbfluteDbm.ColumnJavaType = df.CCI(&dbflute, "java_type", javaTypeSqlName, "", "",
               "String.class", "javaType", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	goTypeSqlName := new(df.ColumnSqlName)
	goTypeSqlName.ColumnSqlName = "go_type"
	goTypeSqlName.IrregularChar = false
	DbfluteDbm.ColumnGoType = df.CCI(&dbflute, "go_type", goTypeSqlName, "", "",
               "String.class", "goType", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	DbfluteDbm.ColumnVersionNo = df.CCI(&dbflute, "version_no", versionNoSqlName, "", "",
               "Integer.class", "versionNo", "", false, false,true, "int4", 10, 0,
               "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")
	delFlagSqlName := new(df.ColumnSqlName)
	delFlagSqlName.ColumnSqlName = "del_flag"
	delFlagSqlName.IrregularChar = false
	DbfluteDbm.ColumnDelFlag = df.CCI(&dbflute, "del_flag", delFlagSqlName, "", "",
               "Integer.class", "delFlag", "", false, false,true, "int4", 10, 0,
               "0",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	DbfluteDbm.ColumnRegisterDatetime = df.CCI(&dbflute, "register_datetime", registerDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	DbfluteDbm.ColumnRegisterUser = df.CCI(&dbflute, "register_user", registerUserSqlName, "", "",
               "String.class", "registerUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	DbfluteDbm.ColumnRegisterProcess = df.CCI(&dbflute, "register_process", registerProcessSqlName, "", "",
               "String.class", "registerProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	DbfluteDbm.ColumnUpdateDatetime = df.CCI(&dbflute, "update_datetime", updateDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	DbfluteDbm.ColumnUpdateUser = df.CCI(&dbflute, "update_user", updateUserSqlName, "", "",
               "String.class", "updateUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	DbfluteDbm.ColumnUpdateProcess = df.CCI(&dbflute, "update_process", updateProcessSqlName, "", "",
               "String.class", "updateProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")

	DbfluteDbm.ColumnInfoList = new(df.List)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnId)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnLoginId)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnDb)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnDbName)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnDbTable)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnName)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnDbType)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnRequired)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnJavaType)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnGoType)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnVersionNo)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnDelFlag)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnRegisterDatetime)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnRegisterUser)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnRegisterProcess)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnUpdateDatetime)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnUpdateUser)
	DbfluteDbm.ColumnInfoList.Add(DbfluteDbm.ColumnUpdateProcess)


	DbfluteDbm.ColumnInfoMap=make(map[string]int)
	DbfluteDbm.ColumnInfoMap["id"]=0
		DbfluteDbm.ColumnInfoMap["loginId"]=1
		DbfluteDbm.ColumnInfoMap["db"]=2
		DbfluteDbm.ColumnInfoMap["dbName"]=3
		DbfluteDbm.ColumnInfoMap["dbTable"]=4
		DbfluteDbm.ColumnInfoMap["name"]=5
		DbfluteDbm.ColumnInfoMap["dbType"]=6
		DbfluteDbm.ColumnInfoMap["required"]=7
		DbfluteDbm.ColumnInfoMap["javaType"]=8
		DbfluteDbm.ColumnInfoMap["goType"]=9
		DbfluteDbm.ColumnInfoMap["versionNo"]=10
		DbfluteDbm.ColumnInfoMap["delFlag"]=11
		DbfluteDbm.ColumnInfoMap["registerDatetime"]=12
		DbfluteDbm.ColumnInfoMap["registerUser"]=13
		DbfluteDbm.ColumnInfoMap["registerProcess"]=14
		DbfluteDbm.ColumnInfoMap["updateDatetime"]=15
		DbfluteDbm.ColumnInfoMap["updateUser"]=16
		DbfluteDbm.ColumnInfoMap["updateProcess"]=17
	    DbfluteDbm.PrimaryKey = true
    DbfluteDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &dbflute
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(DbfluteDbm.ColumnId)

	DbfluteDbm.PrimaryInfo = new(df.PrimaryInfo)
	DbfluteDbm.PrimaryInfo.UniqueInfo = ui
	
	DbfluteDbm.SequenceFlag = true
	DbfluteDbm.SequenceName = "dbflute_id_seq"
	DbfluteDbm.SequenceIncrementSize = 1
	DbfluteDbm.SequenceCacheSize = 0
	DbfluteDbm.VersionNoFlag = true
	DbfluteDbm.VersionNoColumnInfo = DbfluteDbm.ColumnVersionNo
	
	var dbfluteMeta df.DBMeta = DbfluteDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Dbflute"] = &dbfluteMeta
}

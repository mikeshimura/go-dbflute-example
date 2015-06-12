package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type UserTableDbm_T struct {
	df.BaseDBMeta
	ColumnId *df.ColumnInfo
	ColumnTableName *df.ColumnInfo
	ColumnKey1 *df.ColumnInfo
	ColumnKey2 *df.ColumnInfo
	ColumnS1Data *df.ColumnInfo
	ColumnS2Data *df.ColumnInfo
	ColumnS3Data *df.ColumnInfo
	ColumnN1Data *df.ColumnInfo
	ColumnN2Data *df.ColumnInfo
	ColumnN3Data *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
	ColumnDelFlag *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
}

func (b *UserTableDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *UserTableDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
}

func (b *UserTableDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var UserTableDbm *UserTableDbm_T

func Create_UserTableDbm() {
	UserTableDbm = new(UserTableDbm_T)
	UserTableDbm.TableDbName = "user_table"
	UserTableDbm.TableDispName = "user_table"
	UserTableDbm.TablePropertyName = "userTable"
	UserTableDbm.TableSqlName = new(df.TableSqlName)
	UserTableDbm.TableSqlName.TableSqlName = "user_table"
	UserTableDbm.TableSqlName.CorrespondingDbName = UserTableDbm.TableDbName

	var userTable df.DBMeta
	userTable = UserTableDbm
	UserTableDbm.DBMeta=&userTable
	idSqlName := new(df.ColumnSqlName)
	idSqlName.ColumnSqlName = "id"
	idSqlName.IrregularChar = false
	UserTableDbm.ColumnId = df.CCI(&userTable, "id", idSqlName, "", "",
               "Integer.class", "id", "", true, true,true, "serial", 10, 0,
               "nextval('user_table_id_seq'::regclass)",false,"","", "","employeeList","",false,"int64")
	tableNameSqlName := new(df.ColumnSqlName)
	tableNameSqlName.ColumnSqlName = "table_name"
	tableNameSqlName.IrregularChar = false
	UserTableDbm.ColumnTableName = df.CCI(&userTable, "table_name", tableNameSqlName, "", "",
               "String.class", "tableName", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	key1SqlName := new(df.ColumnSqlName)
	key1SqlName.ColumnSqlName = "key_1"
	key1SqlName.IrregularChar = false
	UserTableDbm.ColumnKey1 = df.CCI(&userTable, "key_1", key1SqlName, "", "",
               "String.class", "key1", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	key2SqlName := new(df.ColumnSqlName)
	key2SqlName.ColumnSqlName = "key_2"
	key2SqlName.IrregularChar = false
	UserTableDbm.ColumnKey2 = df.CCI(&userTable, "key_2", key2SqlName, "", "",
               "String.class", "key2", "", false, false,true, "varchar", 100, 0,
               "",false,"","", "","","",false,"string")
	s1DataSqlName := new(df.ColumnSqlName)
	s1DataSqlName.ColumnSqlName = "s1_data"
	s1DataSqlName.IrregularChar = false
	UserTableDbm.ColumnS1Data = df.CCI(&userTable, "s1_data", s1DataSqlName, "", "",
               "String.class", "s1Data", "", false, false,false, "text", 2147483647, 0,
               "",false,"","", "","","",false,"sql.NullString")
	s2DataSqlName := new(df.ColumnSqlName)
	s2DataSqlName.ColumnSqlName = "s2_data"
	s2DataSqlName.IrregularChar = false
	UserTableDbm.ColumnS2Data = df.CCI(&userTable, "s2_data", s2DataSqlName, "", "",
               "String.class", "s2Data", "", false, false,false, "varchar", 100, 0,
               "",false,"","", "","","",false,"sql.NullString")
	s3DataSqlName := new(df.ColumnSqlName)
	s3DataSqlName.ColumnSqlName = "s3_data"
	s3DataSqlName.IrregularChar = false
	UserTableDbm.ColumnS3Data = df.CCI(&userTable, "s3_data", s3DataSqlName, "", "",
               "String.class", "s3Data", "", false, false,false, "varchar", 100, 0,
               "",false,"","", "","","",false,"sql.NullString")
	n1DataSqlName := new(df.ColumnSqlName)
	n1DataSqlName.ColumnSqlName = "n1_data"
	n1DataSqlName.IrregularChar = false
	UserTableDbm.ColumnN1Data = df.CCI(&userTable, "n1_data", n1DataSqlName, "", "",
               "java.math.BigDecimal.class", "n1Data", "", false, false,false, "numeric", 20, 2,
               "",false,"","", "","","",false,"df.NullNumeric")
	n2DataSqlName := new(df.ColumnSqlName)
	n2DataSqlName.ColumnSqlName = "n2_data"
	n2DataSqlName.IrregularChar = false
	UserTableDbm.ColumnN2Data = df.CCI(&userTable, "n2_data", n2DataSqlName, "", "",
               "java.math.BigDecimal.class", "n2Data", "", false, false,false, "numeric", 20, 2,
               "",false,"","", "","","",false,"df.NullNumeric")
	n3DataSqlName := new(df.ColumnSqlName)
	n3DataSqlName.ColumnSqlName = "n3_data"
	n3DataSqlName.IrregularChar = false
	UserTableDbm.ColumnN3Data = df.CCI(&userTable, "n3_data", n3DataSqlName, "", "",
               "java.math.BigDecimal.class", "n3Data", "", false, false,false, "numeric", 20, 2,
               "",false,"","", "","","",false,"df.NullNumeric")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	UserTableDbm.ColumnVersionNo = df.CCI(&userTable, "version_no", versionNoSqlName, "", "",
               "Integer.class", "versionNo", "", false, false,true, "int4", 10, 0,
               "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")
	delFlagSqlName := new(df.ColumnSqlName)
	delFlagSqlName.ColumnSqlName = "del_flag"
	delFlagSqlName.IrregularChar = false
	UserTableDbm.ColumnDelFlag = df.CCI(&userTable, "del_flag", delFlagSqlName, "", "",
               "Integer.class", "delFlag", "", false, false,true, "int4", 10, 0,
               "0",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	UserTableDbm.ColumnRegisterDatetime = df.CCI(&userTable, "register_datetime", registerDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	UserTableDbm.ColumnRegisterUser = df.CCI(&userTable, "register_user", registerUserSqlName, "", "",
               "String.class", "registerUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	UserTableDbm.ColumnRegisterProcess = df.CCI(&userTable, "register_process", registerProcessSqlName, "", "",
               "String.class", "registerProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	UserTableDbm.ColumnUpdateDatetime = df.CCI(&userTable, "update_datetime", updateDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	UserTableDbm.ColumnUpdateUser = df.CCI(&userTable, "update_user", updateUserSqlName, "", "",
               "String.class", "updateUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	UserTableDbm.ColumnUpdateProcess = df.CCI(&userTable, "update_process", updateProcessSqlName, "", "",
               "String.class", "updateProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")

	UserTableDbm.ColumnInfoList = new(df.List)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnId)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnTableName)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnKey1)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnKey2)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnS1Data)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnS2Data)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnS3Data)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnN1Data)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnN2Data)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnN3Data)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnVersionNo)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnDelFlag)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnRegisterDatetime)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnRegisterUser)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnRegisterProcess)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnUpdateDatetime)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnUpdateUser)
	UserTableDbm.ColumnInfoList.Add(UserTableDbm.ColumnUpdateProcess)


	UserTableDbm.ColumnInfoMap=make(map[string]int)
	UserTableDbm.ColumnInfoMap["id"]=0
		UserTableDbm.ColumnInfoMap["tableName"]=1
		UserTableDbm.ColumnInfoMap["key1"]=2
		UserTableDbm.ColumnInfoMap["key2"]=3
		UserTableDbm.ColumnInfoMap["s1Data"]=4
		UserTableDbm.ColumnInfoMap["s2Data"]=5
		UserTableDbm.ColumnInfoMap["s3Data"]=6
		UserTableDbm.ColumnInfoMap["n1Data"]=7
		UserTableDbm.ColumnInfoMap["n2Data"]=8
		UserTableDbm.ColumnInfoMap["n3Data"]=9
		UserTableDbm.ColumnInfoMap["versionNo"]=10
		UserTableDbm.ColumnInfoMap["delFlag"]=11
		UserTableDbm.ColumnInfoMap["registerDatetime"]=12
		UserTableDbm.ColumnInfoMap["registerUser"]=13
		UserTableDbm.ColumnInfoMap["registerProcess"]=14
		UserTableDbm.ColumnInfoMap["updateDatetime"]=15
		UserTableDbm.ColumnInfoMap["updateUser"]=16
		UserTableDbm.ColumnInfoMap["updateProcess"]=17
	    UserTableDbm.PrimaryKey = true
    UserTableDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &userTable
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(UserTableDbm.ColumnId)

	UserTableDbm.PrimaryInfo = new(df.PrimaryInfo)
	UserTableDbm.PrimaryInfo.UniqueInfo = ui
	
	UserTableDbm.SequenceFlag = true
	UserTableDbm.SequenceName = "user_table_id_seq"
	UserTableDbm.SequenceIncrementSize = 1
	UserTableDbm.SequenceCacheSize = 0
	UserTableDbm.VersionNoFlag = true
	UserTableDbm.VersionNoColumnInfo = UserTableDbm.ColumnVersionNo
	
	var userTableMeta df.DBMeta = UserTableDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["UserTable"] = &userTableMeta
}

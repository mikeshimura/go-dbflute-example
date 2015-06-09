package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type SysTableDbm_T struct {
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

func (b *SysTableDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *SysTableDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
}

func (b *SysTableDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var SysTableDbm *SysTableDbm_T

func Create_SysTableDbm() {
	SysTableDbm = new(SysTableDbm_T)
	SysTableDbm.TableDbName = "sys_table"
	SysTableDbm.TableDispName = "sys_table"
	SysTableDbm.TablePropertyName = "sysTable"
	SysTableDbm.TableSqlName = new(df.TableSqlName)
	SysTableDbm.TableSqlName.TableSqlName = "sys_table"
	SysTableDbm.TableSqlName.CorrespondingDbName = SysTableDbm.TableDbName

	var sysTable df.DBMeta
	sysTable = SysTableDbm
	SysTableDbm.DBMeta=&sysTable
	idSqlName := new(df.ColumnSqlName)
	idSqlName.ColumnSqlName = "id"
	idSqlName.IrregularChar = false
	SysTableDbm.ColumnId = df.CCI(&sysTable, "id", idSqlName, "", "",
               "Integer.class", "id", "", true, true,true, "serial", 10, 0,
               "nextval('sys_table_id_seq'::regclass)",false,"","", "","","",false,"int64")
	tableNameSqlName := new(df.ColumnSqlName)
	tableNameSqlName.ColumnSqlName = "table_name"
	tableNameSqlName.IrregularChar = false
	SysTableDbm.ColumnTableName = df.CCI(&sysTable, "table_name", tableNameSqlName, "", "",
               "String.class", "tableName", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	key1SqlName := new(df.ColumnSqlName)
	key1SqlName.ColumnSqlName = "key_1"
	key1SqlName.IrregularChar = false
	SysTableDbm.ColumnKey1 = df.CCI(&sysTable, "key_1", key1SqlName, "", "",
               "String.class", "key1", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	key2SqlName := new(df.ColumnSqlName)
	key2SqlName.ColumnSqlName = "key_2"
	key2SqlName.IrregularChar = false
	SysTableDbm.ColumnKey2 = df.CCI(&sysTable, "key_2", key2SqlName, "", "",
               "String.class", "key2", "", false, false,true, "varchar", 100, 0,
               "",false,"","", "","","",false,"string")
	s1DataSqlName := new(df.ColumnSqlName)
	s1DataSqlName.ColumnSqlName = "s1_data"
	s1DataSqlName.IrregularChar = false
	SysTableDbm.ColumnS1Data = df.CCI(&sysTable, "s1_data", s1DataSqlName, "", "",
               "String.class", "s1Data", "", false, false,false, "text", 2147483647, 0,
               "",false,"","", "","","",false,"sql.NullString")
	s2DataSqlName := new(df.ColumnSqlName)
	s2DataSqlName.ColumnSqlName = "s2_data"
	s2DataSqlName.IrregularChar = false
	SysTableDbm.ColumnS2Data = df.CCI(&sysTable, "s2_data", s2DataSqlName, "", "",
               "String.class", "s2Data", "", false, false,false, "varchar", 100, 0,
               "",false,"","", "","","",false,"sql.NullString")
	s3DataSqlName := new(df.ColumnSqlName)
	s3DataSqlName.ColumnSqlName = "s3_data"
	s3DataSqlName.IrregularChar = false
	SysTableDbm.ColumnS3Data = df.CCI(&sysTable, "s3_data", s3DataSqlName, "", "",
               "String.class", "s3Data", "", false, false,false, "varchar", 100, 0,
               "",false,"","", "","","",false,"sql.NullString")
	n1DataSqlName := new(df.ColumnSqlName)
	n1DataSqlName.ColumnSqlName = "n1_data"
	n1DataSqlName.IrregularChar = false
	SysTableDbm.ColumnN1Data = df.CCI(&sysTable, "n1_data", n1DataSqlName, "", "",
               "java.math.BigDecimal.class", "n1Data", "", false, false,false, "numeric", 20, 2,
               "",false,"","", "","","",false,"df.NullNumeric")
	n2DataSqlName := new(df.ColumnSqlName)
	n2DataSqlName.ColumnSqlName = "n2_data"
	n2DataSqlName.IrregularChar = false
	SysTableDbm.ColumnN2Data = df.CCI(&sysTable, "n2_data", n2DataSqlName, "", "",
               "java.math.BigDecimal.class", "n2Data", "", false, false,false, "numeric", 20, 2,
               "",false,"","", "","","",false,"df.NullNumeric")
	n3DataSqlName := new(df.ColumnSqlName)
	n3DataSqlName.ColumnSqlName = "n3_data"
	n3DataSqlName.IrregularChar = false
	SysTableDbm.ColumnN3Data = df.CCI(&sysTable, "n3_data", n3DataSqlName, "", "",
               "java.math.BigDecimal.class", "n3Data", "", false, false,false, "numeric", 20, 2,
               "",false,"","", "","","",false,"df.NullNumeric")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	SysTableDbm.ColumnVersionNo = df.CCI(&sysTable, "version_no", versionNoSqlName, "", "",
               "Integer.class", "versionNo", "", false, false,true, "int4", 10, 0,
               "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")
	delFlagSqlName := new(df.ColumnSqlName)
	delFlagSqlName.ColumnSqlName = "del_flag"
	delFlagSqlName.IrregularChar = false
	SysTableDbm.ColumnDelFlag = df.CCI(&sysTable, "del_flag", delFlagSqlName, "", "",
               "Integer.class", "delFlag", "", false, false,true, "int4", 10, 0,
               "0",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	SysTableDbm.ColumnRegisterDatetime = df.CCI(&sysTable, "register_datetime", registerDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	SysTableDbm.ColumnRegisterUser = df.CCI(&sysTable, "register_user", registerUserSqlName, "", "",
               "String.class", "registerUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	SysTableDbm.ColumnRegisterProcess = df.CCI(&sysTable, "register_process", registerProcessSqlName, "", "",
               "String.class", "registerProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	SysTableDbm.ColumnUpdateDatetime = df.CCI(&sysTable, "update_datetime", updateDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	SysTableDbm.ColumnUpdateUser = df.CCI(&sysTable, "update_user", updateUserSqlName, "", "",
               "String.class", "updateUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	SysTableDbm.ColumnUpdateProcess = df.CCI(&sysTable, "update_process", updateProcessSqlName, "", "",
               "String.class", "updateProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")

	SysTableDbm.ColumnInfoList = new(df.List)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnId)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnTableName)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnKey1)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnKey2)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnS1Data)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnS2Data)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnS3Data)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnN1Data)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnN2Data)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnN3Data)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnVersionNo)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnDelFlag)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnRegisterDatetime)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnRegisterUser)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnRegisterProcess)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnUpdateDatetime)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnUpdateUser)
	SysTableDbm.ColumnInfoList.Add(SysTableDbm.ColumnUpdateProcess)


	SysTableDbm.ColumnInfoMap=make(map[string]int)
	SysTableDbm.ColumnInfoMap["id"]=0
		SysTableDbm.ColumnInfoMap["tableName"]=1
		SysTableDbm.ColumnInfoMap["key1"]=2
		SysTableDbm.ColumnInfoMap["key2"]=3
		SysTableDbm.ColumnInfoMap["s1Data"]=4
		SysTableDbm.ColumnInfoMap["s2Data"]=5
		SysTableDbm.ColumnInfoMap["s3Data"]=6
		SysTableDbm.ColumnInfoMap["n1Data"]=7
		SysTableDbm.ColumnInfoMap["n2Data"]=8
		SysTableDbm.ColumnInfoMap["n3Data"]=9
		SysTableDbm.ColumnInfoMap["versionNo"]=10
		SysTableDbm.ColumnInfoMap["delFlag"]=11
		SysTableDbm.ColumnInfoMap["registerDatetime"]=12
		SysTableDbm.ColumnInfoMap["registerUser"]=13
		SysTableDbm.ColumnInfoMap["registerProcess"]=14
		SysTableDbm.ColumnInfoMap["updateDatetime"]=15
		SysTableDbm.ColumnInfoMap["updateUser"]=16
		SysTableDbm.ColumnInfoMap["updateProcess"]=17
	    SysTableDbm.PrimaryKey = true
    SysTableDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &sysTable
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(SysTableDbm.ColumnId)

	SysTableDbm.PrimaryInfo = new(df.PrimaryInfo)
	SysTableDbm.PrimaryInfo.UniqueInfo = ui
	
	SysTableDbm.SequenceFlag = true
	SysTableDbm.SequenceName = "sys_table_id_seq"
	SysTableDbm.SequenceIncrementSize = 1
	SysTableDbm.SequenceCacheSize = 0
	SysTableDbm.VersionNoFlag = true
	SysTableDbm.VersionNoColumnInfo = SysTableDbm.ColumnVersionNo
	
	var sysTableMeta df.DBMeta = SysTableDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["SysTable"] = &sysTableMeta
}

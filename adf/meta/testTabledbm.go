package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type TestTableDbm_T struct {
	df.BaseDBMeta
	ColumnId *df.ColumnInfo
	ColumnTestId *df.ColumnInfo
	ColumnTestDate *df.ColumnInfo
	ColumnTestTimestamp *df.ColumnInfo
	ColumnTestNbr *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
	ColumnDelFlag *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
}

func (b *TestTableDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *TestTableDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
}

func (b *TestTableDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var TestTableDbm *TestTableDbm_T

func Create_TestTableDbm() {
	TestTableDbm = new(TestTableDbm_T)
	TestTableDbm.TableDbName = "test_table"
	TestTableDbm.TableDispName = "test_table"
	TestTableDbm.TablePropertyName = "testTable"
	TestTableDbm.TableSqlName = new(df.TableSqlName)
	TestTableDbm.TableSqlName.TableSqlName = "test_table"
	TestTableDbm.TableSqlName.CorrespondingDbName = TestTableDbm.TableDbName

	var testTable df.DBMeta
	testTable = TestTableDbm
	TestTableDbm.DBMeta=&testTable
	idSqlName := new(df.ColumnSqlName)
	idSqlName.ColumnSqlName = "id"
	idSqlName.IrregularChar = false
	TestTableDbm.ColumnId = df.CCI(&testTable, "id", idSqlName, "", "",
               "Integer.class", "id", "", true, true,true, "serial", 10, 0,
               "nextval('test_table_id_seq'::regclass)",false,"","", "","","",false,"int64")
	testIdSqlName := new(df.ColumnSqlName)
	testIdSqlName.ColumnSqlName = "test_id"
	testIdSqlName.IrregularChar = false
	TestTableDbm.ColumnTestId = df.CCI(&testTable, "test_id", testIdSqlName, "", "",
               "String.class", "testId", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	testDateSqlName := new(df.ColumnSqlName)
	testDateSqlName.ColumnSqlName = "test_date"
	testDateSqlName.IrregularChar = false
	TestTableDbm.ColumnTestDate = df.CCI(&testTable, "test_date", testDateSqlName, "", "",
               "java.time.LocalDate.class", "testDate", "", false, false,false, "date", 13, 0,
               "",false,"","", "","","",false,"df.NullDate")
	testTimestampSqlName := new(df.ColumnSqlName)
	testTimestampSqlName.ColumnSqlName = "test_timestamp"
	testTimestampSqlName.IrregularChar = false
	TestTableDbm.ColumnTestTimestamp = df.CCI(&testTable, "test_timestamp", testTimestampSqlName, "", "",
               "java.time.LocalDateTime.class", "testTimestamp", "", false, false,false, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.NullTimestamp")
	testNbrSqlName := new(df.ColumnSqlName)
	testNbrSqlName.ColumnSqlName = "test_nbr"
	testNbrSqlName.IrregularChar = false
	TestTableDbm.ColumnTestNbr = df.CCI(&testTable, "test_nbr", testNbrSqlName, "", "",
               "Integer.class", "testNbr", "", false, false,false, "int4", 10, 0,
               "",false,"","", "","","",false,"sql.NullInt64")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	TestTableDbm.ColumnVersionNo = df.CCI(&testTable, "version_no", versionNoSqlName, "", "",
               "Integer.class", "versionNo", "", false, false,true, "int4", 10, 0,
               "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")
	delFlagSqlName := new(df.ColumnSqlName)
	delFlagSqlName.ColumnSqlName = "del_flag"
	delFlagSqlName.IrregularChar = false
	TestTableDbm.ColumnDelFlag = df.CCI(&testTable, "del_flag", delFlagSqlName, "", "",
               "Integer.class", "delFlag", "", false, false,true, "int4", 10, 0,
               "0",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	TestTableDbm.ColumnRegisterDatetime = df.CCI(&testTable, "register_datetime", registerDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	TestTableDbm.ColumnRegisterUser = df.CCI(&testTable, "register_user", registerUserSqlName, "", "",
               "String.class", "registerUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	TestTableDbm.ColumnRegisterProcess = df.CCI(&testTable, "register_process", registerProcessSqlName, "", "",
               "String.class", "registerProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	TestTableDbm.ColumnUpdateDatetime = df.CCI(&testTable, "update_datetime", updateDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	TestTableDbm.ColumnUpdateUser = df.CCI(&testTable, "update_user", updateUserSqlName, "", "",
               "String.class", "updateUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	TestTableDbm.ColumnUpdateProcess = df.CCI(&testTable, "update_process", updateProcessSqlName, "", "",
               "String.class", "updateProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")

	TestTableDbm.ColumnInfoList = new(df.List)
	TestTableDbm.ColumnInfoList.Add(TestTableDbm.ColumnId)
	TestTableDbm.ColumnInfoList.Add(TestTableDbm.ColumnTestId)
	TestTableDbm.ColumnInfoList.Add(TestTableDbm.ColumnTestDate)
	TestTableDbm.ColumnInfoList.Add(TestTableDbm.ColumnTestTimestamp)
	TestTableDbm.ColumnInfoList.Add(TestTableDbm.ColumnTestNbr)
	TestTableDbm.ColumnInfoList.Add(TestTableDbm.ColumnVersionNo)
	TestTableDbm.ColumnInfoList.Add(TestTableDbm.ColumnDelFlag)
	TestTableDbm.ColumnInfoList.Add(TestTableDbm.ColumnRegisterDatetime)
	TestTableDbm.ColumnInfoList.Add(TestTableDbm.ColumnRegisterUser)
	TestTableDbm.ColumnInfoList.Add(TestTableDbm.ColumnRegisterProcess)
	TestTableDbm.ColumnInfoList.Add(TestTableDbm.ColumnUpdateDatetime)
	TestTableDbm.ColumnInfoList.Add(TestTableDbm.ColumnUpdateUser)
	TestTableDbm.ColumnInfoList.Add(TestTableDbm.ColumnUpdateProcess)


	TestTableDbm.ColumnInfoMap=make(map[string]int)
	TestTableDbm.ColumnInfoMap["id"]=0
		TestTableDbm.ColumnInfoMap["testId"]=1
		TestTableDbm.ColumnInfoMap["testDate"]=2
		TestTableDbm.ColumnInfoMap["testTimestamp"]=3
		TestTableDbm.ColumnInfoMap["testNbr"]=4
		TestTableDbm.ColumnInfoMap["versionNo"]=5
		TestTableDbm.ColumnInfoMap["delFlag"]=6
		TestTableDbm.ColumnInfoMap["registerDatetime"]=7
		TestTableDbm.ColumnInfoMap["registerUser"]=8
		TestTableDbm.ColumnInfoMap["registerProcess"]=9
		TestTableDbm.ColumnInfoMap["updateDatetime"]=10
		TestTableDbm.ColumnInfoMap["updateUser"]=11
		TestTableDbm.ColumnInfoMap["updateProcess"]=12
	    TestTableDbm.PrimaryKey = true
    TestTableDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &testTable
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(TestTableDbm.ColumnId)

	TestTableDbm.PrimaryInfo = new(df.PrimaryInfo)
	TestTableDbm.PrimaryInfo.UniqueInfo = ui
	
	TestTableDbm.SequenceFlag = true
	TestTableDbm.SequenceName = "test_table_id_seq"
	TestTableDbm.SequenceIncrementSize = 1
	TestTableDbm.SequenceCacheSize = 0
	TestTableDbm.VersionNoFlag = true
	TestTableDbm.VersionNoColumnInfo = TestTableDbm.ColumnVersionNo
	
	var testTableMeta df.DBMeta = TestTableDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["TestTable"] = &testTableMeta
}

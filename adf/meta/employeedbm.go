package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type EmployeeDbm_T struct {
	df.BaseDBMeta
	ColumnId *df.ColumnInfo
	ColumnEmpCd *df.ColumnInfo
	ColumnSecId *df.ColumnInfo
	ColumnName *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
	ColumnDelFlag *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
}

func (b *EmployeeDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *EmployeeDbm_T) foreignUserTable() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		EmployeeDbm.GetColumnInfoByPropertyName("secId"),
		UserTableDbm.GetColumnInfoByPropertyName("id"),
	}

	return b.BaseDBMeta.Cfi("employee_sec_id_fkey", "UserTable",
		columns, 0, false, false, false, false,
		"", nil, false, "employeeList")
}	
func (b *EmployeeDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
	b.ForeignInfoMap["UserTable"] = b.foreignUserTable()
}

func (b *EmployeeDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var EmployeeDbm *EmployeeDbm_T

func Create_EmployeeDbm() {
	EmployeeDbm = new(EmployeeDbm_T)
	EmployeeDbm.TableDbName = "employee"
	EmployeeDbm.TableDispName = "employee"
	EmployeeDbm.TablePropertyName = "employee"
	EmployeeDbm.TableSqlName = new(df.TableSqlName)
	EmployeeDbm.TableSqlName.TableSqlName = "employee"
	EmployeeDbm.TableSqlName.CorrespondingDbName = EmployeeDbm.TableDbName

	var employee df.DBMeta
	employee = EmployeeDbm
	EmployeeDbm.DBMeta=&employee
	idSqlName := new(df.ColumnSqlName)
	idSqlName.ColumnSqlName = "id"
	idSqlName.IrregularChar = false
	EmployeeDbm.ColumnId = df.CCI(&employee, "id", idSqlName, "", "",
               "Integer.class", "id", "", true, true,true, "serial", 10, 0,
               "nextval('employee_id_seq'::regclass)",false,"","", "","salesSlipList","",false,"int64")
	empCdSqlName := new(df.ColumnSqlName)
	empCdSqlName.ColumnSqlName = "emp_cd"
	empCdSqlName.IrregularChar = false
	EmployeeDbm.ColumnEmpCd = df.CCI(&employee, "emp_cd", empCdSqlName, "", "",
               "String.class", "empCd", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	secIdSqlName := new(df.ColumnSqlName)
	secIdSqlName.ColumnSqlName = "sec_id"
	secIdSqlName.IrregularChar = false
	EmployeeDbm.ColumnSecId = df.CCI(&employee, "sec_id", secIdSqlName, "", "",
               "Integer.class", "secId", "", false, false,true, "int4", 10, 0,
               "",false,"","", "userTable","","",false,"int64")
	nameSqlName := new(df.ColumnSqlName)
	nameSqlName.ColumnSqlName = "name"
	nameSqlName.IrregularChar = false
	EmployeeDbm.ColumnName = df.CCI(&employee, "name", nameSqlName, "", "",
               "String.class", "name", "", false, false,true, "varchar", 100, 0,
               "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	EmployeeDbm.ColumnVersionNo = df.CCI(&employee, "version_no", versionNoSqlName, "", "",
               "Integer.class", "versionNo", "", false, false,true, "int4", 10, 0,
               "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")
	delFlagSqlName := new(df.ColumnSqlName)
	delFlagSqlName.ColumnSqlName = "del_flag"
	delFlagSqlName.IrregularChar = false
	EmployeeDbm.ColumnDelFlag = df.CCI(&employee, "del_flag", delFlagSqlName, "", "",
               "Integer.class", "delFlag", "", false, false,true, "int4", 10, 0,
               "0",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	EmployeeDbm.ColumnRegisterDatetime = df.CCI(&employee, "register_datetime", registerDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	EmployeeDbm.ColumnRegisterUser = df.CCI(&employee, "register_user", registerUserSqlName, "", "",
               "String.class", "registerUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	EmployeeDbm.ColumnRegisterProcess = df.CCI(&employee, "register_process", registerProcessSqlName, "", "",
               "String.class", "registerProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	EmployeeDbm.ColumnUpdateDatetime = df.CCI(&employee, "update_datetime", updateDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	EmployeeDbm.ColumnUpdateUser = df.CCI(&employee, "update_user", updateUserSqlName, "", "",
               "String.class", "updateUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	EmployeeDbm.ColumnUpdateProcess = df.CCI(&employee, "update_process", updateProcessSqlName, "", "",
               "String.class", "updateProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")

	EmployeeDbm.ColumnInfoList = new(df.List)
	EmployeeDbm.ColumnInfoList.Add(EmployeeDbm.ColumnId)
	EmployeeDbm.ColumnInfoList.Add(EmployeeDbm.ColumnEmpCd)
	EmployeeDbm.ColumnInfoList.Add(EmployeeDbm.ColumnSecId)
	EmployeeDbm.ColumnInfoList.Add(EmployeeDbm.ColumnName)
	EmployeeDbm.ColumnInfoList.Add(EmployeeDbm.ColumnVersionNo)
	EmployeeDbm.ColumnInfoList.Add(EmployeeDbm.ColumnDelFlag)
	EmployeeDbm.ColumnInfoList.Add(EmployeeDbm.ColumnRegisterDatetime)
	EmployeeDbm.ColumnInfoList.Add(EmployeeDbm.ColumnRegisterUser)
	EmployeeDbm.ColumnInfoList.Add(EmployeeDbm.ColumnRegisterProcess)
	EmployeeDbm.ColumnInfoList.Add(EmployeeDbm.ColumnUpdateDatetime)
	EmployeeDbm.ColumnInfoList.Add(EmployeeDbm.ColumnUpdateUser)
	EmployeeDbm.ColumnInfoList.Add(EmployeeDbm.ColumnUpdateProcess)


	EmployeeDbm.ColumnInfoMap=make(map[string]int)
	EmployeeDbm.ColumnInfoMap["id"]=0
		EmployeeDbm.ColumnInfoMap["empCd"]=1
		EmployeeDbm.ColumnInfoMap["secId"]=2
		EmployeeDbm.ColumnInfoMap["name"]=3
		EmployeeDbm.ColumnInfoMap["versionNo"]=4
		EmployeeDbm.ColumnInfoMap["delFlag"]=5
		EmployeeDbm.ColumnInfoMap["registerDatetime"]=6
		EmployeeDbm.ColumnInfoMap["registerUser"]=7
		EmployeeDbm.ColumnInfoMap["registerProcess"]=8
		EmployeeDbm.ColumnInfoMap["updateDatetime"]=9
		EmployeeDbm.ColumnInfoMap["updateUser"]=10
		EmployeeDbm.ColumnInfoMap["updateProcess"]=11
	    EmployeeDbm.PrimaryKey = true
    EmployeeDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &employee
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(EmployeeDbm.ColumnId)

	EmployeeDbm.PrimaryInfo = new(df.PrimaryInfo)
	EmployeeDbm.PrimaryInfo.UniqueInfo = ui
	
	EmployeeDbm.SequenceFlag = true
	EmployeeDbm.SequenceName = "employee_id_seq"
	EmployeeDbm.SequenceIncrementSize = 1
	EmployeeDbm.SequenceCacheSize = 0
	EmployeeDbm.VersionNoFlag = true
	EmployeeDbm.VersionNoColumnInfo = EmployeeDbm.ColumnVersionNo
	
	var employeeMeta df.DBMeta = EmployeeDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Employee"] = &employeeMeta
}

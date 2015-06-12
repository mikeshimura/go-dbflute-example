package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type LoginDbm_T struct {
	df.BaseDBMeta
	ColumnId *df.ColumnInfo
	ColumnLoginId *df.ColumnInfo
	ColumnName *df.ColumnInfo
	ColumnPassword *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
	ColumnDelFlag *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
}

func (b *LoginDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *LoginDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
}

func (b *LoginDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var LoginDbm *LoginDbm_T

func Create_LoginDbm() {
	LoginDbm = new(LoginDbm_T)
	LoginDbm.TableDbName = "login"
	LoginDbm.TableDispName = "login"
	LoginDbm.TablePropertyName = "login"
	LoginDbm.TableSqlName = new(df.TableSqlName)
	LoginDbm.TableSqlName.TableSqlName = "login"
	LoginDbm.TableSqlName.CorrespondingDbName = LoginDbm.TableDbName

	var login df.DBMeta
	login = LoginDbm
	LoginDbm.DBMeta=&login
	idSqlName := new(df.ColumnSqlName)
	idSqlName.ColumnSqlName = "id"
	idSqlName.IrregularChar = false
	LoginDbm.ColumnId = df.CCI(&login, "id", idSqlName, "", "",
               "Integer.class", "id", "", true, true,true, "serial", 10, 0,
               "nextval('login_id_seq'::regclass)",false,"","", "","sessionList","",false,"int64")
	loginIdSqlName := new(df.ColumnSqlName)
	loginIdSqlName.ColumnSqlName = "login_id"
	loginIdSqlName.IrregularChar = false
	LoginDbm.ColumnLoginId = df.CCI(&login, "login_id", loginIdSqlName, "", "",
               "String.class", "loginId", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	nameSqlName := new(df.ColumnSqlName)
	nameSqlName.ColumnSqlName = "name"
	nameSqlName.IrregularChar = false
	LoginDbm.ColumnName = df.CCI(&login, "name", nameSqlName, "", "",
               "String.class", "name", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	passwordSqlName := new(df.ColumnSqlName)
	passwordSqlName.ColumnSqlName = "password"
	passwordSqlName.IrregularChar = false
	LoginDbm.ColumnPassword = df.CCI(&login, "password", passwordSqlName, "", "",
               "String.class", "password", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	LoginDbm.ColumnVersionNo = df.CCI(&login, "version_no", versionNoSqlName, "", "",
               "Integer.class", "versionNo", "", false, false,true, "int4", 10, 0,
               "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")
	delFlagSqlName := new(df.ColumnSqlName)
	delFlagSqlName.ColumnSqlName = "del_flag"
	delFlagSqlName.IrregularChar = false
	LoginDbm.ColumnDelFlag = df.CCI(&login, "del_flag", delFlagSqlName, "", "",
               "Integer.class", "delFlag", "", false, false,true, "int4", 10, 0,
               "0",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	LoginDbm.ColumnRegisterDatetime = df.CCI(&login, "register_datetime", registerDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	LoginDbm.ColumnRegisterUser = df.CCI(&login, "register_user", registerUserSqlName, "", "",
               "String.class", "registerUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	LoginDbm.ColumnRegisterProcess = df.CCI(&login, "register_process", registerProcessSqlName, "", "",
               "String.class", "registerProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	LoginDbm.ColumnUpdateDatetime = df.CCI(&login, "update_datetime", updateDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	LoginDbm.ColumnUpdateUser = df.CCI(&login, "update_user", updateUserSqlName, "", "",
               "String.class", "updateUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	LoginDbm.ColumnUpdateProcess = df.CCI(&login, "update_process", updateProcessSqlName, "", "",
               "String.class", "updateProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")

	LoginDbm.ColumnInfoList = new(df.List)
	LoginDbm.ColumnInfoList.Add(LoginDbm.ColumnId)
	LoginDbm.ColumnInfoList.Add(LoginDbm.ColumnLoginId)
	LoginDbm.ColumnInfoList.Add(LoginDbm.ColumnName)
	LoginDbm.ColumnInfoList.Add(LoginDbm.ColumnPassword)
	LoginDbm.ColumnInfoList.Add(LoginDbm.ColumnVersionNo)
	LoginDbm.ColumnInfoList.Add(LoginDbm.ColumnDelFlag)
	LoginDbm.ColumnInfoList.Add(LoginDbm.ColumnRegisterDatetime)
	LoginDbm.ColumnInfoList.Add(LoginDbm.ColumnRegisterUser)
	LoginDbm.ColumnInfoList.Add(LoginDbm.ColumnRegisterProcess)
	LoginDbm.ColumnInfoList.Add(LoginDbm.ColumnUpdateDatetime)
	LoginDbm.ColumnInfoList.Add(LoginDbm.ColumnUpdateUser)
	LoginDbm.ColumnInfoList.Add(LoginDbm.ColumnUpdateProcess)


	LoginDbm.ColumnInfoMap=make(map[string]int)
	LoginDbm.ColumnInfoMap["id"]=0
		LoginDbm.ColumnInfoMap["loginId"]=1
		LoginDbm.ColumnInfoMap["name"]=2
		LoginDbm.ColumnInfoMap["password"]=3
		LoginDbm.ColumnInfoMap["versionNo"]=4
		LoginDbm.ColumnInfoMap["delFlag"]=5
		LoginDbm.ColumnInfoMap["registerDatetime"]=6
		LoginDbm.ColumnInfoMap["registerUser"]=7
		LoginDbm.ColumnInfoMap["registerProcess"]=8
		LoginDbm.ColumnInfoMap["updateDatetime"]=9
		LoginDbm.ColumnInfoMap["updateUser"]=10
		LoginDbm.ColumnInfoMap["updateProcess"]=11
	    LoginDbm.PrimaryKey = true
    LoginDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &login
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(LoginDbm.ColumnId)

	LoginDbm.PrimaryInfo = new(df.PrimaryInfo)
	LoginDbm.PrimaryInfo.UniqueInfo = ui
	
	LoginDbm.SequenceFlag = true
	LoginDbm.SequenceName = "login_id_seq"
	LoginDbm.SequenceIncrementSize = 1
	LoginDbm.SequenceCacheSize = 0
	LoginDbm.VersionNoFlag = true
	LoginDbm.VersionNoColumnInfo = LoginDbm.ColumnVersionNo
	
	var loginMeta df.DBMeta = LoginDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Login"] = &loginMeta
}

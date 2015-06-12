package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type SessionDbm_T struct {
	df.BaseDBMeta
	ColumnId *df.ColumnInfo
	ColumnUuid *df.ColumnInfo
	ColumnLoginId *df.ColumnInfo
	ColumnCompId *df.ColumnInfo
	ColumnData *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
	ColumnDelFlag *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
}

func (b *SessionDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *SessionDbm_T) foreignLogin() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		SessionDbm.GetColumnInfoByPropertyName("loginId"),
		LoginDbm.GetColumnInfoByPropertyName("id"),
	}

	return b.BaseDBMeta.Cfi("session_login_fkey", "Login",
		columns, 0, false, false, false, false,
		"", nil, false, "sessionList")
}	
func (b *SessionDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
	b.ForeignInfoMap["Login"] = b.foreignLogin()
}

func (b *SessionDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var SessionDbm *SessionDbm_T

func Create_SessionDbm() {
	SessionDbm = new(SessionDbm_T)
	SessionDbm.TableDbName = "session"
	SessionDbm.TableDispName = "session"
	SessionDbm.TablePropertyName = "session"
	SessionDbm.TableSqlName = new(df.TableSqlName)
	SessionDbm.TableSqlName.TableSqlName = "session"
	SessionDbm.TableSqlName.CorrespondingDbName = SessionDbm.TableDbName

	var session df.DBMeta
	session = SessionDbm
	SessionDbm.DBMeta=&session
	idSqlName := new(df.ColumnSqlName)
	idSqlName.ColumnSqlName = "id"
	idSqlName.IrregularChar = false
	SessionDbm.ColumnId = df.CCI(&session, "id", idSqlName, "", "",
               "Integer.class", "id", "", true, true,true, "serial", 10, 0,
               "nextval('session_id_seq'::regclass)",false,"","", "","","",false,"int64")
	uuidSqlName := new(df.ColumnSqlName)
	uuidSqlName.ColumnSqlName = "uuid"
	uuidSqlName.IrregularChar = false
	SessionDbm.ColumnUuid = df.CCI(&session, "uuid", uuidSqlName, "", "",
               "String.class", "uuid", "", false, false,true, "varchar", 50, 0,
               "",false,"","", "","","",false,"string")
	loginIdSqlName := new(df.ColumnSqlName)
	loginIdSqlName.ColumnSqlName = "login_id"
	loginIdSqlName.IrregularChar = false
	SessionDbm.ColumnLoginId = df.CCI(&session, "login_id", loginIdSqlName, "", "",
               "Integer.class", "loginId", "", false, false,false, "int4", 10, 0,
               "",false,"","", "login","","",false,"sql.NullInt64")
	compIdSqlName := new(df.ColumnSqlName)
	compIdSqlName.ColumnSqlName = "comp_id"
	compIdSqlName.IrregularChar = false
	SessionDbm.ColumnCompId = df.CCI(&session, "comp_id", compIdSqlName, "", "",
               "Integer.class", "compId", "", false, false,false, "int4", 10, 0,
               "",false,"","", "","","",false,"sql.NullInt64")
	dataSqlName := new(df.ColumnSqlName)
	dataSqlName.ColumnSqlName = "data"
	dataSqlName.IrregularChar = false
	SessionDbm.ColumnData = df.CCI(&session, "data", dataSqlName, "", "",
               "String.class", "data", "", false, false,false, "varchar", 255, 0,
               "",false,"","", "","","",false,"sql.NullString")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	SessionDbm.ColumnVersionNo = df.CCI(&session, "version_no", versionNoSqlName, "", "",
               "Integer.class", "versionNo", "", false, false,true, "int4", 10, 0,
               "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")
	delFlagSqlName := new(df.ColumnSqlName)
	delFlagSqlName.ColumnSqlName = "del_flag"
	delFlagSqlName.IrregularChar = false
	SessionDbm.ColumnDelFlag = df.CCI(&session, "del_flag", delFlagSqlName, "", "",
               "Integer.class", "delFlag", "", false, false,true, "int4", 10, 0,
               "0",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	SessionDbm.ColumnRegisterDatetime = df.CCI(&session, "register_datetime", registerDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	SessionDbm.ColumnRegisterUser = df.CCI(&session, "register_user", registerUserSqlName, "", "",
               "String.class", "registerUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	SessionDbm.ColumnRegisterProcess = df.CCI(&session, "register_process", registerProcessSqlName, "", "",
               "String.class", "registerProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	SessionDbm.ColumnUpdateDatetime = df.CCI(&session, "update_datetime", updateDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	SessionDbm.ColumnUpdateUser = df.CCI(&session, "update_user", updateUserSqlName, "", "",
               "String.class", "updateUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	SessionDbm.ColumnUpdateProcess = df.CCI(&session, "update_process", updateProcessSqlName, "", "",
               "String.class", "updateProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")

	SessionDbm.ColumnInfoList = new(df.List)
	SessionDbm.ColumnInfoList.Add(SessionDbm.ColumnId)
	SessionDbm.ColumnInfoList.Add(SessionDbm.ColumnUuid)
	SessionDbm.ColumnInfoList.Add(SessionDbm.ColumnLoginId)
	SessionDbm.ColumnInfoList.Add(SessionDbm.ColumnCompId)
	SessionDbm.ColumnInfoList.Add(SessionDbm.ColumnData)
	SessionDbm.ColumnInfoList.Add(SessionDbm.ColumnVersionNo)
	SessionDbm.ColumnInfoList.Add(SessionDbm.ColumnDelFlag)
	SessionDbm.ColumnInfoList.Add(SessionDbm.ColumnRegisterDatetime)
	SessionDbm.ColumnInfoList.Add(SessionDbm.ColumnRegisterUser)
	SessionDbm.ColumnInfoList.Add(SessionDbm.ColumnRegisterProcess)
	SessionDbm.ColumnInfoList.Add(SessionDbm.ColumnUpdateDatetime)
	SessionDbm.ColumnInfoList.Add(SessionDbm.ColumnUpdateUser)
	SessionDbm.ColumnInfoList.Add(SessionDbm.ColumnUpdateProcess)


	SessionDbm.ColumnInfoMap=make(map[string]int)
	SessionDbm.ColumnInfoMap["id"]=0
		SessionDbm.ColumnInfoMap["uuid"]=1
		SessionDbm.ColumnInfoMap["loginId"]=2
		SessionDbm.ColumnInfoMap["compId"]=3
		SessionDbm.ColumnInfoMap["data"]=4
		SessionDbm.ColumnInfoMap["versionNo"]=5
		SessionDbm.ColumnInfoMap["delFlag"]=6
		SessionDbm.ColumnInfoMap["registerDatetime"]=7
		SessionDbm.ColumnInfoMap["registerUser"]=8
		SessionDbm.ColumnInfoMap["registerProcess"]=9
		SessionDbm.ColumnInfoMap["updateDatetime"]=10
		SessionDbm.ColumnInfoMap["updateUser"]=11
		SessionDbm.ColumnInfoMap["updateProcess"]=12
	    SessionDbm.PrimaryKey = true
    SessionDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &session
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(SessionDbm.ColumnId)

	SessionDbm.PrimaryInfo = new(df.PrimaryInfo)
	SessionDbm.PrimaryInfo.UniqueInfo = ui
	
	SessionDbm.SequenceFlag = true
	SessionDbm.SequenceName = "session_id_seq"
	SessionDbm.SequenceIncrementSize = 1
	SessionDbm.SequenceCacheSize = 0
	SessionDbm.VersionNoFlag = true
	SessionDbm.VersionNoColumnInfo = SessionDbm.ColumnVersionNo
	
	var sessionMeta df.DBMeta = SessionDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Session"] = &sessionMeta
}

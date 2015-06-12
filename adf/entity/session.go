package entity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type Session struct {
	id int64
	uuid string
	loginId sql.NullInt64
	compId sql.NullInt64
	data sql.NullString
	versionNo int64
	delFlag int64
	registerDatetime df.Timestamp
	registerUser string
	registerProcess string
	updateDatetime df.Timestamp
	updateUser string
	updateProcess string
	df.BaseEntity
Login_R  *Login

}

func CreateSession() *Session{
	session :=new(Session)
	session.SetUp()
	return session 
}

func (l *Session) GetId () int64 {
	return l.id
}
func (l *Session) GetUuid () string {
	return l.uuid
}
func (l *Session) GetLoginId () sql.NullInt64 {
	return l.loginId
}
func (l *Session) GetCompId () sql.NullInt64 {
	return l.compId
}
func (l *Session) GetData () sql.NullString {
	return l.data
}
func (l *Session) GetVersionNo () int64 {
	return l.versionNo
}
func (l *Session) GetDelFlag () int64 {
	return l.delFlag
}
func (l *Session) GetRegisterDatetime () df.Timestamp {
	return l.registerDatetime
}
func (l *Session) GetRegisterUser () string {
	return l.registerUser
}
func (l *Session) GetRegisterProcess () string {
	return l.registerProcess
}
func (l *Session) GetUpdateDatetime () df.Timestamp {
	return l.updateDatetime
}
func (l *Session) GetUpdateUser () string {
	return l.updateUser
}
func (l *Session) GetUpdateProcess () string {
	return l.updateProcess
}

func (t *Session) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 13)
	i[0] = &(t.id)
	i[1] = &(t.uuid)
	i[2] = &(t.loginId)
	i[3] = &(t.compId)
	i[4] = &(t.data)
	i[5] = &(t.versionNo)
	i[6] = &(t.delFlag)
	i[7] = &(t.registerDatetime)
	i[8] = &(t.registerUser)
	i[9] = &(t.registerProcess)
	i[10] = &(t.updateDatetime)
	i[11] = &(t.updateUser)
	i[12] = &(t.updateProcess)
	return i
}


func (t *Session) AsTableDbName() string {
	return "Session"
}

func (t *Session) HasPrimaryKeyValue() bool{
        if t.id == 0 {
            return false 
        }
        return true
}
func (t *Session) SetId(id int64) {
	t.AddPropertyName("id")
	t.id = id
}
func (t *Session) SetUuid(uuid string) {
	t.AddPropertyName("uuid")
	t.uuid = uuid
}
func (t *Session) SetLoginId(loginId sql.NullInt64) {
	t.AddPropertyName("loginId")
	t.loginId = loginId
}
func (t *Session) SetCompId(compId sql.NullInt64) {
	t.AddPropertyName("compId")
	t.compId = compId
}
func (t *Session) SetData(data sql.NullString) {
	t.AddPropertyName("data")
	t.data = data
}
func (t *Session) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}
func (t *Session) SetDelFlag(delFlag int64) {
	t.AddPropertyName("delFlag")
	t.delFlag = delFlag
}
func (t *Session) SetRegisterDatetime(registerDatetime df.Timestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *Session) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *Session) SetRegisterProcess(registerProcess string) {
	t.AddPropertyName("registerProcess")
	t.registerProcess = registerProcess
}
func (t *Session) SetUpdateDatetime(updateDatetime df.Timestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *Session) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *Session) SetUpdateProcess(updateProcess string) {
	t.AddPropertyName("updateProcess")
	t.updateProcess = updateProcess
}
func (t *Session) GetLogin_R() *Login{
	return t.Login_R
}
func (t *Session) SetLogin_R(value *Login) {
    t.Login_R = value
}
func (t *Session) SetUp(){
	
}
func (t *Session)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type Login struct {
	id int64
	loginId string
	name string
	password string
	versionNo int64
	delFlag int64
	registerDatetime df.Timestamp
	registerUser string
	registerProcess string
	updateDatetime df.Timestamp
	updateUser string
	updateProcess string
	df.BaseEntity
}

func CreateLogin() *Login{
	login :=new(Login)
	login.SetUp()
	return login 
}

func (l *Login) GetId () int64 {
	return l.id
}
func (l *Login) GetLoginId () string {
	return l.loginId
}
func (l *Login) GetName () string {
	return l.name
}
func (l *Login) GetPassword () string {
	return l.password
}
func (l *Login) GetVersionNo () int64 {
	return l.versionNo
}
func (l *Login) GetDelFlag () int64 {
	return l.delFlag
}
func (l *Login) GetRegisterDatetime () df.Timestamp {
	return l.registerDatetime
}
func (l *Login) GetRegisterUser () string {
	return l.registerUser
}
func (l *Login) GetRegisterProcess () string {
	return l.registerProcess
}
func (l *Login) GetUpdateDatetime () df.Timestamp {
	return l.updateDatetime
}
func (l *Login) GetUpdateUser () string {
	return l.updateUser
}
func (l *Login) GetUpdateProcess () string {
	return l.updateProcess
}

func (t *Login) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 12)
	i[0] = &(t.id)
	i[1] = &(t.loginId)
	i[2] = &(t.name)
	i[3] = &(t.password)
	i[4] = &(t.versionNo)
	i[5] = &(t.delFlag)
	i[6] = &(t.registerDatetime)
	i[7] = &(t.registerUser)
	i[8] = &(t.registerProcess)
	i[9] = &(t.updateDatetime)
	i[10] = &(t.updateUser)
	i[11] = &(t.updateProcess)
	return i
}


func (t *Login) AsTableDbName() string {
	return "Login"
}

func (t *Login) HasPrimaryKeyValue() bool{
        if t.id == 0 {
            return false 
        }
        return true
}
func (t *Login) SetId(id int64) {
	t.AddPropertyName("id")
	t.id = id
}
func (t *Login) SetLoginId(loginId string) {
	t.AddPropertyName("loginId")
	t.loginId = loginId
}
func (t *Login) SetName(name string) {
	t.AddPropertyName("name")
	t.name = name
}
func (t *Login) SetPassword(password string) {
	t.AddPropertyName("password")
	t.password = password
}
func (t *Login) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}
func (t *Login) SetDelFlag(delFlag int64) {
	t.AddPropertyName("delFlag")
	t.delFlag = delFlag
}
func (t *Login) SetRegisterDatetime(registerDatetime df.Timestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *Login) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *Login) SetRegisterProcess(registerProcess string) {
	t.AddPropertyName("registerProcess")
	t.registerProcess = registerProcess
}
func (t *Login) SetUpdateDatetime(updateDatetime df.Timestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *Login) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *Login) SetUpdateProcess(updateProcess string) {
	t.AddPropertyName("updateProcess")
	t.updateProcess = updateProcess
}
func (t *Login) SetUp(){
	
}
func (t *Login)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
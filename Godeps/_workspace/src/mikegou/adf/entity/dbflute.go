package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type Dbflute struct {
	id int64
	loginId int64
	db string
	dbName string
	dbTable string
	name string
	dbType string
	required bool
	javaType string
	goType string
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

func CreateDbflute() *Dbflute{
	dbflute :=new(Dbflute)
	dbflute.SetUp()
	return dbflute 
}

func (l *Dbflute) GetId () int64 {
	return l.id
}
func (l *Dbflute) GetLoginId () int64 {
	return l.loginId
}
func (l *Dbflute) GetDb () string {
	return l.db
}
func (l *Dbflute) GetDbName () string {
	return l.dbName
}
func (l *Dbflute) GetDbTable () string {
	return l.dbTable
}
func (l *Dbflute) GetName () string {
	return l.name
}
func (l *Dbflute) GetDbType () string {
	return l.dbType
}
func (l *Dbflute) GetRequired () bool {
	return l.required
}
func (l *Dbflute) GetJavaType () string {
	return l.javaType
}
func (l *Dbflute) GetGoType () string {
	return l.goType
}
func (l *Dbflute) GetVersionNo () int64 {
	return l.versionNo
}
func (l *Dbflute) GetDelFlag () int64 {
	return l.delFlag
}
func (l *Dbflute) GetRegisterDatetime () df.Timestamp {
	return l.registerDatetime
}
func (l *Dbflute) GetRegisterUser () string {
	return l.registerUser
}
func (l *Dbflute) GetRegisterProcess () string {
	return l.registerProcess
}
func (l *Dbflute) GetUpdateDatetime () df.Timestamp {
	return l.updateDatetime
}
func (l *Dbflute) GetUpdateUser () string {
	return l.updateUser
}
func (l *Dbflute) GetUpdateProcess () string {
	return l.updateProcess
}

func (t *Dbflute) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 18)
	i[0] = &(t.id)
	i[1] = &(t.loginId)
	i[2] = &(t.db)
	i[3] = &(t.dbName)
	i[4] = &(t.dbTable)
	i[5] = &(t.name)
	i[6] = &(t.dbType)
	i[7] = &(t.required)
	i[8] = &(t.javaType)
	i[9] = &(t.goType)
	i[10] = &(t.versionNo)
	i[11] = &(t.delFlag)
	i[12] = &(t.registerDatetime)
	i[13] = &(t.registerUser)
	i[14] = &(t.registerProcess)
	i[15] = &(t.updateDatetime)
	i[16] = &(t.updateUser)
	i[17] = &(t.updateProcess)
	return i
}


func (t *Dbflute) AsTableDbName() string {
	return "Dbflute"
}

func (t *Dbflute) HasPrimaryKeyValue() bool{
        if t.id == 0 {
            return false 
        }
        return true
}
func (t *Dbflute) SetId(id int64) {
	t.AddPropertyName("id")
	t.id = id
}
func (t *Dbflute) SetLoginId(loginId int64) {
	t.AddPropertyName("loginId")
	t.loginId = loginId
}
func (t *Dbflute) SetDb(db string) {
	t.AddPropertyName("db")
	t.db = db
}
func (t *Dbflute) SetDbName(dbName string) {
	t.AddPropertyName("dbName")
	t.dbName = dbName
}
func (t *Dbflute) SetDbTable(dbTable string) {
	t.AddPropertyName("dbTable")
	t.dbTable = dbTable
}
func (t *Dbflute) SetName(name string) {
	t.AddPropertyName("name")
	t.name = name
}
func (t *Dbflute) SetDbType(dbType string) {
	t.AddPropertyName("dbType")
	t.dbType = dbType
}
func (t *Dbflute) SetRequired(required bool) {
	t.AddPropertyName("required")
	t.required = required
}
func (t *Dbflute) SetJavaType(javaType string) {
	t.AddPropertyName("javaType")
	t.javaType = javaType
}
func (t *Dbflute) SetGoType(goType string) {
	t.AddPropertyName("goType")
	t.goType = goType
}
func (t *Dbflute) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}
func (t *Dbflute) SetDelFlag(delFlag int64) {
	t.AddPropertyName("delFlag")
	t.delFlag = delFlag
}
func (t *Dbflute) SetRegisterDatetime(registerDatetime df.Timestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *Dbflute) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *Dbflute) SetRegisterProcess(registerProcess string) {
	t.AddPropertyName("registerProcess")
	t.registerProcess = registerProcess
}
func (t *Dbflute) SetUpdateDatetime(updateDatetime df.Timestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *Dbflute) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *Dbflute) SetUpdateProcess(updateProcess string) {
	t.AddPropertyName("updateProcess")
	t.updateProcess = updateProcess
}
func (t *Dbflute) GetLogin_R() *Login{
	return t.Login_R
}
func (t *Dbflute) SetLogin_R(value *Login) {
    t.Login_R = value
}
func (t *Dbflute) SetUp(){
	
}
func (t *Dbflute)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}

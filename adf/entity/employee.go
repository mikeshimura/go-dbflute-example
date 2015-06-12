package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type Employee struct {
	id int64
	empCd string
	secId int64
	name string
	versionNo int64
	delFlag int64
	registerDatetime df.Timestamp
	registerUser string
	registerProcess string
	updateDatetime df.Timestamp
	updateUser string
	updateProcess string
	df.BaseEntity
UserTable_R  *UserTable

}

func CreateEmployee() *Employee{
	employee :=new(Employee)
	employee.SetUp()
	return employee 
}

func (l *Employee) GetId () int64 {
	return l.id
}
func (l *Employee) GetEmpCd () string {
	return l.empCd
}
func (l *Employee) GetSecId () int64 {
	return l.secId
}
func (l *Employee) GetName () string {
	return l.name
}
func (l *Employee) GetVersionNo () int64 {
	return l.versionNo
}
func (l *Employee) GetDelFlag () int64 {
	return l.delFlag
}
func (l *Employee) GetRegisterDatetime () df.Timestamp {
	return l.registerDatetime
}
func (l *Employee) GetRegisterUser () string {
	return l.registerUser
}
func (l *Employee) GetRegisterProcess () string {
	return l.registerProcess
}
func (l *Employee) GetUpdateDatetime () df.Timestamp {
	return l.updateDatetime
}
func (l *Employee) GetUpdateUser () string {
	return l.updateUser
}
func (l *Employee) GetUpdateProcess () string {
	return l.updateProcess
}

func (t *Employee) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 12)
	i[0] = &(t.id)
	i[1] = &(t.empCd)
	i[2] = &(t.secId)
	i[3] = &(t.name)
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


func (t *Employee) AsTableDbName() string {
	return "Employee"
}

func (t *Employee) HasPrimaryKeyValue() bool{
        if t.id == 0 {
            return false 
        }
        return true
}
func (t *Employee) SetId(id int64) {
	t.AddPropertyName("id")
	t.id = id
}
func (t *Employee) SetEmpCd(empCd string) {
	t.AddPropertyName("empCd")
	t.empCd = empCd
}
func (t *Employee) SetSecId(secId int64) {
	t.AddPropertyName("secId")
	t.secId = secId
}
func (t *Employee) SetName(name string) {
	t.AddPropertyName("name")
	t.name = name
}
func (t *Employee) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}
func (t *Employee) SetDelFlag(delFlag int64) {
	t.AddPropertyName("delFlag")
	t.delFlag = delFlag
}
func (t *Employee) SetRegisterDatetime(registerDatetime df.Timestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *Employee) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *Employee) SetRegisterProcess(registerProcess string) {
	t.AddPropertyName("registerProcess")
	t.registerProcess = registerProcess
}
func (t *Employee) SetUpdateDatetime(updateDatetime df.Timestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *Employee) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *Employee) SetUpdateProcess(updateProcess string) {
	t.AddPropertyName("updateProcess")
	t.updateProcess = updateProcess
}
func (t *Employee) GetUserTable_R() *UserTable{
	return t.UserTable_R
}
func (t *Employee) SetUserTable_R(value *UserTable) {
    t.UserTable_R = value
}
func (t *Employee) SetUp(){
	
}
func (t *Employee)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
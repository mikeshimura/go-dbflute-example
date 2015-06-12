package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type Customer struct {
	id int64
	cusCd string
	name string
	addr string
	bldg string
	cusConSec string
	cusConName string
	tel string
	salesAmount int64
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

func CreateCustomer() *Customer{
	customer :=new(Customer)
	customer.SetUp()
	return customer 
}

func (l *Customer) GetId () int64 {
	return l.id
}
func (l *Customer) GetCusCd () string {
	return l.cusCd
}
func (l *Customer) GetName () string {
	return l.name
}
func (l *Customer) GetAddr () string {
	return l.addr
}
func (l *Customer) GetBldg () string {
	return l.bldg
}
func (l *Customer) GetCusConSec () string {
	return l.cusConSec
}
func (l *Customer) GetCusConName () string {
	return l.cusConName
}
func (l *Customer) GetTel () string {
	return l.tel
}
func (l *Customer) GetSalesAmount () int64 {
	return l.salesAmount
}
func (l *Customer) GetVersionNo () int64 {
	return l.versionNo
}
func (l *Customer) GetDelFlag () int64 {
	return l.delFlag
}
func (l *Customer) GetRegisterDatetime () df.Timestamp {
	return l.registerDatetime
}
func (l *Customer) GetRegisterUser () string {
	return l.registerUser
}
func (l *Customer) GetRegisterProcess () string {
	return l.registerProcess
}
func (l *Customer) GetUpdateDatetime () df.Timestamp {
	return l.updateDatetime
}
func (l *Customer) GetUpdateUser () string {
	return l.updateUser
}
func (l *Customer) GetUpdateProcess () string {
	return l.updateProcess
}

func (t *Customer) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 17)
	i[0] = &(t.id)
	i[1] = &(t.cusCd)
	i[2] = &(t.name)
	i[3] = &(t.addr)
	i[4] = &(t.bldg)
	i[5] = &(t.cusConSec)
	i[6] = &(t.cusConName)
	i[7] = &(t.tel)
	i[8] = &(t.salesAmount)
	i[9] = &(t.versionNo)
	i[10] = &(t.delFlag)
	i[11] = &(t.registerDatetime)
	i[12] = &(t.registerUser)
	i[13] = &(t.registerProcess)
	i[14] = &(t.updateDatetime)
	i[15] = &(t.updateUser)
	i[16] = &(t.updateProcess)
	return i
}


func (t *Customer) AsTableDbName() string {
	return "Customer"
}

func (t *Customer) HasPrimaryKeyValue() bool{
        if t.id == 0 {
            return false 
        }
        return true
}
func (t *Customer) SetId(id int64) {
	t.AddPropertyName("id")
	t.id = id
}
func (t *Customer) SetCusCd(cusCd string) {
	t.AddPropertyName("cusCd")
	t.cusCd = cusCd
}
func (t *Customer) SetName(name string) {
	t.AddPropertyName("name")
	t.name = name
}
func (t *Customer) SetAddr(addr string) {
	t.AddPropertyName("addr")
	t.addr = addr
}
func (t *Customer) SetBldg(bldg string) {
	t.AddPropertyName("bldg")
	t.bldg = bldg
}
func (t *Customer) SetCusConSec(cusConSec string) {
	t.AddPropertyName("cusConSec")
	t.cusConSec = cusConSec
}
func (t *Customer) SetCusConName(cusConName string) {
	t.AddPropertyName("cusConName")
	t.cusConName = cusConName
}
func (t *Customer) SetTel(tel string) {
	t.AddPropertyName("tel")
	t.tel = tel
}
func (t *Customer) SetSalesAmount(salesAmount int64) {
	t.AddPropertyName("salesAmount")
	t.salesAmount = salesAmount
}
func (t *Customer) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}
func (t *Customer) SetDelFlag(delFlag int64) {
	t.AddPropertyName("delFlag")
	t.delFlag = delFlag
}
func (t *Customer) SetRegisterDatetime(registerDatetime df.Timestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *Customer) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *Customer) SetRegisterProcess(registerProcess string) {
	t.AddPropertyName("registerProcess")
	t.registerProcess = registerProcess
}
func (t *Customer) SetUpdateDatetime(updateDatetime df.Timestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *Customer) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *Customer) SetUpdateProcess(updateProcess string) {
	t.AddPropertyName("updateProcess")
	t.updateProcess = updateProcess
}
func (t *Customer) SetUp(){
	
}
func (t *Customer)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
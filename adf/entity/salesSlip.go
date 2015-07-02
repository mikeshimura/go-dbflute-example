package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type SalesSlip struct {
	id int64
	slipNo int64
	line int64
	salesDate df.Date
	salesId int64
	cusId int64
	prdId int64
	qty int64
	unit string
	unitPrice int64
	salesAmount int64
	slipAmount int64
	slipCons int64
	versionNo int64
	delFlag int64
	registerDatetime df.Timestamp
	registerUser string
	registerProcess string
	updateDatetime df.Timestamp
	updateUser string
	updateProcess string
	df.BaseEntity
Product_R  *Product

Employee_R  *Employee

}

func CreateSalesSlip() *SalesSlip{
	salesSlip :=new(SalesSlip)
	salesSlip.SetUp()
	return salesSlip 
}

func (l *SalesSlip) GetId () int64 {
	return l.id
}
func (l *SalesSlip) GetSlipNo () int64 {
	return l.slipNo
}
func (l *SalesSlip) GetLine () int64 {
	return l.line
}
func (l *SalesSlip) GetSalesDate () df.Date {
	return l.salesDate
}
func (l *SalesSlip) GetSalesId () int64 {
	return l.salesId
}
func (l *SalesSlip) GetCusId () int64 {
	return l.cusId
}
func (l *SalesSlip) GetPrdId () int64 {
	return l.prdId
}
func (l *SalesSlip) GetQty () int64 {
	return l.qty
}
func (l *SalesSlip) GetUnit () string {
	return l.unit
}
func (l *SalesSlip) GetUnitPrice () int64 {
	return l.unitPrice
}
func (l *SalesSlip) GetSalesAmount () int64 {
	return l.salesAmount
}
func (l *SalesSlip) GetSlipAmount () int64 {
	return l.slipAmount
}
func (l *SalesSlip) GetSlipCons () int64 {
	return l.slipCons
}
func (l *SalesSlip) GetVersionNo () int64 {
	return l.versionNo
}
func (l *SalesSlip) GetDelFlag () int64 {
	return l.delFlag
}
func (l *SalesSlip) GetRegisterDatetime () df.Timestamp {
	return l.registerDatetime
}
func (l *SalesSlip) GetRegisterUser () string {
	return l.registerUser
}
func (l *SalesSlip) GetRegisterProcess () string {
	return l.registerProcess
}
func (l *SalesSlip) GetUpdateDatetime () df.Timestamp {
	return l.updateDatetime
}
func (l *SalesSlip) GetUpdateUser () string {
	return l.updateUser
}
func (l *SalesSlip) GetUpdateProcess () string {
	return l.updateProcess
}

func (t *SalesSlip) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 21)
	i[0] = &(t.id)
	i[1] = &(t.slipNo)
	i[2] = &(t.line)
	i[3] = &(t.salesDate)
	i[4] = &(t.salesId)
	i[5] = &(t.cusId)
	i[6] = &(t.prdId)
	i[7] = &(t.qty)
	i[8] = &(t.unit)
	i[9] = &(t.unitPrice)
	i[10] = &(t.salesAmount)
	i[11] = &(t.slipAmount)
	i[12] = &(t.slipCons)
	i[13] = &(t.versionNo)
	i[14] = &(t.delFlag)
	i[15] = &(t.registerDatetime)
	i[16] = &(t.registerUser)
	i[17] = &(t.registerProcess)
	i[18] = &(t.updateDatetime)
	i[19] = &(t.updateUser)
	i[20] = &(t.updateProcess)
	return i
}


func (t *SalesSlip) AsTableDbName() string {
	return "SalesSlip"
}

func (t *SalesSlip) HasPrimaryKeyValue() bool{
        if t.id == 0 {
            return false 
        }
        return true
}
func (t *SalesSlip) SetId(id int64) {
	t.AddPropertyName("id")
	t.id = id
}
func (t *SalesSlip) SetSlipNo(slipNo int64) {
	t.AddPropertyName("slipNo")
	t.slipNo = slipNo
}
func (t *SalesSlip) SetLine(line int64) {
	t.AddPropertyName("line")
	t.line = line
}
func (t *SalesSlip) SetSalesDate(salesDate df.Date) {
	t.AddPropertyName("salesDate")
	t.salesDate = salesDate
}
func (t *SalesSlip) SetSalesId(salesId int64) {
	t.AddPropertyName("salesId")
	t.salesId = salesId
}
func (t *SalesSlip) SetCusId(cusId int64) {
	t.AddPropertyName("cusId")
	t.cusId = cusId
}
func (t *SalesSlip) SetPrdId(prdId int64) {
	t.AddPropertyName("prdId")
	t.prdId = prdId
}
func (t *SalesSlip) SetQty(qty int64) {
	t.AddPropertyName("qty")
	t.qty = qty
}
func (t *SalesSlip) SetUnit(unit string) {
	t.AddPropertyName("unit")
	t.unit = unit
}
func (t *SalesSlip) SetUnitPrice(unitPrice int64) {
	t.AddPropertyName("unitPrice")
	t.unitPrice = unitPrice
}
func (t *SalesSlip) SetSalesAmount(salesAmount int64) {
	t.AddPropertyName("salesAmount")
	t.salesAmount = salesAmount
}
func (t *SalesSlip) SetSlipAmount(slipAmount int64) {
	t.AddPropertyName("slipAmount")
	t.slipAmount = slipAmount
}
func (t *SalesSlip) SetSlipCons(slipCons int64) {
	t.AddPropertyName("slipCons")
	t.slipCons = slipCons
}
func (t *SalesSlip) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}
func (t *SalesSlip) SetDelFlag(delFlag int64) {
	t.AddPropertyName("delFlag")
	t.delFlag = delFlag
}
func (t *SalesSlip) SetRegisterDatetime(registerDatetime df.Timestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *SalesSlip) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *SalesSlip) SetRegisterProcess(registerProcess string) {
	t.AddPropertyName("registerProcess")
	t.registerProcess = registerProcess
}
func (t *SalesSlip) SetUpdateDatetime(updateDatetime df.Timestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *SalesSlip) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *SalesSlip) SetUpdateProcess(updateProcess string) {
	t.AddPropertyName("updateProcess")
	t.updateProcess = updateProcess
}
func (t *SalesSlip) GetProduct_R() *Product{
	return t.Product_R
}
func (t *SalesSlip) SetProduct_R(value *Product) {
    t.Product_R = value
}
func (t *SalesSlip) GetEmployee_R() *Employee{
	return t.Employee_R
}
func (t *SalesSlip) SetEmployee_R(value *Employee) {
    t.Employee_R = value
}
func (t *SalesSlip) SetUp(){
	
}
func (t *SalesSlip)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
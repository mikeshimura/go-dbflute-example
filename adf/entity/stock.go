package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type Stock struct {
	id int64
	yyyymm string
	prdId int64
	name string
	cost int64
	price int64
	unit string
	beforeQty int64
	purQty int64
	rcvQty int64
	salesQty int64
	outQty int64
	stockQty int64
	stockAmount int64
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

}

func CreateStock() *Stock{
	stock :=new(Stock)
	stock.SetUp()
	return stock 
}

func (l *Stock) GetId () int64 {
	return l.id
}
func (l *Stock) GetYyyymm () string {
	return l.yyyymm
}
func (l *Stock) GetPrdId () int64 {
	return l.prdId
}
func (l *Stock) GetName () string {
	return l.name
}
func (l *Stock) GetCost () int64 {
	return l.cost
}
func (l *Stock) GetPrice () int64 {
	return l.price
}
func (l *Stock) GetUnit () string {
	return l.unit
}
func (l *Stock) GetBeforeQty () int64 {
	return l.beforeQty
}
func (l *Stock) GetPurQty () int64 {
	return l.purQty
}
func (l *Stock) GetRcvQty () int64 {
	return l.rcvQty
}
func (l *Stock) GetSalesQty () int64 {
	return l.salesQty
}
func (l *Stock) GetOutQty () int64 {
	return l.outQty
}
func (l *Stock) GetStockQty () int64 {
	return l.stockQty
}
func (l *Stock) GetStockAmount () int64 {
	return l.stockAmount
}
func (l *Stock) GetVersionNo () int64 {
	return l.versionNo
}
func (l *Stock) GetDelFlag () int64 {
	return l.delFlag
}
func (l *Stock) GetRegisterDatetime () df.Timestamp {
	return l.registerDatetime
}
func (l *Stock) GetRegisterUser () string {
	return l.registerUser
}
func (l *Stock) GetRegisterProcess () string {
	return l.registerProcess
}
func (l *Stock) GetUpdateDatetime () df.Timestamp {
	return l.updateDatetime
}
func (l *Stock) GetUpdateUser () string {
	return l.updateUser
}
func (l *Stock) GetUpdateProcess () string {
	return l.updateProcess
}

func (t *Stock) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 22)
	i[0] = &(t.id)
	i[1] = &(t.yyyymm)
	i[2] = &(t.prdId)
	i[3] = &(t.name)
	i[4] = &(t.cost)
	i[5] = &(t.price)
	i[6] = &(t.unit)
	i[7] = &(t.beforeQty)
	i[8] = &(t.purQty)
	i[9] = &(t.rcvQty)
	i[10] = &(t.salesQty)
	i[11] = &(t.outQty)
	i[12] = &(t.stockQty)
	i[13] = &(t.stockAmount)
	i[14] = &(t.versionNo)
	i[15] = &(t.delFlag)
	i[16] = &(t.registerDatetime)
	i[17] = &(t.registerUser)
	i[18] = &(t.registerProcess)
	i[19] = &(t.updateDatetime)
	i[20] = &(t.updateUser)
	i[21] = &(t.updateProcess)
	return i
}


func (t *Stock) AsTableDbName() string {
	return "Stock"
}

func (t *Stock) HasPrimaryKeyValue() bool{
        if t.id == 0 {
            return false 
        }
        return true
}
func (t *Stock) SetId(id int64) {
	t.AddPropertyName("id")
	t.id = id
}
func (t *Stock) SetYyyymm(yyyymm string) {
	t.AddPropertyName("yyyymm")
	t.yyyymm = yyyymm
}
func (t *Stock) SetPrdId(prdId int64) {
	t.AddPropertyName("prdId")
	t.prdId = prdId
}
func (t *Stock) SetName(name string) {
	t.AddPropertyName("name")
	t.name = name
}
func (t *Stock) SetCost(cost int64) {
	t.AddPropertyName("cost")
	t.cost = cost
}
func (t *Stock) SetPrice(price int64) {
	t.AddPropertyName("price")
	t.price = price
}
func (t *Stock) SetUnit(unit string) {
	t.AddPropertyName("unit")
	t.unit = unit
}
func (t *Stock) SetBeforeQty(beforeQty int64) {
	t.AddPropertyName("beforeQty")
	t.beforeQty = beforeQty
}
func (t *Stock) SetPurQty(purQty int64) {
	t.AddPropertyName("purQty")
	t.purQty = purQty
}
func (t *Stock) SetRcvQty(rcvQty int64) {
	t.AddPropertyName("rcvQty")
	t.rcvQty = rcvQty
}
func (t *Stock) SetSalesQty(salesQty int64) {
	t.AddPropertyName("salesQty")
	t.salesQty = salesQty
}
func (t *Stock) SetOutQty(outQty int64) {
	t.AddPropertyName("outQty")
	t.outQty = outQty
}
func (t *Stock) SetStockQty(stockQty int64) {
	t.AddPropertyName("stockQty")
	t.stockQty = stockQty
}
func (t *Stock) SetStockAmount(stockAmount int64) {
	t.AddPropertyName("stockAmount")
	t.stockAmount = stockAmount
}
func (t *Stock) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}
func (t *Stock) SetDelFlag(delFlag int64) {
	t.AddPropertyName("delFlag")
	t.delFlag = delFlag
}
func (t *Stock) SetRegisterDatetime(registerDatetime df.Timestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *Stock) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *Stock) SetRegisterProcess(registerProcess string) {
	t.AddPropertyName("registerProcess")
	t.registerProcess = registerProcess
}
func (t *Stock) SetUpdateDatetime(updateDatetime df.Timestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *Stock) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *Stock) SetUpdateProcess(updateProcess string) {
	t.AddPropertyName("updateProcess")
	t.updateProcess = updateProcess
}
func (t *Stock) GetProduct_R() *Product{
	return t.Product_R
}
func (t *Stock) SetProduct_R(value *Product) {
    t.Product_R = value
}
func (t *Stock) SetUp(){
	
}
func (t *Stock)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
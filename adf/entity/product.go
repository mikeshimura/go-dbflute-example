package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type Product struct {
	id int64
	prdCd string
	prdCat string
	name string
	cost int64
	price int64
	perQty int64
	unit string
	purQty int64
	purAmount int64
	salesQty int64
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

func CreateProduct() *Product{
	product :=new(Product)
	product.SetUp()
	return product 
}

func (l *Product) GetId () int64 {
	return l.id
}
func (l *Product) GetPrdCd () string {
	return l.prdCd
}
func (l *Product) GetPrdCat () string {
	return l.prdCat
}
func (l *Product) GetName () string {
	return l.name
}
func (l *Product) GetCost () int64 {
	return l.cost
}
func (l *Product) GetPrice () int64 {
	return l.price
}
func (l *Product) GetPerQty () int64 {
	return l.perQty
}
func (l *Product) GetUnit () string {
	return l.unit
}
func (l *Product) GetPurQty () int64 {
	return l.purQty
}
func (l *Product) GetPurAmount () int64 {
	return l.purAmount
}
func (l *Product) GetSalesQty () int64 {
	return l.salesQty
}
func (l *Product) GetSalesAmount () int64 {
	return l.salesAmount
}
func (l *Product) GetVersionNo () int64 {
	return l.versionNo
}
func (l *Product) GetDelFlag () int64 {
	return l.delFlag
}
func (l *Product) GetRegisterDatetime () df.Timestamp {
	return l.registerDatetime
}
func (l *Product) GetRegisterUser () string {
	return l.registerUser
}
func (l *Product) GetRegisterProcess () string {
	return l.registerProcess
}
func (l *Product) GetUpdateDatetime () df.Timestamp {
	return l.updateDatetime
}
func (l *Product) GetUpdateUser () string {
	return l.updateUser
}
func (l *Product) GetUpdateProcess () string {
	return l.updateProcess
}

func (t *Product) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 20)
	i[0] = &(t.id)
	i[1] = &(t.prdCd)
	i[2] = &(t.prdCat)
	i[3] = &(t.name)
	i[4] = &(t.cost)
	i[5] = &(t.price)
	i[6] = &(t.perQty)
	i[7] = &(t.unit)
	i[8] = &(t.purQty)
	i[9] = &(t.purAmount)
	i[10] = &(t.salesQty)
	i[11] = &(t.salesAmount)
	i[12] = &(t.versionNo)
	i[13] = &(t.delFlag)
	i[14] = &(t.registerDatetime)
	i[15] = &(t.registerUser)
	i[16] = &(t.registerProcess)
	i[17] = &(t.updateDatetime)
	i[18] = &(t.updateUser)
	i[19] = &(t.updateProcess)
	return i
}


func (t *Product) AsTableDbName() string {
	return "Product"
}

func (t *Product) HasPrimaryKeyValue() bool{
        if t.id == 0 {
            return false 
        }
        return true
}
func (t *Product) SetId(id int64) {
	t.AddPropertyName("id")
	t.id = id
}
func (t *Product) SetPrdCd(prdCd string) {
	t.AddPropertyName("prdCd")
	t.prdCd = prdCd
}
func (t *Product) SetPrdCat(prdCat string) {
	t.AddPropertyName("prdCat")
	t.prdCat = prdCat
}
func (t *Product) SetName(name string) {
	t.AddPropertyName("name")
	t.name = name
}
func (t *Product) SetCost(cost int64) {
	t.AddPropertyName("cost")
	t.cost = cost
}
func (t *Product) SetPrice(price int64) {
	t.AddPropertyName("price")
	t.price = price
}
func (t *Product) SetPerQty(perQty int64) {
	t.AddPropertyName("perQty")
	t.perQty = perQty
}
func (t *Product) SetUnit(unit string) {
	t.AddPropertyName("unit")
	t.unit = unit
}
func (t *Product) SetPurQty(purQty int64) {
	t.AddPropertyName("purQty")
	t.purQty = purQty
}
func (t *Product) SetPurAmount(purAmount int64) {
	t.AddPropertyName("purAmount")
	t.purAmount = purAmount
}
func (t *Product) SetSalesQty(salesQty int64) {
	t.AddPropertyName("salesQty")
	t.salesQty = salesQty
}
func (t *Product) SetSalesAmount(salesAmount int64) {
	t.AddPropertyName("salesAmount")
	t.salesAmount = salesAmount
}
func (t *Product) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}
func (t *Product) SetDelFlag(delFlag int64) {
	t.AddPropertyName("delFlag")
	t.delFlag = delFlag
}
func (t *Product) SetRegisterDatetime(registerDatetime df.Timestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *Product) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *Product) SetRegisterProcess(registerProcess string) {
	t.AddPropertyName("registerProcess")
	t.registerProcess = registerProcess
}
func (t *Product) SetUpdateDatetime(updateDatetime df.Timestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *Product) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *Product) SetUpdateProcess(updateProcess string) {
	t.AddPropertyName("updateProcess")
	t.updateProcess = updateProcess
}
func (t *Product) SetUp(){
	
}
func (t *Product)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
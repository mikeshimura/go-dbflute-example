package web

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/mikeshimura/dbflute/df"
	"github.com/mikeshimura/dfweb"
	"godbfexam/adf/bhv"
	"godbfexam/adf/cb"
	"godbfexam/adf/entity"
	"godbfexam/util"
)

func CustomerExec(c *gin.Context) {
	usertable := new(Customer)
	usertable.context = c
	usertable.Exec()
}

type Customer struct {
	session *entity.Session
	login   *entity.Login
	tx      *sql.Tx
	ctx     *df.Context
	context *gin.Context
}

func (p *Customer) Exec() {
	p.session, p.login, p.tx, p.ctx = util.GetSessionLogin(p.context, "godbfexan")
	defer dfweb.ErrorRecover(p.context, p.tx)
	json := dfweb.GetBodyJson(p.context)
	transactions := json["transactions"]
	if transactions != nil {
		p.Transaction(transactions.([]interface{}))
		return
	}
	operationType := (json["operationType"]).(string)
	var data map[string]interface{}
	if df.GetType(json["data"])[0:3] == "map" {
		data = (json["data"]).(map[string]interface{})
	}
	var id float64
	if df.GetType(json["data"]) == "float64" {
		id = (json["data"]).(float64)
	}
	switch operationType {
	case "fetch":
		p.Fetch(data)
	case "add":
		p.Insert(data)
	case "update":
		p.Update(data)
	case "remove":
		p.Delete(id)
	default:
		panic("Customer operationType が無効です。 :" + operationType)
	}
}
func (p *Customer) Transaction(transactions []interface{}) {
	transactionResult := new(df.List)
	for _, transaction := range transactions {
		var tran map[string]interface{} = transaction.(map[string]interface{})
		res := p.executeEachTransaction(tran)
		transactionResult.Add(res)
	}
	p.context.JSON(200, dfweb.SetNormalFetchResult(transactionResult))
}
func (p *Customer) executeEachTransaction(tran map[string]interface{}) map[string]interface{} {
	op := (tran["operationType"]).(string)
	data := (tran["data"]).(map[string]interface{})
	if op == "add" {
		return p.InsertSub(data)
	}
	if op == "update" {
		return p.UpdateSub(data)
	}
	return nil
}
func (p *Customer) Fetch(data map[string]interface{}) {
	cbx := cb.CreateCustomerCB()
	cbx.Query().SetDelFlag_Equal(0)
	p.SetupCriteria(cbx, data)
	//add Order By
	cbx.Query().AddOrderBy_CusCd_Asc()
	res, err := bhv.CustomerBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	list := p.ResultToMap(res)
	p.context.JSON(200, dfweb.SetNormalFetchResult(list))
}
func (p *Customer) Insert(data map[string]interface{}) {
	p.context.JSON(200, p.InsertSub(data))
}
func (p *Customer) InsertSub(data map[string]interface{}) map[string]interface{} {
	p.ctx.Put("Process", "godbfexan:insert")
	entityx := entity.CreateCustomer()
	var e df.Entity = entityx
	dfweb.MapToEntity(data, &e, "Customer", false)
	//DupCheck
	if p.DupCheck(entityx) {
		panic("この CustomerCodeは既に使用されています。")
	}
	_, err := bhv.CustomerBhv_I.Insert(entityx, p.tx, p.ctx)
	if err != nil {
		panic("Insert Error:" + err.Error())
	}
	var ee df.Entity = entityx
	rmap := p.EntityToMap(&ee)
	return dfweb.SetSingleFetchResult(rmap)
}
func (p *Customer) Update(data map[string]interface{}) {
	p.context.JSON(200, p.UpdateSub(data))
}
func (p *Customer) UpdateSub(data map[string]interface{}) map[string]interface{} {
	p.ctx.Put("Process", "godbfexan:update")
	entityx := entity.CreateCustomer()
	var e df.Entity = entityx
	dfweb.MapToEntity(data, &e, "Customer", true)
	//DupCheck
	old := p.getOld(entityx.GetId())
	if old.GetCusCd() != entityx.GetCusCd() &&
		p.DupCheck(entityx) {
		panic("この CustomerCodeは既に使用されています。")
	}
	_, err := bhv.CustomerBhv_I.Update(entityx, p.tx, p.ctx)
	if err != nil {
		panic("Update Error:" + err.Error())
	}
	var ee df.Entity = entityx
	rmap := p.EntityToMap(&ee)
	return dfweb.SetSingleFetchResult(rmap)
}
func (p *Customer) Delete(id float64) {
	p.ctx.Put("Process", "godbfexan:delete")
	entityx := p.getOld(int64(id))
	entityx.SetDelFlag(p.getDelFlagMaxValue(entityx) + 1)
	_, err := bhv.CustomerBhv_I.Update(entityx, p.tx, p.ctx)
	if err != nil {
		panic("Delete Error:" + err.Error())
	}
	var ee df.Entity = entityx
	rmap := p.EntityToMap(&ee)
	p.context.JSON(200, dfweb.SetSingleFetchResult(rmap))
}
func (p *Customer) getDelFlagMaxValue(entityx *entity.Customer) int64 {
	cbx := cb.CreateCustomerCB()
	cbx.Query().SetCusCd_Equal(entityx.GetCusCd())
	cbx.Query().AddOrderBy_DelFlag_Desc()
	cbx.FetchFirst(1)
	res, err := bhv.CustomerBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	if res.AllRecordCount == 1 {
		return (res.List.Get(0)).(*entity.Customer).GetDelFlag()
	}
	panic("DelFlag Max not found")
	return 1
}

func (p *Customer) getOld(id int64) *entity.Customer {
	cbx := cb.CreateCustomerCB()
	cbx.Query().SetDelFlag_Equal(0)
	cbx.Query().SetId_Equal(id)
	res, err := bhv.CustomerBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	if res.AllRecordCount == 1 {
		return res.List.Get(0).(*entity.Customer)
	} else {
		return nil
	}
}

func (p *Customer) SetupCriteria(cbx *cb.CustomerCB, data map[string]interface{}) {
	if data == nil {
		return
	}
	maxRecord := data["maxRecord"]
	if maxRecord != nil {
		cbx.FetchFirst(int((dfweb.ConvFromWebDataInd(maxRecord, "int64")).(int64)))
	}
	criteria := data["criteria"]
	if criteria == nil {
		return
	}
	cr := criteria.([]interface{})
	for _, each := range cr {
		emap := each.(map[string]interface{})
		dfweb.SetCriteria(cbx.Query(), emap, "Customer")
	}
}

func (p *Customer) ResultToMap(res *df.ListResultBean) *df.List {
	reslist := df.CreateList()
	for _, e := range res.List.GetAsArray() {
		entityx := e.(*entity.Customer)
		var entityi df.Entity = entityx
		reslist.Add(p.EntityToMap(&entityi))
	}
	return reslist
}
func (p *Customer) EntityToMap(entityi *df.Entity) map[string]interface{} {
	return dfweb.NewEntityToMap(entityi, []string{"cusCd", "name", "addr", "bldg", "cusConSec", "cusConName", "tel", "salesAmount", "id", "versionNo"})
}
func (p *Customer) DupCheck(entityx *entity.Customer) bool {
	cbx := cb.CreateCustomerCB()
	cbx.Query().SetDelFlag_Equal(0)
	cbx.Query().SetCusCd_Equal(entityx.GetCusCd())
	res, err := bhv.CustomerBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	return res.AllRecordCount == 1
}

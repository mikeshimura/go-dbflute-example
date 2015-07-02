package web

import (
	"database/sql"
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mikeshimura/dbflute/df"
	"github.com/mikeshimura/dfweb"
	"godbfexam/adf/bhv"
	"godbfexam/adf/cb"
	"godbfexam/adf/entity"
	"godbfexam/util"
)

func EmployeeExec(c *gin.Context) {
	usertable := new(Employee)
	usertable.context = c
	usertable.Exec()
}

type Employee struct {
	session *entity.Session
	login   *entity.Login
	tx      *sql.Tx
	ctx     *df.Context
	context *gin.Context
}

func (p *Employee) Exec() {
	p.session, p.login, p.tx, p.ctx = util.GetSessionLogin(p.context, "userTable")
	defer dfweb.ErrorRecover(p.context, p.tx)
	json := dfweb.GetBodyJson(p.context)
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
	case "getSecData":
		p.getSecData()
	default:
		panic("Employee operationType が無効です。 :" + operationType)
	}
}
func (p *Employee) getSecData() {
	cbx := cb.CreateUserTableCB()
	cbx.Query().SetDelFlag_Equal(0)
	cbx.Query().SetTableName_Equal("section")
	cbx.Query().AddOrderBy_Key1_Asc()
	res, err := bhv.UserTableBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	reslist := new(df.List)
	cmap := make(map[string]interface{})
	cmap["value"] = 0
	cmap["label"] = ""
	reslist.Add(cmap)
	for _, ent := range res.List.GetAsArray() {
		entx := ent.(*entity.UserTable)
		cmap := make(map[string]interface{})
		cmap["value"] = entx.GetId()
		cmap["label"] = entx.GetS1Data().String
		reslist.Add(cmap)
	}
	p.context.JSON(200, dfweb.SetNormalFetchResult(reslist))
}

func (p *Employee) Fetch(data map[string]interface{}) {
	cbx := cb.CreateEmployeeCB()
	cbx.SetupSelect_UserTable()
	cbx.Query().SetDelFlag_Equal(0)
	p.SetupCriteria(cbx, data)
	//add Order By
	res, err := bhv.EmployeeBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	list := p.ResultToMap(res)
	p.context.JSON(200, dfweb.SetNormalFetchResult(list))
}
func (p *Employee) Insert(data map[string]interface{}) {
	p.ctx.Put("Process", "employee:insert")
	entityx := entity.CreateEmployee()
	var e df.Entity = entityx
	dfweb.MapToEntity(data, &e, "Employee", false)
	//DupCheck
	if p.DupCheck(entityx) {
		panic("この XXXXは既に使用されています。")
	}
	_, err := bhv.EmployeeBhv_I.Insert(entityx, p.tx, p.ctx)
	if err != nil {
		panic("Insert Error:" + err.Error())
	}
	var ee df.Entity = entityx
	rmap := p.EntityToMap(&ee)
	p.context.JSON(200, dfweb.SetSingleFetchResult(rmap))
}
func (p *Employee) Update(data map[string]interface{}) {
	p.ctx.Put("Process", "employee:update")
	entityx := entity.CreateEmployee()
	var e df.Entity = entityx
	dfweb.MapToEntity(data, &e, "Employee", true)
	//DupCheck
	old := p.getOld(entityx.GetId())
	if entityx.GetEmpCd() != old.GetEmpCd() &&
		p.DupCheck(entityx) {
		panic("この EmpCdは既に使用されています。")
	}
	_, err := bhv.EmployeeBhv_I.Update(entityx, p.tx, p.ctx)
	if err != nil {
		panic("Update Error:" + err.Error())
	}
	ut := p.GetUserTableSec(entityx.GetSecId())
	var ee df.Entity = entityx
	rmap := p.EntityToMap(&ee)
	rmap["sec"] = dfweb.ConvWebData(ut.GetS1Data())
	p.context.JSON(200, dfweb.SetSingleFetchResult(rmap))
}
func (p *Employee) GetUserTableSec(id int64) *entity.UserTable {
	cbx := cb.CreateUserTableCB()
	cbx.Query().SetId_Equal(id)
	res, err := bhv.UserTableBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic("Insert Error:" + err.Error())
	}
	if res.AllRecordCount != 1 {
		return nil
	}
	return (res.List.Get(0)).(*entity.UserTable)
}
func (p *Employee) Delete(id float64) {
	p.ctx.Put("Process", "employee:delete")
	entityx := p.getOld(int64(id))
	entityx.SetDelFlag(p.getDelFlagMaxValue(entityx) + 1)
	_, err := bhv.EmployeeBhv_I.Update(entityx, p.tx, p.ctx)
	if err != nil {
		panic("Delete Error:" + err.Error())
	}
	var ee df.Entity = entityx
	rmap := p.EntityToMap(&ee)
	p.context.JSON(200, dfweb.SetSingleFetchResult(rmap))
}
func (p *Employee) getDelFlagMaxValue(entityx *entity.Employee) int64 {
	cbx := cb.CreateEmployeeCB()
	cbx.Query().SetEmpCd_Equal(entityx.GetEmpCd())
	cbx.Query().AddOrderBy_DelFlag_Desc()
	cbx.FetchFirst(1)
	res, err := bhv.EmployeeBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	if res.AllRecordCount == 1 {
		return (res.List.Get(0)).(*entity.Employee).GetDelFlag()
	}
	panic("DelFlag Max not found")
	return 1
}

func (p *Employee) getOld(id int64) *entity.Employee {
	cbx := cb.CreateEmployeeCB()
	cbx.Query().SetDelFlag_Equal(0)
	cbx.Query().SetId_Equal(id)
	res, err := bhv.EmployeeBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	if res.AllRecordCount == 1 {
		return res.List.Get(0).(*entity.Employee)
	} else {
		return nil
	}
}

func (p *Employee) SetupCriteria(cbx *cb.EmployeeCB, data map[string]interface{}) {
	if data == nil {
		return
	}
	criteria := data["criteria"]
	if criteria == nil {
		return
	}
	cr := criteria.([]interface{})
	for _, each := range cr {
		emap := each.(map[string]interface{})
		fn:=emap["fieldName"].(string)
		if fn == "s1Data" {
			dfweb.SetCriteria(cbx.Query().QueryUserTable(), emap, "UserTable")
		} else {
			dfweb.SetCriteria(cbx.Query(), emap, "Employee")
		}
	}
}

func (p *Employee) ResultToMap(res *df.ListResultBean) *df.List {
	reslist := df.CreateList()
	for _, e := range res.List.GetAsArray() {
		entityx := e.(*entity.Employee)
		var entityi df.Entity = entityx
		emap := p.EntityToMap(&entityi)
		emap["sec"] = dfweb.ConvWebData(entityx.GetUserTable_R().GetS1Data())
		reslist.Add(emap)
	}
	return reslist
}
func (p *Employee) EntityToMap(entityi *df.Entity) map[string]interface{} {
	return dfweb.NewEntityToMap(entityi, []string{"empCd", "secId", "name", "id", "versionNo"})
}
func (p *Employee) DupCheck(entityx *entity.Employee) bool {
	cbx := cb.CreateEmployeeCB()
	cbx.Query().SetDelFlag_Equal(0)
	cbx.Query().SetEmpCd_Equal(entityx.GetEmpCd())
	res, err := bhv.EmployeeBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	return res.AllRecordCount == 1
}

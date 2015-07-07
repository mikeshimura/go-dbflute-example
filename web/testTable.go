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
func TestTableExec(c *gin.Context) {
	usertable:=new(TestTable)
	usertable.context=c
	usertable.Exec()
}
type TestTable struct{
	session *entity.Session
	login *entity.Login
	tx *sql.Tx
	ctx *df.Context
	context *gin.Context
}
func (p *TestTable) Exec() {
	p.session, p.login, p.tx, p.ctx = util.GetSessionLogin(p.context, "testTable")
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
	default:
		panic("TestTable operationType が無効です。 :" + operationType)
	}
}

func (p *TestTable) Fetch(data map[string]interface{}) {
	cbx := cb.CreateTestTableCB()
	cbx.Query().SetDelFlag_Equal(0)
	p.SetupCriteria(cbx, data)
	cbx.Query().AddOrderBy_TestId_Asc()
	res, err := bhv.TestTableBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	list := p.ResultToMap(res)
	p.context.JSON(200, dfweb.SetNormalFetchResult(list))
}
func (p *TestTable) Insert(data map[string]interface{}) {
	p.ctx.Put("Process", "testTable:insert")
	entityx := entity.CreateTestTable()
	var e df.Entity = entityx
	dfweb.MapToEntity(data, &e, "TestTable", false)
	//DupCheck
	if p.DupCheck(entityx) {
		panic("この TestId既に使用されています。")
	}
	_, err := bhv.TestTableBhv_I.Insert(entityx, p.tx, p.ctx)
	if err != nil {
		panic("Insert Error:" + err.Error())
	}
	var ee df.Entity = entityx
	rmap := p.EntityToMap(&ee)
	p.context.JSON(200, dfweb.SetSingleFetchResult(rmap))
}
func (p *TestTable) Update(data map[string]interface{}) {
	p.ctx.Put("Process", "testTable:update")
	entityx := entity.CreateTestTable()
	var e df.Entity = entityx
	dfweb.MapToEntity(data, &e, "TestTable", true)
	//DupCheck
	old := p.getOld(entityx.GetId())
	if (entityx.GetTestId() != old.GetTestId() &&
		p.DupCheck(entityx)) {
		panic("このTestIdは既に使用されています。")
	}
	_, err := bhv.TestTableBhv_I.Update(entityx, p.tx, p.ctx)
	if err != nil {
		panic("Update Error:" + err.Error())
	}
	var ee df.Entity = entityx
	rmap := p.EntityToMap(&ee)
	p.context.JSON(200, dfweb.SetSingleFetchResult(rmap))
}
func (p *TestTable) Delete(id float64) {
	p.ctx.Put("Process", "testTable:delete")
	entityx := p.getOld(int64(id))
	entityx.SetDelFlag(p.getDelFlagMaxValue(entityx) + 1)
	_, err := bhv.TestTableBhv_I.Update(entityx, p.tx, p.ctx)
	if err != nil {
		panic("Delete Error:" + err.Error())
	}
	var ee df.Entity = entityx
	rmap := p.EntityToMap(&ee)
	p.context.JSON(200, dfweb.SetSingleFetchResult(rmap))
}
func (p *TestTable) getDelFlagMaxValue(entityx *entity.TestTable) int64 {
	cbx := cb.CreateTestTableCB()
	cbx.Query().SetTestId_Equal(entityx.GetTestId())
	cbx.Query().AddOrderBy_DelFlag_Desc()
	cbx.FetchFirst(1)
	res, err := bhv.TestTableBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	if res.AllRecordCount == 1 {
		return (res.List.Get(0)).(*entity.TestTable).GetDelFlag()
	}
	panic("DelFlag Max not found")
	return 1
}

func (p *TestTable) getOld(id int64) *entity.TestTable {
	cbx := cb.CreateTestTableCB()
	cbx.Query().SetDelFlag_Equal(0)
	cbx.Query().SetId_Equal(id)
	res, err := bhv.TestTableBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	if res.AllRecordCount == 1 {
		return res.List.Get(0).(*entity.TestTable)
	} else {
		return nil
	}
}

func (p *TestTable) SetupCriteria(cbx *cb.TestTableCB, data map[string]interface{}) {
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
		dfweb.SetCriteria(cbx.Query(), emap, "TestTable")
	}
}

func (p *TestTable) ResultToMap(res *df.ListResultBean) *df.List {
	reslist := df.CreateList()
	for _, e := range res.List.GetAsArray() {
		entityx := e.(*entity.TestTable)
		var entityi df.Entity = entityx
		reslist.Add(p.EntityToMap(&entityi))
	}
	return reslist
}
func (p *TestTable) EntityToMap(entityi *df.Entity) map[string]interface{} {
	return dfweb.NewEntityToMap(entityi, []string{ "testId","testDate","testTimestamp","testNbr","id","versionNo" })
}
func (p *TestTable) DupCheck(entityx *entity.TestTable) bool {
	cbx := cb.CreateTestTableCB()
	cbx.Query().SetDelFlag_Equal(0)
	cbx.Query().SetTestId_Equal(entityx.GetTestId())
	res, err := bhv.TestTableBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	return res.AllRecordCount == 1
}

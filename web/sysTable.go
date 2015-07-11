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
func SysTableExec(c *gin.Context) {
	usertable:=new(SysTable)
	usertable.context=c
	usertable.Exec()
}
type SysTable struct{
	session *entity.Session
	login *entity.Login
	tx *sql.Tx
	ctx *df.Context
	context *gin.Context
}
func (p *SysTable) Exec() {
	p.session, p.login, p.tx, p.ctx = util.GetSessionLogin(p.context, "sysTable")
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
	case "getdbs":
		p.GetDbs()
	default:
		panic("SysTable operationType が無効です。 :" + operationType)
	}
}
func (p *SysTable) GetDbs() {
	cbx := cb.CreateSysTableCB()
	cbx.Query().SetDelFlag_Equal(0)
	cbx.Query().SetTableName_Equal("DbName")
	cbx.Query().AddOrderBy_Key1_Asc()
	res, err := bhv.SysTableBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	reslist:=new(df.List)
	for _,ent:=range res.List.GetAsArray(){
		entx:=ent.(*entity.SysTable)
		cmap:=make(map[string]string)
		cmap["value"]=entx.GetKey1()
		cmap["label"]=entx.GetKey1()
		reslist.Add(cmap)
	}
	p.context.JSON(200, dfweb.SetNormalFetchResult(reslist))
}
func (p *SysTable) Fetch(data map[string]interface{}) {
	cbx := cb.CreateSysTableCB()
	cbx.Query().SetDelFlag_Equal(0)
	p.SetupCriteria(cbx, data)
	cbx.Query().AddOrderBy_TableName_Asc().AddOrderBy_Key1_Asc().AddOrderBy_Key2_Asc()
	res, err := bhv.SysTableBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	list := p.ResultToMap(res)
	p.context.JSON(200, dfweb.SetNormalFetchResult(list))
}
func (p *SysTable) Insert(data map[string]interface{}) {
	p.ctx.Put("Process", "login:insert")
	entityx := entity.CreateSysTable()
	var e df.Entity = entityx
	dfweb.MapToEntity(data, &e, "SysTable", false)
	//DupCheck
	if p.DupCheck(entityx) {
		panic("この tableName,Key1,Key2 は既に使用されています。")
	}
	_, err := bhv.SysTableBhv_I.Insert(entityx, p.tx, p.ctx)
	if err != nil {
		panic("Insert Error:" + err.Error())
	}
	var ee df.Entity = entityx
	rmap := p.EntityToMap(&ee)
	p.context.JSON(200, dfweb.SetSingleFetchResult(rmap))
}
func (p *SysTable) Update(data map[string]interface{}) {
	p.ctx.Put("Process", "login:update")
	entityx := entity.CreateSysTable()
	var e df.Entity = entityx
	dfweb.MapToEntity(data, &e, "SysTable", true)
	//DupCheck
	old := p.getOld(entityx.GetId())
	if (entityx.GetTableName() != old.GetTableName() ||
		entityx.GetKey1() != old.GetKey1() ||
		entityx.GetKey2() != old.GetKey2()) &&
		p.DupCheck(entityx) {
		panic("この tableName,Key1,Key2 は既に使用されています。")
	}
	_, err := bhv.SysTableBhv_I.Update(entityx, p.tx, p.ctx)
	if err != nil {
		panic("Update Error:" + err.Error())
	}
	var ee df.Entity = entityx
	rmap := p.EntityToMap(&ee)
	p.context.JSON(200, dfweb.SetSingleFetchResult(rmap))
}
func (p *SysTable) Delete(id float64) {
	p.ctx.Put("Process", "login:delete")
	entityx := p.getOld(int64(id))
	entityx.SetDelFlag(p.getDelFlagMaxValue(entityx) + 1)
	_, err := bhv.SysTableBhv_I.Update(entityx, p.tx, p.ctx)
	if err != nil {
		panic("Delete Error:" + err.Error())
	}
	var ee df.Entity = entityx
	rmap := p.EntityToMap(&ee)
	p.context.JSON(200, dfweb.SetSingleFetchResult(rmap))
}
func (p *SysTable) getDelFlagMaxValue(entityx *entity.SysTable) int64 {
	cbx := cb.CreateSysTableCB()
	cbx.Query().SetTableName_Equal(entityx.GetTableName())
	cbx.Query().SetKey1_Equal(entityx.GetKey1())
	cbx.Query().SetKey2_Equal(entityx.GetKey2())
	cbx.Query().AddOrderBy_DelFlag_Desc()
	cbx.FetchFirst(1)
	res, err := bhv.SysTableBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	if res.AllRecordCount == 1 {
		return (res.List.Get(0)).(*entity.SysTable).GetDelFlag()
	}
	panic("DelFlag Max not found")
	return 1
}

func (p *SysTable) getOld(id int64) *entity.SysTable {
	cbx := cb.CreateSysTableCB()
	cbx.Query().SetDelFlag_Equal(0)
	cbx.Query().SetId_Equal(id)
	res, err := bhv.SysTableBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	if res.AllRecordCount == 1 {
		return res.List.Get(0).(*entity.SysTable)
	} else {
		return nil
	}
}

func (p *SysTable) SetupCriteria(cbx *cb.SysTableCB, data map[string]interface{}) {
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
		dfweb.SetCriteria(cbx.Query(), emap, "SysTable")
	}
}

func (p *SysTable) ResultToMap(res *df.ListResultBean) *df.List {
	reslist := df.CreateList()
	for _, e := range res.List.GetAsArray() {
		login := e.(*entity.SysTable)
		var entity df.Entity = login
		reslist.Add(p.EntityToMap(&entity))
	}
	return reslist
}
func (p *SysTable) EntityToMap(entity *df.Entity) map[string]interface{} {
	return dfweb.NewEntityToMap(entity, []string{
		"tableName", "key1", "key2", "s1Data", "s2Data", "s3Data",
		"n1Data", "n2Data", "n3Data", "id", "versionNo"})
}
func (p *SysTable) DupCheck(entityx *entity.SysTable) bool {
	cbx := cb.CreateSysTableCB()
	cbx.Query().SetDelFlag_Equal(0)
	cbx.Query().SetTableName_Equal(entityx.GetTableName())
	cbx.Query().SetKey1_Equal(entityx.GetKey1())
	cbx.Query().SetKey2_Equal(entityx.GetKey2())
	res, err := bhv.SysTableBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	return res.AllRecordCount > 0
}

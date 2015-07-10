package web

import (
	"github.com/pborman/uuid"
	"database/sql"
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mikeshimura/dbflute/df"
	"github.com/mikeshimura/dfweb"
	"godbfexam/adf/bhv"
	"godbfexam/adf/cb"
	"godbfexam/adf/entity"
	"godbfexam/util"
	"net/http"
)
func LoginExec(c *gin.Context) {
	login:=new(Login)
	login.context=c
	login.Exec()
}
type Login struct{
	session *entity.Session
	login *entity.Login
	tx *sql.Tx
	ctx *df.Context
	context *gin.Context
}
func (p *Login) Exec () {
	p.session, p.login, p.tx, p.ctx = util.GetSessionLogin(p.context, "login")
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
		p.Fetch( data)
	case "add":
		p.Insert(data)
	case "update":
		p.Update( data)
	case "remove":
		p.Delete(id)
	default:
		panic("Login operationType が無効です。 :" + operationType)
	}
}
func (p *Login) Fetch( data map[string]interface{}) {
	cbx := cb.CreateLoginCB()
	cbx.Query().SetDelFlag_Equal(0)
	p.SetupCriteria(cbx, data)
	cbx.Query().AddOrderBy_LoginId_Asc()
	res, err := bhv.LoginBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	list := p.ResultToMap(res)
	p.context.JSON(200, dfweb.SetNormalFetchResult(list))
}
func (p *Login) Insert(data map[string]interface{}) {
	p.ctx.Put("Process", "login:insert")
	entityx := entity.CreateLogin()
	var e df.Entity = entityx
	dfweb.MapToEntity(data, &e, "Login", false)
	if p.DupCheck(entityx) {
		panic("この 社員番号は既に使用されています。")
	}
	entityx.SetPassword(dfweb.CreateMd5(entityx.GetPassword()))
	_, err := bhv.LoginBhv_I.Insert(entityx, p.tx, p.ctx)
	if err != nil {
		panic("Insert Error:" + err.Error())
	}
	var ee df.Entity = entityx
	rmap := p.EntityToMap(&ee)
	p.context.JSON(200, dfweb.SetSingleFetchResult(rmap))
}
func (p *Login) Update( data map[string]interface{}) {
	p.ctx.Put("Process", "login:update")
	entityx := entity.CreateLogin()
	var e df.Entity = entityx
	dfweb.MapToEntity(data, &e, "Login", true)
	p.convertPassword(entityx)
	old := p.getOld(entityx.GetId())
	if entityx.GetLoginId() != old.GetLoginId() && p.DupCheck(entityx) {
		panic("この 社員番号は既に使用されています。")
	}
	_, err := bhv.LoginBhv_I.Update(entityx, p.tx, p.ctx)
	if err != nil {
		panic("Update Error:" + err.Error())
	}
	var ee df.Entity = entityx
	rmap := p.EntityToMap(&ee)
	p.context.JSON(200, dfweb.SetSingleFetchResult(rmap))
}
func (p *Login) Delete( id float64) {
	p.ctx.Put("Process", "login:delete")
	entityx := p.getOld(int64(id))
	entityx.SetDelFlag(p.getDelFlagMaxValue(entityx) + 1)
	_, err := bhv.LoginBhv_I.Update(entityx, p.tx, p.ctx)
	if err != nil {
		panic("Delete Error:" + err.Error())
	}
	var ee df.Entity = entityx
	rmap := p.EntityToMap(&ee)
	p.context.JSON(200, dfweb.SetSingleFetchResult(rmap))
}
func(p *Login)  getDelFlagMaxValue(entityx *entity.Login) int64 {
	cb := cb.CreateLoginCB()
	cb.Query().SetLoginId_Equal(entityx.GetLoginId())
	cb.Query().AddOrderBy_DelFlag_Desc()
	cb.FetchFirst(1)
	res, err := bhv.LoginBhv_I.SelectList(cb, p.tx)
	if err != nil {
		panic(err.Error())
	}
	if res.AllRecordCount == 1 {
		return (res.List.Get(0)).(*entity.Login).GetDelFlag()
	}
	panic("DelFlag Max not found")
	return 1
}
func(p *Login)  convertPassword(entity *entity.Login) {
	if entity.IsModifiedProperty("password") {
		entity.SetPassword(dfweb.CreateMd5(entity.GetPassword()))
	}
}
func(p *Login)  getOld(id int64) *entity.Login {
	cbx := cb.CreateLoginCB()
	cbx.Query().SetDelFlag_Equal(0)
	cbx.Query().SetId_Equal(id)
	res, err := bhv.LoginBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	if res.AllRecordCount == 1 {
		return res.List.Get(0).(*entity.Login)
	} else {
		return nil
	}
}
func(p *Login)  DupCheck(entity *entity.Login) bool {
	cbx := cb.CreateLoginCB()
	cbx.Query().SetDelFlag_Equal(0)
	cbx.Query().SetLoginId_Equal(entity.GetLoginId())
	res, err := bhv.LoginBhv_I.SelectList(cbx, p.tx)
	if err != nil {
		panic(err.Error())
	}
	return res.AllRecordCount == 1
}
func (p *Login) SetupCriteria(cbx *cb.LoginCB, data map[string]interface{}) {
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
		dfweb.SetCriteria(cbx.Query(), emap, "Login")
	}
}

func (p *Login) ResultToMap(res *df.ListResultBean) *df.List {
	reslist := df.CreateList()
	for _, e := range res.List.GetAsArray() {
		login := e.(*entity.Login)
		var entity df.Entity = login
		reslist.Add(p.EntityToMap(&entity))
	}
	return reslist
}
func(p *Login)  EntityToMap(entity *df.Entity) map[string]interface{} {
	return dfweb.NewEntityToMap(entity, []string{"id", "loginId", "name",
		"versionNo"})
}
func Loginauth(c *gin.Context) {
	login:=new(Login)
	login.LoginauthExec(c)
}
func (p *Login) LoginauthExec(c *gin.Context) {
	tx, err := util.GetTx()
	defer dfweb.ErrorRecover(c, tx)
	if err != nil {
		panic("Get Transaction Error:" + err.Error())
	}
	json := dfweb.GetBodyJson(c)
	login := p.getLoginFromLoginIｄ(json, tx)
	if login == nil {
		c.JSON(200, dfweb.SetErrorMessage("ユーザー IDが見つかりません"))
		return
	}
	password := json["password"]
	if password == nil {
		c.JSON(200, dfweb.SetErrorMessage("パスワードがありません"))
		return
	}
	pwd := dfweb.CreateMd5(password.(string))
	if login.GetPassword() != pwd {
		c.JSON(200, dfweb.SetErrorMessage("passwordが違います"))
		return
	}
	ctx := util.CreateContextFromLogin("loginauth", login)
	session := p.createNewSession(login, tx, ctx)
	loginInfo := p.createLoginInfo(login)
	p.addUuidCookie(c, session.GetUuid())
	c.JSON(200, dfweb.SetSingleFetchResult(loginInfo))
}
func Logout(c *gin.Context) {
	login:=new(Login)
	login.LogoutExec(c)
}
func (p *Login) LogoutExec(c *gin.Context) {
	session, _, tx, ctx := util.GetSessionLogin(c, "logout")
	defer dfweb.ErrorRecover(c, tx)
	session.SetDelFlag(1)
	bhv.SessionBhv_I.Update(session, tx, ctx)
	c.JSON(200, dfweb.SetSingleFetchResult("OK"))
}

func(p *Login)  addUuidCookie(c *gin.Context, uuid string) {
	cookie := new(http.Cookie)
	cookie.Name = "uuid"
	cookie.Value = uuid
	cookie.Path = "/"
	http.SetCookie(c.Writer, cookie)
}
func(p *Login)  createLoginInfo(login *entity.Login) map[string]interface{} {
	rmap := make(map[string]interface{})
	rmap["uid"] = login.GetId()
	rmap["name"] = login.GetName()
	return rmap
}
func(p *Login) createNewSession(login *entity.Login, tx *sql.Tx, ctx *df.Context) *entity.Session {
	uuid := uuid.New()
	session := entity.CreateSession()
	session.SetUuid(uuid)
	session.SetLoginId(df.CreateNullInt64(login.GetId()))
	bhv.SessionBhv_I.Insert(session, tx, ctx)
	return session
}
func (p *Login) getLoginFromLoginIｄ(json map[string]interface{}, tx *sql.Tx) *entity.Login {

	cb := cb.CreateLoginCB()
	loginId := json["loginId"]
	if loginId == nil {
		return nil
	}
	cb.Query().SetLoginId_Equal((loginId).(string))
	cb.Query().SetDelFlag_Equal(0)
	res, err2 := bhv.LoginBhv_I.SelectList(cb, tx)

	if err2 != nil {
		panic(err2.Error())
	}
	if res.AllRecordCount == 1 {
		return (res.List.Get(0)).(*entity.Login)
	}
	return nil
}

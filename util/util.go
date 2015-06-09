package util

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/mikeshimura/dbflute/df"
	"github.com/mikeshimura/dfweb"
	"mikegou/adf/bhv"
	"mikegou/adf/cb"
	"mikegou/adf/entity"
	"strconv"
)

var Db *sql.DB
var Sslmode string
func CreateContextFromLogin(proc string, login *entity.Login) *df.Context {
	ctx := df.CreateContext()
	ctx.Put("Process", proc)
	ctx.Put("User", login.GetLoginId())
	return ctx
}
func GetLoginFromSession(session *entity.Session, tx *sql.Tx) *entity.Login {
	cb := cb.CreateLoginCB()
	cb.Query().SetId_Equal(session.GetLoginId().Int64)
	cb.Query().SetDelFlag_Equal(0)
	res, err := bhv.LoginBhv_I.SelectList(cb, tx)
	if err != nil {
		panic("Login Select Error:" + err.Error())
	}
	if res.AllRecordCount != 1 {
		panic("Login Not Found")
	}
	return (res.List.Get(0)).(*entity.Login)
}
func GetSessionLogin(c *gin.Context, proc string) (*entity.Session, *entity.Login, *sql.Tx, *df.Context) {
	tx, err := GetTx()
	if err != nil {
		panic("Get Transaction Error:" + err.Error())
	}
	session := GetSessionFromRequestCookie(c, tx)
	if session == nil {
		panic("セッションがありません。再度ログインして下さい。")
	}
	login := GetLoginFromSession(session, tx)
	if login == nil {
		panic("ログインが無効です。再度ログインして下さい。")
	}
	ctx := CreateContextFromLogin(proc, login)
	return session, login, tx, ctx
}
func GetTx() (*sql.Tx, error) {
	db, err := GetDb()
	if err != nil {
		return nil, err
	}
	tx, err := df.TxBegin(db)
	return tx, err
}
func GetSessionFromRequestCookie(c *gin.Context, tx *sql.Tx) *entity.Session {
	cookie, err := c.Request.Cookie("uuid")
	if err != nil {
		return nil
	}
	uuid := cookie.Value
	uuid = uuid
	cb := cb.CreateSessionCB()
	cb.Query().SetUuid_Equal(uuid)
	cb.Query().SetDelFlag_Equal(0)
	res, err2 := bhv.SessionBhv_I.SelectList(cb, tx)
	if err2 != nil {
		panic("Session Select Error:" + err2.Error())
	}
	if res.AllRecordCount != 1 {
		return nil
	}
	return (res.List.Get(0)).(*entity.Session)
}

func GetDb() (*sql.DB, error) {
	if Db == nil {
		var err error
		Db, err = dfweb.PgOpenDatabaseUrl(Sslmode) 
		if err != nil {
			return nil, err
		}
	}
	return Db, nil
}
func CheckAuth(screen string, session *entity.Session, login *entity.Login, tx *sql.Tx) bool {
	return true
}
func CreateScript(login *entity.Login, script []string) []string {
	slist := new(df.StringList)
	slist.Add("<script>")
	slist.Add("if (!(window[\"$c\"] != null)) {window[\"$c\"] = {};}")
	slist.Add("$c.login = {};")
	slist.Add("$c.login.uid = " + strconv.Itoa(int(login.GetId())) + ";")
	slist.Add("$c.login.name = '" + login.GetName() + "';")
	for _, s := range script {
		slist.Add(s)
	}
	slist.Add("</script>\n")
	return slist.GetAsArray()
}

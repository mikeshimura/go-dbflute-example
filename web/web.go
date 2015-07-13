package web

import (
	//"fmt"
	//"database/sql"
	"github.com/gin-gonic/gin"
	"godbfexam/util"
)

func IndexHtml(c *gin.Context) {
	html := createHtml("メニュー画面", []string{"/js/$c.js", "/js/$v.js", "/js/index.js", "/js/_cjsx.js", "/js/indexjsx.js"},
		[]string{"/js/lib/fluxxor.js", "/js/lib/react.js", "/js/lib/react-bootstrap.js",
			"/js/lib/jquery-1.11.1.js", "/js/lib/lodash.js"},
		[]string{"/css/bootstrap.css", "/css/main.css"}, []string{})
	c.Header("Content-Type", "text/html")
	c.String(200,html)
}

func UserHtml(c *gin.Context) {
	title := "USER管理"
	js := []string{"/js/$c.js", "/js/$v.js", "/js/user.js", "/js/_cjsx.js", "/js/userjsx.js"}
	jslib := []string{"/js/lib/fluxxor.js", "/js/lib/react.js", "/js/lib/react-bootstrap.js",
		"/js/lib/jquery-1.11.1.js", "/js/lib/lodash.js"}
	css := []string{"/css/bootstrap.css", "/css/main.css"}
	GenerateAuthorizedScreen(c, title, js, jslib, css, []string{})
}
func LoginHtml(c *gin.Context) {
	title := "Login管理"
	js := []string{"/js/$c.js", "/js/$v.js", "/js/login.js", "/js/_cjsx.js", "/js/loginjsx.js"}
	jslib := []string{"/js/lib/fluxxor.js", "/js/lib/react.js", "/js/lib/react-bootstrap.js",
		"/js/lib/jquery-1.11.1.js", "/js/lib/lodash.js"}
	css := []string{"/css/bootstrap.css", "/css/main.css"}
	GenerateAuthorizedScreen(c, title, js, jslib, css, []string{})
}
func UserTblHtml(c *gin.Context) {
	title := "USER TABLE"
	js := []string{"/js/$c.js", "/js/$v.js", "/js/usertbl.js", "/js/_cjsx.js", "/js/usertbljsx.js"}
	jslib := []string{"/js/lib/fluxxor.js", "/js/lib/react.js", "/js/lib/react-bootstrap.js",
		"/js/lib/jquery-1.11.1.js", "/js/lib/lodash.js"}
	css := []string{"/css/bootstrap.css", "/css/main.css"}
	GenerateAuthorizedScreen(c, title, js, jslib, css, []string{})
}
func SysTblHtml(c *gin.Context) {
	title := "SYSTEM TABLE"
	js := []string{"/js/$c.js", "/js/$v.js", "/js/systbl.js", "/js/_cjsx.js", "/js/systbljsx.js"}
	jslib := []string{"/js/lib/fluxxor.js", "/js/lib/react.js", "/js/lib/react-bootstrap.js",
		"/js/lib/jquery-1.11.1.js", "/js/lib/lodash.js"}
	css := []string{"/css/bootstrap.css", "/css/main.css"}
	GenerateAuthorizedScreen(c, title, js, jslib, css, []string{})
}

func EmployeeHtml(c *gin.Context) {
	title := "従業員マスター"
	js := []string{"/js/$c.js", "/js/$v.js", "/js/employee.js", "/js/_cjsx.js", "/js/employeejsx.js"}
	jslib := []string{"/js/lib/fluxxor.js", "/js/lib/react.js", "/js/lib/react-bootstrap.js",
		"/js/lib/jquery-1.11.1.js", "/js/lib/lodash.js"}
	css := []string{"/css/bootstrap.css", "/css/main.css"}
	GenerateAuthorizedScreen(c, title, js, jslib, css, []string{})
}

func CustomerHtml(c *gin.Context) {
	title := "CUSTOMER管理"
	js := []string{"/js/$c.js", "/js/$v.js", "/js/customer.js", "/js/_cjsx.js", "/js/customerjsx.js"}
	jslib := []string{"/js/lib/fluxxor.js", "/js/lib/react.js", "/js/lib/react-bootstrap.js",
		"/js/lib/jquery-1.11.1.js", "/js/lib/lodash.js"}
	css := []string{"/css/bootstrap.css", "/css/main.css"}
	GenerateAuthorizedScreen(c, title, js, jslib, css, []string{})
}
func CustomertHtml(c *gin.Context) {
	title := "CUSTOMER管理"
	js := []string{"/js/$c.js", "/js/$v.js", "/js/customert.js", "/js/_cjsx.js", "/js/customertjsx.js"}
	jslib := []string{"/js/lib/fluxxor.js", "/js/lib/react.js", "/js/lib/react-bootstrap.js",
		"/js/lib/jquery-1.11.1.js", "/js/lib/lodash.js"}
	css := []string{"/css/bootstrap.css", "/css/main.css"}
	GenerateAuthorizedScreen(c, title, js, jslib, css, []string{})
}
func CustomerinHtml(c *gin.Context) {
	title := "CUSTOMER管理"
	js := []string{"/js/$c.js", "/js/$v.js", "/js/customerin.js", "/js/_cjsx.js", "/js/customerinjsx.js"}
	jslib := []string{"/js/lib/fluxxor.js", "/js/lib/react.js", "/js/lib/react-bootstrap.js",
		"/js/lib/jquery-1.11.1.js", "/js/lib/lodash.js"}
	css := []string{"/css/bootstrap.css", "/css/main.css"}
	GenerateAuthorizedScreen(c, title, js, jslib, css, []string{})
}

func TestTableHtml(c *gin.Context) {
	title := "TEST TABLE 管理"
	js := []string{"/js/$c.js", "/js/$v.js", "/js/testtable.js", "/js/_cjsx.js", "/js/testtablejsx.js"}
	jslib := []string{"/js/lib/fluxxor.js", "/js/lib/react.js", "/js/lib/react-bootstrap.js",
		"/js/lib/jquery-1.11.1.js", "/js/lib/lodash.js"}
	css := []string{"/css/bootstrap.css", "/css/main.css"}
	GenerateAuthorizedScreen(c, title, js, jslib, css, []string{})
}
func GenerateAuthorizedScreen(
	c *gin.Context, title string, js []string, jslib []string, css []string, script []string) {
	defer func() {
		errx := recover()
		if errx != nil {
			c.String(200, errx.(string))
		}
	}()
	screen := GetScreen(c)
	session, login, tx, _ := util.GetSessionLogin(c, "web")
	if util.CheckAuth(screen, session, login, tx) == false {
		panic("この画面は権限がありません。")
	}
	scripts := util.CreateScript(login, script)
	html := createHtml(title, js, jslib, css, scripts)
	c.Header("Content-Type", "text/html")
	c.String(200, html)
}
func GetScreen(c *gin.Context) string {
	path := c.Request.URL.Path
	return path[1:]
}

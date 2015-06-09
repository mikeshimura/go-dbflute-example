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
	c.HTMLString(200, html)
}

func TemplateHtml(c *gin.Context) {
	title := "TEMPLATE"
	js := []string{"/js/$c.js", "/js/$v.js", "/js/template.js", "/js/_cjsx.js", "/js/templatejsx.js"}
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
	c.HTMLString(200, html)
}
func GetScreen(c *gin.Context) string {
	path := c.Request.URL.Path
	return path[1:]
}

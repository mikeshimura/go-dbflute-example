package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mikeshimura/dbflute/df"
	"github.com/mikeshimura/dbflute/log"
	"godbfexam/util"
	"godbfexam/web"
	"os"
)

func main() {
	df.DISP_SQL_DEFAULT_TIMESTAMP_FORMAT = "2006/01/02 15:04:05"
	df.DISP_SQL_DEFAULT_DATE_FORMAT = "2006/01/02"
	log.SetCharCode("utf-8")
	util.Sslmode = "disable"
	r := gin.Default()
	r.GET("/", web.IndexHtml)
	r.GET("/user", web.UserHtml)
	r.GET("/usertbl", web.UserTblHtml)
	r.GET("/systbl", web.SysTblHtml)
	r.GET("/employee", web.EmployeeHtml)
	r.GET("/login", web.LoginHtml)
	r.GET("/customer", web.CustomerHtml)
	r.GET("/customert", web.CustomertHtml)
	r.GET("/customerin", web.CustomerinHtml)
	r.GET("/testtable", web.TestTableHtml)
	ajax := r.Group("/ajax")

	{
		ajax.POST("/loginauth", web.Loginauth)
		ajax.POST("/logout", web.Logout)
		ajax.POST("/login", web.LoginExec)
		ajax.POST("/usertbl", web.UserTableExec)
		ajax.POST("/systbl", web.SysTableExec)
		ajax.POST("/employee", web.EmployeeExec)
		ajax.POST("/customer", web.CustomerExec)
		ajax.POST("/testtable", web.TestTableExec)
	}
	r.Static("/static", "./static")
	r.Run(":" + os.Getenv("PORT"))
}

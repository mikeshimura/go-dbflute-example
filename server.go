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
	ajax := r.Group("/ajax")

	{
		ajax = ajax
	}
	r.Static("/static", "./static")
	r.Run(":" + os.Getenv("PORT"))
}

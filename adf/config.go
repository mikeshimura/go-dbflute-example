package adf

import (
	"github.com/mikeshimura/dbflute/df"
)
func DBCurrentCreate() {
	di := new(df.DBCurrent)
	di.ProjectName = "postgresql"
	di.ProjectPrefix = ""
	di.PagingCountLater = true
	di.PagingCountLeastJoin = true
	di.InnerJoinAutoDetect =  true
	di.ThatsBadTimingDetect = true
	di.NullOrEmptyQueryAllowed = false
	di.EmptyStringQueryAllowed = false
	di.EmptyStringParameterAllowed = false
	di.OverridingQueryAllowed = false
	di.NonSpecifiedColumnAccessAllowed = false
	di.ColumnNullObjectAllowed = false
	di.ColumnNullObjectGearedToSpecify = false
	di.DisableSelectIndex = false
	di.QueryUpdateCountPreCheck = false
	var dw df.DBWay = new(df.WayOfPostgreSQL)
	di.DBWay = &dw
	var dd df.DBDef = new(df.PostgreSQL)
	di.DBDef = &dd
	df.DBCurrent_I = di
}



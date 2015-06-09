package cmeta

import (
	"github.com/mikeshimura/dbflute/df"
)

func init() {
	Create_C_DatabaseSelectDbm()
	var DatabaseSelect_meta df.DBMeta = C_DatabaseSelectDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["C_DatabaseSelect"] = &DatabaseSelect_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("C_DatabaseSelect", "C_DatabaseSelect")
	Create_C_DbTableSelectDbm()
	var DbTableSelect_meta df.DBMeta = C_DbTableSelectDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["C_DbTableSelect"] = &DbTableSelect_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("C_DbTableSelect", "C_DbTableSelect")
}

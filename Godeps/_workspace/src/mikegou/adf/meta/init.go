package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

func MetaInit() {
	Create_DbfluteDbm()
	var Dbflute_meta df.DBMeta = DbfluteDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["Dbflute"] = &Dbflute_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("Dbflute", "Dbflute")
	Create_LoginDbm()
	var Login_meta df.DBMeta = LoginDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["Login"] = &Login_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("Login", "Login")
	Create_SessionDbm()
	var Session_meta df.DBMeta = SessionDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["Session"] = &Session_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("Session", "Session")
	Create_SysTableDbm()
	var SysTable_meta df.DBMeta = SysTableDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["SysTable"] = &SysTable_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("SysTable", "SysTable")
	Create_UserTableDbm()
	var UserTable_meta df.DBMeta = UserTableDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["UserTable"] = &UserTable_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("UserTable", "UserTable")
}

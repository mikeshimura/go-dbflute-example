package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

func MetaInit() {
	Create_CustomerDbm()
	var Customer_meta df.DBMeta = CustomerDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["Customer"] = &Customer_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("Customer", "Customer")
	Create_EmployeeDbm()
	var Employee_meta df.DBMeta = EmployeeDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["Employee"] = &Employee_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("Employee", "Employee")
	Create_LoginDbm()
	var Login_meta df.DBMeta = LoginDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["Login"] = &Login_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("Login", "Login")
	Create_ProductDbm()
	var Product_meta df.DBMeta = ProductDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["Product"] = &Product_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("Product", "Product")
	Create_SalesSlipDbm()
	var SalesSlip_meta df.DBMeta = SalesSlipDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["SalesSlip"] = &SalesSlip_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("SalesSlip", "SalesSlip")
	Create_SessionDbm()
	var Session_meta df.DBMeta = SessionDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["Session"] = &Session_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("Session", "Session")
	Create_StockDbm()
	var Stock_meta df.DBMeta = StockDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["Stock"] = &Stock_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("Stock", "Stock")
	Create_SysTableDbm()
	var SysTable_meta df.DBMeta = SysTableDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["SysTable"] = &SysTable_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("SysTable", "SysTable")
	Create_UserTableDbm()
	var UserTable_meta df.DBMeta = UserTableDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["UserTable"] = &UserTable_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("UserTable", "UserTable")
}
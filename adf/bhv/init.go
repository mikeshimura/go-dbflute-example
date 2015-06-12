package bhv

import (
	"godbfexam/adf"
	"github.com/mikeshimura/dbflute/df"
	"godbfexam/adf/entity"
	"godbfexam/adf/meta"
	_	"github.com/lib/pq"  //Don't eraze require to use this package
)
//To fix init sequence, all init routines are called from here.
func init() {
	adf.AdfInit()
	entity.EntityInit()
	meta.MetaInit()
	CustomerBhv_I = new(CustomerBhv)
	CustomerBhv_I.BaseBehavior = new(df.BaseBehavior)
	CustomerBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	CustomerBhv_I.BaseBehavior.TableDbName = "Customer"
	var customer df.Behavior =CustomerBhv_I
	CustomerBhv_I.BaseBehavior.Behavior=&customer
	CustomerBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	EmployeeBhv_I = new(EmployeeBhv)
	EmployeeBhv_I.BaseBehavior = new(df.BaseBehavior)
	EmployeeBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	EmployeeBhv_I.BaseBehavior.TableDbName = "Employee"
	var employee df.Behavior =EmployeeBhv_I
	EmployeeBhv_I.BaseBehavior.Behavior=&employee
	EmployeeBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	LoginBhv_I = new(LoginBhv)
	LoginBhv_I.BaseBehavior = new(df.BaseBehavior)
	LoginBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	LoginBhv_I.BaseBehavior.TableDbName = "Login"
	var login df.Behavior =LoginBhv_I
	LoginBhv_I.BaseBehavior.Behavior=&login
	LoginBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	ProductBhv_I = new(ProductBhv)
	ProductBhv_I.BaseBehavior = new(df.BaseBehavior)
	ProductBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	ProductBhv_I.BaseBehavior.TableDbName = "Product"
	var product df.Behavior =ProductBhv_I
	ProductBhv_I.BaseBehavior.Behavior=&product
	ProductBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	SalesSlipBhv_I = new(SalesSlipBhv)
	SalesSlipBhv_I.BaseBehavior = new(df.BaseBehavior)
	SalesSlipBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	SalesSlipBhv_I.BaseBehavior.TableDbName = "SalesSlip"
	var salesSlip df.Behavior =SalesSlipBhv_I
	SalesSlipBhv_I.BaseBehavior.Behavior=&salesSlip
	SalesSlipBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	SessionBhv_I = new(SessionBhv)
	SessionBhv_I.BaseBehavior = new(df.BaseBehavior)
	SessionBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	SessionBhv_I.BaseBehavior.TableDbName = "Session"
	var session df.Behavior =SessionBhv_I
	SessionBhv_I.BaseBehavior.Behavior=&session
	SessionBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	StockBhv_I = new(StockBhv)
	StockBhv_I.BaseBehavior = new(df.BaseBehavior)
	StockBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	StockBhv_I.BaseBehavior.TableDbName = "Stock"
	var stock df.Behavior =StockBhv_I
	StockBhv_I.BaseBehavior.Behavior=&stock
	StockBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	SysTableBhv_I = new(SysTableBhv)
	SysTableBhv_I.BaseBehavior = new(df.BaseBehavior)
	SysTableBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	SysTableBhv_I.BaseBehavior.TableDbName = "SysTable"
	var sysTable df.Behavior =SysTableBhv_I
	SysTableBhv_I.BaseBehavior.Behavior=&sysTable
	SysTableBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	UserTableBhv_I = new(UserTableBhv)
	UserTableBhv_I.BaseBehavior = new(df.BaseBehavior)
	UserTableBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	UserTableBhv_I.BaseBehavior.TableDbName = "UserTable"
	var userTable df.Behavior =UserTableBhv_I
	UserTableBhv_I.BaseBehavior.Behavior=&userTable
	UserTableBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
}
package bhv

import (
	"mikegou/adf"
	"github.com/mikeshimura/dbflute/df"
	"mikegou/adf/entity"
	"mikegou/adf/meta"
	_	"github.com/lib/pq"  //Don't eraze require to use this package
)
//To fix init sequence, all init routines are called from here.
func init() {
	adf.AdfInit()
	entity.EntityInit()
	meta.MetaInit()
	DbfluteBhv_I = new(DbfluteBhv)
	DbfluteBhv_I.BaseBehavior = new(df.BaseBehavior)
	DbfluteBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	DbfluteBhv_I.BaseBehavior.TableDbName = "Dbflute"
	var dbflute df.Behavior =DbfluteBhv_I
	DbfluteBhv_I.BaseBehavior.Behavior=&dbflute
	DbfluteBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	LoginBhv_I = new(LoginBhv)
	LoginBhv_I.BaseBehavior = new(df.BaseBehavior)
	LoginBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	LoginBhv_I.BaseBehavior.TableDbName = "Login"
	var login df.Behavior =LoginBhv_I
	LoginBhv_I.BaseBehavior.Behavior=&login
	LoginBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	SessionBhv_I = new(SessionBhv)
	SessionBhv_I.BaseBehavior = new(df.BaseBehavior)
	SessionBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	SessionBhv_I.BaseBehavior.TableDbName = "Session"
	var session df.Behavior =SessionBhv_I
	SessionBhv_I.BaseBehavior.Behavior=&session
	SessionBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
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

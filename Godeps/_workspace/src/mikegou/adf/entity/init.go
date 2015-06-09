package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

func EntityInit() {
	Dbflute := func() *df.Entity {
		var te df.Entity = new(Dbflute)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Dbflute", Dbflute)
	Login := func() *df.Entity {
		var te df.Entity = new(Login)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Login", Login)
	Session := func() *df.Entity {
		var te df.Entity = new(Session)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Session", Session)
	SysTable := func() *df.Entity {
		var te df.Entity = new(SysTable)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("SysTable", SysTable)
	UserTable := func() *df.Entity {
		var te df.Entity = new(UserTable)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("UserTable", UserTable)
}

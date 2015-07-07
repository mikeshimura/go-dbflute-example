package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

func EntityInit() {
	Customer := func() *df.Entity {
		var te df.Entity = new(Customer)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Customer", Customer)
	Employee := func() *df.Entity {
		var te df.Entity = new(Employee)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Employee", Employee)
	Login := func() *df.Entity {
		var te df.Entity = new(Login)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Login", Login)
	Product := func() *df.Entity {
		var te df.Entity = new(Product)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Product", Product)
	SalesSlip := func() *df.Entity {
		var te df.Entity = new(SalesSlip)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("SalesSlip", SalesSlip)
	Session := func() *df.Entity {
		var te df.Entity = new(Session)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Session", Session)
	Stock := func() *df.Entity {
		var te df.Entity = new(Stock)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Stock", Stock)
	SysTable := func() *df.Entity {
		var te df.Entity = new(SysTable)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("SysTable", SysTable)
	TestTable := func() *df.Entity {
		var te df.Entity = new(TestTable)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("TestTable", TestTable)
	UserTable := func() *df.Entity {
		var te df.Entity = new(UserTable)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("UserTable", UserTable)
}
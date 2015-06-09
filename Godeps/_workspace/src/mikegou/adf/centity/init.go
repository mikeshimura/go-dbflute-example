package centity

import (
	"github.com/mikeshimura/dbflute/df"
)

func init() {
	DatabaseSelect := func() *df.Entity {
		var te df.Entity = new(C_DatabaseSelect)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("C_DatabaseSelect", DatabaseSelect)
	DbTableSelect := func() *df.Entity {
		var te df.Entity = new(C_DbTableSelect)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("C_DbTableSelect", DbTableSelect)
}

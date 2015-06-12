package adf

import (
	"github.com/mikeshimura/dbflute/df"
)

func AdfInit() {
	DBCurrentCreate()
	df.Bci = new(df.BehaviorCommandInvoker)
	impl := new(InvokerAssistantImpl)
	InvokerAssistant_I = impl
	InvokerAssistant_I.Create()
	df.Bci.InvokerAssistant = &InvokerAssistant_I
	setupper:=new(CommonColumnAutoSetupperImpl)
	var su df.CommonColumnAutoSetupper = setupper
	setupper.CommonColumnAutoSetupper=&su
	df.CommonColumnAutoSetupper_I=&su
}
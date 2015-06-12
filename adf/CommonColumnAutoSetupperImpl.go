package adf

import (
	"github.com/mikeshimura/dbflute/df"
	"time"
)

type CommonColumnAutoSetupperImpl struct {
	df.BaseCommonColumnAutoSetupper
}

func (p *CommonColumnAutoSetupperImpl) DoHandleCommonColumnOfInsertIfNeeds(
	entity *df.Entity, context *df.Context) {
	tempRegisterDatetime:= df.CreateTimestamp(time.Now())
	if &tempRegisterDatetime != nil {
		df.SetEntityValue(entity, "registerDatetime", tempRegisterDatetime)
	}
	tempRegisterUser:= context.Get("User")
	if &tempRegisterUser != nil {
		df.SetEntityValue(entity, "registerUser", tempRegisterUser)
	}
	tempRegisterProcess:= context.Get("Process")
	if &tempRegisterProcess != nil {
		df.SetEntityValue(entity, "registerProcess", tempRegisterProcess)
	}
	tempUpdateDatetime:= df.GetEntityValue(entity, "registerDatetime")
	if &tempUpdateDatetime != nil {
		df.SetEntityValue(entity, "updateDatetime", tempUpdateDatetime)
	}
	tempUpdateUser:= df.GetEntityValue(entity, "registerUser")
	if &tempUpdateUser != nil {
		df.SetEntityValue(entity, "updateUser", tempUpdateUser)
	}
	tempUpdateProcess:= df.GetEntityValue(entity, "registerProcess")
	if &tempUpdateProcess != nil {
		df.SetEntityValue(entity, "updateProcess", tempUpdateProcess)
	}

}
func (p *CommonColumnAutoSetupperImpl) DoHandleCommonColumnOfUpdateIfNeeds(
	entity *df.Entity, context *df.Context) {
	tempUpdateDatetime:= df.CreateTimestamp(time.Now())
	if &tempUpdateDatetime != nil {
		df.SetEntityValue(entity, "updateDatetime", tempUpdateDatetime)
	}
	tempUpdateUser:= context.Get("User")
	if &tempUpdateUser != nil {
		df.SetEntityValue(entity, "updateUser", tempUpdateUser)
	}
	tempUpdateProcess:= context.Get("Process")
	if &tempUpdateProcess != nil {
		df.SetEntityValue(entity, "updateProcess", tempUpdateProcess)
	}
}
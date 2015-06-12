package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type EmployeeCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	Id *df.ConditionValue
	EmpCd *df.ConditionValue
	SecId *df.ConditionValue
	Name *df.ConditionValue
	VersionNo *df.ConditionValue
	DelFlag *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterUser *df.ConditionValue
	RegisterProcess *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateUser *df.ConditionValue
	UpdateProcess *df.ConditionValue
    conditionQueryUserTable *UserTableCQ
}

func (q *EmployeeCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *EmployeeCQ) getCValueId() *df.ConditionValue {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	return q.Id
}



func (q *EmployeeCQ) SetId_Equal(value int64) *EmployeeCQ {
	q.regId(df.CK_EQ_C, value)
	return q
}
func (q *EmployeeCQ) SetId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueId(), "id")
}
func (q *EmployeeCQ) SetId_NotEqual(value int64) *EmployeeCQ {
	q.regId(df.CK_NE_C, value)
	return q
}

func (q *EmployeeCQ) SetId_GreaterThan(value int64) *EmployeeCQ {
	q.regId(df.CK_GT_C, value)
	return q
}

func (q *EmployeeCQ) SetId_LessThan(value int64) *EmployeeCQ {
	q.regId(df.CK_LT_C, value)
	return q
}

func (q *EmployeeCQ) SetId_GreaterEqual(value int64) *EmployeeCQ {
	q.regId(df.CK_GE_C, value)
	return q
}

func (q *EmployeeCQ) SetId_LessEqual(value int64) *EmployeeCQ {
	q.regId(df.CK_LE_C, value)
	return q
}
func (q *EmployeeCQ) SetId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueId(),"id",rangeOfOption)
}	


func (q *EmployeeCQ) SetId_IsNull() *EmployeeCQ {
	q.regId(df.CK_ISN_C, 0)
	return q
}
func (q *EmployeeCQ) SetId_IsNotNull() *EmployeeCQ {
	q.regId(df.CK_ISNN_C, 0)
	return q
}
func (q *EmployeeCQ) AddOrderBy_Id_Asc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBA("id")
	return q
}
func (q *EmployeeCQ) AddOrderBy_Id_Desc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBD("id")
	return q
}
func (q *EmployeeCQ) regId(key *df.ConditionKey, value interface{}) {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Id, "id")
}

func (q *EmployeeCQ) getCValueEmpCd() *df.ConditionValue {
	if q.EmpCd == nil {
		q.EmpCd = new(df.ConditionValue)
	}
	return q.EmpCd
}


func (q *EmployeeCQ) SetEmpCd_Equal(value string) *EmployeeCQ {
	q.regEmpCd(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *EmployeeCQ) SetEmpCd_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueEmpCd(), "empCd")
}
func (q *EmployeeCQ) SetEmpCd_NotEqual(value string) *EmployeeCQ {
	q.regEmpCd(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetEmpCd_GreaterThan(value string) *EmployeeCQ {
	q.regEmpCd(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetEmpCd_LessThan(value string) *EmployeeCQ {
	q.regEmpCd(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetEmpCd_GreaterEqual(value string) *EmployeeCQ {
	q.regEmpCd(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *EmployeeCQ) SetEmpCd_LessEqual(value string) *EmployeeCQ {
	q.regEmpCd(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetEmpCd_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueEmpCd(), "empCd", option)
}

func (q *EmployeeCQ) SetEmpCd_PrefixSearch(value string) error {
	return q.SetEmpCd_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *EmployeeCQ) SetEmpCd_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueEmpCd(), "empCd", option)
}



func (q *EmployeeCQ) AddOrderBy_EmpCd_Asc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBA("empCd")
	return q
}
func (q *EmployeeCQ) AddOrderBy_EmpCd_Desc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBD("empCd")
	return q
}
func (q *EmployeeCQ) regEmpCd(key *df.ConditionKey, value interface{}) {
	if q.EmpCd == nil {
		q.EmpCd = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.EmpCd, "empCd")
}

func (q *EmployeeCQ) getCValueSecId() *df.ConditionValue {
	if q.SecId == nil {
		q.SecId = new(df.ConditionValue)
	}
	return q.SecId
}



func (q *EmployeeCQ) SetSecId_Equal(value int64) *EmployeeCQ {
	q.regSecId(df.CK_EQ_C, value)
	return q
}
func (q *EmployeeCQ) SetSecId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueSecId(), "secId")
}
func (q *EmployeeCQ) SetSecId_NotEqual(value int64) *EmployeeCQ {
	q.regSecId(df.CK_NE_C, value)
	return q
}

func (q *EmployeeCQ) SetSecId_GreaterThan(value int64) *EmployeeCQ {
	q.regSecId(df.CK_GT_C, value)
	return q
}

func (q *EmployeeCQ) SetSecId_LessThan(value int64) *EmployeeCQ {
	q.regSecId(df.CK_LT_C, value)
	return q
}

func (q *EmployeeCQ) SetSecId_GreaterEqual(value int64) *EmployeeCQ {
	q.regSecId(df.CK_GE_C, value)
	return q
}

func (q *EmployeeCQ) SetSecId_LessEqual(value int64) *EmployeeCQ {
	q.regSecId(df.CK_LE_C, value)
	return q
}
func (q *EmployeeCQ) SetSecId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueSecId(),"secId",rangeOfOption)
}	


func (q *EmployeeCQ) AddOrderBy_SecId_Asc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBA("secId")
	return q
}
func (q *EmployeeCQ) AddOrderBy_SecId_Desc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBD("secId")
	return q
}
func (q *EmployeeCQ) regSecId(key *df.ConditionKey, value interface{}) {
	if q.SecId == nil {
		q.SecId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SecId, "secId")
}

func (q *EmployeeCQ) getCValueName() *df.ConditionValue {
	if q.Name == nil {
		q.Name = new(df.ConditionValue)
	}
	return q.Name
}


func (q *EmployeeCQ) SetName_Equal(value string) *EmployeeCQ {
	q.regName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *EmployeeCQ) SetName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueName(), "name")
}
func (q *EmployeeCQ) SetName_NotEqual(value string) *EmployeeCQ {
	q.regName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetName_GreaterThan(value string) *EmployeeCQ {
	q.regName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetName_LessThan(value string) *EmployeeCQ {
	q.regName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetName_GreaterEqual(value string) *EmployeeCQ {
	q.regName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *EmployeeCQ) SetName_LessEqual(value string) *EmployeeCQ {
	q.regName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueName(), "name", option)
}

func (q *EmployeeCQ) SetName_PrefixSearch(value string) error {
	return q.SetName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *EmployeeCQ) SetName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueName(), "name", option)
}



func (q *EmployeeCQ) AddOrderBy_Name_Asc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBA("name")
	return q
}
func (q *EmployeeCQ) AddOrderBy_Name_Desc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBD("name")
	return q
}
func (q *EmployeeCQ) regName(key *df.ConditionKey, value interface{}) {
	if q.Name == nil {
		q.Name = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Name, "name")
}

func (q *EmployeeCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *EmployeeCQ) SetVersionNo_Equal(value int64) *EmployeeCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}
func (q *EmployeeCQ) SetVersionNo_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueVersionNo(), "versionNo")
}
func (q *EmployeeCQ) SetVersionNo_NotEqual(value int64) *EmployeeCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *EmployeeCQ) SetVersionNo_GreaterThan(value int64) *EmployeeCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *EmployeeCQ) SetVersionNo_LessThan(value int64) *EmployeeCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *EmployeeCQ) SetVersionNo_GreaterEqual(value int64) *EmployeeCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *EmployeeCQ) SetVersionNo_LessEqual(value int64) *EmployeeCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *EmployeeCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *EmployeeCQ) AddOrderBy_VersionNo_Asc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *EmployeeCQ) AddOrderBy_VersionNo_Desc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *EmployeeCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}

func (q *EmployeeCQ) getCValueDelFlag() *df.ConditionValue {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	return q.DelFlag
}



func (q *EmployeeCQ) SetDelFlag_Equal(value int64) *EmployeeCQ {
	q.regDelFlag(df.CK_EQ_C, value)
	return q
}
func (q *EmployeeCQ) SetDelFlag_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDelFlag(), "delFlag")
}
func (q *EmployeeCQ) SetDelFlag_NotEqual(value int64) *EmployeeCQ {
	q.regDelFlag(df.CK_NE_C, value)
	return q
}

func (q *EmployeeCQ) SetDelFlag_GreaterThan(value int64) *EmployeeCQ {
	q.regDelFlag(df.CK_GT_C, value)
	return q
}

func (q *EmployeeCQ) SetDelFlag_LessThan(value int64) *EmployeeCQ {
	q.regDelFlag(df.CK_LT_C, value)
	return q
}

func (q *EmployeeCQ) SetDelFlag_GreaterEqual(value int64) *EmployeeCQ {
	q.regDelFlag(df.CK_GE_C, value)
	return q
}

func (q *EmployeeCQ) SetDelFlag_LessEqual(value int64) *EmployeeCQ {
	q.regDelFlag(df.CK_LE_C, value)
	return q
}
func (q *EmployeeCQ) SetDelFlag_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDelFlag(),"delFlag",rangeOfOption)
}	


func (q *EmployeeCQ) AddOrderBy_DelFlag_Asc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBA("delFlag")
	return q
}
func (q *EmployeeCQ) AddOrderBy_DelFlag_Desc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBD("delFlag")
	return q
}
func (q *EmployeeCQ) regDelFlag(key *df.ConditionKey, value interface{}) {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DelFlag, "delFlag")
}

func (q *EmployeeCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *EmployeeCQ) SetRegisterDatetime_Equal(value df.Timestamp) *EmployeeCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *EmployeeCQ) SetRegisterDatetime_GreaterThan(value df.Timestamp) *EmployeeCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *EmployeeCQ) SetRegisterDatetime_LessThan(value df.Timestamp) *EmployeeCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *EmployeeCQ) SetRegisterDatetime_GreaterEqual(value df.Timestamp) *EmployeeCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *EmployeeCQ) SetRegisterDatetime_LessEqual(value df.Timestamp) *EmployeeCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *EmployeeCQ) AddOrderBy_RegisterDatetime_Asc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *EmployeeCQ) AddOrderBy_RegisterDatetime_Desc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *EmployeeCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *EmployeeCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *EmployeeCQ) SetRegisterUser_Equal(value string) *EmployeeCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *EmployeeCQ) SetRegisterUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterUser(), "registerUser")
}
func (q *EmployeeCQ) SetRegisterUser_NotEqual(value string) *EmployeeCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetRegisterUser_GreaterThan(value string) *EmployeeCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetRegisterUser_LessThan(value string) *EmployeeCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetRegisterUser_GreaterEqual(value string) *EmployeeCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *EmployeeCQ) SetRegisterUser_LessEqual(value string) *EmployeeCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *EmployeeCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *EmployeeCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *EmployeeCQ) AddOrderBy_RegisterUser_Asc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *EmployeeCQ) AddOrderBy_RegisterUser_Desc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *EmployeeCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *EmployeeCQ) getCValueRegisterProcess() *df.ConditionValue {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	return q.RegisterProcess
}


func (q *EmployeeCQ) SetRegisterProcess_Equal(value string) *EmployeeCQ {
	q.regRegisterProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *EmployeeCQ) SetRegisterProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterProcess(), "registerProcess")
}
func (q *EmployeeCQ) SetRegisterProcess_NotEqual(value string) *EmployeeCQ {
	q.regRegisterProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetRegisterProcess_GreaterThan(value string) *EmployeeCQ {
	q.regRegisterProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetRegisterProcess_LessThan(value string) *EmployeeCQ {
	q.regRegisterProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetRegisterProcess_GreaterEqual(value string) *EmployeeCQ {
	q.regRegisterProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *EmployeeCQ) SetRegisterProcess_LessEqual(value string) *EmployeeCQ {
	q.regRegisterProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetRegisterProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}

func (q *EmployeeCQ) SetRegisterProcess_PrefixSearch(value string) error {
	return q.SetRegisterProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *EmployeeCQ) SetRegisterProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}



func (q *EmployeeCQ) AddOrderBy_RegisterProcess_Asc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBA("registerProcess")
	return q
}
func (q *EmployeeCQ) AddOrderBy_RegisterProcess_Desc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBD("registerProcess")
	return q
}
func (q *EmployeeCQ) regRegisterProcess(key *df.ConditionKey, value interface{}) {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterProcess, "registerProcess")
}

func (q *EmployeeCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *EmployeeCQ) SetUpdateDatetime_Equal(value df.Timestamp) *EmployeeCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *EmployeeCQ) SetUpdateDatetime_GreaterThan(value df.Timestamp) *EmployeeCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *EmployeeCQ) SetUpdateDatetime_LessThan(value df.Timestamp) *EmployeeCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *EmployeeCQ) SetUpdateDatetime_GreaterEqual(value df.Timestamp) *EmployeeCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *EmployeeCQ) SetUpdateDatetime_LessEqual(value df.Timestamp) *EmployeeCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *EmployeeCQ) AddOrderBy_UpdateDatetime_Asc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *EmployeeCQ) AddOrderBy_UpdateDatetime_Desc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *EmployeeCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *EmployeeCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *EmployeeCQ) SetUpdateUser_Equal(value string) *EmployeeCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *EmployeeCQ) SetUpdateUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateUser(), "updateUser")
}
func (q *EmployeeCQ) SetUpdateUser_NotEqual(value string) *EmployeeCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetUpdateUser_GreaterThan(value string) *EmployeeCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetUpdateUser_LessThan(value string) *EmployeeCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetUpdateUser_GreaterEqual(value string) *EmployeeCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *EmployeeCQ) SetUpdateUser_LessEqual(value string) *EmployeeCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *EmployeeCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *EmployeeCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *EmployeeCQ) AddOrderBy_UpdateUser_Asc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *EmployeeCQ) AddOrderBy_UpdateUser_Desc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *EmployeeCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *EmployeeCQ) getCValueUpdateProcess() *df.ConditionValue {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	return q.UpdateProcess
}


func (q *EmployeeCQ) SetUpdateProcess_Equal(value string) *EmployeeCQ {
	q.regUpdateProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *EmployeeCQ) SetUpdateProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateProcess(), "updateProcess")
}
func (q *EmployeeCQ) SetUpdateProcess_NotEqual(value string) *EmployeeCQ {
	q.regUpdateProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetUpdateProcess_GreaterThan(value string) *EmployeeCQ {
	q.regUpdateProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetUpdateProcess_LessThan(value string) *EmployeeCQ {
	q.regUpdateProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetUpdateProcess_GreaterEqual(value string) *EmployeeCQ {
	q.regUpdateProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *EmployeeCQ) SetUpdateProcess_LessEqual(value string) *EmployeeCQ {
	q.regUpdateProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *EmployeeCQ) SetUpdateProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}

func (q *EmployeeCQ) SetUpdateProcess_PrefixSearch(value string) error {
	return q.SetUpdateProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *EmployeeCQ) SetUpdateProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}



func (q *EmployeeCQ) AddOrderBy_UpdateProcess_Asc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBA("updateProcess")
	return q
}
func (q *EmployeeCQ) AddOrderBy_UpdateProcess_Desc() *EmployeeCQ {
	q.BaseConditionQuery.RegOBD("updateProcess")
	return q
}
func (q *EmployeeCQ) regUpdateProcess(key *df.ConditionKey, value interface{}) {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateProcess, "updateProcess")
}


func (q *EmployeeCQ) QueryUserTable() *UserTableCQ {
	if q.conditionQueryUserTable == nil {
		q.conditionQueryUserTable = q.xcreateQueryUserTable()
		q.xsetupOuterJoinUserTable()
	}
	return q.conditionQueryUserTable
}

func (q *EmployeeCQ) xcreateQueryUserTable() *UserTableCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("Employee", "UserTable")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreateUserTableCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "UserTable"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *EmployeeCQ) xsetupOuterJoinUserTable() {
	    cq := q.QueryUserTable()
        joinOnMap := make(map[string]string)
        joinOnMap["secId"]="id"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "UserTable");
}	
	
func CreateEmployeeCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *EmployeeCQ {
	cq := new(EmployeeCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "Employee"
	cq.BaseConditionQuery.ReferrerQuery = referrerQuery
	cq.BaseConditionQuery.SqlClause = sqlClause
	cq.BaseConditionQuery.AliasName = aliasName
	cq.BaseConditionQuery.NestLevel = nestlevel
	cq.BaseConditionQuery.DBMetaProvider = df.DBMetaProvider_I
	cq.BaseConditionQuery.CQ_PROPERTY = "Query"
	var cqi df.ConditionQuery = cq
	cq.BaseConditionQuery.ConditionQuery=&cqi
	return cq
}	
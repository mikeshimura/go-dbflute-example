package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type SessionCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	Id *df.ConditionValue
	Uuid *df.ConditionValue
	LoginId *df.ConditionValue
	CompId *df.ConditionValue
	Data *df.ConditionValue
	VersionNo *df.ConditionValue
	DelFlag *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterUser *df.ConditionValue
	RegisterProcess *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateUser *df.ConditionValue
	UpdateProcess *df.ConditionValue
    conditionQueryLogin *LoginCQ
}

func (q *SessionCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *SessionCQ) getCValueId() *df.ConditionValue {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	return q.Id
}



func (q *SessionCQ) SetId_Equal(value int64) *SessionCQ {
	q.regId(df.CK_EQ_C, value)
	return q
}
func (q *SessionCQ) SetId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueId(), "id")
}
func (q *SessionCQ) SetId_NotEqual(value int64) *SessionCQ {
	q.regId(df.CK_NE_C, value)
	return q
}

func (q *SessionCQ) SetId_GreaterThan(value int64) *SessionCQ {
	q.regId(df.CK_GT_C, value)
	return q
}

func (q *SessionCQ) SetId_LessThan(value int64) *SessionCQ {
	q.regId(df.CK_LT_C, value)
	return q
}

func (q *SessionCQ) SetId_GreaterEqual(value int64) *SessionCQ {
	q.regId(df.CK_GE_C, value)
	return q
}

func (q *SessionCQ) SetId_LessEqual(value int64) *SessionCQ {
	q.regId(df.CK_LE_C, value)
	return q
}
func (q *SessionCQ) SetId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueId(),"id",rangeOfOption)
}	


func (q *SessionCQ) SetId_IsNull() *SessionCQ {
	q.regId(df.CK_ISN_C, 0)
	return q
}
func (q *SessionCQ) SetId_IsNotNull() *SessionCQ {
	q.regId(df.CK_ISNN_C, 0)
	return q
}
func (q *SessionCQ) AddOrderBy_Id_Asc() *SessionCQ {
	q.BaseConditionQuery.RegOBA("id")
	return q
}
func (q *SessionCQ) AddOrderBy_Id_Desc() *SessionCQ {
	q.BaseConditionQuery.RegOBD("id")
	return q
}
func (q *SessionCQ) regId(key *df.ConditionKey, value interface{}) {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Id, "id")
}

func (q *SessionCQ) getCValueUuid() *df.ConditionValue {
	if q.Uuid == nil {
		q.Uuid = new(df.ConditionValue)
	}
	return q.Uuid
}


func (q *SessionCQ) SetUuid_Equal(value string) *SessionCQ {
	q.regUuid(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SessionCQ) SetUuid_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUuid(), "uuid")
}
func (q *SessionCQ) SetUuid_NotEqual(value string) *SessionCQ {
	q.regUuid(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetUuid_GreaterThan(value string) *SessionCQ {
	q.regUuid(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetUuid_LessThan(value string) *SessionCQ {
	q.regUuid(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetUuid_GreaterEqual(value string) *SessionCQ {
	q.regUuid(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SessionCQ) SetUuid_LessEqual(value string) *SessionCQ {
	q.regUuid(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetUuid_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUuid(), "uuid", option)
}

func (q *SessionCQ) SetUuid_PrefixSearch(value string) error {
	return q.SetUuid_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SessionCQ) SetUuid_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUuid(), "uuid", option)
}



func (q *SessionCQ) AddOrderBy_Uuid_Asc() *SessionCQ {
	q.BaseConditionQuery.RegOBA("uuid")
	return q
}
func (q *SessionCQ) AddOrderBy_Uuid_Desc() *SessionCQ {
	q.BaseConditionQuery.RegOBD("uuid")
	return q
}
func (q *SessionCQ) regUuid(key *df.ConditionKey, value interface{}) {
	if q.Uuid == nil {
		q.Uuid = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Uuid, "uuid")
}

func (q *SessionCQ) getCValueLoginId() *df.ConditionValue {
	if q.LoginId == nil {
		q.LoginId = new(df.ConditionValue)
	}
	return q.LoginId
}



func (q *SessionCQ) SetLoginId_Equal(value int64) *SessionCQ {
	q.regLoginId(df.CK_EQ_C, value)
	return q
}
func (q *SessionCQ) SetLoginId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueLoginId(), "loginId")
}
func (q *SessionCQ) SetLoginId_NotEqual(value int64) *SessionCQ {
	q.regLoginId(df.CK_NE_C, value)
	return q
}

func (q *SessionCQ) SetLoginId_GreaterThan(value int64) *SessionCQ {
	q.regLoginId(df.CK_GT_C, value)
	return q
}

func (q *SessionCQ) SetLoginId_LessThan(value int64) *SessionCQ {
	q.regLoginId(df.CK_LT_C, value)
	return q
}

func (q *SessionCQ) SetLoginId_GreaterEqual(value int64) *SessionCQ {
	q.regLoginId(df.CK_GE_C, value)
	return q
}

func (q *SessionCQ) SetLoginId_LessEqual(value int64) *SessionCQ {
	q.regLoginId(df.CK_LE_C, value)
	return q
}
func (q *SessionCQ) SetLoginId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueLoginId(),"loginId",rangeOfOption)
}	


func (q *SessionCQ) SetLoginId_IsNull() *SessionCQ {
	q.regLoginId(df.CK_ISN_C, 0)
	return q
}
func (q *SessionCQ) SetLoginId_IsNotNull() *SessionCQ {
	q.regLoginId(df.CK_ISNN_C, 0)
	return q
}
func (q *SessionCQ) AddOrderBy_LoginId_Asc() *SessionCQ {
	q.BaseConditionQuery.RegOBA("loginId")
	return q
}
func (q *SessionCQ) AddOrderBy_LoginId_Desc() *SessionCQ {
	q.BaseConditionQuery.RegOBD("loginId")
	return q
}
func (q *SessionCQ) regLoginId(key *df.ConditionKey, value interface{}) {
	if q.LoginId == nil {
		q.LoginId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.LoginId, "loginId")
}

func (q *SessionCQ) getCValueCompId() *df.ConditionValue {
	if q.CompId == nil {
		q.CompId = new(df.ConditionValue)
	}
	return q.CompId
}



func (q *SessionCQ) SetCompId_Equal(value int64) *SessionCQ {
	q.regCompId(df.CK_EQ_C, value)
	return q
}
func (q *SessionCQ) SetCompId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueCompId(), "compId")
}
func (q *SessionCQ) SetCompId_NotEqual(value int64) *SessionCQ {
	q.regCompId(df.CK_NE_C, value)
	return q
}

func (q *SessionCQ) SetCompId_GreaterThan(value int64) *SessionCQ {
	q.regCompId(df.CK_GT_C, value)
	return q
}

func (q *SessionCQ) SetCompId_LessThan(value int64) *SessionCQ {
	q.regCompId(df.CK_LT_C, value)
	return q
}

func (q *SessionCQ) SetCompId_GreaterEqual(value int64) *SessionCQ {
	q.regCompId(df.CK_GE_C, value)
	return q
}

func (q *SessionCQ) SetCompId_LessEqual(value int64) *SessionCQ {
	q.regCompId(df.CK_LE_C, value)
	return q
}
func (q *SessionCQ) SetCompId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueCompId(),"compId",rangeOfOption)
}	


func (q *SessionCQ) SetCompId_IsNull() *SessionCQ {
	q.regCompId(df.CK_ISN_C, 0)
	return q
}
func (q *SessionCQ) SetCompId_IsNotNull() *SessionCQ {
	q.regCompId(df.CK_ISNN_C, 0)
	return q
}
func (q *SessionCQ) AddOrderBy_CompId_Asc() *SessionCQ {
	q.BaseConditionQuery.RegOBA("compId")
	return q
}
func (q *SessionCQ) AddOrderBy_CompId_Desc() *SessionCQ {
	q.BaseConditionQuery.RegOBD("compId")
	return q
}
func (q *SessionCQ) regCompId(key *df.ConditionKey, value interface{}) {
	if q.CompId == nil {
		q.CompId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.CompId, "compId")
}

func (q *SessionCQ) getCValueData() *df.ConditionValue {
	if q.Data == nil {
		q.Data = new(df.ConditionValue)
	}
	return q.Data
}


func (q *SessionCQ) SetData_Equal(value string) *SessionCQ {
	q.regData(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SessionCQ) SetData_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueData(), "data")
}
func (q *SessionCQ) SetData_NotEqual(value string) *SessionCQ {
	q.regData(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetData_GreaterThan(value string) *SessionCQ {
	q.regData(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetData_LessThan(value string) *SessionCQ {
	q.regData(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetData_GreaterEqual(value string) *SessionCQ {
	q.regData(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SessionCQ) SetData_LessEqual(value string) *SessionCQ {
	q.regData(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetData_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueData(), "data", option)
}

func (q *SessionCQ) SetData_PrefixSearch(value string) error {
	return q.SetData_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SessionCQ) SetData_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueData(), "data", option)
}



func (q *SessionCQ) SetData_IsNull() *SessionCQ {
	q.regData(df.CK_ISN_C, 0)
	return q
}
func (q *SessionCQ) SetData_IsNullOrEmpty() *SessionCQ {
	q.regData(df.CK_ISNOE_C, 0)
	return q
}
func (q *SessionCQ) SetData_IsNotNull() *SessionCQ {
	q.regData(df.CK_ISNN_C, 0)
	return q
}
func (q *SessionCQ) AddOrderBy_Data_Asc() *SessionCQ {
	q.BaseConditionQuery.RegOBA("data")
	return q
}
func (q *SessionCQ) AddOrderBy_Data_Desc() *SessionCQ {
	q.BaseConditionQuery.RegOBD("data")
	return q
}
func (q *SessionCQ) regData(key *df.ConditionKey, value interface{}) {
	if q.Data == nil {
		q.Data = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Data, "data")
}

func (q *SessionCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *SessionCQ) SetVersionNo_Equal(value int64) *SessionCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}
func (q *SessionCQ) SetVersionNo_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueVersionNo(), "versionNo")
}
func (q *SessionCQ) SetVersionNo_NotEqual(value int64) *SessionCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *SessionCQ) SetVersionNo_GreaterThan(value int64) *SessionCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *SessionCQ) SetVersionNo_LessThan(value int64) *SessionCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *SessionCQ) SetVersionNo_GreaterEqual(value int64) *SessionCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *SessionCQ) SetVersionNo_LessEqual(value int64) *SessionCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *SessionCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *SessionCQ) AddOrderBy_VersionNo_Asc() *SessionCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *SessionCQ) AddOrderBy_VersionNo_Desc() *SessionCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *SessionCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}

func (q *SessionCQ) getCValueDelFlag() *df.ConditionValue {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	return q.DelFlag
}



func (q *SessionCQ) SetDelFlag_Equal(value int64) *SessionCQ {
	q.regDelFlag(df.CK_EQ_C, value)
	return q
}
func (q *SessionCQ) SetDelFlag_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDelFlag(), "delFlag")
}
func (q *SessionCQ) SetDelFlag_NotEqual(value int64) *SessionCQ {
	q.regDelFlag(df.CK_NE_C, value)
	return q
}

func (q *SessionCQ) SetDelFlag_GreaterThan(value int64) *SessionCQ {
	q.regDelFlag(df.CK_GT_C, value)
	return q
}

func (q *SessionCQ) SetDelFlag_LessThan(value int64) *SessionCQ {
	q.regDelFlag(df.CK_LT_C, value)
	return q
}

func (q *SessionCQ) SetDelFlag_GreaterEqual(value int64) *SessionCQ {
	q.regDelFlag(df.CK_GE_C, value)
	return q
}

func (q *SessionCQ) SetDelFlag_LessEqual(value int64) *SessionCQ {
	q.regDelFlag(df.CK_LE_C, value)
	return q
}
func (q *SessionCQ) SetDelFlag_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDelFlag(),"delFlag",rangeOfOption)
}	


func (q *SessionCQ) AddOrderBy_DelFlag_Asc() *SessionCQ {
	q.BaseConditionQuery.RegOBA("delFlag")
	return q
}
func (q *SessionCQ) AddOrderBy_DelFlag_Desc() *SessionCQ {
	q.BaseConditionQuery.RegOBD("delFlag")
	return q
}
func (q *SessionCQ) regDelFlag(key *df.ConditionKey, value interface{}) {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DelFlag, "delFlag")
}

func (q *SessionCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *SessionCQ) SetRegisterDatetime_Equal(value df.Timestamp) *SessionCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *SessionCQ) SetRegisterDatetime_GreaterThan(value df.Timestamp) *SessionCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *SessionCQ) SetRegisterDatetime_LessThan(value df.Timestamp) *SessionCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *SessionCQ) SetRegisterDatetime_GreaterEqual(value df.Timestamp) *SessionCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *SessionCQ) SetRegisterDatetime_LessEqual(value df.Timestamp) *SessionCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *SessionCQ) AddOrderBy_RegisterDatetime_Asc() *SessionCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *SessionCQ) AddOrderBy_RegisterDatetime_Desc() *SessionCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *SessionCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *SessionCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *SessionCQ) SetRegisterUser_Equal(value string) *SessionCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SessionCQ) SetRegisterUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterUser(), "registerUser")
}
func (q *SessionCQ) SetRegisterUser_NotEqual(value string) *SessionCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetRegisterUser_GreaterThan(value string) *SessionCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetRegisterUser_LessThan(value string) *SessionCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetRegisterUser_GreaterEqual(value string) *SessionCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SessionCQ) SetRegisterUser_LessEqual(value string) *SessionCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *SessionCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SessionCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *SessionCQ) AddOrderBy_RegisterUser_Asc() *SessionCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *SessionCQ) AddOrderBy_RegisterUser_Desc() *SessionCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *SessionCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *SessionCQ) getCValueRegisterProcess() *df.ConditionValue {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	return q.RegisterProcess
}


func (q *SessionCQ) SetRegisterProcess_Equal(value string) *SessionCQ {
	q.regRegisterProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SessionCQ) SetRegisterProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterProcess(), "registerProcess")
}
func (q *SessionCQ) SetRegisterProcess_NotEqual(value string) *SessionCQ {
	q.regRegisterProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetRegisterProcess_GreaterThan(value string) *SessionCQ {
	q.regRegisterProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetRegisterProcess_LessThan(value string) *SessionCQ {
	q.regRegisterProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetRegisterProcess_GreaterEqual(value string) *SessionCQ {
	q.regRegisterProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SessionCQ) SetRegisterProcess_LessEqual(value string) *SessionCQ {
	q.regRegisterProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetRegisterProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}

func (q *SessionCQ) SetRegisterProcess_PrefixSearch(value string) error {
	return q.SetRegisterProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SessionCQ) SetRegisterProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}



func (q *SessionCQ) AddOrderBy_RegisterProcess_Asc() *SessionCQ {
	q.BaseConditionQuery.RegOBA("registerProcess")
	return q
}
func (q *SessionCQ) AddOrderBy_RegisterProcess_Desc() *SessionCQ {
	q.BaseConditionQuery.RegOBD("registerProcess")
	return q
}
func (q *SessionCQ) regRegisterProcess(key *df.ConditionKey, value interface{}) {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterProcess, "registerProcess")
}

func (q *SessionCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *SessionCQ) SetUpdateDatetime_Equal(value df.Timestamp) *SessionCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *SessionCQ) SetUpdateDatetime_GreaterThan(value df.Timestamp) *SessionCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *SessionCQ) SetUpdateDatetime_LessThan(value df.Timestamp) *SessionCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *SessionCQ) SetUpdateDatetime_GreaterEqual(value df.Timestamp) *SessionCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *SessionCQ) SetUpdateDatetime_LessEqual(value df.Timestamp) *SessionCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *SessionCQ) AddOrderBy_UpdateDatetime_Asc() *SessionCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *SessionCQ) AddOrderBy_UpdateDatetime_Desc() *SessionCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *SessionCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *SessionCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *SessionCQ) SetUpdateUser_Equal(value string) *SessionCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SessionCQ) SetUpdateUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateUser(), "updateUser")
}
func (q *SessionCQ) SetUpdateUser_NotEqual(value string) *SessionCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetUpdateUser_GreaterThan(value string) *SessionCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetUpdateUser_LessThan(value string) *SessionCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetUpdateUser_GreaterEqual(value string) *SessionCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SessionCQ) SetUpdateUser_LessEqual(value string) *SessionCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *SessionCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SessionCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *SessionCQ) AddOrderBy_UpdateUser_Asc() *SessionCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *SessionCQ) AddOrderBy_UpdateUser_Desc() *SessionCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *SessionCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *SessionCQ) getCValueUpdateProcess() *df.ConditionValue {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	return q.UpdateProcess
}


func (q *SessionCQ) SetUpdateProcess_Equal(value string) *SessionCQ {
	q.regUpdateProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SessionCQ) SetUpdateProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateProcess(), "updateProcess")
}
func (q *SessionCQ) SetUpdateProcess_NotEqual(value string) *SessionCQ {
	q.regUpdateProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetUpdateProcess_GreaterThan(value string) *SessionCQ {
	q.regUpdateProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetUpdateProcess_LessThan(value string) *SessionCQ {
	q.regUpdateProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetUpdateProcess_GreaterEqual(value string) *SessionCQ {
	q.regUpdateProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SessionCQ) SetUpdateProcess_LessEqual(value string) *SessionCQ {
	q.regUpdateProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SessionCQ) SetUpdateProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}

func (q *SessionCQ) SetUpdateProcess_PrefixSearch(value string) error {
	return q.SetUpdateProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SessionCQ) SetUpdateProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}



func (q *SessionCQ) AddOrderBy_UpdateProcess_Asc() *SessionCQ {
	q.BaseConditionQuery.RegOBA("updateProcess")
	return q
}
func (q *SessionCQ) AddOrderBy_UpdateProcess_Desc() *SessionCQ {
	q.BaseConditionQuery.RegOBD("updateProcess")
	return q
}
func (q *SessionCQ) regUpdateProcess(key *df.ConditionKey, value interface{}) {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateProcess, "updateProcess")
}


func (q *SessionCQ) QueryLogin() *LoginCQ {
	if q.conditionQueryLogin == nil {
		q.conditionQueryLogin = q.xcreateQueryLogin()
		q.xsetupOuterJoinLogin()
	}
	return q.conditionQueryLogin
}

func (q *SessionCQ) xcreateQueryLogin() *LoginCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("Session", "Login")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreateLoginCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "Login"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *SessionCQ) xsetupOuterJoinLogin() {
	    cq := q.QueryLogin()
        joinOnMap := make(map[string]string)
        joinOnMap["loginId"]="id"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "Login");
}	
	
func CreateSessionCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *SessionCQ {
	cq := new(SessionCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "Session"
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

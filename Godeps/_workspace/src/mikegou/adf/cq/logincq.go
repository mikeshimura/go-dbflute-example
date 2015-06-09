package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type LoginCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	Id *df.ConditionValue
	LoginId *df.ConditionValue
	Name *df.ConditionValue
	Password *df.ConditionValue
	VersionNo *df.ConditionValue
	DelFlag *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterUser *df.ConditionValue
	RegisterProcess *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateUser *df.ConditionValue
	UpdateProcess *df.ConditionValue
}

func (q *LoginCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *LoginCQ) getCValueId() *df.ConditionValue {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	return q.Id
}



func (q *LoginCQ) SetId_Equal(value int64) *LoginCQ {
	q.regId(df.CK_EQ_C, value)
	return q
}
func (q *LoginCQ) SetId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueId(), "id")
}
func (q *LoginCQ) SetId_NotEqual(value int64) *LoginCQ {
	q.regId(df.CK_NE_C, value)
	return q
}

func (q *LoginCQ) SetId_GreaterThan(value int64) *LoginCQ {
	q.regId(df.CK_GT_C, value)
	return q
}

func (q *LoginCQ) SetId_LessThan(value int64) *LoginCQ {
	q.regId(df.CK_LT_C, value)
	return q
}

func (q *LoginCQ) SetId_GreaterEqual(value int64) *LoginCQ {
	q.regId(df.CK_GE_C, value)
	return q
}

func (q *LoginCQ) SetId_LessEqual(value int64) *LoginCQ {
	q.regId(df.CK_LE_C, value)
	return q
}
func (q *LoginCQ) SetId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueId(),"id",rangeOfOption)
}	


func (q *LoginCQ) SetId_IsNull() *LoginCQ {
	q.regId(df.CK_ISN_C, 0)
	return q
}
func (q *LoginCQ) SetId_IsNotNull() *LoginCQ {
	q.regId(df.CK_ISNN_C, 0)
	return q
}
func (q *LoginCQ) AddOrderBy_Id_Asc() *LoginCQ {
	q.BaseConditionQuery.RegOBA("id")
	return q
}
func (q *LoginCQ) AddOrderBy_Id_Desc() *LoginCQ {
	q.BaseConditionQuery.RegOBD("id")
	return q
}
func (q *LoginCQ) regId(key *df.ConditionKey, value interface{}) {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Id, "id")
}

func (q *LoginCQ) getCValueLoginId() *df.ConditionValue {
	if q.LoginId == nil {
		q.LoginId = new(df.ConditionValue)
	}
	return q.LoginId
}


func (q *LoginCQ) SetLoginId_Equal(value string) *LoginCQ {
	q.regLoginId(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *LoginCQ) SetLoginId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueLoginId(), "loginId")
}
func (q *LoginCQ) SetLoginId_NotEqual(value string) *LoginCQ {
	q.regLoginId(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetLoginId_GreaterThan(value string) *LoginCQ {
	q.regLoginId(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetLoginId_LessThan(value string) *LoginCQ {
	q.regLoginId(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetLoginId_GreaterEqual(value string) *LoginCQ {
	q.regLoginId(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *LoginCQ) SetLoginId_LessEqual(value string) *LoginCQ {
	q.regLoginId(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetLoginId_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueLoginId(), "loginId", option)
}

func (q *LoginCQ) SetLoginId_PrefixSearch(value string) error {
	return q.SetLoginId_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *LoginCQ) SetLoginId_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueLoginId(), "loginId", option)
}



func (q *LoginCQ) AddOrderBy_LoginId_Asc() *LoginCQ {
	q.BaseConditionQuery.RegOBA("loginId")
	return q
}
func (q *LoginCQ) AddOrderBy_LoginId_Desc() *LoginCQ {
	q.BaseConditionQuery.RegOBD("loginId")
	return q
}
func (q *LoginCQ) regLoginId(key *df.ConditionKey, value interface{}) {
	if q.LoginId == nil {
		q.LoginId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.LoginId, "loginId")
}

func (q *LoginCQ) getCValueName() *df.ConditionValue {
	if q.Name == nil {
		q.Name = new(df.ConditionValue)
	}
	return q.Name
}


func (q *LoginCQ) SetName_Equal(value string) *LoginCQ {
	q.regName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *LoginCQ) SetName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueName(), "name")
}
func (q *LoginCQ) SetName_NotEqual(value string) *LoginCQ {
	q.regName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetName_GreaterThan(value string) *LoginCQ {
	q.regName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetName_LessThan(value string) *LoginCQ {
	q.regName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetName_GreaterEqual(value string) *LoginCQ {
	q.regName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *LoginCQ) SetName_LessEqual(value string) *LoginCQ {
	q.regName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueName(), "name", option)
}

func (q *LoginCQ) SetName_PrefixSearch(value string) error {
	return q.SetName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *LoginCQ) SetName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueName(), "name", option)
}



func (q *LoginCQ) AddOrderBy_Name_Asc() *LoginCQ {
	q.BaseConditionQuery.RegOBA("name")
	return q
}
func (q *LoginCQ) AddOrderBy_Name_Desc() *LoginCQ {
	q.BaseConditionQuery.RegOBD("name")
	return q
}
func (q *LoginCQ) regName(key *df.ConditionKey, value interface{}) {
	if q.Name == nil {
		q.Name = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Name, "name")
}

func (q *LoginCQ) getCValuePassword() *df.ConditionValue {
	if q.Password == nil {
		q.Password = new(df.ConditionValue)
	}
	return q.Password
}


func (q *LoginCQ) SetPassword_Equal(value string) *LoginCQ {
	q.regPassword(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *LoginCQ) SetPassword_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePassword(), "password")
}
func (q *LoginCQ) SetPassword_NotEqual(value string) *LoginCQ {
	q.regPassword(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetPassword_GreaterThan(value string) *LoginCQ {
	q.regPassword(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetPassword_LessThan(value string) *LoginCQ {
	q.regPassword(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetPassword_GreaterEqual(value string) *LoginCQ {
	q.regPassword(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *LoginCQ) SetPassword_LessEqual(value string) *LoginCQ {
	q.regPassword(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetPassword_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValuePassword(), "password", option)
}

func (q *LoginCQ) SetPassword_PrefixSearch(value string) error {
	return q.SetPassword_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *LoginCQ) SetPassword_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValuePassword(), "password", option)
}



func (q *LoginCQ) AddOrderBy_Password_Asc() *LoginCQ {
	q.BaseConditionQuery.RegOBA("password")
	return q
}
func (q *LoginCQ) AddOrderBy_Password_Desc() *LoginCQ {
	q.BaseConditionQuery.RegOBD("password")
	return q
}
func (q *LoginCQ) regPassword(key *df.ConditionKey, value interface{}) {
	if q.Password == nil {
		q.Password = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Password, "password")
}

func (q *LoginCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *LoginCQ) SetVersionNo_Equal(value int64) *LoginCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}
func (q *LoginCQ) SetVersionNo_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueVersionNo(), "versionNo")
}
func (q *LoginCQ) SetVersionNo_NotEqual(value int64) *LoginCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *LoginCQ) SetVersionNo_GreaterThan(value int64) *LoginCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *LoginCQ) SetVersionNo_LessThan(value int64) *LoginCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *LoginCQ) SetVersionNo_GreaterEqual(value int64) *LoginCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *LoginCQ) SetVersionNo_LessEqual(value int64) *LoginCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *LoginCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *LoginCQ) AddOrderBy_VersionNo_Asc() *LoginCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *LoginCQ) AddOrderBy_VersionNo_Desc() *LoginCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *LoginCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}

func (q *LoginCQ) getCValueDelFlag() *df.ConditionValue {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	return q.DelFlag
}



func (q *LoginCQ) SetDelFlag_Equal(value int64) *LoginCQ {
	q.regDelFlag(df.CK_EQ_C, value)
	return q
}
func (q *LoginCQ) SetDelFlag_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDelFlag(), "delFlag")
}
func (q *LoginCQ) SetDelFlag_NotEqual(value int64) *LoginCQ {
	q.regDelFlag(df.CK_NE_C, value)
	return q
}

func (q *LoginCQ) SetDelFlag_GreaterThan(value int64) *LoginCQ {
	q.regDelFlag(df.CK_GT_C, value)
	return q
}

func (q *LoginCQ) SetDelFlag_LessThan(value int64) *LoginCQ {
	q.regDelFlag(df.CK_LT_C, value)
	return q
}

func (q *LoginCQ) SetDelFlag_GreaterEqual(value int64) *LoginCQ {
	q.regDelFlag(df.CK_GE_C, value)
	return q
}

func (q *LoginCQ) SetDelFlag_LessEqual(value int64) *LoginCQ {
	q.regDelFlag(df.CK_LE_C, value)
	return q
}
func (q *LoginCQ) SetDelFlag_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDelFlag(),"delFlag",rangeOfOption)
}	


func (q *LoginCQ) AddOrderBy_DelFlag_Asc() *LoginCQ {
	q.BaseConditionQuery.RegOBA("delFlag")
	return q
}
func (q *LoginCQ) AddOrderBy_DelFlag_Desc() *LoginCQ {
	q.BaseConditionQuery.RegOBD("delFlag")
	return q
}
func (q *LoginCQ) regDelFlag(key *df.ConditionKey, value interface{}) {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DelFlag, "delFlag")
}

func (q *LoginCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *LoginCQ) SetRegisterDatetime_Equal(value df.Timestamp) *LoginCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *LoginCQ) SetRegisterDatetime_GreaterThan(value df.Timestamp) *LoginCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *LoginCQ) SetRegisterDatetime_LessThan(value df.Timestamp) *LoginCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *LoginCQ) SetRegisterDatetime_GreaterEqual(value df.Timestamp) *LoginCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *LoginCQ) SetRegisterDatetime_LessEqual(value df.Timestamp) *LoginCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *LoginCQ) AddOrderBy_RegisterDatetime_Asc() *LoginCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *LoginCQ) AddOrderBy_RegisterDatetime_Desc() *LoginCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *LoginCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *LoginCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *LoginCQ) SetRegisterUser_Equal(value string) *LoginCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *LoginCQ) SetRegisterUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterUser(), "registerUser")
}
func (q *LoginCQ) SetRegisterUser_NotEqual(value string) *LoginCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetRegisterUser_GreaterThan(value string) *LoginCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetRegisterUser_LessThan(value string) *LoginCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetRegisterUser_GreaterEqual(value string) *LoginCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *LoginCQ) SetRegisterUser_LessEqual(value string) *LoginCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *LoginCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *LoginCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *LoginCQ) AddOrderBy_RegisterUser_Asc() *LoginCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *LoginCQ) AddOrderBy_RegisterUser_Desc() *LoginCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *LoginCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *LoginCQ) getCValueRegisterProcess() *df.ConditionValue {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	return q.RegisterProcess
}


func (q *LoginCQ) SetRegisterProcess_Equal(value string) *LoginCQ {
	q.regRegisterProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *LoginCQ) SetRegisterProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterProcess(), "registerProcess")
}
func (q *LoginCQ) SetRegisterProcess_NotEqual(value string) *LoginCQ {
	q.regRegisterProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetRegisterProcess_GreaterThan(value string) *LoginCQ {
	q.regRegisterProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetRegisterProcess_LessThan(value string) *LoginCQ {
	q.regRegisterProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetRegisterProcess_GreaterEqual(value string) *LoginCQ {
	q.regRegisterProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *LoginCQ) SetRegisterProcess_LessEqual(value string) *LoginCQ {
	q.regRegisterProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetRegisterProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}

func (q *LoginCQ) SetRegisterProcess_PrefixSearch(value string) error {
	return q.SetRegisterProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *LoginCQ) SetRegisterProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}



func (q *LoginCQ) AddOrderBy_RegisterProcess_Asc() *LoginCQ {
	q.BaseConditionQuery.RegOBA("registerProcess")
	return q
}
func (q *LoginCQ) AddOrderBy_RegisterProcess_Desc() *LoginCQ {
	q.BaseConditionQuery.RegOBD("registerProcess")
	return q
}
func (q *LoginCQ) regRegisterProcess(key *df.ConditionKey, value interface{}) {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterProcess, "registerProcess")
}

func (q *LoginCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *LoginCQ) SetUpdateDatetime_Equal(value df.Timestamp) *LoginCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *LoginCQ) SetUpdateDatetime_GreaterThan(value df.Timestamp) *LoginCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *LoginCQ) SetUpdateDatetime_LessThan(value df.Timestamp) *LoginCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *LoginCQ) SetUpdateDatetime_GreaterEqual(value df.Timestamp) *LoginCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *LoginCQ) SetUpdateDatetime_LessEqual(value df.Timestamp) *LoginCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *LoginCQ) AddOrderBy_UpdateDatetime_Asc() *LoginCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *LoginCQ) AddOrderBy_UpdateDatetime_Desc() *LoginCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *LoginCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *LoginCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *LoginCQ) SetUpdateUser_Equal(value string) *LoginCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *LoginCQ) SetUpdateUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateUser(), "updateUser")
}
func (q *LoginCQ) SetUpdateUser_NotEqual(value string) *LoginCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetUpdateUser_GreaterThan(value string) *LoginCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetUpdateUser_LessThan(value string) *LoginCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetUpdateUser_GreaterEqual(value string) *LoginCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *LoginCQ) SetUpdateUser_LessEqual(value string) *LoginCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *LoginCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *LoginCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *LoginCQ) AddOrderBy_UpdateUser_Asc() *LoginCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *LoginCQ) AddOrderBy_UpdateUser_Desc() *LoginCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *LoginCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *LoginCQ) getCValueUpdateProcess() *df.ConditionValue {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	return q.UpdateProcess
}


func (q *LoginCQ) SetUpdateProcess_Equal(value string) *LoginCQ {
	q.regUpdateProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *LoginCQ) SetUpdateProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateProcess(), "updateProcess")
}
func (q *LoginCQ) SetUpdateProcess_NotEqual(value string) *LoginCQ {
	q.regUpdateProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetUpdateProcess_GreaterThan(value string) *LoginCQ {
	q.regUpdateProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetUpdateProcess_LessThan(value string) *LoginCQ {
	q.regUpdateProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetUpdateProcess_GreaterEqual(value string) *LoginCQ {
	q.regUpdateProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *LoginCQ) SetUpdateProcess_LessEqual(value string) *LoginCQ {
	q.regUpdateProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *LoginCQ) SetUpdateProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}

func (q *LoginCQ) SetUpdateProcess_PrefixSearch(value string) error {
	return q.SetUpdateProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *LoginCQ) SetUpdateProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}



func (q *LoginCQ) AddOrderBy_UpdateProcess_Asc() *LoginCQ {
	q.BaseConditionQuery.RegOBA("updateProcess")
	return q
}
func (q *LoginCQ) AddOrderBy_UpdateProcess_Desc() *LoginCQ {
	q.BaseConditionQuery.RegOBD("updateProcess")
	return q
}
func (q *LoginCQ) regUpdateProcess(key *df.ConditionKey, value interface{}) {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateProcess, "updateProcess")
}


func CreateLoginCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *LoginCQ {
	cq := new(LoginCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "Login"
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

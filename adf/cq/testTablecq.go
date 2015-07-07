package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type TestTableCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	Id *df.ConditionValue
	TestId *df.ConditionValue
	TestDate *df.ConditionValue
	TestTimestamp *df.ConditionValue
	TestNbr *df.ConditionValue
	VersionNo *df.ConditionValue
	DelFlag *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterUser *df.ConditionValue
	RegisterProcess *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateUser *df.ConditionValue
	UpdateProcess *df.ConditionValue
}

func (q *TestTableCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *TestTableCQ) getCValueId() *df.ConditionValue {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	return q.Id
}



func (q *TestTableCQ) SetId_Equal(value int64) *TestTableCQ {
	q.regId(df.CK_EQ_C, value)
	return q
}
func (q *TestTableCQ) SetId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueId(), "id")
}
func (q *TestTableCQ) SetId_NotEqual(value int64) *TestTableCQ {
	q.regId(df.CK_NE_C, value)
	return q
}

func (q *TestTableCQ) SetId_GreaterThan(value int64) *TestTableCQ {
	q.regId(df.CK_GT_C, value)
	return q
}

func (q *TestTableCQ) SetId_LessThan(value int64) *TestTableCQ {
	q.regId(df.CK_LT_C, value)
	return q
}

func (q *TestTableCQ) SetId_GreaterEqual(value int64) *TestTableCQ {
	q.regId(df.CK_GE_C, value)
	return q
}

func (q *TestTableCQ) SetId_LessEqual(value int64) *TestTableCQ {
	q.regId(df.CK_LE_C, value)
	return q
}
func (q *TestTableCQ) SetId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueId(),"id",rangeOfOption)
}	


func (q *TestTableCQ) SetId_IsNull() *TestTableCQ {
	q.regId(df.CK_ISN_C, 0)
	return q
}
func (q *TestTableCQ) SetId_IsNotNull() *TestTableCQ {
	q.regId(df.CK_ISNN_C, 0)
	return q
}
func (q *TestTableCQ) AddOrderBy_Id_Asc() *TestTableCQ {
	q.BaseConditionQuery.RegOBA("id")
	return q
}
func (q *TestTableCQ) AddOrderBy_Id_Desc() *TestTableCQ {
	q.BaseConditionQuery.RegOBD("id")
	return q
}
func (q *TestTableCQ) regId(key *df.ConditionKey, value interface{}) {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Id, "id")
}

func (q *TestTableCQ) getCValueTestId() *df.ConditionValue {
	if q.TestId == nil {
		q.TestId = new(df.ConditionValue)
	}
	return q.TestId
}


func (q *TestTableCQ) SetTestId_Equal(value string) *TestTableCQ {
	q.regTestId(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *TestTableCQ) SetTestId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueTestId(), "testId")
}
func (q *TestTableCQ) SetTestId_NotEqual(value string) *TestTableCQ {
	q.regTestId(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetTestId_GreaterThan(value string) *TestTableCQ {
	q.regTestId(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetTestId_LessThan(value string) *TestTableCQ {
	q.regTestId(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetTestId_GreaterEqual(value string) *TestTableCQ {
	q.regTestId(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *TestTableCQ) SetTestId_LessEqual(value string) *TestTableCQ {
	q.regTestId(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetTestId_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueTestId(), "testId", option)
}

func (q *TestTableCQ) SetTestId_PrefixSearch(value string) error {
	return q.SetTestId_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *TestTableCQ) SetTestId_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueTestId(), "testId", option)
}



func (q *TestTableCQ) AddOrderBy_TestId_Asc() *TestTableCQ {
	q.BaseConditionQuery.RegOBA("testId")
	return q
}
func (q *TestTableCQ) AddOrderBy_TestId_Desc() *TestTableCQ {
	q.BaseConditionQuery.RegOBD("testId")
	return q
}
func (q *TestTableCQ) regTestId(key *df.ConditionKey, value interface{}) {
	if q.TestId == nil {
		q.TestId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TestId, "testId")
}

func (q *TestTableCQ) getCValueTestDate() *df.ConditionValue {
	if q.TestDate == nil {
		q.TestDate = new(df.ConditionValue)
	}
	return q.TestDate
}




func (q *TestTableCQ) SetTestDate_Equal(value df.Date) *TestTableCQ {
	q.regTestDate(df.CK_EQ_C, value)
	return q
}


func (q *TestTableCQ) SetTestDate_GreaterThan(value df.Date) *TestTableCQ {
	q.regTestDate(df.CK_GT_C, value)
	return q
}

func (q *TestTableCQ) SetTestDate_LessThan(value df.Date) *TestTableCQ {
	q.regTestDate(df.CK_LT_C, value)
	return q
}

func (q *TestTableCQ) SetTestDate_GreaterEqual(value df.Date) *TestTableCQ {
	q.regTestDate(df.CK_GE_C, value)
	return q
}

func (q *TestTableCQ) SetTestDate_LessEqual(value df.Date) *TestTableCQ {
	q.regTestDate(df.CK_LE_C, value)
	return q
}

func (q *TestTableCQ) SetTestDate_IsNull() *TestTableCQ {
	q.regTestDate(df.CK_ISN_C, 0)
	return q
}
func (q *TestTableCQ) SetTestDate_IsNotNull() *TestTableCQ {
	q.regTestDate(df.CK_ISNN_C, 0)
	return q
}
func (q *TestTableCQ) AddOrderBy_TestDate_Asc() *TestTableCQ {
	q.BaseConditionQuery.RegOBA("testDate")
	return q
}
func (q *TestTableCQ) AddOrderBy_TestDate_Desc() *TestTableCQ {
	q.BaseConditionQuery.RegOBD("testDate")
	return q
}
func (q *TestTableCQ) regTestDate(key *df.ConditionKey, value interface{}) {
	if q.TestDate == nil {
		q.TestDate = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TestDate, "testDate")
}

func (q *TestTableCQ) getCValueTestTimestamp() *df.ConditionValue {
	if q.TestTimestamp == nil {
		q.TestTimestamp = new(df.ConditionValue)
	}
	return q.TestTimestamp
}




func (q *TestTableCQ) SetTestTimestamp_Equal(value df.Timestamp) *TestTableCQ {
	q.regTestTimestamp(df.CK_EQ_C, value)
	return q
}


func (q *TestTableCQ) SetTestTimestamp_GreaterThan(value df.Timestamp) *TestTableCQ {
	q.regTestTimestamp(df.CK_GT_C, value)
	return q
}

func (q *TestTableCQ) SetTestTimestamp_LessThan(value df.Timestamp) *TestTableCQ {
	q.regTestTimestamp(df.CK_LT_C, value)
	return q
}

func (q *TestTableCQ) SetTestTimestamp_GreaterEqual(value df.Timestamp) *TestTableCQ {
	q.regTestTimestamp(df.CK_GE_C, value)
	return q
}

func (q *TestTableCQ) SetTestTimestamp_LessEqual(value df.Timestamp) *TestTableCQ {
	q.regTestTimestamp(df.CK_LE_C, value)
	return q
}

func (q *TestTableCQ) SetTestTimestamp_IsNull() *TestTableCQ {
	q.regTestTimestamp(df.CK_ISN_C, 0)
	return q
}
func (q *TestTableCQ) SetTestTimestamp_IsNotNull() *TestTableCQ {
	q.regTestTimestamp(df.CK_ISNN_C, 0)
	return q
}
func (q *TestTableCQ) AddOrderBy_TestTimestamp_Asc() *TestTableCQ {
	q.BaseConditionQuery.RegOBA("testTimestamp")
	return q
}
func (q *TestTableCQ) AddOrderBy_TestTimestamp_Desc() *TestTableCQ {
	q.BaseConditionQuery.RegOBD("testTimestamp")
	return q
}
func (q *TestTableCQ) regTestTimestamp(key *df.ConditionKey, value interface{}) {
	if q.TestTimestamp == nil {
		q.TestTimestamp = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TestTimestamp, "testTimestamp")
}

func (q *TestTableCQ) getCValueTestNbr() *df.ConditionValue {
	if q.TestNbr == nil {
		q.TestNbr = new(df.ConditionValue)
	}
	return q.TestNbr
}



func (q *TestTableCQ) SetTestNbr_Equal(value int64) *TestTableCQ {
	q.regTestNbr(df.CK_EQ_C, value)
	return q
}
func (q *TestTableCQ) SetTestNbr_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueTestNbr(), "testNbr")
}
func (q *TestTableCQ) SetTestNbr_NotEqual(value int64) *TestTableCQ {
	q.regTestNbr(df.CK_NE_C, value)
	return q
}

func (q *TestTableCQ) SetTestNbr_GreaterThan(value int64) *TestTableCQ {
	q.regTestNbr(df.CK_GT_C, value)
	return q
}

func (q *TestTableCQ) SetTestNbr_LessThan(value int64) *TestTableCQ {
	q.regTestNbr(df.CK_LT_C, value)
	return q
}

func (q *TestTableCQ) SetTestNbr_GreaterEqual(value int64) *TestTableCQ {
	q.regTestNbr(df.CK_GE_C, value)
	return q
}

func (q *TestTableCQ) SetTestNbr_LessEqual(value int64) *TestTableCQ {
	q.regTestNbr(df.CK_LE_C, value)
	return q
}
func (q *TestTableCQ) SetTestNbr_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueTestNbr(),"testNbr",rangeOfOption)
}	


func (q *TestTableCQ) SetTestNbr_IsNull() *TestTableCQ {
	q.regTestNbr(df.CK_ISN_C, 0)
	return q
}
func (q *TestTableCQ) SetTestNbr_IsNotNull() *TestTableCQ {
	q.regTestNbr(df.CK_ISNN_C, 0)
	return q
}
func (q *TestTableCQ) AddOrderBy_TestNbr_Asc() *TestTableCQ {
	q.BaseConditionQuery.RegOBA("testNbr")
	return q
}
func (q *TestTableCQ) AddOrderBy_TestNbr_Desc() *TestTableCQ {
	q.BaseConditionQuery.RegOBD("testNbr")
	return q
}
func (q *TestTableCQ) regTestNbr(key *df.ConditionKey, value interface{}) {
	if q.TestNbr == nil {
		q.TestNbr = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TestNbr, "testNbr")
}

func (q *TestTableCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *TestTableCQ) SetVersionNo_Equal(value int64) *TestTableCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}
func (q *TestTableCQ) SetVersionNo_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueVersionNo(), "versionNo")
}
func (q *TestTableCQ) SetVersionNo_NotEqual(value int64) *TestTableCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *TestTableCQ) SetVersionNo_GreaterThan(value int64) *TestTableCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *TestTableCQ) SetVersionNo_LessThan(value int64) *TestTableCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *TestTableCQ) SetVersionNo_GreaterEqual(value int64) *TestTableCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *TestTableCQ) SetVersionNo_LessEqual(value int64) *TestTableCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *TestTableCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *TestTableCQ) AddOrderBy_VersionNo_Asc() *TestTableCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *TestTableCQ) AddOrderBy_VersionNo_Desc() *TestTableCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *TestTableCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}

func (q *TestTableCQ) getCValueDelFlag() *df.ConditionValue {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	return q.DelFlag
}



func (q *TestTableCQ) SetDelFlag_Equal(value int64) *TestTableCQ {
	q.regDelFlag(df.CK_EQ_C, value)
	return q
}
func (q *TestTableCQ) SetDelFlag_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDelFlag(), "delFlag")
}
func (q *TestTableCQ) SetDelFlag_NotEqual(value int64) *TestTableCQ {
	q.regDelFlag(df.CK_NE_C, value)
	return q
}

func (q *TestTableCQ) SetDelFlag_GreaterThan(value int64) *TestTableCQ {
	q.regDelFlag(df.CK_GT_C, value)
	return q
}

func (q *TestTableCQ) SetDelFlag_LessThan(value int64) *TestTableCQ {
	q.regDelFlag(df.CK_LT_C, value)
	return q
}

func (q *TestTableCQ) SetDelFlag_GreaterEqual(value int64) *TestTableCQ {
	q.regDelFlag(df.CK_GE_C, value)
	return q
}

func (q *TestTableCQ) SetDelFlag_LessEqual(value int64) *TestTableCQ {
	q.regDelFlag(df.CK_LE_C, value)
	return q
}
func (q *TestTableCQ) SetDelFlag_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDelFlag(),"delFlag",rangeOfOption)
}	


func (q *TestTableCQ) AddOrderBy_DelFlag_Asc() *TestTableCQ {
	q.BaseConditionQuery.RegOBA("delFlag")
	return q
}
func (q *TestTableCQ) AddOrderBy_DelFlag_Desc() *TestTableCQ {
	q.BaseConditionQuery.RegOBD("delFlag")
	return q
}
func (q *TestTableCQ) regDelFlag(key *df.ConditionKey, value interface{}) {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DelFlag, "delFlag")
}

func (q *TestTableCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *TestTableCQ) SetRegisterDatetime_Equal(value df.Timestamp) *TestTableCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *TestTableCQ) SetRegisterDatetime_GreaterThan(value df.Timestamp) *TestTableCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *TestTableCQ) SetRegisterDatetime_LessThan(value df.Timestamp) *TestTableCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *TestTableCQ) SetRegisterDatetime_GreaterEqual(value df.Timestamp) *TestTableCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *TestTableCQ) SetRegisterDatetime_LessEqual(value df.Timestamp) *TestTableCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *TestTableCQ) AddOrderBy_RegisterDatetime_Asc() *TestTableCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *TestTableCQ) AddOrderBy_RegisterDatetime_Desc() *TestTableCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *TestTableCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *TestTableCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *TestTableCQ) SetRegisterUser_Equal(value string) *TestTableCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *TestTableCQ) SetRegisterUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterUser(), "registerUser")
}
func (q *TestTableCQ) SetRegisterUser_NotEqual(value string) *TestTableCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetRegisterUser_GreaterThan(value string) *TestTableCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetRegisterUser_LessThan(value string) *TestTableCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetRegisterUser_GreaterEqual(value string) *TestTableCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *TestTableCQ) SetRegisterUser_LessEqual(value string) *TestTableCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *TestTableCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *TestTableCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *TestTableCQ) AddOrderBy_RegisterUser_Asc() *TestTableCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *TestTableCQ) AddOrderBy_RegisterUser_Desc() *TestTableCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *TestTableCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *TestTableCQ) getCValueRegisterProcess() *df.ConditionValue {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	return q.RegisterProcess
}


func (q *TestTableCQ) SetRegisterProcess_Equal(value string) *TestTableCQ {
	q.regRegisterProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *TestTableCQ) SetRegisterProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterProcess(), "registerProcess")
}
func (q *TestTableCQ) SetRegisterProcess_NotEqual(value string) *TestTableCQ {
	q.regRegisterProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetRegisterProcess_GreaterThan(value string) *TestTableCQ {
	q.regRegisterProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetRegisterProcess_LessThan(value string) *TestTableCQ {
	q.regRegisterProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetRegisterProcess_GreaterEqual(value string) *TestTableCQ {
	q.regRegisterProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *TestTableCQ) SetRegisterProcess_LessEqual(value string) *TestTableCQ {
	q.regRegisterProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetRegisterProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}

func (q *TestTableCQ) SetRegisterProcess_PrefixSearch(value string) error {
	return q.SetRegisterProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *TestTableCQ) SetRegisterProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}



func (q *TestTableCQ) AddOrderBy_RegisterProcess_Asc() *TestTableCQ {
	q.BaseConditionQuery.RegOBA("registerProcess")
	return q
}
func (q *TestTableCQ) AddOrderBy_RegisterProcess_Desc() *TestTableCQ {
	q.BaseConditionQuery.RegOBD("registerProcess")
	return q
}
func (q *TestTableCQ) regRegisterProcess(key *df.ConditionKey, value interface{}) {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterProcess, "registerProcess")
}

func (q *TestTableCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *TestTableCQ) SetUpdateDatetime_Equal(value df.Timestamp) *TestTableCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *TestTableCQ) SetUpdateDatetime_GreaterThan(value df.Timestamp) *TestTableCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *TestTableCQ) SetUpdateDatetime_LessThan(value df.Timestamp) *TestTableCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *TestTableCQ) SetUpdateDatetime_GreaterEqual(value df.Timestamp) *TestTableCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *TestTableCQ) SetUpdateDatetime_LessEqual(value df.Timestamp) *TestTableCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *TestTableCQ) AddOrderBy_UpdateDatetime_Asc() *TestTableCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *TestTableCQ) AddOrderBy_UpdateDatetime_Desc() *TestTableCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *TestTableCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *TestTableCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *TestTableCQ) SetUpdateUser_Equal(value string) *TestTableCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *TestTableCQ) SetUpdateUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateUser(), "updateUser")
}
func (q *TestTableCQ) SetUpdateUser_NotEqual(value string) *TestTableCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetUpdateUser_GreaterThan(value string) *TestTableCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetUpdateUser_LessThan(value string) *TestTableCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetUpdateUser_GreaterEqual(value string) *TestTableCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *TestTableCQ) SetUpdateUser_LessEqual(value string) *TestTableCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *TestTableCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *TestTableCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *TestTableCQ) AddOrderBy_UpdateUser_Asc() *TestTableCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *TestTableCQ) AddOrderBy_UpdateUser_Desc() *TestTableCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *TestTableCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *TestTableCQ) getCValueUpdateProcess() *df.ConditionValue {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	return q.UpdateProcess
}


func (q *TestTableCQ) SetUpdateProcess_Equal(value string) *TestTableCQ {
	q.regUpdateProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *TestTableCQ) SetUpdateProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateProcess(), "updateProcess")
}
func (q *TestTableCQ) SetUpdateProcess_NotEqual(value string) *TestTableCQ {
	q.regUpdateProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetUpdateProcess_GreaterThan(value string) *TestTableCQ {
	q.regUpdateProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetUpdateProcess_LessThan(value string) *TestTableCQ {
	q.regUpdateProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetUpdateProcess_GreaterEqual(value string) *TestTableCQ {
	q.regUpdateProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *TestTableCQ) SetUpdateProcess_LessEqual(value string) *TestTableCQ {
	q.regUpdateProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *TestTableCQ) SetUpdateProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}

func (q *TestTableCQ) SetUpdateProcess_PrefixSearch(value string) error {
	return q.SetUpdateProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *TestTableCQ) SetUpdateProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}



func (q *TestTableCQ) AddOrderBy_UpdateProcess_Asc() *TestTableCQ {
	q.BaseConditionQuery.RegOBA("updateProcess")
	return q
}
func (q *TestTableCQ) AddOrderBy_UpdateProcess_Desc() *TestTableCQ {
	q.BaseConditionQuery.RegOBD("updateProcess")
	return q
}
func (q *TestTableCQ) regUpdateProcess(key *df.ConditionKey, value interface{}) {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateProcess, "updateProcess")
}


func CreateTestTableCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *TestTableCQ {
	cq := new(TestTableCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "TestTable"
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
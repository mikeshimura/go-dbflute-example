package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type SysTableCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	Id *df.ConditionValue
	TableName *df.ConditionValue
	Key1 *df.ConditionValue
	Key2 *df.ConditionValue
	S1Data *df.ConditionValue
	S2Data *df.ConditionValue
	S3Data *df.ConditionValue
	N1Data *df.ConditionValue
	N2Data *df.ConditionValue
	N3Data *df.ConditionValue
	VersionNo *df.ConditionValue
	DelFlag *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterUser *df.ConditionValue
	RegisterProcess *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateUser *df.ConditionValue
	UpdateProcess *df.ConditionValue
}

func (q *SysTableCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *SysTableCQ) getCValueId() *df.ConditionValue {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	return q.Id
}



func (q *SysTableCQ) SetId_Equal(value int64) *SysTableCQ {
	q.regId(df.CK_EQ_C, value)
	return q
}
func (q *SysTableCQ) SetId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueId(), "id")
}
func (q *SysTableCQ) SetId_NotEqual(value int64) *SysTableCQ {
	q.regId(df.CK_NE_C, value)
	return q
}

func (q *SysTableCQ) SetId_GreaterThan(value int64) *SysTableCQ {
	q.regId(df.CK_GT_C, value)
	return q
}

func (q *SysTableCQ) SetId_LessThan(value int64) *SysTableCQ {
	q.regId(df.CK_LT_C, value)
	return q
}

func (q *SysTableCQ) SetId_GreaterEqual(value int64) *SysTableCQ {
	q.regId(df.CK_GE_C, value)
	return q
}

func (q *SysTableCQ) SetId_LessEqual(value int64) *SysTableCQ {
	q.regId(df.CK_LE_C, value)
	return q
}
func (q *SysTableCQ) SetId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueId(),"id",rangeOfOption)
}	


func (q *SysTableCQ) SetId_IsNull() *SysTableCQ {
	q.regId(df.CK_ISN_C, 0)
	return q
}
func (q *SysTableCQ) SetId_IsNotNull() *SysTableCQ {
	q.regId(df.CK_ISNN_C, 0)
	return q
}
func (q *SysTableCQ) AddOrderBy_Id_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("id")
	return q
}
func (q *SysTableCQ) AddOrderBy_Id_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("id")
	return q
}
func (q *SysTableCQ) regId(key *df.ConditionKey, value interface{}) {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Id, "id")
}

func (q *SysTableCQ) getCValueTableName() *df.ConditionValue {
	if q.TableName == nil {
		q.TableName = new(df.ConditionValue)
	}
	return q.TableName
}


func (q *SysTableCQ) SetTableName_Equal(value string) *SysTableCQ {
	q.regTableName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SysTableCQ) SetTableName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueTableName(), "tableName")
}
func (q *SysTableCQ) SetTableName_NotEqual(value string) *SysTableCQ {
	q.regTableName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetTableName_GreaterThan(value string) *SysTableCQ {
	q.regTableName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetTableName_LessThan(value string) *SysTableCQ {
	q.regTableName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetTableName_GreaterEqual(value string) *SysTableCQ {
	q.regTableName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SysTableCQ) SetTableName_LessEqual(value string) *SysTableCQ {
	q.regTableName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetTableName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueTableName(), "tableName", option)
}

func (q *SysTableCQ) SetTableName_PrefixSearch(value string) error {
	return q.SetTableName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SysTableCQ) SetTableName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueTableName(), "tableName", option)
}



func (q *SysTableCQ) AddOrderBy_TableName_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("tableName")
	return q
}
func (q *SysTableCQ) AddOrderBy_TableName_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("tableName")
	return q
}
func (q *SysTableCQ) regTableName(key *df.ConditionKey, value interface{}) {
	if q.TableName == nil {
		q.TableName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TableName, "tableName")
}

func (q *SysTableCQ) getCValueKey1() *df.ConditionValue {
	if q.Key1 == nil {
		q.Key1 = new(df.ConditionValue)
	}
	return q.Key1
}


func (q *SysTableCQ) SetKey1_Equal(value string) *SysTableCQ {
	q.regKey1(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SysTableCQ) SetKey1_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueKey1(), "key1")
}
func (q *SysTableCQ) SetKey1_NotEqual(value string) *SysTableCQ {
	q.regKey1(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetKey1_GreaterThan(value string) *SysTableCQ {
	q.regKey1(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetKey1_LessThan(value string) *SysTableCQ {
	q.regKey1(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetKey1_GreaterEqual(value string) *SysTableCQ {
	q.regKey1(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SysTableCQ) SetKey1_LessEqual(value string) *SysTableCQ {
	q.regKey1(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetKey1_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueKey1(), "key1", option)
}

func (q *SysTableCQ) SetKey1_PrefixSearch(value string) error {
	return q.SetKey1_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SysTableCQ) SetKey1_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueKey1(), "key1", option)
}



func (q *SysTableCQ) AddOrderBy_Key1_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("key1")
	return q
}
func (q *SysTableCQ) AddOrderBy_Key1_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("key1")
	return q
}
func (q *SysTableCQ) regKey1(key *df.ConditionKey, value interface{}) {
	if q.Key1 == nil {
		q.Key1 = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Key1, "key1")
}

func (q *SysTableCQ) getCValueKey2() *df.ConditionValue {
	if q.Key2 == nil {
		q.Key2 = new(df.ConditionValue)
	}
	return q.Key2
}


func (q *SysTableCQ) SetKey2_Equal(value string) *SysTableCQ {
	q.regKey2(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SysTableCQ) SetKey2_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueKey2(), "key2")
}
func (q *SysTableCQ) SetKey2_NotEqual(value string) *SysTableCQ {
	q.regKey2(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetKey2_GreaterThan(value string) *SysTableCQ {
	q.regKey2(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetKey2_LessThan(value string) *SysTableCQ {
	q.regKey2(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetKey2_GreaterEqual(value string) *SysTableCQ {
	q.regKey2(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SysTableCQ) SetKey2_LessEqual(value string) *SysTableCQ {
	q.regKey2(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetKey2_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueKey2(), "key2", option)
}

func (q *SysTableCQ) SetKey2_PrefixSearch(value string) error {
	return q.SetKey2_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SysTableCQ) SetKey2_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueKey2(), "key2", option)
}



func (q *SysTableCQ) AddOrderBy_Key2_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("key2")
	return q
}
func (q *SysTableCQ) AddOrderBy_Key2_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("key2")
	return q
}
func (q *SysTableCQ) regKey2(key *df.ConditionKey, value interface{}) {
	if q.Key2 == nil {
		q.Key2 = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Key2, "key2")
}

func (q *SysTableCQ) getCValueS1Data() *df.ConditionValue {
	if q.S1Data == nil {
		q.S1Data = new(df.ConditionValue)
	}
	return q.S1Data
}


func (q *SysTableCQ) SetS1Data_Equal(value string) *SysTableCQ {
	q.regS1Data(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SysTableCQ) SetS1Data_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueS1Data(), "s1Data")
}
func (q *SysTableCQ) SetS1Data_NotEqual(value string) *SysTableCQ {
	q.regS1Data(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetS1Data_GreaterThan(value string) *SysTableCQ {
	q.regS1Data(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetS1Data_LessThan(value string) *SysTableCQ {
	q.regS1Data(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetS1Data_GreaterEqual(value string) *SysTableCQ {
	q.regS1Data(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SysTableCQ) SetS1Data_LessEqual(value string) *SysTableCQ {
	q.regS1Data(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetS1Data_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueS1Data(), "s1Data", option)
}

func (q *SysTableCQ) SetS1Data_PrefixSearch(value string) error {
	return q.SetS1Data_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SysTableCQ) SetS1Data_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueS1Data(), "s1Data", option)
}



func (q *SysTableCQ) SetS1Data_IsNull() *SysTableCQ {
	q.regS1Data(df.CK_ISN_C, 0)
	return q
}
func (q *SysTableCQ) SetS1Data_IsNullOrEmpty() *SysTableCQ {
	q.regS1Data(df.CK_ISNOE_C, 0)
	return q
}
func (q *SysTableCQ) SetS1Data_IsNotNull() *SysTableCQ {
	q.regS1Data(df.CK_ISNN_C, 0)
	return q
}
func (q *SysTableCQ) AddOrderBy_S1Data_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("s1Data")
	return q
}
func (q *SysTableCQ) AddOrderBy_S1Data_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("s1Data")
	return q
}
func (q *SysTableCQ) regS1Data(key *df.ConditionKey, value interface{}) {
	if q.S1Data == nil {
		q.S1Data = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.S1Data, "s1Data")
}

func (q *SysTableCQ) getCValueS2Data() *df.ConditionValue {
	if q.S2Data == nil {
		q.S2Data = new(df.ConditionValue)
	}
	return q.S2Data
}


func (q *SysTableCQ) SetS2Data_Equal(value string) *SysTableCQ {
	q.regS2Data(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SysTableCQ) SetS2Data_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueS2Data(), "s2Data")
}
func (q *SysTableCQ) SetS2Data_NotEqual(value string) *SysTableCQ {
	q.regS2Data(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetS2Data_GreaterThan(value string) *SysTableCQ {
	q.regS2Data(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetS2Data_LessThan(value string) *SysTableCQ {
	q.regS2Data(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetS2Data_GreaterEqual(value string) *SysTableCQ {
	q.regS2Data(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SysTableCQ) SetS2Data_LessEqual(value string) *SysTableCQ {
	q.regS2Data(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetS2Data_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueS2Data(), "s2Data", option)
}

func (q *SysTableCQ) SetS2Data_PrefixSearch(value string) error {
	return q.SetS2Data_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SysTableCQ) SetS2Data_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueS2Data(), "s2Data", option)
}



func (q *SysTableCQ) SetS2Data_IsNull() *SysTableCQ {
	q.regS2Data(df.CK_ISN_C, 0)
	return q
}
func (q *SysTableCQ) SetS2Data_IsNullOrEmpty() *SysTableCQ {
	q.regS2Data(df.CK_ISNOE_C, 0)
	return q
}
func (q *SysTableCQ) SetS2Data_IsNotNull() *SysTableCQ {
	q.regS2Data(df.CK_ISNN_C, 0)
	return q
}
func (q *SysTableCQ) AddOrderBy_S2Data_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("s2Data")
	return q
}
func (q *SysTableCQ) AddOrderBy_S2Data_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("s2Data")
	return q
}
func (q *SysTableCQ) regS2Data(key *df.ConditionKey, value interface{}) {
	if q.S2Data == nil {
		q.S2Data = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.S2Data, "s2Data")
}

func (q *SysTableCQ) getCValueS3Data() *df.ConditionValue {
	if q.S3Data == nil {
		q.S3Data = new(df.ConditionValue)
	}
	return q.S3Data
}


func (q *SysTableCQ) SetS3Data_Equal(value string) *SysTableCQ {
	q.regS3Data(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SysTableCQ) SetS3Data_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueS3Data(), "s3Data")
}
func (q *SysTableCQ) SetS3Data_NotEqual(value string) *SysTableCQ {
	q.regS3Data(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetS3Data_GreaterThan(value string) *SysTableCQ {
	q.regS3Data(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetS3Data_LessThan(value string) *SysTableCQ {
	q.regS3Data(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetS3Data_GreaterEqual(value string) *SysTableCQ {
	q.regS3Data(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SysTableCQ) SetS3Data_LessEqual(value string) *SysTableCQ {
	q.regS3Data(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetS3Data_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueS3Data(), "s3Data", option)
}

func (q *SysTableCQ) SetS3Data_PrefixSearch(value string) error {
	return q.SetS3Data_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SysTableCQ) SetS3Data_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueS3Data(), "s3Data", option)
}



func (q *SysTableCQ) SetS3Data_IsNull() *SysTableCQ {
	q.regS3Data(df.CK_ISN_C, 0)
	return q
}
func (q *SysTableCQ) SetS3Data_IsNullOrEmpty() *SysTableCQ {
	q.regS3Data(df.CK_ISNOE_C, 0)
	return q
}
func (q *SysTableCQ) SetS3Data_IsNotNull() *SysTableCQ {
	q.regS3Data(df.CK_ISNN_C, 0)
	return q
}
func (q *SysTableCQ) AddOrderBy_S3Data_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("s3Data")
	return q
}
func (q *SysTableCQ) AddOrderBy_S3Data_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("s3Data")
	return q
}
func (q *SysTableCQ) regS3Data(key *df.ConditionKey, value interface{}) {
	if q.S3Data == nil {
		q.S3Data = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.S3Data, "s3Data")
}

func (q *SysTableCQ) getCValueN1Data() *df.ConditionValue {
	if q.N1Data == nil {
		q.N1Data = new(df.ConditionValue)
	}
	return q.N1Data
}



func (q *SysTableCQ) SetN1Data_Equal(value df.Numeric) *SysTableCQ {
	q.regN1Data(df.CK_EQ_C, value)
	return q
}
func (q *SysTableCQ) SetN1Data_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueN1Data(), "n1Data")
}
func (q *SysTableCQ) SetN1Data_NotEqual(value df.Numeric) *SysTableCQ {
	q.regN1Data(df.CK_NE_C, value)
	return q
}

func (q *SysTableCQ) SetN1Data_GreaterThan(value df.Numeric) *SysTableCQ {
	q.regN1Data(df.CK_GT_C, value)
	return q
}

func (q *SysTableCQ) SetN1Data_LessThan(value df.Numeric) *SysTableCQ {
	q.regN1Data(df.CK_LT_C, value)
	return q
}

func (q *SysTableCQ) SetN1Data_GreaterEqual(value df.Numeric) *SysTableCQ {
	q.regN1Data(df.CK_GE_C, value)
	return q
}

func (q *SysTableCQ) SetN1Data_LessEqual(value df.Numeric) *SysTableCQ {
	q.regN1Data(df.CK_LE_C, value)
	return q
}
func (q *SysTableCQ) SetN1Data_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueN1Data(),"n1Data",rangeOfOption)
}	


func (q *SysTableCQ) SetN1Data_IsNull() *SysTableCQ {
	q.regN1Data(df.CK_ISN_C, 0)
	return q
}
func (q *SysTableCQ) SetN1Data_IsNotNull() *SysTableCQ {
	q.regN1Data(df.CK_ISNN_C, 0)
	return q
}
func (q *SysTableCQ) AddOrderBy_N1Data_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("n1Data")
	return q
}
func (q *SysTableCQ) AddOrderBy_N1Data_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("n1Data")
	return q
}
func (q *SysTableCQ) regN1Data(key *df.ConditionKey, value interface{}) {
	if q.N1Data == nil {
		q.N1Data = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.N1Data, "n1Data")
}

func (q *SysTableCQ) getCValueN2Data() *df.ConditionValue {
	if q.N2Data == nil {
		q.N2Data = new(df.ConditionValue)
	}
	return q.N2Data
}



func (q *SysTableCQ) SetN2Data_Equal(value df.Numeric) *SysTableCQ {
	q.regN2Data(df.CK_EQ_C, value)
	return q
}
func (q *SysTableCQ) SetN2Data_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueN2Data(), "n2Data")
}
func (q *SysTableCQ) SetN2Data_NotEqual(value df.Numeric) *SysTableCQ {
	q.regN2Data(df.CK_NE_C, value)
	return q
}

func (q *SysTableCQ) SetN2Data_GreaterThan(value df.Numeric) *SysTableCQ {
	q.regN2Data(df.CK_GT_C, value)
	return q
}

func (q *SysTableCQ) SetN2Data_LessThan(value df.Numeric) *SysTableCQ {
	q.regN2Data(df.CK_LT_C, value)
	return q
}

func (q *SysTableCQ) SetN2Data_GreaterEqual(value df.Numeric) *SysTableCQ {
	q.regN2Data(df.CK_GE_C, value)
	return q
}

func (q *SysTableCQ) SetN2Data_LessEqual(value df.Numeric) *SysTableCQ {
	q.regN2Data(df.CK_LE_C, value)
	return q
}
func (q *SysTableCQ) SetN2Data_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueN2Data(),"n2Data",rangeOfOption)
}	


func (q *SysTableCQ) SetN2Data_IsNull() *SysTableCQ {
	q.regN2Data(df.CK_ISN_C, 0)
	return q
}
func (q *SysTableCQ) SetN2Data_IsNotNull() *SysTableCQ {
	q.regN2Data(df.CK_ISNN_C, 0)
	return q
}
func (q *SysTableCQ) AddOrderBy_N2Data_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("n2Data")
	return q
}
func (q *SysTableCQ) AddOrderBy_N2Data_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("n2Data")
	return q
}
func (q *SysTableCQ) regN2Data(key *df.ConditionKey, value interface{}) {
	if q.N2Data == nil {
		q.N2Data = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.N2Data, "n2Data")
}

func (q *SysTableCQ) getCValueN3Data() *df.ConditionValue {
	if q.N3Data == nil {
		q.N3Data = new(df.ConditionValue)
	}
	return q.N3Data
}



func (q *SysTableCQ) SetN3Data_Equal(value df.Numeric) *SysTableCQ {
	q.regN3Data(df.CK_EQ_C, value)
	return q
}
func (q *SysTableCQ) SetN3Data_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueN3Data(), "n3Data")
}
func (q *SysTableCQ) SetN3Data_NotEqual(value df.Numeric) *SysTableCQ {
	q.regN3Data(df.CK_NE_C, value)
	return q
}

func (q *SysTableCQ) SetN3Data_GreaterThan(value df.Numeric) *SysTableCQ {
	q.regN3Data(df.CK_GT_C, value)
	return q
}

func (q *SysTableCQ) SetN3Data_LessThan(value df.Numeric) *SysTableCQ {
	q.regN3Data(df.CK_LT_C, value)
	return q
}

func (q *SysTableCQ) SetN3Data_GreaterEqual(value df.Numeric) *SysTableCQ {
	q.regN3Data(df.CK_GE_C, value)
	return q
}

func (q *SysTableCQ) SetN3Data_LessEqual(value df.Numeric) *SysTableCQ {
	q.regN3Data(df.CK_LE_C, value)
	return q
}
func (q *SysTableCQ) SetN3Data_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueN3Data(),"n3Data",rangeOfOption)
}	


func (q *SysTableCQ) SetN3Data_IsNull() *SysTableCQ {
	q.regN3Data(df.CK_ISN_C, 0)
	return q
}
func (q *SysTableCQ) SetN3Data_IsNotNull() *SysTableCQ {
	q.regN3Data(df.CK_ISNN_C, 0)
	return q
}
func (q *SysTableCQ) AddOrderBy_N3Data_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("n3Data")
	return q
}
func (q *SysTableCQ) AddOrderBy_N3Data_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("n3Data")
	return q
}
func (q *SysTableCQ) regN3Data(key *df.ConditionKey, value interface{}) {
	if q.N3Data == nil {
		q.N3Data = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.N3Data, "n3Data")
}

func (q *SysTableCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *SysTableCQ) SetVersionNo_Equal(value int64) *SysTableCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}
func (q *SysTableCQ) SetVersionNo_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueVersionNo(), "versionNo")
}
func (q *SysTableCQ) SetVersionNo_NotEqual(value int64) *SysTableCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *SysTableCQ) SetVersionNo_GreaterThan(value int64) *SysTableCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *SysTableCQ) SetVersionNo_LessThan(value int64) *SysTableCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *SysTableCQ) SetVersionNo_GreaterEqual(value int64) *SysTableCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *SysTableCQ) SetVersionNo_LessEqual(value int64) *SysTableCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *SysTableCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *SysTableCQ) AddOrderBy_VersionNo_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *SysTableCQ) AddOrderBy_VersionNo_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *SysTableCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}

func (q *SysTableCQ) getCValueDelFlag() *df.ConditionValue {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	return q.DelFlag
}



func (q *SysTableCQ) SetDelFlag_Equal(value int64) *SysTableCQ {
	q.regDelFlag(df.CK_EQ_C, value)
	return q
}
func (q *SysTableCQ) SetDelFlag_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDelFlag(), "delFlag")
}
func (q *SysTableCQ) SetDelFlag_NotEqual(value int64) *SysTableCQ {
	q.regDelFlag(df.CK_NE_C, value)
	return q
}

func (q *SysTableCQ) SetDelFlag_GreaterThan(value int64) *SysTableCQ {
	q.regDelFlag(df.CK_GT_C, value)
	return q
}

func (q *SysTableCQ) SetDelFlag_LessThan(value int64) *SysTableCQ {
	q.regDelFlag(df.CK_LT_C, value)
	return q
}

func (q *SysTableCQ) SetDelFlag_GreaterEqual(value int64) *SysTableCQ {
	q.regDelFlag(df.CK_GE_C, value)
	return q
}

func (q *SysTableCQ) SetDelFlag_LessEqual(value int64) *SysTableCQ {
	q.regDelFlag(df.CK_LE_C, value)
	return q
}
func (q *SysTableCQ) SetDelFlag_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDelFlag(),"delFlag",rangeOfOption)
}	


func (q *SysTableCQ) AddOrderBy_DelFlag_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("delFlag")
	return q
}
func (q *SysTableCQ) AddOrderBy_DelFlag_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("delFlag")
	return q
}
func (q *SysTableCQ) regDelFlag(key *df.ConditionKey, value interface{}) {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DelFlag, "delFlag")
}

func (q *SysTableCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *SysTableCQ) SetRegisterDatetime_Equal(value df.Timestamp) *SysTableCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *SysTableCQ) SetRegisterDatetime_GreaterThan(value df.Timestamp) *SysTableCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *SysTableCQ) SetRegisterDatetime_LessThan(value df.Timestamp) *SysTableCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *SysTableCQ) SetRegisterDatetime_GreaterEqual(value df.Timestamp) *SysTableCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *SysTableCQ) SetRegisterDatetime_LessEqual(value df.Timestamp) *SysTableCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *SysTableCQ) AddOrderBy_RegisterDatetime_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *SysTableCQ) AddOrderBy_RegisterDatetime_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *SysTableCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *SysTableCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *SysTableCQ) SetRegisterUser_Equal(value string) *SysTableCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SysTableCQ) SetRegisterUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterUser(), "registerUser")
}
func (q *SysTableCQ) SetRegisterUser_NotEqual(value string) *SysTableCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetRegisterUser_GreaterThan(value string) *SysTableCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetRegisterUser_LessThan(value string) *SysTableCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetRegisterUser_GreaterEqual(value string) *SysTableCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SysTableCQ) SetRegisterUser_LessEqual(value string) *SysTableCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *SysTableCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SysTableCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *SysTableCQ) AddOrderBy_RegisterUser_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *SysTableCQ) AddOrderBy_RegisterUser_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *SysTableCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *SysTableCQ) getCValueRegisterProcess() *df.ConditionValue {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	return q.RegisterProcess
}


func (q *SysTableCQ) SetRegisterProcess_Equal(value string) *SysTableCQ {
	q.regRegisterProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SysTableCQ) SetRegisterProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterProcess(), "registerProcess")
}
func (q *SysTableCQ) SetRegisterProcess_NotEqual(value string) *SysTableCQ {
	q.regRegisterProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetRegisterProcess_GreaterThan(value string) *SysTableCQ {
	q.regRegisterProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetRegisterProcess_LessThan(value string) *SysTableCQ {
	q.regRegisterProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetRegisterProcess_GreaterEqual(value string) *SysTableCQ {
	q.regRegisterProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SysTableCQ) SetRegisterProcess_LessEqual(value string) *SysTableCQ {
	q.regRegisterProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetRegisterProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}

func (q *SysTableCQ) SetRegisterProcess_PrefixSearch(value string) error {
	return q.SetRegisterProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SysTableCQ) SetRegisterProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}



func (q *SysTableCQ) AddOrderBy_RegisterProcess_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("registerProcess")
	return q
}
func (q *SysTableCQ) AddOrderBy_RegisterProcess_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("registerProcess")
	return q
}
func (q *SysTableCQ) regRegisterProcess(key *df.ConditionKey, value interface{}) {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterProcess, "registerProcess")
}

func (q *SysTableCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *SysTableCQ) SetUpdateDatetime_Equal(value df.Timestamp) *SysTableCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *SysTableCQ) SetUpdateDatetime_GreaterThan(value df.Timestamp) *SysTableCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *SysTableCQ) SetUpdateDatetime_LessThan(value df.Timestamp) *SysTableCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *SysTableCQ) SetUpdateDatetime_GreaterEqual(value df.Timestamp) *SysTableCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *SysTableCQ) SetUpdateDatetime_LessEqual(value df.Timestamp) *SysTableCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *SysTableCQ) AddOrderBy_UpdateDatetime_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *SysTableCQ) AddOrderBy_UpdateDatetime_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *SysTableCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *SysTableCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *SysTableCQ) SetUpdateUser_Equal(value string) *SysTableCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SysTableCQ) SetUpdateUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateUser(), "updateUser")
}
func (q *SysTableCQ) SetUpdateUser_NotEqual(value string) *SysTableCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetUpdateUser_GreaterThan(value string) *SysTableCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetUpdateUser_LessThan(value string) *SysTableCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetUpdateUser_GreaterEqual(value string) *SysTableCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SysTableCQ) SetUpdateUser_LessEqual(value string) *SysTableCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *SysTableCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SysTableCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *SysTableCQ) AddOrderBy_UpdateUser_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *SysTableCQ) AddOrderBy_UpdateUser_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *SysTableCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *SysTableCQ) getCValueUpdateProcess() *df.ConditionValue {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	return q.UpdateProcess
}


func (q *SysTableCQ) SetUpdateProcess_Equal(value string) *SysTableCQ {
	q.regUpdateProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SysTableCQ) SetUpdateProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateProcess(), "updateProcess")
}
func (q *SysTableCQ) SetUpdateProcess_NotEqual(value string) *SysTableCQ {
	q.regUpdateProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetUpdateProcess_GreaterThan(value string) *SysTableCQ {
	q.regUpdateProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetUpdateProcess_LessThan(value string) *SysTableCQ {
	q.regUpdateProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetUpdateProcess_GreaterEqual(value string) *SysTableCQ {
	q.regUpdateProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SysTableCQ) SetUpdateProcess_LessEqual(value string) *SysTableCQ {
	q.regUpdateProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SysTableCQ) SetUpdateProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}

func (q *SysTableCQ) SetUpdateProcess_PrefixSearch(value string) error {
	return q.SetUpdateProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SysTableCQ) SetUpdateProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}



func (q *SysTableCQ) AddOrderBy_UpdateProcess_Asc() *SysTableCQ {
	q.BaseConditionQuery.RegOBA("updateProcess")
	return q
}
func (q *SysTableCQ) AddOrderBy_UpdateProcess_Desc() *SysTableCQ {
	q.BaseConditionQuery.RegOBD("updateProcess")
	return q
}
func (q *SysTableCQ) regUpdateProcess(key *df.ConditionKey, value interface{}) {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateProcess, "updateProcess")
}


func CreateSysTableCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *SysTableCQ {
	cq := new(SysTableCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "SysTable"
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

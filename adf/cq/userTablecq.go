package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type UserTableCQ struct {
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

func (q *UserTableCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *UserTableCQ) getCValueId() *df.ConditionValue {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	return q.Id
}



func (q *UserTableCQ) SetId_Equal(value int64) *UserTableCQ {
	q.regId(df.CK_EQ_C, value)
	return q
}
func (q *UserTableCQ) SetId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueId(), "id")
}
func (q *UserTableCQ) SetId_NotEqual(value int64) *UserTableCQ {
	q.regId(df.CK_NE_C, value)
	return q
}

func (q *UserTableCQ) SetId_GreaterThan(value int64) *UserTableCQ {
	q.regId(df.CK_GT_C, value)
	return q
}

func (q *UserTableCQ) SetId_LessThan(value int64) *UserTableCQ {
	q.regId(df.CK_LT_C, value)
	return q
}

func (q *UserTableCQ) SetId_GreaterEqual(value int64) *UserTableCQ {
	q.regId(df.CK_GE_C, value)
	return q
}

func (q *UserTableCQ) SetId_LessEqual(value int64) *UserTableCQ {
	q.regId(df.CK_LE_C, value)
	return q
}
func (q *UserTableCQ) SetId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueId(),"id",rangeOfOption)
}	


func (q *UserTableCQ) SetId_IsNull() *UserTableCQ {
	q.regId(df.CK_ISN_C, 0)
	return q
}
func (q *UserTableCQ) SetId_IsNotNull() *UserTableCQ {
	q.regId(df.CK_ISNN_C, 0)
	return q
}
func (q *UserTableCQ) AddOrderBy_Id_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("id")
	return q
}
func (q *UserTableCQ) AddOrderBy_Id_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("id")
	return q
}
func (q *UserTableCQ) regId(key *df.ConditionKey, value interface{}) {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Id, "id")
}

func (q *UserTableCQ) getCValueTableName() *df.ConditionValue {
	if q.TableName == nil {
		q.TableName = new(df.ConditionValue)
	}
	return q.TableName
}


func (q *UserTableCQ) SetTableName_Equal(value string) *UserTableCQ {
	q.regTableName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *UserTableCQ) SetTableName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueTableName(), "tableName")
}
func (q *UserTableCQ) SetTableName_NotEqual(value string) *UserTableCQ {
	q.regTableName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetTableName_GreaterThan(value string) *UserTableCQ {
	q.regTableName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetTableName_LessThan(value string) *UserTableCQ {
	q.regTableName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetTableName_GreaterEqual(value string) *UserTableCQ {
	q.regTableName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *UserTableCQ) SetTableName_LessEqual(value string) *UserTableCQ {
	q.regTableName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetTableName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueTableName(), "tableName", option)
}

func (q *UserTableCQ) SetTableName_PrefixSearch(value string) error {
	return q.SetTableName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *UserTableCQ) SetTableName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueTableName(), "tableName", option)
}



func (q *UserTableCQ) AddOrderBy_TableName_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("tableName")
	return q
}
func (q *UserTableCQ) AddOrderBy_TableName_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("tableName")
	return q
}
func (q *UserTableCQ) regTableName(key *df.ConditionKey, value interface{}) {
	if q.TableName == nil {
		q.TableName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TableName, "tableName")
}

func (q *UserTableCQ) getCValueKey1() *df.ConditionValue {
	if q.Key1 == nil {
		q.Key1 = new(df.ConditionValue)
	}
	return q.Key1
}


func (q *UserTableCQ) SetKey1_Equal(value string) *UserTableCQ {
	q.regKey1(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *UserTableCQ) SetKey1_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueKey1(), "key1")
}
func (q *UserTableCQ) SetKey1_NotEqual(value string) *UserTableCQ {
	q.regKey1(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetKey1_GreaterThan(value string) *UserTableCQ {
	q.regKey1(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetKey1_LessThan(value string) *UserTableCQ {
	q.regKey1(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetKey1_GreaterEqual(value string) *UserTableCQ {
	q.regKey1(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *UserTableCQ) SetKey1_LessEqual(value string) *UserTableCQ {
	q.regKey1(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetKey1_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueKey1(), "key1", option)
}

func (q *UserTableCQ) SetKey1_PrefixSearch(value string) error {
	return q.SetKey1_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *UserTableCQ) SetKey1_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueKey1(), "key1", option)
}



func (q *UserTableCQ) AddOrderBy_Key1_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("key1")
	return q
}
func (q *UserTableCQ) AddOrderBy_Key1_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("key1")
	return q
}
func (q *UserTableCQ) regKey1(key *df.ConditionKey, value interface{}) {
	if q.Key1 == nil {
		q.Key1 = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Key1, "key1")
}

func (q *UserTableCQ) getCValueKey2() *df.ConditionValue {
	if q.Key2 == nil {
		q.Key2 = new(df.ConditionValue)
	}
	return q.Key2
}


func (q *UserTableCQ) SetKey2_Equal(value string) *UserTableCQ {
	q.regKey2(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *UserTableCQ) SetKey2_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueKey2(), "key2")
}
func (q *UserTableCQ) SetKey2_NotEqual(value string) *UserTableCQ {
	q.regKey2(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetKey2_GreaterThan(value string) *UserTableCQ {
	q.regKey2(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetKey2_LessThan(value string) *UserTableCQ {
	q.regKey2(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetKey2_GreaterEqual(value string) *UserTableCQ {
	q.regKey2(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *UserTableCQ) SetKey2_LessEqual(value string) *UserTableCQ {
	q.regKey2(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetKey2_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueKey2(), "key2", option)
}

func (q *UserTableCQ) SetKey2_PrefixSearch(value string) error {
	return q.SetKey2_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *UserTableCQ) SetKey2_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueKey2(), "key2", option)
}



func (q *UserTableCQ) AddOrderBy_Key2_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("key2")
	return q
}
func (q *UserTableCQ) AddOrderBy_Key2_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("key2")
	return q
}
func (q *UserTableCQ) regKey2(key *df.ConditionKey, value interface{}) {
	if q.Key2 == nil {
		q.Key2 = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Key2, "key2")
}

func (q *UserTableCQ) getCValueS1Data() *df.ConditionValue {
	if q.S1Data == nil {
		q.S1Data = new(df.ConditionValue)
	}
	return q.S1Data
}


func (q *UserTableCQ) SetS1Data_Equal(value string) *UserTableCQ {
	q.regS1Data(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *UserTableCQ) SetS1Data_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueS1Data(), "s1Data")
}
func (q *UserTableCQ) SetS1Data_NotEqual(value string) *UserTableCQ {
	q.regS1Data(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetS1Data_GreaterThan(value string) *UserTableCQ {
	q.regS1Data(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetS1Data_LessThan(value string) *UserTableCQ {
	q.regS1Data(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetS1Data_GreaterEqual(value string) *UserTableCQ {
	q.regS1Data(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *UserTableCQ) SetS1Data_LessEqual(value string) *UserTableCQ {
	q.regS1Data(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetS1Data_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueS1Data(), "s1Data", option)
}

func (q *UserTableCQ) SetS1Data_PrefixSearch(value string) error {
	return q.SetS1Data_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *UserTableCQ) SetS1Data_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueS1Data(), "s1Data", option)
}



func (q *UserTableCQ) SetS1Data_IsNull() *UserTableCQ {
	q.regS1Data(df.CK_ISN_C, 0)
	return q
}
func (q *UserTableCQ) SetS1Data_IsNullOrEmpty() *UserTableCQ {
	q.regS1Data(df.CK_ISNOE_C, 0)
	return q
}
func (q *UserTableCQ) SetS1Data_IsNotNull() *UserTableCQ {
	q.regS1Data(df.CK_ISNN_C, 0)
	return q
}
func (q *UserTableCQ) AddOrderBy_S1Data_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("s1Data")
	return q
}
func (q *UserTableCQ) AddOrderBy_S1Data_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("s1Data")
	return q
}
func (q *UserTableCQ) regS1Data(key *df.ConditionKey, value interface{}) {
	if q.S1Data == nil {
		q.S1Data = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.S1Data, "s1Data")
}

func (q *UserTableCQ) getCValueS2Data() *df.ConditionValue {
	if q.S2Data == nil {
		q.S2Data = new(df.ConditionValue)
	}
	return q.S2Data
}


func (q *UserTableCQ) SetS2Data_Equal(value string) *UserTableCQ {
	q.regS2Data(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *UserTableCQ) SetS2Data_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueS2Data(), "s2Data")
}
func (q *UserTableCQ) SetS2Data_NotEqual(value string) *UserTableCQ {
	q.regS2Data(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetS2Data_GreaterThan(value string) *UserTableCQ {
	q.regS2Data(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetS2Data_LessThan(value string) *UserTableCQ {
	q.regS2Data(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetS2Data_GreaterEqual(value string) *UserTableCQ {
	q.regS2Data(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *UserTableCQ) SetS2Data_LessEqual(value string) *UserTableCQ {
	q.regS2Data(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetS2Data_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueS2Data(), "s2Data", option)
}

func (q *UserTableCQ) SetS2Data_PrefixSearch(value string) error {
	return q.SetS2Data_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *UserTableCQ) SetS2Data_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueS2Data(), "s2Data", option)
}



func (q *UserTableCQ) SetS2Data_IsNull() *UserTableCQ {
	q.regS2Data(df.CK_ISN_C, 0)
	return q
}
func (q *UserTableCQ) SetS2Data_IsNullOrEmpty() *UserTableCQ {
	q.regS2Data(df.CK_ISNOE_C, 0)
	return q
}
func (q *UserTableCQ) SetS2Data_IsNotNull() *UserTableCQ {
	q.regS2Data(df.CK_ISNN_C, 0)
	return q
}
func (q *UserTableCQ) AddOrderBy_S2Data_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("s2Data")
	return q
}
func (q *UserTableCQ) AddOrderBy_S2Data_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("s2Data")
	return q
}
func (q *UserTableCQ) regS2Data(key *df.ConditionKey, value interface{}) {
	if q.S2Data == nil {
		q.S2Data = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.S2Data, "s2Data")
}

func (q *UserTableCQ) getCValueS3Data() *df.ConditionValue {
	if q.S3Data == nil {
		q.S3Data = new(df.ConditionValue)
	}
	return q.S3Data
}


func (q *UserTableCQ) SetS3Data_Equal(value string) *UserTableCQ {
	q.regS3Data(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *UserTableCQ) SetS3Data_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueS3Data(), "s3Data")
}
func (q *UserTableCQ) SetS3Data_NotEqual(value string) *UserTableCQ {
	q.regS3Data(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetS3Data_GreaterThan(value string) *UserTableCQ {
	q.regS3Data(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetS3Data_LessThan(value string) *UserTableCQ {
	q.regS3Data(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetS3Data_GreaterEqual(value string) *UserTableCQ {
	q.regS3Data(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *UserTableCQ) SetS3Data_LessEqual(value string) *UserTableCQ {
	q.regS3Data(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetS3Data_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueS3Data(), "s3Data", option)
}

func (q *UserTableCQ) SetS3Data_PrefixSearch(value string) error {
	return q.SetS3Data_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *UserTableCQ) SetS3Data_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueS3Data(), "s3Data", option)
}



func (q *UserTableCQ) SetS3Data_IsNull() *UserTableCQ {
	q.regS3Data(df.CK_ISN_C, 0)
	return q
}
func (q *UserTableCQ) SetS3Data_IsNullOrEmpty() *UserTableCQ {
	q.regS3Data(df.CK_ISNOE_C, 0)
	return q
}
func (q *UserTableCQ) SetS3Data_IsNotNull() *UserTableCQ {
	q.regS3Data(df.CK_ISNN_C, 0)
	return q
}
func (q *UserTableCQ) AddOrderBy_S3Data_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("s3Data")
	return q
}
func (q *UserTableCQ) AddOrderBy_S3Data_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("s3Data")
	return q
}
func (q *UserTableCQ) regS3Data(key *df.ConditionKey, value interface{}) {
	if q.S3Data == nil {
		q.S3Data = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.S3Data, "s3Data")
}

func (q *UserTableCQ) getCValueN1Data() *df.ConditionValue {
	if q.N1Data == nil {
		q.N1Data = new(df.ConditionValue)
	}
	return q.N1Data
}



func (q *UserTableCQ) SetN1Data_Equal(value df.Numeric) *UserTableCQ {
	q.regN1Data(df.CK_EQ_C, value)
	return q
}
func (q *UserTableCQ) SetN1Data_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueN1Data(), "n1Data")
}
func (q *UserTableCQ) SetN1Data_NotEqual(value df.Numeric) *UserTableCQ {
	q.regN1Data(df.CK_NE_C, value)
	return q
}

func (q *UserTableCQ) SetN1Data_GreaterThan(value df.Numeric) *UserTableCQ {
	q.regN1Data(df.CK_GT_C, value)
	return q
}

func (q *UserTableCQ) SetN1Data_LessThan(value df.Numeric) *UserTableCQ {
	q.regN1Data(df.CK_LT_C, value)
	return q
}

func (q *UserTableCQ) SetN1Data_GreaterEqual(value df.Numeric) *UserTableCQ {
	q.regN1Data(df.CK_GE_C, value)
	return q
}

func (q *UserTableCQ) SetN1Data_LessEqual(value df.Numeric) *UserTableCQ {
	q.regN1Data(df.CK_LE_C, value)
	return q
}
func (q *UserTableCQ) SetN1Data_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueN1Data(),"n1Data",rangeOfOption)
}	


func (q *UserTableCQ) SetN1Data_IsNull() *UserTableCQ {
	q.regN1Data(df.CK_ISN_C, 0)
	return q
}
func (q *UserTableCQ) SetN1Data_IsNotNull() *UserTableCQ {
	q.regN1Data(df.CK_ISNN_C, 0)
	return q
}
func (q *UserTableCQ) AddOrderBy_N1Data_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("n1Data")
	return q
}
func (q *UserTableCQ) AddOrderBy_N1Data_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("n1Data")
	return q
}
func (q *UserTableCQ) regN1Data(key *df.ConditionKey, value interface{}) {
	if q.N1Data == nil {
		q.N1Data = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.N1Data, "n1Data")
}

func (q *UserTableCQ) getCValueN2Data() *df.ConditionValue {
	if q.N2Data == nil {
		q.N2Data = new(df.ConditionValue)
	}
	return q.N2Data
}



func (q *UserTableCQ) SetN2Data_Equal(value df.Numeric) *UserTableCQ {
	q.regN2Data(df.CK_EQ_C, value)
	return q
}
func (q *UserTableCQ) SetN2Data_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueN2Data(), "n2Data")
}
func (q *UserTableCQ) SetN2Data_NotEqual(value df.Numeric) *UserTableCQ {
	q.regN2Data(df.CK_NE_C, value)
	return q
}

func (q *UserTableCQ) SetN2Data_GreaterThan(value df.Numeric) *UserTableCQ {
	q.regN2Data(df.CK_GT_C, value)
	return q
}

func (q *UserTableCQ) SetN2Data_LessThan(value df.Numeric) *UserTableCQ {
	q.regN2Data(df.CK_LT_C, value)
	return q
}

func (q *UserTableCQ) SetN2Data_GreaterEqual(value df.Numeric) *UserTableCQ {
	q.regN2Data(df.CK_GE_C, value)
	return q
}

func (q *UserTableCQ) SetN2Data_LessEqual(value df.Numeric) *UserTableCQ {
	q.regN2Data(df.CK_LE_C, value)
	return q
}
func (q *UserTableCQ) SetN2Data_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueN2Data(),"n2Data",rangeOfOption)
}	


func (q *UserTableCQ) SetN2Data_IsNull() *UserTableCQ {
	q.regN2Data(df.CK_ISN_C, 0)
	return q
}
func (q *UserTableCQ) SetN2Data_IsNotNull() *UserTableCQ {
	q.regN2Data(df.CK_ISNN_C, 0)
	return q
}
func (q *UserTableCQ) AddOrderBy_N2Data_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("n2Data")
	return q
}
func (q *UserTableCQ) AddOrderBy_N2Data_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("n2Data")
	return q
}
func (q *UserTableCQ) regN2Data(key *df.ConditionKey, value interface{}) {
	if q.N2Data == nil {
		q.N2Data = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.N2Data, "n2Data")
}

func (q *UserTableCQ) getCValueN3Data() *df.ConditionValue {
	if q.N3Data == nil {
		q.N3Data = new(df.ConditionValue)
	}
	return q.N3Data
}



func (q *UserTableCQ) SetN3Data_Equal(value df.Numeric) *UserTableCQ {
	q.regN3Data(df.CK_EQ_C, value)
	return q
}
func (q *UserTableCQ) SetN3Data_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueN3Data(), "n3Data")
}
func (q *UserTableCQ) SetN3Data_NotEqual(value df.Numeric) *UserTableCQ {
	q.regN3Data(df.CK_NE_C, value)
	return q
}

func (q *UserTableCQ) SetN3Data_GreaterThan(value df.Numeric) *UserTableCQ {
	q.regN3Data(df.CK_GT_C, value)
	return q
}

func (q *UserTableCQ) SetN3Data_LessThan(value df.Numeric) *UserTableCQ {
	q.regN3Data(df.CK_LT_C, value)
	return q
}

func (q *UserTableCQ) SetN3Data_GreaterEqual(value df.Numeric) *UserTableCQ {
	q.regN3Data(df.CK_GE_C, value)
	return q
}

func (q *UserTableCQ) SetN3Data_LessEqual(value df.Numeric) *UserTableCQ {
	q.regN3Data(df.CK_LE_C, value)
	return q
}
func (q *UserTableCQ) SetN3Data_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueN3Data(),"n3Data",rangeOfOption)
}	


func (q *UserTableCQ) SetN3Data_IsNull() *UserTableCQ {
	q.regN3Data(df.CK_ISN_C, 0)
	return q
}
func (q *UserTableCQ) SetN3Data_IsNotNull() *UserTableCQ {
	q.regN3Data(df.CK_ISNN_C, 0)
	return q
}
func (q *UserTableCQ) AddOrderBy_N3Data_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("n3Data")
	return q
}
func (q *UserTableCQ) AddOrderBy_N3Data_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("n3Data")
	return q
}
func (q *UserTableCQ) regN3Data(key *df.ConditionKey, value interface{}) {
	if q.N3Data == nil {
		q.N3Data = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.N3Data, "n3Data")
}

func (q *UserTableCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *UserTableCQ) SetVersionNo_Equal(value int64) *UserTableCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}
func (q *UserTableCQ) SetVersionNo_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueVersionNo(), "versionNo")
}
func (q *UserTableCQ) SetVersionNo_NotEqual(value int64) *UserTableCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *UserTableCQ) SetVersionNo_GreaterThan(value int64) *UserTableCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *UserTableCQ) SetVersionNo_LessThan(value int64) *UserTableCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *UserTableCQ) SetVersionNo_GreaterEqual(value int64) *UserTableCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *UserTableCQ) SetVersionNo_LessEqual(value int64) *UserTableCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *UserTableCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *UserTableCQ) AddOrderBy_VersionNo_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *UserTableCQ) AddOrderBy_VersionNo_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *UserTableCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}

func (q *UserTableCQ) getCValueDelFlag() *df.ConditionValue {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	return q.DelFlag
}



func (q *UserTableCQ) SetDelFlag_Equal(value int64) *UserTableCQ {
	q.regDelFlag(df.CK_EQ_C, value)
	return q
}
func (q *UserTableCQ) SetDelFlag_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDelFlag(), "delFlag")
}
func (q *UserTableCQ) SetDelFlag_NotEqual(value int64) *UserTableCQ {
	q.regDelFlag(df.CK_NE_C, value)
	return q
}

func (q *UserTableCQ) SetDelFlag_GreaterThan(value int64) *UserTableCQ {
	q.regDelFlag(df.CK_GT_C, value)
	return q
}

func (q *UserTableCQ) SetDelFlag_LessThan(value int64) *UserTableCQ {
	q.regDelFlag(df.CK_LT_C, value)
	return q
}

func (q *UserTableCQ) SetDelFlag_GreaterEqual(value int64) *UserTableCQ {
	q.regDelFlag(df.CK_GE_C, value)
	return q
}

func (q *UserTableCQ) SetDelFlag_LessEqual(value int64) *UserTableCQ {
	q.regDelFlag(df.CK_LE_C, value)
	return q
}
func (q *UserTableCQ) SetDelFlag_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDelFlag(),"delFlag",rangeOfOption)
}	


func (q *UserTableCQ) AddOrderBy_DelFlag_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("delFlag")
	return q
}
func (q *UserTableCQ) AddOrderBy_DelFlag_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("delFlag")
	return q
}
func (q *UserTableCQ) regDelFlag(key *df.ConditionKey, value interface{}) {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DelFlag, "delFlag")
}

func (q *UserTableCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *UserTableCQ) SetRegisterDatetime_Equal(value df.Timestamp) *UserTableCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *UserTableCQ) SetRegisterDatetime_GreaterThan(value df.Timestamp) *UserTableCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *UserTableCQ) SetRegisterDatetime_LessThan(value df.Timestamp) *UserTableCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *UserTableCQ) SetRegisterDatetime_GreaterEqual(value df.Timestamp) *UserTableCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *UserTableCQ) SetRegisterDatetime_LessEqual(value df.Timestamp) *UserTableCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *UserTableCQ) AddOrderBy_RegisterDatetime_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *UserTableCQ) AddOrderBy_RegisterDatetime_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *UserTableCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *UserTableCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *UserTableCQ) SetRegisterUser_Equal(value string) *UserTableCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *UserTableCQ) SetRegisterUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterUser(), "registerUser")
}
func (q *UserTableCQ) SetRegisterUser_NotEqual(value string) *UserTableCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetRegisterUser_GreaterThan(value string) *UserTableCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetRegisterUser_LessThan(value string) *UserTableCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetRegisterUser_GreaterEqual(value string) *UserTableCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *UserTableCQ) SetRegisterUser_LessEqual(value string) *UserTableCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *UserTableCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *UserTableCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *UserTableCQ) AddOrderBy_RegisterUser_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *UserTableCQ) AddOrderBy_RegisterUser_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *UserTableCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *UserTableCQ) getCValueRegisterProcess() *df.ConditionValue {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	return q.RegisterProcess
}


func (q *UserTableCQ) SetRegisterProcess_Equal(value string) *UserTableCQ {
	q.regRegisterProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *UserTableCQ) SetRegisterProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterProcess(), "registerProcess")
}
func (q *UserTableCQ) SetRegisterProcess_NotEqual(value string) *UserTableCQ {
	q.regRegisterProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetRegisterProcess_GreaterThan(value string) *UserTableCQ {
	q.regRegisterProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetRegisterProcess_LessThan(value string) *UserTableCQ {
	q.regRegisterProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetRegisterProcess_GreaterEqual(value string) *UserTableCQ {
	q.regRegisterProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *UserTableCQ) SetRegisterProcess_LessEqual(value string) *UserTableCQ {
	q.regRegisterProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetRegisterProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}

func (q *UserTableCQ) SetRegisterProcess_PrefixSearch(value string) error {
	return q.SetRegisterProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *UserTableCQ) SetRegisterProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}



func (q *UserTableCQ) AddOrderBy_RegisterProcess_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("registerProcess")
	return q
}
func (q *UserTableCQ) AddOrderBy_RegisterProcess_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("registerProcess")
	return q
}
func (q *UserTableCQ) regRegisterProcess(key *df.ConditionKey, value interface{}) {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterProcess, "registerProcess")
}

func (q *UserTableCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *UserTableCQ) SetUpdateDatetime_Equal(value df.Timestamp) *UserTableCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *UserTableCQ) SetUpdateDatetime_GreaterThan(value df.Timestamp) *UserTableCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *UserTableCQ) SetUpdateDatetime_LessThan(value df.Timestamp) *UserTableCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *UserTableCQ) SetUpdateDatetime_GreaterEqual(value df.Timestamp) *UserTableCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *UserTableCQ) SetUpdateDatetime_LessEqual(value df.Timestamp) *UserTableCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *UserTableCQ) AddOrderBy_UpdateDatetime_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *UserTableCQ) AddOrderBy_UpdateDatetime_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *UserTableCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *UserTableCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *UserTableCQ) SetUpdateUser_Equal(value string) *UserTableCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *UserTableCQ) SetUpdateUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateUser(), "updateUser")
}
func (q *UserTableCQ) SetUpdateUser_NotEqual(value string) *UserTableCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetUpdateUser_GreaterThan(value string) *UserTableCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetUpdateUser_LessThan(value string) *UserTableCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetUpdateUser_GreaterEqual(value string) *UserTableCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *UserTableCQ) SetUpdateUser_LessEqual(value string) *UserTableCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *UserTableCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *UserTableCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *UserTableCQ) AddOrderBy_UpdateUser_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *UserTableCQ) AddOrderBy_UpdateUser_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *UserTableCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *UserTableCQ) getCValueUpdateProcess() *df.ConditionValue {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	return q.UpdateProcess
}


func (q *UserTableCQ) SetUpdateProcess_Equal(value string) *UserTableCQ {
	q.regUpdateProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *UserTableCQ) SetUpdateProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateProcess(), "updateProcess")
}
func (q *UserTableCQ) SetUpdateProcess_NotEqual(value string) *UserTableCQ {
	q.regUpdateProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetUpdateProcess_GreaterThan(value string) *UserTableCQ {
	q.regUpdateProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetUpdateProcess_LessThan(value string) *UserTableCQ {
	q.regUpdateProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetUpdateProcess_GreaterEqual(value string) *UserTableCQ {
	q.regUpdateProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *UserTableCQ) SetUpdateProcess_LessEqual(value string) *UserTableCQ {
	q.regUpdateProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *UserTableCQ) SetUpdateProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}

func (q *UserTableCQ) SetUpdateProcess_PrefixSearch(value string) error {
	return q.SetUpdateProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *UserTableCQ) SetUpdateProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}



func (q *UserTableCQ) AddOrderBy_UpdateProcess_Asc() *UserTableCQ {
	q.BaseConditionQuery.RegOBA("updateProcess")
	return q
}
func (q *UserTableCQ) AddOrderBy_UpdateProcess_Desc() *UserTableCQ {
	q.BaseConditionQuery.RegOBD("updateProcess")
	return q
}
func (q *UserTableCQ) regUpdateProcess(key *df.ConditionKey, value interface{}) {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateProcess, "updateProcess")
}


func CreateUserTableCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *UserTableCQ {
	cq := new(UserTableCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "UserTable"
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
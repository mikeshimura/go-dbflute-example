package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type DbfluteCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	Id *df.ConditionValue
	LoginId *df.ConditionValue
	Db *df.ConditionValue
	DbName *df.ConditionValue
	DbTable *df.ConditionValue
	Name *df.ConditionValue
	DbType *df.ConditionValue
	Required *df.ConditionValue
	JavaType *df.ConditionValue
	GoType *df.ConditionValue
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

func (q *DbfluteCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *DbfluteCQ) getCValueId() *df.ConditionValue {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	return q.Id
}



func (q *DbfluteCQ) SetId_Equal(value int64) *DbfluteCQ {
	q.regId(df.CK_EQ_C, value)
	return q
}
func (q *DbfluteCQ) SetId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueId(), "id")
}
func (q *DbfluteCQ) SetId_NotEqual(value int64) *DbfluteCQ {
	q.regId(df.CK_NE_C, value)
	return q
}

func (q *DbfluteCQ) SetId_GreaterThan(value int64) *DbfluteCQ {
	q.regId(df.CK_GT_C, value)
	return q
}

func (q *DbfluteCQ) SetId_LessThan(value int64) *DbfluteCQ {
	q.regId(df.CK_LT_C, value)
	return q
}

func (q *DbfluteCQ) SetId_GreaterEqual(value int64) *DbfluteCQ {
	q.regId(df.CK_GE_C, value)
	return q
}

func (q *DbfluteCQ) SetId_LessEqual(value int64) *DbfluteCQ {
	q.regId(df.CK_LE_C, value)
	return q
}
func (q *DbfluteCQ) SetId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueId(),"id",rangeOfOption)
}	


func (q *DbfluteCQ) SetId_IsNull() *DbfluteCQ {
	q.regId(df.CK_ISN_C, 0)
	return q
}
func (q *DbfluteCQ) SetId_IsNotNull() *DbfluteCQ {
	q.regId(df.CK_ISNN_C, 0)
	return q
}
func (q *DbfluteCQ) AddOrderBy_Id_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("id")
	return q
}
func (q *DbfluteCQ) AddOrderBy_Id_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("id")
	return q
}
func (q *DbfluteCQ) regId(key *df.ConditionKey, value interface{}) {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Id, "id")
}

func (q *DbfluteCQ) getCValueLoginId() *df.ConditionValue {
	if q.LoginId == nil {
		q.LoginId = new(df.ConditionValue)
	}
	return q.LoginId
}



func (q *DbfluteCQ) SetLoginId_Equal(value int64) *DbfluteCQ {
	q.regLoginId(df.CK_EQ_C, value)
	return q
}
func (q *DbfluteCQ) SetLoginId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueLoginId(), "loginId")
}
func (q *DbfluteCQ) SetLoginId_NotEqual(value int64) *DbfluteCQ {
	q.regLoginId(df.CK_NE_C, value)
	return q
}

func (q *DbfluteCQ) SetLoginId_GreaterThan(value int64) *DbfluteCQ {
	q.regLoginId(df.CK_GT_C, value)
	return q
}

func (q *DbfluteCQ) SetLoginId_LessThan(value int64) *DbfluteCQ {
	q.regLoginId(df.CK_LT_C, value)
	return q
}

func (q *DbfluteCQ) SetLoginId_GreaterEqual(value int64) *DbfluteCQ {
	q.regLoginId(df.CK_GE_C, value)
	return q
}

func (q *DbfluteCQ) SetLoginId_LessEqual(value int64) *DbfluteCQ {
	q.regLoginId(df.CK_LE_C, value)
	return q
}
func (q *DbfluteCQ) SetLoginId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueLoginId(),"loginId",rangeOfOption)
}	


func (q *DbfluteCQ) AddOrderBy_LoginId_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("loginId")
	return q
}
func (q *DbfluteCQ) AddOrderBy_LoginId_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("loginId")
	return q
}
func (q *DbfluteCQ) regLoginId(key *df.ConditionKey, value interface{}) {
	if q.LoginId == nil {
		q.LoginId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.LoginId, "loginId")
}

func (q *DbfluteCQ) getCValueDb() *df.ConditionValue {
	if q.Db == nil {
		q.Db = new(df.ConditionValue)
	}
	return q.Db
}


func (q *DbfluteCQ) SetDb_Equal(value string) *DbfluteCQ {
	q.regDb(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *DbfluteCQ) SetDb_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDb(), "db")
}
func (q *DbfluteCQ) SetDb_NotEqual(value string) *DbfluteCQ {
	q.regDb(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDb_GreaterThan(value string) *DbfluteCQ {
	q.regDb(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDb_LessThan(value string) *DbfluteCQ {
	q.regDb(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDb_GreaterEqual(value string) *DbfluteCQ {
	q.regDb(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *DbfluteCQ) SetDb_LessEqual(value string) *DbfluteCQ {
	q.regDb(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDb_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueDb(), "db", option)
}

func (q *DbfluteCQ) SetDb_PrefixSearch(value string) error {
	return q.SetDb_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *DbfluteCQ) SetDb_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueDb(), "db", option)
}



func (q *DbfluteCQ) AddOrderBy_Db_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("db")
	return q
}
func (q *DbfluteCQ) AddOrderBy_Db_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("db")
	return q
}
func (q *DbfluteCQ) regDb(key *df.ConditionKey, value interface{}) {
	if q.Db == nil {
		q.Db = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Db, "db")
}

func (q *DbfluteCQ) getCValueDbName() *df.ConditionValue {
	if q.DbName == nil {
		q.DbName = new(df.ConditionValue)
	}
	return q.DbName
}


func (q *DbfluteCQ) SetDbName_Equal(value string) *DbfluteCQ {
	q.regDbName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *DbfluteCQ) SetDbName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDbName(), "dbName")
}
func (q *DbfluteCQ) SetDbName_NotEqual(value string) *DbfluteCQ {
	q.regDbName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDbName_GreaterThan(value string) *DbfluteCQ {
	q.regDbName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDbName_LessThan(value string) *DbfluteCQ {
	q.regDbName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDbName_GreaterEqual(value string) *DbfluteCQ {
	q.regDbName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *DbfluteCQ) SetDbName_LessEqual(value string) *DbfluteCQ {
	q.regDbName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDbName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueDbName(), "dbName", option)
}

func (q *DbfluteCQ) SetDbName_PrefixSearch(value string) error {
	return q.SetDbName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *DbfluteCQ) SetDbName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueDbName(), "dbName", option)
}



func (q *DbfluteCQ) AddOrderBy_DbName_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("dbName")
	return q
}
func (q *DbfluteCQ) AddOrderBy_DbName_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("dbName")
	return q
}
func (q *DbfluteCQ) regDbName(key *df.ConditionKey, value interface{}) {
	if q.DbName == nil {
		q.DbName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DbName, "dbName")
}

func (q *DbfluteCQ) getCValueDbTable() *df.ConditionValue {
	if q.DbTable == nil {
		q.DbTable = new(df.ConditionValue)
	}
	return q.DbTable
}


func (q *DbfluteCQ) SetDbTable_Equal(value string) *DbfluteCQ {
	q.regDbTable(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *DbfluteCQ) SetDbTable_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDbTable(), "dbTable")
}
func (q *DbfluteCQ) SetDbTable_NotEqual(value string) *DbfluteCQ {
	q.regDbTable(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDbTable_GreaterThan(value string) *DbfluteCQ {
	q.regDbTable(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDbTable_LessThan(value string) *DbfluteCQ {
	q.regDbTable(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDbTable_GreaterEqual(value string) *DbfluteCQ {
	q.regDbTable(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *DbfluteCQ) SetDbTable_LessEqual(value string) *DbfluteCQ {
	q.regDbTable(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDbTable_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueDbTable(), "dbTable", option)
}

func (q *DbfluteCQ) SetDbTable_PrefixSearch(value string) error {
	return q.SetDbTable_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *DbfluteCQ) SetDbTable_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueDbTable(), "dbTable", option)
}



func (q *DbfluteCQ) AddOrderBy_DbTable_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("dbTable")
	return q
}
func (q *DbfluteCQ) AddOrderBy_DbTable_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("dbTable")
	return q
}
func (q *DbfluteCQ) regDbTable(key *df.ConditionKey, value interface{}) {
	if q.DbTable == nil {
		q.DbTable = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DbTable, "dbTable")
}

func (q *DbfluteCQ) getCValueName() *df.ConditionValue {
	if q.Name == nil {
		q.Name = new(df.ConditionValue)
	}
	return q.Name
}


func (q *DbfluteCQ) SetName_Equal(value string) *DbfluteCQ {
	q.regName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *DbfluteCQ) SetName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueName(), "name")
}
func (q *DbfluteCQ) SetName_NotEqual(value string) *DbfluteCQ {
	q.regName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetName_GreaterThan(value string) *DbfluteCQ {
	q.regName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetName_LessThan(value string) *DbfluteCQ {
	q.regName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetName_GreaterEqual(value string) *DbfluteCQ {
	q.regName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *DbfluteCQ) SetName_LessEqual(value string) *DbfluteCQ {
	q.regName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueName(), "name", option)
}

func (q *DbfluteCQ) SetName_PrefixSearch(value string) error {
	return q.SetName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *DbfluteCQ) SetName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueName(), "name", option)
}



func (q *DbfluteCQ) AddOrderBy_Name_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("name")
	return q
}
func (q *DbfluteCQ) AddOrderBy_Name_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("name")
	return q
}
func (q *DbfluteCQ) regName(key *df.ConditionKey, value interface{}) {
	if q.Name == nil {
		q.Name = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Name, "name")
}

func (q *DbfluteCQ) getCValueDbType() *df.ConditionValue {
	if q.DbType == nil {
		q.DbType = new(df.ConditionValue)
	}
	return q.DbType
}


func (q *DbfluteCQ) SetDbType_Equal(value string) *DbfluteCQ {
	q.regDbType(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *DbfluteCQ) SetDbType_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDbType(), "dbType")
}
func (q *DbfluteCQ) SetDbType_NotEqual(value string) *DbfluteCQ {
	q.regDbType(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDbType_GreaterThan(value string) *DbfluteCQ {
	q.regDbType(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDbType_LessThan(value string) *DbfluteCQ {
	q.regDbType(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDbType_GreaterEqual(value string) *DbfluteCQ {
	q.regDbType(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *DbfluteCQ) SetDbType_LessEqual(value string) *DbfluteCQ {
	q.regDbType(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetDbType_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueDbType(), "dbType", option)
}

func (q *DbfluteCQ) SetDbType_PrefixSearch(value string) error {
	return q.SetDbType_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *DbfluteCQ) SetDbType_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueDbType(), "dbType", option)
}



func (q *DbfluteCQ) AddOrderBy_DbType_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("dbType")
	return q
}
func (q *DbfluteCQ) AddOrderBy_DbType_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("dbType")
	return q
}
func (q *DbfluteCQ) regDbType(key *df.ConditionKey, value interface{}) {
	if q.DbType == nil {
		q.DbType = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DbType, "dbType")
}

func (q *DbfluteCQ) getCValueRequired() *df.ConditionValue {
	if q.Required == nil {
		q.Required = new(df.ConditionValue)
	}
	return q.Required
}





func (q *DbfluteCQ) AddOrderBy_Required_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("required")
	return q
}
func (q *DbfluteCQ) AddOrderBy_Required_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("required")
	return q
}
func (q *DbfluteCQ) regRequired(key *df.ConditionKey, value interface{}) {
	if q.Required == nil {
		q.Required = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Required, "required")
}

func (q *DbfluteCQ) getCValueJavaType() *df.ConditionValue {
	if q.JavaType == nil {
		q.JavaType = new(df.ConditionValue)
	}
	return q.JavaType
}


func (q *DbfluteCQ) SetJavaType_Equal(value string) *DbfluteCQ {
	q.regJavaType(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *DbfluteCQ) SetJavaType_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueJavaType(), "javaType")
}
func (q *DbfluteCQ) SetJavaType_NotEqual(value string) *DbfluteCQ {
	q.regJavaType(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetJavaType_GreaterThan(value string) *DbfluteCQ {
	q.regJavaType(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetJavaType_LessThan(value string) *DbfluteCQ {
	q.regJavaType(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetJavaType_GreaterEqual(value string) *DbfluteCQ {
	q.regJavaType(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *DbfluteCQ) SetJavaType_LessEqual(value string) *DbfluteCQ {
	q.regJavaType(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetJavaType_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueJavaType(), "javaType", option)
}

func (q *DbfluteCQ) SetJavaType_PrefixSearch(value string) error {
	return q.SetJavaType_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *DbfluteCQ) SetJavaType_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueJavaType(), "javaType", option)
}



func (q *DbfluteCQ) AddOrderBy_JavaType_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("javaType")
	return q
}
func (q *DbfluteCQ) AddOrderBy_JavaType_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("javaType")
	return q
}
func (q *DbfluteCQ) regJavaType(key *df.ConditionKey, value interface{}) {
	if q.JavaType == nil {
		q.JavaType = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.JavaType, "javaType")
}

func (q *DbfluteCQ) getCValueGoType() *df.ConditionValue {
	if q.GoType == nil {
		q.GoType = new(df.ConditionValue)
	}
	return q.GoType
}


func (q *DbfluteCQ) SetGoType_Equal(value string) *DbfluteCQ {
	q.regGoType(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *DbfluteCQ) SetGoType_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueGoType(), "goType")
}
func (q *DbfluteCQ) SetGoType_NotEqual(value string) *DbfluteCQ {
	q.regGoType(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetGoType_GreaterThan(value string) *DbfluteCQ {
	q.regGoType(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetGoType_LessThan(value string) *DbfluteCQ {
	q.regGoType(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetGoType_GreaterEqual(value string) *DbfluteCQ {
	q.regGoType(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *DbfluteCQ) SetGoType_LessEqual(value string) *DbfluteCQ {
	q.regGoType(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetGoType_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueGoType(), "goType", option)
}

func (q *DbfluteCQ) SetGoType_PrefixSearch(value string) error {
	return q.SetGoType_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *DbfluteCQ) SetGoType_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueGoType(), "goType", option)
}



func (q *DbfluteCQ) AddOrderBy_GoType_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("goType")
	return q
}
func (q *DbfluteCQ) AddOrderBy_GoType_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("goType")
	return q
}
func (q *DbfluteCQ) regGoType(key *df.ConditionKey, value interface{}) {
	if q.GoType == nil {
		q.GoType = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.GoType, "goType")
}

func (q *DbfluteCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *DbfluteCQ) SetVersionNo_Equal(value int64) *DbfluteCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}
func (q *DbfluteCQ) SetVersionNo_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueVersionNo(), "versionNo")
}
func (q *DbfluteCQ) SetVersionNo_NotEqual(value int64) *DbfluteCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *DbfluteCQ) SetVersionNo_GreaterThan(value int64) *DbfluteCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *DbfluteCQ) SetVersionNo_LessThan(value int64) *DbfluteCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *DbfluteCQ) SetVersionNo_GreaterEqual(value int64) *DbfluteCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *DbfluteCQ) SetVersionNo_LessEqual(value int64) *DbfluteCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *DbfluteCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *DbfluteCQ) AddOrderBy_VersionNo_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *DbfluteCQ) AddOrderBy_VersionNo_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *DbfluteCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}

func (q *DbfluteCQ) getCValueDelFlag() *df.ConditionValue {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	return q.DelFlag
}



func (q *DbfluteCQ) SetDelFlag_Equal(value int64) *DbfluteCQ {
	q.regDelFlag(df.CK_EQ_C, value)
	return q
}
func (q *DbfluteCQ) SetDelFlag_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDelFlag(), "delFlag")
}
func (q *DbfluteCQ) SetDelFlag_NotEqual(value int64) *DbfluteCQ {
	q.regDelFlag(df.CK_NE_C, value)
	return q
}

func (q *DbfluteCQ) SetDelFlag_GreaterThan(value int64) *DbfluteCQ {
	q.regDelFlag(df.CK_GT_C, value)
	return q
}

func (q *DbfluteCQ) SetDelFlag_LessThan(value int64) *DbfluteCQ {
	q.regDelFlag(df.CK_LT_C, value)
	return q
}

func (q *DbfluteCQ) SetDelFlag_GreaterEqual(value int64) *DbfluteCQ {
	q.regDelFlag(df.CK_GE_C, value)
	return q
}

func (q *DbfluteCQ) SetDelFlag_LessEqual(value int64) *DbfluteCQ {
	q.regDelFlag(df.CK_LE_C, value)
	return q
}
func (q *DbfluteCQ) SetDelFlag_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDelFlag(),"delFlag",rangeOfOption)
}	


func (q *DbfluteCQ) AddOrderBy_DelFlag_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("delFlag")
	return q
}
func (q *DbfluteCQ) AddOrderBy_DelFlag_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("delFlag")
	return q
}
func (q *DbfluteCQ) regDelFlag(key *df.ConditionKey, value interface{}) {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DelFlag, "delFlag")
}

func (q *DbfluteCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *DbfluteCQ) SetRegisterDatetime_Equal(value df.Timestamp) *DbfluteCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *DbfluteCQ) SetRegisterDatetime_GreaterThan(value df.Timestamp) *DbfluteCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *DbfluteCQ) SetRegisterDatetime_LessThan(value df.Timestamp) *DbfluteCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *DbfluteCQ) SetRegisterDatetime_GreaterEqual(value df.Timestamp) *DbfluteCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *DbfluteCQ) SetRegisterDatetime_LessEqual(value df.Timestamp) *DbfluteCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *DbfluteCQ) AddOrderBy_RegisterDatetime_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *DbfluteCQ) AddOrderBy_RegisterDatetime_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *DbfluteCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *DbfluteCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *DbfluteCQ) SetRegisterUser_Equal(value string) *DbfluteCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *DbfluteCQ) SetRegisterUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterUser(), "registerUser")
}
func (q *DbfluteCQ) SetRegisterUser_NotEqual(value string) *DbfluteCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetRegisterUser_GreaterThan(value string) *DbfluteCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetRegisterUser_LessThan(value string) *DbfluteCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetRegisterUser_GreaterEqual(value string) *DbfluteCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *DbfluteCQ) SetRegisterUser_LessEqual(value string) *DbfluteCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *DbfluteCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *DbfluteCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *DbfluteCQ) AddOrderBy_RegisterUser_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *DbfluteCQ) AddOrderBy_RegisterUser_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *DbfluteCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *DbfluteCQ) getCValueRegisterProcess() *df.ConditionValue {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	return q.RegisterProcess
}


func (q *DbfluteCQ) SetRegisterProcess_Equal(value string) *DbfluteCQ {
	q.regRegisterProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *DbfluteCQ) SetRegisterProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterProcess(), "registerProcess")
}
func (q *DbfluteCQ) SetRegisterProcess_NotEqual(value string) *DbfluteCQ {
	q.regRegisterProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetRegisterProcess_GreaterThan(value string) *DbfluteCQ {
	q.regRegisterProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetRegisterProcess_LessThan(value string) *DbfluteCQ {
	q.regRegisterProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetRegisterProcess_GreaterEqual(value string) *DbfluteCQ {
	q.regRegisterProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *DbfluteCQ) SetRegisterProcess_LessEqual(value string) *DbfluteCQ {
	q.regRegisterProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetRegisterProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}

func (q *DbfluteCQ) SetRegisterProcess_PrefixSearch(value string) error {
	return q.SetRegisterProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *DbfluteCQ) SetRegisterProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}



func (q *DbfluteCQ) AddOrderBy_RegisterProcess_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("registerProcess")
	return q
}
func (q *DbfluteCQ) AddOrderBy_RegisterProcess_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("registerProcess")
	return q
}
func (q *DbfluteCQ) regRegisterProcess(key *df.ConditionKey, value interface{}) {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterProcess, "registerProcess")
}

func (q *DbfluteCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *DbfluteCQ) SetUpdateDatetime_Equal(value df.Timestamp) *DbfluteCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *DbfluteCQ) SetUpdateDatetime_GreaterThan(value df.Timestamp) *DbfluteCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *DbfluteCQ) SetUpdateDatetime_LessThan(value df.Timestamp) *DbfluteCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *DbfluteCQ) SetUpdateDatetime_GreaterEqual(value df.Timestamp) *DbfluteCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *DbfluteCQ) SetUpdateDatetime_LessEqual(value df.Timestamp) *DbfluteCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *DbfluteCQ) AddOrderBy_UpdateDatetime_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *DbfluteCQ) AddOrderBy_UpdateDatetime_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *DbfluteCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *DbfluteCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *DbfluteCQ) SetUpdateUser_Equal(value string) *DbfluteCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *DbfluteCQ) SetUpdateUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateUser(), "updateUser")
}
func (q *DbfluteCQ) SetUpdateUser_NotEqual(value string) *DbfluteCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetUpdateUser_GreaterThan(value string) *DbfluteCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetUpdateUser_LessThan(value string) *DbfluteCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetUpdateUser_GreaterEqual(value string) *DbfluteCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *DbfluteCQ) SetUpdateUser_LessEqual(value string) *DbfluteCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *DbfluteCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *DbfluteCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *DbfluteCQ) AddOrderBy_UpdateUser_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *DbfluteCQ) AddOrderBy_UpdateUser_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *DbfluteCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *DbfluteCQ) getCValueUpdateProcess() *df.ConditionValue {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	return q.UpdateProcess
}


func (q *DbfluteCQ) SetUpdateProcess_Equal(value string) *DbfluteCQ {
	q.regUpdateProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *DbfluteCQ) SetUpdateProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateProcess(), "updateProcess")
}
func (q *DbfluteCQ) SetUpdateProcess_NotEqual(value string) *DbfluteCQ {
	q.regUpdateProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetUpdateProcess_GreaterThan(value string) *DbfluteCQ {
	q.regUpdateProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetUpdateProcess_LessThan(value string) *DbfluteCQ {
	q.regUpdateProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetUpdateProcess_GreaterEqual(value string) *DbfluteCQ {
	q.regUpdateProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *DbfluteCQ) SetUpdateProcess_LessEqual(value string) *DbfluteCQ {
	q.regUpdateProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *DbfluteCQ) SetUpdateProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}

func (q *DbfluteCQ) SetUpdateProcess_PrefixSearch(value string) error {
	return q.SetUpdateProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *DbfluteCQ) SetUpdateProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}



func (q *DbfluteCQ) AddOrderBy_UpdateProcess_Asc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBA("updateProcess")
	return q
}
func (q *DbfluteCQ) AddOrderBy_UpdateProcess_Desc() *DbfluteCQ {
	q.BaseConditionQuery.RegOBD("updateProcess")
	return q
}
func (q *DbfluteCQ) regUpdateProcess(key *df.ConditionKey, value interface{}) {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateProcess, "updateProcess")
}


func (q *DbfluteCQ) QueryLogin() *LoginCQ {
	if q.conditionQueryLogin == nil {
		q.conditionQueryLogin = q.xcreateQueryLogin()
		q.xsetupOuterJoinLogin()
	}
	return q.conditionQueryLogin
}

func (q *DbfluteCQ) xcreateQueryLogin() *LoginCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("Dbflute", "Login")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreateLoginCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "Login"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *DbfluteCQ) xsetupOuterJoinLogin() {
	    cq := q.QueryLogin()
        joinOnMap := make(map[string]string)
        joinOnMap["loginId"]="id"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "Login");
}	
	
func CreateDbfluteCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *DbfluteCQ {
	cq := new(DbfluteCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "Dbflute"
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

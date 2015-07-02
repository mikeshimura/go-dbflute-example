package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type CustomerCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	Id *df.ConditionValue
	CusCd *df.ConditionValue
	Name *df.ConditionValue
	Addr *df.ConditionValue
	Bldg *df.ConditionValue
	CusConSec *df.ConditionValue
	CusConName *df.ConditionValue
	Tel *df.ConditionValue
	SalesAmount *df.ConditionValue
	VersionNo *df.ConditionValue
	DelFlag *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterUser *df.ConditionValue
	RegisterProcess *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateUser *df.ConditionValue
	UpdateProcess *df.ConditionValue
}

func (q *CustomerCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *CustomerCQ) getCValueId() *df.ConditionValue {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	return q.Id
}



func (q *CustomerCQ) SetId_Equal(value int64) *CustomerCQ {
	q.regId(df.CK_EQ_C, value)
	return q
}
func (q *CustomerCQ) SetId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueId(), "id")
}
func (q *CustomerCQ) SetId_NotEqual(value int64) *CustomerCQ {
	q.regId(df.CK_NE_C, value)
	return q
}

func (q *CustomerCQ) SetId_GreaterThan(value int64) *CustomerCQ {
	q.regId(df.CK_GT_C, value)
	return q
}

func (q *CustomerCQ) SetId_LessThan(value int64) *CustomerCQ {
	q.regId(df.CK_LT_C, value)
	return q
}

func (q *CustomerCQ) SetId_GreaterEqual(value int64) *CustomerCQ {
	q.regId(df.CK_GE_C, value)
	return q
}

func (q *CustomerCQ) SetId_LessEqual(value int64) *CustomerCQ {
	q.regId(df.CK_LE_C, value)
	return q
}
func (q *CustomerCQ) SetId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueId(),"id",rangeOfOption)
}	


func (q *CustomerCQ) SetId_IsNull() *CustomerCQ {
	q.regId(df.CK_ISN_C, 0)
	return q
}
func (q *CustomerCQ) SetId_IsNotNull() *CustomerCQ {
	q.regId(df.CK_ISNN_C, 0)
	return q
}
func (q *CustomerCQ) AddOrderBy_Id_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("id")
	return q
}
func (q *CustomerCQ) AddOrderBy_Id_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("id")
	return q
}
func (q *CustomerCQ) regId(key *df.ConditionKey, value interface{}) {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Id, "id")
}

func (q *CustomerCQ) getCValueCusCd() *df.ConditionValue {
	if q.CusCd == nil {
		q.CusCd = new(df.ConditionValue)
	}
	return q.CusCd
}


func (q *CustomerCQ) SetCusCd_Equal(value string) *CustomerCQ {
	q.regCusCd(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *CustomerCQ) SetCusCd_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueCusCd(), "cusCd")
}
func (q *CustomerCQ) SetCusCd_NotEqual(value string) *CustomerCQ {
	q.regCusCd(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetCusCd_GreaterThan(value string) *CustomerCQ {
	q.regCusCd(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetCusCd_LessThan(value string) *CustomerCQ {
	q.regCusCd(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetCusCd_GreaterEqual(value string) *CustomerCQ {
	q.regCusCd(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *CustomerCQ) SetCusCd_LessEqual(value string) *CustomerCQ {
	q.regCusCd(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetCusCd_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueCusCd(), "cusCd", option)
}

func (q *CustomerCQ) SetCusCd_PrefixSearch(value string) error {
	return q.SetCusCd_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *CustomerCQ) SetCusCd_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueCusCd(), "cusCd", option)
}



func (q *CustomerCQ) AddOrderBy_CusCd_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("cusCd")
	return q
}
func (q *CustomerCQ) AddOrderBy_CusCd_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("cusCd")
	return q
}
func (q *CustomerCQ) regCusCd(key *df.ConditionKey, value interface{}) {
	if q.CusCd == nil {
		q.CusCd = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.CusCd, "cusCd")
}

func (q *CustomerCQ) getCValueName() *df.ConditionValue {
	if q.Name == nil {
		q.Name = new(df.ConditionValue)
	}
	return q.Name
}


func (q *CustomerCQ) SetName_Equal(value string) *CustomerCQ {
	q.regName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *CustomerCQ) SetName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueName(), "name")
}
func (q *CustomerCQ) SetName_NotEqual(value string) *CustomerCQ {
	q.regName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetName_GreaterThan(value string) *CustomerCQ {
	q.regName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetName_LessThan(value string) *CustomerCQ {
	q.regName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetName_GreaterEqual(value string) *CustomerCQ {
	q.regName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *CustomerCQ) SetName_LessEqual(value string) *CustomerCQ {
	q.regName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueName(), "name", option)
}

func (q *CustomerCQ) SetName_PrefixSearch(value string) error {
	return q.SetName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *CustomerCQ) SetName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueName(), "name", option)
}



func (q *CustomerCQ) AddOrderBy_Name_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("name")
	return q
}
func (q *CustomerCQ) AddOrderBy_Name_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("name")
	return q
}
func (q *CustomerCQ) regName(key *df.ConditionKey, value interface{}) {
	if q.Name == nil {
		q.Name = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Name, "name")
}

func (q *CustomerCQ) getCValueAddr() *df.ConditionValue {
	if q.Addr == nil {
		q.Addr = new(df.ConditionValue)
	}
	return q.Addr
}


func (q *CustomerCQ) SetAddr_Equal(value string) *CustomerCQ {
	q.regAddr(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *CustomerCQ) SetAddr_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueAddr(), "addr")
}
func (q *CustomerCQ) SetAddr_NotEqual(value string) *CustomerCQ {
	q.regAddr(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetAddr_GreaterThan(value string) *CustomerCQ {
	q.regAddr(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetAddr_LessThan(value string) *CustomerCQ {
	q.regAddr(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetAddr_GreaterEqual(value string) *CustomerCQ {
	q.regAddr(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *CustomerCQ) SetAddr_LessEqual(value string) *CustomerCQ {
	q.regAddr(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetAddr_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueAddr(), "addr", option)
}

func (q *CustomerCQ) SetAddr_PrefixSearch(value string) error {
	return q.SetAddr_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *CustomerCQ) SetAddr_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueAddr(), "addr", option)
}



func (q *CustomerCQ) SetAddr_IsNull() *CustomerCQ {
	q.regAddr(df.CK_ISN_C, 0)
	return q
}
func (q *CustomerCQ) SetAddr_IsNullOrEmpty() *CustomerCQ {
	q.regAddr(df.CK_ISNOE_C, 0)
	return q
}
func (q *CustomerCQ) SetAddr_IsNotNull() *CustomerCQ {
	q.regAddr(df.CK_ISNN_C, 0)
	return q
}
func (q *CustomerCQ) AddOrderBy_Addr_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("addr")
	return q
}
func (q *CustomerCQ) AddOrderBy_Addr_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("addr")
	return q
}
func (q *CustomerCQ) regAddr(key *df.ConditionKey, value interface{}) {
	if q.Addr == nil {
		q.Addr = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Addr, "addr")
}

func (q *CustomerCQ) getCValueBldg() *df.ConditionValue {
	if q.Bldg == nil {
		q.Bldg = new(df.ConditionValue)
	}
	return q.Bldg
}


func (q *CustomerCQ) SetBldg_Equal(value string) *CustomerCQ {
	q.regBldg(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *CustomerCQ) SetBldg_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueBldg(), "bldg")
}
func (q *CustomerCQ) SetBldg_NotEqual(value string) *CustomerCQ {
	q.regBldg(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetBldg_GreaterThan(value string) *CustomerCQ {
	q.regBldg(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetBldg_LessThan(value string) *CustomerCQ {
	q.regBldg(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetBldg_GreaterEqual(value string) *CustomerCQ {
	q.regBldg(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *CustomerCQ) SetBldg_LessEqual(value string) *CustomerCQ {
	q.regBldg(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetBldg_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueBldg(), "bldg", option)
}

func (q *CustomerCQ) SetBldg_PrefixSearch(value string) error {
	return q.SetBldg_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *CustomerCQ) SetBldg_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueBldg(), "bldg", option)
}



func (q *CustomerCQ) SetBldg_IsNull() *CustomerCQ {
	q.regBldg(df.CK_ISN_C, 0)
	return q
}
func (q *CustomerCQ) SetBldg_IsNullOrEmpty() *CustomerCQ {
	q.regBldg(df.CK_ISNOE_C, 0)
	return q
}
func (q *CustomerCQ) SetBldg_IsNotNull() *CustomerCQ {
	q.regBldg(df.CK_ISNN_C, 0)
	return q
}
func (q *CustomerCQ) AddOrderBy_Bldg_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("bldg")
	return q
}
func (q *CustomerCQ) AddOrderBy_Bldg_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("bldg")
	return q
}
func (q *CustomerCQ) regBldg(key *df.ConditionKey, value interface{}) {
	if q.Bldg == nil {
		q.Bldg = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Bldg, "bldg")
}

func (q *CustomerCQ) getCValueCusConSec() *df.ConditionValue {
	if q.CusConSec == nil {
		q.CusConSec = new(df.ConditionValue)
	}
	return q.CusConSec
}


func (q *CustomerCQ) SetCusConSec_Equal(value string) *CustomerCQ {
	q.regCusConSec(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *CustomerCQ) SetCusConSec_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueCusConSec(), "cusConSec")
}
func (q *CustomerCQ) SetCusConSec_NotEqual(value string) *CustomerCQ {
	q.regCusConSec(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetCusConSec_GreaterThan(value string) *CustomerCQ {
	q.regCusConSec(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetCusConSec_LessThan(value string) *CustomerCQ {
	q.regCusConSec(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetCusConSec_GreaterEqual(value string) *CustomerCQ {
	q.regCusConSec(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *CustomerCQ) SetCusConSec_LessEqual(value string) *CustomerCQ {
	q.regCusConSec(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetCusConSec_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueCusConSec(), "cusConSec", option)
}

func (q *CustomerCQ) SetCusConSec_PrefixSearch(value string) error {
	return q.SetCusConSec_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *CustomerCQ) SetCusConSec_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueCusConSec(), "cusConSec", option)
}



func (q *CustomerCQ) SetCusConSec_IsNull() *CustomerCQ {
	q.regCusConSec(df.CK_ISN_C, 0)
	return q
}
func (q *CustomerCQ) SetCusConSec_IsNullOrEmpty() *CustomerCQ {
	q.regCusConSec(df.CK_ISNOE_C, 0)
	return q
}
func (q *CustomerCQ) SetCusConSec_IsNotNull() *CustomerCQ {
	q.regCusConSec(df.CK_ISNN_C, 0)
	return q
}
func (q *CustomerCQ) AddOrderBy_CusConSec_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("cusConSec")
	return q
}
func (q *CustomerCQ) AddOrderBy_CusConSec_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("cusConSec")
	return q
}
func (q *CustomerCQ) regCusConSec(key *df.ConditionKey, value interface{}) {
	if q.CusConSec == nil {
		q.CusConSec = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.CusConSec, "cusConSec")
}

func (q *CustomerCQ) getCValueCusConName() *df.ConditionValue {
	if q.CusConName == nil {
		q.CusConName = new(df.ConditionValue)
	}
	return q.CusConName
}


func (q *CustomerCQ) SetCusConName_Equal(value string) *CustomerCQ {
	q.regCusConName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *CustomerCQ) SetCusConName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueCusConName(), "cusConName")
}
func (q *CustomerCQ) SetCusConName_NotEqual(value string) *CustomerCQ {
	q.regCusConName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetCusConName_GreaterThan(value string) *CustomerCQ {
	q.regCusConName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetCusConName_LessThan(value string) *CustomerCQ {
	q.regCusConName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetCusConName_GreaterEqual(value string) *CustomerCQ {
	q.regCusConName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *CustomerCQ) SetCusConName_LessEqual(value string) *CustomerCQ {
	q.regCusConName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetCusConName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueCusConName(), "cusConName", option)
}

func (q *CustomerCQ) SetCusConName_PrefixSearch(value string) error {
	return q.SetCusConName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *CustomerCQ) SetCusConName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueCusConName(), "cusConName", option)
}



func (q *CustomerCQ) SetCusConName_IsNull() *CustomerCQ {
	q.regCusConName(df.CK_ISN_C, 0)
	return q
}
func (q *CustomerCQ) SetCusConName_IsNullOrEmpty() *CustomerCQ {
	q.regCusConName(df.CK_ISNOE_C, 0)
	return q
}
func (q *CustomerCQ) SetCusConName_IsNotNull() *CustomerCQ {
	q.regCusConName(df.CK_ISNN_C, 0)
	return q
}
func (q *CustomerCQ) AddOrderBy_CusConName_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("cusConName")
	return q
}
func (q *CustomerCQ) AddOrderBy_CusConName_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("cusConName")
	return q
}
func (q *CustomerCQ) regCusConName(key *df.ConditionKey, value interface{}) {
	if q.CusConName == nil {
		q.CusConName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.CusConName, "cusConName")
}

func (q *CustomerCQ) getCValueTel() *df.ConditionValue {
	if q.Tel == nil {
		q.Tel = new(df.ConditionValue)
	}
	return q.Tel
}


func (q *CustomerCQ) SetTel_Equal(value string) *CustomerCQ {
	q.regTel(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *CustomerCQ) SetTel_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueTel(), "tel")
}
func (q *CustomerCQ) SetTel_NotEqual(value string) *CustomerCQ {
	q.regTel(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetTel_GreaterThan(value string) *CustomerCQ {
	q.regTel(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetTel_LessThan(value string) *CustomerCQ {
	q.regTel(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetTel_GreaterEqual(value string) *CustomerCQ {
	q.regTel(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *CustomerCQ) SetTel_LessEqual(value string) *CustomerCQ {
	q.regTel(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetTel_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueTel(), "tel", option)
}

func (q *CustomerCQ) SetTel_PrefixSearch(value string) error {
	return q.SetTel_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *CustomerCQ) SetTel_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueTel(), "tel", option)
}



func (q *CustomerCQ) SetTel_IsNull() *CustomerCQ {
	q.regTel(df.CK_ISN_C, 0)
	return q
}
func (q *CustomerCQ) SetTel_IsNullOrEmpty() *CustomerCQ {
	q.regTel(df.CK_ISNOE_C, 0)
	return q
}
func (q *CustomerCQ) SetTel_IsNotNull() *CustomerCQ {
	q.regTel(df.CK_ISNN_C, 0)
	return q
}
func (q *CustomerCQ) AddOrderBy_Tel_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("tel")
	return q
}
func (q *CustomerCQ) AddOrderBy_Tel_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("tel")
	return q
}
func (q *CustomerCQ) regTel(key *df.ConditionKey, value interface{}) {
	if q.Tel == nil {
		q.Tel = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Tel, "tel")
}

func (q *CustomerCQ) getCValueSalesAmount() *df.ConditionValue {
	if q.SalesAmount == nil {
		q.SalesAmount = new(df.ConditionValue)
	}
	return q.SalesAmount
}



func (q *CustomerCQ) SetSalesAmount_Equal(value int64) *CustomerCQ {
	q.regSalesAmount(df.CK_EQ_C, value)
	return q
}
func (q *CustomerCQ) SetSalesAmount_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueSalesAmount(), "salesAmount")
}
func (q *CustomerCQ) SetSalesAmount_NotEqual(value int64) *CustomerCQ {
	q.regSalesAmount(df.CK_NE_C, value)
	return q
}

func (q *CustomerCQ) SetSalesAmount_GreaterThan(value int64) *CustomerCQ {
	q.regSalesAmount(df.CK_GT_C, value)
	return q
}

func (q *CustomerCQ) SetSalesAmount_LessThan(value int64) *CustomerCQ {
	q.regSalesAmount(df.CK_LT_C, value)
	return q
}

func (q *CustomerCQ) SetSalesAmount_GreaterEqual(value int64) *CustomerCQ {
	q.regSalesAmount(df.CK_GE_C, value)
	return q
}

func (q *CustomerCQ) SetSalesAmount_LessEqual(value int64) *CustomerCQ {
	q.regSalesAmount(df.CK_LE_C, value)
	return q
}
func (q *CustomerCQ) SetSalesAmount_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueSalesAmount(),"salesAmount",rangeOfOption)
}	


func (q *CustomerCQ) SetSalesAmount_IsNull() *CustomerCQ {
	q.regSalesAmount(df.CK_ISN_C, 0)
	return q
}
func (q *CustomerCQ) SetSalesAmount_IsNotNull() *CustomerCQ {
	q.regSalesAmount(df.CK_ISNN_C, 0)
	return q
}
func (q *CustomerCQ) AddOrderBy_SalesAmount_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("salesAmount")
	return q
}
func (q *CustomerCQ) AddOrderBy_SalesAmount_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("salesAmount")
	return q
}
func (q *CustomerCQ) regSalesAmount(key *df.ConditionKey, value interface{}) {
	if q.SalesAmount == nil {
		q.SalesAmount = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SalesAmount, "salesAmount")
}

func (q *CustomerCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *CustomerCQ) SetVersionNo_Equal(value int64) *CustomerCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}
func (q *CustomerCQ) SetVersionNo_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueVersionNo(), "versionNo")
}
func (q *CustomerCQ) SetVersionNo_NotEqual(value int64) *CustomerCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *CustomerCQ) SetVersionNo_GreaterThan(value int64) *CustomerCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *CustomerCQ) SetVersionNo_LessThan(value int64) *CustomerCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *CustomerCQ) SetVersionNo_GreaterEqual(value int64) *CustomerCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *CustomerCQ) SetVersionNo_LessEqual(value int64) *CustomerCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *CustomerCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *CustomerCQ) AddOrderBy_VersionNo_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *CustomerCQ) AddOrderBy_VersionNo_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *CustomerCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}

func (q *CustomerCQ) getCValueDelFlag() *df.ConditionValue {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	return q.DelFlag
}



func (q *CustomerCQ) SetDelFlag_Equal(value int64) *CustomerCQ {
	q.regDelFlag(df.CK_EQ_C, value)
	return q
}
func (q *CustomerCQ) SetDelFlag_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDelFlag(), "delFlag")
}
func (q *CustomerCQ) SetDelFlag_NotEqual(value int64) *CustomerCQ {
	q.regDelFlag(df.CK_NE_C, value)
	return q
}

func (q *CustomerCQ) SetDelFlag_GreaterThan(value int64) *CustomerCQ {
	q.regDelFlag(df.CK_GT_C, value)
	return q
}

func (q *CustomerCQ) SetDelFlag_LessThan(value int64) *CustomerCQ {
	q.regDelFlag(df.CK_LT_C, value)
	return q
}

func (q *CustomerCQ) SetDelFlag_GreaterEqual(value int64) *CustomerCQ {
	q.regDelFlag(df.CK_GE_C, value)
	return q
}

func (q *CustomerCQ) SetDelFlag_LessEqual(value int64) *CustomerCQ {
	q.regDelFlag(df.CK_LE_C, value)
	return q
}
func (q *CustomerCQ) SetDelFlag_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDelFlag(),"delFlag",rangeOfOption)
}	


func (q *CustomerCQ) AddOrderBy_DelFlag_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("delFlag")
	return q
}
func (q *CustomerCQ) AddOrderBy_DelFlag_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("delFlag")
	return q
}
func (q *CustomerCQ) regDelFlag(key *df.ConditionKey, value interface{}) {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DelFlag, "delFlag")
}

func (q *CustomerCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *CustomerCQ) SetRegisterDatetime_Equal(value df.Timestamp) *CustomerCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *CustomerCQ) SetRegisterDatetime_GreaterThan(value df.Timestamp) *CustomerCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *CustomerCQ) SetRegisterDatetime_LessThan(value df.Timestamp) *CustomerCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *CustomerCQ) SetRegisterDatetime_GreaterEqual(value df.Timestamp) *CustomerCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *CustomerCQ) SetRegisterDatetime_LessEqual(value df.Timestamp) *CustomerCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *CustomerCQ) AddOrderBy_RegisterDatetime_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *CustomerCQ) AddOrderBy_RegisterDatetime_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *CustomerCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *CustomerCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *CustomerCQ) SetRegisterUser_Equal(value string) *CustomerCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *CustomerCQ) SetRegisterUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterUser(), "registerUser")
}
func (q *CustomerCQ) SetRegisterUser_NotEqual(value string) *CustomerCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetRegisterUser_GreaterThan(value string) *CustomerCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetRegisterUser_LessThan(value string) *CustomerCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetRegisterUser_GreaterEqual(value string) *CustomerCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *CustomerCQ) SetRegisterUser_LessEqual(value string) *CustomerCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *CustomerCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *CustomerCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *CustomerCQ) AddOrderBy_RegisterUser_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *CustomerCQ) AddOrderBy_RegisterUser_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *CustomerCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *CustomerCQ) getCValueRegisterProcess() *df.ConditionValue {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	return q.RegisterProcess
}


func (q *CustomerCQ) SetRegisterProcess_Equal(value string) *CustomerCQ {
	q.regRegisterProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *CustomerCQ) SetRegisterProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterProcess(), "registerProcess")
}
func (q *CustomerCQ) SetRegisterProcess_NotEqual(value string) *CustomerCQ {
	q.regRegisterProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetRegisterProcess_GreaterThan(value string) *CustomerCQ {
	q.regRegisterProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetRegisterProcess_LessThan(value string) *CustomerCQ {
	q.regRegisterProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetRegisterProcess_GreaterEqual(value string) *CustomerCQ {
	q.regRegisterProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *CustomerCQ) SetRegisterProcess_LessEqual(value string) *CustomerCQ {
	q.regRegisterProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetRegisterProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}

func (q *CustomerCQ) SetRegisterProcess_PrefixSearch(value string) error {
	return q.SetRegisterProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *CustomerCQ) SetRegisterProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}



func (q *CustomerCQ) AddOrderBy_RegisterProcess_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("registerProcess")
	return q
}
func (q *CustomerCQ) AddOrderBy_RegisterProcess_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("registerProcess")
	return q
}
func (q *CustomerCQ) regRegisterProcess(key *df.ConditionKey, value interface{}) {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterProcess, "registerProcess")
}

func (q *CustomerCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *CustomerCQ) SetUpdateDatetime_Equal(value df.Timestamp) *CustomerCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *CustomerCQ) SetUpdateDatetime_GreaterThan(value df.Timestamp) *CustomerCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *CustomerCQ) SetUpdateDatetime_LessThan(value df.Timestamp) *CustomerCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *CustomerCQ) SetUpdateDatetime_GreaterEqual(value df.Timestamp) *CustomerCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *CustomerCQ) SetUpdateDatetime_LessEqual(value df.Timestamp) *CustomerCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *CustomerCQ) AddOrderBy_UpdateDatetime_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *CustomerCQ) AddOrderBy_UpdateDatetime_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *CustomerCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *CustomerCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *CustomerCQ) SetUpdateUser_Equal(value string) *CustomerCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *CustomerCQ) SetUpdateUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateUser(), "updateUser")
}
func (q *CustomerCQ) SetUpdateUser_NotEqual(value string) *CustomerCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetUpdateUser_GreaterThan(value string) *CustomerCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetUpdateUser_LessThan(value string) *CustomerCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetUpdateUser_GreaterEqual(value string) *CustomerCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *CustomerCQ) SetUpdateUser_LessEqual(value string) *CustomerCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *CustomerCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *CustomerCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *CustomerCQ) AddOrderBy_UpdateUser_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *CustomerCQ) AddOrderBy_UpdateUser_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *CustomerCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *CustomerCQ) getCValueUpdateProcess() *df.ConditionValue {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	return q.UpdateProcess
}


func (q *CustomerCQ) SetUpdateProcess_Equal(value string) *CustomerCQ {
	q.regUpdateProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *CustomerCQ) SetUpdateProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateProcess(), "updateProcess")
}
func (q *CustomerCQ) SetUpdateProcess_NotEqual(value string) *CustomerCQ {
	q.regUpdateProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetUpdateProcess_GreaterThan(value string) *CustomerCQ {
	q.regUpdateProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetUpdateProcess_LessThan(value string) *CustomerCQ {
	q.regUpdateProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetUpdateProcess_GreaterEqual(value string) *CustomerCQ {
	q.regUpdateProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *CustomerCQ) SetUpdateProcess_LessEqual(value string) *CustomerCQ {
	q.regUpdateProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *CustomerCQ) SetUpdateProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}

func (q *CustomerCQ) SetUpdateProcess_PrefixSearch(value string) error {
	return q.SetUpdateProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *CustomerCQ) SetUpdateProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}



func (q *CustomerCQ) AddOrderBy_UpdateProcess_Asc() *CustomerCQ {
	q.BaseConditionQuery.RegOBA("updateProcess")
	return q
}
func (q *CustomerCQ) AddOrderBy_UpdateProcess_Desc() *CustomerCQ {
	q.BaseConditionQuery.RegOBD("updateProcess")
	return q
}
func (q *CustomerCQ) regUpdateProcess(key *df.ConditionKey, value interface{}) {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateProcess, "updateProcess")
}


func CreateCustomerCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *CustomerCQ {
	cq := new(CustomerCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "Customer"
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
package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type StockCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	Id *df.ConditionValue
	Yyyymm *df.ConditionValue
	PrdId *df.ConditionValue
	Name *df.ConditionValue
	Cost *df.ConditionValue
	Price *df.ConditionValue
	Unit *df.ConditionValue
	BeforeQty *df.ConditionValue
	PurQty *df.ConditionValue
	RcvQty *df.ConditionValue
	SalesQty *df.ConditionValue
	OutQty *df.ConditionValue
	StockQty *df.ConditionValue
	StockAmount *df.ConditionValue
	VersionNo *df.ConditionValue
	DelFlag *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterUser *df.ConditionValue
	RegisterProcess *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateUser *df.ConditionValue
	UpdateProcess *df.ConditionValue
    conditionQueryProduct *ProductCQ
}

func (q *StockCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *StockCQ) getCValueId() *df.ConditionValue {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	return q.Id
}



func (q *StockCQ) SetId_Equal(value int64) *StockCQ {
	q.regId(df.CK_EQ_C, value)
	return q
}
func (q *StockCQ) SetId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueId(), "id")
}
func (q *StockCQ) SetId_NotEqual(value int64) *StockCQ {
	q.regId(df.CK_NE_C, value)
	return q
}

func (q *StockCQ) SetId_GreaterThan(value int64) *StockCQ {
	q.regId(df.CK_GT_C, value)
	return q
}

func (q *StockCQ) SetId_LessThan(value int64) *StockCQ {
	q.regId(df.CK_LT_C, value)
	return q
}

func (q *StockCQ) SetId_GreaterEqual(value int64) *StockCQ {
	q.regId(df.CK_GE_C, value)
	return q
}

func (q *StockCQ) SetId_LessEqual(value int64) *StockCQ {
	q.regId(df.CK_LE_C, value)
	return q
}
func (q *StockCQ) SetId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueId(),"id",rangeOfOption)
}	


func (q *StockCQ) SetId_IsNull() *StockCQ {
	q.regId(df.CK_ISN_C, 0)
	return q
}
func (q *StockCQ) SetId_IsNotNull() *StockCQ {
	q.regId(df.CK_ISNN_C, 0)
	return q
}
func (q *StockCQ) AddOrderBy_Id_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("id")
	return q
}
func (q *StockCQ) AddOrderBy_Id_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("id")
	return q
}
func (q *StockCQ) regId(key *df.ConditionKey, value interface{}) {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Id, "id")
}

func (q *StockCQ) getCValueYyyymm() *df.ConditionValue {
	if q.Yyyymm == nil {
		q.Yyyymm = new(df.ConditionValue)
	}
	return q.Yyyymm
}


func (q *StockCQ) SetYyyymm_Equal(value string) *StockCQ {
	q.regYyyymm(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *StockCQ) SetYyyymm_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueYyyymm(), "yyyymm")
}
func (q *StockCQ) SetYyyymm_NotEqual(value string) *StockCQ {
	q.regYyyymm(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetYyyymm_GreaterThan(value string) *StockCQ {
	q.regYyyymm(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetYyyymm_LessThan(value string) *StockCQ {
	q.regYyyymm(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetYyyymm_GreaterEqual(value string) *StockCQ {
	q.regYyyymm(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *StockCQ) SetYyyymm_LessEqual(value string) *StockCQ {
	q.regYyyymm(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetYyyymm_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueYyyymm(), "yyyymm", option)
}

func (q *StockCQ) SetYyyymm_PrefixSearch(value string) error {
	return q.SetYyyymm_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *StockCQ) SetYyyymm_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueYyyymm(), "yyyymm", option)
}



func (q *StockCQ) AddOrderBy_Yyyymm_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("yyyymm")
	return q
}
func (q *StockCQ) AddOrderBy_Yyyymm_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("yyyymm")
	return q
}
func (q *StockCQ) regYyyymm(key *df.ConditionKey, value interface{}) {
	if q.Yyyymm == nil {
		q.Yyyymm = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Yyyymm, "yyyymm")
}

func (q *StockCQ) getCValuePrdId() *df.ConditionValue {
	if q.PrdId == nil {
		q.PrdId = new(df.ConditionValue)
	}
	return q.PrdId
}



func (q *StockCQ) SetPrdId_Equal(value int64) *StockCQ {
	q.regPrdId(df.CK_EQ_C, value)
	return q
}
func (q *StockCQ) SetPrdId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePrdId(), "prdId")
}
func (q *StockCQ) SetPrdId_NotEqual(value int64) *StockCQ {
	q.regPrdId(df.CK_NE_C, value)
	return q
}

func (q *StockCQ) SetPrdId_GreaterThan(value int64) *StockCQ {
	q.regPrdId(df.CK_GT_C, value)
	return q
}

func (q *StockCQ) SetPrdId_LessThan(value int64) *StockCQ {
	q.regPrdId(df.CK_LT_C, value)
	return q
}

func (q *StockCQ) SetPrdId_GreaterEqual(value int64) *StockCQ {
	q.regPrdId(df.CK_GE_C, value)
	return q
}

func (q *StockCQ) SetPrdId_LessEqual(value int64) *StockCQ {
	q.regPrdId(df.CK_LE_C, value)
	return q
}
func (q *StockCQ) SetPrdId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePrdId(),"prdId",rangeOfOption)
}	


func (q *StockCQ) AddOrderBy_PrdId_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("prdId")
	return q
}
func (q *StockCQ) AddOrderBy_PrdId_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("prdId")
	return q
}
func (q *StockCQ) regPrdId(key *df.ConditionKey, value interface{}) {
	if q.PrdId == nil {
		q.PrdId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PrdId, "prdId")
}

func (q *StockCQ) getCValueName() *df.ConditionValue {
	if q.Name == nil {
		q.Name = new(df.ConditionValue)
	}
	return q.Name
}


func (q *StockCQ) SetName_Equal(value string) *StockCQ {
	q.regName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *StockCQ) SetName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueName(), "name")
}
func (q *StockCQ) SetName_NotEqual(value string) *StockCQ {
	q.regName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetName_GreaterThan(value string) *StockCQ {
	q.regName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetName_LessThan(value string) *StockCQ {
	q.regName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetName_GreaterEqual(value string) *StockCQ {
	q.regName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *StockCQ) SetName_LessEqual(value string) *StockCQ {
	q.regName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueName(), "name", option)
}

func (q *StockCQ) SetName_PrefixSearch(value string) error {
	return q.SetName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *StockCQ) SetName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueName(), "name", option)
}



func (q *StockCQ) AddOrderBy_Name_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("name")
	return q
}
func (q *StockCQ) AddOrderBy_Name_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("name")
	return q
}
func (q *StockCQ) regName(key *df.ConditionKey, value interface{}) {
	if q.Name == nil {
		q.Name = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Name, "name")
}

func (q *StockCQ) getCValueCost() *df.ConditionValue {
	if q.Cost == nil {
		q.Cost = new(df.ConditionValue)
	}
	return q.Cost
}



func (q *StockCQ) SetCost_Equal(value int64) *StockCQ {
	q.regCost(df.CK_EQ_C, value)
	return q
}
func (q *StockCQ) SetCost_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueCost(), "cost")
}
func (q *StockCQ) SetCost_NotEqual(value int64) *StockCQ {
	q.regCost(df.CK_NE_C, value)
	return q
}

func (q *StockCQ) SetCost_GreaterThan(value int64) *StockCQ {
	q.regCost(df.CK_GT_C, value)
	return q
}

func (q *StockCQ) SetCost_LessThan(value int64) *StockCQ {
	q.regCost(df.CK_LT_C, value)
	return q
}

func (q *StockCQ) SetCost_GreaterEqual(value int64) *StockCQ {
	q.regCost(df.CK_GE_C, value)
	return q
}

func (q *StockCQ) SetCost_LessEqual(value int64) *StockCQ {
	q.regCost(df.CK_LE_C, value)
	return q
}
func (q *StockCQ) SetCost_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueCost(),"cost",rangeOfOption)
}	


func (q *StockCQ) AddOrderBy_Cost_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("cost")
	return q
}
func (q *StockCQ) AddOrderBy_Cost_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("cost")
	return q
}
func (q *StockCQ) regCost(key *df.ConditionKey, value interface{}) {
	if q.Cost == nil {
		q.Cost = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Cost, "cost")
}

func (q *StockCQ) getCValuePrice() *df.ConditionValue {
	if q.Price == nil {
		q.Price = new(df.ConditionValue)
	}
	return q.Price
}



func (q *StockCQ) SetPrice_Equal(value int64) *StockCQ {
	q.regPrice(df.CK_EQ_C, value)
	return q
}
func (q *StockCQ) SetPrice_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePrice(), "price")
}
func (q *StockCQ) SetPrice_NotEqual(value int64) *StockCQ {
	q.regPrice(df.CK_NE_C, value)
	return q
}

func (q *StockCQ) SetPrice_GreaterThan(value int64) *StockCQ {
	q.regPrice(df.CK_GT_C, value)
	return q
}

func (q *StockCQ) SetPrice_LessThan(value int64) *StockCQ {
	q.regPrice(df.CK_LT_C, value)
	return q
}

func (q *StockCQ) SetPrice_GreaterEqual(value int64) *StockCQ {
	q.regPrice(df.CK_GE_C, value)
	return q
}

func (q *StockCQ) SetPrice_LessEqual(value int64) *StockCQ {
	q.regPrice(df.CK_LE_C, value)
	return q
}
func (q *StockCQ) SetPrice_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePrice(),"price",rangeOfOption)
}	


func (q *StockCQ) AddOrderBy_Price_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("price")
	return q
}
func (q *StockCQ) AddOrderBy_Price_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("price")
	return q
}
func (q *StockCQ) regPrice(key *df.ConditionKey, value interface{}) {
	if q.Price == nil {
		q.Price = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Price, "price")
}

func (q *StockCQ) getCValueUnit() *df.ConditionValue {
	if q.Unit == nil {
		q.Unit = new(df.ConditionValue)
	}
	return q.Unit
}


func (q *StockCQ) SetUnit_Equal(value string) *StockCQ {
	q.regUnit(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *StockCQ) SetUnit_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUnit(), "unit")
}
func (q *StockCQ) SetUnit_NotEqual(value string) *StockCQ {
	q.regUnit(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetUnit_GreaterThan(value string) *StockCQ {
	q.regUnit(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetUnit_LessThan(value string) *StockCQ {
	q.regUnit(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetUnit_GreaterEqual(value string) *StockCQ {
	q.regUnit(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *StockCQ) SetUnit_LessEqual(value string) *StockCQ {
	q.regUnit(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetUnit_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUnit(), "unit", option)
}

func (q *StockCQ) SetUnit_PrefixSearch(value string) error {
	return q.SetUnit_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *StockCQ) SetUnit_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUnit(), "unit", option)
}



func (q *StockCQ) AddOrderBy_Unit_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("unit")
	return q
}
func (q *StockCQ) AddOrderBy_Unit_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("unit")
	return q
}
func (q *StockCQ) regUnit(key *df.ConditionKey, value interface{}) {
	if q.Unit == nil {
		q.Unit = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Unit, "unit")
}

func (q *StockCQ) getCValueBeforeQty() *df.ConditionValue {
	if q.BeforeQty == nil {
		q.BeforeQty = new(df.ConditionValue)
	}
	return q.BeforeQty
}



func (q *StockCQ) SetBeforeQty_Equal(value int64) *StockCQ {
	q.regBeforeQty(df.CK_EQ_C, value)
	return q
}
func (q *StockCQ) SetBeforeQty_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueBeforeQty(), "beforeQty")
}
func (q *StockCQ) SetBeforeQty_NotEqual(value int64) *StockCQ {
	q.regBeforeQty(df.CK_NE_C, value)
	return q
}

func (q *StockCQ) SetBeforeQty_GreaterThan(value int64) *StockCQ {
	q.regBeforeQty(df.CK_GT_C, value)
	return q
}

func (q *StockCQ) SetBeforeQty_LessThan(value int64) *StockCQ {
	q.regBeforeQty(df.CK_LT_C, value)
	return q
}

func (q *StockCQ) SetBeforeQty_GreaterEqual(value int64) *StockCQ {
	q.regBeforeQty(df.CK_GE_C, value)
	return q
}

func (q *StockCQ) SetBeforeQty_LessEqual(value int64) *StockCQ {
	q.regBeforeQty(df.CK_LE_C, value)
	return q
}
func (q *StockCQ) SetBeforeQty_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueBeforeQty(),"beforeQty",rangeOfOption)
}	


func (q *StockCQ) AddOrderBy_BeforeQty_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("beforeQty")
	return q
}
func (q *StockCQ) AddOrderBy_BeforeQty_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("beforeQty")
	return q
}
func (q *StockCQ) regBeforeQty(key *df.ConditionKey, value interface{}) {
	if q.BeforeQty == nil {
		q.BeforeQty = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.BeforeQty, "beforeQty")
}

func (q *StockCQ) getCValuePurQty() *df.ConditionValue {
	if q.PurQty == nil {
		q.PurQty = new(df.ConditionValue)
	}
	return q.PurQty
}



func (q *StockCQ) SetPurQty_Equal(value int64) *StockCQ {
	q.regPurQty(df.CK_EQ_C, value)
	return q
}
func (q *StockCQ) SetPurQty_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePurQty(), "purQty")
}
func (q *StockCQ) SetPurQty_NotEqual(value int64) *StockCQ {
	q.regPurQty(df.CK_NE_C, value)
	return q
}

func (q *StockCQ) SetPurQty_GreaterThan(value int64) *StockCQ {
	q.regPurQty(df.CK_GT_C, value)
	return q
}

func (q *StockCQ) SetPurQty_LessThan(value int64) *StockCQ {
	q.regPurQty(df.CK_LT_C, value)
	return q
}

func (q *StockCQ) SetPurQty_GreaterEqual(value int64) *StockCQ {
	q.regPurQty(df.CK_GE_C, value)
	return q
}

func (q *StockCQ) SetPurQty_LessEqual(value int64) *StockCQ {
	q.regPurQty(df.CK_LE_C, value)
	return q
}
func (q *StockCQ) SetPurQty_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePurQty(),"purQty",rangeOfOption)
}	


func (q *StockCQ) AddOrderBy_PurQty_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("purQty")
	return q
}
func (q *StockCQ) AddOrderBy_PurQty_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("purQty")
	return q
}
func (q *StockCQ) regPurQty(key *df.ConditionKey, value interface{}) {
	if q.PurQty == nil {
		q.PurQty = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PurQty, "purQty")
}

func (q *StockCQ) getCValueRcvQty() *df.ConditionValue {
	if q.RcvQty == nil {
		q.RcvQty = new(df.ConditionValue)
	}
	return q.RcvQty
}



func (q *StockCQ) SetRcvQty_Equal(value int64) *StockCQ {
	q.regRcvQty(df.CK_EQ_C, value)
	return q
}
func (q *StockCQ) SetRcvQty_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRcvQty(), "rcvQty")
}
func (q *StockCQ) SetRcvQty_NotEqual(value int64) *StockCQ {
	q.regRcvQty(df.CK_NE_C, value)
	return q
}

func (q *StockCQ) SetRcvQty_GreaterThan(value int64) *StockCQ {
	q.regRcvQty(df.CK_GT_C, value)
	return q
}

func (q *StockCQ) SetRcvQty_LessThan(value int64) *StockCQ {
	q.regRcvQty(df.CK_LT_C, value)
	return q
}

func (q *StockCQ) SetRcvQty_GreaterEqual(value int64) *StockCQ {
	q.regRcvQty(df.CK_GE_C, value)
	return q
}

func (q *StockCQ) SetRcvQty_LessEqual(value int64) *StockCQ {
	q.regRcvQty(df.CK_LE_C, value)
	return q
}
func (q *StockCQ) SetRcvQty_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueRcvQty(),"rcvQty",rangeOfOption)
}	


func (q *StockCQ) AddOrderBy_RcvQty_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("rcvQty")
	return q
}
func (q *StockCQ) AddOrderBy_RcvQty_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("rcvQty")
	return q
}
func (q *StockCQ) regRcvQty(key *df.ConditionKey, value interface{}) {
	if q.RcvQty == nil {
		q.RcvQty = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RcvQty, "rcvQty")
}

func (q *StockCQ) getCValueSalesQty() *df.ConditionValue {
	if q.SalesQty == nil {
		q.SalesQty = new(df.ConditionValue)
	}
	return q.SalesQty
}



func (q *StockCQ) SetSalesQty_Equal(value int64) *StockCQ {
	q.regSalesQty(df.CK_EQ_C, value)
	return q
}
func (q *StockCQ) SetSalesQty_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueSalesQty(), "salesQty")
}
func (q *StockCQ) SetSalesQty_NotEqual(value int64) *StockCQ {
	q.regSalesQty(df.CK_NE_C, value)
	return q
}

func (q *StockCQ) SetSalesQty_GreaterThan(value int64) *StockCQ {
	q.regSalesQty(df.CK_GT_C, value)
	return q
}

func (q *StockCQ) SetSalesQty_LessThan(value int64) *StockCQ {
	q.regSalesQty(df.CK_LT_C, value)
	return q
}

func (q *StockCQ) SetSalesQty_GreaterEqual(value int64) *StockCQ {
	q.regSalesQty(df.CK_GE_C, value)
	return q
}

func (q *StockCQ) SetSalesQty_LessEqual(value int64) *StockCQ {
	q.regSalesQty(df.CK_LE_C, value)
	return q
}
func (q *StockCQ) SetSalesQty_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueSalesQty(),"salesQty",rangeOfOption)
}	


func (q *StockCQ) AddOrderBy_SalesQty_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("salesQty")
	return q
}
func (q *StockCQ) AddOrderBy_SalesQty_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("salesQty")
	return q
}
func (q *StockCQ) regSalesQty(key *df.ConditionKey, value interface{}) {
	if q.SalesQty == nil {
		q.SalesQty = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SalesQty, "salesQty")
}

func (q *StockCQ) getCValueOutQty() *df.ConditionValue {
	if q.OutQty == nil {
		q.OutQty = new(df.ConditionValue)
	}
	return q.OutQty
}



func (q *StockCQ) SetOutQty_Equal(value int64) *StockCQ {
	q.regOutQty(df.CK_EQ_C, value)
	return q
}
func (q *StockCQ) SetOutQty_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueOutQty(), "outQty")
}
func (q *StockCQ) SetOutQty_NotEqual(value int64) *StockCQ {
	q.regOutQty(df.CK_NE_C, value)
	return q
}

func (q *StockCQ) SetOutQty_GreaterThan(value int64) *StockCQ {
	q.regOutQty(df.CK_GT_C, value)
	return q
}

func (q *StockCQ) SetOutQty_LessThan(value int64) *StockCQ {
	q.regOutQty(df.CK_LT_C, value)
	return q
}

func (q *StockCQ) SetOutQty_GreaterEqual(value int64) *StockCQ {
	q.regOutQty(df.CK_GE_C, value)
	return q
}

func (q *StockCQ) SetOutQty_LessEqual(value int64) *StockCQ {
	q.regOutQty(df.CK_LE_C, value)
	return q
}
func (q *StockCQ) SetOutQty_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueOutQty(),"outQty",rangeOfOption)
}	


func (q *StockCQ) AddOrderBy_OutQty_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("outQty")
	return q
}
func (q *StockCQ) AddOrderBy_OutQty_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("outQty")
	return q
}
func (q *StockCQ) regOutQty(key *df.ConditionKey, value interface{}) {
	if q.OutQty == nil {
		q.OutQty = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.OutQty, "outQty")
}

func (q *StockCQ) getCValueStockQty() *df.ConditionValue {
	if q.StockQty == nil {
		q.StockQty = new(df.ConditionValue)
	}
	return q.StockQty
}



func (q *StockCQ) SetStockQty_Equal(value int64) *StockCQ {
	q.regStockQty(df.CK_EQ_C, value)
	return q
}
func (q *StockCQ) SetStockQty_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueStockQty(), "stockQty")
}
func (q *StockCQ) SetStockQty_NotEqual(value int64) *StockCQ {
	q.regStockQty(df.CK_NE_C, value)
	return q
}

func (q *StockCQ) SetStockQty_GreaterThan(value int64) *StockCQ {
	q.regStockQty(df.CK_GT_C, value)
	return q
}

func (q *StockCQ) SetStockQty_LessThan(value int64) *StockCQ {
	q.regStockQty(df.CK_LT_C, value)
	return q
}

func (q *StockCQ) SetStockQty_GreaterEqual(value int64) *StockCQ {
	q.regStockQty(df.CK_GE_C, value)
	return q
}

func (q *StockCQ) SetStockQty_LessEqual(value int64) *StockCQ {
	q.regStockQty(df.CK_LE_C, value)
	return q
}
func (q *StockCQ) SetStockQty_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueStockQty(),"stockQty",rangeOfOption)
}	


func (q *StockCQ) AddOrderBy_StockQty_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("stockQty")
	return q
}
func (q *StockCQ) AddOrderBy_StockQty_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("stockQty")
	return q
}
func (q *StockCQ) regStockQty(key *df.ConditionKey, value interface{}) {
	if q.StockQty == nil {
		q.StockQty = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.StockQty, "stockQty")
}

func (q *StockCQ) getCValueStockAmount() *df.ConditionValue {
	if q.StockAmount == nil {
		q.StockAmount = new(df.ConditionValue)
	}
	return q.StockAmount
}



func (q *StockCQ) SetStockAmount_Equal(value int64) *StockCQ {
	q.regStockAmount(df.CK_EQ_C, value)
	return q
}
func (q *StockCQ) SetStockAmount_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueStockAmount(), "stockAmount")
}
func (q *StockCQ) SetStockAmount_NotEqual(value int64) *StockCQ {
	q.regStockAmount(df.CK_NE_C, value)
	return q
}

func (q *StockCQ) SetStockAmount_GreaterThan(value int64) *StockCQ {
	q.regStockAmount(df.CK_GT_C, value)
	return q
}

func (q *StockCQ) SetStockAmount_LessThan(value int64) *StockCQ {
	q.regStockAmount(df.CK_LT_C, value)
	return q
}

func (q *StockCQ) SetStockAmount_GreaterEqual(value int64) *StockCQ {
	q.regStockAmount(df.CK_GE_C, value)
	return q
}

func (q *StockCQ) SetStockAmount_LessEqual(value int64) *StockCQ {
	q.regStockAmount(df.CK_LE_C, value)
	return q
}
func (q *StockCQ) SetStockAmount_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueStockAmount(),"stockAmount",rangeOfOption)
}	


func (q *StockCQ) AddOrderBy_StockAmount_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("stockAmount")
	return q
}
func (q *StockCQ) AddOrderBy_StockAmount_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("stockAmount")
	return q
}
func (q *StockCQ) regStockAmount(key *df.ConditionKey, value interface{}) {
	if q.StockAmount == nil {
		q.StockAmount = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.StockAmount, "stockAmount")
}

func (q *StockCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *StockCQ) SetVersionNo_Equal(value int64) *StockCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}
func (q *StockCQ) SetVersionNo_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueVersionNo(), "versionNo")
}
func (q *StockCQ) SetVersionNo_NotEqual(value int64) *StockCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *StockCQ) SetVersionNo_GreaterThan(value int64) *StockCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *StockCQ) SetVersionNo_LessThan(value int64) *StockCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *StockCQ) SetVersionNo_GreaterEqual(value int64) *StockCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *StockCQ) SetVersionNo_LessEqual(value int64) *StockCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *StockCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *StockCQ) AddOrderBy_VersionNo_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *StockCQ) AddOrderBy_VersionNo_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *StockCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}

func (q *StockCQ) getCValueDelFlag() *df.ConditionValue {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	return q.DelFlag
}



func (q *StockCQ) SetDelFlag_Equal(value int64) *StockCQ {
	q.regDelFlag(df.CK_EQ_C, value)
	return q
}
func (q *StockCQ) SetDelFlag_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDelFlag(), "delFlag")
}
func (q *StockCQ) SetDelFlag_NotEqual(value int64) *StockCQ {
	q.regDelFlag(df.CK_NE_C, value)
	return q
}

func (q *StockCQ) SetDelFlag_GreaterThan(value int64) *StockCQ {
	q.regDelFlag(df.CK_GT_C, value)
	return q
}

func (q *StockCQ) SetDelFlag_LessThan(value int64) *StockCQ {
	q.regDelFlag(df.CK_LT_C, value)
	return q
}

func (q *StockCQ) SetDelFlag_GreaterEqual(value int64) *StockCQ {
	q.regDelFlag(df.CK_GE_C, value)
	return q
}

func (q *StockCQ) SetDelFlag_LessEqual(value int64) *StockCQ {
	q.regDelFlag(df.CK_LE_C, value)
	return q
}
func (q *StockCQ) SetDelFlag_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDelFlag(),"delFlag",rangeOfOption)
}	


func (q *StockCQ) AddOrderBy_DelFlag_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("delFlag")
	return q
}
func (q *StockCQ) AddOrderBy_DelFlag_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("delFlag")
	return q
}
func (q *StockCQ) regDelFlag(key *df.ConditionKey, value interface{}) {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DelFlag, "delFlag")
}

func (q *StockCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *StockCQ) SetRegisterDatetime_Equal(value df.Timestamp) *StockCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *StockCQ) SetRegisterDatetime_GreaterThan(value df.Timestamp) *StockCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *StockCQ) SetRegisterDatetime_LessThan(value df.Timestamp) *StockCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *StockCQ) SetRegisterDatetime_GreaterEqual(value df.Timestamp) *StockCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *StockCQ) SetRegisterDatetime_LessEqual(value df.Timestamp) *StockCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *StockCQ) AddOrderBy_RegisterDatetime_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *StockCQ) AddOrderBy_RegisterDatetime_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *StockCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *StockCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *StockCQ) SetRegisterUser_Equal(value string) *StockCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *StockCQ) SetRegisterUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterUser(), "registerUser")
}
func (q *StockCQ) SetRegisterUser_NotEqual(value string) *StockCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetRegisterUser_GreaterThan(value string) *StockCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetRegisterUser_LessThan(value string) *StockCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetRegisterUser_GreaterEqual(value string) *StockCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *StockCQ) SetRegisterUser_LessEqual(value string) *StockCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *StockCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *StockCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *StockCQ) AddOrderBy_RegisterUser_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *StockCQ) AddOrderBy_RegisterUser_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *StockCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *StockCQ) getCValueRegisterProcess() *df.ConditionValue {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	return q.RegisterProcess
}


func (q *StockCQ) SetRegisterProcess_Equal(value string) *StockCQ {
	q.regRegisterProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *StockCQ) SetRegisterProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterProcess(), "registerProcess")
}
func (q *StockCQ) SetRegisterProcess_NotEqual(value string) *StockCQ {
	q.regRegisterProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetRegisterProcess_GreaterThan(value string) *StockCQ {
	q.regRegisterProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetRegisterProcess_LessThan(value string) *StockCQ {
	q.regRegisterProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetRegisterProcess_GreaterEqual(value string) *StockCQ {
	q.regRegisterProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *StockCQ) SetRegisterProcess_LessEqual(value string) *StockCQ {
	q.regRegisterProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetRegisterProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}

func (q *StockCQ) SetRegisterProcess_PrefixSearch(value string) error {
	return q.SetRegisterProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *StockCQ) SetRegisterProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}



func (q *StockCQ) AddOrderBy_RegisterProcess_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("registerProcess")
	return q
}
func (q *StockCQ) AddOrderBy_RegisterProcess_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("registerProcess")
	return q
}
func (q *StockCQ) regRegisterProcess(key *df.ConditionKey, value interface{}) {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterProcess, "registerProcess")
}

func (q *StockCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *StockCQ) SetUpdateDatetime_Equal(value df.Timestamp) *StockCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *StockCQ) SetUpdateDatetime_GreaterThan(value df.Timestamp) *StockCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *StockCQ) SetUpdateDatetime_LessThan(value df.Timestamp) *StockCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *StockCQ) SetUpdateDatetime_GreaterEqual(value df.Timestamp) *StockCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *StockCQ) SetUpdateDatetime_LessEqual(value df.Timestamp) *StockCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *StockCQ) AddOrderBy_UpdateDatetime_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *StockCQ) AddOrderBy_UpdateDatetime_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *StockCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *StockCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *StockCQ) SetUpdateUser_Equal(value string) *StockCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *StockCQ) SetUpdateUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateUser(), "updateUser")
}
func (q *StockCQ) SetUpdateUser_NotEqual(value string) *StockCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetUpdateUser_GreaterThan(value string) *StockCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetUpdateUser_LessThan(value string) *StockCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetUpdateUser_GreaterEqual(value string) *StockCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *StockCQ) SetUpdateUser_LessEqual(value string) *StockCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *StockCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *StockCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *StockCQ) AddOrderBy_UpdateUser_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *StockCQ) AddOrderBy_UpdateUser_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *StockCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *StockCQ) getCValueUpdateProcess() *df.ConditionValue {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	return q.UpdateProcess
}


func (q *StockCQ) SetUpdateProcess_Equal(value string) *StockCQ {
	q.regUpdateProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *StockCQ) SetUpdateProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateProcess(), "updateProcess")
}
func (q *StockCQ) SetUpdateProcess_NotEqual(value string) *StockCQ {
	q.regUpdateProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetUpdateProcess_GreaterThan(value string) *StockCQ {
	q.regUpdateProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetUpdateProcess_LessThan(value string) *StockCQ {
	q.regUpdateProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetUpdateProcess_GreaterEqual(value string) *StockCQ {
	q.regUpdateProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *StockCQ) SetUpdateProcess_LessEqual(value string) *StockCQ {
	q.regUpdateProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *StockCQ) SetUpdateProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}

func (q *StockCQ) SetUpdateProcess_PrefixSearch(value string) error {
	return q.SetUpdateProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *StockCQ) SetUpdateProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}



func (q *StockCQ) AddOrderBy_UpdateProcess_Asc() *StockCQ {
	q.BaseConditionQuery.RegOBA("updateProcess")
	return q
}
func (q *StockCQ) AddOrderBy_UpdateProcess_Desc() *StockCQ {
	q.BaseConditionQuery.RegOBD("updateProcess")
	return q
}
func (q *StockCQ) regUpdateProcess(key *df.ConditionKey, value interface{}) {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateProcess, "updateProcess")
}


func (q *StockCQ) QueryProduct() *ProductCQ {
	if q.conditionQueryProduct == nil {
		q.conditionQueryProduct = q.xcreateQueryProduct()
		q.xsetupOuterJoinProduct()
	}
	return q.conditionQueryProduct
}

func (q *StockCQ) xcreateQueryProduct() *ProductCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("Stock", "Product")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreateProductCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "Product"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *StockCQ) xsetupOuterJoinProduct() {
	    cq := q.QueryProduct()
        joinOnMap := make(map[string]string)
        joinOnMap["prdId"]="id"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "Product");
}	
	
func CreateStockCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *StockCQ {
	cq := new(StockCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "Stock"
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
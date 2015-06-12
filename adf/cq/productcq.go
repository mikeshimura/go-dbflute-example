package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type ProductCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	Id *df.ConditionValue
	PrdCd *df.ConditionValue
	PrdCat *df.ConditionValue
	Name *df.ConditionValue
	Cost *df.ConditionValue
	Price *df.ConditionValue
	PerQty *df.ConditionValue
	Unit *df.ConditionValue
	PurQty *df.ConditionValue
	PurAmount *df.ConditionValue
	SalesQty *df.ConditionValue
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

func (q *ProductCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *ProductCQ) getCValueId() *df.ConditionValue {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	return q.Id
}



func (q *ProductCQ) SetId_Equal(value int64) *ProductCQ {
	q.regId(df.CK_EQ_C, value)
	return q
}
func (q *ProductCQ) SetId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueId(), "id")
}
func (q *ProductCQ) SetId_NotEqual(value int64) *ProductCQ {
	q.regId(df.CK_NE_C, value)
	return q
}

func (q *ProductCQ) SetId_GreaterThan(value int64) *ProductCQ {
	q.regId(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetId_LessThan(value int64) *ProductCQ {
	q.regId(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetId_GreaterEqual(value int64) *ProductCQ {
	q.regId(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetId_LessEqual(value int64) *ProductCQ {
	q.regId(df.CK_LE_C, value)
	return q
}
func (q *ProductCQ) SetId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueId(),"id",rangeOfOption)
}	


func (q *ProductCQ) SetId_IsNull() *ProductCQ {
	q.regId(df.CK_ISN_C, 0)
	return q
}
func (q *ProductCQ) SetId_IsNotNull() *ProductCQ {
	q.regId(df.CK_ISNN_C, 0)
	return q
}
func (q *ProductCQ) AddOrderBy_Id_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("id")
	return q
}
func (q *ProductCQ) AddOrderBy_Id_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("id")
	return q
}
func (q *ProductCQ) regId(key *df.ConditionKey, value interface{}) {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Id, "id")
}

func (q *ProductCQ) getCValuePrdCd() *df.ConditionValue {
	if q.PrdCd == nil {
		q.PrdCd = new(df.ConditionValue)
	}
	return q.PrdCd
}


func (q *ProductCQ) SetPrdCd_Equal(value string) *ProductCQ {
	q.regPrdCd(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *ProductCQ) SetPrdCd_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePrdCd(), "prdCd")
}
func (q *ProductCQ) SetPrdCd_NotEqual(value string) *ProductCQ {
	q.regPrdCd(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetPrdCd_GreaterThan(value string) *ProductCQ {
	q.regPrdCd(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetPrdCd_LessThan(value string) *ProductCQ {
	q.regPrdCd(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetPrdCd_GreaterEqual(value string) *ProductCQ {
	q.regPrdCd(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCQ) SetPrdCd_LessEqual(value string) *ProductCQ {
	q.regPrdCd(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetPrdCd_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValuePrdCd(), "prdCd", option)
}

func (q *ProductCQ) SetPrdCd_PrefixSearch(value string) error {
	return q.SetPrdCd_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCQ) SetPrdCd_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValuePrdCd(), "prdCd", option)
}



func (q *ProductCQ) AddOrderBy_PrdCd_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("prdCd")
	return q
}
func (q *ProductCQ) AddOrderBy_PrdCd_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("prdCd")
	return q
}
func (q *ProductCQ) regPrdCd(key *df.ConditionKey, value interface{}) {
	if q.PrdCd == nil {
		q.PrdCd = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PrdCd, "prdCd")
}

func (q *ProductCQ) getCValuePrdCat() *df.ConditionValue {
	if q.PrdCat == nil {
		q.PrdCat = new(df.ConditionValue)
	}
	return q.PrdCat
}


func (q *ProductCQ) SetPrdCat_Equal(value string) *ProductCQ {
	q.regPrdCat(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *ProductCQ) SetPrdCat_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePrdCat(), "prdCat")
}
func (q *ProductCQ) SetPrdCat_NotEqual(value string) *ProductCQ {
	q.regPrdCat(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetPrdCat_GreaterThan(value string) *ProductCQ {
	q.regPrdCat(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetPrdCat_LessThan(value string) *ProductCQ {
	q.regPrdCat(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetPrdCat_GreaterEqual(value string) *ProductCQ {
	q.regPrdCat(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCQ) SetPrdCat_LessEqual(value string) *ProductCQ {
	q.regPrdCat(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetPrdCat_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValuePrdCat(), "prdCat", option)
}

func (q *ProductCQ) SetPrdCat_PrefixSearch(value string) error {
	return q.SetPrdCat_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCQ) SetPrdCat_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValuePrdCat(), "prdCat", option)
}



func (q *ProductCQ) AddOrderBy_PrdCat_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("prdCat")
	return q
}
func (q *ProductCQ) AddOrderBy_PrdCat_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("prdCat")
	return q
}
func (q *ProductCQ) regPrdCat(key *df.ConditionKey, value interface{}) {
	if q.PrdCat == nil {
		q.PrdCat = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PrdCat, "prdCat")
}

func (q *ProductCQ) getCValueName() *df.ConditionValue {
	if q.Name == nil {
		q.Name = new(df.ConditionValue)
	}
	return q.Name
}


func (q *ProductCQ) SetName_Equal(value string) *ProductCQ {
	q.regName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *ProductCQ) SetName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueName(), "name")
}
func (q *ProductCQ) SetName_NotEqual(value string) *ProductCQ {
	q.regName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetName_GreaterThan(value string) *ProductCQ {
	q.regName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetName_LessThan(value string) *ProductCQ {
	q.regName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetName_GreaterEqual(value string) *ProductCQ {
	q.regName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCQ) SetName_LessEqual(value string) *ProductCQ {
	q.regName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueName(), "name", option)
}

func (q *ProductCQ) SetName_PrefixSearch(value string) error {
	return q.SetName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCQ) SetName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueName(), "name", option)
}



func (q *ProductCQ) AddOrderBy_Name_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("name")
	return q
}
func (q *ProductCQ) AddOrderBy_Name_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("name")
	return q
}
func (q *ProductCQ) regName(key *df.ConditionKey, value interface{}) {
	if q.Name == nil {
		q.Name = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Name, "name")
}

func (q *ProductCQ) getCValueCost() *df.ConditionValue {
	if q.Cost == nil {
		q.Cost = new(df.ConditionValue)
	}
	return q.Cost
}



func (q *ProductCQ) SetCost_Equal(value int64) *ProductCQ {
	q.regCost(df.CK_EQ_C, value)
	return q
}
func (q *ProductCQ) SetCost_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueCost(), "cost")
}
func (q *ProductCQ) SetCost_NotEqual(value int64) *ProductCQ {
	q.regCost(df.CK_NE_C, value)
	return q
}

func (q *ProductCQ) SetCost_GreaterThan(value int64) *ProductCQ {
	q.regCost(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetCost_LessThan(value int64) *ProductCQ {
	q.regCost(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetCost_GreaterEqual(value int64) *ProductCQ {
	q.regCost(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetCost_LessEqual(value int64) *ProductCQ {
	q.regCost(df.CK_LE_C, value)
	return q
}
func (q *ProductCQ) SetCost_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueCost(),"cost",rangeOfOption)
}	


func (q *ProductCQ) AddOrderBy_Cost_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("cost")
	return q
}
func (q *ProductCQ) AddOrderBy_Cost_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("cost")
	return q
}
func (q *ProductCQ) regCost(key *df.ConditionKey, value interface{}) {
	if q.Cost == nil {
		q.Cost = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Cost, "cost")
}

func (q *ProductCQ) getCValuePrice() *df.ConditionValue {
	if q.Price == nil {
		q.Price = new(df.ConditionValue)
	}
	return q.Price
}



func (q *ProductCQ) SetPrice_Equal(value int64) *ProductCQ {
	q.regPrice(df.CK_EQ_C, value)
	return q
}
func (q *ProductCQ) SetPrice_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePrice(), "price")
}
func (q *ProductCQ) SetPrice_NotEqual(value int64) *ProductCQ {
	q.regPrice(df.CK_NE_C, value)
	return q
}

func (q *ProductCQ) SetPrice_GreaterThan(value int64) *ProductCQ {
	q.regPrice(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetPrice_LessThan(value int64) *ProductCQ {
	q.regPrice(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetPrice_GreaterEqual(value int64) *ProductCQ {
	q.regPrice(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetPrice_LessEqual(value int64) *ProductCQ {
	q.regPrice(df.CK_LE_C, value)
	return q
}
func (q *ProductCQ) SetPrice_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePrice(),"price",rangeOfOption)
}	


func (q *ProductCQ) AddOrderBy_Price_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("price")
	return q
}
func (q *ProductCQ) AddOrderBy_Price_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("price")
	return q
}
func (q *ProductCQ) regPrice(key *df.ConditionKey, value interface{}) {
	if q.Price == nil {
		q.Price = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Price, "price")
}

func (q *ProductCQ) getCValuePerQty() *df.ConditionValue {
	if q.PerQty == nil {
		q.PerQty = new(df.ConditionValue)
	}
	return q.PerQty
}



func (q *ProductCQ) SetPerQty_Equal(value int64) *ProductCQ {
	q.regPerQty(df.CK_EQ_C, value)
	return q
}
func (q *ProductCQ) SetPerQty_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePerQty(), "perQty")
}
func (q *ProductCQ) SetPerQty_NotEqual(value int64) *ProductCQ {
	q.regPerQty(df.CK_NE_C, value)
	return q
}

func (q *ProductCQ) SetPerQty_GreaterThan(value int64) *ProductCQ {
	q.regPerQty(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetPerQty_LessThan(value int64) *ProductCQ {
	q.regPerQty(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetPerQty_GreaterEqual(value int64) *ProductCQ {
	q.regPerQty(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetPerQty_LessEqual(value int64) *ProductCQ {
	q.regPerQty(df.CK_LE_C, value)
	return q
}
func (q *ProductCQ) SetPerQty_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePerQty(),"perQty",rangeOfOption)
}	


func (q *ProductCQ) AddOrderBy_PerQty_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("perQty")
	return q
}
func (q *ProductCQ) AddOrderBy_PerQty_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("perQty")
	return q
}
func (q *ProductCQ) regPerQty(key *df.ConditionKey, value interface{}) {
	if q.PerQty == nil {
		q.PerQty = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PerQty, "perQty")
}

func (q *ProductCQ) getCValueUnit() *df.ConditionValue {
	if q.Unit == nil {
		q.Unit = new(df.ConditionValue)
	}
	return q.Unit
}


func (q *ProductCQ) SetUnit_Equal(value string) *ProductCQ {
	q.regUnit(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *ProductCQ) SetUnit_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUnit(), "unit")
}
func (q *ProductCQ) SetUnit_NotEqual(value string) *ProductCQ {
	q.regUnit(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUnit_GreaterThan(value string) *ProductCQ {
	q.regUnit(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUnit_LessThan(value string) *ProductCQ {
	q.regUnit(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUnit_GreaterEqual(value string) *ProductCQ {
	q.regUnit(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCQ) SetUnit_LessEqual(value string) *ProductCQ {
	q.regUnit(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUnit_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUnit(), "unit", option)
}

func (q *ProductCQ) SetUnit_PrefixSearch(value string) error {
	return q.SetUnit_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCQ) SetUnit_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUnit(), "unit", option)
}



func (q *ProductCQ) AddOrderBy_Unit_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("unit")
	return q
}
func (q *ProductCQ) AddOrderBy_Unit_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("unit")
	return q
}
func (q *ProductCQ) regUnit(key *df.ConditionKey, value interface{}) {
	if q.Unit == nil {
		q.Unit = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Unit, "unit")
}

func (q *ProductCQ) getCValuePurQty() *df.ConditionValue {
	if q.PurQty == nil {
		q.PurQty = new(df.ConditionValue)
	}
	return q.PurQty
}



func (q *ProductCQ) SetPurQty_Equal(value int64) *ProductCQ {
	q.regPurQty(df.CK_EQ_C, value)
	return q
}
func (q *ProductCQ) SetPurQty_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePurQty(), "purQty")
}
func (q *ProductCQ) SetPurQty_NotEqual(value int64) *ProductCQ {
	q.regPurQty(df.CK_NE_C, value)
	return q
}

func (q *ProductCQ) SetPurQty_GreaterThan(value int64) *ProductCQ {
	q.regPurQty(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetPurQty_LessThan(value int64) *ProductCQ {
	q.regPurQty(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetPurQty_GreaterEqual(value int64) *ProductCQ {
	q.regPurQty(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetPurQty_LessEqual(value int64) *ProductCQ {
	q.regPurQty(df.CK_LE_C, value)
	return q
}
func (q *ProductCQ) SetPurQty_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePurQty(),"purQty",rangeOfOption)
}	


func (q *ProductCQ) AddOrderBy_PurQty_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("purQty")
	return q
}
func (q *ProductCQ) AddOrderBy_PurQty_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("purQty")
	return q
}
func (q *ProductCQ) regPurQty(key *df.ConditionKey, value interface{}) {
	if q.PurQty == nil {
		q.PurQty = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PurQty, "purQty")
}

func (q *ProductCQ) getCValuePurAmount() *df.ConditionValue {
	if q.PurAmount == nil {
		q.PurAmount = new(df.ConditionValue)
	}
	return q.PurAmount
}



func (q *ProductCQ) SetPurAmount_Equal(value int64) *ProductCQ {
	q.regPurAmount(df.CK_EQ_C, value)
	return q
}
func (q *ProductCQ) SetPurAmount_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePurAmount(), "purAmount")
}
func (q *ProductCQ) SetPurAmount_NotEqual(value int64) *ProductCQ {
	q.regPurAmount(df.CK_NE_C, value)
	return q
}

func (q *ProductCQ) SetPurAmount_GreaterThan(value int64) *ProductCQ {
	q.regPurAmount(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetPurAmount_LessThan(value int64) *ProductCQ {
	q.regPurAmount(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetPurAmount_GreaterEqual(value int64) *ProductCQ {
	q.regPurAmount(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetPurAmount_LessEqual(value int64) *ProductCQ {
	q.regPurAmount(df.CK_LE_C, value)
	return q
}
func (q *ProductCQ) SetPurAmount_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePurAmount(),"purAmount",rangeOfOption)
}	


func (q *ProductCQ) AddOrderBy_PurAmount_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("purAmount")
	return q
}
func (q *ProductCQ) AddOrderBy_PurAmount_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("purAmount")
	return q
}
func (q *ProductCQ) regPurAmount(key *df.ConditionKey, value interface{}) {
	if q.PurAmount == nil {
		q.PurAmount = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PurAmount, "purAmount")
}

func (q *ProductCQ) getCValueSalesQty() *df.ConditionValue {
	if q.SalesQty == nil {
		q.SalesQty = new(df.ConditionValue)
	}
	return q.SalesQty
}



func (q *ProductCQ) SetSalesQty_Equal(value int64) *ProductCQ {
	q.regSalesQty(df.CK_EQ_C, value)
	return q
}
func (q *ProductCQ) SetSalesQty_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueSalesQty(), "salesQty")
}
func (q *ProductCQ) SetSalesQty_NotEqual(value int64) *ProductCQ {
	q.regSalesQty(df.CK_NE_C, value)
	return q
}

func (q *ProductCQ) SetSalesQty_GreaterThan(value int64) *ProductCQ {
	q.regSalesQty(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetSalesQty_LessThan(value int64) *ProductCQ {
	q.regSalesQty(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetSalesQty_GreaterEqual(value int64) *ProductCQ {
	q.regSalesQty(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetSalesQty_LessEqual(value int64) *ProductCQ {
	q.regSalesQty(df.CK_LE_C, value)
	return q
}
func (q *ProductCQ) SetSalesQty_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueSalesQty(),"salesQty",rangeOfOption)
}	


func (q *ProductCQ) AddOrderBy_SalesQty_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("salesQty")
	return q
}
func (q *ProductCQ) AddOrderBy_SalesQty_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("salesQty")
	return q
}
func (q *ProductCQ) regSalesQty(key *df.ConditionKey, value interface{}) {
	if q.SalesQty == nil {
		q.SalesQty = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SalesQty, "salesQty")
}

func (q *ProductCQ) getCValueSalesAmount() *df.ConditionValue {
	if q.SalesAmount == nil {
		q.SalesAmount = new(df.ConditionValue)
	}
	return q.SalesAmount
}



func (q *ProductCQ) SetSalesAmount_Equal(value int64) *ProductCQ {
	q.regSalesAmount(df.CK_EQ_C, value)
	return q
}
func (q *ProductCQ) SetSalesAmount_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueSalesAmount(), "salesAmount")
}
func (q *ProductCQ) SetSalesAmount_NotEqual(value int64) *ProductCQ {
	q.regSalesAmount(df.CK_NE_C, value)
	return q
}

func (q *ProductCQ) SetSalesAmount_GreaterThan(value int64) *ProductCQ {
	q.regSalesAmount(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetSalesAmount_LessThan(value int64) *ProductCQ {
	q.regSalesAmount(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetSalesAmount_GreaterEqual(value int64) *ProductCQ {
	q.regSalesAmount(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetSalesAmount_LessEqual(value int64) *ProductCQ {
	q.regSalesAmount(df.CK_LE_C, value)
	return q
}
func (q *ProductCQ) SetSalesAmount_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueSalesAmount(),"salesAmount",rangeOfOption)
}	


func (q *ProductCQ) AddOrderBy_SalesAmount_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("salesAmount")
	return q
}
func (q *ProductCQ) AddOrderBy_SalesAmount_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("salesAmount")
	return q
}
func (q *ProductCQ) regSalesAmount(key *df.ConditionKey, value interface{}) {
	if q.SalesAmount == nil {
		q.SalesAmount = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SalesAmount, "salesAmount")
}

func (q *ProductCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *ProductCQ) SetVersionNo_Equal(value int64) *ProductCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}
func (q *ProductCQ) SetVersionNo_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueVersionNo(), "versionNo")
}
func (q *ProductCQ) SetVersionNo_NotEqual(value int64) *ProductCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *ProductCQ) SetVersionNo_GreaterThan(value int64) *ProductCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetVersionNo_LessThan(value int64) *ProductCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetVersionNo_GreaterEqual(value int64) *ProductCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetVersionNo_LessEqual(value int64) *ProductCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *ProductCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *ProductCQ) AddOrderBy_VersionNo_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *ProductCQ) AddOrderBy_VersionNo_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *ProductCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}

func (q *ProductCQ) getCValueDelFlag() *df.ConditionValue {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	return q.DelFlag
}



func (q *ProductCQ) SetDelFlag_Equal(value int64) *ProductCQ {
	q.regDelFlag(df.CK_EQ_C, value)
	return q
}
func (q *ProductCQ) SetDelFlag_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDelFlag(), "delFlag")
}
func (q *ProductCQ) SetDelFlag_NotEqual(value int64) *ProductCQ {
	q.regDelFlag(df.CK_NE_C, value)
	return q
}

func (q *ProductCQ) SetDelFlag_GreaterThan(value int64) *ProductCQ {
	q.regDelFlag(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetDelFlag_LessThan(value int64) *ProductCQ {
	q.regDelFlag(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetDelFlag_GreaterEqual(value int64) *ProductCQ {
	q.regDelFlag(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetDelFlag_LessEqual(value int64) *ProductCQ {
	q.regDelFlag(df.CK_LE_C, value)
	return q
}
func (q *ProductCQ) SetDelFlag_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDelFlag(),"delFlag",rangeOfOption)
}	


func (q *ProductCQ) AddOrderBy_DelFlag_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("delFlag")
	return q
}
func (q *ProductCQ) AddOrderBy_DelFlag_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("delFlag")
	return q
}
func (q *ProductCQ) regDelFlag(key *df.ConditionKey, value interface{}) {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DelFlag, "delFlag")
}

func (q *ProductCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *ProductCQ) SetRegisterDatetime_Equal(value df.Timestamp) *ProductCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *ProductCQ) SetRegisterDatetime_GreaterThan(value df.Timestamp) *ProductCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetRegisterDatetime_LessThan(value df.Timestamp) *ProductCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetRegisterDatetime_GreaterEqual(value df.Timestamp) *ProductCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetRegisterDatetime_LessEqual(value df.Timestamp) *ProductCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *ProductCQ) AddOrderBy_RegisterDatetime_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *ProductCQ) AddOrderBy_RegisterDatetime_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *ProductCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *ProductCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *ProductCQ) SetRegisterUser_Equal(value string) *ProductCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *ProductCQ) SetRegisterUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterUser(), "registerUser")
}
func (q *ProductCQ) SetRegisterUser_NotEqual(value string) *ProductCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetRegisterUser_GreaterThan(value string) *ProductCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetRegisterUser_LessThan(value string) *ProductCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetRegisterUser_GreaterEqual(value string) *ProductCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCQ) SetRegisterUser_LessEqual(value string) *ProductCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *ProductCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *ProductCQ) AddOrderBy_RegisterUser_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *ProductCQ) AddOrderBy_RegisterUser_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *ProductCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *ProductCQ) getCValueRegisterProcess() *df.ConditionValue {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	return q.RegisterProcess
}


func (q *ProductCQ) SetRegisterProcess_Equal(value string) *ProductCQ {
	q.regRegisterProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *ProductCQ) SetRegisterProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterProcess(), "registerProcess")
}
func (q *ProductCQ) SetRegisterProcess_NotEqual(value string) *ProductCQ {
	q.regRegisterProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetRegisterProcess_GreaterThan(value string) *ProductCQ {
	q.regRegisterProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetRegisterProcess_LessThan(value string) *ProductCQ {
	q.regRegisterProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetRegisterProcess_GreaterEqual(value string) *ProductCQ {
	q.regRegisterProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCQ) SetRegisterProcess_LessEqual(value string) *ProductCQ {
	q.regRegisterProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetRegisterProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}

func (q *ProductCQ) SetRegisterProcess_PrefixSearch(value string) error {
	return q.SetRegisterProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCQ) SetRegisterProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}



func (q *ProductCQ) AddOrderBy_RegisterProcess_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("registerProcess")
	return q
}
func (q *ProductCQ) AddOrderBy_RegisterProcess_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("registerProcess")
	return q
}
func (q *ProductCQ) regRegisterProcess(key *df.ConditionKey, value interface{}) {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterProcess, "registerProcess")
}

func (q *ProductCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *ProductCQ) SetUpdateDatetime_Equal(value df.Timestamp) *ProductCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *ProductCQ) SetUpdateDatetime_GreaterThan(value df.Timestamp) *ProductCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetUpdateDatetime_LessThan(value df.Timestamp) *ProductCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetUpdateDatetime_GreaterEqual(value df.Timestamp) *ProductCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetUpdateDatetime_LessEqual(value df.Timestamp) *ProductCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *ProductCQ) AddOrderBy_UpdateDatetime_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *ProductCQ) AddOrderBy_UpdateDatetime_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *ProductCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *ProductCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *ProductCQ) SetUpdateUser_Equal(value string) *ProductCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *ProductCQ) SetUpdateUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateUser(), "updateUser")
}
func (q *ProductCQ) SetUpdateUser_NotEqual(value string) *ProductCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUpdateUser_GreaterThan(value string) *ProductCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUpdateUser_LessThan(value string) *ProductCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUpdateUser_GreaterEqual(value string) *ProductCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCQ) SetUpdateUser_LessEqual(value string) *ProductCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *ProductCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *ProductCQ) AddOrderBy_UpdateUser_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *ProductCQ) AddOrderBy_UpdateUser_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *ProductCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *ProductCQ) getCValueUpdateProcess() *df.ConditionValue {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	return q.UpdateProcess
}


func (q *ProductCQ) SetUpdateProcess_Equal(value string) *ProductCQ {
	q.regUpdateProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *ProductCQ) SetUpdateProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateProcess(), "updateProcess")
}
func (q *ProductCQ) SetUpdateProcess_NotEqual(value string) *ProductCQ {
	q.regUpdateProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUpdateProcess_GreaterThan(value string) *ProductCQ {
	q.regUpdateProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUpdateProcess_LessThan(value string) *ProductCQ {
	q.regUpdateProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUpdateProcess_GreaterEqual(value string) *ProductCQ {
	q.regUpdateProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCQ) SetUpdateProcess_LessEqual(value string) *ProductCQ {
	q.regUpdateProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUpdateProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}

func (q *ProductCQ) SetUpdateProcess_PrefixSearch(value string) error {
	return q.SetUpdateProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCQ) SetUpdateProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}



func (q *ProductCQ) AddOrderBy_UpdateProcess_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("updateProcess")
	return q
}
func (q *ProductCQ) AddOrderBy_UpdateProcess_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("updateProcess")
	return q
}
func (q *ProductCQ) regUpdateProcess(key *df.ConditionKey, value interface{}) {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateProcess, "updateProcess")
}


func CreateProductCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *ProductCQ {
	cq := new(ProductCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "Product"
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
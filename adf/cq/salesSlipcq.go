package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type SalesSlipCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	Id *df.ConditionValue
	SlipNo *df.ConditionValue
	Line *df.ConditionValue
	SalesDate *df.ConditionValue
	SalesId *df.ConditionValue
	CusId *df.ConditionValue
	PrdId *df.ConditionValue
	Qty *df.ConditionValue
	Unit *df.ConditionValue
	UnitPrice *df.ConditionValue
	SalesAmount *df.ConditionValue
	SlipAmount *df.ConditionValue
	SlipCons *df.ConditionValue
	VersionNo *df.ConditionValue
	DelFlag *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterUser *df.ConditionValue
	RegisterProcess *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateUser *df.ConditionValue
	UpdateProcess *df.ConditionValue
    conditionQueryCustomer *CustomerCQ
    conditionQueryProduct *ProductCQ
    conditionQueryEmployee *EmployeeCQ
}

func (q *SalesSlipCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *SalesSlipCQ) getCValueId() *df.ConditionValue {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	return q.Id
}



func (q *SalesSlipCQ) SetId_Equal(value int64) *SalesSlipCQ {
	q.regId(df.CK_EQ_C, value)
	return q
}
func (q *SalesSlipCQ) SetId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueId(), "id")
}
func (q *SalesSlipCQ) SetId_NotEqual(value int64) *SalesSlipCQ {
	q.regId(df.CK_NE_C, value)
	return q
}

func (q *SalesSlipCQ) SetId_GreaterThan(value int64) *SalesSlipCQ {
	q.regId(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetId_LessThan(value int64) *SalesSlipCQ {
	q.regId(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetId_GreaterEqual(value int64) *SalesSlipCQ {
	q.regId(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetId_LessEqual(value int64) *SalesSlipCQ {
	q.regId(df.CK_LE_C, value)
	return q
}
func (q *SalesSlipCQ) SetId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueId(),"id",rangeOfOption)
}	


func (q *SalesSlipCQ) SetId_IsNull() *SalesSlipCQ {
	q.regId(df.CK_ISN_C, 0)
	return q
}
func (q *SalesSlipCQ) SetId_IsNotNull() *SalesSlipCQ {
	q.regId(df.CK_ISNN_C, 0)
	return q
}
func (q *SalesSlipCQ) AddOrderBy_Id_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("id")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_Id_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("id")
	return q
}
func (q *SalesSlipCQ) regId(key *df.ConditionKey, value interface{}) {
	if q.Id == nil {
		q.Id = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Id, "id")
}

func (q *SalesSlipCQ) getCValueSlipNo() *df.ConditionValue {
	if q.SlipNo == nil {
		q.SlipNo = new(df.ConditionValue)
	}
	return q.SlipNo
}



func (q *SalesSlipCQ) SetSlipNo_Equal(value int64) *SalesSlipCQ {
	q.regSlipNo(df.CK_EQ_C, value)
	return q
}
func (q *SalesSlipCQ) SetSlipNo_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueSlipNo(), "slipNo")
}
func (q *SalesSlipCQ) SetSlipNo_NotEqual(value int64) *SalesSlipCQ {
	q.regSlipNo(df.CK_NE_C, value)
	return q
}

func (q *SalesSlipCQ) SetSlipNo_GreaterThan(value int64) *SalesSlipCQ {
	q.regSlipNo(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetSlipNo_LessThan(value int64) *SalesSlipCQ {
	q.regSlipNo(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetSlipNo_GreaterEqual(value int64) *SalesSlipCQ {
	q.regSlipNo(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetSlipNo_LessEqual(value int64) *SalesSlipCQ {
	q.regSlipNo(df.CK_LE_C, value)
	return q
}
func (q *SalesSlipCQ) SetSlipNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueSlipNo(),"slipNo",rangeOfOption)
}	


func (q *SalesSlipCQ) AddOrderBy_SlipNo_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("slipNo")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_SlipNo_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("slipNo")
	return q
}
func (q *SalesSlipCQ) regSlipNo(key *df.ConditionKey, value interface{}) {
	if q.SlipNo == nil {
		q.SlipNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SlipNo, "slipNo")
}

func (q *SalesSlipCQ) getCValueLine() *df.ConditionValue {
	if q.Line == nil {
		q.Line = new(df.ConditionValue)
	}
	return q.Line
}



func (q *SalesSlipCQ) SetLine_Equal(value int64) *SalesSlipCQ {
	q.regLine(df.CK_EQ_C, value)
	return q
}
func (q *SalesSlipCQ) SetLine_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueLine(), "line")
}
func (q *SalesSlipCQ) SetLine_NotEqual(value int64) *SalesSlipCQ {
	q.regLine(df.CK_NE_C, value)
	return q
}

func (q *SalesSlipCQ) SetLine_GreaterThan(value int64) *SalesSlipCQ {
	q.regLine(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetLine_LessThan(value int64) *SalesSlipCQ {
	q.regLine(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetLine_GreaterEqual(value int64) *SalesSlipCQ {
	q.regLine(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetLine_LessEqual(value int64) *SalesSlipCQ {
	q.regLine(df.CK_LE_C, value)
	return q
}
func (q *SalesSlipCQ) SetLine_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueLine(),"line",rangeOfOption)
}	


func (q *SalesSlipCQ) AddOrderBy_Line_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("line")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_Line_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("line")
	return q
}
func (q *SalesSlipCQ) regLine(key *df.ConditionKey, value interface{}) {
	if q.Line == nil {
		q.Line = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Line, "line")
}

func (q *SalesSlipCQ) getCValueSalesDate() *df.ConditionValue {
	if q.SalesDate == nil {
		q.SalesDate = new(df.ConditionValue)
	}
	return q.SalesDate
}




func (q *SalesSlipCQ) SetSalesDate_Equal(value df.Date) *SalesSlipCQ {
	q.regSalesDate(df.CK_EQ_C, value)
	return q
}


func (q *SalesSlipCQ) SetSalesDate_GreaterThan(value df.Date) *SalesSlipCQ {
	q.regSalesDate(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetSalesDate_LessThan(value df.Date) *SalesSlipCQ {
	q.regSalesDate(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetSalesDate_GreaterEqual(value df.Date) *SalesSlipCQ {
	q.regSalesDate(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetSalesDate_LessEqual(value df.Date) *SalesSlipCQ {
	q.regSalesDate(df.CK_LE_C, value)
	return q
}

func (q *SalesSlipCQ) AddOrderBy_SalesDate_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("salesDate")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_SalesDate_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("salesDate")
	return q
}
func (q *SalesSlipCQ) regSalesDate(key *df.ConditionKey, value interface{}) {
	if q.SalesDate == nil {
		q.SalesDate = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SalesDate, "salesDate")
}

func (q *SalesSlipCQ) getCValueSalesId() *df.ConditionValue {
	if q.SalesId == nil {
		q.SalesId = new(df.ConditionValue)
	}
	return q.SalesId
}



func (q *SalesSlipCQ) SetSalesId_Equal(value int64) *SalesSlipCQ {
	q.regSalesId(df.CK_EQ_C, value)
	return q
}
func (q *SalesSlipCQ) SetSalesId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueSalesId(), "salesId")
}
func (q *SalesSlipCQ) SetSalesId_NotEqual(value int64) *SalesSlipCQ {
	q.regSalesId(df.CK_NE_C, value)
	return q
}

func (q *SalesSlipCQ) SetSalesId_GreaterThan(value int64) *SalesSlipCQ {
	q.regSalesId(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetSalesId_LessThan(value int64) *SalesSlipCQ {
	q.regSalesId(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetSalesId_GreaterEqual(value int64) *SalesSlipCQ {
	q.regSalesId(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetSalesId_LessEqual(value int64) *SalesSlipCQ {
	q.regSalesId(df.CK_LE_C, value)
	return q
}
func (q *SalesSlipCQ) SetSalesId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueSalesId(),"salesId",rangeOfOption)
}	


func (q *SalesSlipCQ) AddOrderBy_SalesId_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("salesId")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_SalesId_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("salesId")
	return q
}
func (q *SalesSlipCQ) regSalesId(key *df.ConditionKey, value interface{}) {
	if q.SalesId == nil {
		q.SalesId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SalesId, "salesId")
}

func (q *SalesSlipCQ) getCValueCusId() *df.ConditionValue {
	if q.CusId == nil {
		q.CusId = new(df.ConditionValue)
	}
	return q.CusId
}



func (q *SalesSlipCQ) SetCusId_Equal(value int64) *SalesSlipCQ {
	q.regCusId(df.CK_EQ_C, value)
	return q
}
func (q *SalesSlipCQ) SetCusId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueCusId(), "cusId")
}
func (q *SalesSlipCQ) SetCusId_NotEqual(value int64) *SalesSlipCQ {
	q.regCusId(df.CK_NE_C, value)
	return q
}

func (q *SalesSlipCQ) SetCusId_GreaterThan(value int64) *SalesSlipCQ {
	q.regCusId(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetCusId_LessThan(value int64) *SalesSlipCQ {
	q.regCusId(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetCusId_GreaterEqual(value int64) *SalesSlipCQ {
	q.regCusId(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetCusId_LessEqual(value int64) *SalesSlipCQ {
	q.regCusId(df.CK_LE_C, value)
	return q
}
func (q *SalesSlipCQ) SetCusId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueCusId(),"cusId",rangeOfOption)
}	


func (q *SalesSlipCQ) AddOrderBy_CusId_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("cusId")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_CusId_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("cusId")
	return q
}
func (q *SalesSlipCQ) regCusId(key *df.ConditionKey, value interface{}) {
	if q.CusId == nil {
		q.CusId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.CusId, "cusId")
}

func (q *SalesSlipCQ) getCValuePrdId() *df.ConditionValue {
	if q.PrdId == nil {
		q.PrdId = new(df.ConditionValue)
	}
	return q.PrdId
}



func (q *SalesSlipCQ) SetPrdId_Equal(value int64) *SalesSlipCQ {
	q.regPrdId(df.CK_EQ_C, value)
	return q
}
func (q *SalesSlipCQ) SetPrdId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePrdId(), "prdId")
}
func (q *SalesSlipCQ) SetPrdId_NotEqual(value int64) *SalesSlipCQ {
	q.regPrdId(df.CK_NE_C, value)
	return q
}

func (q *SalesSlipCQ) SetPrdId_GreaterThan(value int64) *SalesSlipCQ {
	q.regPrdId(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetPrdId_LessThan(value int64) *SalesSlipCQ {
	q.regPrdId(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetPrdId_GreaterEqual(value int64) *SalesSlipCQ {
	q.regPrdId(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetPrdId_LessEqual(value int64) *SalesSlipCQ {
	q.regPrdId(df.CK_LE_C, value)
	return q
}
func (q *SalesSlipCQ) SetPrdId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePrdId(),"prdId",rangeOfOption)
}	


func (q *SalesSlipCQ) AddOrderBy_PrdId_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("prdId")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_PrdId_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("prdId")
	return q
}
func (q *SalesSlipCQ) regPrdId(key *df.ConditionKey, value interface{}) {
	if q.PrdId == nil {
		q.PrdId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PrdId, "prdId")
}

func (q *SalesSlipCQ) getCValueQty() *df.ConditionValue {
	if q.Qty == nil {
		q.Qty = new(df.ConditionValue)
	}
	return q.Qty
}



func (q *SalesSlipCQ) SetQty_Equal(value int64) *SalesSlipCQ {
	q.regQty(df.CK_EQ_C, value)
	return q
}
func (q *SalesSlipCQ) SetQty_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueQty(), "qty")
}
func (q *SalesSlipCQ) SetQty_NotEqual(value int64) *SalesSlipCQ {
	q.regQty(df.CK_NE_C, value)
	return q
}

func (q *SalesSlipCQ) SetQty_GreaterThan(value int64) *SalesSlipCQ {
	q.regQty(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetQty_LessThan(value int64) *SalesSlipCQ {
	q.regQty(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetQty_GreaterEqual(value int64) *SalesSlipCQ {
	q.regQty(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetQty_LessEqual(value int64) *SalesSlipCQ {
	q.regQty(df.CK_LE_C, value)
	return q
}
func (q *SalesSlipCQ) SetQty_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueQty(),"qty",rangeOfOption)
}	


func (q *SalesSlipCQ) AddOrderBy_Qty_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("qty")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_Qty_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("qty")
	return q
}
func (q *SalesSlipCQ) regQty(key *df.ConditionKey, value interface{}) {
	if q.Qty == nil {
		q.Qty = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Qty, "qty")
}

func (q *SalesSlipCQ) getCValueUnit() *df.ConditionValue {
	if q.Unit == nil {
		q.Unit = new(df.ConditionValue)
	}
	return q.Unit
}


func (q *SalesSlipCQ) SetUnit_Equal(value string) *SalesSlipCQ {
	q.regUnit(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SalesSlipCQ) SetUnit_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUnit(), "unit")
}
func (q *SalesSlipCQ) SetUnit_NotEqual(value string) *SalesSlipCQ {
	q.regUnit(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetUnit_GreaterThan(value string) *SalesSlipCQ {
	q.regUnit(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetUnit_LessThan(value string) *SalesSlipCQ {
	q.regUnit(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetUnit_GreaterEqual(value string) *SalesSlipCQ {
	q.regUnit(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SalesSlipCQ) SetUnit_LessEqual(value string) *SalesSlipCQ {
	q.regUnit(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetUnit_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUnit(), "unit", option)
}

func (q *SalesSlipCQ) SetUnit_PrefixSearch(value string) error {
	return q.SetUnit_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SalesSlipCQ) SetUnit_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUnit(), "unit", option)
}



func (q *SalesSlipCQ) AddOrderBy_Unit_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("unit")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_Unit_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("unit")
	return q
}
func (q *SalesSlipCQ) regUnit(key *df.ConditionKey, value interface{}) {
	if q.Unit == nil {
		q.Unit = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Unit, "unit")
}

func (q *SalesSlipCQ) getCValueUnitPrice() *df.ConditionValue {
	if q.UnitPrice == nil {
		q.UnitPrice = new(df.ConditionValue)
	}
	return q.UnitPrice
}



func (q *SalesSlipCQ) SetUnitPrice_Equal(value int64) *SalesSlipCQ {
	q.regUnitPrice(df.CK_EQ_C, value)
	return q
}
func (q *SalesSlipCQ) SetUnitPrice_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUnitPrice(), "unitPrice")
}
func (q *SalesSlipCQ) SetUnitPrice_NotEqual(value int64) *SalesSlipCQ {
	q.regUnitPrice(df.CK_NE_C, value)
	return q
}

func (q *SalesSlipCQ) SetUnitPrice_GreaterThan(value int64) *SalesSlipCQ {
	q.regUnitPrice(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetUnitPrice_LessThan(value int64) *SalesSlipCQ {
	q.regUnitPrice(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetUnitPrice_GreaterEqual(value int64) *SalesSlipCQ {
	q.regUnitPrice(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetUnitPrice_LessEqual(value int64) *SalesSlipCQ {
	q.regUnitPrice(df.CK_LE_C, value)
	return q
}
func (q *SalesSlipCQ) SetUnitPrice_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueUnitPrice(),"unitPrice",rangeOfOption)
}	


func (q *SalesSlipCQ) AddOrderBy_UnitPrice_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("unitPrice")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_UnitPrice_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("unitPrice")
	return q
}
func (q *SalesSlipCQ) regUnitPrice(key *df.ConditionKey, value interface{}) {
	if q.UnitPrice == nil {
		q.UnitPrice = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UnitPrice, "unitPrice")
}

func (q *SalesSlipCQ) getCValueSalesAmount() *df.ConditionValue {
	if q.SalesAmount == nil {
		q.SalesAmount = new(df.ConditionValue)
	}
	return q.SalesAmount
}



func (q *SalesSlipCQ) SetSalesAmount_Equal(value int64) *SalesSlipCQ {
	q.regSalesAmount(df.CK_EQ_C, value)
	return q
}
func (q *SalesSlipCQ) SetSalesAmount_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueSalesAmount(), "salesAmount")
}
func (q *SalesSlipCQ) SetSalesAmount_NotEqual(value int64) *SalesSlipCQ {
	q.regSalesAmount(df.CK_NE_C, value)
	return q
}

func (q *SalesSlipCQ) SetSalesAmount_GreaterThan(value int64) *SalesSlipCQ {
	q.regSalesAmount(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetSalesAmount_LessThan(value int64) *SalesSlipCQ {
	q.regSalesAmount(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetSalesAmount_GreaterEqual(value int64) *SalesSlipCQ {
	q.regSalesAmount(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetSalesAmount_LessEqual(value int64) *SalesSlipCQ {
	q.regSalesAmount(df.CK_LE_C, value)
	return q
}
func (q *SalesSlipCQ) SetSalesAmount_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueSalesAmount(),"salesAmount",rangeOfOption)
}	


func (q *SalesSlipCQ) AddOrderBy_SalesAmount_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("salesAmount")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_SalesAmount_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("salesAmount")
	return q
}
func (q *SalesSlipCQ) regSalesAmount(key *df.ConditionKey, value interface{}) {
	if q.SalesAmount == nil {
		q.SalesAmount = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SalesAmount, "salesAmount")
}

func (q *SalesSlipCQ) getCValueSlipAmount() *df.ConditionValue {
	if q.SlipAmount == nil {
		q.SlipAmount = new(df.ConditionValue)
	}
	return q.SlipAmount
}



func (q *SalesSlipCQ) SetSlipAmount_Equal(value int64) *SalesSlipCQ {
	q.regSlipAmount(df.CK_EQ_C, value)
	return q
}
func (q *SalesSlipCQ) SetSlipAmount_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueSlipAmount(), "slipAmount")
}
func (q *SalesSlipCQ) SetSlipAmount_NotEqual(value int64) *SalesSlipCQ {
	q.regSlipAmount(df.CK_NE_C, value)
	return q
}

func (q *SalesSlipCQ) SetSlipAmount_GreaterThan(value int64) *SalesSlipCQ {
	q.regSlipAmount(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetSlipAmount_LessThan(value int64) *SalesSlipCQ {
	q.regSlipAmount(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetSlipAmount_GreaterEqual(value int64) *SalesSlipCQ {
	q.regSlipAmount(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetSlipAmount_LessEqual(value int64) *SalesSlipCQ {
	q.regSlipAmount(df.CK_LE_C, value)
	return q
}
func (q *SalesSlipCQ) SetSlipAmount_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueSlipAmount(),"slipAmount",rangeOfOption)
}	


func (q *SalesSlipCQ) AddOrderBy_SlipAmount_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("slipAmount")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_SlipAmount_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("slipAmount")
	return q
}
func (q *SalesSlipCQ) regSlipAmount(key *df.ConditionKey, value interface{}) {
	if q.SlipAmount == nil {
		q.SlipAmount = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SlipAmount, "slipAmount")
}

func (q *SalesSlipCQ) getCValueSlipCons() *df.ConditionValue {
	if q.SlipCons == nil {
		q.SlipCons = new(df.ConditionValue)
	}
	return q.SlipCons
}



func (q *SalesSlipCQ) SetSlipCons_Equal(value int64) *SalesSlipCQ {
	q.regSlipCons(df.CK_EQ_C, value)
	return q
}
func (q *SalesSlipCQ) SetSlipCons_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueSlipCons(), "slipCons")
}
func (q *SalesSlipCQ) SetSlipCons_NotEqual(value int64) *SalesSlipCQ {
	q.regSlipCons(df.CK_NE_C, value)
	return q
}

func (q *SalesSlipCQ) SetSlipCons_GreaterThan(value int64) *SalesSlipCQ {
	q.regSlipCons(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetSlipCons_LessThan(value int64) *SalesSlipCQ {
	q.regSlipCons(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetSlipCons_GreaterEqual(value int64) *SalesSlipCQ {
	q.regSlipCons(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetSlipCons_LessEqual(value int64) *SalesSlipCQ {
	q.regSlipCons(df.CK_LE_C, value)
	return q
}
func (q *SalesSlipCQ) SetSlipCons_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueSlipCons(),"slipCons",rangeOfOption)
}	


func (q *SalesSlipCQ) AddOrderBy_SlipCons_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("slipCons")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_SlipCons_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("slipCons")
	return q
}
func (q *SalesSlipCQ) regSlipCons(key *df.ConditionKey, value interface{}) {
	if q.SlipCons == nil {
		q.SlipCons = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SlipCons, "slipCons")
}

func (q *SalesSlipCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *SalesSlipCQ) SetVersionNo_Equal(value int64) *SalesSlipCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}
func (q *SalesSlipCQ) SetVersionNo_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueVersionNo(), "versionNo")
}
func (q *SalesSlipCQ) SetVersionNo_NotEqual(value int64) *SalesSlipCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *SalesSlipCQ) SetVersionNo_GreaterThan(value int64) *SalesSlipCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetVersionNo_LessThan(value int64) *SalesSlipCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetVersionNo_GreaterEqual(value int64) *SalesSlipCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetVersionNo_LessEqual(value int64) *SalesSlipCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *SalesSlipCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *SalesSlipCQ) AddOrderBy_VersionNo_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_VersionNo_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *SalesSlipCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}

func (q *SalesSlipCQ) getCValueDelFlag() *df.ConditionValue {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	return q.DelFlag
}



func (q *SalesSlipCQ) SetDelFlag_Equal(value int64) *SalesSlipCQ {
	q.regDelFlag(df.CK_EQ_C, value)
	return q
}
func (q *SalesSlipCQ) SetDelFlag_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDelFlag(), "delFlag")
}
func (q *SalesSlipCQ) SetDelFlag_NotEqual(value int64) *SalesSlipCQ {
	q.regDelFlag(df.CK_NE_C, value)
	return q
}

func (q *SalesSlipCQ) SetDelFlag_GreaterThan(value int64) *SalesSlipCQ {
	q.regDelFlag(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetDelFlag_LessThan(value int64) *SalesSlipCQ {
	q.regDelFlag(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetDelFlag_GreaterEqual(value int64) *SalesSlipCQ {
	q.regDelFlag(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetDelFlag_LessEqual(value int64) *SalesSlipCQ {
	q.regDelFlag(df.CK_LE_C, value)
	return q
}
func (q *SalesSlipCQ) SetDelFlag_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDelFlag(),"delFlag",rangeOfOption)
}	


func (q *SalesSlipCQ) AddOrderBy_DelFlag_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("delFlag")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_DelFlag_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("delFlag")
	return q
}
func (q *SalesSlipCQ) regDelFlag(key *df.ConditionKey, value interface{}) {
	if q.DelFlag == nil {
		q.DelFlag = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DelFlag, "delFlag")
}

func (q *SalesSlipCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *SalesSlipCQ) SetRegisterDatetime_Equal(value df.Timestamp) *SalesSlipCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *SalesSlipCQ) SetRegisterDatetime_GreaterThan(value df.Timestamp) *SalesSlipCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetRegisterDatetime_LessThan(value df.Timestamp) *SalesSlipCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetRegisterDatetime_GreaterEqual(value df.Timestamp) *SalesSlipCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetRegisterDatetime_LessEqual(value df.Timestamp) *SalesSlipCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *SalesSlipCQ) AddOrderBy_RegisterDatetime_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_RegisterDatetime_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *SalesSlipCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *SalesSlipCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *SalesSlipCQ) SetRegisterUser_Equal(value string) *SalesSlipCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SalesSlipCQ) SetRegisterUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterUser(), "registerUser")
}
func (q *SalesSlipCQ) SetRegisterUser_NotEqual(value string) *SalesSlipCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetRegisterUser_GreaterThan(value string) *SalesSlipCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetRegisterUser_LessThan(value string) *SalesSlipCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetRegisterUser_GreaterEqual(value string) *SalesSlipCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SalesSlipCQ) SetRegisterUser_LessEqual(value string) *SalesSlipCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *SalesSlipCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SalesSlipCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *SalesSlipCQ) AddOrderBy_RegisterUser_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_RegisterUser_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *SalesSlipCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *SalesSlipCQ) getCValueRegisterProcess() *df.ConditionValue {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	return q.RegisterProcess
}


func (q *SalesSlipCQ) SetRegisterProcess_Equal(value string) *SalesSlipCQ {
	q.regRegisterProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SalesSlipCQ) SetRegisterProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterProcess(), "registerProcess")
}
func (q *SalesSlipCQ) SetRegisterProcess_NotEqual(value string) *SalesSlipCQ {
	q.regRegisterProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetRegisterProcess_GreaterThan(value string) *SalesSlipCQ {
	q.regRegisterProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetRegisterProcess_LessThan(value string) *SalesSlipCQ {
	q.regRegisterProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetRegisterProcess_GreaterEqual(value string) *SalesSlipCQ {
	q.regRegisterProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SalesSlipCQ) SetRegisterProcess_LessEqual(value string) *SalesSlipCQ {
	q.regRegisterProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetRegisterProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}

func (q *SalesSlipCQ) SetRegisterProcess_PrefixSearch(value string) error {
	return q.SetRegisterProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SalesSlipCQ) SetRegisterProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}



func (q *SalesSlipCQ) AddOrderBy_RegisterProcess_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("registerProcess")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_RegisterProcess_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("registerProcess")
	return q
}
func (q *SalesSlipCQ) regRegisterProcess(key *df.ConditionKey, value interface{}) {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterProcess, "registerProcess")
}

func (q *SalesSlipCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *SalesSlipCQ) SetUpdateDatetime_Equal(value df.Timestamp) *SalesSlipCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *SalesSlipCQ) SetUpdateDatetime_GreaterThan(value df.Timestamp) *SalesSlipCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *SalesSlipCQ) SetUpdateDatetime_LessThan(value df.Timestamp) *SalesSlipCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *SalesSlipCQ) SetUpdateDatetime_GreaterEqual(value df.Timestamp) *SalesSlipCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *SalesSlipCQ) SetUpdateDatetime_LessEqual(value df.Timestamp) *SalesSlipCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *SalesSlipCQ) AddOrderBy_UpdateDatetime_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_UpdateDatetime_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *SalesSlipCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *SalesSlipCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *SalesSlipCQ) SetUpdateUser_Equal(value string) *SalesSlipCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SalesSlipCQ) SetUpdateUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateUser(), "updateUser")
}
func (q *SalesSlipCQ) SetUpdateUser_NotEqual(value string) *SalesSlipCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetUpdateUser_GreaterThan(value string) *SalesSlipCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetUpdateUser_LessThan(value string) *SalesSlipCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetUpdateUser_GreaterEqual(value string) *SalesSlipCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SalesSlipCQ) SetUpdateUser_LessEqual(value string) *SalesSlipCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *SalesSlipCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SalesSlipCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *SalesSlipCQ) AddOrderBy_UpdateUser_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_UpdateUser_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *SalesSlipCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *SalesSlipCQ) getCValueUpdateProcess() *df.ConditionValue {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	return q.UpdateProcess
}


func (q *SalesSlipCQ) SetUpdateProcess_Equal(value string) *SalesSlipCQ {
	q.regUpdateProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SalesSlipCQ) SetUpdateProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateProcess(), "updateProcess")
}
func (q *SalesSlipCQ) SetUpdateProcess_NotEqual(value string) *SalesSlipCQ {
	q.regUpdateProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetUpdateProcess_GreaterThan(value string) *SalesSlipCQ {
	q.regUpdateProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetUpdateProcess_LessThan(value string) *SalesSlipCQ {
	q.regUpdateProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetUpdateProcess_GreaterEqual(value string) *SalesSlipCQ {
	q.regUpdateProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SalesSlipCQ) SetUpdateProcess_LessEqual(value string) *SalesSlipCQ {
	q.regUpdateProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SalesSlipCQ) SetUpdateProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}

func (q *SalesSlipCQ) SetUpdateProcess_PrefixSearch(value string) error {
	return q.SetUpdateProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SalesSlipCQ) SetUpdateProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}



func (q *SalesSlipCQ) AddOrderBy_UpdateProcess_Asc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBA("updateProcess")
	return q
}
func (q *SalesSlipCQ) AddOrderBy_UpdateProcess_Desc() *SalesSlipCQ {
	q.BaseConditionQuery.RegOBD("updateProcess")
	return q
}
func (q *SalesSlipCQ) regUpdateProcess(key *df.ConditionKey, value interface{}) {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateProcess, "updateProcess")
}


func (q *SalesSlipCQ) QueryCustomer() *CustomerCQ {
	if q.conditionQueryCustomer == nil {
		q.conditionQueryCustomer = q.xcreateQueryCustomer()
		q.xsetupOuterJoinCustomer()
	}
	return q.conditionQueryCustomer
}

func (q *SalesSlipCQ) xcreateQueryCustomer() *CustomerCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("SalesSlip", "Customer")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreateCustomerCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "Customer"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *SalesSlipCQ) xsetupOuterJoinCustomer() {
	    cq := q.QueryCustomer()
        joinOnMap := make(map[string]string)
        joinOnMap["cusId"]="id"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "Customer");
}	
	
func (q *SalesSlipCQ) QueryProduct() *ProductCQ {
	if q.conditionQueryProduct == nil {
		q.conditionQueryProduct = q.xcreateQueryProduct()
		q.xsetupOuterJoinProduct()
	}
	return q.conditionQueryProduct
}

func (q *SalesSlipCQ) xcreateQueryProduct() *ProductCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("SalesSlip", "Product")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreateProductCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "Product"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *SalesSlipCQ) xsetupOuterJoinProduct() {
	    cq := q.QueryProduct()
        joinOnMap := make(map[string]string)
        joinOnMap["prdId"]="id"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "Product");
}	
	
func (q *SalesSlipCQ) QueryEmployee() *EmployeeCQ {
	if q.conditionQueryEmployee == nil {
		q.conditionQueryEmployee = q.xcreateQueryEmployee()
		q.xsetupOuterJoinEmployee()
	}
	return q.conditionQueryEmployee
}

func (q *SalesSlipCQ) xcreateQueryEmployee() *EmployeeCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("SalesSlip", "Employee")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreateEmployeeCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "Employee"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *SalesSlipCQ) xsetupOuterJoinEmployee() {
	    cq := q.QueryEmployee()
        joinOnMap := make(map[string]string)
        joinOnMap["salesId"]="id"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "Employee");
}	
	
func CreateSalesSlipCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *SalesSlipCQ {
	cq := new(SalesSlipCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "SalesSlip"
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
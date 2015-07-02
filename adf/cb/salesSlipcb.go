package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"godbfexam/adf/cq"
	"godbfexam/adf/meta"
)

type SalesSlipCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.SalesSlipCQ
    NssProduct *ProductNss
    NssEmployee *EmployeeNss
}

func CreateSalesSlipCB() *SalesSlipCB {
	cb := new(SalesSlipCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "SalesSlip"
	cb.BaseConditionBean.Name = "SalesSlipCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.SalesSlipDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *SalesSlipCB) Query() *cq.SalesSlipCQ {
	if l.query == nil {
		l.query = cq.CreateSalesSlipCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *SalesSlipCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *SalesSlipCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *SalesSlipCB) SetupSelect_Product() *ProductNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryProduct().GetBaseConditionQuery())
	if l.NssProduct == nil || !l.NssProduct.hasConditionQuery() {
		l.NssProduct = new(ProductNss)
		l.NssProduct.Query = l.Query().QueryProduct()
	}
	return l.NssProduct
}
func (l *SalesSlipCB) SetupSelect_Employee() *EmployeeNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryEmployee().GetBaseConditionQuery())
	if l.NssEmployee == nil || !l.NssEmployee.hasConditionQuery() {
		l.NssEmployee = new(EmployeeNss)
		l.NssEmployee.Query = l.Query().QueryEmployee()
	}
	return l.NssEmployee
}

func (l *SalesSlipCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *SalesSlipCB) OrScopeQuery(fquery func(*SalesSlipCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type SalesSlipNss struct {
	Query *cq.SalesSlipCQ
    NssProduct *ProductNss
    NssEmployee *EmployeeNss
}
func (p *SalesSlipNss) WithProduct() *ProductNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryProduct().GetBaseConditionQuery())
	if p.NssProduct == nil || !p.NssProduct.hasConditionQuery() {
		p.NssProduct = new(ProductNss)
		p.NssProduct.Query = p.Query.QueryProduct()
	}
	return p.NssProduct
}
func (p *SalesSlipNss) WithEmployee() *EmployeeNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryEmployee().GetBaseConditionQuery())
	if p.NssEmployee == nil || !p.NssEmployee.hasConditionQuery() {
		p.NssEmployee = new(EmployeeNss)
		p.NssEmployee.Query = p.Query.QueryEmployee()
	}
	return p.NssEmployee
}
func (p *SalesSlipNss) hasConditionQuery() bool {
	return p.Query != nil
}
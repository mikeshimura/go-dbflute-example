package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"godbfexam/adf/cq"
	"godbfexam/adf/meta"
)

type StockCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.StockCQ
    NssProduct *ProductNss
}

func CreateStockCB() *StockCB {
	cb := new(StockCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "Stock"
	cb.BaseConditionBean.Name = "StockCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.StockDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *StockCB) Query() *cq.StockCQ {
	if l.query == nil {
		l.query = cq.CreateStockCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *StockCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *StockCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *StockCB) SetupSelect_Product() *ProductNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryProduct().GetBaseConditionQuery())
	if l.NssProduct == nil || !l.NssProduct.hasConditionQuery() {
		l.NssProduct = new(ProductNss)
		l.NssProduct.Query = l.Query().QueryProduct()
	}
	return l.NssProduct
}

func (l *StockCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *StockCB) OrScopeQuery(fquery func(*StockCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type StockNss struct {
	Query *cq.StockCQ
    NssProduct *ProductNss
}
func (p *StockNss) WithProduct() *ProductNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryProduct().GetBaseConditionQuery())
	if p.NssProduct == nil || !p.NssProduct.hasConditionQuery() {
		p.NssProduct = new(ProductNss)
		p.NssProduct.Query = p.Query.QueryProduct()
	}
	return p.NssProduct
}
func (p *StockNss) hasConditionQuery() bool {
	return p.Query != nil
}
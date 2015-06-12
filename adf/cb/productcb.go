package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"godbfexam/adf/cq"
	"godbfexam/adf/meta"
)

type ProductCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.ProductCQ
}

func CreateProductCB() *ProductCB {
	cb := new(ProductCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "Product"
	cb.BaseConditionBean.Name = "ProductCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.ProductDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *ProductCB) Query() *cq.ProductCQ {
	if l.query == nil {
		l.query = cq.CreateProductCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *ProductCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *ProductCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}


func (l *ProductCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *ProductCB) OrScopeQuery(fquery func(*ProductCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type ProductNss struct {
	Query *cq.ProductCQ
}
func (p *ProductNss) hasConditionQuery() bool {
	return p.Query != nil
}
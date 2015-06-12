package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"godbfexam/adf/cq"
	"godbfexam/adf/meta"
)

type CustomerCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.CustomerCQ
}

func CreateCustomerCB() *CustomerCB {
	cb := new(CustomerCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "Customer"
	cb.BaseConditionBean.Name = "CustomerCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.CustomerDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *CustomerCB) Query() *cq.CustomerCQ {
	if l.query == nil {
		l.query = cq.CreateCustomerCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *CustomerCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *CustomerCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}


func (l *CustomerCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *CustomerCB) OrScopeQuery(fquery func(*CustomerCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type CustomerNss struct {
	Query *cq.CustomerCQ
}
func (p *CustomerNss) hasConditionQuery() bool {
	return p.Query != nil
}
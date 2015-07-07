package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"godbfexam/adf/cq"
	"godbfexam/adf/meta"
)

type TestTableCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.TestTableCQ
}

func CreateTestTableCB() *TestTableCB {
	cb := new(TestTableCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "TestTable"
	cb.BaseConditionBean.Name = "TestTableCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.TestTableDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *TestTableCB) Query() *cq.TestTableCQ {
	if l.query == nil {
		l.query = cq.CreateTestTableCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *TestTableCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *TestTableCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}


func (l *TestTableCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *TestTableCB) OrScopeQuery(fquery func(*TestTableCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type TestTableNss struct {
	Query *cq.TestTableCQ
}
func (p *TestTableNss) hasConditionQuery() bool {
	return p.Query != nil
}
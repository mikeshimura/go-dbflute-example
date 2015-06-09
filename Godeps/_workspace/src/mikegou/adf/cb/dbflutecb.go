package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"mikegou/adf/cq"
	"mikegou/adf/meta"
)

type DbfluteCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.DbfluteCQ
    NssLogin *LoginNss
}

func CreateDbfluteCB() *DbfluteCB {
	cb := new(DbfluteCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "Dbflute"
	cb.BaseConditionBean.Name = "DbfluteCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.DbfluteDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *DbfluteCB) Query() *cq.DbfluteCQ {
	if l.query == nil {
		l.query = cq.CreateDbfluteCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *DbfluteCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *DbfluteCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *DbfluteCB) SetupSelect_Login() *LoginNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryLogin().GetBaseConditionQuery())
	if l.NssLogin == nil || !l.NssLogin.hasConditionQuery() {
		l.NssLogin = new(LoginNss)
		l.NssLogin.Query = l.Query().QueryLogin()
	}
	return l.NssLogin
}

func (l *DbfluteCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *DbfluteCB) OrScopeQuery(fquery func(*DbfluteCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type DbfluteNss struct {
	Query *cq.DbfluteCQ
    NssLogin *LoginNss
}
func (p *DbfluteNss) WithLogin() *LoginNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryLogin().GetBaseConditionQuery())
	if p.NssLogin == nil || !p.NssLogin.hasConditionQuery() {
		p.NssLogin = new(LoginNss)
		p.NssLogin.Query = p.Query.QueryLogin()
	}
	return p.NssLogin
}
func (p *DbfluteNss) hasConditionQuery() bool {
	return p.Query != nil
}

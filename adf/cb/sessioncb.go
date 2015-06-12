package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"godbfexam/adf/cq"
	"godbfexam/adf/meta"
)

type SessionCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.SessionCQ
    NssLogin *LoginNss
}

func CreateSessionCB() *SessionCB {
	cb := new(SessionCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "Session"
	cb.BaseConditionBean.Name = "SessionCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.SessionDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *SessionCB) Query() *cq.SessionCQ {
	if l.query == nil {
		l.query = cq.CreateSessionCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *SessionCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *SessionCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *SessionCB) SetupSelect_Login() *LoginNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryLogin().GetBaseConditionQuery())
	if l.NssLogin == nil || !l.NssLogin.hasConditionQuery() {
		l.NssLogin = new(LoginNss)
		l.NssLogin.Query = l.Query().QueryLogin()
	}
	return l.NssLogin
}

func (l *SessionCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *SessionCB) OrScopeQuery(fquery func(*SessionCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type SessionNss struct {
	Query *cq.SessionCQ
    NssLogin *LoginNss
}
func (p *SessionNss) WithLogin() *LoginNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryLogin().GetBaseConditionQuery())
	if p.NssLogin == nil || !p.NssLogin.hasConditionQuery() {
		p.NssLogin = new(LoginNss)
		p.NssLogin.Query = p.Query.QueryLogin()
	}
	return p.NssLogin
}
func (p *SessionNss) hasConditionQuery() bool {
	return p.Query != nil
}
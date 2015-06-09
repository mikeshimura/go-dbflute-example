package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"mikegou/adf/cq"
	"mikegou/adf/meta"
)

type UserTableCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.UserTableCQ
    NssLogin *LoginNss
}

func CreateUserTableCB() *UserTableCB {
	cb := new(UserTableCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "UserTable"
	cb.BaseConditionBean.Name = "UserTableCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.UserTableDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *UserTableCB) Query() *cq.UserTableCQ {
	if l.query == nil {
		l.query = cq.CreateUserTableCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *UserTableCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *UserTableCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *UserTableCB) SetupSelect_Login() *LoginNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryLogin().GetBaseConditionQuery())
	if l.NssLogin == nil || !l.NssLogin.hasConditionQuery() {
		l.NssLogin = new(LoginNss)
		l.NssLogin.Query = l.Query().QueryLogin()
	}
	return l.NssLogin
}

func (l *UserTableCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *UserTableCB) OrScopeQuery(fquery func(*UserTableCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type UserTableNss struct {
	Query *cq.UserTableCQ
    NssLogin *LoginNss
}
func (p *UserTableNss) WithLogin() *LoginNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryLogin().GetBaseConditionQuery())
	if p.NssLogin == nil || !p.NssLogin.hasConditionQuery() {
		p.NssLogin = new(LoginNss)
		p.NssLogin.Query = p.Query.QueryLogin()
	}
	return p.NssLogin
}
func (p *UserTableNss) hasConditionQuery() bool {
	return p.Query != nil
}

package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"godbfexam/adf/cq"
	"godbfexam/adf/meta"
)

type EmployeeCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.EmployeeCQ
    NssUserTable *UserTableNss
}

func CreateEmployeeCB() *EmployeeCB {
	cb := new(EmployeeCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "Employee"
	cb.BaseConditionBean.Name = "EmployeeCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.EmployeeDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *EmployeeCB) Query() *cq.EmployeeCQ {
	if l.query == nil {
		l.query = cq.CreateEmployeeCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *EmployeeCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *EmployeeCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *EmployeeCB) SetupSelect_UserTable() *UserTableNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryUserTable().GetBaseConditionQuery())
	if l.NssUserTable == nil || !l.NssUserTable.hasConditionQuery() {
		l.NssUserTable = new(UserTableNss)
		l.NssUserTable.Query = l.Query().QueryUserTable()
	}
	return l.NssUserTable
}

func (l *EmployeeCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *EmployeeCB) OrScopeQuery(fquery func(*EmployeeCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type EmployeeNss struct {
	Query *cq.EmployeeCQ
    NssUserTable *UserTableNss
}
func (p *EmployeeNss) WithUserTable() *UserTableNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryUserTable().GetBaseConditionQuery())
	if p.NssUserTable == nil || !p.NssUserTable.hasConditionQuery() {
		p.NssUserTable = new(UserTableNss)
		p.NssUserTable.Query = p.Query.QueryUserTable()
	}
	return p.NssUserTable
}
func (p *EmployeeNss) hasConditionQuery() bool {
	return p.Query != nil
}
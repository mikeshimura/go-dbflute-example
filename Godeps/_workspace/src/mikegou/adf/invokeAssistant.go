package adf

import (
	"github.com/mikeshimura/dbflute/df"
)

type InvokerAssistantImpl struct {
	StatementFactory          *df.StatementFactory
	OutsideSqlExecutorFactory *df.OutsideSqlExecutorFactory
}

func (i *InvokerAssistantImpl) GetDBCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}
func (i *InvokerAssistantImpl) Create() {
	tn := new(df.TnStatementFactoryImpl)
	var sf df.StatementFactory = tn
	i.StatementFactory = &sf
}
func (i *InvokerAssistantImpl) GetDBMetaProvider() *df.DBMetaProvider {
	return df.DBMetaProvider_I
}
func (i *InvokerAssistantImpl) GetStatementFactory() *df.StatementFactory {
	return i.StatementFactory
}
func (i *InvokerAssistantImpl) AssistOutsideSqlExecutorFactory() *df.OutsideSqlExecutorFactory {
	if i.OutsideSqlExecutorFactory != nil {
		return i.OutsideSqlExecutorFactory
	}
	var factory = new(df.DefaultOutsideSqlExecutorFactory)
	var of df.OutsideSqlExecutorFactory = factory
	i.OutsideSqlExecutorFactory = &of
	return &of
}

var InvokerAssistant_I df.InvokerAssistant

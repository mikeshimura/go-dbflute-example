package entity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type TestTable struct {
	id int64
	testId string
	testDate df.NullDate
	testTimestamp df.NullTimestamp
	testNbr sql.NullInt64
	versionNo int64
	delFlag int64
	registerDatetime df.Timestamp
	registerUser string
	registerProcess string
	updateDatetime df.Timestamp
	updateUser string
	updateProcess string
	df.BaseEntity
}

func CreateTestTable() *TestTable{
	testTable :=new(TestTable)
	testTable.SetUp()
	return testTable 
}

func (l *TestTable) GetId () int64 {
	return l.id
}
func (l *TestTable) GetTestId () string {
	return l.testId
}
func (l *TestTable) GetTestDate () df.NullDate {
	return l.testDate
}
func (l *TestTable) GetTestTimestamp () df.NullTimestamp {
	return l.testTimestamp
}
func (l *TestTable) GetTestNbr () sql.NullInt64 {
	return l.testNbr
}
func (l *TestTable) GetVersionNo () int64 {
	return l.versionNo
}
func (l *TestTable) GetDelFlag () int64 {
	return l.delFlag
}
func (l *TestTable) GetRegisterDatetime () df.Timestamp {
	return l.registerDatetime
}
func (l *TestTable) GetRegisterUser () string {
	return l.registerUser
}
func (l *TestTable) GetRegisterProcess () string {
	return l.registerProcess
}
func (l *TestTable) GetUpdateDatetime () df.Timestamp {
	return l.updateDatetime
}
func (l *TestTable) GetUpdateUser () string {
	return l.updateUser
}
func (l *TestTable) GetUpdateProcess () string {
	return l.updateProcess
}

func (t *TestTable) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 13)
	i[0] = &(t.id)
	i[1] = &(t.testId)
	i[2] = &(t.testDate)
	i[3] = &(t.testTimestamp)
	i[4] = &(t.testNbr)
	i[5] = &(t.versionNo)
	i[6] = &(t.delFlag)
	i[7] = &(t.registerDatetime)
	i[8] = &(t.registerUser)
	i[9] = &(t.registerProcess)
	i[10] = &(t.updateDatetime)
	i[11] = &(t.updateUser)
	i[12] = &(t.updateProcess)
	return i
}


func (t *TestTable) AsTableDbName() string {
	return "TestTable"
}

func (t *TestTable) HasPrimaryKeyValue() bool{
        if t.id == 0 {
            return false 
        }
        return true
}
func (t *TestTable) SetId(id int64) {
	t.AddPropertyName("id")
	t.id = id
}
func (t *TestTable) SetTestId(testId string) {
	t.AddPropertyName("testId")
	t.testId = testId
}
func (t *TestTable) SetTestDate(testDate df.NullDate) {
	t.AddPropertyName("testDate")
	t.testDate = testDate
}
func (t *TestTable) SetTestTimestamp(testTimestamp df.NullTimestamp) {
	t.AddPropertyName("testTimestamp")
	t.testTimestamp = testTimestamp
}
func (t *TestTable) SetTestNbr(testNbr sql.NullInt64) {
	t.AddPropertyName("testNbr")
	t.testNbr = testNbr
}
func (t *TestTable) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}
func (t *TestTable) SetDelFlag(delFlag int64) {
	t.AddPropertyName("delFlag")
	t.delFlag = delFlag
}
func (t *TestTable) SetRegisterDatetime(registerDatetime df.Timestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *TestTable) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *TestTable) SetRegisterProcess(registerProcess string) {
	t.AddPropertyName("registerProcess")
	t.registerProcess = registerProcess
}
func (t *TestTable) SetUpdateDatetime(updateDatetime df.Timestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *TestTable) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *TestTable) SetUpdateProcess(updateProcess string) {
	t.AddPropertyName("updateProcess")
	t.updateProcess = updateProcess
}
func (t *TestTable) SetUp(){
	
}
func (t *TestTable)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
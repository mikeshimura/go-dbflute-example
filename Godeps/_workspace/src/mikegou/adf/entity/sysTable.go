package entity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type SysTable struct {
	id int64
	tableName string
	key1 string
	key2 string
	s1Data sql.NullString
	s2Data sql.NullString
	s3Data sql.NullString
	n1Data df.NullNumeric
	n2Data df.NullNumeric
	n3Data df.NullNumeric
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

func CreateSysTable() *SysTable{
	sysTable :=new(SysTable)
	sysTable.SetUp()
	return sysTable 
}

func (l *SysTable) GetId () int64 {
	return l.id
}
func (l *SysTable) GetTableName () string {
	return l.tableName
}
func (l *SysTable) GetKey1 () string {
	return l.key1
}
func (l *SysTable) GetKey2 () string {
	return l.key2
}
func (l *SysTable) GetS1Data () sql.NullString {
	return l.s1Data
}
func (l *SysTable) GetS2Data () sql.NullString {
	return l.s2Data
}
func (l *SysTable) GetS3Data () sql.NullString {
	return l.s3Data
}
func (l *SysTable) GetN1Data () df.NullNumeric {
	return l.n1Data
}
func (l *SysTable) GetN2Data () df.NullNumeric {
	return l.n2Data
}
func (l *SysTable) GetN3Data () df.NullNumeric {
	return l.n3Data
}
func (l *SysTable) GetVersionNo () int64 {
	return l.versionNo
}
func (l *SysTable) GetDelFlag () int64 {
	return l.delFlag
}
func (l *SysTable) GetRegisterDatetime () df.Timestamp {
	return l.registerDatetime
}
func (l *SysTable) GetRegisterUser () string {
	return l.registerUser
}
func (l *SysTable) GetRegisterProcess () string {
	return l.registerProcess
}
func (l *SysTable) GetUpdateDatetime () df.Timestamp {
	return l.updateDatetime
}
func (l *SysTable) GetUpdateUser () string {
	return l.updateUser
}
func (l *SysTable) GetUpdateProcess () string {
	return l.updateProcess
}

func (t *SysTable) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 18)
	i[0] = &(t.id)
	i[1] = &(t.tableName)
	i[2] = &(t.key1)
	i[3] = &(t.key2)
	i[4] = &(t.s1Data)
	i[5] = &(t.s2Data)
	i[6] = &(t.s3Data)
	i[7] = &(t.n1Data)
	i[8] = &(t.n2Data)
	i[9] = &(t.n3Data)
	i[10] = &(t.versionNo)
	i[11] = &(t.delFlag)
	i[12] = &(t.registerDatetime)
	i[13] = &(t.registerUser)
	i[14] = &(t.registerProcess)
	i[15] = &(t.updateDatetime)
	i[16] = &(t.updateUser)
	i[17] = &(t.updateProcess)
	return i
}


func (t *SysTable) AsTableDbName() string {
	return "SysTable"
}

func (t *SysTable) HasPrimaryKeyValue() bool{
        if t.id == 0 {
            return false 
        }
        return true
}
func (t *SysTable) SetId(id int64) {
	t.AddPropertyName("id")
	t.id = id
}
func (t *SysTable) SetTableName(tableName string) {
	t.AddPropertyName("tableName")
	t.tableName = tableName
}
func (t *SysTable) SetKey1(key1 string) {
	t.AddPropertyName("key1")
	t.key1 = key1
}
func (t *SysTable) SetKey2(key2 string) {
	t.AddPropertyName("key2")
	t.key2 = key2
}
func (t *SysTable) SetS1Data(s1Data sql.NullString) {
	t.AddPropertyName("s1Data")
	t.s1Data = s1Data
}
func (t *SysTable) SetS2Data(s2Data sql.NullString) {
	t.AddPropertyName("s2Data")
	t.s2Data = s2Data
}
func (t *SysTable) SetS3Data(s3Data sql.NullString) {
	t.AddPropertyName("s3Data")
	t.s3Data = s3Data
}
func (t *SysTable) SetN1Data(n1Data df.NullNumeric) {
	t.AddPropertyName("n1Data")
	t.n1Data = n1Data
}
func (t *SysTable) SetN2Data(n2Data df.NullNumeric) {
	t.AddPropertyName("n2Data")
	t.n2Data = n2Data
}
func (t *SysTable) SetN3Data(n3Data df.NullNumeric) {
	t.AddPropertyName("n3Data")
	t.n3Data = n3Data
}
func (t *SysTable) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}
func (t *SysTable) SetDelFlag(delFlag int64) {
	t.AddPropertyName("delFlag")
	t.delFlag = delFlag
}
func (t *SysTable) SetRegisterDatetime(registerDatetime df.Timestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *SysTable) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *SysTable) SetRegisterProcess(registerProcess string) {
	t.AddPropertyName("registerProcess")
	t.registerProcess = registerProcess
}
func (t *SysTable) SetUpdateDatetime(updateDatetime df.Timestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *SysTable) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *SysTable) SetUpdateProcess(updateProcess string) {
	t.AddPropertyName("updateProcess")
	t.updateProcess = updateProcess
}
func (t *SysTable) SetUp(){
	t.n1Data.DecPoint = 2
	t.n2Data.DecPoint = 2
	t.n3Data.DecPoint = 2
	
}
func (t *SysTable)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}

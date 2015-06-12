package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type CustomerDbm_T struct {
	df.BaseDBMeta
	ColumnId *df.ColumnInfo
	ColumnCusCd *df.ColumnInfo
	ColumnName *df.ColumnInfo
	ColumnAddr *df.ColumnInfo
	ColumnBldg *df.ColumnInfo
	ColumnCusConSec *df.ColumnInfo
	ColumnCusConName *df.ColumnInfo
	ColumnTel *df.ColumnInfo
	ColumnSalesAmount *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
	ColumnDelFlag *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
}

func (b *CustomerDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *CustomerDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
}

func (b *CustomerDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var CustomerDbm *CustomerDbm_T

func Create_CustomerDbm() {
	CustomerDbm = new(CustomerDbm_T)
	CustomerDbm.TableDbName = "customer"
	CustomerDbm.TableDispName = "customer"
	CustomerDbm.TablePropertyName = "customer"
	CustomerDbm.TableSqlName = new(df.TableSqlName)
	CustomerDbm.TableSqlName.TableSqlName = "customer"
	CustomerDbm.TableSqlName.CorrespondingDbName = CustomerDbm.TableDbName

	var customer df.DBMeta
	customer = CustomerDbm
	CustomerDbm.DBMeta=&customer
	idSqlName := new(df.ColumnSqlName)
	idSqlName.ColumnSqlName = "id"
	idSqlName.IrregularChar = false
	CustomerDbm.ColumnId = df.CCI(&customer, "id", idSqlName, "", "",
               "Integer.class", "id", "", true, true,true, "serial", 10, 0,
               "nextval('customer_id_seq'::regclass)",false,"","", "","salesSlipList","",false,"int64")
	cusCdSqlName := new(df.ColumnSqlName)
	cusCdSqlName.ColumnSqlName = "cus_cd"
	cusCdSqlName.IrregularChar = false
	CustomerDbm.ColumnCusCd = df.CCI(&customer, "cus_cd", cusCdSqlName, "", "",
               "String.class", "cusCd", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	nameSqlName := new(df.ColumnSqlName)
	nameSqlName.ColumnSqlName = "name"
	nameSqlName.IrregularChar = false
	CustomerDbm.ColumnName = df.CCI(&customer, "name", nameSqlName, "", "",
               "String.class", "name", "", false, false,true, "varchar", 100, 0,
               "",false,"","", "","","",false,"string")
	addrSqlName := new(df.ColumnSqlName)
	addrSqlName.ColumnSqlName = "addr"
	addrSqlName.IrregularChar = false
	CustomerDbm.ColumnAddr = df.CCI(&customer, "addr", addrSqlName, "", "",
               "String.class", "addr", "", false, false,true, "varchar", 100, 0,
               "",false,"","", "","","",false,"string")
	bldgSqlName := new(df.ColumnSqlName)
	bldgSqlName.ColumnSqlName = "bldg"
	bldgSqlName.IrregularChar = false
	CustomerDbm.ColumnBldg = df.CCI(&customer, "bldg", bldgSqlName, "", "",
               "String.class", "bldg", "", false, false,true, "varchar", 100, 0,
               "",false,"","", "","","",false,"string")
	cusConSecSqlName := new(df.ColumnSqlName)
	cusConSecSqlName.ColumnSqlName = "cus_con_sec"
	cusConSecSqlName.IrregularChar = false
	CustomerDbm.ColumnCusConSec = df.CCI(&customer, "cus_con_sec", cusConSecSqlName, "", "",
               "String.class", "cusConSec", "", false, false,true, "varchar", 100, 0,
               "",false,"","", "","","",false,"string")
	cusConNameSqlName := new(df.ColumnSqlName)
	cusConNameSqlName.ColumnSqlName = "cus_con_name"
	cusConNameSqlName.IrregularChar = false
	CustomerDbm.ColumnCusConName = df.CCI(&customer, "cus_con_name", cusConNameSqlName, "", "",
               "String.class", "cusConName", "", false, false,true, "varchar", 100, 0,
               "",false,"","", "","","",false,"string")
	telSqlName := new(df.ColumnSqlName)
	telSqlName.ColumnSqlName = "tel"
	telSqlName.IrregularChar = false
	CustomerDbm.ColumnTel = df.CCI(&customer, "tel", telSqlName, "", "",
               "String.class", "tel", "", false, false,true, "varchar", 100, 0,
               "",false,"","", "","","",false,"string")
	salesAmountSqlName := new(df.ColumnSqlName)
	salesAmountSqlName.ColumnSqlName = "sales_amount"
	salesAmountSqlName.IrregularChar = false
	CustomerDbm.ColumnSalesAmount = df.CCI(&customer, "sales_amount", salesAmountSqlName, "", "",
               "Long.class", "salesAmount", "", false, false,true, "int8", 19, 0,
               "",false,"","", "","","",false,"int64")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	CustomerDbm.ColumnVersionNo = df.CCI(&customer, "version_no", versionNoSqlName, "", "",
               "Integer.class", "versionNo", "", false, false,true, "int4", 10, 0,
               "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")
	delFlagSqlName := new(df.ColumnSqlName)
	delFlagSqlName.ColumnSqlName = "del_flag"
	delFlagSqlName.IrregularChar = false
	CustomerDbm.ColumnDelFlag = df.CCI(&customer, "del_flag", delFlagSqlName, "", "",
               "Integer.class", "delFlag", "", false, false,true, "int4", 10, 0,
               "0",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	CustomerDbm.ColumnRegisterDatetime = df.CCI(&customer, "register_datetime", registerDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	CustomerDbm.ColumnRegisterUser = df.CCI(&customer, "register_user", registerUserSqlName, "", "",
               "String.class", "registerUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	CustomerDbm.ColumnRegisterProcess = df.CCI(&customer, "register_process", registerProcessSqlName, "", "",
               "String.class", "registerProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	CustomerDbm.ColumnUpdateDatetime = df.CCI(&customer, "update_datetime", updateDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	CustomerDbm.ColumnUpdateUser = df.CCI(&customer, "update_user", updateUserSqlName, "", "",
               "String.class", "updateUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	CustomerDbm.ColumnUpdateProcess = df.CCI(&customer, "update_process", updateProcessSqlName, "", "",
               "String.class", "updateProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")

	CustomerDbm.ColumnInfoList = new(df.List)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnId)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnCusCd)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnName)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnAddr)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnBldg)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnCusConSec)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnCusConName)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnTel)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnSalesAmount)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnVersionNo)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnDelFlag)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnRegisterDatetime)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnRegisterUser)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnRegisterProcess)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnUpdateDatetime)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnUpdateUser)
	CustomerDbm.ColumnInfoList.Add(CustomerDbm.ColumnUpdateProcess)


	CustomerDbm.ColumnInfoMap=make(map[string]int)
	CustomerDbm.ColumnInfoMap["id"]=0
		CustomerDbm.ColumnInfoMap["cusCd"]=1
		CustomerDbm.ColumnInfoMap["name"]=2
		CustomerDbm.ColumnInfoMap["addr"]=3
		CustomerDbm.ColumnInfoMap["bldg"]=4
		CustomerDbm.ColumnInfoMap["cusConSec"]=5
		CustomerDbm.ColumnInfoMap["cusConName"]=6
		CustomerDbm.ColumnInfoMap["tel"]=7
		CustomerDbm.ColumnInfoMap["salesAmount"]=8
		CustomerDbm.ColumnInfoMap["versionNo"]=9
		CustomerDbm.ColumnInfoMap["delFlag"]=10
		CustomerDbm.ColumnInfoMap["registerDatetime"]=11
		CustomerDbm.ColumnInfoMap["registerUser"]=12
		CustomerDbm.ColumnInfoMap["registerProcess"]=13
		CustomerDbm.ColumnInfoMap["updateDatetime"]=14
		CustomerDbm.ColumnInfoMap["updateUser"]=15
		CustomerDbm.ColumnInfoMap["updateProcess"]=16
	    CustomerDbm.PrimaryKey = true
    CustomerDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &customer
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(CustomerDbm.ColumnId)

	CustomerDbm.PrimaryInfo = new(df.PrimaryInfo)
	CustomerDbm.PrimaryInfo.UniqueInfo = ui
	
	CustomerDbm.SequenceFlag = true
	CustomerDbm.SequenceName = "customer_id_seq"
	CustomerDbm.SequenceIncrementSize = 1
	CustomerDbm.SequenceCacheSize = 0
	CustomerDbm.VersionNoFlag = true
	CustomerDbm.VersionNoColumnInfo = CustomerDbm.ColumnVersionNo
	
	var customerMeta df.DBMeta = CustomerDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Customer"] = &customerMeta
}

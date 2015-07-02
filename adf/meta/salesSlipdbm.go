package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type SalesSlipDbm_T struct {
	df.BaseDBMeta
	ColumnId *df.ColumnInfo
	ColumnSlipNo *df.ColumnInfo
	ColumnLine *df.ColumnInfo
	ColumnSalesDate *df.ColumnInfo
	ColumnSalesId *df.ColumnInfo
	ColumnCusId *df.ColumnInfo
	ColumnPrdId *df.ColumnInfo
	ColumnQty *df.ColumnInfo
	ColumnUnit *df.ColumnInfo
	ColumnUnitPrice *df.ColumnInfo
	ColumnSalesAmount *df.ColumnInfo
	ColumnSlipAmount *df.ColumnInfo
	ColumnSlipCons *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
	ColumnDelFlag *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
}

func (b *SalesSlipDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *SalesSlipDbm_T) foreignProduct() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		SalesSlipDbm.GetColumnInfoByPropertyName("prdId"),
		ProductDbm.GetColumnInfoByPropertyName("id"),
	}

	return b.BaseDBMeta.Cfi("sales_slip_product_fkey", "Product",
		columns, 0, false, false, false, false,
		"", nil, false, "salesSlipList")
}	
func (b *SalesSlipDbm_T) foreignEmployee() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		SalesSlipDbm.GetColumnInfoByPropertyName("salesId"),
		EmployeeDbm.GetColumnInfoByPropertyName("id"),
	}

	return b.BaseDBMeta.Cfi("sales_slip_sales_id_fkey", "Employee",
		columns, 1, false, false, false, false,
		"", nil, false, "salesSlipList")
}	
func (b *SalesSlipDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
	b.ForeignInfoMap["Product"] = b.foreignProduct()
	b.ForeignInfoMap["Employee"] = b.foreignEmployee()
}

func (b *SalesSlipDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var SalesSlipDbm *SalesSlipDbm_T

func Create_SalesSlipDbm() {
	SalesSlipDbm = new(SalesSlipDbm_T)
	SalesSlipDbm.TableDbName = "sales_slip"
	SalesSlipDbm.TableDispName = "sales_slip"
	SalesSlipDbm.TablePropertyName = "salesSlip"
	SalesSlipDbm.TableSqlName = new(df.TableSqlName)
	SalesSlipDbm.TableSqlName.TableSqlName = "sales_slip"
	SalesSlipDbm.TableSqlName.CorrespondingDbName = SalesSlipDbm.TableDbName

	var salesSlip df.DBMeta
	salesSlip = SalesSlipDbm
	SalesSlipDbm.DBMeta=&salesSlip
	idSqlName := new(df.ColumnSqlName)
	idSqlName.ColumnSqlName = "id"
	idSqlName.IrregularChar = false
	SalesSlipDbm.ColumnId = df.CCI(&salesSlip, "id", idSqlName, "", "",
               "Integer.class", "id", "", true, true,true, "serial", 10, 0,
               "nextval('sales_slip_id_seq'::regclass)",false,"","", "","","",false,"int64")
	slipNoSqlName := new(df.ColumnSqlName)
	slipNoSqlName.ColumnSqlName = "slip_no"
	slipNoSqlName.IrregularChar = false
	SalesSlipDbm.ColumnSlipNo = df.CCI(&salesSlip, "slip_no", slipNoSqlName, "", "",
               "Integer.class", "slipNo", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	lineSqlName := new(df.ColumnSqlName)
	lineSqlName.ColumnSqlName = "line"
	lineSqlName.IrregularChar = false
	SalesSlipDbm.ColumnLine = df.CCI(&salesSlip, "line", lineSqlName, "", "",
               "Integer.class", "line", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	salesDateSqlName := new(df.ColumnSqlName)
	salesDateSqlName.ColumnSqlName = "sales_date"
	salesDateSqlName.IrregularChar = false
	SalesSlipDbm.ColumnSalesDate = df.CCI(&salesSlip, "sales_date", salesDateSqlName, "", "",
               "java.time.LocalDate.class", "salesDate", "", false, false,true, "date", 13, 0,
               "",false,"","", "","","",false,"df.Date")
	salesIdSqlName := new(df.ColumnSqlName)
	salesIdSqlName.ColumnSqlName = "sales_id"
	salesIdSqlName.IrregularChar = false
	SalesSlipDbm.ColumnSalesId = df.CCI(&salesSlip, "sales_id", salesIdSqlName, "", "",
               "Integer.class", "salesId", "", false, false,true, "int4", 10, 0,
               "",false,"","", "employee","","",false,"int64")
	cusIdSqlName := new(df.ColumnSqlName)
	cusIdSqlName.ColumnSqlName = "cus_id"
	cusIdSqlName.IrregularChar = false
	SalesSlipDbm.ColumnCusId = df.CCI(&salesSlip, "cus_id", cusIdSqlName, "", "",
               "Integer.class", "cusId", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	prdIdSqlName := new(df.ColumnSqlName)
	prdIdSqlName.ColumnSqlName = "prd_id"
	prdIdSqlName.IrregularChar = false
	SalesSlipDbm.ColumnPrdId = df.CCI(&salesSlip, "prd_id", prdIdSqlName, "", "",
               "Integer.class", "prdId", "", false, false,true, "int4", 10, 0,
               "",false,"","", "product","","",false,"int64")
	qtySqlName := new(df.ColumnSqlName)
	qtySqlName.ColumnSqlName = "qty"
	qtySqlName.IrregularChar = false
	SalesSlipDbm.ColumnQty = df.CCI(&salesSlip, "qty", qtySqlName, "", "",
               "Integer.class", "qty", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	unitSqlName := new(df.ColumnSqlName)
	unitSqlName.ColumnSqlName = "unit"
	unitSqlName.IrregularChar = false
	SalesSlipDbm.ColumnUnit = df.CCI(&salesSlip, "unit", unitSqlName, "", "",
               "String.class", "unit", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	unitPriceSqlName := new(df.ColumnSqlName)
	unitPriceSqlName.ColumnSqlName = "unit_price"
	unitPriceSqlName.IrregularChar = false
	SalesSlipDbm.ColumnUnitPrice = df.CCI(&salesSlip, "unit_price", unitPriceSqlName, "", "",
               "Integer.class", "unitPrice", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	salesAmountSqlName := new(df.ColumnSqlName)
	salesAmountSqlName.ColumnSqlName = "sales_amount"
	salesAmountSqlName.IrregularChar = false
	SalesSlipDbm.ColumnSalesAmount = df.CCI(&salesSlip, "sales_amount", salesAmountSqlName, "", "",
               "Long.class", "salesAmount", "", false, false,true, "int8", 19, 0,
               "",false,"","", "","","",false,"int64")
	slipAmountSqlName := new(df.ColumnSqlName)
	slipAmountSqlName.ColumnSqlName = "slip_amount"
	slipAmountSqlName.IrregularChar = false
	SalesSlipDbm.ColumnSlipAmount = df.CCI(&salesSlip, "slip_amount", slipAmountSqlName, "", "",
               "Long.class", "slipAmount", "", false, false,true, "int8", 19, 0,
               "",false,"","", "","","",false,"int64")
	slipConsSqlName := new(df.ColumnSqlName)
	slipConsSqlName.ColumnSqlName = "slip_cons"
	slipConsSqlName.IrregularChar = false
	SalesSlipDbm.ColumnSlipCons = df.CCI(&salesSlip, "slip_cons", slipConsSqlName, "", "",
               "Long.class", "slipCons", "", false, false,true, "int8", 19, 0,
               "",false,"","", "","","",false,"int64")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	SalesSlipDbm.ColumnVersionNo = df.CCI(&salesSlip, "version_no", versionNoSqlName, "", "",
               "Integer.class", "versionNo", "", false, false,true, "int4", 10, 0,
               "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")
	delFlagSqlName := new(df.ColumnSqlName)
	delFlagSqlName.ColumnSqlName = "del_flag"
	delFlagSqlName.IrregularChar = false
	SalesSlipDbm.ColumnDelFlag = df.CCI(&salesSlip, "del_flag", delFlagSqlName, "", "",
               "Integer.class", "delFlag", "", false, false,true, "int4", 10, 0,
               "0",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	SalesSlipDbm.ColumnRegisterDatetime = df.CCI(&salesSlip, "register_datetime", registerDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	SalesSlipDbm.ColumnRegisterUser = df.CCI(&salesSlip, "register_user", registerUserSqlName, "", "",
               "String.class", "registerUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	SalesSlipDbm.ColumnRegisterProcess = df.CCI(&salesSlip, "register_process", registerProcessSqlName, "", "",
               "String.class", "registerProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	SalesSlipDbm.ColumnUpdateDatetime = df.CCI(&salesSlip, "update_datetime", updateDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	SalesSlipDbm.ColumnUpdateUser = df.CCI(&salesSlip, "update_user", updateUserSqlName, "", "",
               "String.class", "updateUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	SalesSlipDbm.ColumnUpdateProcess = df.CCI(&salesSlip, "update_process", updateProcessSqlName, "", "",
               "String.class", "updateProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")

	SalesSlipDbm.ColumnInfoList = new(df.List)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnId)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnSlipNo)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnLine)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnSalesDate)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnSalesId)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnCusId)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnPrdId)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnQty)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnUnit)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnUnitPrice)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnSalesAmount)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnSlipAmount)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnSlipCons)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnVersionNo)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnDelFlag)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnRegisterDatetime)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnRegisterUser)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnRegisterProcess)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnUpdateDatetime)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnUpdateUser)
	SalesSlipDbm.ColumnInfoList.Add(SalesSlipDbm.ColumnUpdateProcess)


	SalesSlipDbm.ColumnInfoMap=make(map[string]int)
	SalesSlipDbm.ColumnInfoMap["id"]=0
		SalesSlipDbm.ColumnInfoMap["slipNo"]=1
		SalesSlipDbm.ColumnInfoMap["line"]=2
		SalesSlipDbm.ColumnInfoMap["salesDate"]=3
		SalesSlipDbm.ColumnInfoMap["salesId"]=4
		SalesSlipDbm.ColumnInfoMap["cusId"]=5
		SalesSlipDbm.ColumnInfoMap["prdId"]=6
		SalesSlipDbm.ColumnInfoMap["qty"]=7
		SalesSlipDbm.ColumnInfoMap["unit"]=8
		SalesSlipDbm.ColumnInfoMap["unitPrice"]=9
		SalesSlipDbm.ColumnInfoMap["salesAmount"]=10
		SalesSlipDbm.ColumnInfoMap["slipAmount"]=11
		SalesSlipDbm.ColumnInfoMap["slipCons"]=12
		SalesSlipDbm.ColumnInfoMap["versionNo"]=13
		SalesSlipDbm.ColumnInfoMap["delFlag"]=14
		SalesSlipDbm.ColumnInfoMap["registerDatetime"]=15
		SalesSlipDbm.ColumnInfoMap["registerUser"]=16
		SalesSlipDbm.ColumnInfoMap["registerProcess"]=17
		SalesSlipDbm.ColumnInfoMap["updateDatetime"]=18
		SalesSlipDbm.ColumnInfoMap["updateUser"]=19
		SalesSlipDbm.ColumnInfoMap["updateProcess"]=20
	    SalesSlipDbm.PrimaryKey = true
    SalesSlipDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &salesSlip
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(SalesSlipDbm.ColumnId)

	SalesSlipDbm.PrimaryInfo = new(df.PrimaryInfo)
	SalesSlipDbm.PrimaryInfo.UniqueInfo = ui
	
	SalesSlipDbm.SequenceFlag = true
	SalesSlipDbm.SequenceName = "sales_slip_id_seq"
	SalesSlipDbm.SequenceIncrementSize = 1
	SalesSlipDbm.SequenceCacheSize = 0
	SalesSlipDbm.VersionNoFlag = true
	SalesSlipDbm.VersionNoColumnInfo = SalesSlipDbm.ColumnVersionNo
	
	var salesSlipMeta df.DBMeta = SalesSlipDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["SalesSlip"] = &salesSlipMeta
}

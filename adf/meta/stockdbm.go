package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type StockDbm_T struct {
	df.BaseDBMeta
	ColumnId *df.ColumnInfo
	ColumnYyyymm *df.ColumnInfo
	ColumnPrdId *df.ColumnInfo
	ColumnName *df.ColumnInfo
	ColumnCost *df.ColumnInfo
	ColumnPrice *df.ColumnInfo
	ColumnUnit *df.ColumnInfo
	ColumnBeforeQty *df.ColumnInfo
	ColumnPurQty *df.ColumnInfo
	ColumnRcvQty *df.ColumnInfo
	ColumnSalesQty *df.ColumnInfo
	ColumnOutQty *df.ColumnInfo
	ColumnStockQty *df.ColumnInfo
	ColumnStockAmount *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
	ColumnDelFlag *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
}

func (b *StockDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *StockDbm_T) foreignProduct() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		StockDbm.GetColumnInfoByPropertyName("prdId"),
		ProductDbm.GetColumnInfoByPropertyName("id"),
	}

	return b.BaseDBMeta.Cfi("stock_product_fkey", "Product",
		columns, 0, false, false, false, false,
		"", nil, false, "stockList")
}	
func (b *StockDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
	b.ForeignInfoMap["Product"] = b.foreignProduct()
}

func (b *StockDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var StockDbm *StockDbm_T

func Create_StockDbm() {
	StockDbm = new(StockDbm_T)
	StockDbm.TableDbName = "stock"
	StockDbm.TableDispName = "stock"
	StockDbm.TablePropertyName = "stock"
	StockDbm.TableSqlName = new(df.TableSqlName)
	StockDbm.TableSqlName.TableSqlName = "stock"
	StockDbm.TableSqlName.CorrespondingDbName = StockDbm.TableDbName

	var stock df.DBMeta
	stock = StockDbm
	StockDbm.DBMeta=&stock
	idSqlName := new(df.ColumnSqlName)
	idSqlName.ColumnSqlName = "id"
	idSqlName.IrregularChar = false
	StockDbm.ColumnId = df.CCI(&stock, "id", idSqlName, "", "",
               "Integer.class", "id", "", true, true,true, "serial", 10, 0,
               "nextval('stock_id_seq'::regclass)",false,"","", "","","",false,"int64")
	yyyymmSqlName := new(df.ColumnSqlName)
	yyyymmSqlName.ColumnSqlName = "yyyymm"
	yyyymmSqlName.IrregularChar = false
	StockDbm.ColumnYyyymm = df.CCI(&stock, "yyyymm", yyyymmSqlName, "", "",
               "String.class", "yyyymm", "", false, false,true, "varchar", 6, 0,
               "",false,"","", "","","",false,"string")
	prdIdSqlName := new(df.ColumnSqlName)
	prdIdSqlName.ColumnSqlName = "prd_id"
	prdIdSqlName.IrregularChar = false
	StockDbm.ColumnPrdId = df.CCI(&stock, "prd_id", prdIdSqlName, "", "",
               "Integer.class", "prdId", "", false, false,true, "int4", 10, 0,
               "",false,"","", "product","","",false,"int64")
	nameSqlName := new(df.ColumnSqlName)
	nameSqlName.ColumnSqlName = "name"
	nameSqlName.IrregularChar = false
	StockDbm.ColumnName = df.CCI(&stock, "name", nameSqlName, "", "",
               "String.class", "name", "", false, false,true, "varchar", 100, 0,
               "",false,"","", "","","",false,"string")
	costSqlName := new(df.ColumnSqlName)
	costSqlName.ColumnSqlName = "cost"
	costSqlName.IrregularChar = false
	StockDbm.ColumnCost = df.CCI(&stock, "cost", costSqlName, "", "",
               "Integer.class", "cost", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	priceSqlName := new(df.ColumnSqlName)
	priceSqlName.ColumnSqlName = "price"
	priceSqlName.IrregularChar = false
	StockDbm.ColumnPrice = df.CCI(&stock, "price", priceSqlName, "", "",
               "Integer.class", "price", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	unitSqlName := new(df.ColumnSqlName)
	unitSqlName.ColumnSqlName = "unit"
	unitSqlName.IrregularChar = false
	StockDbm.ColumnUnit = df.CCI(&stock, "unit", unitSqlName, "", "",
               "String.class", "unit", "", false, false,true, "varchar", 100, 0,
               "",false,"","", "","","",false,"string")
	beforeQtySqlName := new(df.ColumnSqlName)
	beforeQtySqlName.ColumnSqlName = "before_qty"
	beforeQtySqlName.IrregularChar = false
	StockDbm.ColumnBeforeQty = df.CCI(&stock, "before_qty", beforeQtySqlName, "", "",
               "Integer.class", "beforeQty", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	purQtySqlName := new(df.ColumnSqlName)
	purQtySqlName.ColumnSqlName = "pur_qty"
	purQtySqlName.IrregularChar = false
	StockDbm.ColumnPurQty = df.CCI(&stock, "pur_qty", purQtySqlName, "", "",
               "Integer.class", "purQty", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	rcvQtySqlName := new(df.ColumnSqlName)
	rcvQtySqlName.ColumnSqlName = "rcv_qty"
	rcvQtySqlName.IrregularChar = false
	StockDbm.ColumnRcvQty = df.CCI(&stock, "rcv_qty", rcvQtySqlName, "", "",
               "Integer.class", "rcvQty", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	salesQtySqlName := new(df.ColumnSqlName)
	salesQtySqlName.ColumnSqlName = "sales_qty"
	salesQtySqlName.IrregularChar = false
	StockDbm.ColumnSalesQty = df.CCI(&stock, "sales_qty", salesQtySqlName, "", "",
               "Integer.class", "salesQty", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	outQtySqlName := new(df.ColumnSqlName)
	outQtySqlName.ColumnSqlName = "out_qty"
	outQtySqlName.IrregularChar = false
	StockDbm.ColumnOutQty = df.CCI(&stock, "out_qty", outQtySqlName, "", "",
               "Integer.class", "outQty", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	stockQtySqlName := new(df.ColumnSqlName)
	stockQtySqlName.ColumnSqlName = "stock_qty"
	stockQtySqlName.IrregularChar = false
	StockDbm.ColumnStockQty = df.CCI(&stock, "stock_qty", stockQtySqlName, "", "",
               "Integer.class", "stockQty", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	stockAmountSqlName := new(df.ColumnSqlName)
	stockAmountSqlName.ColumnSqlName = "stock_amount"
	stockAmountSqlName.IrregularChar = false
	StockDbm.ColumnStockAmount = df.CCI(&stock, "stock_amount", stockAmountSqlName, "", "",
               "Long.class", "stockAmount", "", false, false,true, "int8", 19, 0,
               "",false,"","", "","","",false,"int64")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	StockDbm.ColumnVersionNo = df.CCI(&stock, "version_no", versionNoSqlName, "", "",
               "Integer.class", "versionNo", "", false, false,true, "int4", 10, 0,
               "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")
	delFlagSqlName := new(df.ColumnSqlName)
	delFlagSqlName.ColumnSqlName = "del_flag"
	delFlagSqlName.IrregularChar = false
	StockDbm.ColumnDelFlag = df.CCI(&stock, "del_flag", delFlagSqlName, "", "",
               "Integer.class", "delFlag", "", false, false,true, "int4", 10, 0,
               "0",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	StockDbm.ColumnRegisterDatetime = df.CCI(&stock, "register_datetime", registerDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	StockDbm.ColumnRegisterUser = df.CCI(&stock, "register_user", registerUserSqlName, "", "",
               "String.class", "registerUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	StockDbm.ColumnRegisterProcess = df.CCI(&stock, "register_process", registerProcessSqlName, "", "",
               "String.class", "registerProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	StockDbm.ColumnUpdateDatetime = df.CCI(&stock, "update_datetime", updateDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	StockDbm.ColumnUpdateUser = df.CCI(&stock, "update_user", updateUserSqlName, "", "",
               "String.class", "updateUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	StockDbm.ColumnUpdateProcess = df.CCI(&stock, "update_process", updateProcessSqlName, "", "",
               "String.class", "updateProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")

	StockDbm.ColumnInfoList = new(df.List)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnId)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnYyyymm)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnPrdId)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnName)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnCost)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnPrice)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnUnit)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnBeforeQty)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnPurQty)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnRcvQty)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnSalesQty)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnOutQty)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnStockQty)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnStockAmount)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnVersionNo)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnDelFlag)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnRegisterDatetime)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnRegisterUser)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnRegisterProcess)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnUpdateDatetime)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnUpdateUser)
	StockDbm.ColumnInfoList.Add(StockDbm.ColumnUpdateProcess)


	StockDbm.ColumnInfoMap=make(map[string]int)
	StockDbm.ColumnInfoMap["id"]=0
		StockDbm.ColumnInfoMap["yyyymm"]=1
		StockDbm.ColumnInfoMap["prdId"]=2
		StockDbm.ColumnInfoMap["name"]=3
		StockDbm.ColumnInfoMap["cost"]=4
		StockDbm.ColumnInfoMap["price"]=5
		StockDbm.ColumnInfoMap["unit"]=6
		StockDbm.ColumnInfoMap["beforeQty"]=7
		StockDbm.ColumnInfoMap["purQty"]=8
		StockDbm.ColumnInfoMap["rcvQty"]=9
		StockDbm.ColumnInfoMap["salesQty"]=10
		StockDbm.ColumnInfoMap["outQty"]=11
		StockDbm.ColumnInfoMap["stockQty"]=12
		StockDbm.ColumnInfoMap["stockAmount"]=13
		StockDbm.ColumnInfoMap["versionNo"]=14
		StockDbm.ColumnInfoMap["delFlag"]=15
		StockDbm.ColumnInfoMap["registerDatetime"]=16
		StockDbm.ColumnInfoMap["registerUser"]=17
		StockDbm.ColumnInfoMap["registerProcess"]=18
		StockDbm.ColumnInfoMap["updateDatetime"]=19
		StockDbm.ColumnInfoMap["updateUser"]=20
		StockDbm.ColumnInfoMap["updateProcess"]=21
	    StockDbm.PrimaryKey = true
    StockDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &stock
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(StockDbm.ColumnId)

	StockDbm.PrimaryInfo = new(df.PrimaryInfo)
	StockDbm.PrimaryInfo.UniqueInfo = ui
	
	StockDbm.SequenceFlag = true
	StockDbm.SequenceName = "stock_id_seq"
	StockDbm.SequenceIncrementSize = 1
	StockDbm.SequenceCacheSize = 0
	StockDbm.VersionNoFlag = true
	StockDbm.VersionNoColumnInfo = StockDbm.ColumnVersionNo
	
	var stockMeta df.DBMeta = StockDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Stock"] = &stockMeta
}

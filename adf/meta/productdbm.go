package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type ProductDbm_T struct {
	df.BaseDBMeta
	ColumnId *df.ColumnInfo
	ColumnPrdCd *df.ColumnInfo
	ColumnPrdCat *df.ColumnInfo
	ColumnName *df.ColumnInfo
	ColumnCost *df.ColumnInfo
	ColumnPrice *df.ColumnInfo
	ColumnPerQty *df.ColumnInfo
	ColumnUnit *df.ColumnInfo
	ColumnPurQty *df.ColumnInfo
	ColumnPurAmount *df.ColumnInfo
	ColumnSalesQty *df.ColumnInfo
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

func (b *ProductDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *ProductDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
}

func (b *ProductDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var ProductDbm *ProductDbm_T

func Create_ProductDbm() {
	ProductDbm = new(ProductDbm_T)
	ProductDbm.TableDbName = "product"
	ProductDbm.TableDispName = "product"
	ProductDbm.TablePropertyName = "product"
	ProductDbm.TableSqlName = new(df.TableSqlName)
	ProductDbm.TableSqlName.TableSqlName = "product"
	ProductDbm.TableSqlName.CorrespondingDbName = ProductDbm.TableDbName

	var product df.DBMeta
	product = ProductDbm
	ProductDbm.DBMeta=&product
	idSqlName := new(df.ColumnSqlName)
	idSqlName.ColumnSqlName = "id"
	idSqlName.IrregularChar = false
	ProductDbm.ColumnId = df.CCI(&product, "id", idSqlName, "", "",
               "Integer.class", "id", "", true, true,true, "serial", 10, 0,
               "nextval('product_id_seq'::regclass)",false,"","", "","salesSlipList,stockList","",false,"int64")
	prdCdSqlName := new(df.ColumnSqlName)
	prdCdSqlName.ColumnSqlName = "prd_cd"
	prdCdSqlName.IrregularChar = false
	ProductDbm.ColumnPrdCd = df.CCI(&product, "prd_cd", prdCdSqlName, "", "",
               "String.class", "prdCd", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	prdCatSqlName := new(df.ColumnSqlName)
	prdCatSqlName.ColumnSqlName = "prd_cat"
	prdCatSqlName.IrregularChar = false
	ProductDbm.ColumnPrdCat = df.CCI(&product, "prd_cat", prdCatSqlName, "", "",
               "String.class", "prdCat", "", false, false,true, "varchar", 40, 0,
               "",false,"","", "","","",false,"string")
	nameSqlName := new(df.ColumnSqlName)
	nameSqlName.ColumnSqlName = "name"
	nameSqlName.IrregularChar = false
	ProductDbm.ColumnName = df.CCI(&product, "name", nameSqlName, "", "",
               "String.class", "name", "", false, false,true, "varchar", 100, 0,
               "",false,"","", "","","",false,"string")
	costSqlName := new(df.ColumnSqlName)
	costSqlName.ColumnSqlName = "cost"
	costSqlName.IrregularChar = false
	ProductDbm.ColumnCost = df.CCI(&product, "cost", costSqlName, "", "",
               "Integer.class", "cost", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	priceSqlName := new(df.ColumnSqlName)
	priceSqlName.ColumnSqlName = "price"
	priceSqlName.IrregularChar = false
	ProductDbm.ColumnPrice = df.CCI(&product, "price", priceSqlName, "", "",
               "Integer.class", "price", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	perQtySqlName := new(df.ColumnSqlName)
	perQtySqlName.ColumnSqlName = "per_qty"
	perQtySqlName.IrregularChar = false
	ProductDbm.ColumnPerQty = df.CCI(&product, "per_qty", perQtySqlName, "", "",
               "Integer.class", "perQty", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	unitSqlName := new(df.ColumnSqlName)
	unitSqlName.ColumnSqlName = "unit"
	unitSqlName.IrregularChar = false
	ProductDbm.ColumnUnit = df.CCI(&product, "unit", unitSqlName, "", "",
               "String.class", "unit", "", false, false,true, "varchar", 100, 0,
               "",false,"","", "","","",false,"string")
	purQtySqlName := new(df.ColumnSqlName)
	purQtySqlName.ColumnSqlName = "pur_qty"
	purQtySqlName.IrregularChar = false
	ProductDbm.ColumnPurQty = df.CCI(&product, "pur_qty", purQtySqlName, "", "",
               "Integer.class", "purQty", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	purAmountSqlName := new(df.ColumnSqlName)
	purAmountSqlName.ColumnSqlName = "pur_amount"
	purAmountSqlName.IrregularChar = false
	ProductDbm.ColumnPurAmount = df.CCI(&product, "pur_amount", purAmountSqlName, "", "",
               "Long.class", "purAmount", "", false, false,true, "int8", 19, 0,
               "",false,"","", "","","",false,"int64")
	salesQtySqlName := new(df.ColumnSqlName)
	salesQtySqlName.ColumnSqlName = "sales_qty"
	salesQtySqlName.IrregularChar = false
	ProductDbm.ColumnSalesQty = df.CCI(&product, "sales_qty", salesQtySqlName, "", "",
               "Integer.class", "salesQty", "", false, false,true, "int4", 10, 0,
               "",false,"","", "","","",false,"int64")
	salesAmountSqlName := new(df.ColumnSqlName)
	salesAmountSqlName.ColumnSqlName = "sales_amount"
	salesAmountSqlName.IrregularChar = false
	ProductDbm.ColumnSalesAmount = df.CCI(&product, "sales_amount", salesAmountSqlName, "", "",
               "Long.class", "salesAmount", "", false, false,true, "int8", 19, 0,
               "",false,"","", "","","",false,"int64")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	ProductDbm.ColumnVersionNo = df.CCI(&product, "version_no", versionNoSqlName, "", "",
               "Integer.class", "versionNo", "", false, false,true, "int4", 10, 0,
               "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")
	delFlagSqlName := new(df.ColumnSqlName)
	delFlagSqlName.ColumnSqlName = "del_flag"
	delFlagSqlName.IrregularChar = false
	ProductDbm.ColumnDelFlag = df.CCI(&product, "del_flag", delFlagSqlName, "", "",
               "Integer.class", "delFlag", "", false, false,true, "int4", 10, 0,
               "0",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	ProductDbm.ColumnRegisterDatetime = df.CCI(&product, "register_datetime", registerDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	ProductDbm.ColumnRegisterUser = df.CCI(&product, "register_user", registerUserSqlName, "", "",
               "String.class", "registerUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	ProductDbm.ColumnRegisterProcess = df.CCI(&product, "register_process", registerProcessSqlName, "", "",
               "String.class", "registerProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	ProductDbm.ColumnUpdateDatetime = df.CCI(&product, "update_datetime", updateDatetimeSqlName, "", "",
               "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 29, 6,
               "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	ProductDbm.ColumnUpdateUser = df.CCI(&product, "update_user", updateUserSqlName, "", "",
               "String.class", "updateUser", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	ProductDbm.ColumnUpdateProcess = df.CCI(&product, "update_process", updateProcessSqlName, "", "",
               "String.class", "updateProcess", "", false, false,true, "varchar", 30, 0,
               "",false,"","", "","","",false,"string")

	ProductDbm.ColumnInfoList = new(df.List)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnId)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnPrdCd)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnPrdCat)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnName)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnCost)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnPrice)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnPerQty)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnUnit)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnPurQty)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnPurAmount)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnSalesQty)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnSalesAmount)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnVersionNo)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnDelFlag)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnRegisterDatetime)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnRegisterUser)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnRegisterProcess)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnUpdateDatetime)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnUpdateUser)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnUpdateProcess)


	ProductDbm.ColumnInfoMap=make(map[string]int)
	ProductDbm.ColumnInfoMap["id"]=0
		ProductDbm.ColumnInfoMap["prdCd"]=1
		ProductDbm.ColumnInfoMap["prdCat"]=2
		ProductDbm.ColumnInfoMap["name"]=3
		ProductDbm.ColumnInfoMap["cost"]=4
		ProductDbm.ColumnInfoMap["price"]=5
		ProductDbm.ColumnInfoMap["perQty"]=6
		ProductDbm.ColumnInfoMap["unit"]=7
		ProductDbm.ColumnInfoMap["purQty"]=8
		ProductDbm.ColumnInfoMap["purAmount"]=9
		ProductDbm.ColumnInfoMap["salesQty"]=10
		ProductDbm.ColumnInfoMap["salesAmount"]=11
		ProductDbm.ColumnInfoMap["versionNo"]=12
		ProductDbm.ColumnInfoMap["delFlag"]=13
		ProductDbm.ColumnInfoMap["registerDatetime"]=14
		ProductDbm.ColumnInfoMap["registerUser"]=15
		ProductDbm.ColumnInfoMap["registerProcess"]=16
		ProductDbm.ColumnInfoMap["updateDatetime"]=17
		ProductDbm.ColumnInfoMap["updateUser"]=18
		ProductDbm.ColumnInfoMap["updateProcess"]=19
	    ProductDbm.PrimaryKey = true
    ProductDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &product
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(ProductDbm.ColumnId)

	ProductDbm.PrimaryInfo = new(df.PrimaryInfo)
	ProductDbm.PrimaryInfo.UniqueInfo = ui
	
	ProductDbm.SequenceFlag = true
	ProductDbm.SequenceName = "product_id_seq"
	ProductDbm.SequenceIncrementSize = 1
	ProductDbm.SequenceCacheSize = 0
	ProductDbm.VersionNoFlag = true
	ProductDbm.VersionNoColumnInfo = ProductDbm.ColumnVersionNo
	
	var productMeta df.DBMeta = ProductDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Product"] = &productMeta
}

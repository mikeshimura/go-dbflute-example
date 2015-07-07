
$c.checkAndCreate("$w")
 
$w.handleChange = (e) ->
  $c.handleChange($w.app,e.target.name,e.target.value);

$w.handleClick = (e) ->
  name=e.target.name
  if name=="alert#CloseBtn"
     $w.flux.actions.$c_alertHide()
  if name=="deleteCfm#CloseBtn"
     $w.flux.actions.$c_deleteCfmHide()
  if name=="deleteCfm#YesBtn"
     $w.flux.actions.$c_deleteCfmHide()
     $w.app.state.common.deleteCfm.callback()
  if name == "btntestTable_New"
    $w.testTable_formClear()
  if name == "btntestTable_Search"
    $w.testTable_formSearch()
  if name == "btntestTable_Update"
    $w.testTable_formUpdate()
  if name == "btntestTable_Delete"
    $w.testTable_formDelete()
  if typeof(e.target.id)=="undefined"
    return
  ids = e.target.id.split("#");
  if (ids[0] == "testTable_row")
    temp={ testTable:$w.app.state.testTable }
    selRow = Number(ids[2])
    temp.testTable.selRow=selRow
    temp.testTable_form=_.cloneDeep(temp.testTable.rcds[selRow])
    $w.app.setState(temp)
$w.testTable_formSearch = () ->
  criteria=$c.createCriteria($w.app.state.testTable_search,["testId","testDate","testTimestamp","testNbr"])
  maxRecord=$w.app.state.testTable_search.maxRecord
  $w.flux.actions.$c_rcd_fetch($w.app.state.testTable,$w.app.state.testTable_form,"testTable",criteria,maxRecord)
$w.testTable_formUpdate = () ->
  $w.flux.actions.$c_rcd_update($w.app.state.testTable,$w.app.state.testTable_form,"testTable")
$w.testTable_formDelete = () ->
  if $w.app.state.testTable_form.id == ""
    $w.flux.actions.$c_rcd_delete_id_blank()
    return
  $w.flux.actions.$c_deleteCfmShow($w.testTable_formDeleteCfm)
$w.testTable_formDeleteCfm = () ->
  $w.flux.actions.$c_rcd_delete($w.app.state.testTable,$w.app.state.testTable_form,"testTable")
$w.testTable_formClear = () ->
  formtemp={
    testTable_form:_.cloneDeep($w.app.state.testTable.blank)
    testTable:$w.app.state.testTable
  }
  formtemp.testTable.selRow=-1
  $w.app.setState(formtemp)
$w.constants =
  $W_LOGIN_SUCCESS: "$W_LOGIN_SUCCESS"


$w.actions = {
} 

$w.PageStore = Fluxxor.createStore(
  initialize: ->
    @data = 
      {
      }
    #@bindActions $w.constants.$W_LOGIN_SUCCESS, @onLoginSuccess,

    return
) 

$w.flux = new Fluxxor.Flux()
$w.pageStore=new $w.PageStore;
$w.flux.addStore("PAGE",$w.pageStore)
$w.flux.addActions($w.actions)
$w.commonStore=new $c.CommonStore;
$w.flux.addStore("COMMON",$w.commonStore)
$w.flux.addActions($c.actions)
$w.rcdStore=new $c.RcdStore;
$w.flux.addStore("RCD",$w.rcdStore)
$w.flux.addActions($c.rcdActions)
rcdStore = $w.flux.store("RCD")
rcdStore.addTable("testTable")
$w.FluxMixin = Fluxxor.FluxMixin(React)
$w.StoreWatchMixin = Fluxxor.StoreWatchMixin
$w.rcdStore.on("rcdComplete_testTable", ->
  rcdTemp=_.cloneDeep($w.app.state.rcd.testTable)
  temp={
    testTable:$w.app.state.testTable
  }
  temp.testTable.rcds=rcdTemp.rcds
  temp.testTable_form=rcdTemp.rcd
  temp.testTable.selRow=rcdTemp.selRow
  $w.app.setState(temp) 
)
$w.getDom = (refname) ->
  return $w.app.refs[refname].getDOMNode()
  
$w.testTable_scroll = ->
  $w.getDom("testTable_tableHead").scrollLeft=$w.getDom("testTable_tableBody").scrollLeft
$w.testTable_onscroll = ->
  $w.getDom("testTable_tableBody").onscroll=$w.testTable_scroll
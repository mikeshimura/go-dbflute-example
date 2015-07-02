
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
  if name == "btncustomer_New"
    $w.customer_formClear()
  if name == "btncustomer_Search"
    $w.customer_formSearch()
  if name == "btncustomer_Update"
    $w.customer_formUpdate()
  if name == "btncustomer_Delete"
    $w.customer_formDelete()
  if typeof(e.target.id)=="undefined"
    return
  ids = e.target.id.split("#");
  if (ids[0] == "customer_row")
    temp={ 
      customer:$w.app.state.customer 
      tabkey:2
    }
    selRow = Number(ids[2])
    temp.customer.selRow=selRow
    temp.customer_form=_.cloneDeep(temp.customer.rcds[selRow])
    $w.app.setState(temp)
$w.handleTabSelect = (tab) ->
  tabkey={
      tabkey:tab
    }
  $w.app.setState(tabkey)
  return
$w.customer_formSearch = () ->
  criteria=$c.createCriteria($w.app.state.customer_search,["cusCd","name"])
  maxRecord=$w.app.state.customer_search.maxRecord
  $w.flux.actions.$c_rcd_fetch($w.app.state.customer,$w.app.state.customer_form,"customer",criteria,maxRecord)
$w.customer_formUpdate = () ->
  $w.flux.actions.$c_rcd_update($w.app.state.customer,$w.app.state.customer_form,"customer")
$w.customer_formDelete = () ->
  if $w.app.state.customer_form.id == ""
    $w.flux.actions.$c_rcd_delete_id_blank()
    return
  $w.flux.actions.$c_deleteCfmShow($w.customer_formDeleteCfm)
$w.customer_formDeleteCfm = () ->
  $w.flux.actions.$c_rcd_delete($w.app.state.customer,$w.app.state.customer_form,"customer")
$w.customer_formClear = () ->
  formtemp={
    customer_form:_.cloneDeep($w.app.state.customer.blank)
    customer:$w.app.state.customer
  }
  formtemp.customer.selRow=-1
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
rcdStore.addTable("customer")
$w.FluxMixin = Fluxxor.FluxMixin(React)
$w.StoreWatchMixin = Fluxxor.StoreWatchMixin
$w.rcdStore.on("rcdComplete_customer", ->
  rcdTemp=_.cloneDeep($w.app.state.rcd.customer)
  temp={
    customer:$w.app.state.customer
  }
  temp.customer.rcds=rcdTemp.rcds
  temp.customer_form=rcdTemp.rcd
  temp.customer.selRow=rcdTemp.selRow
  $w.app.setState(temp) 
)
$w.getDom = (refname) ->
  return $w.app.refs[refname].getDOMNode()
  
$w.customer_scroll = ->
  $w.getDom("customer_tableHead").scrollLeft=$w.getDom("customer_tableBody").scrollLeft
$w.customer_onscroll = ->
  $w.getDom("customer_tableBody").onscroll=$w.customer_scroll
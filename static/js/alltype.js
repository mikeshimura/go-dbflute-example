(function() {
  $c.checkAndCreate("$w");
  $w.handleChange = function(e) {
    var jsx;
    jsx = $w.app;
    if (e.target.type === "checkbox") {
      return $c.handleChange(jsx, e.target.name, e.target.checked);
    } else {
      return $c.handleChange(jsx, e.target.name, e.target.value);
    }
  };
  $w.handleClick = function(e) {
    var ids, jsx, logintemp, name, selRow;
    jsx = $w.app;
    name = e.target.name;
    if (name === "alert#CloseBtn") {
      $w.flux.actions.$c_alertHide();
    }
    if (name === "deleteCfm#CloseBtn") {
      $w.flux.actions.$c_deleteCfmHide();
    }
    if (name === "deleteCfm#YesBtn") {
      $w.flux.actions.$c_deleteCfmHide();
      $w.formDeleteCfm(jsx);
    }
    if (name === "btnNew") {
      $w.formClear(jsx);
    }
    if (name === "btnSearch") {
      $w.formSearch(jsx);
    }
    if (name === "btnUpdate") {
      $w.formUpdate(jsx);
    }
    if (name === "btnDelete") {
      $w.formDelete(jsx);
    }
    if (typeof e.target.id === "undefined") {
      return;
    }
    ids = e.target.id.split("#");
    if (ids[0] === "loginrow") {
      logintemp = {
        login: jsx.state.alltype
      };
      selRow = Number(ids[2]);
      logintemp.login.selRow = selRow;
      logintemp.form = _.cloneDeep(logintemp.login.rcds[selRow]);
      logintemp.form.password = "";
      logintemp.form.passwordcfm = "";
      return jsx.setState(logintemp);
    }
  };
  $w.formSearch = function(jsx) {
    var criteria;
    criteria = $c.createCriteria(jsx.state.search, ["loginId", "name"]);
    return $w.flux.actions.$c_rcd_fetch(jsx.state.alltype, jsx.state.form, "alltype", criteria);
  };
  $w.formUpdate = function(jsx) {
    var form, res, rules;
    form = jsx.state.form;
    res = "";
    if (form.id === "") {
      rules = [];
    } else {
      rules = [];
    }
    if (res.length > 0) {
      $w.flux.actions.$c_alertShow(res);
      return;
    }
    return $w.flux.actions.$c_rcd_update(jsx.state.alltype, jsx.state.form, "alltype");
  };
  $w.formDelete = function(jsx) {
    if (jsx.state.form.id === "") {
      $w.flux.actions.$c_rcd_delete_id_blank();
      return;
    }
    return $w.flux.actions.$c_deleteCfmShow();
  };
  $w.formDeleteCfm = function(jsx) {
    return $w.flux.actions.$c_rcd_delete(jsx.state.alltype, jsx.state.form, "alltype");
  };
  $w.formUpdateCheck = function(form) {
    if (form.password > "" || form.passwordcfm > "") {
      if (form.password !== form.passwordcfm) {
        return [["", "パスワードとパスワード（確認）が一致しません"]];
      }
    }
    return "";
  };
  $w.formClear = function(jsx) {
    var formtemp;
    formtemp = {
      form: _.cloneDeep(jsx.state.alltype.blank)
    };
    return jsx.setState(formtemp);
  };
  $w.constants = {
    $W_LOGIN_SUCCESS: "$W_LOGIN_SUCCESS"
  };
  $w.flux = new Fluxxor.Flux();
  $w.commonStore = new $c.CommonStore;
  $w.flux.addStore("COMMON", $w.commonStore);
  $w.flux.addActions($c.actions);
  $w.rcdStore = new $c.RcdStore;
  $w.flux.addStore("RCD", $w.rcdStore);
  $w.flux.addActions($c.rcdActions);
  $w.FluxMixin = Fluxxor.FluxMixin(React);
  $w.StoreWatchMixin = Fluxxor.StoreWatchMixin;
  $w.common = $w.flux.stores.COMMON;
  $w.rcd = $w.flux.stores.RCD;
  $w.rcd.addTable("alltype");
  $w.rcdStore.on("rcdComplete_alltype", function() {
    var allTypeTemp, rcdAllType;
    rcdAllType = _.cloneDeep($w.app.state.rcd.alltype);
    allTypeTemp = {
      alltype: $w.app.state.alltype
    };
    allTypeTemp.alltype.rcds = rcdAllType.rcds;
    allTypeTemp.form = rcdAllType.rcd;
    allTypeTemp.alltype.selRow = rcdAllType.selRow;
    return $w.app.setState(allTypeTemp);
  });
  $w.getDom = function(refname) {
    return $w.app.refs[refname].getDOMNode();
  };
  $w.scroll = function() {
    return $w.getDom("tableHead").scrollLeft = $w.getDom("tableBody").scrollLeft;
  };
  $w.onscroll = function() {
    return $w.getDom("tableBody").onscroll = $w.scroll;
  };
}).call(this);

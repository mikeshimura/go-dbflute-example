(function() {
  $c.checkAndCreate("$w");
  $w.handleChange = function(e) {
    var name;
    name = e.target.name;
    if (name === "form#db") {
      $w.flux.actions.getDatabases(e.target.value);
    }
    if (name === "form#database") {
      $w.flux.actions.getTables(e.target.value);
    }
    if (name === "form#template") {
      $w.flux.actions.getParameter(e.target.value);
    }
    return $c.handleChange($w.app, e.target.name, e.target.value);
  };
  $w.handleClick = function(e) {
    var name;
    name = e.target.name;
    if (name === "alert#CloseBtn") {
      $w.flux.actions.$c_alertHide();
    }
    if (name === "btnSearch") {
      $w.flux.actions.search();
    }
    if (name === "btnExec") {
      $w.flux.actions.exec();
    }
    if (typeof e.target.id === "undefined") {
      ;
    }
  };
  $w.constants = {
    $W_getdbs_SUCCESS: "W_getdbs_SUCCESS",
    $W_getDatabases_SUCCESS: "W_getDatabases_SUCCESS",
    $W_getTables_SUCCESS: "W_getTables_SUCCESS",
    $W_getColInfo_SUCCESS: "W_getColInfo_SUCCESS",
    $W_getParameter_SUCCESS: "W_getParameter_SUCCESS",
    $W_search_SUCCESS: "W_search_SUCCESS",
    $W_exec_SUCCESS: "W_exec_SUCCESS"
  };
  $w.actions = {
    getdbs: function() {
      var params;
      params = {
        operationType: "getdbs"
      };
      return $c.ajaxPostJson("/ajax/systbl", params, "application/json", $c.ajaxCallback.bind(this, params, $w.constants.$W_getdbs_SUCCESS));
    },
    getDatabases: function(value) {
      var params;
      params = {
        operationType: "getDatabases",
        db: value
      };
      return $c.ajaxPostJson("/ajax/dbflute", params, "application/json", $c.ajaxCallback.bind(this, params, $w.constants.$W_getDatabases_SUCCESS));
    },
    getTables: function(value) {
      var params;
      params = {
        operationType: "getTables",
        database: value,
        db: $w.app.state.form.db
      };
      return $c.ajaxPostJson("/ajax/dbflute", params, "application/json", $c.ajaxCallback.bind(this, params, $w.constants.$W_getTables_SUCCESS));
    },
    getParameter: function(value) {
      var params;
      params = {
        operationType: "getParameter",
        template: value
      };
      return $c.ajaxPostJson("/ajax/template", params, "application/json", $c.ajaxCallback.bind(this, params, $w.constants.$W_getParameter_SUCCESS));
    },
    search: function(value) {
      var params;
      params = {
        operationType: "search",
        tempsel: $w.app.state.form.tempsel
      };
      return $c.ajaxPostJson("/ajax/template", params, "application/json", $c.ajaxCallback.bind(this, params, $w.constants.$W_search_SUCCESS));
    },
    exec: function() {
      var params;
      params = {
        operationType: "exec",
        db: $w.app.state.form.db,
        database: $w.app.state.form.database,
        table: $w.app.state.form.table,
        template: $w.app.state.form.template,
        items: $w.app.state.form.items,
        p1: $w.app.state.form.p1,
        p2: $w.app.state.form.p2,
        p3: $w.app.state.form.p3,
        l1: $w.app.state.form.l1,
        l2: $w.app.state.form.l2,
        l3: $w.app.state.form.l3
      };
      return $c.ajaxPostJson("/ajax/template", params, "application/json", $c.ajaxCallback.bind(this, params, $w.constants.$W_exec_SUCCESS));
    }
  };
  $w.PageStore = Fluxxor.createStore({
    initialize: function() {
      this.data = {
        dbs: [],
        database: [],
        table: [],
        items: [],
        templates: [],
        parameter: "",
        res: ""
      };
      this.bindActions($w.constants.$W_getdbs_SUCCESS, this.getdbsSuccess, $w.constants.$W_getDatabases_SUCCESS, this.getDatabasesSuccess, $w.constants.$W_getTables_SUCCESS, this.getTablesSuccess, $w.constants.$W_getColInfo_SUCCESS, this.getColInfoSuccess, $w.constants.$W_getParameter_SUCCESS, this.getParameterSuccess, $w.constants.$W_search_SUCCESS, this.searchSuccess, $w.constants.$W_exec_SUCCESS, this.execSuccess);
    },
    getdbsSuccess: function(res) {
      this.data.dbs = res.response.data;
      this.data.database = [];
      this.data.table = [];
      this.data.items = [];
      this.emit("change");
    },
    getDatabasesSuccess: function(res) {
      this.data.database = res.response.data;
      this.data.table = [];
      this.data.items = [];
      this.emit("change");
    },
    getTablesSuccess: function(res) {
      this.data.table = res.response.data;
      this.data.items = [];
      this.emit("change");
    },
    getColInfoSuccess: function(res) {
      this.data.items = res.response.data;
      this.emit("change");
      this.emit("colinfo");
    },
    getParameterSuccess: function(res) {
      this.data.parameter = res.response.data;
      this.emit("change");
    },
    searchSuccess: function(res) {
      this.data.templates = res.response.data;
      this.emit("change");
    },
    execSuccess: function(res) {
      this.data.res = res.response.data;
      this.emit("change");
    }
  });
  $w.flux = new Fluxxor.Flux();
  $w.pageStore = new $w.PageStore;
  $w.flux.addStore("PAGE", $w.pageStore);
  $w.flux.addActions($w.actions);
  $w.commonStore = new $c.CommonStore;
  $w.flux.addStore("COMMON", $w.commonStore);
  $w.flux.addActions($c.actions);
  $w.rcdStore = new $c.RcdStore;
  $w.flux.addStore("RCD", $w.rcdStore);
  $w.flux.addActions($c.rcdActions);
  $w.FluxMixin = Fluxxor.FluxMixin(React);
  $w.StoreWatchMixin = Fluxxor.StoreWatchMixin;
  $w.page = $w.flux.stores.PAGE;
  $w.common = $w.flux.stores.COMMON;
  $w.rcd = $w.flux.stores.RCD;
  $w.rcd.addTable("xmlup");
  $w.pageStore.on("colinfo", function() {
    var formtemp, items;
    items = $w.app.state.page.items;
    formtemp = {
      form: _.cloneDeep($w.app.state.form)
    };
    formtemp.form.all = items[0];
    formtemp.form.sel = items[1];
    return $w.app.setState(formtemp);
  });
}).call(this);

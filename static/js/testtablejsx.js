 
  
 var b = ReactBootstrap;
$w.width=900;
$w.height=800;
$w.tableHeight=340;
$w.tableColW={c1:50,c2:90,c3:120,c4:50,c5:50,c6:50}
$w.testTableRows = React.createClass({displayName: 'testTableRows',

    render: function() {
        var rows = this.props.rcds.map(function(rcd, i){
        var bgcolor="";
        if (i==this.props.selRow)
          bgcolor="#d9edf7";
        else
          if (i%2 == 1)
            bgcolor="#F8F8F8";
          else
            bgcolor="#FFFFFF";
        return (
          React.createElement("tr", {key: i}, 
 React.createElement("td", {id: "testTable_row#testId#"+i, 
            style: {width:this.props.cw.c1,backgroundColor:bgcolor,textAlign:"left"}}, rcd.testId), 
 React.createElement("td", {id: "testTable_row#testDate#"+i, 
            style: {width:this.props.cw.c2,backgroundColor:bgcolor,textAlign:"left"}}, rcd.testDate), 
 React.createElement("td", {id: "testTable_row#testTimestamp#"+i, 
            style: {width:this.props.cw.c3,backgroundColor:bgcolor,textAlign:"left"}}, rcd.testTimestamp), 
 React.createElement("td", {id: "testTable_row#testNbr#"+i, 
            style: {width:this.props.cw.c4,backgroundColor:bgcolor,textAlign:"right"}}, rcd.testNbr), 
 React.createElement("td", {id: "testTable_row#id#"+i, 
            style: {width:this.props.cw.c5,backgroundColor:bgcolor,textAlign:"right"}}, rcd.id), 
 React.createElement("td", {id: "testTable_row#versionNo#"+i, 
            style: {width:this.props.cw.c6,backgroundColor:bgcolor,textAlign:"right"}}, rcd.versionNo)

         )
        )
        }, this);
          return (
            React.createElement("tbody", null, 
                rows
            )
          );
    }
  });
$w.Application = React.createClass({displayName: 'Application',
  mixins: [$w.FluxMixin, $w.StoreWatchMixin("PAGE","COMMON","RCD")],
  getInitialState: function() {
  $w.app = this;
  testTable_blank={
          testId:"",
          testDate:"",
          testTimestamp:"",
          testNbr:"",
          id:"",
          versionNo:"",

  };
      return {
                user:$c.login.name,
                testTable_search:{
                  maxRecord:"300",

                  testId:"starts with",
                  testId_s:"",
                  testId_e:"",
                  testDate:"=",
                  testDate_s:"",
                  testDate_e:"",
                  testTimestamp:"=",
                  testTimestamp_s:"",
                  testTimestamp_e:"",
                  testNbr:"=",
                  testNbr_s:"",
                  testNbr_e:"",
                },
                testTable:{
                  url:"/ajax/testTable",
                  cw:$w.tableColW,
                  totalW:$c.totalW($w.tableColW)+2,
                  rcds:[],
                  blank:_.cloneDeep(testTable_blank),
                  selRow:-1
                },
                testTable_form:_.cloneDeep(testTable_blank)
              };
  },
  getStateFromFlux: function() {
    //this.props.flux=$w.flux;
    var pageStore = $w.flux.stores.PAGE;
    var commonStore = $w.flux.stores.COMMON;
    var rcdStore = $w.flux.stores.RCD;
    return {
      page: _.cloneDeep(pageStore.data),
      common:_.cloneDeep(commonStore.data),
      rcd:_.cloneDeep(rcdStore.data)
      };
  },
  render: function() {
    return (
      React.createElement("div", {className: "container-fixed", 
          style: {fontSize:12,border:1,borderStyle:"solid",
            width:$w.width,height:$w.height,backgroundColor: "#F0F0F0"}}, 
      React.createElement(b.Row, {className: "darkBgLarge", 
          style: {margin:0,height:40,lineHeight:"40px",verticalAlign: "middle"}}, 
        React.createElement(b.Col, {xs: 5, style: {textAlign: "center"}}, "TEST"
        ), 
        React.createElement(b.Col, {xs: 5, className: "darkBgMid", style: {textAlign: "center"}}, 
        this.state.user
        ), 
        React.createElement(b.Col, {xs: 1, className: "darkBgMid"}
        ), 
         React.createElement(b.Col, {xs: 1}, 
        React.createElement($c.Loader, {src: "../static/img/ajax-loader.gif", isLoading: this.state.common.loading})
        )
      ), 
      React.createElement(b.Row, {style: {margin:5}}, 
      React.createElement(b.Button, {bsSize: "xsmall", bsStyle: "primary", onClick: $w.handleClick, 
        name: "btntestTable_Search", style: {width:60,marginLeft:10}}, "検索")
      ), 

      React.createElement(b.Row, {style: {verticalAlign:"middle", lineHeight:"26px",marginLeft:0}}, 
         React.createElement(b.Col, {xs: 1, style: {textAlign: "right"}}, "testId"
          ), 
          React.createElement(b.Col, {xs: 2}, 
          React.createElement($c.SelectOption, {options: $c.stringOption, style: {height:24,  fontSize:12}, 
               name: "testTable_search#testId", 
              defaultValue: this.state.testTable_search.testId, onChange: $w.handleChange})
          ), 
          React.createElement(b.Col, {xs: 3}, 
          React.createElement(b.Input, {type: "text", value: this.state.testTable_search.testId_s, 
            name: "testTable_search#testId_s", onChange: $w.handleChange, 
            style: {height:24,fontSize:12,width:"100%"}})
          ), 
          React.createElement(b.Col, {xs: 3}, 
          React.createElement(b.Input, {type: "text", value: this.state.testTable_search.testId_e, 
            name: "testTable_search#testId_e", onChange: $w.handleChange, 
            style: {height:24,fontSize:12,width:"100%"}})
          ), 

          React.createElement(b.Col, {xs: 1, style: {textAlign: "right"}}, "MaxRecord"
          ), 
          React.createElement(b.Col, {xs: 1}, 
            React.createElement(b.Input, {type: "text", value: this.state.testTable_search.maxRecord, 
            name: "testTable_search#maxRecord", onChange: $w.handleChange, 
            style: {height:24,fontSize:12,width:"100%"}})
          )
      ), 

      React.createElement(b.Row, {style: {verticalAlign:"middle", lineHeight:"26px",marginLeft:0}}, 
         React.createElement(b.Col, {xs: 1, style: {textAlign: "right"}}, "testDate"
          ), 
          React.createElement(b.Col, {xs: 2}, 
          React.createElement($c.SelectOption, {options: $c.numberOption, style: {height:24,  fontSize:12}, 
               name: "testTable_search#testDate", 
              defaultValue: this.state.testTable_search.testDate, onChange: $w.handleChange})
          ), 
          React.createElement(b.Col, {xs: 3}, 
          React.createElement(b.Input, {type: "text", value: this.state.testTable_search.testDate_s, 
            name: "testTable_search#testDate_s", onChange: $w.handleChange, 
            style: {height:24,fontSize:12,width:"100%"}})
          ), 
          React.createElement(b.Col, {xs: 3}, 
          React.createElement(b.Input, {type: "text", value: this.state.testTable_search.testDate_e, 
            name: "testTable_search#testDate_e", onChange: $w.handleChange, 
            style: {height:24,fontSize:12,width:"100%"}})
          )

      ), 

      React.createElement(b.Row, {style: {verticalAlign:"middle", lineHeight:"26px",marginLeft:0}}, 
         React.createElement(b.Col, {xs: 1, style: {textAlign: "right"}}, "testTimestamp"
          ), 
          React.createElement(b.Col, {xs: 2}, 
          React.createElement($c.SelectOption, {options: $c.numberOption, style: {height:24,  fontSize:12}, 
               name: "testTable_search#testTimestamp", 
              defaultValue: this.state.testTable_search.testTimestamp, onChange: $w.handleChange})
          ), 
          React.createElement(b.Col, {xs: 3}, 
          React.createElement(b.Input, {type: "text", value: this.state.testTable_search.testTimestamp_s, 
            name: "testTable_search#testTimestamp_s", onChange: $w.handleChange, 
            style: {height:24,fontSize:12,width:"100%"}})
          ), 
          React.createElement(b.Col, {xs: 3}, 
          React.createElement(b.Input, {type: "text", value: this.state.testTable_search.testTimestamp_e, 
            name: "testTable_search#testTimestamp_e", onChange: $w.handleChange, 
            style: {height:24,fontSize:12,width:"100%"}})
          )

      ), 

      React.createElement(b.Row, {style: {verticalAlign:"middle", lineHeight:"26px",marginLeft:0}}, 
         React.createElement(b.Col, {xs: 1, style: {textAlign: "right"}}, "testNbr"
          ), 
          React.createElement(b.Col, {xs: 2}, 
          React.createElement($c.SelectOption, {options: $c.numberOption, style: {height:24,  fontSize:12}, 
               name: "testTable_search#testNbr", 
              defaultValue: this.state.testTable_search.testNbr, onChange: $w.handleChange})
          ), 
          React.createElement(b.Col, {xs: 3}, 
          React.createElement(b.Input, {type: "text", value: this.state.testTable_search.testNbr_s, 
            name: "testTable_search#testNbr_s", onChange: $w.handleChange, 
            style: {height:24,fontSize:12,width:"100%"}})
          ), 
          React.createElement(b.Col, {xs: 3}, 
          React.createElement(b.Input, {type: "text", value: this.state.testTable_search.testNbr_e, 
            name: "testTable_search#testNbr_e", onChange: $w.handleChange, 
            style: {height:24,fontSize:12,width:"100%"}})
          )

      ), 

      React.createElement("div", {style: {width:$w.width-2,border:1,borderStyle:"solid",
          borderColor:"black",height:$w.tableHeight,backgroundColor: "#FFFFFF"}}, 
      React.createElement("div", {ref: "testTable_tableHead", 
          style: {width:$w.width-20,height:20,overflowX:"hidden",overflowY:"hidden"}}, 
      React.createElement(b.Table, {bordered: true, condensed: true, className: "wscrolltable"}, 
       React.createElement("thead", {style: {width:this.state.testTable.totalW,overflowX:"hidden",overflowY:"hidden"}}, 
        React.createElement("tr", null, 
          React.createElement("th", {style: {width:this.state.testTable.cw.c1}}, "testId"), 
          React.createElement("th", {style: {width:this.state.testTable.cw.c2}}, "testDate"), 
          React.createElement("th", {style: {width:this.state.testTable.cw.c3}}, "testTimestamp"), 
          React.createElement("th", {style: {width:this.state.testTable.cw.c4}}, "testNbr"), 
          React.createElement("th", {style: {width:this.state.testTable.cw.c5}}, "id"), 
          React.createElement("th", {style: {width:this.state.testTable.cw.c6}}, "versionNo")

        )
      )
      )
       ), 
      React.createElement("div", {ref: "testTable_tableBody", 
        style: {width:$w.width-4,height:$w.tableHeight-22,overflowX:"scroll",overflowY:"scroll"}}, 
      React.createElement("div", {style: {width:this.state.testTable.totalW,overflowX:"hidden",overflowY:"hidden"}}, 
      React.createElement(b.Table, {bordered: true, condensed: true, className: "wscrolltable", 
       onClick: $w.handleClick}, 
      React.createElement($w.testTableRows, {rcds: this.state.testTable.rcds, cw: this.state.testTable.cw, 
          selRow: this.state.testTable.selRow})
      )
      )
      )
      ), 
      React.createElement(b.Row, {style: {margin:5}}, 
        React.createElement(b.Button, {bsSize: "xsmall", bsStyle: "primary", onClick: $w.handleClick, 
            name: "btntestTable_New", style: {width:60,marginLeft:10}}, "新規"), 
        React.createElement(b.Button, {bsSize: "xsmall", bsStyle: "primary", onClick: $w.handleClick, 
            name: "btntestTable_Update", style: {width:60,marginLeft:10}}, "更新"), 
        React.createElement(b.Button, {bsSize: "xsmall", bsStyle: "primary", onClick: $w.handleClick, 
            name: "btntestTable_Delete", style: {width:60,marginLeft:10}}, "削除")

      ), 
      React.createElement(b.Row, {style: {verticalAlign:"middle", lineHeight:"26px",marginLeft:0,marginRight:5}}, 
         React.createElement(b.Col, {xs: 1, style: {textAlign: "right"}}, "testId"
          ), 
          React.createElement(b.Col, {xs: 3}, 
          React.createElement(b.Input, {type: "text", value: this.state.testTable_form.testId, 
            name: "testTable_form#testId", onChange: $w.handleChange, 
            style: {height:24,fontSize:12,width:"100%"}})
          )
      ), 
      React.createElement(b.Row, {style: {verticalAlign:"middle", lineHeight:"26px",marginLeft:0,marginRight:5}}, 
         React.createElement(b.Col, {xs: 1, style: {textAlign: "right"}}, "testDate"
          ), 
          React.createElement(b.Col, {xs: 3}, 
          React.createElement(b.Input, {type: "text", value: this.state.testTable_form.testDate, 
            name: "testTable_form#testDate", onChange: $w.handleChange, 
            style: {height:24,fontSize:12,width:"100%"}})
          )
      ), 
      React.createElement(b.Row, {style: {verticalAlign:"middle", lineHeight:"26px",marginLeft:0,marginRight:5}}, 
         React.createElement(b.Col, {xs: 1, style: {textAlign: "right"}}, "testTimestamp"
          ), 
          React.createElement(b.Col, {xs: 3}, 
          React.createElement(b.Input, {type: "text", value: this.state.testTable_form.testTimestamp, 
            name: "testTable_form#testTimestamp", onChange: $w.handleChange, 
            style: {height:24,fontSize:12,width:"100%"}})
          )
      ), 
      React.createElement(b.Row, {style: {verticalAlign:"middle", lineHeight:"26px",marginLeft:0,marginRight:5}}, 
         React.createElement(b.Col, {xs: 1, style: {textAlign: "right"}}, "testNbr"
          ), 
          React.createElement(b.Col, {xs: 3}, 
          React.createElement(b.Input, {type: "text", value: this.state.testTable_form.testNbr, 
            name: "testTable_form#testNbr", onChange: $w.handleChange, 
            style: {height:24,fontSize:12,width:"100%"}})
          )
      ), 
      React.createElement(b.Row, {style: {verticalAlign:"middle", lineHeight:"26px",marginLeft:0,marginRight:5}}, 
         React.createElement(b.Col, {xs: 1, style: {textAlign: "right"}}, "id"
          ), 
          React.createElement(b.Col, {xs: 3}, 
          React.createElement(b.Input, {type: "text", value: this.state.testTable_form.id, 
            name: "testTable_form#id", onChange: $w.handleChange, 
            disabled: true, 
            style: {height:24,fontSize:12,width:"100%"}})
          )
      ), 
      React.createElement(b.Row, {style: {verticalAlign:"middle", lineHeight:"26px",marginLeft:0,marginRight:5}}, 
         React.createElement(b.Col, {xs: 1, style: {textAlign: "right"}}, "versionNo"
          ), 
          React.createElement(b.Col, {xs: 3}, 
          React.createElement(b.Input, {type: "text", value: this.state.testTable_form.versionNo, 
            name: "testTable_form#versionNo", onChange: $w.handleChange, 
            disabled: true, 
            style: {height:24,fontSize:12,width:"100%"}})
          )
      ), 

      React.createElement($c.Alert, {isShow: this.state.common.alert.isShow, 
          message: this.state.common.alert.message, onClick: $w.handleClick}), 
      React.createElement($c.DeleteConfirm, {isShow: this.state.common.deleteCfm.isShow, 
          onClick: $w.handleClick})
      )
    );
  },
  componentDidMount: function() {
    $w.testTable_onscroll();
  }
});

React.render(React.createElement($w.Application, {flux: $w.flux}), document.getElementById('content'));

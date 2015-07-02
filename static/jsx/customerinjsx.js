 
  
 var b = ReactBootstrap;
$w.width=900;
$w.height=800;
$w.tableHeight=340;
$w.tableColW={c1:50,c2:50,c3:50,c4:50,c5:50,c6:50,c7:50,c8:50,c9:50,c10:50}
$w.customerRows = React.createClass({

    render: function() {
        $w._customerRows = this;
        var rows = this.props.rcds.map(function(rcd, i){
        var bgcolor="";
        if (i==this.props.selRow)
          bgcolor="#d9edf7";
        else
          if (i%2 == 1)
            bgcolor="#F8F8F8";
          else
            bgcolor="#FFFFFF";
        if (i==this.props.selRow) {
        return (
          <tr key={i} >
 

            <td style={{width:this.props.cw.c1,backgroundColor:bgcolor,padding:0,margin:0}}>
             <b.Input type="text" value={rcd.cusCd}  name={"customer_row#cusCd#"+i}
                  ref={"customer_row#cusCd#"+i} 
                  onChange={this.props.onChange} onKeyDown={this.props.onKeyDown}
                  style={{height:20,fontSize:12,width:"100%",padding:0,margin:0}}/>
            </td>

 

            <td style={{width:this.props.cw.c2,backgroundColor:bgcolor,padding:0,margin:0}}>
             <b.Input type="text" value={rcd.name}  name={"customer_row#name#"+i}
                  ref={"customer_row#name#"+i} 
                  onChange={this.props.onChange} onKeyDown={this.props.onKeyDown}
                  style={{height:20,fontSize:12,width:"100%",padding:0,margin:0}}/>
            </td>

 

            <td style={{width:this.props.cw.c3,backgroundColor:bgcolor,padding:0,margin:0}}>
             <b.Input type="text" value={rcd.addr}  name={"customer_row#addr#"+i}
                  ref={"customer_row#addr#"+i} 
                  onChange={this.props.onChange} onKeyDown={this.props.onKeyDown}
                  style={{height:20,fontSize:12,width:"100%",padding:0,margin:0}}/>
            </td>

 

            <td style={{width:this.props.cw.c4,backgroundColor:bgcolor,padding:0,margin:0}}>
             <b.Input type="text" value={rcd.bldg}  name={"customer_row#bldg#"+i}
                  ref={"customer_row#bldg#"+i} 
                  onChange={this.props.onChange} onKeyDown={this.props.onKeyDown}
                  style={{height:20,fontSize:12,width:"100%",padding:0,margin:0}}/>
            </td>

 

            <td style={{width:this.props.cw.c5,backgroundColor:bgcolor,padding:0,margin:0}}>
             <b.Input type="text" value={rcd.cusConSec}  name={"customer_row#cusConSec#"+i}
                  ref={"customer_row#cusConSec#"+i} 
                  onChange={this.props.onChange} onKeyDown={this.props.onKeyDown}
                  style={{height:20,fontSize:12,width:"100%",padding:0,margin:0}}/>
            </td>

 

            <td style={{width:this.props.cw.c6,backgroundColor:bgcolor,padding:0,margin:0}}>
             <b.Input type="text" value={rcd.cusConName}  name={"customer_row#cusConName#"+i}
                  ref={"customer_row#cusConName#"+i} 
                  onChange={this.props.onChange} onKeyDown={this.props.onKeyDown}
                  style={{height:20,fontSize:12,width:"100%",padding:0,margin:0}}/>
            </td>

 

            <td style={{width:this.props.cw.c7,backgroundColor:bgcolor,padding:0,margin:0}}>
             <b.Input type="text" value={rcd.tel}  name={"customer_row#tel#"+i}
                  ref={"customer_row#tel#"+i} 
                  onChange={this.props.onChange} onKeyDown={this.props.onKeyDown}
                  style={{height:20,fontSize:12,width:"100%",padding:0,margin:0}}/>
            </td>

 

            <td style={{width:this.props.cw.c8,backgroundColor:bgcolor,padding:0,margin:0}}>
             <b.Input type="text" value={rcd.salesAmount}  name={"customer_row#salesAmount#"+i}
                  ref={"customer_row#salesAmount#"+i} 
                  onChange={this.props.onChange} onKeyDown={this.props.onKeyDown}
                  style={{height:20,fontSize:12,width:"100%",padding:0,margin:0}}/>
            </td>

 

            <td id={"customer_row#id#"+i}
            style={{width:this.props.cw.c9,backgroundColor:bgcolor,textAlign:"right"}}>{rcd.id}</td>

 

            <td id={"customer_row#versionNo#"+i}
            style={{width:this.props.cw.c10,backgroundColor:bgcolor,textAlign:"right"}}>{rcd.versionNo}</td>


         </tr>
        )
    }
    else {
        return (
          <tr key={i} >
 <td id={"customer_row#cusCd#"+i}
            style={{width:this.props.cw.c1,backgroundColor:bgcolor,textAlign:"left"}}>{rcd.cusCd}</td>
 <td id={"customer_row#name#"+i}
            style={{width:this.props.cw.c2,backgroundColor:bgcolor,textAlign:"left"}}>{rcd.name}</td>
 <td id={"customer_row#addr#"+i}
            style={{width:this.props.cw.c3,backgroundColor:bgcolor,textAlign:"left"}}>{rcd.addr}</td>
 <td id={"customer_row#bldg#"+i}
            style={{width:this.props.cw.c4,backgroundColor:bgcolor,textAlign:"left"}}>{rcd.bldg}</td>
 <td id={"customer_row#cusConSec#"+i}
            style={{width:this.props.cw.c5,backgroundColor:bgcolor,textAlign:"left"}}>{rcd.cusConSec}</td>
 <td id={"customer_row#cusConName#"+i}
            style={{width:this.props.cw.c6,backgroundColor:bgcolor,textAlign:"left"}}>{rcd.cusConName}</td>
 <td id={"customer_row#tel#"+i}
            style={{width:this.props.cw.c7,backgroundColor:bgcolor,textAlign:"left"}}>{rcd.tel}</td>
 <td id={"customer_row#salesAmount#"+i}
            style={{width:this.props.cw.c8,backgroundColor:bgcolor,textAlign:"right"}}>{rcd.salesAmount}</td>
 <td id={"customer_row#id#"+i}
            style={{width:this.props.cw.c9,backgroundColor:bgcolor,textAlign:"right"}}>{rcd.id}</td>
 <td id={"customer_row#versionNo#"+i}
            style={{width:this.props.cw.c10,backgroundColor:bgcolor,textAlign:"right"}}>{rcd.versionNo}</td>

         </tr>
        )
    }
        }, this);
          return (
            <tbody>
                {rows}
            </tbody>
          );
    }
  });
$w.Application = React.createClass({
  mixins: [$w.FluxMixin, $w.StoreWatchMixin("PAGE","COMMON","RCD")],
  getInitialState: function() {
  $w.app = this;
  customer_blank={
          cusCd:"",
          name:"",
          addr:"",
          bldg:"",
          cusConSec:"",
          cusConName:"",
          tel:"",
          salesAmount:"",
          id:"",
          versionNo:"",

  };
      return {
                user:$c.login.name,
                customer_search:{
                  maxRecord:"300",

                  cusCd:"starts with",
                  cusCd_s:"",
                  cusCd_e:"",
                  name:"starts with",
                  name_s:"",
                  name_e:"",
                },
                customer:{
                  url:"/ajax/customer",
                  cw:$w.tableColW,
                  totalW:$c.totalW($w.tableColW)+2,
                  rcds:[],
                  blank:_.cloneDeep(customer_blank),
                  selRow:-1
                },
                customer_form:_.cloneDeep(customer_blank)
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
      <div className="container-fixed"
          style={{fontSize:12,border:1,borderStyle:"solid",
            width:$w.width,height:$w.height,backgroundColor: "#F0F0F0"}}>
      <b.Row className="darkBgLarge"
          style={{margin:0,height:40,lineHeight:"40px",verticalAlign: "middle"}}>
        <b.Col xs={5} style={{textAlign: "center"}}>カストマー管理
        </b.Col>
        <b.Col xs={5} className="darkBgMid" style={{textAlign: "center"}}>
        {this.state.user}
        </b.Col>
        <b.Col xs={1} className="darkBgMid" >
        </b.Col>
         <b.Col xs={1} >
        <$c.Loader src="../static/img/ajax-loader.gif" isLoading={this.state.common.loading}/>
        </b.Col>
      </b.Row>
      <b.Row style={{margin:5}}>
      <b.Button bsSize="xsmall" bsStyle="primary" onClick={$w.handleClick}
        name="btncustomer_Search" style={{width:60,marginLeft:10}}>検索</b.Button>
      </b.Row>

      <b.Row style={{verticalAlign:"middle", lineHeight:"26px",marginLeft:0}}>
         <b.Col xs={1} style={{textAlign: "right"}}>cusCd
          </b.Col>
          <b.Col xs={2} >
          <$c.SelectOption options={$c.stringOption} style={{height:24,  fontSize:12}}
               name={"customer_search#cusCd"}
              defaultValue={this.state.customer_search.cusCd} onChange={$w.handleChange} />
          </b.Col>
          <b.Col xs={3}>
          <b.Input type="text" value={this.state.customer_search.cusCd_s}
            name="customer_search#cusCd_s" onChange={$w.handleChange}
            style={{height:24,fontSize:12,width:"100%"}}/>
          </b.Col>
          <b.Col xs={3}>
          <b.Input type="text" value={this.state.customer_search.cusCd_e}
            name="customer_search#cusCd_e" onChange={$w.handleChange}
            style={{height:24,fontSize:12,width:"100%"}}/>
          </b.Col>

          <b.Col xs={1} style={{textAlign: "right"}}>MaxRecord
          </b.Col>
          <b.Col xs={1} >
            <b.Input type="text" value={this.state.customer_search.maxRecord}
            name="customer_search#maxRecord" onChange={$w.handleChange}
            style={{height:24,fontSize:12,width:"100%"}}/>
          </b.Col>
      </b.Row>

      <b.Row style={{verticalAlign:"middle", lineHeight:"26px",marginLeft:0}}>
         <b.Col xs={1} style={{textAlign: "right"}}>name
          </b.Col>
          <b.Col xs={2} >
          <$c.SelectOption options={$c.stringOption} style={{height:24,  fontSize:12}}
               name={"customer_search#name"}
              defaultValue={this.state.customer_search.name} onChange={$w.handleChange} />
          </b.Col>
          <b.Col xs={3}>
          <b.Input type="text" value={this.state.customer_search.name_s}
            name="customer_search#name_s" onChange={$w.handleChange}
            style={{height:24,fontSize:12,width:"100%"}}/>
          </b.Col>
          <b.Col xs={3}>
          <b.Input type="text" value={this.state.customer_search.name_e}
            name="customer_search#name_e" onChange={$w.handleChange}
            style={{height:24,fontSize:12,width:"100%"}}/>
          </b.Col>

      </b.Row>

      <b.Row style={{margin:5}}>
        <b.Button bsSize="xsmall" bsStyle="primary" onClick={$w.handleClick}
            name="btncustomer_Update" style={{width:60,marginLeft:10}}>更新</b.Button>
        <b.Button bsSize="xsmall" bsStyle="primary" onClick={$w.handleClick}
            name="btncustomer_Delete" style={{width:60,marginLeft:10}}>削除</b.Button>
        <b.Button bsSize="xsmall" bsStyle="primary" onClick={$w.handleClick}
            name="btncustomer_Cancel" style={{width:60,marginLeft:10}}>取消</b.Button>
      </b.Row>
      <div style={{width:$w.width-2,border:1,borderStyle:"solid",
          borderColor:"black",height:$w.tableHeight,backgroundColor: "#FFFFFF"}}>
      <div ref="customer_tableHead"
          style={{width:$w.width-20,height:20,overflowX:"hidden",overflowY:"hidden"}}>
      <b.Table bordered condensed className="wscrolltable" >
       <thead style={{width:this.state.customer.totalW,overflowX:"hidden",overflowY:"hidden"}}>
        <tr >
          <th style={{width:this.state.customer.cw.c1}}>cusCd</th>
          <th style={{width:this.state.customer.cw.c2}}>name</th>
          <th style={{width:this.state.customer.cw.c3}}>addr</th>
          <th style={{width:this.state.customer.cw.c4}}>bldg</th>
          <th style={{width:this.state.customer.cw.c5}}>cusConSec</th>
          <th style={{width:this.state.customer.cw.c6}}>cusConName</th>
          <th style={{width:this.state.customer.cw.c7}}>tel</th>
          <th style={{width:this.state.customer.cw.c8}}>salesAmount</th>
          <th style={{width:this.state.customer.cw.c9}}>id</th>
          <th style={{width:this.state.customer.cw.c10}}>versionNo</th>

        </tr>
      </thead>
      </b.Table>
       </div>
      <div ref="customer_tableBody"
        style={{width:$w.width-4,height:$w.tableHeight-22,overflowX:"scroll",overflowY:"scroll"}}>
      <div style={{width:this.state.customer.totalW,overflowX:"hidden",overflowY:"hidden"}}>
      <b.Table bordered condensed className="wscrolltable"
       onClick={$w.handleClick}>
      <$w.customerRows rcds={this.state.customer.rcds} cw={this.state.customer.cw}
          selRow={this.state.customer.selRow}  onChange={$w.handleChange}
          onKeyDown={$w.handleRowKeyDown} />
      </b.Table>
      </div>
      </div>
      </div>


      <$c.Alert isShow={this.state.common.alert.isShow}
          message={this.state.common.alert.message} onClick={$w.handleClick} />
      <$c.DeleteConfirm isShow={this.state.common.deleteCfm.isShow}
          onClick={$w.handleClick}/>
      </div>
    );
  },
  componentDidMount: function() {
    $w.customer_onscroll();
  }
});

React.render(<$w.Application flux={$w.flux}/>, document.getElementById('content'));

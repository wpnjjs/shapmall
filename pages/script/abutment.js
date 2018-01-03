// register the grid component
Vue.component('demo-grid', {
    template: '#grid-template',
    props: {
        checked: false,
        checkedboxarr: Array,
        count: 0,
        data: Array,
        keys: Array,
        columns: Array,
        filterKey: String,
    },
    data: function () {
        this.checkedboxarr = new Array();
        console.log("data>>init");
        this.retrive();
    },
    computed: {
        filteredData: function () {
            console.log("computed>>filteredData");
            console.log("this.data",this.data)
            // return this.data
            return this.data;
        }
    },
    methods: {
        modify: _.debounce( function (key) {
            console.log("modify",key)
        },500),

        delone: _.debounce(  function (key) {
            console.log("delete",key)
            $.ajax({
                type: "POST",
                url: "/abutment/d",
                traditional: true,
                data: {
                    id: key
                },
                success: function(msg) {
                    console.log("/abutment/d",msg);
                    // 获取修改后的数据
                    this.retrive();
                }.bind(this),
                error: function(msg){
                    alert("服务器异常！");
                }.bind(this)            
            });
        },500),

        select: _.debounce(  function (key) {
            console.log("select",key);
            var obj = {};
            this.data.forEach(record => {
                if(record.Id_ == key){
                    obj = record;
                }
            });
            if(obj.checked){
                this.checkedboxarr.push(key);
            }else{
                var count = -1;
                this.checkedboxarr.forEach(ele => {
                    count++;
                    if(ele == key){
                        this.checkedboxarr.splice(count,1);
                    }
                });
            }
            console.log("select >>", this.checkedboxarr);
        },500),

        selectall: _.debounce(  function () {
            console.log("selectall");
            var _this = this;
            console.log(this.checked);
            if(this.checked){
                this.checkedboxarr = [];
                // 全选
                this.data.forEach(element => {
                    element.checked = true;
                    this.checkedboxarr.push(element.Id_);
                });
            }else{
                // 取消全选
                this.data.forEach(element => {
                    element.checked = false;
                });
                this.checkedboxarr = [];
            }
            console.log(this.checkedboxarr);
        },500),

        batchdel: _.debounce(  function () {
            console.log("batchdel");
            console.log(this.checkedboxarr);
            $.ajax({
                type: "POST",
                url: "/abutment/bd",
                traditional: true,
                data: {
                    bid: JSON.stringify(this.checkedboxarr)
                },
                success: function(msg) {
                    console.log("/abutment/bd",msg);
                    // 获取修改后的数据
                    this.retrive();
                }.bind(this),
                error: function(msg){
                    alert("服务器异常！");
                }.bind(this)            
            });
        },500),

        create: _.debounce(  function () {
            console.log("create");
            document.getElementById("createandupdate").style.display="block";
            document.getElementById("fade").style.display="block";
        },500),

        retrive: _.debounce(  function () {
            console.log("retrive");
            $.ajax({
                type: "POST",
                url: "/abutment/r",
                traditional: true,
                data: {},
                success: function(msg) {
                    console.log("/abutment/r",msg);
                    if("null" != msg){
                        msgjson = JSON.parse(msg)
                        msgjson.forEach(record => {
                            record.checked = false
                        });
                        console.log(msgjson);
                        this.data = msgjson;
                    }else{
                        this.data = [];
                    }
                }.bind(this),
                error: function(msg){
                    alert("服务器异常！");
                }.bind(this)            
            });
        },500),
    }
  })
  
  // bootstrap the demo
  var demo = new Vue({
    el: '#demo',
    data: {
        columnsCount: 6,
        searchQuery: '',
        columnNames: ['对接系统名称','对接系统码','创建者','创建日期','操作'],
        gridColumns: ['AbutmentSystemName', 'AbutmentSystemCode', 'CreatorCode', 'CreateDate'],
        gridData: [],
        // add and update pages data
        abutmentsystemname: "name",
        abutmentsystemcode: "code",
    },
    methods: {
        retrive: _.debounce(  function () {
            console.log("retrive");
            $.ajax({
                type: "POST",
                url: "/abutment/r",
                traditional: true,
                data: {},
                success: function(msg) {
                    console.log("/abutment/r",msg);
                    if(msg){
                        msgjson = JSON.parse(msg)
                        msgjson.forEach(record => {
                            record.checked = false
                        });
                        console.log(msgjson);
                        this.gridData = msgjson;
                    }
                }.bind(this),
                error: function(msg){
                    alert("服务器异常！");
                }.bind(this)            
            });
        },500),

        insert: _.debounce(  function () {
            console.log("insert");
            // 插入数据
            $.ajax({
                type: "POST",
                url: "/abutment/c",
                traditional: true,
                data: {
                    abutmentsystemname: this.abutmentsystemname,
                    abutmentsystemcode: this.abutmentsystemcode,
                    uid: localStorage.uid
                },
                success: function(msg) {
                    console.log("/abutment/c",msg);
                    document.getElementById("createandupdate").style.display="none";
                    document.getElementById("fade").style.display="none";
                    // 插入数据成功，获取新数据列表
                    this.retrive();
                }.bind(this),
                error: function(msg){
                    alert("服务器异常！");
                }.bind(this)            
            });
        },500),
        
        close: _.debounce( function () {
            console.log("close");
            document.getElementById("createandupdate").style.display="none";
            document.getElementById("fade").style.display="none";
        },500),
    }
  })
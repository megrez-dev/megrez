(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-ac32efe0"],{3122:function(t,e,i){"use strict";i("a513")},a513:function(t,e,i){},eec4:function(t,e,i){"use strict";i.r(e);var a=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("PageView",[i("template",{slot:"content"},[i("t-row",[i("t-col",{attrs:{flex:"34%"}},[i("t-card",{attrs:{title:"添加友链",bordered:!1,"header-bordered":""}},[i("t-form",{ref:"form",attrs:{labelAlign:"top",colon:!0}},[i("t-form-item",{attrs:{label:"网站名称",name:"name"}},[i("t-input",{model:{value:t.link.name,callback:function(e){t.$set(t.link,"name",e)},expression:"link.name"}})],1),i("t-form-item",{attrs:{label:"网站地址",name:"url"}},[i("t-input",{attrs:{placeholder:"http://"},model:{value:t.link.url,callback:function(e){t.$set(t.link,"url",e)},expression:"link.url"}})],1),i("t-form-item",{attrs:{label:"Logo",name:"logo"}},[i("t-input",{model:{value:t.link.logo,callback:function(e){t.$set(t.link,"logo",e)},expression:"link.logo"}})],1),i("t-form-item",{attrs:{label:"排序",name:"priority"}},[i("t-input-number",{attrs:{theme:"row"},model:{value:t.link.priority,callback:function(e){t.$set(t.link,"priority",e)},expression:"link.priority"}})],1),i("t-form-item",{attrs:{label:"网站描述",name:"description"}},[i("t-textarea",{attrs:{autosize:{minRows:3}},model:{value:t.link.description,callback:function(e){t.$set(t.link,"description",e)},expression:"link.description"}})],1)],1),i("template",{slot:"footer"},["add"===t.mode?i("t-button",{attrs:{theme:"primary"},on:{click:t.handleClickAdd}},[t._v(" 添加 ")]):t._e(),"edit"===t.mode?i("t-button",{attrs:{theme:"primary"},on:{click:t.handleClickUpdate}},[t._v(" 更新 ")]):t._e(),"edit"===t.mode?i("t-button",{attrs:{theme:"primary",variant:"dashed"},on:{click:t.handleClickReturn}},[t._v(" 返回添加 ")]):t._e()],1)],2)],1),i("t-col",{attrs:{flex:"1%"}}),i("t-col",{attrs:{flex:"65%"}},[i("t-card",{attrs:{title:"友链列表",bordered:!1,"header-bordered":""}},[i("t-table",{attrs:{data:t.links,columns:t.columns,rowKey:"id",verticalAlign:"middle",loading:t.isLoading,pagination:t.pagination},on:{change:t.rehandleChange},scopedSlots:t._u([{key:"url",fn:function(e){var a=e.row;return[i("a",{staticClass:"t-button-link",staticStyle:{"text-overflow":"ellipsis"},attrs:{href:a.url,target:"_blank"}},[t._v(t._s(a.url))])]}},{key:"logo",fn:function(t){var e=t.row;return[i("div",{staticClass:"logo-wrapper"},[i("img",{attrs:{src:e.logo,alt:e.name}})])]}},{key:"priority",fn:function(t){var e=t.row;return[i("t-badge",{attrs:{count:e.priority,shape:"round",offset:[-14,-5],showZero:""}})]}},{key:"op",fn:function(e){return[i("a",{staticClass:"t-button-link",on:{click:function(i){return t.handleClickEdit(e)}}},[t._v("编辑")]),i("t-divider",{attrs:{layout:"vertical"}}),i("a",{staticClass:"t-button-link",on:{click:function(i){return t.handleClickDelete(e)}}},[t._v("删除")])]}}])})],1)],1)],1)],1)],2)},n=[],l=(i("14d9"),i("c8c3")),s={name:"Links",data(){return{mode:"add",link:{name:"",url:"",logo:"",priority:0,description:""},links:[],isLoading:!1,columns:[{colKey:"name",title:"名称",width:"150px"},{colKey:"url",title:"URL"},{colKey:"logo",title:"LOGO",width:"100px"},{colKey:"priority",title:"排序",width:"100px"},{fixed:"right",colKey:"op",title:"操作"}],tableLayout:"auto",rowClassName:"property-class",pagination:{current:1,pageSize:5}}},async mounted(){await this.fetchData(this.pagination)},methods:{handleClickAdd(){""!==this.link.name?""!==this.link.url?""!==this.link.logo?this.$request.post("link",this.link).then(t=>{this.$message.success("添加成功"),this.links.push(t.data),this.clearForm()}).catch(()=>{this.$message.warning("添加失败")}):this.$message.warning("LOGO不能为空"):this.$message.warning("URL不能为空"):this.$message.warning("名称不能为空")},async fetchData(t=this.pagination){try{this.isLoading=!0,this.$request.get("links",{params:{pageNum:t.current,pageSize:t.pageSize}}).then(e=>{this.isLoading=!1,this.links=e.data.list,this.pagination={...t,total:e.data.total}})}catch(e){this.links=[]}},async rehandleChange(t){const{current:e,pageSize:i}=t.pagination,a={current:e,pageSize:i};await this.fetchData(a)},handleClickEdit({row:t}){this.mode="edit",this.link={...t}},handleClickDelete({row:t}){this.$request.delete("link/"+t.id).then(()=>{this.$message.info("删除成功");for(let e=0;e<this.links.length;e++)if(this.links[e].id===t.id){this.links.splice(e,1);break}}).catch(()=>{this.$message.error("删除失败")})},handleClickUpdate(){""!==this.link.name?""!==this.link.url?""!==this.link.logo?this.$request.put("link/"+this.link.id,this.link).then(()=>{this.$message.info("更新成功"),this.mode="add",this.clearForm(),this.fetchData()}).catch(()=>{this.$message.error("更新失败")}):this.$message.warning("LOGO不能为空"):this.$message.warning("URL不能为空"):this.$message.warning("名称不能为空")},handleClickReturn(){this.mode="add",this.clearForm()},clearForm(){this.link.id=0,this.link.name="",this.link.url="",this.link.logo="",this.link.priority=0,this.link.description=""}},components:{PageView:l["a"]}},r=s,o=(i("3122"),i("2877")),c=Object(o["a"])(r,a,n,!1,null,"587ef8a0",null);e["default"]=c.exports}}]);
//# sourceMappingURL=chunk-ac32efe0.b6fcfc90.js.map
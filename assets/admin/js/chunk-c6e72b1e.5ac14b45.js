(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-c6e72b1e"],{3429:function(e,t,a){"use strict";a("3baa")},"3baa":function(e,t,a){},b12b:function(e,t,a){"use strict";a.r(t);var s=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("PageView",[a("template",{slot:"header"},[a("div",{staticClass:"page-header-bar"},[a("div",{staticClass:"page-header-bar-text"},[a("span",{staticClass:"page-header-bar-title"},[e._v("当前主题")]),a("span",{staticClass:"page-header-bar-description"},[e._v(e._s(e.currentTheme))])]),a("span",{staticClass:"page-header-bar-operator"},[a("span",{staticClass:"page-header-bar-operator-item"},[a("t-button",{attrs:{theme:"primary",variant:"base",loading:e.saveBtnLoading},on:{click:e.onClickSave}},[e._v("保存设置")])],1)])])]),a("template",{slot:"content"},[a("div",{staticClass:"theme-setting-container"},[0!=e.themeConfig.tabs.length?a("t-tabs",{attrs:{defaultValue:e.themeConfig.tabs[0].key}},e._l(e.themeConfig.tabs,(function(t){return a("t-tab-panel",{key:t.key,attrs:{value:t.key},scopedSlots:e._u([{key:"label",fn:function(){return[a("icon",{staticStyle:{"margin-right":"4px"},attrs:{name:"setting"}}),e._v(" "+e._s(t.name)+" ")]},proxy:!0}],null,!0)},[a("div",{staticClass:"setting-form-container"},[a("t-form",{ref:"form",refInFor:!0,attrs:{labelAlign:"top",colon:!0}},e._l(t.items,(function(s){return a("t-form-item",{key:s.key,attrs:{label:s.name,name:s.key,help:s.description}},["input"===s.type?a("t-input",{attrs:{placeholder:s.placeholder},model:{value:s.value,callback:function(t){e.$set(s,"value",t)},expression:"item.value"}}):e._e(),"textarea"===s.type?a("t-textarea",{attrs:{placeholder:s.placeholder,name:s.key,autosize:{minRows:6}},model:{value:s.value,callback:function(t){e.$set(s,"value",t)},expression:"item.value"}}):e._e(),"select"===s.type?a("t-select",{attrs:{options:s.options},model:{value:s.value,callback:function(t){e.$set(s,"value",t)},expression:"item.value"}}):e._e(),"multiSelect"===s.type?a("t-select",{attrs:{multiple:"",minCollapsedNum:3,options:s.options},model:{value:s.value,callback:function(t){e.$set(s,"value",t)},expression:"item.value"}}):e._e(),"switch"===s.type?a("t-switch",{model:{value:s.value,callback:function(t){e.$set(s,"value",t)},expression:"item.value"}}):e._e(),"tags"===s.type?a("t-tag-input",{attrs:{clearable:""},model:{value:s.value,callback:function(t){e.$set(s,"value",t)},expression:"item.value"}}):e._e(),"image"===s.type?a("t-input-group",{attrs:{separate:""}},[a("t-input",{style:{width:"500px"},attrs:{clearable:"",placeholder:s.placeholder},model:{value:s.value,callback:function(t){e.$set(s,"value",t)},expression:"item.value"}}),a("t-button",{attrs:{theme:"primary",shape:"square",variant:"outline"},on:{click:function(a){return e.openAttachSelectDrawer(t,s)}}},[a("image-icon",{attrs:{slot:"icon"},slot:"icon"})],1)],1):e._e()],1)})),1)],1)])})),1):e._e(),a("AttachSelectDrawer",{ref:"attachSelectDrawer",attrs:{mode:"single",visible:e.drawerVisible},on:{"update:visible":function(t){e.drawerVisible=t},select:e.selectAttach}})],1)])],2)},i=[],n=a("f646"),l=a("c8e4"),c=a("f06c"),r=a("c8c3"),o={name:"ThemeSetting",data(){return{saveBtnLoading:!1,themeConfig:{tabs:[]},drawerVisible:!1,currentTheme:"",selectedAttachTabIndex:0,selectedAttachItemIndex:0}},methods:{openAttachSelectDrawer(e,t){this.drawerVisible=!0,this.selectedAttachTabIndex=this.themeConfig.tabs.indexOf(e),this.selectedAttachItemIndex=this.themeConfig.tabs[this.selectedAttachTabIndex].items.indexOf(t)},onClickSave(){this.saveBtnLoading=!0,this.$request.put("/theme/current/config",this.themeConfig).then(()=>{this.$message.success("保存成功")}).finally(()=>{this.saveBtnLoading=!1})},fetchConfig(){this.$request.get("/theme/current/config").then(e=>{for(let t=0;t<this.themeConfig.tabs.length;t++)for(let e=0;e<this.themeConfig.tabs[t].items.length;e++)"multiSelect"!==this.themeConfig.tabs[t].items[e].type&&"tags"!==this.themeConfig.tabs[t].items[e].type||(this.themeConfig.tabs[t].items[e].value=[]);this.themeConfig=e.data})},fetchTheme(){this.$request.get("/theme/current/id").then(e=>{this.currentTheme=e.data})},selectAttach(e){this.themeConfig.tabs[this.selectedAttachTabIndex].items[this.selectedAttachItemIndex].value=e.url}},beforeMount(){this.fetchConfig(),this.fetchTheme()},components:{Icon:n["a"],ImageIcon:l["a"],AttachSelectDrawer:c["a"],PageView:r["a"]}},h=o,u=(a("3429"),a("2877")),m=Object(u["a"])(h,s,i,!1,null,"8d37be74",null);t["default"]=m.exports}}]);
//# sourceMappingURL=chunk-c6e72b1e.5ac14b45.js.map
(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-f616cc8e"],{"513f":function(e,t,a){},b12b:function(e,t,a){"use strict";a.r(t);var n=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("PageView",[a("template",{slot:"header"},[a("div",{staticClass:"page-header-bar"},[a("div",{staticClass:"page-header-bar-text"},[a("span",{staticClass:"page-header-bar-title"},[e._v("当前主题")]),a("span",{staticClass:"page-header-bar-description"},[e._v(e._s(e.currentTheme))])]),a("span",{staticClass:"page-header-bar-operator"},[a("span",{staticClass:"page-header-bar-operator-item"},[a("t-button",{attrs:{theme:"primary",variant:"base"},on:{click:e.onClickSave}},[e._v("保存设置")])],1)])])]),a("template",{slot:"content"},[a("div",{staticClass:"theme-setting-container"},[0!=e.themeConfig.tabs.length?a("t-tabs",{attrs:{defaultValue:e.themeConfig.tabs[0].key}},e._l(e.themeConfig.tabs,(function(t){return a("t-tab-panel",{key:t.key,attrs:{value:t.key},scopedSlots:e._u([{key:"label",fn:function(){return[a("icon",{staticStyle:{"margin-right":"4px"},attrs:{name:"setting"}}),e._v(" "+e._s(t.name)+" ")]},proxy:!0}],null,!0)},[a("div",{staticClass:"setting-form-container"},[a("t-form",{ref:"form",refInFor:!0,attrs:{labelAlign:"top",colon:!0}},e._l(t.items,(function(n){return a("t-form-item",{key:n.key,attrs:{label:n.name,name:n.key,help:n.description}},["input"===n.type?a("t-input",{attrs:{placeholder:n.placeholder},model:{value:n.value,callback:function(t){e.$set(n,"value",t)},expression:"item.value"}}):e._e(),"textarea"===n.type?a("t-textarea",{attrs:{placeholder:n.placeholder,name:n.key,autosize:{minRows:6}},model:{value:n.value,callback:function(t){e.$set(n,"value",t)},expression:"item.value"}}):e._e(),"select"===n.type?a("t-select",{attrs:{options:n.options},model:{value:n.value,callback:function(t){e.$set(n,"value",t)},expression:"item.value"}}):e._e(),"multiSelect"===n.type?a("t-select",{attrs:{multiple:"",minCollapsedNum:3,options:n.options},model:{value:n.value,callback:function(t){e.$set(n,"value",t)},expression:"item.value"}}):e._e(),"switch"===n.type?a("t-switch",{model:{value:n.value,callback:function(t){e.$set(n,"value",t)},expression:"item.value"}}):e._e(),"tags"===n.type?a("t-tag-input",{attrs:{clearable:""},model:{value:n.value,callback:function(t){e.$set(n,"value",t)},expression:"item.value"}}):e._e(),"image"===n.type?a("t-input-group",{attrs:{separate:""}},[a("t-input",{style:{width:"500px"},attrs:{placeholder:n.placeholder},model:{value:n.value,callback:function(t){e.$set(n,"value",t)},expression:"item.value"}}),a("t-button",{attrs:{theme:"primary",shape:"square",variant:"outline"},on:{click:function(a){return e.openAttachSelectDrawer(t,n)}}},[a("image-icon",{attrs:{slot:"icon"},slot:"icon"})],1)],1):e._e()],1)})),1)],1)])})),1):e._e(),a("AttachSelectDrawer",{ref:"attachSelectDrawer",on:{select:e.selectAttach}})],1)])],2)},s=[],i=a("f646"),r=a("87eb"),c=a("c09c"),l=a("2b0e"),o=a("b25f"),u=a("6be3"),h=(a("4d26"),a("febf"),["size"]);function p(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,n)}return a}function f(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?p(Object(a),!0).forEach((function(t){Object(r["a"])(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):p(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}var m={tag:"svg",attrs:{fill:"none",viewBox:"0 0 16 16",width:"1em",height:"1em"},children:[{tag:"path",attrs:{fill:"currentColor",d:"M10 8a2 2 0 100-4 2 2 0 000 4zm0-1a1 1 0 100-2 1 1 0 000 2z",fillOpacity:.9,fillRule:"evenodd",clipRule:"evenodd"}},{tag:"path",attrs:{fill:"currentColor",d:"M2 13a1 1 0 001 1h10a1 1 0 001-1V3a1 1 0 00-1-1H3a1 1 0 00-1 1v10zm1-1.3l3-3 4.3 4.3H3v-1.3zm0-1.4V3h10v10h-1.3L6 7.3l-3 3z",fillOpacity:.9}}]},d=l["default"].extend({name:"ImageIcon",functional:!0,props:{size:{type:String},onClick:{type:Function}},render:function(e,t){var a=t.props,n=t.data,s=a.size,i=Object(c["a"])(a,h),r=Object(u["a"])(s),l=r.className,p=r.style,d=f(f({},i||{}),{},{id:"image",icon:m,staticClass:l,style:p});return n.props=d,e(o["a"],n)}}),b=a("f06c"),v=a("c8c3"),g={name:"ThemeSetting",data:function(){return{themeConfig:{tabs:[]},currentTheme:"",selectedAttachTabIndex:0,selectedAttachItemIndex:0}},methods:{openAttachSelectDrawer:function(e,t){this.$refs.attachSelectDrawer.open(),this.selectedAttachTabIndex=this.themeConfig.tabs.indexOf(e),this.selectedAttachItemIndex=this.themeConfig.tabs[this.selectedAttachTabIndex].items.indexOf(t)},onClickSave:function(){var e=this;this.$request.put("/theme/current/config",this.themeConfig).then((function(){e.$message.success("保存成功")})).catch((function(){e.$message.error("保存失败")}))},fetchConfig:function(){var e=this;this.$request.get("/theme/current/config").then((function(t){for(var a=0;a<e.themeConfig.tabs.length;a++)for(var n=0;n<e.themeConfig.tabs[a].items.length;n++)"multiSelect"!==e.themeConfig.tabs[a].items[n].type&&"tags"!==e.themeConfig.tabs[a].items[n].type||(e.themeConfig.tabs[a].items[n].value=[]);e.themeConfig=t.data}))},fetchTheme:function(){var e=this;this.$request.get("/theme/current/id").then((function(t){e.currentTheme=t.data}))},selectAttach:function(e){this.themeConfig.tabs[this.selectedAttachTabIndex].items[this.selectedAttachItemIndex].value=e.url}},beforeMount:function(){this.fetchConfig(),this.fetchTheme()},components:{Icon:i["a"],ImageIcon:d,AttachSelectDrawer:b["a"],PageView:v["a"]}},y=g,C=(a("ff55"),a("2877")),w=Object(C["a"])(y,n,s,!1,null,"13868432",null);t["default"]=w.exports},ff55:function(e,t,a){"use strict";a("513f")}}]);
//# sourceMappingURL=chunk-f616cc8e.b2f5f1d1.js.map
import{q as t,o as e,c as i,w as l,a,b as s,r as o,e as n,d as r}from"./index.8d4a8305.js";import{_ as d,a as u}from"./uni-list.723f536b.js";import{d as c,_ as f}from"./time.0102cd3d.js";import{_ as m}from"./uni-icons.51822eb1.js";import{_ as p,a as _}from"./uni-grid.ef650232.js";import{_ as h}from"./uni-section.96f3ba19.js";import{r as j}from"./const.9e5d4232.js";import{_ as g}from"./plugin-vue_export-helper.21dcd24c.js";import"./_commonjsHelpers.4e997714.js";var b=g({data:()=>({data:{}}),onLoad(t){this.id=t.id,this.load()},onPullDownRefresh(){this.load()},methods:{format:c,load(){j({url:"project/"+this.id,success:t=>{this.data=t},complete(){uni.stopPullDownRefresh()}})},remove(){uni.showModal({title:"提示",content:"确定删除？",success:e=>{e.confirm&&(t("log","at pages/project/detail.vue:89","用户点击确定"),j({url:"project/"+this.id+"/delete",success:t=>{uni.navigateBack(),uni.showToast({title:"删除成功"})}}))},fail:console.error})}}},[["render",function(t,c,j,g,b,x){const T=o(n("uni-list-item"),d),v=o(n("uni-list"),u),k=o(n("uni-card"),f),w=o(n("uni-icons"),m),y=o(n("uni-grid-item"),p),D=o(n("uni-grid"),_),P=o(n("uni-section"),h),C=r;return e(),i(C,null,{default:l((()=>[a(k,{title:b.data.name,subTitle:b.data.id,note:"Tips",thumbnail:"/static/icons/project.svg"},{default:l((()=>[a(v,{border:!1},{default:l((()=>[a(T,{title:"ID",rightText:b.data.id},null,8,["rightText"]),a(T,{title:"创建时间",rightText:x.format(b.data.created)},null,8,["rightText"])])),_:1})])),_:1},8,["title","subTitle"]),a(v,null,{default:l((()=>[a(T,{title:"可视化组态",link:"",to:"./interface?id="+t.id},{header:l((()=>[a(w,{class:"list-icon",type:"color"})])),_:1},8,["to"]),a(T,{title:"编辑项目",link:"",to:"./edit?id="+t.id},{header:l((()=>[a(w,{class:"list-icon",customPrefix:"iconfont",type:"icon-pen"})])),_:1},8,["to"]),a(T,{title:"删除项目",onClick:x.remove,clickable:!0},{header:l((()=>[a(w,{class:"list-icon",customPrefix:"iconfont",type:"icon-dustbin"})])),_:1},8,["onClick"])])),_:1}),a(P,{title:"变量",type:"line"},{default:l((()=>[a(D,{column:3},{default:l((()=>[a(y,null,{default:l((()=>[s("温度：30")])),_:1}),a(y,null,{default:l((()=>[s("温度：30")])),_:1}),a(y,null,{default:l((()=>[s("温度：30")])),_:1}),a(y,null,{default:l((()=>[s("温度：30")])),_:1}),a(y,null,{default:l((()=>[s("温度：30")])),_:1}),a(y,null,{default:l((()=>[s("温度：30")])),_:1}),a(y,null,{default:l((()=>[s("温度：30")])),_:1})])),_:1})])),_:1}),a(P,{title:"日志",type:"line"},{default:l((()=>[s(" TODO：Event ")])),_:1})])),_:1})}]]);export{b as default};
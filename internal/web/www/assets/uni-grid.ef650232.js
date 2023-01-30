import{o as i,c as e,w as t,a as r,k as o,n as d,h as s,g as h,d as l}from"./index.8d4a8305.js";import{_ as a}from"./plugin-vue_export-helper.21dcd24c.js";var n=a({name:"UniGridItem",inject:["grid"],props:{index:{type:Number,default:0}},data:()=>({column:0,showBorder:!0,square:!0,highlight:!0,left:0,top:0,openNum:2,width:0,borderColor:"#e5e5e5"}),created(){this.column=this.grid.column,this.showBorder=this.grid.showBorder,this.square=this.grid.square,this.highlight=this.grid.highlight,this.top=0===this.hor?this.grid.hor:this.hor,this.left=0===this.ver?this.grid.ver:this.ver,this.borderColor=this.grid.borderColor,this.grid.children.push(this),this.width=this.grid.width},beforeDestroy(){this.grid.children.forEach(((i,e)=>{i===this&&this.grid.children.splice(e,1)}))},methods:{_onClick(){this.grid.change({detail:{index:this.index}})}}},[["render",function(a,n,c,u,g,p){const m=l;return g.width?(i(),e(m,{key:0,style:s("width:"+g.width+";"+(g.square?"height:"+g.width:"")),class:"uni-grid-item"},{default:t((()=>[r(m,{class:d([{"uni-grid-item--border":g.showBorder,"uni-grid-item--border-top":g.showBorder&&c.index<g.column,"uni-highlight":g.highlight},"uni-grid-item__box"]),style:s({"border-right-color":g.borderColor,"border-bottom-color":g.borderColor,"border-top-color":g.borderColor}),onClick:p._onClick},{default:t((()=>[o(a.$slots,"default",{},void 0,!0)])),_:3},8,["class","style","onClick"])])),_:3},8,["style"])):h("",!0)}],["__scopeId","data-v-0bfb1cac"]]);var c=a({name:"UniGrid",emits:["change"],props:{column:{type:Number,default:3},showBorder:{type:Boolean,default:!0},borderColor:{type:String,default:"#D2D2D2"},square:{type:Boolean,default:!0},highlight:{type:Boolean,default:!0}},provide(){return{grid:this}},data:()=>({elId:`Uni_${Math.ceil(1e6*Math.random()).toString(36)}`,width:0}),created(){this.children=[]},mounted(){this.$nextTick((()=>{this.init()}))},methods:{init(){setTimeout((()=>{this._getSize((i=>{this.children.forEach(((e,t)=>{e.width=i}))}))}),50)},change(i){this.$emit("change",i)},_getSize(i){uni.createSelectorQuery().in(this).select(`#${this.elId}`).boundingClientRect().exec((e=>{this.width=parseInt((e[0].width-1)/this.column)+"px",i(this.width)}))}}},[["render",function(h,a,n,c,u,g){const p=l;return i(),e(p,{class:"uni-grid-wrap"},{default:t((()=>[r(p,{id:u.elId,ref:"uni-grid",class:d(["uni-grid",{"uni-grid--border":n.showBorder}]),style:s({"border-left-color":n.borderColor})},{default:t((()=>[o(h.$slots,"default",{},void 0,!0)])),_:3},8,["id","class","style"])])),_:3})}],["__scopeId","data-v-b300e6fa"]]);export{n as _,c as a};
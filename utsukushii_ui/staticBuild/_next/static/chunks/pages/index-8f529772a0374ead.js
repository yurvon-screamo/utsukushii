(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[405],{8312:function(e,t,s){(window.__NEXT_P=window.__NEXT_P||[]).push(["/",function(){return s(9185)}])},9185:function(e,t,s){"use strict";s.r(t),s.d(t,{default:function(){return L}});var r,l,n=s(5893),a=s(7294),d=s(512),i=s(8388);function o(){for(var e=arguments.length,t=Array(e),s=0;s<e;s++)t[s]=arguments[s];return(0,i.m6)((0,d.W)(t))}let c=a.forwardRef((e,t)=>{let{className:s,...r}=e;return(0,n.jsx)("div",{ref:t,className:o("rounded-lg border bg-card text-card-foreground shadow-sm",s),...r})});c.displayName="Card";let x=a.forwardRef((e,t)=>{let{className:s,...r}=e;return(0,n.jsx)("div",{ref:t,className:o("flex flex-col space-y-1.5 p-6",s),...r})});x.displayName="CardHeader";let m=a.forwardRef((e,t)=>{let{className:s,...r}=e;return(0,n.jsx)("h3",{ref:t,className:o("text-2xl font-semibold leading-none tracking-tight",s),...r})});m.displayName="CardTitle",a.forwardRef((e,t)=>{let{className:s,...r}=e;return(0,n.jsx)("p",{ref:t,className:o("text-sm text-muted-foreground",s),...r})}).displayName="CardDescription";let u=a.forwardRef((e,t)=>{let{className:s,...r}=e;return(0,n.jsx)("div",{ref:t,className:o("p-6 pt-0",s),...r})});function h(){for(var e=arguments.length,t=Array(e),s=0;s<e;s++)t[s]=arguments[s];return(0,i.m6)((0,d.W)(t))}u.displayName="CardContent",a.forwardRef((e,t)=>{let{className:s,...r}=e;return(0,n.jsx)("div",{ref:t,className:o("flex items-center p-6 pt-0",s),...r})}).displayName="CardFooter";let f=a.forwardRef((e,t)=>{let{className:s,...r}=e;return(0,n.jsx)("div",{className:"relative w-full overflow-auto",children:(0,n.jsx)("table",{ref:t,className:h("w-full caption-bottom text-sm",s),...r})})});f.displayName="Table",a.forwardRef((e,t)=>{let{className:s,...r}=e;return(0,n.jsx)("thead",{ref:t,className:h("[&_tr]:border-b",s),...r})}).displayName="TableHeader";let j=a.forwardRef((e,t)=>{let{className:s,...r}=e;return(0,n.jsx)("tbody",{ref:t,className:h("[&_tr:last-child]:border-0",s),...r})});j.displayName="TableBody",a.forwardRef((e,t)=>{let{className:s,...r}=e;return(0,n.jsx)("tfoot",{ref:t,className:h("border-t bg-muted/50 font-medium [&>tr]:last:border-b-0",s),...r})}).displayName="TableFooter";let p=a.forwardRef((e,t)=>{let{className:s,...r}=e;return(0,n.jsx)("tr",{ref:t,className:h("border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted",s),...r})});p.displayName="TableRow",a.forwardRef((e,t)=>{let{className:s,...r}=e;return(0,n.jsx)("th",{ref:t,className:h("h-12 px-4 text-left align-middle font-medium text-muted-foreground [&:has([role=checkbox])]:pr-0",s),...r})}).displayName="TableHead";let g=a.forwardRef((e,t)=>{let{className:s,...r}=e;return(0,n.jsx)("td",{ref:t,className:h("p-4 align-middle [&:has([role=checkbox])]:pr-0",s),...r})});function w(e){let{testLog:t}=e;return(0,n.jsx)("div",{className:"bg-[#282c34] text-[#bbc2cf] px-4 py-2 text-xs font-mono w-full whitespace-pre transition-opacity ease-linear",children:t})}function N(e){let t=e.records,[s,r]=(0,a.useState)(null),[d,i]=(0,a.useState)(null),o=(e,t)=>{i(e===d?null:e),t&&t.log&&r((null==s?void 0:s.name)===t.name?null:t)};return(0,n.jsx)("div",{children:(0,n.jsx)(f,{children:(0,n.jsx)(j,{children:t.map((e,t)=>(0,n.jsxs)(a.Fragment,{children:[(0,n.jsxs)(p,{onClick:()=>o(t,e),className:"cursor-pointer ".concat(d===t?"bg-muted":""),children:[(0,n.jsx)(g,{className:"font-medium",children:e.name}),(0,n.jsx)(g,{className:"font-medium ".concat(e.state===l.success?"text-green-500":e.state===l.dropped?"text-red-500":"text-yellow-500"),children:e.state}),(0,n.jsx)(g,{children:e.duration})]}),d===t&&e.tests&&e.tests.length>0&&(0,n.jsx)(p,{children:(0,n.jsx)(g,{colSpan:3,children:(0,n.jsx)(N,{records:e.tests})})}),(null==s?void 0:s.log)&&s.name==e.name&&(0,n.jsx)(p,{children:(0,n.jsx)(w,{testName:s.name,testLog:s.log})})]},t))})})})}function v(e){let t=e.content;return(0,n.jsx)("div",{className:"z-10 w-full items-center justify-between font-mono text-sm lg:flex",children:(0,n.jsxs)(c,{className:"w-full",children:[(0,n.jsx)(x,{className:"flex items-left justify-between",children:(0,n.jsxs)("div",{className:"items-center grid grid-cols-3 gap-4",children:[(0,n.jsxs)("div",{className:"flex items-center gap-2",children:[(0,n.jsx)(k,{className:"h-3 w-3 text-muted-foreground"}),(0,n.jsx)("span",{className:"text-xs text-muted-foreground",children:t.timestamp})]}),(0,n.jsx)(m,{children:t.title})]})}),(0,n.jsxs)(u,{children:[(0,n.jsxs)("div",{className:"grid grid-cols-3 gap-4",children:[(0,n.jsxs)("div",{className:"flex items-center gap-2",children:[(0,n.jsx)(R,{className:"h-6 w-6 text-primary"}),(0,n.jsxs)("div",{children:[(0,n.jsx)("div",{className:"text-lg font-medium",children:t.total}),(0,n.jsx)("div",{className:"text-sm text-muted-foreground",children:"Total Tests"})]})]}),(0,n.jsxs)("div",{className:"flex items-center gap-2",children:[(0,n.jsx)(C,{className:"h-6 w-6 text-primary"}),(0,n.jsxs)("div",{children:[(0,n.jsxs)("div",{className:"text-lg font-medium",children:[t.coverage,"%"]}),(0,n.jsx)("div",{className:"text-sm text-muted-foreground",children:"Test Coverage"})]})]}),(0,n.jsxs)("div",{className:"flex flex-col items-start gap-2",children:[(0,n.jsxs)("div",{className:"flex items-center gap-2",children:[(0,n.jsx)(b,{className:"h-5 w-5 text-green-500"}),(0,n.jsxs)("div",{className:"text-sm text-muted-foreground",children:[t.success," Success"]})]}),(0,n.jsxs)("div",{className:"flex items-center gap-2",children:[(0,n.jsx)(T,{className:"h-5 w-5 text-red-500"}),(0,n.jsxs)("div",{className:"text-sm text-muted-foreground",children:[t.dropped," Dropped"]})]}),(0,n.jsxs)("div",{className:"flex items-center gap-2",children:[(0,n.jsx)(y,{className:"h-5 w-5 text-yellow-500"}),(0,n.jsxs)("div",{className:"text-sm text-muted-foreground",children:[t.skipped," Skipped"]})]})]})]}),(0,n.jsx)("div",{className:"mt-6",children:(0,n.jsx)(N,{records:t.tests})})]})]})})}function k(e){return(0,n.jsxs)("svg",{...e,xmlns:"http://www.w3.org/2000/svg",width:"24",height:"24",viewBox:"0 0 24 24",fill:"none",stroke:"currentColor",strokeWidth:"2",strokeLinecap:"round",strokeLinejoin:"round",children:[(0,n.jsx)("path",{d:"M8 2v4"}),(0,n.jsx)("path",{d:"M16 2v4"}),(0,n.jsx)("rect",{width:"18",height:"18",x:"3",y:"4",rx:"2"}),(0,n.jsx)("path",{d:"M3 10h18"})]})}function b(e){return(0,n.jsx)("svg",{...e,xmlns:"http://www.w3.org/2000/svg",width:"24",height:"24",viewBox:"0 0 24 24",fill:"none",stroke:"currentColor",strokeWidth:"2",strokeLinecap:"round",strokeLinejoin:"round",children:(0,n.jsx)("path",{d:"M20 6 9 17l-5-5"})})}function y(e){return(0,n.jsxs)("svg",{...e,xmlns:"http://www.w3.org/2000/svg",width:"24",height:"24",viewBox:"0 0 24 24",fill:"none",stroke:"currentColor",strokeWidth:"2",strokeLinecap:"round",strokeLinejoin:"round",children:[(0,n.jsx)("rect",{x:"14",y:"4",width:"4",height:"16",rx:"1"}),(0,n.jsx)("rect",{x:"6",y:"4",width:"4",height:"16",rx:"1"})]})}function C(e){return(0,n.jsxs)("svg",{...e,xmlns:"http://www.w3.org/2000/svg",width:"24",height:"24",viewBox:"0 0 24 24",fill:"none",stroke:"currentColor",strokeWidth:"2",strokeLinecap:"round",strokeLinejoin:"round",children:[(0,n.jsx)("line",{x1:"19",x2:"5",y1:"5",y2:"19"}),(0,n.jsx)("circle",{cx:"6.5",cy:"6.5",r:"2.5"}),(0,n.jsx)("circle",{cx:"17.5",cy:"17.5",r:"2.5"})]})}function R(e){return(0,n.jsxs)("svg",{...e,xmlns:"http://www.w3.org/2000/svg",width:"24",height:"24",viewBox:"0 0 24 24",fill:"none",stroke:"currentColor",strokeWidth:"2",strokeLinecap:"round",strokeLinejoin:"round",children:[(0,n.jsx)("path",{d:"M14.5 2v17.5c0 1.4-1.1 2.5-2.5 2.5h0c-1.4 0-2.5-1.1-2.5-2.5V2"}),(0,n.jsx)("path",{d:"M8.5 2h7"}),(0,n.jsx)("path",{d:"M14.5 16h-5"})]})}function T(e){return(0,n.jsxs)("svg",{...e,xmlns:"http://www.w3.org/2000/svg",width:"24",height:"24",viewBox:"0 0 24 24",fill:"none",stroke:"currentColor",strokeWidth:"2",strokeLinecap:"round",strokeLinejoin:"round",children:[(0,n.jsx)("path",{d:"M18 6 6 18"}),(0,n.jsx)("path",{d:"m6 6 12 12"})]})}g.displayName="TableCell",a.forwardRef((e,t)=>{let{className:s,...r}=e;return(0,n.jsx)("caption",{ref:t,className:h("mt-4 text-sm text-muted-foreground",s),...r})}).displayName="TableCaption",(r=l||(l={})).success="success",r.dropped="dropped",r.skipped="skipped";var _=s(1758);function L(){let{data:e,error:t,isLoading:s}=(0,_.ZP)("/utsukushii.json",e=>fetch(e).then(e=>e.text()));return t||!e?(0,n.jsx)("div",{children:"failed to load"}):s?(0,n.jsx)("div",{children:"loading..."}):(0,n.jsx)(v,{content:JSON.parse(e)})}}},function(e){e.O(0,[231,888,774,179],function(){return e(e.s=8312)}),_N_E=e.O()}]);
"use strict";(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[231],{3250:function(e,r,t){/**
 * @license React
 * use-sync-external-store-shim.production.min.js
 *
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */var o=t(7294),n="function"==typeof Object.is?Object.is:function(e,r){return e===r&&(0!==e||1/e==1/r)||e!=e&&r!=r},l=o.useState,i=o.useEffect,a=o.useLayoutEffect,s=o.useDebugValue;function d(e){var r=e.getSnapshot;e=e.value;try{var t=r();return!n(e,t)}catch(e){return!0}}var c="undefined"==typeof window||void 0===window.document||void 0===window.document.createElement?function(e,r){return r()}:function(e,r){var t=r(),o=l({inst:{value:t,getSnapshot:r}}),n=o[0].inst,c=o[1];return a(function(){n.value=t,n.getSnapshot=r,d(n)&&c({inst:n})},[e,t,r]),i(function(){return d(n)&&c({inst:n}),e(function(){d(n)&&c({inst:n})})},[e]),s(t),t};r.useSyncExternalStore=void 0!==o.useSyncExternalStore?o.useSyncExternalStore:c},1688:function(e,r,t){e.exports=t(3250)},512:function(e,r,t){t.d(r,{W:function(){return o}});function o(){for(var e,r,t=0,o="",n=arguments.length;t<n;t++)(e=arguments[t])&&(r=function e(r){var t,o,n="";if("string"==typeof r||"number"==typeof r)n+=r;else if("object"==typeof r){if(Array.isArray(r)){var l=r.length;for(t=0;t<l;t++)r[t]&&(o=e(r[t]))&&(n&&(n+=" "),n+=o)}else for(o in r)r[o]&&(n&&(n+=" "),n+=o)}return n}(e))&&(o&&(o+=" "),o+=r);return o}},1758:function(e,r,t){let o;t.d(r,{ZP:function(){return ee}});var n=t(7294),l=t(1688);let i=()=>{},a=i(),s=Object,d=e=>e===a,c=e=>"function"==typeof e,u=(e,r)=>({...e,...r}),p=e=>c(e.then),f=new WeakMap,b=0,g=e=>{let r,t;let o=typeof e,n=e&&e.constructor,l=n==Date;if(s(e)!==e||l||n==RegExp)r=l?e.toJSON():"symbol"==o?e.toString():"string"==o?JSON.stringify(e):""+e;else{if(r=f.get(e))return r;if(r=++b+"~",f.set(e,r),n==Array){for(t=0,r="@";t<e.length;t++)r+=g(e[t])+",";f.set(e,r)}if(n==s){r="#";let o=s.keys(e).sort();for(;!d(t=o.pop());)d(e[t])||(r+=t+":"+g(e[t])+",");f.set(e,r)}}return r},m=new WeakMap,h={},y={},v="undefined",w=typeof window!=v,x=typeof document!=v,k=()=>w&&typeof window.requestAnimationFrame!=v,E=(e,r)=>{let t=m.get(e);return[()=>!d(r)&&e.get(r)||h,o=>{if(!d(r)){let n=e.get(r);r in y||(y[r]=n),t[5](r,u(n,o),n||h)}},t[6],()=>!d(r)&&r in y?y[r]:!d(r)&&e.get(r)||h]},_=!0,[S,R]=w&&window.addEventListener?[window.addEventListener.bind(window),window.removeEventListener.bind(window)]:[i,i],z={initFocus:e=>(x&&document.addEventListener("visibilitychange",e),S("focus",e),()=>{x&&document.removeEventListener("visibilitychange",e),R("focus",e)}),initReconnect:e=>{let r=()=>{_=!0,e()},t=()=>{_=!1};return S("online",r),S("offline",t),()=>{R("online",r),R("offline",t)}}},O=!n.useId,T=!w||"Deno"in window,C=e=>k()?window.requestAnimationFrame(e):setTimeout(e,1),V=T?n.useEffect:n.useLayoutEffect,j="undefined"!=typeof navigator&&navigator.connection,L=!T&&j&&(["slow-2g","2g"].includes(j.effectiveType)||j.saveData),N=e=>{if(c(e))try{e=e()}catch(r){e=""}let r=e;return[e="string"==typeof e?e:(Array.isArray(e)?e.length:e)?g(e):"",r]},P=0,I=()=>++P;var A={ERROR_REVALIDATE_EVENT:3,FOCUS_EVENT:0,MUTATE_EVENT:2,RECONNECT_EVENT:1};async function D(...e){let[r,t,o,n]=e,l=u({populateCache:!0,throwOnError:!0},"boolean"==typeof n?{revalidate:n}:n||{}),i=l.populateCache,s=l.rollbackOnError,f=l.optimisticData,b=e=>"function"==typeof s?s(e):!1!==s,g=l.throwOnError;if(c(t)){let e=[];for(let o of r.keys())!/^\$(inf|sub)\$/.test(o)&&t(r.get(o)._k)&&e.push(o);return Promise.all(e.map(h))}return h(t);async function h(t){let n;let[s]=N(t);if(!s)return;let[u,h]=E(r,s),[y,v,w,x]=m.get(r),k=()=>{let e=y[s];return(c(l.revalidate)?l.revalidate(u().data,t):!1!==l.revalidate)&&(delete w[s],delete x[s],e&&e[0])?e[0](2).then(()=>u().data):u().data};if(e.length<3)return k();let _=o,S=I();v[s]=[S,0];let R=!d(f),z=u(),O=z.data,T=z._c,C=d(T)?O:T;if(R&&h({data:f=c(f)?f(C,O):f,_c:C}),c(_))try{_=_(C)}catch(e){n=e}if(_&&p(_)){if(_=await _.catch(e=>{n=e}),S!==v[s][0]){if(n)throw n;return _}n&&R&&b(n)&&(i=!0,h({data:C,_c:a}))}if(i&&!n&&(c(i)?h({data:i(_,C),error:a,_c:a}):h({data:_,error:a,_c:a})),v[s][1]=I(),Promise.resolve(k()).then(()=>{h({_c:a})}),n){if(g)throw n;return}return _}}let M=(e,r)=>{for(let t in e)e[t][0]&&e[t][0](r)},G=(e,r)=>{if(!m.has(e)){let t=u(z,r),o={},n=D.bind(a,e),l=i,s={},d=(e,r)=>{let t=s[e]||[];return s[e]=t,t.push(r),()=>t.splice(t.indexOf(r),1)},c=(r,t,o)=>{e.set(r,t);let n=s[r];if(n)for(let e of n)e(t,o)},p=()=>{if(!m.has(e)&&(m.set(e,[o,{},{},{},n,c,d]),!T)){let r=t.initFocus(setTimeout.bind(a,M.bind(a,o,0))),n=t.initReconnect(setTimeout.bind(a,M.bind(a,o,1)));l=()=>{r&&r(),n&&n(),m.delete(e)}}};return p(),[e,n,p,l]}return[e,m.get(e)[4]]},[W,$]=G(new Map),F=u({onLoadingSlow:i,onSuccess:i,onError:i,onErrorRetry:(e,r,t,o,n)=>{let l=t.errorRetryCount,i=n.retryCount,a=~~((Math.random()+.5)*(1<<(i<8?i:8)))*t.errorRetryInterval;(d(l)||!(i>l))&&setTimeout(o,a,n)},onDiscarded:i,revalidateOnFocus:!0,revalidateOnReconnect:!0,revalidateIfStale:!0,shouldRetryOnError:!0,errorRetryInterval:L?1e4:5e3,focusThrottleInterval:5e3,dedupingInterval:2e3,loadingTimeout:L?5e3:3e3,compare:(e,r)=>g(e)==g(r),isPaused:()=>!1,cache:W,mutate:$,fallback:{}},{isOnline:()=>_,isVisible:()=>{let e=x&&document.visibilityState;return d(e)||"hidden"!==e}}),U=(e,r)=>{let t=u(e,r);if(r){let{use:o,fallback:n}=e,{use:l,fallback:i}=r;o&&l&&(t.use=o.concat(l)),n&&i&&(t.fallback=u(n,i))}return t},q=(0,n.createContext)({}),J=w&&window.__SWR_DEVTOOLS_USE__,Z=J?window.__SWR_DEVTOOLS_USE__:[],B=e=>c(e[1])?[e[0],e[1],e[2]||{}]:[e[0],null,(null===e[1]?e[2]:e[1])||{}],H=()=>u(F,(0,n.useContext)(q)),K=Z.concat(e=>(r,t,o)=>{let n=t&&((...e)=>{let[o]=N(r),[,,,n]=m.get(W);if(o.startsWith("$inf$"))return t(...e);let l=n[o];return d(l)?t(...e):(delete n[o],l)});return e(r,n,o)}),Q=(e,r,t)=>{let o=r[e]||(r[e]=[]);return o.push(t),()=>{let e=o.indexOf(t);e>=0&&(o[e]=o[o.length-1],o.pop())}};J&&(window.__SWR_DEVTOOLS_REACT__=n);let X=n.use||(e=>{if("pending"===e.status)throw e;if("fulfilled"===e.status)return e.value;if("rejected"===e.status)throw e.reason;throw e.status="pending",e.then(r=>{e.status="fulfilled",e.value=r},r=>{e.status="rejected",e.reason=r}),e}),Y={dedupe:!0};s.defineProperty(e=>{let{value:r}=e,t=(0,n.useContext)(q),o=c(r),l=(0,n.useMemo)(()=>o?r(t):r,[o,t,r]),i=(0,n.useMemo)(()=>o?l:U(t,l),[o,t,l]),s=l&&l.provider,d=(0,n.useRef)(a);s&&!d.current&&(d.current=G(s(i.cache||W),l));let p=d.current;return p&&(i.cache=p[0],i.mutate=p[1]),V(()=>{if(p)return p[2]&&p[2](),p[3]},[]),(0,n.createElement)(q.Provider,u(e,{value:i}))},"defaultValue",{value:F});let ee=(o=(e,r,t)=>{let{cache:o,compare:i,suspense:s,fallbackData:p,revalidateOnMount:f,revalidateIfStale:b,refreshInterval:g,refreshWhenHidden:h,refreshWhenOffline:y,keepPreviousData:v}=t,[w,x,k,_]=m.get(o),[S,R]=N(e),z=(0,n.useRef)(!1),j=(0,n.useRef)(!1),L=(0,n.useRef)(S),P=(0,n.useRef)(r),M=(0,n.useRef)(t),G=()=>M.current,W=()=>G().isVisible()&&G().isOnline(),[$,F,U,q]=E(o,S),J=(0,n.useRef)({}).current,Z=d(p)?t.fallback[S]:p,B=(e,r)=>{for(let t in J)if("data"===t){if(!i(e[t],r[t])&&(!d(e[t])||!i(ei,r[t])))return!1}else if(r[t]!==e[t])return!1;return!0},H=(0,n.useMemo)(()=>{let e=!!S&&!!r&&(d(f)?!G().isPaused()&&!s&&(!!d(b)||b):f),t=r=>{let t=u(r);return(delete t._k,e)?{isValidating:!0,isLoading:!0,...t}:t},o=$(),n=q(),l=t(o),i=o===n?l:t(n),a=l;return[()=>{let e=t($());return B(e,a)?(a.data=e.data,a.isLoading=e.isLoading,a.isValidating=e.isValidating,a.error=e.error,a):(a=e,e)},()=>i]},[o,S]),K=(0,l.useSyncExternalStore)((0,n.useCallback)(e=>U(S,(r,t)=>{B(t,r)||e()}),[o,S]),H[0],H[1]),ee=!z.current,er=w[S]&&w[S].length>0,et=K.data,eo=d(et)?Z:et,en=K.error,el=(0,n.useRef)(eo),ei=v?d(et)?el.current:et:eo,ea=(!er||!!d(en))&&(ee&&!d(f)?f:!G().isPaused()&&(s?!d(eo)&&b:d(eo)||b)),es=!!(S&&r&&ee&&ea),ed=d(K.isValidating)?es:K.isValidating,ec=d(K.isLoading)?es:K.isLoading,eu=(0,n.useCallback)(async e=>{let r,o;let n=P.current;if(!S||!n||j.current||G().isPaused())return!1;let l=!0,s=e||{},u=!k[S]||!s.dedupe,p=()=>O?!j.current&&S===L.current&&z.current:S===L.current,f={isValidating:!1,isLoading:!1},b=()=>{F(f)},g=()=>{let e=k[S];e&&e[1]===o&&delete k[S]},m={isValidating:!0};d($().data)&&(m.isLoading=!0);try{if(u&&(F(m),t.loadingTimeout&&d($().data)&&setTimeout(()=>{l&&p()&&G().onLoadingSlow(S,t)},t.loadingTimeout),k[S]=[n(R),I()]),[r,o]=k[S],r=await r,u&&setTimeout(g,t.dedupingInterval),!k[S]||k[S][1]!==o)return u&&p()&&G().onDiscarded(S),!1;f.error=a;let e=x[S];if(!d(e)&&(o<=e[0]||o<=e[1]||0===e[1]))return b(),u&&p()&&G().onDiscarded(S),!1;let s=$().data;f.data=i(s,r)?s:r,u&&p()&&G().onSuccess(r,S,t)}catch(t){g();let e=G(),{shouldRetryOnError:r}=e;!e.isPaused()&&(f.error=t,u&&p()&&(e.onError(t,S,e),(!0===r||c(r)&&r(t))&&(!G().revalidateOnFocus||!G().revalidateOnReconnect||W())&&e.onErrorRetry(t,S,e,e=>{let r=w[S];r&&r[0]&&r[0](A.ERROR_REVALIDATE_EVENT,e)},{retryCount:(s.retryCount||0)+1,dedupe:!0})))}return l=!1,b(),!0},[S,o]),ep=(0,n.useCallback)((...e)=>D(o,L.current,...e),[]);if(V(()=>{P.current=r,M.current=t,d(et)||(el.current=et)}),V(()=>{if(!S)return;let e=eu.bind(a,Y),r=0,t=Q(S,w,(t,o={})=>{if(t==A.FOCUS_EVENT){let t=Date.now();G().revalidateOnFocus&&t>r&&W()&&(r=t+G().focusThrottleInterval,e())}else if(t==A.RECONNECT_EVENT)G().revalidateOnReconnect&&W()&&e();else if(t==A.MUTATE_EVENT)return eu();else if(t==A.ERROR_REVALIDATE_EVENT)return eu(o)});return j.current=!1,L.current=S,z.current=!0,F({_k:R}),ea&&(d(eo)||T?e():C(e)),()=>{j.current=!0,t()}},[S]),V(()=>{let e;function r(){let r=c(g)?g($().data):g;r&&-1!==e&&(e=setTimeout(t,r))}function t(){!$().error&&(h||G().isVisible())&&(y||G().isOnline())?eu(Y).then(r):r()}return r(),()=>{e&&(clearTimeout(e),e=-1)}},[g,h,y,S]),(0,n.useDebugValue)(ei),s&&d(eo)&&S){if(!O&&T)throw Error("Fallback data is required when using suspense in SSR.");P.current=r,M.current=t,j.current=!1;let e=_[S];if(d(e)||X(ep(e)),d(en)){let e=eu(Y);d(ei)||(e.status="fulfilled",e.value=!0),X(e)}else throw en}return{mutate:ep,get data(){return J.data=!0,ei},get error(){return J.error=!0,en},get isValidating(){return J.isValidating=!0,ed},get isLoading(){return J.isLoading=!0,ec}}},function(...e){let r=H(),[t,n,l]=B(e),i=U(r,l),a=o,{use:s}=i,d=(s||[]).concat(K);for(let e=d.length;e--;)a=d[e](a);return a(t,n||i.fetcher||null,i)})},8388:function(e,r,t){t.d(r,{m6:function(){return B}});let o=e=>{let r=a(e),{conflictingClassGroups:t,conflictingClassGroupModifiers:o}=e;return{getClassGroupId:e=>{let t=e.split("-");return""===t[0]&&1!==t.length&&t.shift(),n(t,r)||i(e)},getConflictingClassGroupIds:(e,r)=>{let n=t[e]||[];return r&&o[e]?[...n,...o[e]]:n}}},n=(e,r)=>{if(0===e.length)return r.classGroupId;let t=e[0],o=r.nextPart.get(t),l=o?n(e.slice(1),o):void 0;if(l)return l;if(0===r.validators.length)return;let i=e.join("-");return r.validators.find(({validator:e})=>e(i))?.classGroupId},l=/^\[(.+)\]$/,i=e=>{if(l.test(e)){let r=l.exec(e)[1],t=r?.substring(0,r.indexOf(":"));if(t)return"arbitrary.."+t}},a=e=>{let{theme:r,prefix:t}=e,o={nextPart:new Map,validators:[]};return u(Object.entries(e.classGroups),t).forEach(([e,t])=>{s(t,o,e,r)}),o},s=(e,r,t,o)=>{e.forEach(e=>{if("string"==typeof e){(""===e?r:d(r,e)).classGroupId=t;return}if("function"==typeof e){if(c(e)){s(e(o),r,t,o);return}r.validators.push({validator:e,classGroupId:t});return}Object.entries(e).forEach(([e,n])=>{s(n,d(r,e),t,o)})})},d=(e,r)=>{let t=e;return r.split("-").forEach(e=>{t.nextPart.has(e)||t.nextPart.set(e,{nextPart:new Map,validators:[]}),t=t.nextPart.get(e)}),t},c=e=>e.isThemeGetter,u=(e,r)=>r?e.map(([e,t])=>[e,t.map(e=>"string"==typeof e?r+e:"object"==typeof e?Object.fromEntries(Object.entries(e).map(([e,t])=>[r+e,t])):e)]):e,p=e=>{if(e<1)return{get:()=>void 0,set:()=>{}};let r=0,t=new Map,o=new Map,n=(n,l)=>{t.set(n,l),++r>e&&(r=0,o=t,t=new Map)};return{get(e){let r=t.get(e);return void 0!==r?r:void 0!==(r=o.get(e))?(n(e,r),r):void 0},set(e,r){t.has(e)?t.set(e,r):n(e,r)}}},f=e=>{let{separator:r,experimentalParseClassName:t}=e,o=1===r.length,n=r[0],l=r.length,i=e=>{let t;let i=[],a=0,s=0;for(let d=0;d<e.length;d++){let c=e[d];if(0===a){if(c===n&&(o||e.slice(d,d+l)===r)){i.push(e.slice(s,d)),s=d+l;continue}if("/"===c){t=d;continue}}"["===c?a++:"]"===c&&a--}let d=0===i.length?e:e.substring(s),c=d.startsWith("!"),u=c?d.substring(1):d;return{modifiers:i,hasImportantModifier:c,baseClassName:u,maybePostfixModifierPosition:t&&t>s?t-s:void 0}};return t?e=>t({className:e,parseClassName:i}):i},b=e=>{if(e.length<=1)return e;let r=[],t=[];return e.forEach(e=>{"["===e[0]?(r.push(...t.sort(),e),t=[]):t.push(e)}),r.push(...t.sort()),r},g=e=>({cache:p(e.cacheSize),parseClassName:f(e),...o(e)}),m=(e,r)=>{let{parseClassName:t,getClassGroupId:o,getConflictingClassGroupIds:n}=r,l=[],i="";for(let r=e.length-1;r>=0;){for(;" "===e[r];)--r;let a=e.lastIndexOf(" ",r),s=e.slice(-1===a?0:a+1,r+1);r=a;let{modifiers:d,hasImportantModifier:c,baseClassName:u,maybePostfixModifierPosition:p}=t(s),f=!!p,g=o(f?u.substring(0,p):u);if(!g){if(!f||!(g=o(u))){i=s+(i.length>0?" "+i:i);continue}f=!1}let m=b(d).join(":"),h=c?m+"!":m,y=h+g;if(l.includes(y))continue;l.push(y);let v=n(g,f);for(let e=0;e<v.length;++e){let r=v[e];l.push(h+r)}i=s+(i.length>0?" "+i:i)}return i};function h(){let e,r,t=0,o="";for(;t<arguments.length;)(e=arguments[t++])&&(r=y(e))&&(o&&(o+=" "),o+=r);return o}let y=e=>{let r;if("string"==typeof e)return e;let t="";for(let o=0;o<e.length;o++)e[o]&&(r=y(e[o]))&&(t&&(t+=" "),t+=r);return t},v=e=>{let r=r=>r[e]||[];return r.isThemeGetter=!0,r},w=/^\[(?:([a-z-]+):)?(.+)\]$/i,x=/^\d+\/\d+$/,k=new Set(["px","full","screen"]),E=/^(\d+(\.\d+)?)?(xs|sm|md|lg|xl)$/,_=/\d+(%|px|r?em|[sdl]?v([hwib]|min|max)|pt|pc|in|cm|mm|cap|ch|ex|r?lh|cq(w|h|i|b|min|max))|\b(calc|min|max|clamp)\(.+\)|^0$/,S=/^(rgba?|hsla?|hwb|(ok)?(lab|lch))\(.+\)$/,R=/^(inset_)?-?((\d+)?\.?(\d+)[a-z]+|0)_-?((\d+)?\.?(\d+)[a-z]+|0)/,z=/^(url|image|image-set|cross-fade|element|(repeating-)?(linear|radial|conic)-gradient)\(.+\)$/,O=e=>C(e)||k.has(e)||x.test(e),T=e=>F(e,"length",U),C=e=>!!e&&!Number.isNaN(Number(e)),V=e=>F(e,"number",C),j=e=>!!e&&Number.isInteger(Number(e)),L=e=>e.endsWith("%")&&C(e.slice(0,-1)),N=e=>w.test(e),P=e=>E.test(e),I=new Set(["length","size","percentage"]),A=e=>F(e,I,q),D=e=>F(e,"position",q),M=new Set(["image","url"]),G=e=>F(e,M,Z),W=e=>F(e,"",J),$=()=>!0,F=(e,r,t)=>{let o=w.exec(e);return!!o&&(o[1]?"string"==typeof r?o[1]===r:r.has(o[1]):t(o[2]))},U=e=>_.test(e)&&!S.test(e),q=()=>!1,J=e=>R.test(e),Z=e=>z.test(e),B=function(e,...r){let t,o,n;let l=function(a){return o=(t=g(r.reduce((e,r)=>r(e),e()))).cache.get,n=t.cache.set,l=i,i(a)};function i(e){let r=o(e);if(r)return r;let l=m(e,t);return n(e,l),l}return function(){return l(h.apply(null,arguments))}}(()=>{let e=v("colors"),r=v("spacing"),t=v("blur"),o=v("brightness"),n=v("borderColor"),l=v("borderRadius"),i=v("borderSpacing"),a=v("borderWidth"),s=v("contrast"),d=v("grayscale"),c=v("hueRotate"),u=v("invert"),p=v("gap"),f=v("gradientColorStops"),b=v("gradientColorStopPositions"),g=v("inset"),m=v("margin"),h=v("opacity"),y=v("padding"),w=v("saturate"),x=v("scale"),k=v("sepia"),E=v("skew"),_=v("space"),S=v("translate"),R=()=>["auto","contain","none"],z=()=>["auto","hidden","clip","visible","scroll"],I=()=>["auto",N,r],M=()=>[N,r],F=()=>["",O,T],U=()=>["auto",C,N],q=()=>["bottom","center","left","left-bottom","left-top","right","right-bottom","right-top","top"],J=()=>["solid","dashed","dotted","double","none"],Z=()=>["normal","multiply","screen","overlay","darken","lighten","color-dodge","color-burn","hard-light","soft-light","difference","exclusion","hue","saturation","color","luminosity"],B=()=>["start","end","center","between","around","evenly","stretch"],H=()=>["","0",N],K=()=>["auto","avoid","all","avoid-page","page","left","right","column"],Q=()=>[C,N];return{cacheSize:500,separator:":",theme:{colors:[$],spacing:[O,T],blur:["none","",P,N],brightness:Q(),borderColor:[e],borderRadius:["none","","full",P,N],borderSpacing:M(),borderWidth:F(),contrast:Q(),grayscale:H(),hueRotate:Q(),invert:H(),gap:M(),gradientColorStops:[e],gradientColorStopPositions:[L,T],inset:I(),margin:I(),opacity:Q(),padding:M(),saturate:Q(),scale:Q(),sepia:H(),skew:Q(),space:M(),translate:M()},classGroups:{aspect:[{aspect:["auto","square","video",N]}],container:["container"],columns:[{columns:[P]}],"break-after":[{"break-after":K()}],"break-before":[{"break-before":K()}],"break-inside":[{"break-inside":["auto","avoid","avoid-page","avoid-column"]}],"box-decoration":[{"box-decoration":["slice","clone"]}],box:[{box:["border","content"]}],display:["block","inline-block","inline","flex","inline-flex","table","inline-table","table-caption","table-cell","table-column","table-column-group","table-footer-group","table-header-group","table-row-group","table-row","flow-root","grid","inline-grid","contents","list-item","hidden"],float:[{float:["right","left","none","start","end"]}],clear:[{clear:["left","right","both","none","start","end"]}],isolation:["isolate","isolation-auto"],"object-fit":[{object:["contain","cover","fill","none","scale-down"]}],"object-position":[{object:[...q(),N]}],overflow:[{overflow:z()}],"overflow-x":[{"overflow-x":z()}],"overflow-y":[{"overflow-y":z()}],overscroll:[{overscroll:R()}],"overscroll-x":[{"overscroll-x":R()}],"overscroll-y":[{"overscroll-y":R()}],position:["static","fixed","absolute","relative","sticky"],inset:[{inset:[g]}],"inset-x":[{"inset-x":[g]}],"inset-y":[{"inset-y":[g]}],start:[{start:[g]}],end:[{end:[g]}],top:[{top:[g]}],right:[{right:[g]}],bottom:[{bottom:[g]}],left:[{left:[g]}],visibility:["visible","invisible","collapse"],z:[{z:["auto",j,N]}],basis:[{basis:I()}],"flex-direction":[{flex:["row","row-reverse","col","col-reverse"]}],"flex-wrap":[{flex:["wrap","wrap-reverse","nowrap"]}],flex:[{flex:["1","auto","initial","none",N]}],grow:[{grow:H()}],shrink:[{shrink:H()}],order:[{order:["first","last","none",j,N]}],"grid-cols":[{"grid-cols":[$]}],"col-start-end":[{col:["auto",{span:["full",j,N]},N]}],"col-start":[{"col-start":U()}],"col-end":[{"col-end":U()}],"grid-rows":[{"grid-rows":[$]}],"row-start-end":[{row:["auto",{span:[j,N]},N]}],"row-start":[{"row-start":U()}],"row-end":[{"row-end":U()}],"grid-flow":[{"grid-flow":["row","col","dense","row-dense","col-dense"]}],"auto-cols":[{"auto-cols":["auto","min","max","fr",N]}],"auto-rows":[{"auto-rows":["auto","min","max","fr",N]}],gap:[{gap:[p]}],"gap-x":[{"gap-x":[p]}],"gap-y":[{"gap-y":[p]}],"justify-content":[{justify:["normal",...B()]}],"justify-items":[{"justify-items":["start","end","center","stretch"]}],"justify-self":[{"justify-self":["auto","start","end","center","stretch"]}],"align-content":[{content:["normal",...B(),"baseline"]}],"align-items":[{items:["start","end","center","baseline","stretch"]}],"align-self":[{self:["auto","start","end","center","stretch","baseline"]}],"place-content":[{"place-content":[...B(),"baseline"]}],"place-items":[{"place-items":["start","end","center","baseline","stretch"]}],"place-self":[{"place-self":["auto","start","end","center","stretch"]}],p:[{p:[y]}],px:[{px:[y]}],py:[{py:[y]}],ps:[{ps:[y]}],pe:[{pe:[y]}],pt:[{pt:[y]}],pr:[{pr:[y]}],pb:[{pb:[y]}],pl:[{pl:[y]}],m:[{m:[m]}],mx:[{mx:[m]}],my:[{my:[m]}],ms:[{ms:[m]}],me:[{me:[m]}],mt:[{mt:[m]}],mr:[{mr:[m]}],mb:[{mb:[m]}],ml:[{ml:[m]}],"space-x":[{"space-x":[_]}],"space-x-reverse":["space-x-reverse"],"space-y":[{"space-y":[_]}],"space-y-reverse":["space-y-reverse"],w:[{w:["auto","min","max","fit","svw","lvw","dvw",N,r]}],"min-w":[{"min-w":[N,r,"min","max","fit"]}],"max-w":[{"max-w":[N,r,"none","full","min","max","fit","prose",{screen:[P]},P]}],h:[{h:[N,r,"auto","min","max","fit","svh","lvh","dvh"]}],"min-h":[{"min-h":[N,r,"min","max","fit","svh","lvh","dvh"]}],"max-h":[{"max-h":[N,r,"min","max","fit","svh","lvh","dvh"]}],size:[{size:[N,r,"auto","min","max","fit"]}],"font-size":[{text:["base",P,T]}],"font-smoothing":["antialiased","subpixel-antialiased"],"font-style":["italic","not-italic"],"font-weight":[{font:["thin","extralight","light","normal","medium","semibold","bold","extrabold","black",V]}],"font-family":[{font:[$]}],"fvn-normal":["normal-nums"],"fvn-ordinal":["ordinal"],"fvn-slashed-zero":["slashed-zero"],"fvn-figure":["lining-nums","oldstyle-nums"],"fvn-spacing":["proportional-nums","tabular-nums"],"fvn-fraction":["diagonal-fractions","stacked-fractons"],tracking:[{tracking:["tighter","tight","normal","wide","wider","widest",N]}],"line-clamp":[{"line-clamp":["none",C,V]}],leading:[{leading:["none","tight","snug","normal","relaxed","loose",O,N]}],"list-image":[{"list-image":["none",N]}],"list-style-type":[{list:["none","disc","decimal",N]}],"list-style-position":[{list:["inside","outside"]}],"placeholder-color":[{placeholder:[e]}],"placeholder-opacity":[{"placeholder-opacity":[h]}],"text-alignment":[{text:["left","center","right","justify","start","end"]}],"text-color":[{text:[e]}],"text-opacity":[{"text-opacity":[h]}],"text-decoration":["underline","overline","line-through","no-underline"],"text-decoration-style":[{decoration:[...J(),"wavy"]}],"text-decoration-thickness":[{decoration:["auto","from-font",O,T]}],"underline-offset":[{"underline-offset":["auto",O,N]}],"text-decoration-color":[{decoration:[e]}],"text-transform":["uppercase","lowercase","capitalize","normal-case"],"text-overflow":["truncate","text-ellipsis","text-clip"],"text-wrap":[{text:["wrap","nowrap","balance","pretty"]}],indent:[{indent:M()}],"vertical-align":[{align:["baseline","top","middle","bottom","text-top","text-bottom","sub","super",N]}],whitespace:[{whitespace:["normal","nowrap","pre","pre-line","pre-wrap","break-spaces"]}],break:[{break:["normal","words","all","keep"]}],hyphens:[{hyphens:["none","manual","auto"]}],content:[{content:["none",N]}],"bg-attachment":[{bg:["fixed","local","scroll"]}],"bg-clip":[{"bg-clip":["border","padding","content","text"]}],"bg-opacity":[{"bg-opacity":[h]}],"bg-origin":[{"bg-origin":["border","padding","content"]}],"bg-position":[{bg:[...q(),D]}],"bg-repeat":[{bg:["no-repeat",{repeat:["","x","y","round","space"]}]}],"bg-size":[{bg:["auto","cover","contain",A]}],"bg-image":[{bg:["none",{"gradient-to":["t","tr","r","br","b","bl","l","tl"]},G]}],"bg-color":[{bg:[e]}],"gradient-from-pos":[{from:[b]}],"gradient-via-pos":[{via:[b]}],"gradient-to-pos":[{to:[b]}],"gradient-from":[{from:[f]}],"gradient-via":[{via:[f]}],"gradient-to":[{to:[f]}],rounded:[{rounded:[l]}],"rounded-s":[{"rounded-s":[l]}],"rounded-e":[{"rounded-e":[l]}],"rounded-t":[{"rounded-t":[l]}],"rounded-r":[{"rounded-r":[l]}],"rounded-b":[{"rounded-b":[l]}],"rounded-l":[{"rounded-l":[l]}],"rounded-ss":[{"rounded-ss":[l]}],"rounded-se":[{"rounded-se":[l]}],"rounded-ee":[{"rounded-ee":[l]}],"rounded-es":[{"rounded-es":[l]}],"rounded-tl":[{"rounded-tl":[l]}],"rounded-tr":[{"rounded-tr":[l]}],"rounded-br":[{"rounded-br":[l]}],"rounded-bl":[{"rounded-bl":[l]}],"border-w":[{border:[a]}],"border-w-x":[{"border-x":[a]}],"border-w-y":[{"border-y":[a]}],"border-w-s":[{"border-s":[a]}],"border-w-e":[{"border-e":[a]}],"border-w-t":[{"border-t":[a]}],"border-w-r":[{"border-r":[a]}],"border-w-b":[{"border-b":[a]}],"border-w-l":[{"border-l":[a]}],"border-opacity":[{"border-opacity":[h]}],"border-style":[{border:[...J(),"hidden"]}],"divide-x":[{"divide-x":[a]}],"divide-x-reverse":["divide-x-reverse"],"divide-y":[{"divide-y":[a]}],"divide-y-reverse":["divide-y-reverse"],"divide-opacity":[{"divide-opacity":[h]}],"divide-style":[{divide:J()}],"border-color":[{border:[n]}],"border-color-x":[{"border-x":[n]}],"border-color-y":[{"border-y":[n]}],"border-color-t":[{"border-t":[n]}],"border-color-r":[{"border-r":[n]}],"border-color-b":[{"border-b":[n]}],"border-color-l":[{"border-l":[n]}],"divide-color":[{divide:[n]}],"outline-style":[{outline:["",...J()]}],"outline-offset":[{"outline-offset":[O,N]}],"outline-w":[{outline:[O,T]}],"outline-color":[{outline:[e]}],"ring-w":[{ring:F()}],"ring-w-inset":["ring-inset"],"ring-color":[{ring:[e]}],"ring-opacity":[{"ring-opacity":[h]}],"ring-offset-w":[{"ring-offset":[O,T]}],"ring-offset-color":[{"ring-offset":[e]}],shadow:[{shadow:["","inner","none",P,W]}],"shadow-color":[{shadow:[$]}],opacity:[{opacity:[h]}],"mix-blend":[{"mix-blend":[...Z(),"plus-lighter","plus-darker"]}],"bg-blend":[{"bg-blend":Z()}],filter:[{filter:["","none"]}],blur:[{blur:[t]}],brightness:[{brightness:[o]}],contrast:[{contrast:[s]}],"drop-shadow":[{"drop-shadow":["","none",P,N]}],grayscale:[{grayscale:[d]}],"hue-rotate":[{"hue-rotate":[c]}],invert:[{invert:[u]}],saturate:[{saturate:[w]}],sepia:[{sepia:[k]}],"backdrop-filter":[{"backdrop-filter":["","none"]}],"backdrop-blur":[{"backdrop-blur":[t]}],"backdrop-brightness":[{"backdrop-brightness":[o]}],"backdrop-contrast":[{"backdrop-contrast":[s]}],"backdrop-grayscale":[{"backdrop-grayscale":[d]}],"backdrop-hue-rotate":[{"backdrop-hue-rotate":[c]}],"backdrop-invert":[{"backdrop-invert":[u]}],"backdrop-opacity":[{"backdrop-opacity":[h]}],"backdrop-saturate":[{"backdrop-saturate":[w]}],"backdrop-sepia":[{"backdrop-sepia":[k]}],"border-collapse":[{border:["collapse","separate"]}],"border-spacing":[{"border-spacing":[i]}],"border-spacing-x":[{"border-spacing-x":[i]}],"border-spacing-y":[{"border-spacing-y":[i]}],"table-layout":[{table:["auto","fixed"]}],caption:[{caption:["top","bottom"]}],transition:[{transition:["none","all","","colors","opacity","shadow","transform",N]}],duration:[{duration:Q()}],ease:[{ease:["linear","in","out","in-out",N]}],delay:[{delay:Q()}],animate:[{animate:["none","spin","ping","pulse","bounce",N]}],transform:[{transform:["","gpu","none"]}],scale:[{scale:[x]}],"scale-x":[{"scale-x":[x]}],"scale-y":[{"scale-y":[x]}],rotate:[{rotate:[j,N]}],"translate-x":[{"translate-x":[S]}],"translate-y":[{"translate-y":[S]}],"skew-x":[{"skew-x":[E]}],"skew-y":[{"skew-y":[E]}],"transform-origin":[{origin:["center","top","top-right","right","bottom-right","bottom","bottom-left","left","top-left",N]}],accent:[{accent:["auto",e]}],appearance:[{appearance:["none","auto"]}],cursor:[{cursor:["auto","default","pointer","wait","text","move","help","not-allowed","none","context-menu","progress","cell","crosshair","vertical-text","alias","copy","no-drop","grab","grabbing","all-scroll","col-resize","row-resize","n-resize","e-resize","s-resize","w-resize","ne-resize","nw-resize","se-resize","sw-resize","ew-resize","ns-resize","nesw-resize","nwse-resize","zoom-in","zoom-out",N]}],"caret-color":[{caret:[e]}],"pointer-events":[{"pointer-events":["none","auto"]}],resize:[{resize:["none","y","x",""]}],"scroll-behavior":[{scroll:["auto","smooth"]}],"scroll-m":[{"scroll-m":M()}],"scroll-mx":[{"scroll-mx":M()}],"scroll-my":[{"scroll-my":M()}],"scroll-ms":[{"scroll-ms":M()}],"scroll-me":[{"scroll-me":M()}],"scroll-mt":[{"scroll-mt":M()}],"scroll-mr":[{"scroll-mr":M()}],"scroll-mb":[{"scroll-mb":M()}],"scroll-ml":[{"scroll-ml":M()}],"scroll-p":[{"scroll-p":M()}],"scroll-px":[{"scroll-px":M()}],"scroll-py":[{"scroll-py":M()}],"scroll-ps":[{"scroll-ps":M()}],"scroll-pe":[{"scroll-pe":M()}],"scroll-pt":[{"scroll-pt":M()}],"scroll-pr":[{"scroll-pr":M()}],"scroll-pb":[{"scroll-pb":M()}],"scroll-pl":[{"scroll-pl":M()}],"snap-align":[{snap:["start","end","center","align-none"]}],"snap-stop":[{snap:["normal","always"]}],"snap-type":[{snap:["none","x","y","both"]}],"snap-strictness":[{snap:["mandatory","proximity"]}],touch:[{touch:["auto","none","manipulation"]}],"touch-x":[{"touch-pan":["x","left","right"]}],"touch-y":[{"touch-pan":["y","up","down"]}],"touch-pz":["touch-pinch-zoom"],select:[{select:["none","text","all","auto"]}],"will-change":[{"will-change":["auto","scroll","contents","transform",N]}],fill:[{fill:[e,"none"]}],"stroke-w":[{stroke:[O,T,V]}],stroke:[{stroke:[e,"none"]}],sr:["sr-only","not-sr-only"],"forced-color-adjust":[{"forced-color-adjust":["auto","none"]}]},conflictingClassGroups:{overflow:["overflow-x","overflow-y"],overscroll:["overscroll-x","overscroll-y"],inset:["inset-x","inset-y","start","end","top","right","bottom","left"],"inset-x":["right","left"],"inset-y":["top","bottom"],flex:["basis","grow","shrink"],gap:["gap-x","gap-y"],p:["px","py","ps","pe","pt","pr","pb","pl"],px:["pr","pl"],py:["pt","pb"],m:["mx","my","ms","me","mt","mr","mb","ml"],mx:["mr","ml"],my:["mt","mb"],size:["w","h"],"font-size":["leading"],"fvn-normal":["fvn-ordinal","fvn-slashed-zero","fvn-figure","fvn-spacing","fvn-fraction"],"fvn-ordinal":["fvn-normal"],"fvn-slashed-zero":["fvn-normal"],"fvn-figure":["fvn-normal"],"fvn-spacing":["fvn-normal"],"fvn-fraction":["fvn-normal"],"line-clamp":["display","overflow"],rounded:["rounded-s","rounded-e","rounded-t","rounded-r","rounded-b","rounded-l","rounded-ss","rounded-se","rounded-ee","rounded-es","rounded-tl","rounded-tr","rounded-br","rounded-bl"],"rounded-s":["rounded-ss","rounded-es"],"rounded-e":["rounded-se","rounded-ee"],"rounded-t":["rounded-tl","rounded-tr"],"rounded-r":["rounded-tr","rounded-br"],"rounded-b":["rounded-br","rounded-bl"],"rounded-l":["rounded-tl","rounded-bl"],"border-spacing":["border-spacing-x","border-spacing-y"],"border-w":["border-w-s","border-w-e","border-w-t","border-w-r","border-w-b","border-w-l"],"border-w-x":["border-w-r","border-w-l"],"border-w-y":["border-w-t","border-w-b"],"border-color":["border-color-t","border-color-r","border-color-b","border-color-l"],"border-color-x":["border-color-r","border-color-l"],"border-color-y":["border-color-t","border-color-b"],"scroll-m":["scroll-mx","scroll-my","scroll-ms","scroll-me","scroll-mt","scroll-mr","scroll-mb","scroll-ml"],"scroll-mx":["scroll-mr","scroll-ml"],"scroll-my":["scroll-mt","scroll-mb"],"scroll-p":["scroll-px","scroll-py","scroll-ps","scroll-pe","scroll-pt","scroll-pr","scroll-pb","scroll-pl"],"scroll-px":["scroll-pr","scroll-pl"],"scroll-py":["scroll-pt","scroll-pb"],touch:["touch-x","touch-y","touch-pz"],"touch-x":["touch"],"touch-y":["touch"],"touch-pz":["touch"]},conflictingClassGroupModifiers:{"font-size":["leading"]}}})}}]);
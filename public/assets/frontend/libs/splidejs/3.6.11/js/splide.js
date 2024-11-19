function _defineProperties(n,t){for(var e=0;e<t.length;e++){var i=t[e];i.enumerable=i.enumerable||!1,i.configurable=!0,"value"in i&&(i.writable=!0),Object.defineProperty(n,i.key,i)}}function _createClass(n,t,e){return t&&_defineProperties(n.prototype,t),e&&_defineProperties(n,e),n}/*!
 * Splide.js
 * Version  : 3.6.11
 * License  : MIT
 * Copyright: 2022 Naotoshi Fujita
 */ !function(n,t){"object"==typeof exports&&"undefined"!=typeof module?module.exports=t():"function"==typeof define&&define.amd?define(t):(n="undefined"!=typeof globalThis?globalThis:n||self).Splide=t()}(this,function(){"use strict";var n="splide",t="data-"+n;function e(n){n.length=0}function i(n){return!s(n)&&"object"==typeof n}function o(n){return Array.isArray(n)}function r(n){return"string"==typeof n}function u(n){return void 0===n}function s(n){return null===n}function c(n){return n instanceof HTMLElement}function a(n){return o(n)?n:[n]}function f(n,t){a(n).forEach(t)}function l(n,t){return n.indexOf(t)>-1}function d(n,t){return n.push.apply(n,a(t)),n}var p=Array.prototype;function v(n,t,e){return p.slice.call(n,t,e)}function h(n,t,e){n&&f(t,function(t){t&&n.classList[e?"add":"remove"](t)})}function g(n,t){h(n,r(t)?t.split(" "):t,!0)}function $(n,t){f(t,n.appendChild.bind(n))}function m(n,t){f(n,function(n){var e=t.parentNode;e&&e.insertBefore(n,t)})}function y(n,t){return c(n)&&(n.msMatchesSelector||n.matches).call(n,t)}function _(n,t){return n?v(n.children).filter(function(n){return y(n,t)}):[]}function b(n,t){return t?_(n,t)[0]:n.firstElementChild}function x(n,t,e){if(n){var i=Object.keys(n);i=e?i.reverse():i;for(var o=0;o<i.length;o++){var r=i[o];if("__proto__"!==r&&!1===t(n[r],r))break}}return n}function w(n){return v(arguments,1).forEach(function(t){x(t,function(e,i){n[i]=t[i]})}),n}function E(n,t){return x(t,function(t,e){o(t)?n[e]=t.slice():i(t)?n[e]=E(i(n[e])?n[e]:{},t):n[e]=t}),n}function S(n,t){n&&f(t,function(t){n.removeAttribute(t)})}function k(n,t,e){i(t)?x(t,function(t,e){k(n,e,t)}):s(e)?S(n,t):n.setAttribute(t,String(e))}function P(n,t,e){var i=document.createElement(n);return t&&(r(t)?g(i,t):k(i,t)),e&&$(e,i),i}function C(n,t,e){if(u(e))return getComputedStyle(n)[t];if(!s(e)){var i=n.style;e=""+e,i[t]!==e&&(i[t]=e)}}function L(n,t){C(n,"display",t)}function A(n,t){return n.getAttribute(t)}function z(n,t){return n&&n.classList.contains(t)}function D(n){return n.getBoundingClientRect()}function I(n){f(n,function(n){n&&n.parentNode&&n.parentNode.removeChild(n)})}function N(n){return b(new DOMParser().parseFromString(n,"text/html").body)}function R(n,t){n.preventDefault(),t&&(n.stopPropagation(),n.stopImmediatePropagation())}function T(n,t){return n&&n.querySelector(t)}function M(n,t){return v(n.querySelectorAll(t))}function F(n,t){h(n,t,!1)}function W(n){return r(n)?n:n?n+"px":""}function B(t,e){if(void 0===e&&(e=""),!t)throw Error("["+n+"] "+e)}function O(n){setTimeout(n)}var j=function n(){};function X(n){return requestAnimationFrame(n)}var H=Math.min,G=Math.max,Y=Math.floor,q=Math.ceil,U=Math.abs;function K(n,t,e){return U(n-t)<e}function Q(n,t,e,i){var o=H(t,e),r=G(t,e);return i?o<n&&n<r:o<=n&&n<=r}function V(n,t,e){var i=H(t,e),o=G(t,e);return H(G(i,n),o)}function J(n){return+(n>0)-+(n<0)}function Z(n,t){return f(t,function(t){n=n.replace("%s",""+t)}),n}function nn(n){return n<10?"0"+n:""+n}var nt={},ne="mounted",ni="ready",no="move",nr="moved",nu="shifted",ns="click",nc="slide:keydown",na="refresh",nf="updated",nl="resize",nd="resized",np="repositioned",nv="scroll",nh="scrolled",ng="destroy",n$="navigation:mounted",nm="lazyload:loaded";function ny(n){var t=n.event,e={},i=[];function o(n,t,e){r(n,t,function(n,t){i=i.filter(function(i){return i[0]!==n||i[1]!==t||!!e&&i[2]!==e||(n.removeEventListener(t,i[2],i[3]),!1)})})}function r(n,t,e){f(n,function(n){n&&t.split(" ").forEach(e.bind(null,n))})}function u(){i=i.filter(function(n){return o(n[0],n[1])}),t.offBy(e)}return t.on(ng,u,e),{on:function n(i,o,r){t.on(i,o,e,r)},off:function n(i){t.off(i,e)},emit:t.emit,bind:function n(t,e,o,u){r(t,e,function(n,t){i.push([n,t,o,u]),n.addEventListener(t,o,u)})},unbind:o,destroy:u}}function n_(n,t,e,i){var o,r,u=Date.now,s=0,c=!0,a=0;function f(){if(!c){var r=u()-o;if(r>=n?(s=1,o=u()):s=r/n,e&&e(s),1===s&&(t(),i&&++a>=i))return l();X(f)}}function l(){c=!0}function d(){cancelAnimationFrame(r),s=0,r=0,c=!0}function p(t){n=t}function v(){return c}return{start:function t(e){e||d(),o=u()-(e?s*n:0),c=!1,X(f)},rewind:function n(){o=u(),s=0,e&&e(s)},pause:l,cancel:d,set:p,isPaused:v}}function n8(n,t){var e;return function i(){var o=arguments,r=this;e||(e=n_(t||0,function(){n.apply(r,o),e=null},null,1)).start()}}var nb={marginRight:["marginBottom","marginLeft"],autoWidth:["autoHeight"],fixedWidth:["fixedHeight"],paddingLeft:["paddingTop","paddingRight"],paddingRight:["paddingBottom","paddingLeft"],width:["height"],left:["top","right"],right:["bottom","left"],x:["y"],X:["Y"],Y:["X"],ArrowLeft:["ArrowUp","ArrowRight"],ArrowRight:["ArrowDown","ArrowLeft"]},nx=n,nw=n+"__slider",nE=n+"__track",nS=n+"__list",nk=n+"__slide",nP=nk+"--clone",nC=nk+"__container",nL=n+"__arrows",n2=n+"__arrow",nA=n2+"--prev",nz=n2+"--next",nD=n+"__pagination",nI=n+"__progress",nN=nI+"__bar",nR=n+"__autoplay",nT=n+"__play",n3=n+"__pause",nM="is-active",n0="is-prev",n1="is-next",nF="is-visible",nW="is-loading",nB=[nM,nF,n0,n1,nW],nO="role",nj="aria-controls",nX="aria-current",n4="aria-label",n6="aria-hidden",nH="tabindex",nG="aria-orientation",nY=[nO,nj,nX,n4,n6,nG,nH,"disabled"],nq="slide",nU="loop",n7="fade",nK=t+"-interval",nQ={passive:!1,capture:!0},nV="touchmove mousemove",nJ="touchend touchcancel mouseup",nZ=["Left","Right","Up","Down"],n5="keydown",n9=t+"-lazy",tn=n9+"-srcset",tt="["+n9+"], ["+tn+"]",te=[" ","Enter","Spacebar"],ti=Object.freeze({__proto__:null,Options:function n(e,i,o){var r,u,s,c=n8(f);function a(n){n&&removeEventListener("resize",c)}function f(){var n,t,i,c,f=(n=u,t=function(n){return n[1].matches},v(n).filter(t)[0]||[]);f[0]!==s&&(i=s=f[0],c=o.breakpoints[i]||r,c.destroy?(e.options=r,e.destroy("completely"===c.destroy)):(e.state.is(5)&&(a(!0),e.mount()),e.options=c))}return{setup:function n(){try{E(o,JSON.parse(A(e.root,t)))}catch(i){B(!1,i.message)}r=E({},o);var s=o.breakpoints;if(s){var c="min"===o.mediaQuery;u=Object.keys(s).sort(function(n,t){return c?+t-+n:+n-+t}).map(function(n){return[n,matchMedia("("+(c?"min":"max")+"-width:"+n+"px)")]}),f()}},mount:function n(){u&&addEventListener("resize",c)},destroy:a}},Direction:function n(t,e,i){return{resolve:function n(t,e){var o=i.direction;return nb[t]["rtl"!==o||e?"ttb"===o?0:-1:1]||t},orient:function n(t){return t*("rtl"===i.direction?1:-1)}}},Elements:function t(i,o,r){var u,s,c,a,f=ny(i).on,l=i.root,p={},v=[];function h(){var t,e,i,o;s=b(l,"."+nw),c=T(l,"."+nE),a=b(c,"."+nS),B(c&&a,"A track/list element is missing."),d(v,_(a,"."+nk+":not(."+nP+")")),t=x("."+nR),e=x("."+nL),w(p,{root:l,slider:s,track:c,list:a,slides:v,arrows:e,autoplay:t,prev:T(e,"."+nA),next:T(e,"."+nz),bar:T(x("."+nI),"."+nN),play:T(t,"."+nT),pause:T(t,"."+n3)}),o=l.id||""+(i=n)+nn(nt[i]=(nt[i]||0)+1),l.id=o,c.id=c.id||o+"-track",a.id=a.id||o+"-list",g(l,u=E())}function $(){[l,c,a].forEach(function(n){S(n,"style")}),e(v),F(l,u)}function m(){$(),h()}function y(){F(l,u),g(l,u=E())}function x(n){return b(l,n)||b(s,n)}function E(){return[nx+"--"+r.type,nx+"--"+r.direction,r.drag&&nx+"--draggable",r.isNavigation&&nx+"--nav",nM]}return w(p,{setup:h,mount:function n(){f(na,m,8),f(nf,y)},destroy:$})},Slides:function n(t,i,o){var u=ny(t),s=u.on,d=u.emit,p=u.bind,v=i.Elements,_=v.slides,x=v.list,w=[];function E(){_.forEach(function(n,t){R(n,t,-1)})}function P(){W(function(n){n.destroy()}),e(w)}function L(){P(),E()}function R(n,e,i){var o=function n(t,e,i,o){var r,u=ny(t),s=u.on,c=u.emit,a=u.bind,f=u.destroy,l=t.Components,d=t.root,p=t.options,v=p.isNavigation,g=p.updateOnMove,$=l.Direction.resolve,m=A(o,"style"),y=i>-1,_=b(o,"."+nC),x=p.focusableNodes&&M(o,p.focusableNodes);function w(){var n=Z(p.i18n.slideX,(y?i:e)+1),r=t.splides.map(function(n){return n.splide.root.id}).join(" ");k(o,n4,n),k(o,nj,r),k(o,nO,"menuitem"),L(I())}function E(){r||P()}function P(){if(!r){var n,i,u=t.index;L(I()),i=!(n=function n(){if(t.is(n7))return I();var e=D(l.Elements.track),i=D(o),r=$("left"),u=$("right");return Y(e[r])<=q(i[r])&&Y(i[u])<=q(e[u])}())&&(!I()||y),k(o,n6,i||null),k(o,nH,!i&&p.slideFocus?0:null),x&&x.forEach(function(n){k(n,nH,i?-1:null)}),n!==z(o,nF)&&(h(o,nF,n),c(n?"visible":"hidden",N)),h(o,n0,e===u-1),h(o,n1,e===u+1)}}function L(n){n!==z(o,nM)&&(h(o,nM,n),v&&k(o,nX,n||null),c(n?"active":"inactive",N))}function I(){var n=t.index;return n===e||p.cloneStatus&&n===i}var N={index:e,slideIndex:i,slide:o,container:_,isClone:y,mount:function n(){y||(o.id=d.id+"-slide"+nn(e+1)),a(o,"click keydown",function(n){c("click"===n.type?ns:nc,N,n)}),s([na,np,nu,nr,nh],P),s(n$,w),g&&s(no,E)},destroy:function n(){r=!0,f(),F(o,nB),S(o,nY),k(o,"style",m)},update:P,style:function n(t,e,i){C(i&&_||o,t,e)},isWithin:function n(i,o){var r=U(i-e);return!y&&(p.rewind||t.is(nU))&&(r=H(r,t.length-r)),r<=o}};return N}(t,e,i,n);o.mount(),w.push(o)}function T(n){return n?B(function(n){return!n.isClone}):w}function W(n,t){T(t).forEach(n)}function B(n){var t;return w.filter("function"==typeof(t=n)?n:function(t){return r(n)?y(t.slide,n):l(a(n),t.index)})}return{mount:function n(){E(),s(na,L),s([ne,na],function(){w.sort(function(n,t){return n.index-t.index})})},destroy:P,update:function n(){W(function(n){n.update()})},register:R,get:T,getIn:function n(t){var e=i.Controller,r=e.toIndex(t),u=e.hasFocus()?1:o.perPage;return B(function(n){return Q(n.index,r,r+u-1)})},getAt:function n(t){return B(t)[0]},add:function n(t,e){f(t,function(n){if(r(n)&&(n=N(n)),c(n)){var t,i,u,s,a=_[e];a?m(n,a):$(x,n),g(n,o.classes.slide),t=n,i=d.bind(null,nl),(s=(u=M(t,"img")).length)?u.forEach(function(n){p(n,"load error",function(){--s||i()})}):i()}}),d(na)},remove:function n(t){I(B(t).map(function(n){return n.slide})),d(na)},forEach:W,filter:B,style:function n(t,e,i){W(function(n){n.style(t,e,i)})},getLength:function n(t){return t?_.length:w.length},isEnough:function n(){return w.length>o.perPage}}},Layout:function n(t,e,o){var r,u,s=ny(t),c=s.on,a=s.bind,f=s.emit,l=e.Slides,d=e.Direction.resolve,p=e.Elements,v=p.root,h=p.track,g=p.list,$=l.getAt;function m(){u=null,r="ttb"===o.direction,C(v,"maxWidth",W(o.width)),C(h,d("paddingLeft"),_(!1)),C(h,d("paddingRight"),_(!0)),y()}function y(){var n,t=D(v);u&&u.width===t.width&&u.height===t.height||(C(h,"height",(n="",r&&(n=b(),B(n,"height or heightRatio is missing."),n="calc("+n+" - "+_(!1)+" - "+_(!0)+")"),n)),l.style(d("marginRight"),W(o.gap)),l.style("width",(o.autoWidth?"":W(o.fixedWidth)||(r?"":x()))||null),l.style("height",W(o.fixedHeight)||(r?o.autoHeight?"":x():b())||null,!0),u=t,f(nd))}function _(n){var t=o.padding,e=d(n?"right":"left");return t&&W(t[e]||(i(t)?0:t))||"0px"}function b(){return W(o.height||D(g).width*o.heightRatio)}function x(){var n=W(o.gap);return"calc((100%"+(n&&" + "+n)+")/"+(o.perPage||1)+(n&&" - "+n)+")"}function w(n,t){var e=$(n);if(e){var i=D(e.slide)[d("right")],o=D(g)[d("left")];return U(i-o)+(t?0:E())}return 0}function E(){var n=$(0);return n&&parseFloat(C(n.slide,d("marginRight")))||0}return{mount:function n(){m(),a(window,"resize load",n8(f.bind(this,nl))),c([nf,na],m),c(nl,y)},listSize:function n(){return D(g)[d("width")]},slideSize:function n(t,e){var i=$(t||0);return i?D(i.slide)[d("width")]+(e?0:E()):0},sliderSize:function n(){return w(t.length-1,!0)-w(-1,!0)},totalSize:w,getPadding:function n(t){return parseFloat(C(h,d("padding"+(t?"Right":"Left"))))||0}}},Clones:function n(t,i,o){var u,s=ny(t),c=s.on,a=s.emit,f=i.Elements,l=i.Slides,p=i.Direction.resolve,v=[];function h(){(u=x())&&(function n(e){var i=l.get().slice(),r=i.length;if(r){for(;i.length<e;)d(i,i);d(i.slice(-e),i.slice(0,e)).forEach(function(n,u){var s,c,a,p=u<e,h=(s=n.slide,c=u,a=s.cloneNode(!0),g(a,o.classes.clone),a.id=t.root.id+"-clone"+nn(c+1),a);p?m(h,i[0].slide):$(f.list,h),d(v,h),l.register(h,u-e+(p?0:r),n.index)})}}(u),a(nl))}function y(){I(v),e(v)}function _(){y(),h()}function b(){u<x()&&a(na)}function x(){var n=o.clones;if(t.is(nU)){if(!n){var e=function n(t,e){if(r(e)){var i=P("div",{style:"width: "+e+"; position: absolute;"},t);e=D(i).width,I(i)}return e}(f.list,o[p("fixedWidth")]);n=(e&&q(D(f.track)[p("width")]/e)||o[p("autoWidth")]&&t.length||o.perPage)*(o.drag?(o.flickMaxPages||1)+1:2)}}else n=0;return n}return{mount:function n(){h(),c(na,_),c([nf,nl],b)},destroy:y}},Move:function n(t,e,i){var o,r=ny(t),s=r.on,c=r.emit,a=e.Layout,f=a.slideSize,l=a.getPadding,d=a.totalSize,p=a.listSize,v=a.sliderSize,h=e.Direction,g=h.resolve,$=h.orient,m=e.Elements,y=m.list,_=m.track;function b(){L()||(e.Scroll.cancel(),x(t.index),c(np))}function x(n){w(k(n,!0))}function w(n,e){if(!t.is(n7)){var i=e?n:function n(e){if(t.is(nU)){var i=$(e-P()),o=A(!1,e)&&i<0,r=A(!0,e)&&i>0;(o||r)&&(e=E(e,r))}return e}(n);y.style.transform="translate"+g("X")+"("+i+"px)",n!==i&&c(nu)}}function E(n,t){var e=n-C(t),i=v();return n-$(i*(q(U(e)/i)||1))*(t?1:-1)}function k(n,e){var o,r,u,s=$(d(n-1)-(o=n,r=i.focus,"center"===r?(p()-f(o,!0))/2:+r*f(o)||0));return e?(u=s,i.trimSpace&&t.is(nq)&&(u=V(u,0,$(v()-p()))),u):s}function P(){var n=g("left");return D(y)[n]-D(_)[n]+$(l(!1))}function C(n){return k(n?e.Controller.getEnd():0,!!i.trimSpace)}function L(){return t.state.is(4)&&i.waitForTransition}function A(n,t){t=u(t)?P():t;var e=!0!==n&&$(t)<$(C(!1)),i=!1!==n&&$(t)>$(C(!0));return e||i}return{mount:function n(){o=e.Transition,s([ne,nd,nf,na],b)},destroy:function n(){S(y,"style")},move:function n(r,u,s,a){if(!L()){var f=t.state.set,l=P();r!==u&&(o.cancel(),w(E(l,r>u),!0)),f(4),c(no,u,s,r),o.start(u,function(){f(3),c(nr,u,s,r),"move"===i.trimSpace&&r!==s&&l===P()?e.Controller.go(r>s?">":"<",!1,a):a&&a()})}},jump:x,translate:w,shift:E,cancel:function n(){w(P()),o.cancel()},toIndex:function n(t){for(var i=e.Slides.get(),o=0,r=1/0,u=0;u<i.length;u++){var s=i[u].index,c=U(k(s,!0)-t);if(c<=r)r=c,o=s;else break}return o},toPosition:k,getPosition:P,getLimit:C,isBusy:L,exceededLimit:A}},Controller:function n(t,e,i){var o,s,c,a=ny(t).on,f=e.Move,l=f.getPosition,d=f.getLimit,p=e.Slides,v=p.isEnough,h=p.getLength,g=t.is(nU),$=t.is(nq),m=i.start||0,y=m;function _(){o=h(!0),s=i.perMove,c=i.perPage,m=V(m,0,o-1)}function b(n,t,i,o,r){var u=t?n:A(n);e.Scroll.scroll(t||i?f.toPosition(u,!0):n,o,function(){z(f.toIndex(f.getPosition())),r&&r()})}function x(n){return E(!1,n)}function w(n){return E(!0,n)}function E(n,t){var e,i,o=S(m+(s||(D()?1:c))*(n?-1:1),m);if(-1===o&&$){if(e=l(),!(1>U(e-(i=d(!n)))))return n?0:k()}return t?o:P(o)}function S(n,t,e){if(v()){var r=k();n<0||n>r?n=Q(0,n,t,!0)||Q(r,t,n,!0)?C(L(n)):g?s||D()?n:n<0?-(o%c||c):o:i.rewind?n<0?r:0:-1:e||n===t||(n=s?n:C(L(t)+(n<t?-1:1)))}else n=-1;return n}function k(){var n=o-c;return(D()||g&&s)&&(n=o-1),G(n,0)}function P(n){return g?v()?n%o+(n<0?o:0):-1:n}function C(n){return V(D()?n:c*n,0,k())}function L(n){return D()||(n=Y((n=Q(n,o-c,o-1)?o-1:n)/c)),n}function A(n){var t=f.toIndex(n);return $?V(t,0,k()):t}function z(n){n!==m&&(y=m,m=n)}function D(){return!u(i.focus)||i.isNavigation}return{mount:function n(){_(),a([nf,na],_,9)},go:function n(t,e,o){var u=function n(t){var e=m;if(r(t)){var i=t.match(/([+\-<>])(\d+)?/)||[],o=i[1],u=i[2];"+"===o||"-"===o?e=S(m+ +(""+o+(+u||1)),m,!0):">"===o?e=u?C(+u):x(!0):"<"===o&&(e=w(!0))}else e=g?t:V(t,0,k());return e}(t);if(i.useScroll)b(u,!0,!0,i.speed,o);else{var s=P(u);s>-1&&!f.isBusy()&&(e||s!==m)&&(z(s),f.move(u,s,y,o))}},scroll:b,getNext:x,getPrev:w,getAdjacent:E,getEnd:k,setIndex:z,getIndex:function n(t){return t?y:m},toIndex:C,toPage:L,toDest:A,hasFocus:D}},Arrows:function n(t,e,i){var o,r=ny(t),u=r.on,s=r.bind,c=r.emit,a=i.classes,f=i.i18n,l=e.Elements,d=e.Controller,p=l.arrows,v=l.prev,h=l.next,g={};function y(){if(!i.arrows||v&&h||(p=P("div",a.arrows),v=_(!0),h=_(!1),o=!0,$(p,[v,h]),m(p,b("slider"===i.arrows&&l.slider||t.root))),v&&h){if(g.prev)L(p,!1===i.arrows?"none":"");else{var n,e=l.track.id;k(v,nj,e),k(h,nj,e),g.prev=v,g.next=h,n=d.go,u([ne,nr,nf,na,nh],x),s(h,"click",function(){n(">",!0)}),s(v,"click",function(){n("<",!0)}),c("arrows:mounted",v,h)}}}function _(n){return N('<button class="'+a.arrow+" "+(n?a.prev:a.next)+'" type="button"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 40 40" width="40" height="40"><path d="'+(i.arrowPath||"m15.5 0.932-4.3 4.38 14.5 14.6-14.5 14.5 4.3 4.4 14.6-14.6 4.4-4.3-4.4-4.4-14.6-14.6z")+'" />')}function x(){var n=t.index,e=d.getPrev(),i=d.getNext(),o=e>-1&&n<e?f.last:f.prev,r=i>-1&&n>i?f.first:f.next;v.disabled=e<0,h.disabled=i<0,k(v,n4,o),k(h,n4,r),c("arrows:updated",v,h,e,i)}return{arrows:g,mount:function n(){y(),u(nf,y)},destroy:function n(){o?I(p):(S(v,nY),S(h,nY))}}},Autoplay:function n(t,e,i){var o,r,u,s=ny(t),c=s.on,a=s.bind,f=s.emit,l=n_(i.interval,t.go.bind(t,">"),function n(t){var e=p.bar;e&&C(e,"width",100*t+"%"),f("autoplay:playing",t)}),d=l.isPaused,p=e.Elements;function v(n){var t=n?"pause":"play",e=p[t];e&&(k(e,nj,p.track.id),k(e,n4,i.i18n[t]),a(e,"click",n?g:h))}function h(){d()&&e.Slides.isEnough()&&(l.start(!i.resetProgress),r=o=u=!1,f("autoplay:play"))}function g(n){void 0===n&&(n=!0),d()||(l.pause(),f("autoplay:pause")),u=n}function $(){u||(o||r?g(!1):h())}function m(){var n=e.Slides.getAt(t.index);l.set(n&&+A(n.slide,nK)||i.interval)}return{mount:function n(){var t,e=i.autoplay;e&&(v(!0),v(!1),t=p.root,i.pauseOnHover&&a(t,"mouseenter mouseleave",function(n){o="mouseenter"===n.type,$()}),i.pauseOnFocus&&a(t,"focusin focusout",function(n){r="focusin"===n.type,$()}),c([no,nv,na],l.rewind),c(no,m),"pause"!==e&&h())},destroy:l.cancel,play:h,pause:g,isPaused:d}},Cover:function n(t,e,i){var o=ny(t).on;function r(n){e.Slides.forEach(function(t){var e=b(t.container||t.slide,"img");e&&e.src&&u(n,e,t)})}function u(n,t,e){e.style("background",n?'center/cover no-repeat url("'+t.src+'")':"",!0),L(t,n?"none":"")}return{mount:function n(){i.cover&&(o(nm,function(n,t){u(!0,n,t)}),o([ne,nf,na],r.bind(null,!0)))},destroy:function n(){r(!1)}}},Scroll:function n(t,e,i){var o,r,u=ny(t),s=u.on,c=u.emit,a=e.Move,f=a.getPosition,l=a.getLimit,d=a.exceededLimit;function p(){var n=f(),e=a.toIndex(n);Q(e,0,t.length-1)||a.translate(a.shift(n,e>0),!0),r&&r(),c(nh)}function v(){o&&o.cancel()}function h(){o&&!o.isPaused()&&(v(),p())}return{mount:function n(){s(no,v),s([nf,na],h)},destroy:v,scroll:function n(e,u,s,h){var g,$=f(),m=1;u=u||(g=U(e-$),G(g/1.5,800)),r=s,v(),o=n_(u,p,function(o){var r,u,s,c=f(),p=($+(e-$)*(r=o,u=i.easingFunc,u?u(r):1-Math.pow(1-r,4))-f())*m;a.translate(c+p),t.is(nq)&&!h&&d()&&(m*=.6,10>U(p)&&(s=d(!1),n(l(!s),600,null,!0)))},1),c(nv),o.start()},cancel:h}},Drag:function n(t,e,o){var r,u,s,c,a,f,l,d,p,v=ny(t),h=v.on,g=v.emit,$=v.bind,m=v.unbind,_=e.Move,b=e.Scroll,x=e.Controller,w=e.Elements.track,E=e.Direction,S=E.resolve,k=E.orient,P=_.getPosition,C=_.exceededLimit,L=!1;function A(){var n,t=o.drag;d=n=!t,a="free"===t}function z(n){if(!d){var t=o.noDrag,e=B(n),i=!t||!y(n.target,t);l=!1,i&&(e||!n.button)&&(_.isBusy()?R(n,!0):(s=null,c=null,$(p=e?w:window,nV,D,nQ),$(p,nJ,I,nQ),_.cancel(),b.cancel(),N(n)))}}function D(n){if(c||g("drag"),c=n,n.cancelable){var e=F(n)-F(u);if(f){_.translate(r+(s=e,s/(L&&t.is(nq)?5:1)));var s,a=W(n)-W(u)>200,d=L!==(L=C());(a||d)&&N(n),g("dragging"),l=!0,R(n)}else{var p=o.dragMinThreshold;p=i(p)?p:{mouse:0,touch:+p||10},f=U(e)>(B(n)?p.touch:p.mouse),M()&&R(n)}}}function I(n){m(p,nV,D),m(p,nJ,I);var i=t.index;if(c){if(f||n.cancelable&&M()){var r,l=function n(e){if(t.is(nU)||!L){var i=u===c&&s||u,o=F(c)-F(i),r=W(e)-W(i),a=W(e)-W(c)<200;if(r&&a)return o/r}return 0}(n),d=(r=l,P()+J(r)*H(U(r)*(o.flickPower||600),a?1/0:e.Layout.listSize()*(o.flickMaxPages||1)));a?x.scroll(d):t.is(n7)?x.go(i+k(J(l))):x.go(x.toDest(d),!0),R(n)}g("dragged")}else a||P()===_.toPosition(i)||x.go(i,!0);f=!1}function N(n){s=u,u=n,r=P()}function T(n){!d&&l&&R(n,!0)}function M(){var n=U(F(c)-F(u)),t=U(F(c,!0)-F(u,!0));return n>t}function F(n,t){return(B(n)?n.touches[0]:n)["page"+S(t?"Y":"X")]}function W(n){return n.timeStamp}function B(n){return"undefined"!=typeof TouchEvent&&n instanceof TouchEvent}function O(){return f}function X(n){d=n}return{mount:function n(){$(w,nV,j,nQ),$(w,nJ,j,nQ),$(w,"touchstart mousedown",z,nQ),$(w,"click",T,{capture:!0}),$(w,"dragstart",R),h([ne,nf],A)},disable:X,isDragging:O}},Keyboard:function n(t,e,i){var o,r,u=ny(t),s=u.on,a=u.bind,f=u.unbind,d=t.root,p=e.Direction.resolve;function v(){var n=i.keyboard;n&&("focused"===n?(o=d,k(d,nH,0)):o=window,a(o,n5,y))}function h(){f(o,n5),c(o)&&S(o,nH)}function g(n){r=n}function $(){var n=r;r=!0,O(function(){r=n})}function m(){h(),v()}function y(n){if(!r){var e=n.key,i=l(nZ,e)?"Arrow"+e:e;i===p("ArrowLeft")?t.go("<"):i===p("ArrowRight")&&t.go(">")}}return{mount:function n(){v(),s(nf,m),s(no,$)},destroy:h,disable:g}},LazyLoad:function n(t,e,i){var o=ny(t),r=o.on,u=o.off,s=o.bind,c=o.emit,a="sequential"===i.lazyLoad,f=[],l=0;function d(){v(),p()}function p(){e.Slides.forEach(function(n){M(n.slide,tt).forEach(function(t){var e=A(t,n9),o=A(t,tn);if(e!==t.src||o!==t.srcset){var r=i.classes.spinner,u=t.parentElement,s=b(u,"."+r)||P("span",r,u);k(s,nO,"presentation"),f.push({_img:t,_Slide:n,src:e,srcset:o,_spinner:s}),t.src||L(t,"none")}})}),a&&m()}function v(){l=0,f=[]}function h(){(f=f.filter(function(n){var e=i.perPage*((i.preloadPages||1)+1)-1;return!n._Slide.isWithin(t.index,e)||$(n)})).length||u(nr)}function $(n){var t=n._img;g(n._Slide.slide,nW),s(t,"load error",function(t){var e,i,o;e=n,i="error"===t.type,o=e._Slide,F(o.slide,nW),i||(I(e._spinner),L(e._img,""),c(nm,e._img,o),c(nl)),a&&m()}),["src","srcset"].forEach(function(e){n[e]&&(k(t,e,n[e]),S(t,"src"===e?n9:tn))})}function m(){l<f.length&&$(f[l++])}return{mount:function n(){i.lazyLoad&&(p(),r(na,d),a||r([ne,na,nr,nh],h))},destroy:v}},Pagination:function n(t,i,o){var r,u=ny(t),s=u.on,c=u.emit,a=u.bind,f=u.unbind,l=i.Slides,d=i.Elements,p=i.Controller,v=p.hasFocus,h=p.getIndex,$=[];function m(){y(),o.pagination&&l.isEnough()&&(function n(){var e=t.length,i=o.classes,u=o.i18n,s=o.perPage,c="slider"===o.pagination&&d.slider||d.root,f=v()?e:q(e/s);r=P("ul",i.pagination,c);for(var p=0;p<f;p++){var h=P("li",null,r),g=P("button",{class:i.page,type:"button"},h),m=l.getIn(p).map(function(n){return n.slide.id}),y=!v()&&s>1?u.pageX:u.slideX;a(g,"click",_.bind(null,p)),k(g,nj,m.join(" ")),k(g,n4,Z(y,p+1)),$.push({li:h,button:g,page:p})}}(),c("pagination:mounted",{list:r,items:$},b(t.index)),x())}function y(){r&&(I(r),$.forEach(function(n){f(n.button,"click")}),e($),r=null)}function _(n){p.go(">"+n,!0,function(){var t,e=l.getAt(p.toIndex(n));e&&((t=e.slide).setActive&&t.setActive()||t.focus({preventScroll:!0}))})}function b(n){return $[p.toPage(n)]}function x(){var n=b(h(!0)),t=b(h());n&&(F(n.button,nM),S(n.button,nX)),t&&(g(t.button,nM),k(t.button,nX,!0)),c("pagination:updated",{list:r,items:$},n,t)}return{items:$,mount:function n(){m(),s([nf,na],m),s([no,nh],x)},destroy:y,getAt:b,update:x}},Sync:function n(t,i,o){var r=i.Elements.list,u=[];function s(){var n,e;t.splides.forEach(function(n){n.isParent||function n(e){[t,e].forEach(function(n){var i=ny(n),o=n===t?e:t;i.on(no,function(n,t,e){o.go(o.is(nU)?e:n)}),u.push(i)})}(n.splide)}),o.isNavigation&&(n=ny(t),e=n.on,e(ns,f),e(nc,d),e([ne,nf],a),k(r,nO,"menu"),u.push(n),n.emit(n$,t.splides))}function c(){S(r,nY),u.forEach(function(n){n.destroy()}),e(u)}function a(){k(r,nG,"ttb"!==o.direction?"horizontal":null)}function f(n){t.go(n.index)}function d(n,t){l(te,t.key)&&(f(n),R(t))}return{mount:s,destroy:c,remount:function n(){c(),s()}}},Wheel:function n(t,e,i){var o=ny(t).bind;function r(n){if(n.cancelable){var o=n.deltaY;if(o){var r,u=o<0;t.go(u?"<":">"),r=u,(!i.releaseWheel||t.state.is(4)||-1!==e.Controller.getAdjacent(r))&&R(n)}}}return{mount:function n(){i.wheel&&o(e.Elements.track,"wheel",r,nQ)}}}}),to={type:"slide",speed:400,waitForTransition:!0,perPage:1,cloneStatus:!0,arrows:!0,pagination:!0,interval:5e3,pauseOnHover:!0,pauseOnFocus:!0,resetProgress:!0,keyboard:!0,easing:"cubic-bezier(0.25, 1, 0.5, 1)",drag:!0,direction:"ltr",slideFocus:!0,trimSpace:!0,focusableNodes:"a, button, textarea, input, select, iframe",classes:{slide:nk,clone:nP,arrows:nL,arrow:n2,prev:nA,next:nz,pagination:nD,page:nD+"__page",spinner:n+"__spinner"},i18n:{prev:"Previous slide",next:"Next slide",first:"Go to first slide",last:"Go to last slide",slideX:"Go to slide %s",pageX:"Go to page %s",play:"Start autoplay",pause:"Pause autoplay"}};function tr(n,t,e){var i=ny(n).on;return{mount:function n(){i([ne,na],function(){O(function(){t.Slides.style("transition","opacity "+e.speed+"ms "+e.easing)})})},start:function n(e,i){var o=t.Elements.track;C(o,"height",W(D(o).height)),O(function(){i(),C(o,"height","")})},cancel:j}}function tu(n,t,e){var i,o=ny(n).bind,r=t.Move,u=t.Controller,s=t.Elements.list;function c(){a("")}function a(n){C(s,"transition",n)}return{mount:function n(){o(s,"transitionend",function(n){n.target===s&&i&&(c(),i())})},start:function t(o,s){var c=r.toPosition(o,!0),f=r.getPosition(),l=function t(i){var o=e.rewindSpeed;if(n.is(nq)&&o){var r=u.getIndex(!0),s=u.getEnd();if(0===r&&i>=s||r>=s&&0===i)return o}return e.speed}(o);U(c-f)>=1&&l>=1?(a("transform "+l+"ms "+e.easing),r.translate(c,!0),i=s):(r.jump(o),s())},cancel:c}}var ts=function(){function n(t,e){this.event=function n(){var t={};function e(n,e){i(n,function(n,i){var o=t[n];t[n]=o&&o.filter(function(n){return n._key?n._key!==e:e||n._namespace!==i})})}function i(n,t){a(n).join(" ").split(" ").forEach(function(n){var e=n.split(".");t(e[0],e[1])})}return{on:function n(e,o,r,u){void 0===u&&(u=10),i(e,function(n,e){t[n]=t[n]||[],d(t[n],{_event:n,_callback:o,_namespace:e,_priority:u,_key:r}).sort(function(n,t){return n._priority-t._priority})})},off:e,offBy:function n(i){x(t,function(n,t){e(t,i)})},emit:function n(e){var i=arguments;(t[e]||[]).forEach(function(n){n._callback.apply(n,v(i,1))})},destroy:function n(){t={}}}}(),this.Components={},this.state=function n(t){var e=1;function i(n){e=n}return{set:i,is:function n(t){return l(a(t),e)}}}(1),this.splides=[],this._options={},this._Extensions={};var i=r(t)?T(document,t):t;B(i,i+" is invalid."),this.root=i,E(to,n.defaults),E(E(this._options,to),e||{})}var t=n.prototype;return t.mount=function n(t,e){var i=this,o=this.state,r=this.Components;B(o.is([1,5]),"Already mounted!"),o.set(1),this._Components=r,this._Transition=e||this._Transition||(this.is(n7)?tr:tu),this._Extensions=t||this._Extensions;var u=w({},ti,this._Extensions,{Transition:this._Transition});return x(u,function(n,t){var e=n(i,r,i._options);r[t]=e,e.setup&&e.setup()}),x(r,function(n){n.mount&&n.mount()}),this.emit(ne),g(this.root,"is-initialized"),o.set(3),this.emit(ni),this},t.sync=function n(t){return this.splides.push({splide:t}),t.splides.push({splide:this,isParent:!0}),this.state.is(3)&&(this._Components.Sync.remount(),t.Components.Sync.remount()),this},t.go=function n(t){return this._Components.Controller.go(t),this},t.on=function n(t,e){return this.event.on(t,e,null,20),this},t.off=function n(t){return this.event.off(t),this},t.emit=function n(t){var e;return(e=this.event).emit.apply(e,[t].concat(v(arguments,1))),this},t.add=function n(t,e){return this._Components.Slides.add(t,e),this},t.remove=function n(t){return this._Components.Slides.remove(t),this},t.is=function n(t){return this._options.type===t},t.refresh=function n(){return this.emit(na),this},t.destroy=function n(t){void 0===t&&(t=!0);var i=this.event,o=this.state;return o.is(1)?i.on(ni,this.destroy.bind(this,t),this):(x(this._Components,function(n){n.destroy&&n.destroy(t)},!0),i.emit(ng),i.destroy(),t&&e(this.splides),o.set(5)),this},_createClass(n,[{key:"options",get:function n(){return this._options},set:function n(t){var e=this._options;E(e,t),this.state.is(1)||this.emit(nf,e)}},{key:"length",get:function n(){return this._Components.Slides.getLength(!0)}},{key:"index",get:function n(){return this._Components.Controller.getIndex()}}]),n}();return ts.defaults={},ts.STATES={CREATED:1,MOUNTED:2,IDLE:3,MOVING:4,DESTROYED:5},ts});
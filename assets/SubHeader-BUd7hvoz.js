import{_ as c,r as u,i,o as p,c as d,b as t,f as v,g as _,k as g,u as l,p as f,j as h}from"./index-De5ssewt.js";function x(s){const e=s.filter(n=>n.trim()!=="");if(e.length!==0)return e.sort((n,a)=>n.length-a.length),e[0]}const m=s=>s.slice(0,4);console.log(m("123485774"));const b=s=>(f("data-v-0316c9e8"),s=s(),h(),s),w={class:"search-div"},I=b(()=>t("div",{class:"filter-div"},[t("button",{class:"dropdown-button"},"Filtri"),t("div",{class:"dropdown-menu"},[t("span",null," Filtro 1 "),t("span",null," Filtro 2 "),t("span",null," Filtro 3 ")])],-1)),S={__name:"SubHeader",setup(s){const e=u(""),n=i("getUserInput");return(a,o)=>(p(),d("header",null,[t("div",w,[v(t("input",{onKeypress:o[0]||(o[0]=g(r=>l(n)(e.value),["enter"])),"onUpdate:modelValue":o[1]||(o[1]=r=>e.value=r),class:"search-bar",type:"text"},null,544),[[_,e.value]]),t("button",{onClick:o[2]||(o[2]=r=>l(n)(e.value))},"Cerca")]),I]))}},F=c(S,[["__scopeId","data-v-0316c9e8"]]);export{F as S,x as c,m as g};
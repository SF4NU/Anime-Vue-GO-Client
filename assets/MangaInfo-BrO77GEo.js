import{r as d,_ as r}from"./ratingConverter-CZb7_kzj.js";import{_,l as u,c as p,b as t,u as s,t as i,m as l,q as h,o as v,p as g,j as b,i as m}from"./index-De5ssewt.js";const o=e=>(g("data-v-5a719cab"),e=e(),b(),e),f={class:"anime-main-card"},I={class:"top-row-div"},j={class:"poster-div"},T=["src"],M={class:"anime-info-top-right-div"},x={class:"title-div"},C={class:"sub-titles"},D=o(()=>t("span",null,"Titoli alternativi: ",-1)),k={class:"episodes-div"},w={class:"dates-div"},N=o(()=>t("span",null,"In uscita dal: ",-1)),R=o(()=>t("span",null,"Fino al: ",-1)),S={class:"age-rating-div"},y=o(()=>t("span",null,"Guida alle Classificazioni per Età: ",-1)),B={class:"rating-div"},E=o(()=>t("img",{height:"25px",src:r,alt:"anime-rating-star"},null,-1)),V={class:"rating-rank"},z=o(()=>t("span",null," Top: ",-1)),q={class:"popularity-div"},A=o(()=>t("span",null," Popolarità: ",-1)),F=o(()=>t("span",null," utenti",-1)),G={class:"bottom-synopsis-div"},O=o(()=>t("span",null,"Descrizione: ",-1)),P={__name:"MangaInfo",setup(e){const a=m("dataManga"),n=h().params.id,c=()=>{window.scrollTo({top:0,behavior:"smooth"})};return u(()=>{c(),console.log("MOUNTED")}),(H,J)=>(v(),p("section",null,[t("div",f,[t("div",I,[t("div",j,[t("img",{src:s(a)[s(n)].attributes.posterImage.small,alt:""},null,8,T)]),t("div",M,[t("div",x,[t("span",null,i(s(a)[s(n)].attributes.canonicalTitle),1)]),t("div",C,[D,t("span",null,i(s(a)[s(n)].attributes.titles.en_jp?s(a)[s(n)].attributes.titles.en_jp:""),1),l("     "),t("span",null,i(s(a)[s(n)].attributes.titles.ja_jp?s(a)[s(n)].attributes.titles.ja_jp:""),1)]),t("div",k,[t("span",null,"Capitoli: "+i(s(a)[s(n)].attributes.chapterCount),1),l("      "),t("span",null,"Volumi: "+i(s(a)[s(n)].attributes.volumeCount),1)]),t("div",w,[N,t("span",null,i(s(a)[s(n)].attributes.startDate),1),l("      "),R,t("span",null,i(s(a)[s(n)].attributes.endDate!==null?s(a)[s(n)].attributes.endDate:"N/A"),1)]),t("div",S,[y,t("span",null,i(s(a)[s(n)].attributes.ageRating),1)]),t("div",B,[E,t("span",null,i(s(d)(s(a)[s(n)].attributes.averageRating)),1)]),t("div",V,[z,t("span",null,i(s(a)[s(n)].attributes.ratingRank),1)]),t("div",q,[A,t("span",null,i(s(a)[s(n)].attributes.userCount),1),F])])]),t("div",G,[O,t("p",null,i(s(a)[s(n)].attributes.description),1)])])]))}},Q=_(P,[["__scopeId","data-v-5a719cab"]]);export{Q as default};
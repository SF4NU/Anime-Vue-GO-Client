import{r as m,_ as v}from"./ratingConverter-CZb7_kzj.js";import{_ as g}from"./youtube-svgrepo-com-DeR6TUWh.js";import{S as f,g as w,c as y}from"./SubHeader-CRH6a-o7.js";import{_ as k,c as a,a as d,b as t,F as l,r as C,u as e,o,d as I,w as A,t as n,i as _,e as S,p as T,f as V}from"./index-Ch48BUbF.js";const u=i=>(T("data-v-8b242274"),i=i(),V(),i),x={class:"main-anime-card-div"},L={key:0,class:"poster-div"},N=["src","alt","onClick"],$={key:1,class:"loader"},B={class:"card-details"},D={class:"title"},F={class:"year-ep-count"},E={class:"episode-count"},j={class:"year"},H={class:"rating-div"},O=u(()=>t("img",{height:"25px",src:v,alt:""},null,-1)),R=["data-tooltip"],Y=["href"],q=u(()=>t("img",{class:"watch-img",src:g,alt:"watch-youtube-image"},null,-1)),z=[q],G={__name:"AnimeCards",setup(i){const p=_("data"),c=_("isLoading");return(h,J)=>{const b=S("router-link");return o(),a(l,null,[d(f),t("div",x,[(o(!0),a(l,null,C(e(p),(s,r)=>(o(),a("div",{key:r,class:"main-anime-display"},[e(c)?e(c)?(o(),a("div",$)):I("",!0):(o(),a("div",L,[t("img",{class:"anime-poster",src:s.attributes.posterImage.tiny,alt:s.attributes.canonicalTitle+" poster",onClick:K=>h.$emit("focusOnAnime",r)},null,8,N)])),t("div",B,[t("div",null,[d(b,{to:{name:"AnimeInfo",params:{id:r}}},{default:A(()=>[t("span",D,n(s.attributes.canonicalTitle.length>20&&s.attributes.abbreviatedTitles.length>0?e(y)(s.attributes.abbreviatedTitles):s.attributes.canonicalTitle),1)]),_:2},1032,["to"])]),t("div",F,[t("span",E,n(s.attributes.episodeCount>1?"Episodi: "+s.attributes.episodeCount:s.attributes.episodeCount===null?"N/A":"Film"),1),t("span",j,n(s.attributes.endDate?"Anno: "+e(w)(s.attributes.endDate):"In corso"),1)])]),t("div",H,[O,t("span",null,n(e(m)(s.attributes.averageRating)),1)]),t("div",{class:"trailer-link-div","data-tooltip":`https://www.youtube.com/watch?v=${s.attributes.youtubeVideoId}`},[t("span",null,[t("a",{href:`https://www.youtube.com/watch?v=${s.attributes.youtubeVideoId}`,target:"_blank",class:"watch-link"},z,8,Y)])],8,R)]))),128))])],64)}}},W=k(G,[["__scopeId","data-v-8b242274"]]);export{W as default};

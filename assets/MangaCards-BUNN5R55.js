import{r as h,_ as b}from"./ratingConverter-CZb7_kzj.js";import{S as v,c as m}from"./SubHeader-3bM7kA37.js";import{_ as f,c as e,a as l,b as s,F as d,r as g,u as a,o,d as C,w as k,t as r,i as _,e as y,p as I,f as M}from"./index-CLWZhDPB.js";const S=i=>(I("data-v-eb055f1b"),i=i(),M(),i),x={class:"main-anime-card-div"},T={key:0,class:"poster-div"},L=["src","alt"],N={key:1,class:"loader"},w={class:"card-details"},B={class:"title"},F={class:"year-ep-count"},V={class:"episode-count"},j={class:"year"},A={class:"rating-div"},D=S(()=>s("img",{height:"25px",src:b,alt:""},null,-1)),E={__name:"MangaCards",setup(i){const p=_("dataManga"),n=_("isLoading");return(H,R)=>{const u=y("router-link");return o(),e(d,null,[l(v),s("div",x,[(o(!0),e(d,null,g(a(p),(t,c)=>(o(),e("div",{key:c,class:"main-anime-display"},[a(n)?a(n)?(o(),e("div",N)):C("",!0):(o(),e("div",T,[s("img",{class:"anime-poster",src:t.attributes.posterImage.tiny,alt:t.attributes.canonicalTitle+" poster"},null,8,L)])),s("div",w,[s("div",null,[l(u,{to:{name:"MangaInfo",params:{id:c}},class:"router-link"},{default:k(()=>[s("span",B,r(t.attributes.canonicalTitle.length>20&&t.attributes.abbreviatedTitles.length>0?a(m)(t.attributes.abbreviatedTitles):t.attributes.canonicalTitle),1)]),_:2},1032,["to"])]),s("div",F,[s("span",V,r(t.attributes.chapterCount>1?"Capitoli: "+t.attributes.chapterCount:t.attributes.chapterCount===null?"N/A":"Manga"),1),s("span",j,r(t.attributes.status==="finished"?"Finito":"In corso"),1)])]),s("div",A,[D,s("span",null,r(a(h)(t.attributes.averageRating)),1)])]))),128))])],64)}}},J=f(E,[["__scopeId","data-v-eb055f1b"]]);export{J as default};

import{_ as m,r as s,z as f,c,b as t,e as i,t as n,i as x,q as w,x as B,o as l,p as I,j as k}from"./index-De5ssewt.js";const y="data:image/svg+xml,%3c?xml%20version='1.0'%20encoding='utf-8'?%3e%3csvg%20xmlns='http://www.w3.org/2000/svg'%20width='401px'%20height='401px'%20enable-background='new%20312.809%200%20401%20401'%20version='1.1'%20viewBox='312.809%200%20401%20401'%3e%3cg%20transform='matrix(1.223%200%200%201.223%20-467.5%20-843.44)'%3e%3crect%20x='601.45'%20y='653.07'%20width='401'%20height='401'%20fill='%23E4E6E7'/%3e%3cpath%20d='m802.38%20908.08c-84.515%200-153.52%2048.185-157.38%20108.62h314.79c-3.87-60.44-72.9-108.62-157.41-108.62z'%20fill='%23AEB4B7'/%3e%3cpath%20d='m881.37%20818.86c0%2046.746-35.106%2084.641-78.41%2084.641s-78.41-37.895-78.41-84.641%2035.106-84.641%2078.41-84.641c43.31%200%2078.41%2037.9%2078.41%2084.64z'%20fill='%23AEB4B7'/%3e%3c/g%3e%3c/svg%3e",u=a=>(I("data-v-33427818"),a=a(),k(),a),E={class:"main-profile-div"},S={key:0,class:"profile-picture-div"},b=u(()=>t("img",{height:"200px",src:y,alt:"profile-picture"},null,-1)),D=[b],z={key:1,class:"profile-details-div"},A=u(()=>t("br",null,null,-1)),C={__name:"Profile",setup(a){const d=x("changeLoggedStatus"),p=w(),h=s(p.params.userId),o=s(null),r=s(!1),_=s(0),g=s(0),v=async()=>{try{const e=await B.get(`http://localhost:3000/users/${h.value}`);e.status>=200&&e.status<=209&&(o.value=e.data,console.log(e.data))}catch(e){console.error(e)}};return f(()=>{r.value||(r.value=!0,v(),console.log("hello"),d())}),(e,j)=>(l(),c("section",null,[t("div",E,[o.value?(l(),c("div",S,D)):i("",!0),o.value?(l(),c("div",z,[t("h1",null,n(o.value.Username),1),t("span",null,"Anime count: "+n(_.value),1),A,t("span",null,"Manga count: "+n(g.value),1)])):i("",!0)])]))}},M=m(C,[["__scopeId","data-v-33427818"]]);export{M as default};
import{_,r,c as h,b as e,s as w,f as c,g as v,n as b,x as y,o as x,p as I,j as P}from"./index-Dx0eIbeR.js";const u=n=>(I("data-v-1a3a5a1a"),n=n(),P(),n),U={class:"main-auth-div"},g={class:"registration-div"},V=u(()=>e("h1",null,"Iscrizione",-1)),k=u(()=>e("label",{for:"email"},"Email",-1)),M=u(()=>e("label",{for:"username"},"Username",-1)),S=u(()=>e("label",{for:"password"},"Password",-1)),B=u(()=>e("label",{for:"confirmPassword"},"Conferma Password",-1)),C={__name:"Auth",setup(n){const i=r(""),d=r(""),o=r(""),t=r(""),p=r(!0),m=()=>o.value!==t.value&&t.value!==""&&o.value!==""?(p.value=!1,t.value="",setTimeout(()=>{p.value=!0},1e3),!1):(p.value=!0,!0),f=async()=>{try{if(m()){const l=await y.post("http://localhost:3000/register",{email:i.value,username:d.value,password:o.value});console.log(l),l.request.status>=200&&l.request.status<=209&&(i.value="",d.value="",o.value="",t.value="")}}catch(l){console.error(l)}};return(l,s)=>(x(),h("section",null,[e("div",U,[e("div",g,[V,e("form",{onSubmit:s[4]||(s[4]=w(()=>{},["prevent"]))},[e("div",null,[k,c(e("input",{"onUpdate:modelValue":s[0]||(s[0]=a=>i.value=a),type:"email",id:"email",placeholder:"esempio@yourmail.com"},null,512),[[v,i.value]])]),e("div",null,[M,c(e("input",{"onUpdate:modelValue":s[1]||(s[1]=a=>d.value=a),type:"text",id:"username",placeholder:"utente123"},null,512),[[v,d.value]])]),e("div",null,[S,c(e("input",{"onUpdate:modelValue":s[2]||(s[2]=a=>o.value=a),type:"password",id:"password",placeholder:"••••••••"},null,512),[[v,o.value]])]),e("div",null,[B,c(e("input",{"onUpdate:modelValue":s[3]||(s[3]=a=>t.value=a),type:"password",id:"confirmPassword",class:b(p.value?"":"no-match"),placeholder:"••••••••"},null,2),[[v,t.value]])]),e("button",{type:"submit",onClick:f},"Registrati")],32)])])]))}},z=_(C,[["__scopeId","data-v-1a3a5a1a"]]);export{z as default};

import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    name: "Anime",
    component: () => import(".././views/AnimeCards/AnimeCards.vue"),
  },
  {
    path: "/manga",
    name: "Manga",
    component: () => import(".././views/MangaCards/MangaCards.vue"),
  },
  {
    path: "/anime/:id",
    name: "AnimeInfo",
    component: () => import(".././views/AnimeInfo/AnimeInfo.vue"),
  },
  {
    path: "/manga/:id",
    name: "MangaInfo",
    component: () => import(".././views/MangaInfo/MangaInfo.vue"),
  },
  {
    path: "/register",
    name: "Register",
    component: () => import(".././views/Auth/Auth.vue"),
  },
  {
    path: "/login",
    name: "Login",
    component: () => import(".././views/Auth/Login.vue"),
  },
  {
    path: "/profile/:userId",
    name: "Profile",
    component: () => import(".././views/UserProfile/Profile.vue"),
  },
];

const router = createRouter({
  history: createWebHistory("/Anime-Vue-GO-Client/"),
  routes,
});

export default router;

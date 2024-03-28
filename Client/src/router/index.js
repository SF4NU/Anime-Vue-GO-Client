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
];

const router = createRouter({
  history: createWebHistory("/Anime-Vue-GO-Client/"),
  routes,
});

export default router;

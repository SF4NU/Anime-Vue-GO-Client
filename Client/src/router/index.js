import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    name: "Anime",
    component: () => import(".././views/AnimeCards/AnimeCards.vue"),
    props: true,
  },
  {
    path: "/manga",
    name: "Manga",
    component: () => import(".././views/MangaCards/MangaCards.vue"),
    props: true,
  },
];

const router = createRouter({
  history: createWebHistory("/Anime-Vue-GO-Client/"),
  routes,
});

export default router;

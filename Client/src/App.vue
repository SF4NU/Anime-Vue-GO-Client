<template>
  <section class="main-section">
    <Header @backToAnime="backToAnime" @toManga="toManga" />
    <SubHeader
      @getUserInput="getUserInput"
      v-if="!toggleSubHeader"
      :toggleMangaCards="toggleMangaCards" />
    <AnimeCard
      :data="data"
      :isLoading="isLoading"
      @focusOnAnime="focusOnAnime"
      v-if="!toggleAnimeCards" />
    <AnimeInfo
      :data="data"
      :passIndex="passIndex"
      v-else-if="toggleAnimeFocus"
      @backToAnime="backToAnime" />
    <MangaCards
      :dataManga="dataManga"
      :isLoading="isLoading"
      v-else-if="toggleMangaCards"
      @focusOnManga="focusOnManga" />
    <MangaInfo
      :passIndexManga
      :dataManga="dataManga"
      @backToManga="backToManga"
      v-else-if="toggleMangaFocus" />
    <!-- <Auth /> -->
  </section>
</template>

<script setup>
import { onMounted, ref, watch } from "vue";
import AnimeCard from "./components/AnimeCards/AnimeCards.vue";
import Header from "./components/Header/Header.vue";
import SubHeader from "./components/SubHeader/SubHeader.vue";
import AnimeInfo from "./components/AnimeInfo/AnimeInfo.vue";
import MangaCards from "./components/MangaCards/MangaCards.vue";
import MangaInfo from "./components/MangaInfo/MangaInfo.vue";
import Auth from "./components/Auth/Auth.vue";
import axios from "axios";

const data = ref(null);
const searchedAnime = ref("");
const page = ref(0);
const isLoading = ref(false);
const toggleSubHeader = ref(false);
const toggleAnimeCards = ref(false);
const toggleAnimeFocus = ref(false);
const toggleMangaCards = ref(false);
const toggleMangaFocus = ref(false);
const passIndex = ref(null);
const passIndexManga = ref(null);
const dataManga = ref(null);

const focusOnAnime = (i) => {
  toggleSubHeader.value = true;
  toggleAnimeCards.value = true;
  toggleAnimeFocus.value = true;
  toggleMangaCards.value = true;
  passIndex.value = i;
};
const focusOnManga = (i) => {
  toggleSubHeader.value = true;
  toggleAnimeCards.value = true;
  toggleMangaCards.value = false;
  toggleMangaFocus.value = true;
  passIndexManga.value = i;
};
const backToAnime = () => {
  toggleSubHeader.value = false;
  toggleAnimeCards.value = false;
  toggleAnimeFocus.value = false;
  toggleMangaCards.value = false;
};
const backToManga = () => {
  toggleSubHeader.value = false;
  toggleAnimeCards.value = false;
  toggleMangaCards.value = false;
};
const toManga = () => {
  toggleAnimeCards.value = true;
  toggleMangaCards.value = true;
};

const getUserInput = (input) => {
  searchedAnime.value = input;
  FetchAnime();
  FetchManga();
};

const FetchAnime = async () => {
  try {
    isLoading.value = true;
    const res = await axios.get(
      `https://kitsu.io/api/edge/anime?filter[text]=${searchedAnime.value}&page[limit]=15`
    );
    data.value = res.data.data;
    console.log(data.value);
  } catch (error) {
    console.error(error);
  }
};
const FetchManga = async () => {
  try {
    isLoading.value = true;
    if (searchedAnime.value !== "") {
      const res = await axios.get(
        `https://kitsu.io/api/edge/manga?filter[text]=${searchedAnime.value}&page[limit]=15`
      );
      dataManga.value = res.data.data;
    } else {
      const res = await axios.get(
        `https://kitsu.io/api/edge/manga?page[limit]=13&page[offset]=30`
      );
      dataManga.value = res.data.data;
    }

    console.log(dataManga.value);
  } catch (error) {
    console.error(error);
  }
};

watch(
  () => data.value,
  () => {
    isLoading.value = false;
  }
);

onMounted(() => {
  FetchAnime();
  FetchManga();
  console.log(searchedAnime.value);
});
</script>

<style>
:root {
  --green: #43aa8b;
  --dark-blue: #193d5d;
  --yellow: #ffea00;
  --red: #f94144;
  --orange: #ff5900;
}
body {
  margin: 0;
  font-family: Rubik;
  background: linear-gradient(to bottom right, var(--green), var(--dark-blue));
  min-height: 100dvh;
}
</style>

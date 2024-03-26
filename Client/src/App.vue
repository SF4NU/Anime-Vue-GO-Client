<template>
  <section class="main-section">
    <Header />
    <SubHeader @getUserInput="getUserInput" v-if="!toggleSubHeader" />
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
  </section>
</template>

<script setup>
import { onMounted, ref, watch } from "vue";
import AnimeCard from "./components/AnimeCards/AnimeCards.vue";
import Header from "./components/Header/Header.vue";
import SubHeader from "./components/SubHeader/SubHeader.vue";
import AnimeInfo from "./components/AnimeInfo/AnimeInfo.vue";
import axios from "axios";

const data = ref(null);
const searchedAnime = ref("");
const page = ref(0);
const isLoading = ref(false);
const toggleSubHeader = ref(false);
const toggleAnimeCards = ref(false);
const toggleAnimeFocus = ref(false);
const passIndex = ref(null);

const focusOnAnime = (i) => {
  toggleSubHeader.value = true;
  toggleAnimeCards.value = true;
  toggleAnimeFocus.value = true;
  passIndex.value = i;
};
const backToAnime = () => {
  toggleSubHeader.value = false;
  toggleAnimeCards.value = false;
  toggleAnimeFocus.value = false;
};

const getUserInput = (input) => {
  searchedAnime.value = input;
  FetchAnime();
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

watch(
  () => data.value,
  () => {
    isLoading.value = false;
  }
);

onMounted(() => {
  FetchAnime();
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

<template>
  <section class="main-section">
    <Header />
    <SubHeader />
    <AnimeCard :data="data" />
  </section>
</template>

<script setup>
import { onMounted, ref } from "vue";
import AnimeCard from "./components/AnimeCards/AnimeCards.vue";
import Header from "./components/Header/Header.vue";
import SubHeader from "./components/SubHeader/SubHeader.vue";
import axios from "axios";

const data = ref(null);

onMounted(() => {
  const FetchAnime = async () => {
    try {
      const res = await axios.get(
        "https://kitsu.io/api/edge/anime?filter[text]=naruto"
      );
      data.value = res.data.data;
      console.log(data.value);
    } catch (error) {
      console.error(error);
    }
  };
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
}
</style>

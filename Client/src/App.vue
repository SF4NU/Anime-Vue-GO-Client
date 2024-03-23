<template>
  <section>
    <AnimeCard :data="data" />
  </section>
</template>

<script setup>
import { onMounted, ref } from "vue";
import AnimeCard from "./components/AnimeCards/AnimeCards.vue";
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

<style scoped>
.main-div {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(110px, 1fr));
  column-gap: 10px;
}
</style>

<template>
  <section>
    <h1>My list</h1>
    <div class="anime-container">
      <div v-for="(anime, i) in animeListData" :key="i" class="anime-card">
        <h2>{{ anime.title }}</h2>
        <div>
          <h3>{{ anime.episodes }}</h3>
        </div>
        <div>{{ anime.rating }}</div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { onBeforeMount, ref } from "vue";
import { useRoute } from "vue-router";
import axios from "axios";

const router = useRoute();
const userId = ref(router.params.userId);
const checkIfDataFetched = ref(false);
const animeListData = ref(null);

const getAnimeList = async () => {
  try {
    const res = await axios.get(
      `http://localhost:3000/users/${userId.value}/anime`
    );
    if (res.status >= 200 && res.status <= 209) {
      animeListData.value = res.data;
    } else {
    }
  } catch (error) {
    console.error(error);
  }
};

onBeforeMount(() => {
  if (!checkIfDataFetched.value) {
    checkIfDataFetched.value = true;
    getAnimeList();
  }
});
</script>

<style scoped></style>

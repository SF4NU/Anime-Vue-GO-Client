<template>
  <Header />
  <router-view />
</template>

<script setup>
import { onMounted, ref, watch, provide } from "vue";
import Header from "./components/Header/Header.vue";
import axios from "axios";

const data = ref(null);
provide("data", data);

const searchedAnime = ref("");
const page = ref(0);
const isLoading = ref(false);
provide("isLoading", isLoading);
const dataManga = ref(null);
provide("dataManga", dataManga);
const userId = ref(null);
provide("userId", userId);
const getUserId = (id) => {
  userId.value = id;
};
provide("getUserId", getUserId);
const getUserInput = (input) => {
  searchedAnime.value = input;
  FetchAnime();
  FetchManga();
};
provide("getUserInput", getUserInput);

const isLoggedIn = ref(false);

const changeLoggedStatus = () => {
  isLoggedIn.value = true;
};

provide("changeLoggedStatus", changeLoggedStatus);
provide("isLoggedIn", isLoggedIn);

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

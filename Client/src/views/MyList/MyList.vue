<template>
  <div></div>
</template>

<script setup>
import { onBeforeMount, ref } from "vue";
import { useRoute } from "vue-router";
import axios from "axios";

const router = useRoute();
const userId = ref(router.params.userId);
const checkIfDataFetched = ref(false);

const getAnimeList = async () => {
  try {
    const res = await axios.get(
      `http://localhost:3000/users/${userId.value}/anime`
    );
    if (res.status >= 200 && res.status <= 209) {
      console.log(res.data);
    } else {
      console.log(res);
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

<template>
  <div v-if="userId">{{ userId }}</div>
  <div v-if="userData">{{ userData.Username }}</div>
  <div v-if="userData">{{ userData.Email }}</div>
</template>

<script setup>
import { ref, inject, onBeforeMount, provide } from "vue";
import { useRoute } from "vue-router";
import axios from "axios";
const changeLoggedStatus = inject("changeLoggedStatus");

const route = useRoute();
const userId = ref(route.params.userId);
const userData = ref(null);
const checkIfDataFetched = ref(false);

const getUserData = async () => {
  try {
    const res = await axios.get(`http://localhost:3000/users/${userId.value}`);
    if (res.status >= 200 && res.status <= 209) {
      userData.value = res.data;
      console.log(res.data);
    }
  } catch (error) {
    console.error(error);
  }
};

onBeforeMount(() => {
  if (!checkIfDataFetched.value) {
    checkIfDataFetched.value = true;
    getUserData();
    console.log("hello");
    changeLoggedStatus();
  }
});
</script>

<style scoped></style>

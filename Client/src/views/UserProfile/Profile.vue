<template>
  <section>
    <div class="main-profile-div">
      <div v-if="userData" class="profile-picture-div">
        <img
          height="200px"
          src="../../assets/profile-picture.svg"
          alt="profile-picture" />
      </div>
      <div v-if="userData" class="profile-details-div">
        <h1>{{ userData.Username }}</h1>
        <span>Anime count: {{ animeCount }}</span
        ><br />
        <span>Manga count: {{ mangaCount }}</span>
      </div>
    </div>
  </section>
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
const animeCount = ref(0);
const mangaCount = ref(0);

const getUserData = async () => {
  try {
    const res = await axios.get(`http://localhost:3000/users/${userId.value}`, {
      withCredentials: true,
    });
    if (res.status >= 200 && res.status <= 209) {
      userData.value = res.data;
    }
  } catch (error) {
    console.error(error);
  }
};

onBeforeMount(() => {
  if (!checkIfDataFetched.value) {
    checkIfDataFetched.value = true;
    getUserData();
    changeLoggedStatus();
  }
});
</script>

<style scoped>
section {
  margin-top: -100px;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
.main-profile-div {
  display: flex;
  column-gap: 50px;
  background-color: var(--green);
  padding: 20px;
  width: clamp(300px, 70%, 800px);
}
.main-profile-div h1 {
  font-size: 2rem;
  margin-top: 10px;
  color: var(--dark-blue);
}
.main-profile-div span {
  font-size: 1.5rem;
  color: var(--dark-blue);
}
</style>

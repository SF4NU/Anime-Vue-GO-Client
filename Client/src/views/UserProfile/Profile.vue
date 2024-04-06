<template>
  <section>
    <div class="main-profile-div">
      <div v-if="userData" class="profile-picture-div">
        <img src="../../assets/profile-picture.svg" alt="profile-picture" />
      </div>
      <div v-if="userData" class="profile-details-div">
        <h1>{{ userData.Username }}</h1>
        <span>Anime count: {{ animeCount }}</span
        ><br />
        <span>Manga count: {{ mangaCount }}</span>
      </div>
      <img
        @click="logout"
        class="logout"
        height="30px"
        src="@/assets/logout.svg"
        alt="" />
    </div>
  </section>
</template>

<script setup>
import { ref, inject, onBeforeMount, provide } from "vue";
import { useRoute } from "vue-router";
import axios from "axios";
import router from "@/router";
import Cookies from "js-cookie";

const changeLoggedStatus = inject("changeLoggedStatus");
const route = useRoute();
const userId = ref(route.params.userId);
const userData = ref(null);
const checkIfDataFetched = ref(false);
const animeCount = ref(0);
const mangaCount = ref(0);

const getUserData = async () => {
  try {
    const res = await axios.get(
      `https://anime-vue-go-client-production.up.railway.app/users/${userId.value}`,
      {
        withCredentials: true,
      }
    );
    if (res.status >= 200 && res.status <= 209) {
      userData.value = res.data;
    }
  } catch (error) {
    console.error(error);
  }
};

const logout = async () => {
  try {
    const res = await axios.post(
      "https://anime-vue-go-client-production.up.railway.app/logout",
      {
        withCredentials: true,
      }
    );
    if (res.status >= 200 && res.status <= 209) {
      console.log(document.cookie);
      Cookies.remove("jwt");
      await changeLoggedStatus(false);
      await router.push("/login");
    }
  } catch (error) {
    console.error(error);
  }
};

onBeforeMount(() => {
  if (!checkIfDataFetched.value) {
    checkIfDataFetched.value = true;
    getUserData();
    changeLoggedStatus(true);
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
  border-radius: 25px;
  position: relative;
}
.main-profile-div h1 {
  font-size: 2rem;
  margin-top: 20px;
  color: var(--dark-blue);
  word-break: break-all;
}
.main-profile-div span {
  font-size: 1.5rem;
  color: var(--dark-blue);
}
.profile-picture-div img {
  height: 200px;
}
@media (max-width: 768px) {
  .main-profile-div {
    flex-direction: column;
    row-gap: 20px;
    align-items: center;
    justify-content: center;
  }
  .profile-picture-div img {
    height: 150px;
  }
  .profile-picture-div {
    display: flex;
    justify-content: center;
    margin-top: 25px;
  }
}
.logout {
  position: absolute;
  right: 12px;
  top: 12px;
  cursor: pointer;
}
</style>

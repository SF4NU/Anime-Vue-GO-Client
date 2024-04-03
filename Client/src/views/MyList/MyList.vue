<template>
  <section>
    <h1>My list</h1>
    <div class="anime-container">
      <div v-for="(anime, i) in animeListData" :key="i" class="anime-card">
        <h2>{{ anime.title }}</h2>
        <div>
          <span> Episodi guardati: {{ anime.episodes }} </span>
        </div>
        <div>
          <span> Voto assegnato: {{ anime.rating }} </span>
        </div>
        <div class="completed-div">
          <span>Completato: </span>
          <img height="20px" :src="anime.finished ? check : cross" alt="" />
        </div>
        <div class="dates-div">
          <span> Inizio: {{ anime.starting_date }} </span><br />
          <span v-if="!anime.plan_to_watch && anime.finished">
            Fine: {{ anime.ending_date }}
          </span>
        </div>
        <div
          class="plan-to-watch-div"
          v-if="!anime.finished && anime.plan_to_watch">
          <span> Da vedere ?</span>&nbsp;
          <img
            height="20px"
            :src="anime.plan_to_watch ? check : cross"
            alt="" />
        </div>
        <div class="comment-div">
          <button class="comment-button" :data-comment="anime.comment">
            Nota &nbsp; <img height="25px" src="@/assets/cursor.svg" alt="" />
          </button>
          <p v-if="showComment" class="show-comment">{{ anime.comment }}</p>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { onBeforeMount, ref } from "vue";
import { useRoute } from "vue-router";
import axios from "axios";
import check from "@/assets/check.svg";
import cross from "@/assets/cross-circle-svgrepo-com.svg";

const router = useRoute();
const userId = ref(router.params.userId);
const checkIfDataFetched = ref(false);
const animeListData = ref(null);
const showComment = ref(false);

const getAnimeList = async () => {
  try {
    const res = await axios.get(
      `http://localhost:3000/users/${userId.value}/anime`
    );
    if (res.status >= 200 && res.status <= 209) {
      animeListData.value = res.data;
      console.log(res.data);
    } else {
    }
  } catch (error) {
    console.error(error);
  }
};

const updateAnimeCard = () => {
  try {
    const res = axios.put(
      `http://localhost:3000/update/anime/${animeListData.value.ID}`
    );
  } catch (error) {
    console.error(error);
  }
};

const deleteAnimeCard = () => {
  try {
    const res = axios.delete(
      `http://localhost:3000/delete/anime/${animeListData.value.ID}`
    );
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

<style scoped>
h1 {
  text-align: center;
  margin-top: 50px;
  margin-bottom: 50px;
  font-size: 2rem;
  color: var(--green);
  background: rgba(25, 61, 93, 0.9);
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
  backdrop-filter: blur(6.5px);
  -webkit-backdrop-filter: blur(6.5px);
  border-radius: 10px;
  width: fit-content;
  margin-right: auto;
  margin-left: auto;
  padding: 10px 30px;
  border-radius: 35px;
  user-select: none;
}
.anime-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 220px;
  height: fit-content;
  border: 1px solid black;
  margin: 10px;
  padding: 20px;
  background: linear-gradient(145deg, #3c997d, #48b695);
  box-shadow: 9px 9px 18px #41a587, -9px -9px 18px #45af8f;
  border: 1px solid rgba(255, 255, 255, 0.18);
  border-radius: 25px;
  text-shadow: 2px 2px 2px rgba(0, 0, 0, 0.15);
  transition: transform 0.25s ease, background 0.8s ease, box-shadow 0.8s ease;
}
.anime-card:hover {
  transform: scale(1.02);
  animation: background-changer 3s linear infinite alternate;
}
@keyframes background-changer {
  0% {
    box-shadow: 9px 9px 18px #41a587, -9px -9px 18px #45af8f;
    border: 1px solid rgba(255, 255, 255, 0.18);
  }
  33% {
    box-shadow: 10px -10px 20px #3fa083, -10px 10px 20px #47b493;
  }
  66% {
    box-shadow: -10px -10px 20px #3fa083, 10px 10px 20px #47b493;
  }
  100% {
    box-shadow: -10px 10px 20px #3fa083, 10px -10px 20px #47b493;
  }
}
.anime-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  gap: 50px;
  margin-right: 100px;
  margin-left: 100px;
}
.anime-card h2 {
  font-size: 1.5rem;
  margin: 10px;
  text-align: center;
}
.completed-div {
  display: flex;
  align-items: center;
}
.dates-div {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 10px;
}
.plan-to-watch-div {
  display: flex;
  align-items: center;
}
.comment-div {
  margin-top: 10px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.comment-div p {
  text-align: center;
  word-break: break-all;
}
.comment-div button {
  background-color: var(--dark-blue);
  color: var(--yellow);
  border: none;
  border-radius: 15px;
  padding: 5px 10px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.15);
  position: relative;
}
.comment-button::after {
  content: attr(data-comment);
  height: fit-content;
  width: 150px;
  position: absolute;
  background: rgba(25, 61, 93, 0.75);
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
  backdrop-filter: blur(8.5px);
  -webkit-backdrop-filter: blur(8.5px);
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.18);
  word-break: break-all;
  padding: 20px;
  font-weight: 600;
  border-radius: 15px;
  opacity: 0;
  transition: opacity 0.25s ease, scale 0.5s ease;
  scale: 0;
}
.comment-button:hover::after {
  opacity: 1;
  scale: 1;
}
</style>

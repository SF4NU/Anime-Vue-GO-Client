<template>
  <section>
    <h1>My list</h1>
    <div class="anime-container">
      <div v-for="(anime, i) in animeListData" :key="i" class="anime-card">
        <div class="anime-card-wrap" v-if="isEditing !== i">
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
          <div class="edit-div">
            <img
              @click="handleIsEditing(i, anime.anime_id)"
              height="25px"
              src="@/assets/edit.svg"
              alt="" />
          </div>
        </div>

        <div v-if="isEditing === i" class="edit-div-wrapper">
          <div class="edit-div" v-if="!isLoading">
            <img
              @click="handleIsEditing(i)"
              height="25px"
              :src="cross"
              alt="cross-icon" />
          </div>
          <div class="title-div">
            <h2>{{ anime.title }}</h2>
          </div>
          <div class="check-if-plan-to-watch" v-if="!isLoading">
            <div class="checkbox-wrapper-58">
              <label class="switch">
                <input type="checkbox" v-model="plan_to_watch" />
                <span class="slider"></span>
              </label>
            </div>
            <h4>Anime da vedere</h4>
          </div>
          <div class="add-values-div" v-if="!plan_to_watch && !isLoading">
            <div class="episodes-title-div">
              <h4>Episodi guardati</h4>
            </div>
            <div class="add-episodes-div">
              <img
                height="30px"
                src="@/assets/minus-blue.svg"
                @click="
                  decrementEpisodeCount(singleAnimeData.attributes.episodeCount)
                "
                alt="minus-svg" />
              <span>{{ episodes }}</span>
              <img
                height="25px"
                src="@/assets/plus-blue.svg"
                @click="
                  incrementEpisodeCount(singleAnimeData.attributes.episodeCount)
                "
                alt="plus-svg" />
            </div>
          </div>
          <div v-if="!plan_to_watch && !isLoading" class="add-values-div">
            <div class="rating-title-div">
              <h4>Voto</h4>
            </div>
            <div class="rating-div">
              <img
                height="30px"
                src="@/assets/minus-blue.svg"
                @click="decrementUserRating"
                alt="minus-svg" />
              <span>{{ rating }}</span>
              <img
                height="25px"
                src="@/assets/plus-blue.svg"
                @click="incrementUserRating"
                alt="plus-svg" />
            </div>
          </div>
          <div class="edit-comment-div" v-if="!isLoading">
            <button @click="displayTextArea">Nota</button>
            <div class="text-area-div">
              <textarea
                @input="updateLengthCounter"
                v-if="displayText"
                maxlength="200"
                v-model="textAreaLength"
                name="comment"
                id="comment"
                cols="20"
                rows="5"
                placeholder="Questo anime mi è piaciuto/non mi è piaciuto perché...">
              </textarea>
              <span v-if="displayText">{{ lengthCounter }}/200</span>
            </div>
          </div>
          <div class="date-div">
            <h5>Data di inizio (opz.)</h5>
            <input
              type="date"
              v-model="startingDate"
              min="2000-01-01"
              max="2025-12-31" />
            <h5 v-if="finished">Data di fine (opz.)</h5>
            <input
              v-if="finished"
              type="date"
              v-model="endingDate"
              min="2000-01-01"
              max="2025-12-31" />
          </div>
          <div v-if="isLoading" class="loader-wrapper">
            <div class="loader"></div>
          </div>
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
const isEditing = ref(null);
const episodes = ref(0);
const rating = ref(1);
const finished = ref(false);
const starting_date = ref("");
const ending_date = ref("");
const plan_to_watch = ref(false);
const comment = ref("");
const animeID = ref(null);
const isLoading = ref(false);
const displayText = ref(false);
const textAreaLength = ref("");
const lengthCounter = ref(0);

const singleAnimeData = ref(null);

const setCountersToZero = () => {
  episodes.value = 0;
  rating.value = "";
  finished.value = false;
  starting_date.value = "";
  ending_date.value = "";
  plan_to_watch.value = false;
  comment.value = "";
  plan_to_watch.value = false;
};

const displayTextArea = () => {
  displayText.value = !displayText.value;
};

const updateLengthCounter = () => {
  lengthCounter.value = textAreaLength.value.length;
};

const handleIsEditing = async (i, id) => {
  if (isEditing.value === i) {
    isEditing.value = null;
    setCountersToZero();
    return;
  }
  isEditing.value = i;
  setCountersToZero();
  animeID.value = animeListData.value[i].ID;
  episodes.value = animeListData.value[i].episodes;
  rating.value = animeListData.value[i].rating;
  textAreaLength.value = animeListData.value[i].comment;
  lengthCounter.value = animeListData.value[i].comment.length;
  isLoading.value = true;
  await findAnimeByID(id);
  console.log(singleAnimeData.value);
  isLoading.value = false;
};

const checkIfAnimeCompleted = (maxEp) => {
  if (episodes.value > maxEp - 1 || episodes.value === "Completato") {
    finished.value = true;
    if (singleAnimeData.value.attributes.episodeCount === 1) {
      episodes.value = 1;
    }
    return;
  }
  finished.value = false;
};
const incrementEpisodeCount = (maxEp) => {
  if (episodes.value >= maxEp - 1 || episodes.value === "Completato") {
    episodes.value = "Completato";
    checkIfAnimeCompleted(maxEp);
    return;
  }
  checkIfAnimeCompleted(maxEp);
  episodes.value++;
};
const decrementEpisodeCount = (maxEp) => {
  if (episodes.value <= 0) {
    episodes.value = 0;
    checkIfAnimeCompleted(maxEp);
    return;
  }
  if (episodes.value === "Completato") {
    episodes.value = maxEp - 1;
    checkIfAnimeCompleted(maxEp);
    return;
  }
  episodes.value--;
  checkIfAnimeCompleted(maxEp);
};

const getAnimeList = async () => {
  try {
    const res = await axios.get(
      `https://anime-vue-go-client-production.up.railway.app/users/${userId.value}/anime`,
      {
        withCredentials: true,
      }
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
      `https://anime-vue-go-client-production.up.railway.app/update/anime/${animeID.value}`,
      {
        title: title.value,
        episodes: episodes.value,
        rating: rating.value,
        finished: finished.value,
        starting_date: starting_date.value,
        ending_date: ending_date.value,
        plan_to_watch: plan_to_watch.value,
        comment: comment.value,
      }
    );
  } catch (error) {
    console.error(error);
  }
};

const deleteAnimeCard = () => {
  try {
    const res = axios.delete(
      `https://anime-vue-go-client-production.up.railway.app/delete/anime/${animeID.value}`
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

const findAnimeByID = async (id) => {
  try {
    const res = await axios.get(
      `https://kitsu.io/api/edge/anime?filter[id]=${id}&page[limit]=1`
    );
    singleAnimeData.value = res.data.data[0];
  } catch (error) {
    console.error(error);
  }
};
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
  position: relative;
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
.edit-div {
  position: absolute;
  top: 12px;
  right: 15px;
  cursor: pointer;
}
.anime-card-wrap {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.show-comment {
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
}
input {
  margin: 5px;
}
button {
  margin: 5px;
}
.edit-div-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.loader-wrapper {
  height: 200px;
  width: fit-content;
  display: flex;
  justify-content: center;
  align-items: center;
}
.add-values-div {
  display: flex;
  column-gap: 20px;
  align-items: center;
  justify-content: space-between;
  width: 100%;
}
.add-values-div input {
  width: 55px;
  height: 35px;
  border-radius: 15px;
  background-color: var(--green);
  color: var(--dark-blue);
  font-weight: 600;
  text-align: center;
  border: none;
  margin-right: 3px;
  user-select: none;
}
.add-values-div img {
  cursor: pointer;
}
.add-values-div img:hover {
  filter: brightness(0.9);
}
.add-values-div img:active {
  filter: brightness(0.8);
}
.add-values-div input:focus {
  outline: none;
}
.check-if-plan-to-watch {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-top: 10px;
}
.check-if-plan-to-watch h4 {
  margin-left: 10px;
}
.rating-div,
.add-episodes-div {
  display: flex;
  align-items: center;
  justify-content: space-evenly;
  width: 50%;
}
.episodes-title-div,
.rating-title-div {
  width: 50%;
}
.edit-comment-div {
  margin-top: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.edit-comment-div button {
  background-color: var(--dark-blue);
  color: var(--green);
  font-weight: 600;
  padding: 5px 10px;
  border-radius: 15px;
  cursor: pointer;
  margin-bottom: 10px;
  border: none;
}
.edit-comment-div button:hover {
  filter: brightness(0.9);
}
.edit-comment-div button:active {
  filter: brightness(0.8);
}
.edit-comment-div textarea {
  width: 80%;
  height: 100px;
  border-radius: 15px;
  background-color: var(--green);
  color: var(--dark-blue);
  font-weight: 600;
  padding: 10px;
  border: none;
  resize: none;
}
.edit-comment-div textarea:focus {
  outline: none;
}
.text-area-div textarea {
  width: 70%;
  height: 100px;
  border-radius: 15px;
  background-color: var(--dark-blue);
  color: var(--green);
  font-weight: 600;
  padding: 10px;
  border: none;
  resize: none;
}
.text-area-div {
  display: flex;
  position: relative;
  align-items: center;
  justify-content: center;
}
.text-area-div span {
  position: absolute;
  bottom: 5px;
  right: 30px;
  color: var(--yellow);
  font-weight: 400;
  font-size: 0.6rem;
  opacity: 0.8;
}
.loader {
  width: 50px;
  aspect-ratio: 1;
  border-radius: 50%;
  background: radial-gradient(farthest-side, #f03355 95%, #0000) 50% 1px/12px
      12px no-repeat,
    radial-gradient(farthest-side, #0000 calc(100% - 14px), #ccc 0);
  animation: l9 2s infinite linear;
}
@keyframes l9 {
  to {
    transform: rotate(1turn);
  }
}

.checkbox-wrapper-58 input[type="checkbox"] {
  visibility: hidden;
  display: none;
}

.checkbox-wrapper-58 *,
.checkbox-wrapper-58 ::after,
.checkbox-wrapper-58 ::before {
  box-sizing: border-box;
}

.checkbox-wrapper-58 .switch {
  --switch_width: 2em;
  --switch_height: 1em;
  --thumb_color: #e8e8e8;
  --track_color: #e8e8e8;
  --track_active_color: #888;
  --outline_color: #ac3d01;
  font-size: 17px;
  position: relative;
  display: inline-block;
  width: var(--switch_width);
  height: var(--switch_height);
}

.checkbox-wrapper-58 .slider {
  box-sizing: border-box;
  border: 2px solid var(--outline_color);
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: var(--track_color);
  transition: 0.15s;
  border-radius: var(--switch_height);
}

.checkbox-wrapper-58 .slider:before {
  box-sizing: border-box;
  position: absolute;
  content: "";
  height: var(--switch_height);
  width: var(--switch_height);
  border: 2px solid var(--outline_color);
  border-radius: 100%;
  left: -2px;
  bottom: -2px;
  background-color: var(--thumb_color);
  transform: translateY(-0.2em);
  box-shadow: 0 0.2em 0 var(--outline_color);
  transition: 0.15s;
}

.checkbox-wrapper-58 input:checked + .slider {
  background-color: var(--track_active_color);
}

.checkbox-wrapper-58 input:focus-visible + .slider {
  box-shadow: 0 0 0 2px var(--track_active_color);
}

.checkbox-wrapper-58 input:hover + .slider:before {
  transform: translateY(-0.3em);
  box-shadow: 0 0.3em 0 var(--outline_color);
}

.checkbox-wrapper-58 input:checked + .slider:before {
  transform: translateX(calc(var(--switch_width) - var(--switch_height)))
    translateY(-0.2em);
}

.checkbox-wrapper-58 input:hover:checked + .slider:before {
  transform: translateX(calc(var(--switch_width) - var(--switch_height)))
    translateY(-0.3em);
  box-shadow: 0 0.3em 0 var(--outline_color);
}
</style>

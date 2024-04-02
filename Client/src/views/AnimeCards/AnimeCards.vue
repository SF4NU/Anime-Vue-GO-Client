<template>
  <SubHeader />
  <div class="main-anime-card-div">
    <div v-for="(anime, i) in data" :key="i" class="main-anime-display">
      <div v-if="!isLoading && isAdding !== i" class="poster-div">
        <img
          class="anime-poster"
          :src="anime.attributes.posterImage.tiny"
          :alt="anime.attributes.canonicalTitle + ' poster'" />
      </div>
      <div v-else-if="isLoading" class="loader"></div>
      <div class="card-details" v-if="isAdding !== i">
        <div>
          <router-link
            :to="{ name: 'AnimeInfo', params: { id: i } }"
            class="router-link">
            <span class="title">
              {{
                anime.attributes.canonicalTitle.length > 20 &&
                anime.attributes.abbreviatedTitles.length > 0
                  ? compareLengths(anime.attributes.abbreviatedTitles)
                  : cutString(anime.attributes.canonicalTitle)
              }}
            </span>
          </router-link>
        </div>
        <div class="year-ep-count">
          <span class="episode-count">
            {{
              anime.attributes.episodeCount > 1
                ? "Episodi: " + anime.attributes.episodeCount
                : anime.attributes.episodeCount === null
                ? "N/A"
                : "Film"
            }}
          </span>
          <span class="year">
            {{
              anime.attributes.endDate
                ? "Anno: " + getYear(anime.attributes.endDate)
                : "In corso"
            }}
          </span>
        </div>
      </div>
      <div class="rating-div" v-if="isAdding !== i">
        <img
          height="25px"
          src="../../assets/star-circle-svgrepo-com.svg"
          alt="" />
        <span>
          {{ ratingConverter(anime.attributes.averageRating) }}
        </span>
      </div>
      <div
        v-if="isAdding !== i"
        class="trailer-link-div"
        :data-tooltip="`https://www.youtube.com/watch?v=${anime.attributes.youtubeVideoId}`">
        <span
          ><a
            :href="`https://www.youtube.com/watch?v=${anime.attributes.youtubeVideoId}`"
            target="_blank"
            class="watch-link">
            <img
              class="watch-img"
              src="../../assets/youtube-svgrepo-com.svg"
              alt="watch-youtube-image" /> </a
        ></span>
      </div>
      <div
        :class="isAdding !== i ? 'add-to-list' : 'add-to-list-close'"
        data-tooltip="Aggiungi alla tua lista">
        <img
          height="35px"
          :src="isAdding !== i ? plus : close"
          alt="add-to-list"
          @click="isAddingAnime(i)"
          id="add-to-list-button" />
      </div>
      <div class="add-to-list-expanded-div" v-if="isAdding === i">
        <h3>
          Aggiungi: &nbsp;
          {{
            anime.attributes.canonicalTitle.length > 20 &&
            anime.attributes.abbreviatedTitles.length > 0
              ? compareLengths(anime.attributes.abbreviatedTitles)
              : cutString(anime.attributes.canonicalTitle)
          }}
          alla tua lista
        </h3>
        <div class="check-if-plan-to-watch">
          <div class="checkbox-wrapper-58">
            <label class="switch">
              <input type="checkbox" v-model="checkIfPlanToWatch" />
              <span class="slider"></span>
            </label>
          </div>
          <h4>Anime da vedere</h4>
        </div>
        <h4
          v-if="
            !checkIfPlanToWatch &&
            anime.attributes.episodeCount !== null &&
            anime.attributes.episodeCount !== 1
          ">
          Numero di episodi visti
        </h4>
        <div
          v-if="
            !checkIfPlanToWatch &&
            anime.attributes.episodeCount !== null &&
            anime.attributes.episodeCount !== 1
          "
          class="add-episodes-div">
          <img
            height="42px"
            src="@/assets/minus.svg"
            @click="decrementEpisodeCount(anime.attributes.episodeCount)"
            alt="minus-svg" />
          <span>{{ episodeCount }}</span>
          <img
            height="35px"
            src="@/assets/plus.svg"
            @click="incrementEpisodeCount(anime.attributes.episodeCount)"
            alt="plus-svg" />
        </div>
        <h4 v-if="!checkIfPlanToWatch">Voto</h4>
        <div v-if="!checkIfPlanToWatch" class="add-episodes-div">
          <img
            height="42px"
            src="@/assets/minus.svg"
            @click="decrementUserRating"
            alt="minus-svg" />
          <span>{{ userRating }}</span>
          <img
            height="35px"
            src="@/assets/plus.svg"
            @click="incrementUserRating"
            alt="plus-svg" />
        </div>
        <div class="comment-div">
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
          <h5 v-if="animeCompleted">Data di fine (opz.)</h5>
          <input
            v-if="animeCompleted"
            type="date"
            v-model="endingDate"
            min="2000-01-01"
            max="2025-12-31" />
        </div>
        <div class="submit-anime-div">
          <button @click="submitAnime">Aggiungi</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, inject } from "vue";
// const props = defineProps(["data", "isLoading"]);
const data = inject("data");
const isLoading = inject("isLoading");
import { compareLengths } from "../../utils/compareLengths";
import { getYear } from "../../utils/getYear";
import { ratingConverter } from "../../utils/ratingConverter";
import { cutString } from "../../utils/cutString";
import SubHeader from "@/components/SubHeader/SubHeader.vue";
import close from "@/assets/close.svg";
import plus from "@/assets/plus.svg";
import axios from "axios";

const isLoggedIn = inject("isLoggedIn");
const checkIfPlanToWatch = ref(false);
const textAreaLength = ref("");
const lengthCounter = ref(0);
const updateLengthCounter = () => {
  lengthCounter.value = textAreaLength.value.length;
};
const animeCompleted = ref(false);
const isAdding = ref(null);
const isAddingAnime = (i) => {
  if (isAdding.value === i) {
    isAdding.value = null;
    setCountersToZero();
    return;
  }
  isAdding.value = i;
  setCountersToZero();
};

const setCountersToZero = () => {
  episodeCount.value = 0;
  userRating.value = 1;
  textAreaLength.value = "";
  lengthCounter.value = 0;
  displayText.value = false;
  startingDate.value = "";
  endingDate.value = "";
  checkIfPlanToWatch.value = false;
  animeCompleted.value = false;
};

const episodeCount = ref(0);
const checkIfAnimeCompleted = (maxEp) => {
  if (episodeCount.value > maxEp - 1 || episodeCount.value === "Completato") {
    animeCompleted.value = true;
    if (data.value[isAdding.value].attributes.episodeCount === 1) {
      episodeCount.value = 1;
    }
    return;
  }
  animeCompleted.value = false;
};
const incrementEpisodeCount = (maxEp) => {
  if (episodeCount.value >= maxEp - 1 || episodeCount.value === "Completato") {
    episodeCount.value = "Completato";
    checkIfAnimeCompleted(maxEp);
    return;
  }
  checkIfAnimeCompleted(maxEp);
  episodeCount.value++;
};
const decrementEpisodeCount = (maxEp) => {
  if (episodeCount.value <= 0) {
    episodeCount.value = 0;
    checkIfAnimeCompleted(maxEp);
    return;
  }
  if (episodeCount.value === "Completato") {
    episodeCount.value = maxEp - 1;
    checkIfAnimeCompleted(maxEp);
    return;
  }
  episodeCount.value--;
  checkIfAnimeCompleted(maxEp);
};
const userRating = ref(1);

const incrementUserRating = () => {
  if (userRating.value >= 10) {
    userRating.value = 10;
    return;
  }
  userRating.value += 0.5;
};
const decrementUserRating = () => {
  if (userRating.value <= 1) {
    userRating.value = 1;
    return;
  }
  userRating.value -= 0.5;
};
const displayText = ref(false);
const displayTextArea = () => {
  displayText.value = !displayText.value;
};

const startingDate = ref("");
const endingDate = ref("");

const reverseDate = (date) => {
  return date.split("-").reverse().join("-");
};

const userId = inject("userId");
const submitAnime = async () => {
  try {
    console.log(userId);
    if (!isLoggedIn.value) {
      console.log("You need to be logged in to add anime to your list");
      return;
    }
    if (checkIfPlanToWatch.value) {
      const res = await axios.post(
        `http://localhost:3000/add/anime/${userId.value}`,
        {
          title: data.value[isAdding.value].attributes.canonicalTitle,
          finished: false,
          episodes: 0,
          rating: 0,
          comment: "N/A",
          starting_date: reverseDate(startingDate.value),
          ending_date: "N/A",
          plan_to_watch: true,
        }
      );
      console.log(res);
      isAdding.value = null;
      return;
    } else {
      const res = await axios.post(
        `http://localhost:3000/add/anime/${userId.value}`,
        {
          title: data.value[isAdding.value].attributes.canonicalTitle,
          finished: animeCompleted.value,
          episodes:
            episodeCount.value === "Completato"
              ? data.value[isAdding.value].attributes.episodeCount
              : episodeCount.value,
          rating: userRating.value,
          comment: textAreaLength.value === "" ? "N/A" : textAreaLength.value,
          starting_date: reverseDate(startingDate.value),
          ending_date: reverseDate(endingDate.value),
          plan_to_watch: false,
        }
      );
      console.log(res);
      isAdding.value = null;
    }
  } catch (error) {
    console.error(error);
  }
};
const doThis = () => {
  console.log(checkIfPlanToWatch.value);
};
</script>

<style scoped>
.main-anime-card-div {
  margin-top: 55px;

  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(430px, 1fr));
  gap: 10px;
  width: 77%;
  margin-right: auto;
  margin-left: auto;
  user-select: none;
  -moz-user-select: none;
  margin-bottom: 150px;
}
@media (max-width: 550px) {
  .main-anime-card-div {
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  }
  .main-anime-display {
    flex-direction: column;
    align-items: center;
    justify-content: center;
    min-height: 500px;
    height: fit-content;
  }
  .anime-poster {
    height: 250px;
    margin-top: -50px;
  }
  .card-details {
    align-items: center;
    justify-content: center;
    padding-bottom: 20px;
    row-gap: 25px;
  }
}
.main-anime-display {
  display: flex;
  column-gap: 20px;
  padding: 10px 10px;
  padding-left: 20px;
  background-color: var(--dark-blue);
  border-radius: 15px;
  position: relative;
  box-shadow: -5px -5px 10px rgba(0, 0, 0, 0.2);
}
.title {
  font-weight: 500;
  font-size: 1.1rem;
  color: var(--yellow);
}
.title:hover {
  text-decoration: underline;
  cursor: pointer;
}
.title:active {
  filter: brightness(1.1);
}
.router-link {
  text-decoration: none;
}
.card-details {
  position: relative;
  display: flex;
  flex-direction: column;
  line-height: 30px;
  padding-top: 20px;
}
.episode-count,
.year {
  color: var(--dark-blue);
  background-color: var(--green);
  font-weight: 600;
  padding: 5px 10px;
  border-radius: 5px;
  box-shadow: inset -5px -5px 5px rgba(0, 0, 0, 0.15);
}
.year-ep-count {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  gap: 15px;
  align-items: center;
  justify-content: center;
}
.rating-div {
  position: absolute;
  bottom: 15px;
  right: 25px;
  background-color: var(--orange);
  padding: 5px 10px;
  border-radius: 25px;
  font-weight: 600;
  color: var(--dark-blue);
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 5px 5px 5px rgba(0, 0, 0, 0.15);
}
.anime-poster {
  border: 10px var(--green) solid;
  border-radius: 25px;
  box-shadow: 5px 5px 5px rgba(0, 0, 0, 0.15);
  cursor: pointer;
}
.anime-poster:hover {
  filter: brightness(1.1);
}
.watch-img {
  height: 45px;
  cursor: pointer;
}
.watch-img:hover {
  filter: brightness(0.65);
}
.trailer-link-div {
  position: absolute;
  bottom: 6px;
  right: 120px;
}
.trailer-link-div::before,
.add-to-list::before {
  content: attr(data-tooltip);
  position: absolute;
  top: 45px;
  left: -105px;
  z-index: 20;
  background-color: rgb(172, 172, 172);
  padding: 5px;
  border-radius: 10px;
  opacity: 0;
  pointer-events: none;
  scale: 0;
  transition: opacity 0.25s ease, scale 0.15s ease;
}
.add-to-list::before {
  top: 45px;
  left: -45px;
}
.trailer-link-div:hover::before,
.add-to-list:hover::before {
  opacity: 0.95;
  scale: 1;
}
.add-to-list {
  position: absolute;
  bottom: 12px;
  right: 180px;
  cursor: pointer;
  filter: drop-shadow(5px 5px 5px rgba(0, 0, 0, 0.15));
  z-index: 100;
}
.add-to-list-close {
  position: absolute;
  top: 10px;
  right: 12px;
  cursor: pointer;
  filter: drop-shadow(5px 5px 5px rgba(0, 0, 0, 0.15));
  z-index: 100;
}
.add-to-list:hover,
.add-to-list-close:hover {
  filter: drop-shadow(5px 5px 5px rgba(0, 0, 0, 0.15)) brightness(0.9);
}
.add-to-list:active,
.add-to-list-close:active {
  scale: 1.03;
}
.add-to-list-expanded-div {
  width: 85%;
  height: 90%;
  display: flex;
  flex-direction: column;
  align-items: center;
  border-radius: 15px;
  z-index: 20;
  color: var(--yellow);
  font-weight: 600;
  line-height: 25px;
  margin-right: auto;
  margin-left: auto;
}
.add-to-list-expanded-div h4 {
  margin-top: 10px;
  margin-bottom: 15px;
}
.add-episodes-div {
  display: flex;
  column-gap: 20px;
  align-items: center;
  justify-content: center;
}
.add-episodes-div input {
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
.add-episodes-div img {
  cursor: pointer;
}
.add-episodes-div img:hover {
  filter: brightness(0.9);
}
.add-episodes-div img:active {
  filter: brightness(0.8);
}
.add-episodes-div input:focus {
  outline: none;
}
.comment-div {
  margin-top: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.comment-div button {
  background-color: var(--green);
  color: var(--dark-blue);
  font-weight: 600;
  padding: 5px 10px;
  border-radius: 15px;
  cursor: pointer;
  margin-bottom: 10px;
  border: none;
}
.comment-div button:hover {
  filter: brightness(0.9);
}
.comment-div button:active {
  filter: brightness(0.8);
}
.comment-div textarea {
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
.comment-div textarea:focus {
  outline: none;
}
.loader {
  width: 32px;
  height: 16px;
  display: flex;
  margin-top: 50px;
}
.loader:before,
.loader:after {
  content: "";
  flex: 1;
  background: #3fb8af;
  transform-origin: top right;
  animation: l10-1 1s infinite;
}
.loader:after {
  background: #ff3d7f;
  transform-origin: top left;
  animation-delay: 0.15s;
}
.text-area-div {
  display: flex;
  position: relative;
  align-items: center;
  justify-content: center;
}
.text-area-div span {
  position: absolute;
  bottom: 0;
  right: 30px;
  color: var(--yellow);
  font-weight: 400;
  font-size: 0.6rem;
  opacity: 0.8;
}
.date-div {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  margin-top: 10px;
}
.date-div input {
  width: 150px;
  padding-right: 10px;
  height: 35px;
  border-radius: 15px;
  background-color: var(--green);
  color: var(--dark-blue);
  font-weight: 600;
  text-align: center;
  border: none;
  margin-bottom: 10px;
}
.date-div input:focus {
  outline: none;
}
.date-div h5 {
  margin-top: 15px;
  margin-bottom: 5px;
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
.submit-anime-div {
  margin-top: 10px;
  margin-bottom: 25px;
}
.submit-anime-div button {
  background-color: var(--green);
  color: var(--dark-blue);
  font-weight: 600;
  padding: 5px 10px;
  border-radius: 15px;
  cursor: pointer;
  border: none;
}
.submit-anime-div button:hover {
  filter: brightness(0.9);
}
.submit-anime-div button:active {
  filter: brightness(0.8);
}
@keyframes l10-1 {
  0%,
  5% {
    transform: rotate(0);
  }
  20%,
  30% {
    transform: rotate(90deg);
  }
  45%,
  55% {
    transform: rotate(180deg);
  }
  70%,
  80% {
    transform: rotate(270deg);
  }
  95%,
  100% {
    transform: rotate(360deg);
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

<template>
  <SubHeader />
  <div class="main-anime-card-div">
    <div v-for="(manga, i) in dataManga" :key="i" class="main-anime-display">
      <div v-if="!isLoading && isAdding !== i" class="poster-div">
        <router-link
          :to="{ name: 'MangaInfo', params: { id: i } }"
          class="router-link">
          <img
            class="anime-poster"
            :src="manga.attributes.posterImage.tiny"
            :alt="manga.attributes.canonicalTitle + ' poster'" />
        </router-link>
      </div>
      <div v-else-if="isLoading" class="wrap-loader-error">
        <div class="loader2"></div>
      </div>
      <div class="card-details" v-if="isAdding !== i && !isLoading">
        <div>
          <router-link
            :to="{ name: 'MangaInfo', params: { id: i } }"
            class="router-link">
            <span class="title">
              {{
                manga.attributes.canonicalTitle.length > 20 &&
                manga.attributes.abbreviatedTitles.length > 0
                  ? compareLengths(manga.attributes.abbreviatedTitles)
                  : manga.attributes.canonicalTitle
              }}
            </span>
          </router-link>
        </div>
        <div class="year-ep-count">
          <span class="episode-count">
            {{
              manga.attributes.chapterCount > 1
                ? "Capitoli: " + manga.attributes.chapterCount
                : manga.attributes.chapterCount === null
                ? "N/A"
                : "Manga"
            }}
          </span>
          <span class="year">
            {{ manga.attributes.status === "finished" ? "Finito" : "In corso" }}
          </span>
        </div>
      </div>
      <div class="rating-div" v-if="isAdding !== i && !isLoading">
        <img
          height="25px"
          src="../../assets/star-circle-svgrepo-com.svg"
          alt="" />
        <span>
          {{ ratingConverter(manga.attributes.averageRating) }}
        </span>
      </div>
      <div
        :class="isAdding !== i ? 'add-to-list' : 'add-to-list-close'"
        data-tooltip="Aggiungi alla tua lista">
        <img
          height="35px"
          :src="isAdding !== i ? plus : close"
          alt="add-to-list"
          @click="isAddingManga(i)"
          id="add-to-list-button"
          v-if="!isLoading" />
      </div>
      <div
        class="add-to-list-expanded-div"
        v-if="isAdding === i && !isLoadingAdding">
        <h3>
          Aggiungi: &nbsp;
          {{
            manga.attributes.canonicalTitle.length > 20 &&
            manga.attributes.abbreviatedTitles.length > 0
              ? compareLengths(manga.attributes.abbreviatedTitles)
              : cutString(manga.attributes.canonicalTitle)
          }}
          alla tua lista
        </h3>
        <div class="check-if-plan-to-watch">
          <div class="checkbox-wrapper-58">
            <label class="switch">
              <input
                @click="setCountersToZero"
                type="checkbox"
                v-model="checkIfPlanToRead" />
              <span class="slider"></span>
            </label>
          </div>
          <h4>manga da leggere</h4>
        </div>
        <h4
          v-if="
            !checkIfPlanToRead &&
            manga.attributes.chapterCount !== null &&
            manga.attributes.chapterCount !== 1
          ">
          Numero di capitoli letti
        </h4>
        <div
          v-if="
            !checkIfPlanToRead &&
            manga.attributes.chapterCount !== null &&
            manga.attributes.chapterCount !== 1
          "
          class="add-episodes-div">
          <img
            height="42px"
            src="@/assets/minus.svg"
            @click="decrementChapterCount(manga.attributes.chapterCount)"
            alt="minus-svg" />
          <span>{{ chapterCount }}</span>
          <img
            height="35px"
            src="@/assets/plus.svg"
            @click="incrementChapterCount(manga.attributes.chapterCount)"
            alt="plus-svg" />
        </div>
        <h4 v-if="!checkIfPlanToRead">Voto</h4>
        <div v-if="!checkIfPlanToRead" class="add-episodes-div">
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
              placeholder="Questo manga mi è piaciuto/non mi è piaciuto perché...">
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
          <h5 v-if="mangaCompleted">Data di fine (opz.)</h5>
          <input
            v-if="mangaCompleted"
            type="date"
            v-model="endingDate"
            min="2000-01-01"
            max="2025-12-31" />
        </div>
        <div class="submit-manga-div">
          <button @click="submitManga">Aggiungi</button>
        </div>
      </div>
      <div class="wrap-loader-error" v-if="isLoadingAdding && isAdding === i">
        <div class="loader2"></div>
        <div class="error-response-div">
          <span :class="addingResponse === 'Aggiunto!' ? 'error-green' : ''">
            {{ addingResponse }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { inject, ref, watch } from "vue";
const dataManga = inject("dataManga");
const isLoading = inject("isLoading");
import { compareLengths } from "../../utils/compareLengths";
import { getYear } from "../../utils/getYear";
import { ratingConverter } from "../../utils/ratingConverter";
import SubHeader from "@/components/SubHeader/SubHeader.vue";
import close from "@/assets/close.svg";
import plus from "@/assets/plus.svg";
import { cutString } from "../../utils/cutString";
import { timeout } from "../../utils/timeout";
import axios from "axios";

const isLoadingAdding = ref(false);
const checkIfPlanToRead = ref(false);
const textAreaLength = ref("");
const lengthCounter = ref(0);
const updateLengthCounter = () => {
  lengthCounter.value = textAreaLength.value.length;
};
const isAdding = ref(null);
const mangaCompleted = ref(false);
const addingResponse = ref(null);

const isAddingManga = (i) => {
  if (isAdding.value === i) {
    isAdding.value = null;
    setCountersToZero();
    return;
  }
  isAdding.value = i;
  setCountersToZero();
};
const setCountersToZero = () => {
  chapterCount.value = 0;
  userRating.value = 1;
  textAreaLength.value = "";
  lengthCounter.value = 0;
  displayText.value = false;
  startingDate.value = "";
  endingDate.value = "";
  checkIfPlanToRead.value = false;
  mangaCompleted.value = false;
};
const chapterCount = ref(0);
const checkIfMangaCompleted = (maxEp) => {
  if (chapterCount.value > maxEp - 1 || chapterCount.value === "Completato") {
    mangaCompleted.value = true;
    if (dataManga.value[isAdding.value].attributes.chapterCount === 1) {
      chapterCount.value = 1;
    }
    return;
  }
  mangaCompleted.value = false;
};
const incrementChapterCount = (maxEp) => {
  if (chapterCount.value >= maxEp - 1 || chapterCount.value === "Completato") {
    chapterCount.value = "Completato";
    checkIfMangaCompleted(maxEp);
    return;
  }
  checkIfMangaCompleted(maxEp);
  chapterCount.value++;
};
const decrementChapterCount = (maxEp) => {
  if (chapterCount.value <= 0) {
    chapterCount.value = 0;
    checkIfMangaCompleted(maxEp);
    return;
  }
  if (chapterCount.value === "Completato") {
    chapterCount.value = maxEp - 1;
    checkIfMangaCompleted(maxEp);
    return;
  }
  chapterCount.value--;
  checkIfMangaCompleted(maxEp);
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

const handleSubmitErrors = () => {
  isAdding.value = null;
  isLoadingAdding.value = false;
  addingResponse.value = null;
};

const userId = inject("userId");

const submitManga = async () => {
  try {
    isLoadingAdding.value = true;
    if (checkIfPlanToRead.value) {
      const res = await axios.post(
        `https://anime-vue-go-client-production.up.railway.app/add/manga/${userId.value}`,
        {
          title: dataManga.value[isAdding.value].attributes.canonicalTitle,
          finished: false,
          chapters: 0,
          rating: 0,
          comment: "",
          starting_date:
            startingDate.value !== "" ? reverseDate(startingDate.value) : "",
          ending_date: "",
          plan_to_read: true,
          manga_id: parseInt(dataManga.value[isAdding.value].id),
        },
        {
          withCredentials: true,
        }
      );
      if (res.status >= 200 && res.status <= 209) {
        addingResponse.value = "Aggiunto!";
        await timeout(1500);
        handleSubmitErrors();
      }
      return;
    } else {
      const res = await axios.post(
        `https://anime-vue-go-client-production.up.railway.app/add/manga/${userId.value}`,
        {
          title: dataManga.value[isAdding.value].attributes.canonicalTitle,
          finished: mangaCompleted.value,
          chapters:
            chapterCount.value === "Completato"
              ? data.value[isAdding.value].attributes.chapterCount
              : chapterCount.value,
          rating: userRating.value,
          comment: textAreaLength.value,
          starting_date: reverseDate(startingDate.value),
          ending_date: reverseDate(endingDate.value),
          plan_to_read: false,
          manga_id: parseInt(dataManga.value[isAdding.value].id),
        },
        {
          withCredentials: true,
        }
      );
      if (res.status >= 200 && res.status <= 209) {
        addingResponse.value = "Aggiunto!";
        await timeout(1500);
        handleSubmitErrors();
      }
    }
  } catch (error) {
    console.error(error);
    if (error.response.status === 409) {
      addingResponse.value = "Manga già presente nella tua lista";
      await timeout(2500);
      handleSubmitErrors();
    } else if (error.response.status === 400) {
      addingResponse.value = "Errore, riprova più tardi";
      await timeout(2500);
      handleSubmitErrors();
    } else if (error.response.status === 401) {
      addingResponse.value =
        "Devi essere registrato per aggiungere manga alla lista!";
      await timeout(3000);
      handleSubmitErrors();
    } else {
      addingResponse.value = "C'è stato un errore, non si sa perché!";
      await timeout(3000);
      handleSubmitErrors();
    }
  }
};

watch(
  () => dataManga.value,
  () => {
    isAdding.value = null;
    isLoadingAdding.value = false;
    addingResponse.value = null;
  }
);
</script>

<style scoped>
.main-anime-card-div {
  margin-top: 55px;

  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(430px, 1fr));
  gap: 10px;
  width: 80%;
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
.trailer-link-div::before {
  content: attr(data-tooltip);
  position: absolute;
  top: 45px;
  left: -105px;
  z-index: 10;
  background-color: rgb(172, 172, 172);
  padding: 5px;
  border-radius: 10px;
  opacity: 0;
  pointer-events: none;
  scale: 0;
  transition: opacity 0.25s ease, scale 0.15s ease;
}
.trailer-link-div:hover::before {
  opacity: 0.95;
  scale: 1;
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
.submit-manga-div {
  margin-top: 10px;
  margin-bottom: 25px;
}
.submit-manga-div button {
  background-color: var(--green);
  color: var(--dark-blue);
  font-weight: 600;
  padding: 5px 10px;
  border-radius: 15px;
  cursor: pointer;
  border: none;
}
.submit-manga-div button:hover {
  filter: brightness(0.9);
}
.submit-manga-div button:active {
  filter: brightness(0.8);
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
.loader2 {
  height: 35px;
  width: 80px;
  margin-left: 0px;
  background: radial-gradient(farthest-side, #ffd1d1 94%, #0000) 4px 22px/5px
      5px,
    radial-gradient(farthest-side, #ffd1d1 94%, #0000) 12px 18px/5px 5px,
    radial-gradient(farthest-side, #ffd1d1 94%, #0000) 3px 6px/8px 8px,
    radial-gradient(farthest-side, #eb8594 90%, #0000 94%) left/20px 100%,
    #bd3342;
  background-repeat: no-repeat;
  border-radius: 0 50px 50px 0;
  border-top-left-radius: 30px 40px;
  border-bottom-left-radius: 30px 40px;
  animation: l7 2.5s infinite steps(10);
  position: relative;
}
.loader::before {
  content: " ";
  position: absolute;
  inset: calc(50% - 8px) -10px calc(50% - 8px) auto;
  width: 15px;
  background: #bd3342;
  clip-path: polygon(0 50%, 100% 0, 70% 50%, 100% 100%);
}
@keyframes l7 {
  100% {
    width: 20px;
    margin-left: 60px;
  }
}
.error-response-div {
  color: var(--red);
  font-weight: 600;
  margin-top: 10px;
  text-align: center;
  width: 80%;
}
.wrap-loader-error {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
}
.error-green {
  color: greenyellow;
}
</style>

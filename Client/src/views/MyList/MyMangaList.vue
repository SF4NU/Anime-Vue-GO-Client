<template>
  <section>
    <h1>Manga List</h1>
    <!-- <div class="search-div">
      <input
        v-model="animeTitle"
        v-on:input="handleSearch"
        type="text"
        class="search-in-list"
        :placeholder="placeholderText" />
    </div> -->
    <div class="anime-container">
      <div v-for="(manga, i) in mangaListData" :key="i" class="anime-card">
        <div
          class="anime-card-wrap"
          v-if="isEditing !== i && manga.title !== searchedAnime">
          <h2>{{ manga.title }}</h2>
          <div>
            <span> Capitoli letti: {{ manga.chapters }} </span>
          </div>
          <div v-if="!manga.plan_to_read">
            <span> Voto assegnato: {{ manga.rating }} </span>
          </div>
          <div class="completed-div">
            <span>Finito: </span>
            <img height="20px" :src="manga.finished ? check : cross" alt="" />
          </div>
          <div class="dates-div">
            <span v-if="manga.starting_date !== ''">
              Inizio: {{ manga.starting_date }} </span
            ><br />
            <span v-if="!manga.plan_to_watch && manga.finished">
              Fine: {{ manga.ending_date }}
            </span>
          </div>
          <div
            class="plan-to-watch-div"
            v-if="!manga.finished && manga.plan_to_watch">
            <span> Da leggere ?</span>&nbsp;
            <img
              height="20px"
              :src="manga.plan_to_watch ? check : cross"
              alt="" />
          </div>
          <div class="comment-div" v-if="manga.comment !== ''">
            <button class="comment-button" :data-comment="manga.comment">
              Nota &nbsp; <img height="25px" src="@/assets/cursor.svg" alt="" />
            </button>
            <p v-if="showComment" class="show-comment">{{ manga.comment }}</p>
          </div>
          <div class="edit-div">
            <img
              @click="handleIsEditing(i, manga.manga_id)"
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
            <h2>{{ manga.title }}</h2>
          </div>
          <div class="check-if-plan-to-watch" v-if="!isLoading">
            <div class="checkbox-wrapper-58">
              <label class="switch">
                <input
                  @click="setCountersToZero"
                  type="checkbox"
                  v-model="plan_to_read" />
                <span class="slider"></span>
              </label>
            </div>
            <h4>Manga da leggere</h4>
          </div>
          <div class="add-values-div" v-if="!plan_to_read && !isLoading">
            <div class="episodes-title-div">
              <h4>Capitoli letti</h4>
            </div>
            <div class="add-episodes-div">
              <img
                height="30px"
                src="@/assets/minus-blue.svg"
                @click="
                  decrementChapterCount(singleMangaData.attributes.chapterCount)
                "
                alt="minus-svg" />
              <span>{{ chapters }}</span>
              <img
                height="25px"
                src="@/assets/plus-blue.svg"
                @click="
                  incrementChapterCount(singleMangaData.attributes.chapterCount)
                "
                alt="plus-svg" />
            </div>
          </div>
          <div v-if="!plan_to_read && !isLoading" class="add-values-div">
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
          <div class="date-div" v-if="!isLoading">
            <h5>Data di inizio (opz.)</h5>
            <input
              type="date"
              v-model="starting_date"
              min="2000-01-01"
              max="2025-12-31" />
            <h5 v-if="finished">Data di fine (opz.)</h5>
            <input
              v-if="finished"
              type="date"
              v-model="ending_date"
              min="2000-01-01"
              max="2025-12-31" />
          </div>
          <div class="submit-anime-div" v-if="!isLoading">
            <button @click="updateMangaCard">Modifica</button>
          </div>
          <div
            v-if="!isLoading"
            class="remove-anime-div"
            data-remove="Rimuovi dalla lista">
            <img
              @click="deleteMangaCard"
              height="30px"
              src="@/assets/bin.svg"
              alt="" />
          </div>
          <div v-if="isLoading && isEditing === i" class="wrap-loader-error">
            <div class="loader"></div>
            <div class="error-response-div">
              <span
                :class="addingResponse === 'Modificato!' ? 'error-blue' : ''">
                {{ addingResponse }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { onBeforeMount, onMounted, ref, watch } from "vue";
import { useRoute } from "vue-router";
import axios from "axios";
import check from "@/assets/check.svg";
import cross from "@/assets/cross-circle-svgrepo-com.svg";
import { timeout } from "../../utils/timeout";

const router = useRoute();
const userId = ref(router.params.userId);
const checkIfDataFetched = ref(false);
const mangaListData = ref(null);
const showComment = ref(false);
const isEditing = ref(null);
const chapters = ref(0);
const rating = ref(1);
const finished = ref(false);
const starting_date = ref("");
const ending_date = ref("");
const plan_to_read = ref(false);
const mangaID = ref(null);
const isLoading = ref(false);
const displayText = ref(false);
const textAreaLength = ref("");
const lengthCounter = ref(0);
const animeTitle = ref("");
const placeholderText = ref("");
const placeholders = ref([]);
const searchedAnime = ref("");
const addingResponse = ref("");

const singleMangaData = ref(null);

// const handleSearch = () => {
//   searchedAnime.value = mangaListData.value.filter((anime) =>
//     anime.title.toLowerCase().includes(animeTitle.value.toLowerCase())
//   );
//   console.log(searchedAnime.value);
// };

const setCountersToZero = () => {
  //la funzione resetta tutte le variabili che l'utente modifica nel client (quando serve)
  chapters.value = 0;
  rating.value = 1;
  finished.value = false;
  starting_date.value = "";
  ending_date.value = "";
  plan_to_read.value = false;
  textAreaLength.value = "";
};

const displayTextArea = () => {
  //questo pulsante toglie la finestra dei commenti
  displayText.value = !displayText.value;
};

const updateLengthCounter = () => {
  lengthCounter.value = textAreaLength.value.length;
};

const unreverseDate = (date) => {
  //rimette la data in formato (mm/dd/yyyy)
  return date.split("-").reverse().join("-");
};

const handleIsEditing = async (i, id) => {
  //inizia la sessione per fare modifiche all'anime su cui si è cliccato
  if (isEditing.value === i) {
    isEditing.value = null;
    setCountersToZero();
    return;
  }
  isEditing.value = i;
  setCountersToZero();
  mangaID.value = mangaListData.value[i].ID;
  chapters.value = mangaListData.value[i].chapters;
  rating.value = mangaListData.value[i].rating;
  textAreaLength.value = mangaListData.value[i].comment;
  lengthCounter.value = mangaListData.value[i].comment.length;
  starting_date.value = unreverseDate(mangaListData.value[i].starting_date);
  ending_date.value = unreverseDate(mangaListData.value[i].ending_date);
  plan_to_read.value = mangaListData.value[i].plan_to_read;
  isLoading.value = true;
  await findMangaByID(id);
  checkIfMangaCompleted(singleMangaData.value.attributes.chapterCount);
  isLoading.value = false;
};

const checkIfMangaCompleted = (maxEp) => {
  //controlla se i capitoli del manga sono al massimo per cambiare da maxEp a finito
  if (chapters.value > maxEp - 1 || chapters.value === "Finito") {
    finished.value = true;
    if (chapters.value !== "Finito") {
      chapters.value = "Finito";
    }
    if (singleMangaData.value.attributes.chapterCount === 1) {
      chapters.value = 1;
    }
    return;
  }
  finished.value = false;
};
const incrementChapterCount = (maxEp) => {
  //per il pulsante che aumenta il numero di capitoli
  if (chapters.value >= maxEp - 1 || chapters.value === "Finito") {
    chapters.value = "Finito";
    checkIfMangaCompleted(maxEp);
    return;
  }
  checkIfMangaCompleted(maxEp);
  chapters.value++;
};
const decrementChapterCount = (maxEp) => {
  //diminuisce il numero di capitoli
  if (chapters.value <= 0) {
    chapters.value = 0;
    checkIfMangaCompleted(maxEp);
    return;
  }
  if (chapters.value === "Finito") {
    chapters.value = maxEp - 1;
    checkIfMangaCompleted(maxEp);
    return;
  }
  chapters.value--;
  checkIfMangaCompleted(maxEp);
};

const incrementUserRating = () => {
  //per il pulsante che aumenta lo user score
  if (rating.value >= 10) {
    rating.value = 10;
    return;
  }
  rating.value += 0.5;
};

const decrementUserRating = () => {
  //diminuisce lo user score
  if (rating.value <= 1) {
    rating.value = 1;
    return;
  }
  rating.value -= 0.5;
};

const handleSubmitErrors = () => {
  //imposta le variabili della sessione degli errori a zero
  isEditing.value = null;
  isLoading.value = false;
  addingResponse.value = null;
};

const getMangaList = async () => {
  //fetch della lista di manga dell'utente
  try {
    const res = await axios.get(
      `https://anime-vue-go-client-production.up.railway.app/users/${userId.value}/manga`,
      {
        withCredentials: true,
      }
    );
    if (res.status >= 200 && res.status <= 209) {
      mangaListData.value = res.data.sort((a, b) =>
        a.title.localeCompare(b.title)
      );
      isEditing.value = null;
      isLoading.value = false;
    } else {
    }
  } catch (error) {
    console.error(error);
  }
};

const reverseDate = (date) => {
  //formatta la data in formato italiana (dd/mm/yyyy)
  return date.split("-").reverse().join("-");
};

const updateMangaCard = async () => {
  //aggiorna il manga con i nuovi dati inseriti dall'utente
  try {
    isLoading.value = true;
    const res = await axios.put(
      `https://anime-vue-go-client-production.up.railway.app/update/manga/${mangaID.value}`,
      {
        chapters:
          chapters.value === "Finito"
            ? singleMangaData.value.attributes.chapterCount
            : chapters.value,
        rating: rating.value,
        finished: finished.value,
        starting_date:
          starting_date.value !== "" ? reverseDate(starting_date.value) : "",
        ending_date:
          ending_date.value !== "" && finished.value
            ? reverseDate(ending_date.value)
            : "",
        plan_to_read: plan_to_read.value,
        comment: textAreaLength.value,
      },
      {
        withCredentials: true,
      }
    );

    if (res.status >= 200 && res.status <= 209) {
      addingResponse.value = "Modificato!";
      await timeout(2000);
      getMangaList();
      handleSubmitErrors();
    }
  } catch (error) {
    console.error(error);
    if (error.response.status === 400) {
      addingResponse.value = "Errore, riprova più tardi";
      await timeout(2500);
      handleSubmitErrors();
    } else {
      addingResponse.value = "C'è stato un errore, non si sa perché!";
      await timeout(3000);
      handleSubmitErrors();
    }
  }
};

const deleteMangaCard = async () => {
  //rimuove il manga selezionato dall'utente
  try {
    isLoading.value = true;
    const res = await axios.delete(
      `https://anime-vue-go-client-production.up.railway.app/delete/manga/${mangaID.value}`,
      {
        withCredentials: true,
      }
    );
    if (res.status >= 200 && res.status <= 209) {
      getMangaList();
    }
  } catch (error) {
    console.error(error);
    if (error.response.status === 400) {
      addingResponse.value = "Errore, riprova più tardi";
      await timeout(2500);
      handleSubmitErrors();
    } else {
      addingResponse.value = "C'è stato un errore, non si sa perché!";
      await timeout(3000);
      handleSubmitErrors();
    }
  }
};

onBeforeMount(() => {
  //al caricamento del componente fa il fetch di manga
  if (!checkIfDataFetched.value) {
    checkIfDataFetched.value = true;
    getMangaList();
    // let i = 0;
    // let j = 0;
    // let forward = true;
    // setInterval(() => {
    //   if (forward) {
    //     if (j < placeholders.value[i].length) {
    //       placeholderText.value += placeholders.value[i][j];
    //       j++;
    //     } else {
    //       forward = false;
    //     }
    //   } else {
    //     if (j > 0) {
    //       placeholderText.value = placeholderText.value.slice(0, -1);
    //       j--;
    //     } else {
    //       forward = true;
    //       i++;
    //       if (i === placeholders.value.length) {
    //         i = 0;
    //       }
    //     }
    //   }
    // }, 150);
  }
});

const findMangaByID = async (id) => {
  //trova il manga in base all'id
  try {
    const res = await axios.get(
      `https://kitsu.io/api/edge/manga?filter[id]=${id}&page[limit]=1`
    );
    singleMangaData.value = res.data.data[0];
  } catch (error) {
    console.error(error);
  }
};

// watch(
//   () => mangaListData.value,
//   () => {
//     isEditing.value = null;
//     isLoadingAdding.value = false;
//     addingResponse.value = null;
//   }
// );
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
  min-width: 260px;
  height: fit-content;
  border: 1px solid black;
  margin: 10px;
  padding: 20px;
  background: linear-gradient(145deg, #67ebc3, #28886b);
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
  transition: filter 0.25s ease, scale 0.2s ease;
}
.edit-div:hover {
  filter: brightness(1.2);
  scale: 1.2;
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
  transition: filter 0.25s ease, scale 0.2s ease;
}
.add-values-div img:hover {
  scale: 1.2;
  filter: brightness(0.9);
}
.add-values-div img:active {
  scale: 0.95;
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
  color: #67ebc3;
  font-weight: 600;
  padding: 5px 10px;
  border-radius: 15px;
  cursor: pointer;
  margin-bottom: 10px;
  border: none;
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.15);
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
  background-color: var(--dark-blue);
  color: #67ebc3;
  font-weight: 600;
  text-align: center;
  border: none;
  margin-bottom: 10px;
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.15);
}
.date-div input:focus {
  outline: none;
}
.date-div h5 {
  margin-top: 15px;
  margin-bottom: 5px;
}
.submit-anime-div {
  margin-top: 10px;
  margin-bottom: 25px;
}
.submit-anime-div button {
  background-color: var(--dark-blue);
  color: var(--yellow);
  font-weight: 600;
  padding: 10px 15px;
  border-radius: 15px;
  cursor: pointer;
  border: none;
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.15);
}
.submit-anime-div button:hover {
  filter: brightness(0.9);
}
.submit-anime-div button:active {
  filter: brightness(0.8);
}
.remove-anime-div {
  position: absolute;
  bottom: 12px;
  right: 15px;
  cursor: pointer;
  transition: filter 0.25s ease, scale 0.2s ease;
}
.remove-anime-div::after {
  content: attr(data-remove);
  height: 30px;
  width: 150px;
  position: absolute;
  top: 35px;
  right: -65px;
  background: rgba(25, 61, 93, 0.8);
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
  backdrop-filter: blur(8.5px);
  -webkit-backdrop-filter: blur(8.5px);
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.18);
  word-break: break-all;
  padding: 5px;
  font-weight: 600;
  text-align: center;
  color: #67ebc3;
  transition: opacity 0.25s ease, scale 0.2s ease;
  opacity: 0;
  scale: 0;
}
.remove-anime-div:hover::after {
  opacity: 1;
  scale: 1;
}
.remove-anime-div:hover {
  filter: brightness(1.1);
  scale: 1.1;
}

.loader {
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
  animation: l7 1.5s infinite steps(10);
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
  --outline_color: var(--dark-blue);
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
section {
  margin-bottom: 150px;
}
.search-div {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 50px;
}
.search-in-list {
  width: 200px;
  height: 30px;
  border-radius: 20px;
  padding: 5px;
  text-indent: 15px;
  font-size: 14px;
  transition: border-radius 0.525s ease, width 0.525s ease;
  background: linear-gradient(145deg, var(--green), var(--dark-blue));
  font-family: Rubik;
  color: var(--yellow);
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.15);
}
.search-in-list:focus {
  border-radius: 0;
  outline: none;
  width: 300px;
  background: linear-gradient(145deg, var(--dark-blue), var(--green));
}
.search-in-list::placeholder {
  color: rgba(255, 255, 255, 0.671);
  font-family: Rubik;
}
.error-blue {
  color: var(--dark-blue);
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
</style>

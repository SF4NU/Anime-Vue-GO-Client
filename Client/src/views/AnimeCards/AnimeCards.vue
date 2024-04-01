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
                  : anime.attributes.canonicalTitle
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
      <div class="add-to-list" data-tooltip="Aggiungi alla tua lista">
        <img
          height="35px"
          :src="isAdding !== i ? plus : close"
          alt="add-to-list"
          @click="isAddingAnime(i)"
          id="add-to-list-button" />
      </div>
      <div class="add-to-list-expanded-div" v-if="isAdding === i">
        <h3>
          Aggiungi: <br />
          {{
            anime.attributes.canonicalTitle.length > 20 &&
            anime.attributes.abbreviatedTitles.length > 0
              ? compareLengths(anime.attributes.abbreviatedTitles)
              : anime.attributes.canonicalTitle
          }}
          <br />
          alla tua lista
        </h3>
        <h4>Numero di episodi visti</h4>
        <div class="add-episodes-div">
          <img
            height="42px"
            src="@/assets/minus.svg"
            @click="decrement(anime.attributes.episodeCount)"
            alt="minus-svg" />
          <input type="text" :value="episodeCount" />
          <img
            height="35px"
            src="@/assets/plus.svg"
            @click="increment(anime.attributes.episodeCount)"
            alt="plus-svg" />
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
import SubHeader from "@/components/SubHeader/SubHeader.vue";
import close from "@/assets/close.svg";
import plus from "@/assets/plus.svg";

const isAdding = ref(null);
console.log(isAdding.value);
const isAddingAnime = (i) => {
  if (isAdding.value === i) {
    isAdding.value = null;
    document.getElementById("add-to-list-button").style.scale = 1;
    setEpisodeCountToZero();
    return;
  }
  isAdding.value = i;
  document.getElementById("add-to-list-button").style.scale = 1.2;
  setEpisodeCountToZero();
};

const episodeCount = ref(0);
const increment = (maxEp) => {
  if (episodeCount.value >= maxEp - 1 || episodeCount.value === "Completato") {
    episodeCount.value = "Completato";
    return;
  }
  episodeCount.value++;
};
const decrement = (maxEp) => {
  if (episodeCount.value <= 0) {
    episodeCount.value = 0;
    return;
  }
  if (episodeCount.value === "Completato") {
    episodeCount.value = maxEp - 1;
    return;
  }
  episodeCount.value--;
};
const setEpisodeCountToZero = () => {
  episodeCount.value = 0;
};
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
    height: 500px;
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
.loader {
  width: 50px;
  aspect-ratio: 1;
  border-radius: 50%;
  border: 8px solid #514b82;
  animation: l20-1 0.8s infinite linear alternate, l20-2 1.6s infinite linear;
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
.add-to-list:hover {
  filter: drop-shadow(5px 5px 5px rgba(0, 0, 0, 0.15)) brightness(0.9);
}
.add-to-list:active {
  scale: 1.03;
}
.add-to-list-expanded-div {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  border-radius: 15px;
  z-index: 50;
  color: var(--yellow);
  font-weight: 600;
}
.add-to-list-expanded-div h4 {
  margin-top: 20px;
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
@keyframes l20-1 {
  0% {
    clip-path: polygon(50% 50%, 0 0, 50% 0%, 50% 0%, 50% 0%, 50% 0%, 50% 0%);
  }
  12.5% {
    clip-path: polygon(
      50% 50%,
      0 0,
      50% 0%,
      100% 0%,
      100% 0%,
      100% 0%,
      100% 0%
    );
  }
  25% {
    clip-path: polygon(
      50% 50%,
      0 0,
      50% 0%,
      100% 0%,
      100% 100%,
      100% 100%,
      100% 100%
    );
  }
  50% {
    clip-path: polygon(
      50% 50%,
      0 0,
      50% 0%,
      100% 0%,
      100% 100%,
      50% 100%,
      0% 100%
    );
  }
  62.5% {
    clip-path: polygon(
      50% 50%,
      100% 0,
      100% 0%,
      100% 0%,
      100% 100%,
      50% 100%,
      0% 100%
    );
  }
  75% {
    clip-path: polygon(
      50% 50%,
      100% 100%,
      100% 100%,
      100% 100%,
      100% 100%,
      50% 100%,
      0% 100%
    );
  }
  100% {
    clip-path: polygon(
      50% 50%,
      50% 100%,
      50% 100%,
      50% 100%,
      50% 100%,
      50% 100%,
      0% 100%
    );
  }
}
@keyframes l20-2 {
  0% {
    transform: scaleY(1) rotate(0deg);
  }
  49.99% {
    transform: scaleY(1) rotate(135deg);
  }
  50% {
    transform: scaleY(-1) rotate(0deg);
  }
  100% {
    transform: scaleY(-1) rotate(-135deg);
  }
}
</style>

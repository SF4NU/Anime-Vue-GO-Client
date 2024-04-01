<template>
  <SubHeader />
  <div class="main-anime-card-div">
    <div v-for="(manga, i) in dataManga" :key="i" class="main-anime-display">
      <div v-if="!isLoading" class="poster-div">
        <img
          class="anime-poster"
          :src="manga.attributes.posterImage.tiny"
          :alt="manga.attributes.canonicalTitle + ' poster'" />
      </div>
      <div v-else-if="isLoading" class="loader"></div>
      <div class="card-details">
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
      <div class="rating-div">
        <img
          height="25px"
          src="../../assets/star-circle-svgrepo-com.svg"
          alt="" />
        <span>
          {{ ratingConverter(manga.attributes.averageRating) }}
        </span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { inject, ref, watch } from "vue";
// const props = defineProps(["dataManga", "isLoading"]);
const dataManga = inject("dataManga");
const isLoading = inject("isLoading");
import { compareLengths } from "../../utils/compareLengths";
import { getYear } from "../../utils/getYear";
import { ratingConverter } from "../../utils/ratingConverter";
import SubHeader from "@/components/SubHeader/SubHeader.vue";
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
</style>

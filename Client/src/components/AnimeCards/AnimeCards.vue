<template>
  <div class="main-anime-card-div">
    <div v-for="(anime, i) in props.data" :key="i" class="main-anime-display">
      <div v-if="!props.isLoading">
        <img
          class="anime-poster"
          :src="anime.attributes.posterImage.tiny"
          alt="" />
      </div>
      <div v-else-if="props.isLoading" class="loader"></div>
      <div class="card-details">
        <div>
          <span class="title">
            {{
              anime.attributes.canonicalTitle.length > 20 &&
              anime.attributes.abbreviatedTitles.length > 0
                ? compareLengths(anime.attributes.abbreviatedTitles)
                : anime.attributes.canonicalTitle
            }}
          </span>
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
      <div class="rating-div">
        <img
          height="25px"
          src="../../assets/star-circle-svgrepo-com.svg"
          alt="" />
        <span>
          {{ ratingConverter(anime.attributes.averageRating) }}
        </span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from "vue";
const props = defineProps(["data", "isLoading"]);
import { compareLengths } from "../../utils/compareLengths";
import { getYear } from "../../utils/getYear";
import { ratingConverter } from "../../utils/ratingConverter";

watch(
  () => props.isLoading,
  () => {
    console.log(props.isLoading);
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
    margin-top: -75px;
    margin-bottom: 25px;
  }
  .card-details {
    align-items: center;
    justify-content: center;
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
.card-details {
  position: relative;
  display: flex;
  flex-direction: column;
  line-height: 30px;
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

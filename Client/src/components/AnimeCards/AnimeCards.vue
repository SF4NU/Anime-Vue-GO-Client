<template>
  <div class="main-anime-card-div">
    <div v-for="(anime, i) in props.data" :key="i">
      <p>
        {{
          anime.attributes.abbreviatedTitles.length > 0
            ? compareLengths(anime.attributes.abbreviatedTitles)
            : anime.attributes.canonicalTitle
        }}
      </p>
      <div v-if="!props.isLoading">
        <img :src="anime.attributes.posterImage.tiny" alt="" />
      </div>
      <div v-else-if="props.isLoading" class="loader"></div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from "vue";
const props = defineProps(["data", "isLoading"]);
import { compareLengths } from "../../utils/compareLengths";

watch(
  () => props.isLoading,
  () => {
    console.log(props.isLoading);
  }
);
</script>

<style scoped>
.main-anime-card-div {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(110px, 1fr));
  column-gap: 10px;
}
.loader {
  width: 50px;
  aspect-ratio: 1;
  border-radius: 50%;
  border: 8px solid #514b82;
  animation: l20-1 0.8s infinite linear alternate, l20-2 1.6s infinite linear;
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

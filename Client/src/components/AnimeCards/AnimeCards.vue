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
      <img :src="anime.attributes.posterImage.tiny" alt="" />
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
const props = defineProps(["data"]);
import { compareLengths } from "../../utils/compareLengths";
</script>

<style scoped>
.main-anime-card-div {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(110px, 1fr));
  column-gap: 10px;
}
</style>

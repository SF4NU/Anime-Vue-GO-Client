<template>
  <section>
    <div class="anime-main-card">
      <div class="poster-div">
        <img :src="props.data[passIndex].attributes.posterImage.small" alt="" />
      </div>
      <div class="anime-info-top-right-div">
        <div class="title-div">
          <span>{{ props.data[passIndex].attributes.canonicalTitle }}</span>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref } from "vue";

const props = defineProps(["passIndex", "data"]);
</script>

<style scoped>
section {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.anime-main-card {
  margin-top: 100px;
  min-height: 500px;
  width: clamp(250px, 80%, 1400px);
  background-color: var(--dark-blue);
  border-radius: 25px;
  display: flex;
  flex-direction: row;
}
.poster-div {
  margin-left: 50px;
  margin-top: 50px;
}
.poster-div img {
  border: 15px var(--green) solid;
  border-radius: 10px;
  filter: drop-shadow(5px 5px 15px rgba(0, 0, 0, 0.25));
}
</style>

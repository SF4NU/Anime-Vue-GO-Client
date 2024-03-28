<template>
  <header>
    <div class="left-header">Anime</div>
    <div class="right-header">
      <span @click="$emit('backToAnime')">Anime</span>
      <span @click="$emit('toManga')">Manga</span>
      <span>My List</span>
      <span>Registrati</span>
    </div>
  </header>
</template>

<script setup></script>

<style scoped>
header {
  height: 60px;
  background-color: var(--dark-blue); /* Add this line */
  color: var(--yellow);
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-weight: 500;
  font-size: 1rem;
  box-shadow: 0px 5px 5px rgba(0, 0, 0, 0.25);
}
header span {
  cursor: pointer;
  user-select: none;
}

.left-header {
  width: 80px;
  text-align: center;
}
.right-header {
  width: 60%;
  display: flex;
  justify-content: space-around;
}
</style>

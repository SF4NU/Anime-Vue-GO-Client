<template>
  <header>
    <div class="left-header">Anime</div>
    <div class="right-header">
      <router-link to="/" class="router-link">Anime</router-link>
      <router-link to="/manga" class="router-link">Manga</router-link>
      <span>My List</span>
      <router-link to="/register" class="router-link">Registrati</router-link>
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
.router-link {
  color: var(--yellow);
  text-decoration: none;
}
.router-link:hover {
  text-decoration: underline;
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

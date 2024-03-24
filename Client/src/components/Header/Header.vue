<template>
  <header>
    <div class="left-header">Anime</div>
    <div class="right-header">
      <span>Manga</span>
      <span>My List</span>
      <span>Register</span>
    </div>
  </header>
</template>

<script setup></script>

<style scoped>
header {
  height: 60px;
  background-color: #43aa8b;
  color: var(--dark-blue);
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-weight: 500;
  font-size: 1rem;
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

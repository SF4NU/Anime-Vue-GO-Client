<template>
  <header>
    <div class="search-div">
      <input
        @keypress.enter="sendInput($emit)"
        v-model="userInput"
        class="search-bar"
        type="text"
        placeholder="Inserisci Anime" />
      <button @click="sendInput($emit)">Cerca</button>
    </div>
    <div class="filter-div">
      <button class="dropdown-button">Filtri</button>
      <div class="dropdown-menu">
        <span> Filtro 1 </span>
        <span> Filtro 2 </span>
        <span> Filtro 3 </span>
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref } from "vue";
const userInput = ref("");

const sendInput = (emit) => {
  emit("getUserInput", userInput.value);
};
</script>

<style scoped>
header {
  height: 80px;
  background-color: rgb(197, 197, 197);
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-around;
  width: 100%;
}
.search-div {
  display: flex;
  column-gap: 20px;
}
.search-bar {
  text-indent: 15px;
  padding: 10px 20px;
  border: none;
  border-radius: 25px;
  transition: border-radius 0.35s ease;
  font-family: Rubik, Arial;
  box-shadow: inset -5px -5px 5px rgba(0, 0, 0, 0.15);
}
.search-div button {
  border: none;
  padding: 10px 20px;
  border-radius: 25px;
  cursor: pointer;
  transition: filter 0.15s ease, scale 0.15s linear;
  font-family: Rubik, Arial;
  font-weight: 500;
  box-shadow: inset -5px -5px 5px rgba(0, 0, 0, 0.15);
}
.search-div button:hover {
  filter: brightness(0.95);
}
.search-div button:active {
  scale: 1.05;
}
.search-bar:focus {
  outline: none;
  border-radius: 0px;
}
.filter-div {
  position: relative;
  display: inline-block;
}
.dropdown-menu {
  display: none;
  position: absolute;
  background-color: #f9f9f9;
  min-width: 160px;
  box-shadow: 0px 8px 16px 0px rgba(0, 0, 0, 0.2);
  z-index: 1;
}
.dropdown-button {
  border: none;
  background-color: var(--orange);
  color: var(--dark-blue);
  font-weight: 600;
  padding: 8px 30px;
  font-size: 1.05rem;
  border-radius: 25px;
  box-shadow: inset 5px 5px 15px rgba(255, 255, 255, 0.25);
  cursor: pointer;
}
.dropdown-menu span {
  display: block;
  background-color: var(--green);
  padding: 12px 16px;
  text-decoration: none;
  cursor: pointer;
  width: 100%;
  text-align: center;
  color: var(--dark-blue);
  font-weight: 500;
}
.filter-div:hover .dropdown-menu {
  display: block;
  flex-direction: column;
}
.dropdown-menu span:hover {
  filter: brightness(1.1);
}
@media (max-width: 450px) {
  .search-div button {
    display: none;
  }
}
</style>

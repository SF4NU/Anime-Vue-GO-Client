<template>
  <header>
    <div class="left-header">Anime</div>
    <div class="right-header">
      <router-link to="/" class="router-link">Anime</router-link>
      <router-link to="/manga" class="router-link">Manga</router-link>
      <span>My List</span>
      <router-link to="/register" class="router-link">Registrati</router-link>
    </div>
    <div class="hamburger-div">
      <img
        class="hamburger-icon"
        src="../../assets/hamburger-2-svgrepo-com.svg"
        alt="hamburger-icon"
        @click="handleToggleHeader()"
        :class="!toggleHeader ? '' : 'display-icon'" />
      <img
        class="hamburger-icon"
        src="../../assets/cross-svgrepo-com.svg"
        alt="hamburger-icon"
        @click="handleToggleHeader()"
        :class="!toggleHeader ? 'display-icon' : ''" />
    </div>
    <div :class="!toggleHeader ? 'toggle-header' : 'phone-header'">
      <router-link to="/" class="router-link"
        ><span @click="handleToggleHeader()">Anime</span></router-link
      >
      <router-link to="/manga" class="router-link"
        ><span @click="handleToggleHeader()">Manga</span></router-link
      >
      <span>My List</span>
      <router-link to="/register" class="router-link"
        ><span @click="handleToggleHeader()">Registrati</span></router-link
      >
    </div>
  </header>
</template>

<script setup>
import { ref, onUpdated } from "vue";

const toggleHeader = ref(false);

const handleToggleHeader = () => {
  toggleHeader.value = !toggleHeader.value;
};
onUpdated(() => {
  if (toggleHeader.value) {
    document.body.style.overflow = "hidden";
  } else {
    document.body.style.overflow = "auto";
  }
});
</script>

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
.hamburger-div {
  margin-right: 20px;
  display: none;
}
.toggle-header {
  display: none;
}
.phone-header {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: absolute;
  top: 60px;
  left: 0;
  width: 100%;
  height: 100vh;
  background-color: var(--dark-blue);
  z-index: 1;
  line-height: 150px;
  font-size: 2.5rem;
  animation: slideDown 0.85s ease forwards;
}
.hamburger-icon {
  height: 30px;
  cursor: pointer;
}
.display-icon {
  display: none;
}
@keyframes slideDown {
  from {
    top: -100vh;
  }
  to {
    top: 60px;
  }
}
@media (max-width: 550px) {
  .right-header {
    display: none;
  }
  .hamburger-div {
    display: block;
  }
}
</style>

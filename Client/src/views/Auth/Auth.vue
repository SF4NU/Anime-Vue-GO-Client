<template>
  <section>
    <div class="main-auth-div">
      <div class="registration-div">
        <h1>Iscrizione</h1>
        <form @submit.prevent>
          <div>
            <label for="email">Email</label>
            <input
              v-model="email"
              type="email"
              id="email"
              placeholder="esempio@yourmail.com" />
          </div>
          <div>
            <label for="username">Username</label>
            <input
              v-model="username"
              type="text"
              id="username"
              placeholder="utente123" />
          </div>
          <div>
            <label for="password">Password</label>
            <input
              v-model="password"
              type="password"
              id="password"
              placeholder="••••••••" />
          </div>
          <div>
            <label for="confirmPassword">Conferma Password</label>
            <input
              v-model="confirmPassword"
              @change=""
              type="password"
              id="confirmPassword"
              :class="!passwordsMatch ? 'no-match' : ''"
              placeholder="••••••••" />
          </div>
          <button type="submit" @click="createUser">Registrati</button>
        </form>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref } from "vue";
import axios, { Axios } from "axios";
const email = ref("");
const username = ref("");
const password = ref("");
const confirmPassword = ref("");
const passwordsMatch = ref(true);

const checkIfPasswordsMatch = () => {
  if (
    password.value !== confirmPassword.value &&
    confirmPassword.value !== "" &&
    password.value !== ""
  ) {
    passwordsMatch.value = false;
    confirmPassword.value = "";
    setTimeout(() => {
      passwordsMatch.value = true;
    }, 1000);
    return false;
  } else {
    passwordsMatch.value = true;
    return true;
  }
};

const createUser = async () => {
  try {
    if (checkIfPasswordsMatch()) {
      const res = await axios.post("http://localhost:3000/register", {
        email: email.value,
        username: username.value,
        password: password.value,
      });
      console.log(res);
      if (res.request.status >= 200 && res.request.status <= 209) {
        email.value = "";
        username.value = "";
        password.value = "";
        confirmPassword.value = "";
      }
    }
  } catch (error) {
    console.error(error);
  }
};
</script>

<style scoped>
.main-auth-div {
  display: flex;
  justify-content: center;
  align-items: center;
}
.registration-div {
  margin-top: 100px;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  border-radius: 10px;
  background-color: var(--dark-blue);
  color: var(--green);
  font-weight: 500;
  font-size: 1.5rem;
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.5);
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.25);
}
h1 {
  margin: 0;
  margin-bottom: 10px;
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.5);
}
form {
  display: flex;
  flex-direction: column;
  gap: 5px;
}
input {
  padding: 10px;
  border-radius: 25px;
  border: none;
  text-indent: 15px;
  font-size: 14px;
  transition: border-radius 0.525s ease;
}
input:focus {
  border-radius: 0;
  outline: none;
}
button {
  padding: 5px 15px;
  border-radius: 25px;
  border: none;
  background-color: var(--green);
  color: var(--dark-blue);
  font-weight: 600;
  font-size: 1.5rem;
  width: fit-content;
  margin-right: auto;
  margin-left: auto;
  margin-top: 25px;
  margin-bottom: 25px;
}
button:hover {
  cursor: pointer;
  filter: brightness(1.1);
}
button:active {
  filter: brightness(0.9);
}
.registration-div div {
  display: flex;
  flex-direction: column;
  text-align: center;
  line-height: 60px;
}
@media (max-width: 550px) {
  .registration-div {
    margin-top: 20px;
    scale: 0.9;
  }
}
.no-match {
  animation: shake 0.25s ease 2;
  background-color: rgb(255, 142, 142);
}
@keyframes shake {
  0% {
    transform: translateX(0);
  }
  25% {
    transform: translateX(-5px);
  }
  50% {
    transform: translateX(5px);
  }
  75% {
    transform: translateX(-5px);
  }
  100% {
    transform: translateX(0);
  }
}
</style>

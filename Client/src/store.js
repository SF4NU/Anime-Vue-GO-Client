// import { createStore } from "vuex";

// export default createStore({
//   data: null,
//   searchedAnime: "",},
//   mutations: {
//     setData(state, data) {
//       state.data = data;
//     },
//   },
//   actions: {
//     async fetchAnime({ commit }) {
//       const response = await fetch(`https://kitsu.io/api/edge/manga?filter[text]=${searchedAnime.value}&page[limit]=15`);
//       const data = await response.json().data;
//       commit("setData", data.top);
//     },
//   },);

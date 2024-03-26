<template>
  <section>
    <div class="anime-main-card">
      <div class="back-div">
        <img
          @click="$emit('backToAnime')"
          src="../../assets/back-2-svgrepo-com.svg"
          alt="back-icon" />
      </div>
      <div class="top-row-div">
        <div class="poster-div">
          <img
            :src="props.data[passIndex].attributes.posterImage.small"
            alt="" />
        </div>
        <div class="anime-info-top-right-div">
          <div class="title-div">
            <span>{{ props.data[passIndex].attributes.canonicalTitle }}</span>
          </div>
          <div class="sub-titles">
            <span>Titoli alternativi: </span>
            <span>
              {{
                props.data[passIndex].attributes.titles.en_jp
                  ? props.data[passIndex].attributes.titles.en_jp
                  : ""
              }} </span
            >&nbsp;&nbsp;&nbsp;&nbsp;
            <span>{{
              props.data[passIndex].attributes.titles.ja_jp
                ? props.data[passIndex].attributes.titles.ja_jp
                : ""
            }}</span>
          </div>
          <div class="episodes-div">
            <span
              >Episodi:
              {{ props.data[passIndex].attributes.episodeCount }}</span
            >
            &nbsp;&nbsp;&nbsp;&nbsp;
            <span
              >Durata episodio:
              {{ props.data[passIndex].attributes.episodeLength }}</span
            >
          </div>
          <div class="dates-div">
            <span>In onda dal: </span>
            <span>{{ props.data[passIndex].attributes.startDate }}</span>
            &nbsp;&nbsp;&nbsp;&nbsp;
            <span>In onda fino al: </span>
            <span>
              {{
                props.data[passIndex].attributes.endDate !== null
                  ? props.data[passIndex].attributes.endDate
                  : "N/A"
              }}</span
            >
          </div>
          <div class="age-rating-div">
            <span>Guida alle Classificazioni per Età: </span>
            <span>{{ props.data[passIndex].attributes.ageRatingGuide }}</span>
          </div>
          <div class="rating-div">
            <img
              height="25px"
              src="../../assets/star-circle-svgrepo-com.svg"
              alt="anime-rating-star" />
            <span>
              {{
                ratingConverter(props.data[passIndex].attributes.averageRating)
              }}
            </span>
          </div>
          <div class="rating-rank">
            <span> Top: </span
            ><span>
              {{ props.data[passIndex].attributes.ratingRank }}
            </span>
          </div>
          <div class="show-type-div">
            <span>Tipologia: </span
            ><span>
              {{ props.data[passIndex].attributes.showType }}
            </span>
          </div>
          <div class="popularity-div">
            <span> Popolarità: </span
            ><span>
              {{ props.data[passIndex].attributes.userCount }}
            </span>
            <span> utenti</span>
          </div>
          <div class="trailer-div" :data-tooltip1="`Youtube link`">
            <span> Guarda trailer: </span
            ><span
              ><a
                :href="`https://www.youtube.com/watch?v=${props.data[passIndex].attributes.youtubeVideoId}`"
                target="_blank"
                ><img
                  height="35px"
                  class="watch-img"
                  src="../../assets/youtube-svgrepo-com.svg"
                  alt="watch-youtube-icon" /></a
            ></span>
          </div>
        </div>
      </div>

      <div class="bottom-synopsis-div">
        <span>Descrizione: </span>
        <p>{{ props.data[passIndex].attributes.description }}</p>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { ratingConverter } from "@/utils/ratingConverter";

const props = defineProps(["passIndex", "data"]);

const scrollToTop = () => {
  window.scrollTo({
    top: 0,
    behavior: "smooth",
  });
};

onMounted(() => {
  scrollToTop();
});
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
  flex-direction: column;
  margin-bottom: 150px;
  padding-bottom: 50px;
  position: relative;
}
.top-row-div {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  margin-top: 50px;
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
.title-div {
  margin-top: 50px;
  margin-left: 50px;
}
.title-div span {
  font-size: 2.5rem;
  font-weight: 700;
  color: var(--orange);
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.5);
}
.sub-titles {
  margin-top: 10px;
  margin-left: 50px;
}
.sub-titles span {
  font-size: 1.1rem;
  font-weight: 500;
  color: var(--green);
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.5);
}
.episodes-div {
  margin-top: 10px;
  margin-left: 50px;
}
.episodes-div span {
  font-size: 1.1rem;
  font-weight: 500;
  color: var(--green);
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.5);
}
.dates-div {
  margin-top: 10px;
  margin-left: 50px;
}
.dates-div span {
  font-size: 1.1rem;
  font-weight: 500;
  color: var(--green);
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.5);
}
.dates-div span:nth-child(2),
.dates-div span:nth-child(4) {
  color: var(--orange);
}
.age-rating-div {
  margin-top: 10px;
  margin-left: 50px;
}
.age-rating-div span {
  font-size: 1.1rem;
  font-weight: 500;
  color: var(--green);
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.5);
}
.anime-info-top-right-div {
  display: flex;
  flex-direction: column;
  margin-top: 50px;
  position: relative;
  width: 100%;
  padding-right: 20px;
  line-height: 30px;
}
.rating-div {
  position: absolute;
  bottom: 15px;
  width: fit-content;
  right: 25px;
  background-color: var(--orange);
  padding: 5px 10px;
  border-radius: 25px;
  font-weight: 600;
  color: var(--dark-blue);
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 5px 5px 5px rgba(0, 0, 0, 0.15);
  font-size: 1.5rem;
}
.rating-rank {
  position: absolute;
  bottom: 15px;
  width: fit-content;
  right: 145px;
  background-color: var(--orange);
  padding: 5px 10px;
  border-radius: 25px;
  font-weight: 600;
  color: var(--dark-blue);
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 5px 5px 5px rgba(0, 0, 0, 0.15);
  font-size: 1.5rem;
}
.show-type-div {
  margin-top: 10px;
  margin-left: 50px;
}
.show-type-div span {
  font-size: 1.1rem;
  font-weight: 500;
  color: var(--green);
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.5);
}
.popularity-div {
  margin-top: 10px;
  margin-left: 50px;
}
.popularity-div span {
  font-size: 1.1rem;
  font-weight: 500;
  color: var(--green);
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.5);
}
.popularity-div span:nth-child(2) {
  color: var(--orange);
}
.trailer-div {
  margin-top: 10px;
  margin-left: 50px;
  display: flex;
  flex-direction: row;
  align-items: center;
  column-gap: 15px;
  position: relative;
  width: 200px;
}
.trailer-div span {
  font-size: 1.1rem;
  font-weight: 500;
  color: var(--green);
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.5);
}
.watch-img {
  filter: drop-shadow(5px 5px 15px rgba(0, 0, 0, 0.25));
}
.trailer-div::before {
  content: attr(data-tooltip1);
  position: absolute;
  top: 50px;
  right: 0;
  z-index: 10;
  background-color: rgb(172, 172, 172);
  padding: 5px;
  border-radius: 10px;
  opacity: 0;
  pointer-events: none;
  scale: 0;
  transition: opacity 0.25s ease, scale 0.15s ease;
}
.trailer-div:hover::before {
  opacity: 0.95;
  scale: 1;
}
.bottom-synopsis-div {
  margin-top: 50px;
  margin-left: 50px;
  margin-right: 50px;
  margin-bottom: 50px;
  padding: 20px;
  border-radius: 25px;
  background-color: var(--green);
  box-shadow: 5px 5px 5px rgba(0, 0, 0, 0.15);
  font-size: 1.2rem;
  font-weight: 500;
  color: var(--dark-blue);
}
@media (max-width: 768px) {
  .anime-main-card {
    margin-top: 50px;
    min-height: auto;
    width: 90%;
    margin-bottom: 100px;
    padding-left: 25px;
    padding-right: 25px;
    padding-top: 50px;
  }
  .top-row-div {
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }
  .poster-div {
    margin-left: 0;
    margin-top: 0;
    margin-bottom: 20px;
  }
  .title-div {
    margin-top: -10px;
    margin-left: 0;
    text-align: center;
  }
  .sub-titles {
    margin-left: 0;
    text-align: center;
  }
  .episodes-div {
    margin-left: 0;
    text-align: center;
  }
  .dates-div {
    margin-left: 0;
    text-align: center;
    display: flex;
    flex-direction: column;
  }
  .age-rating-div {
    margin-left: 0;
    text-align: center;
  }
  .anime-info-top-right-div {
    margin-top: 0;
    padding-right: 0;
  }
  .rating-div,
  .rating-rank {
    position: static;
    margin-top: 10px;
    margin-bottom: 10px;
    margin-left: auto;
    margin-right: auto;
    width: fit-content;
    right: 0;
  }
  .show-type-div {
    margin-left: 0;
    text-align: center;
  }
  .popularity-div {
    margin-left: 0;
    text-align: center;
  }
  .trailer-div {
    margin-left: auto;
    margin-right: auto;
    justify-content: center;
  }
  .bottom-synopsis-div {
    margin-top: 20px;
    margin-left: 0;
    margin-right: 0;
    margin-bottom: 20px;
  }
  .poster-div img {
    border: 5px var(--green) solid;
    transform: scale(0.9);
  }
  .trailer-div::before {
    right: -18px;
  }
  .trailer-div img {
    margin-top: 10px;
  }
}
.bottom-synopsis-div span {
  color: var(--yellow);
}
.back-div {
  position: absolute;
  top: 25px;
  left: 50px;
  z-index: 10;
}
.back-div img {
  height: 50px;
  cursor: pointer;
}
</style>

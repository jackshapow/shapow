<template>
  <article id="artistInfo" :class="mode">
    <h1 class="name">
      <span>{{ artist.name }}</span>

      <a class="shuffle" @click.prevent="shuffleAll"><i class="fa fa-random"></i></a>
    </h1>

    <div v-if="artist.info">
      <img v-if="artist.info.image" :src="artist.info.image"
        title="They see me posin, they hatin"
        class="cool-guys-posing cover">

      <div class="bio" v-if="artist.info.bio.summary">
        <div class="summary" v-show="showSummary" v-html="artist.info.bio.summary"/>
        <div class="full" v-show="showFull" v-html="artist.info.bio.full"/>

        <button class="more" v-show="showSummary" @click.prevent="showingFullBio = true">
          Full Bio
        </button>
      </div>
      <p class="none" v-else>This artist has no Last.fm biography – yet.</p>

      <footer>Data &copy; <a target="_blank" :href="artist.info.url">Last.fm</a></footer>
    </div>

    <p class="none" v-else>Nothing can be found. This artist is a mystery.</p>
  </article>
</template>

<script>
import { playback } from '../../../services'

export default {
  props: ['artist', 'mode'],

  data () {
    return {
      showingFullBio: false
    }
  },

  watch: {
    /**
     * Whenever a new artist is loaded into this component, we reset the "full bio" state.
     * @return {Boolean}
     */
    artist () {
      this.showingFullBio = false
    }
  },

  computed: {
    showSummary () {
      return this.mode !== 'full' && !this.showingFullBio
    },

    showFull () {
      return this.mode === 'full' || this.showingFullBio
    }
  },

  methods: {
    /**
     * Shuffle all songs performed by the current song's artist.
     */
    shuffleAll () {
      playback.playAllByArtist(this.artist, false)
    }
  }
}
</script>

<style lang="scss">
@import "../../../assets/sass/partials/_vars.scss";
@import "../../../assets/sass/partials/_mixins.scss";

#artistInfo {
  @include artist-album-info();
}
</style>

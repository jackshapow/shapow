<template>
  <section id="playlists">
    <h1>Playlists
      <i class="fa fa-plus-circle control create" :class="{ creating: creating }" @click="creating = !creating"/>
    </h1>

    <form v-if="creating" @submit.prevent="store" class="create">
      <input type="text"
        @keyup.esc.prevent="creating = false"
        v-model="newName"
        v-koel-focus
        placeholder="↵ to save"
        required
      >
    </form>

    <ul class="menu">
      <playlist-item type="favorites" :playlist="{ name: 'Favorites', songs: favoriteState.songs }"/>
      <playlist-item v-for="playlist in playlistState.playlists" type="playlist" :key="playlist.id" :playlist="playlist"/>
    </ul>
  </section>
</template>

<script>
import { playlistStore, favoriteStore } from '../../../stores'
import router from '../../../router'

import playlistItem from './playlist-item.vue'

export default {
  name: 'sidebar--playlists',
  props: ['currentView'],
  components: { playlistItem },

  data () {
    return {
      playlistState: playlistStore.state,
      favoriteState: favoriteStore.state,
      creating: false,
      newName: ''
    }
  },

  methods: {
    /**
     * Store/create a new playlist.
     */
    async store () {
      this.creating = false

      const playlist = await playlistStore.store(this.newName)
      this.newName = ''
      // Activate the new playlist right away
      this.$nextTick(() => router.go(`playlist/${playlist.id}`))
    }
  }
}
</script>

<style lang="scss">
@import "../../../assets/sass/partials/_vars.scss";
@import "../../../assets/sass/partials/_mixins.scss";

#playlists {
  .control.create {
    margin-top: 2px;
    font-size: 16px;
    transition: .3s;

    &.creating {
      transform: rotate(135deg);
    }
  }

  form.create {
    padding: 8px 16px;

    input[type="text"] {
      width: 100%;
    }
  }
}
</style>

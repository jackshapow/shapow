<template>
  <div id="editSongsOverlay" v-if="shown" class="overlay">
    <sound-bar v-if="loading"></sound-bar>
    <form v-else @submit.prevent="submit">
      <header>
        <img :src="coverUrl" width="96" height="96">
        <hgroup class="meta">
          <h1 :class="{ mixed: !editSingle }">{{ displayedTitle }}</h1>
          <h2 :class="{ mixed: !bySameArtist &&  !formData.artistName }">
            {{ bySameArtist || formData.artistName ? formData.artistName : 'Mixed Artists' }}
          </h2>
          <h2 :class="{ mixed: !inSameAlbum && !formData.albumName }">
            {{ inSameAlbum || formData.albumName ? formData.albumName : 'Mixed Albums' }}
          </h2>
        </hgroup>
      </header>

      <div>
        <div class="tabs tabs-white">
          <div class="header clear">
            <a @click.prevent="currentView = 'details'"
              :class="{ active: currentView === 'details' }">Details</a>
            <a @click.prevent="currentView = 'lyrics'" v-show="editSingle"
              :class="{ active: currentView === 'lyrics' }">Lyrics</a>
          </div>

          <div class="panes">
            <div v-show="currentView === 'details'">
              <div class="form-row" v-if="editSingle">
                <label>Title</label>
                <input name="title" type="text" v-model="formData.title">
              </div>
              <div class="form-row">
                <label>Artist</label>
                <typeahead
                  :items="artistState.artists"
                  :options="artistTypeaheadOptions"
                  v-model="formData.artistName"/>
              </div>
              <div class="form-row">
                <label>Album</label>
                <typeahead
                  :items="albumState.albums"
                  :options="albumTypeaheadOptions"
                  v-model="formData.albumName"/>
              </div>
              <div class="form-row">
                <label class="small">
                  <input type="checkbox" @change="changeCompilationState" ref="compilationStateChk" />
                  Album is a compilation of songs by various artists
                </label>
              </div>
              <div class="form-row" v-show="editSingle">
                <label>Track</label>
                <input name="track" type="text" pattern="\d*" v-model="formData.track"
                title="Empty or a number">
              </div>
            </div>
            <div v-show="currentView === 'lyrics' && editSingle">
              <div class="form-row">
                <textarea v-model="formData.lyrics"/>
              </div>
            </div>
          </div>
        </div>
      </div>

      <footer>
        <input type="submit" value="Update">
        <a @click.prevent="close" class="btn btn-white">Cancel</a>
      </footer>
    </form>
  </div>
</template>

<script>
import { every, filter, union } from 'lodash'

import { br2nl } from '../../utils'
import { songInfo } from '../../services/info'
import { artistStore, albumStore, songStore } from '../../stores'
import config from '../../config'

import soundBar from '../shared/sound-bar.vue'
import typeahead from '../shared/typeahead.vue'

const COMPILATION_STATES = {
  NONE: 0, // No songs belong to a compilation album
  ALL: 1, // All songs belong to compilation album(s)
  SOME: 2 // Some of the songs belong to compilation album(s)
}

export default {
  components: { soundBar, typeahead },

  data () {
    return {
      shown: false,
      songs: [],
      currentView: '',
      loading: false,

      artistState: artistStore.state,
      artistTypeaheadOptions: {
        displayKey: 'name',
        filterKey: 'name'
      },

      albumState: albumStore.state,
      albumTypeaheadOptions: {
        displayKey: 'name',
        filterKey: 'name'
      },

      /**
       * In order not to mess up the original songs, we manually assign and manipulate
       * their attributes.
       *
       * @type {Object}
       */
      formData: {
        title: '',
        albumName: '',
        artistName: '',
        lyrics: '',
        track: '',
        compilationState: null
      }
    }
  },

  computed: {
    /**
     * Determine if we're editing but one song.
     *
     * @return {boolean}
     */
    editSingle () {
      return this.songs.length === 1
    },

    /**
     * Determine if all songs we're editing are by the same artist.
     *
     * @return {boolean}
     */
    bySameArtist () {
      return every(this.songs, song => song.artist.id === this.songs[0].artist.id)
    },

    /**
     * Determine if all songs we're editing are from the same album.
     *
     * @return {boolean}
     */
    inSameAlbum () {
      return every(this.songs, song => song.album.id === this.songs[0].album.id)
    },

    /**
     * URL of the cover to display.
     *
     * @return {string}
     */
    coverUrl () {
      return this.inSameAlbum ? this.songs[0].album.cover : config.unknownCover
    },

    /**
     * Determine the compilation state of the songs.
     *
     * @return {Number}
     */
    compilationState () {
      const albums = this.songs.reduce((acc, song) => {
        return union(acc, [song.album])
      }, [])

      const compiledAlbums = filter(albums, album => album.is_compilation)

      if (!compiledAlbums.length) {
        this.formData.compilationState = COMPILATION_STATES.NONE
      } else if (compiledAlbums.length === albums.length) {
        this.formData.compilationState = COMPILATION_STATES.ALL
      } else {
        this.formData.compilationState = COMPILATION_STATES.SOME
      }

      return this.formData.compilationState
    },

    /**
     * The song title to be displayed.
     *
     * @return {string}
     */
    displayedTitle () {
      return this.editSingle ? this.formData.title : `${this.songs.length} songs selected`
    },

    /**
     * The album name to be displayed.
     *
     * @return {string}
     */
    displayedAlbum () {
      if (this.editSingle) {
        return this.formData.albumName
      } else {
        return this.formData.albumName ? this.formData.albumName : 'Mixed Albums'
      }
    },

    /**
     * The artist name to be displayed.
     *
     * @return {string}
     */
    displayedArtist () {
      if (this.editSingle) {
        return this.formData.artistName
      } else {
        return this.formData.artistName ? this.formData.artistName : 'Mixed Artists'
      }
    }
  },

  methods: {
    async open (songs) {
      this.shown = true
      this.songs = songs
      this.currentView = 'details'

      if (this.editSingle) {
        this.formData.title = this.songs[0].title
        this.formData.albumName = this.songs[0].album.name
        this.formData.artistName = this.songs[0].artist.name

        // If we're editing only one song and the song's info (including lyrics)
        // hasn't been loaded, load it now.
        if (!this.songs[0].infoRetrieved) {
          this.loading = true

          await songInfo.fetch(this.songs[0])
          this.loading = false
          this.formData.lyrics = br2nl(this.songs[0].lyrics)
          this.formData.track = this.songs[0].track || ''
          this.initCompilationStateCheckbox()
        } else {
          this.formData.lyrics = br2nl(this.songs[0].lyrics)
          this.formData.track = this.songs[0].track || ''
          this.initCompilationStateCheckbox()
        }
      } else {
        this.formData.albumName = this.inSameAlbum ? this.songs[0].album.name : ''
        this.formData.artistName = this.bySameArtist ? this.songs[0].artist.name : ''
        this.loading = false
        this.initCompilationStateCheckbox()
      }
    },

    /**
     * Initialize the compilation state's checkbox of the editing songs' album(s).
     */
    initCompilationStateCheckbox () {
      // This must be wrapped in a $nextTick callback, because the form is dynamically
      // attached into DOM in conjunction with `this.loading` data binding.
      this.$nextTick(() => {
        const chk = this.$refs.compilationStateChk

        switch (this.compilationState) {
          case COMPILATION_STATES.ALL:
            chk.checked = true
            chk.indeterminate = false
            break
          case COMPILATION_STATES.NONE:
            chk.checked = false
            chk.indeterminate = false
            break
          default:
            chk.checked = false
            chk.indeterminate = true
            break
        }
      })
    },

    /**
     * Manually set the compilation state.
     * We can't use v-model here due to the tri-state nature of the property.
     * Also, following iTunes style, we don't support circular switching of the states -
     * once the user clicks the checkbox, there's no going back to indeterminate state.
     */
    changeCompilationState (e) {
      this.formData.compilationState = e.target.checked ? COMPILATION_STATES.ALL : COMPILATION_STATES.NONE
    },

    /**
     * Close the modal.
     */
    close () {
      this.shown = false
    },

    /**
     * Submit the form.
     */
    async submit () {
      this.loading = true

      try {
        await songStore.update(this.songs, this.formData)
        this.close()
      } finally {
        this.loading = false
      }
    }
  }
}
</script>

<style lang="scss">
@import "../../assets/sass/partials/_vars.scss";
@import "../../assets/sass/partials/_mixins.scss";

#editSongsOverlay {
  form {
    > header {
      img {
        flex: 0 0 96px;
      }

      .meta {
        flex: 1;
        padding-left: 8px;

        .mixed {
          opacity: .5;
        }
      }
    }
  }
}
</style>

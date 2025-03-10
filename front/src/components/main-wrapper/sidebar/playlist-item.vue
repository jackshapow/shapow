<template>
  <li @dblclick.prevent="edit" :class="['playlist', type, editing ? 'editing' : '']">
    <a :href="playlistUrl"
      @dragleave="removeDroppableState"
      @dragenter.prevent="allowDrop"
      @dragover.prevent
      @drop.stop.prevent="handleDrop"
      :class="{ active: active }"
    >{{ playlist.name }}</a>

    <input type="text"
      @keyup.esc="cancelEdit"
      @keyup.enter="update"
      @blur="update"
      v-model="playlist.name"
      v-if="editing"
      v-koel-focus
      required
    >
  </li>
</template>

<script>
import { event, $ } from '../../../utils'
import { songStore, playlistStore, favoriteStore } from '../../../stores'

export default {
  props: ['playlist', 'type'],

  data () {
    return {
      editing: false,
      active: false
    }
  },

  computed: {
    /**
     * Determine if the current menu item is the "Favorites" one.
     *
     * @return {Boolean}
     */
    isFavorites () {
      return this.type === 'favorites'
    },

    playlistUrl () {
      return this.isFavorites ? '/#!/favorites' : `/#!/playlist/${this.playlist.id}`
    }
  },

  methods: {
    /**
     * Show the form to edit the playlist.
     */
    edit () {
      if (this.isFavorites) {
        return
      }

      this.beforeEditCache = this.playlist.name
      this.editing = true
    },

    /**
     * Update the playlist's name.
     */
    update () {
      if (this.isFavorites || !this.editing) {
        return
      }

      this.editing = false

      this.playlist.name = this.playlist.name.trim()
      if (!this.playlist.name) {
        this.playlist.name = this.beforeEditCache
        return
      }

      playlistStore.update(this.playlist)
    },

    /**
     * Cancel editing.
     */
    cancelEdit () {
      this.editing = false
      this.playlist.name = this.beforeEditCache
    },

    /**
     * Remove the droppable state when a dragleave event occurs on the playlist's DOM element.
     *
     * @param {Object} e The dragleave event.
     */
    removeDroppableState (e) {
      $.removeClass(e.target, 'droppable')
    },

    /**
     * Add a "droppable" class and set the drop effect when an item is dragged over the playlist's
     * DOM element.
     *
     * @param {Object} e The dragover event.
     */
    allowDrop (e) {
      $.addClass(e.target, 'droppable')
      e.dataTransfer.dropEffect = 'move'

      return false
    },

    /**
     * Handle songs dropped to our favorite or playlist menu item.
     *
     * @param  {Object}   e    The event
     *
     * @return {Boolean}
     */
    handleDrop (e) {
      this.removeDroppableState(e)

      if (!e.dataTransfer.getData('application/x-koel.text+plain')) {
        return false
      }

      const songs = songStore.byIds(e.dataTransfer.getData('application/x-koel.text+plain').split(','))

      if (!songs.length) {
        return false
      }

      if (this.type === 'favorites') {
        favoriteStore.like(songs)
      } else {
        playlistStore.addSongs(this.playlist, songs)
      }

      return false
    }
  },

  created () {
    event.on('main-content-view:load', (view, playlist) => {
      if (view === 'favorites') {
        this.active = this.isFavorites
      } else if (view === 'playlist') {
        this.active = this.playlist === playlist
      } else {
        this.active = false
      }
    })
  }
}
</script>

<style lang="scss" scoped>
@import "../../../assets/sass/partials/_vars.scss";
@import "../../../assets/sass/partials/_mixins.scss";

.playlist {
  user-select: none;

  a {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;

    span {
      pointer-events: none;
    }

    &::before {
      content: "\f0f6";
    }
  }

  &.favorites a::before {
    content: "\f004";
    color: $colorHeart;
  }

  input {
    width: calc(100% - 32px);
    margin: 5px 16px;
  }

  &.editing {
    a {
      display: none !important;
    }
  }
}
</style>

<template>
  <section id="mainContent">
    <div class="translucent" :style="{ backgroundImage: albumCover ? 'url('+albumCover+')' : 'none' }"></div>
    <home v-show="view === 'home'"/>
    <queue v-show="view === 'queue'"/>
    <songs v-show="view === 'songs'"/>
    <albums v-show="view === 'albums'"/>
    <album v-show="view === 'album'"/>
    <artists v-show="view === 'artists'"/>
    <artist v-show="view === 'artist'"/>
    <users v-show="view === 'users'"/>
    <uploads v-show="view === 'uploads'"/>
    <settings v-show="view === 'settings'"/>
    <playlist v-show="view === 'playlist'"/>
    <favorites v-show="view === 'favorites'"/>
    <profile v-show="view === 'profile'"/>
    <youtube-player v-if="sharedState.useYouTube" v-show="view === 'youtubePlayer'"/>
  </section>
</template>

<script>
import { event } from '../../../utils'
import { albumStore, sharedStore } from '../../../stores'

import albums from './albums.vue'
import album from './album.vue'
import artists from './artists.vue'
import artist from './artist.vue'
import songs from './songs.vue'
import uploads from './uploads.vue'
import settings from './settings.vue'
import users from './users.vue'
import queue from './queue.vue'
import home from './home.vue'
import playlist from './playlist.vue'
import favorites from './favorites.vue'
import profile from './profile.vue'
import youtubePlayer from './youtube-player.vue'

export default {
  components: { albums, album, artists, artist, songs, uploads, settings,
    users, home, queue, playlist, favorites, profile, youtubePlayer },

  data () {
    return {
      view: 'home', // The default view
      albumCover: null,
      sharedState: sharedStore.state
    }
  },

  created () {
    event.on({
      'main-content-view:load': view => {
        this.view = view
      },

      /**
       * When a new song is played, find its cover for the translucent effect.
       *
       * @param  {Object} song
       *
       * @return {Boolean}
       */
      'song:played': song => {
        this.albumCover = song.album.cover === albumStore.stub.cover ? null : song.album.cover
      }
    })
  }
}
</script>

<style lang="scss">
@import "../../../assets/sass/partials/_vars.scss";
@import "../../../assets/sass/partials/_mixins.scss";

#mainContent {
  flex: 1;
  position: relative;
  overflow: hidden;

  > section {
    position: absolute;
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    backface-visibility: hidden;

    .main-scroll-wrap {
      padding: 24px 24px 48px;
      overflow: auto;
      flex: 1;
      -ms-overflow-style: -ms-autohiding-scrollbar;

      html.touchevents & {
        // Enable scroll with momentum on touch devices
        overflow-y: scroll;
        -webkit-overflow-scrolling: touch;
      }
    }
  }

  h1.heading {
    font-weight: $fontWeight_UltraThin;
    font-size: 2.76rem;
    padding: 1rem 1.8rem;
    border-bottom: 1px solid $color2ndBgr;
    min-height: 96px;
    position: relative;
    align-items: center;
    align-content: stretch;
    display: flex;
    line-height: normal;
    background: rgba(0, 0, 0, .1);

    span:first-child {
      flex: 1;
    }

    .meta {
      display: block;
      font-size: .9rem;
      color: $color2ndText;
      margin: 6px 0 0 2px;

      a {
        color: #fff;

        &:hover {
          color: $colorHighlight;
        }
      }
    }

    .buttons {
      text-align: right;
      z-index: 2;

      @include button-group();
    }
  }

  .translucent {
    position: absolute;
    top: -20px;
    left: -20px;
    right: -20px;
    bottom: -20px;
    filter: blur(20px);
    opacity: .07;
    z-index: 2;
    overflow: hidden;
    background-size: cover;
    background-position: center;
    transform: translateZ(0);
    backface-visibility: hidden;
    perspective: 1000;
    pointer-events: none;
  }

  @media only screen and (max-width: 768px) {
    h1.heading {
      font-size: 1.38rem;
      min-height: 0;
      line-height: 1.85rem;
      text-align: center;
      flex-direction: column;

      .meta {
        display: none;
      }

      .buttons {
        justify-content: center;
        margin-top: 8px;
      }

      span:first-child {
        flex: 0 0 28px;
      }
    }

    > section {
      .main-scroll-wrap {
        padding: 12px;
      }
    }
  }

  @media only screen and (max-width: 375px) {
    > section {
      // Leave some space for the "Back to top" button
      .main-scroll-wrap {
        padding-bottom: 64px;
      }
    }
  }
}
</style>

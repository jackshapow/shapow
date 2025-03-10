<template>
  <div id="app">
    <div id="main" tabindex="0" v-show="authenticated"
      @keydown.space="togglePlayback"
      @keydown.j = "playNext"
      @keydown.k = "playPrev"
      @keydown.f = "search"
      @keydown.mediaPrev = "playPrev"
      @keydown.mediaNext = "playNext"
      @keydown.mediaToggle = "togglePlayback"
    >
      <site-header/>
      <main-wrapper/>
      <site-footer/>
      <overlay ref="overlay"/>
      <edit-songs-form ref="editSongsForm"/>
    </div>

    <div class="login-wrapper" v-if="!authenticated">
      <login-form/>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'

import siteHeader from './components/site-header/index.vue'
import siteFooter from './components/site-footer/index.vue'
import mainWrapper from './components/main-wrapper/index.vue'
import overlay from './components/shared/overlay.vue'
import loginForm from './components/auth/login-form.vue'
import editSongsForm from './components/modals/edit-songs-form.vue'

import { event, showOverlay, hideOverlay, forceReloadWindow, $ } from './utils'
import { sharedStore, userStore, preferenceStore as preferences } from './stores'
import store from './store'
import { playback, ls } from './services'
import { focusDirective, clickawayDirective } from './directives'
import router from './router'

// For the future...
//import VueNativeSock from 'vue-native-websocket'
//Vue.use(VueNativeSock, 'ws://localhost:31337/ws', { store: store, format: 'json', reconnection: true, reconnectionAttempts: 5, reconnectionDelay: 3000})

export default {
  components: { siteHeader, siteFooter, mainWrapper, overlay, loginForm, editSongsForm },

  data () {
    return {
      authenticated: false
    }
  },

  mounted () {
    // The app has just been initialized, check if we can get the user data with an already existing token
    const token = ls.get('jwt-token')
    if (token) {
      this.authenticated = true
      this.init()
    }

    // Create the element to be the ghost drag image.
    const dragGhost = document.createElement('div')
    dragGhost.id = 'dragGhost'
    document.body.appendChild(dragGhost)

    // And the textarea to copy stuff
    const copyArea = document.createElement('textarea')
    copyArea.id = 'copyArea'
    document.body.appendChild(copyArea)

    // Add an ugly mac/non-mac class for OS-targeting styles.
    // I'm crying inside.
    $.addClass(document.documentElement, navigator.userAgent.indexOf('Mac') !== -1 ? 'mac' : 'non-mac')
  },

  methods: {
    async init () {
      showOverlay()

      // Make the most important HTTP request to get all necessary data from the server.
      // Afterwards, init all mandatory stores and services.
      try {
        await sharedStore.init()
        playback.init()
        hideOverlay()

        // Ask for user's notification permission.
        this.requestNotifPermission()

        // To confirm or not to confirm closing, it's a question.
        window.onbeforeunload = e => {
          if (!preferences.confirmClosing) {
            return
          }

          // Notice that a custom message like this has ceased to be supported
          // starting from Chrome 51.
          return 'You asked to confirm before closing, so here it is.'
        }

        // Let all other components know we're ready.
        event.emit('koel:ready')
      } catch (err) {
        this.authenticated = false
      }
    },

    /**
     * Toggle playback when user presses Space key.
     *
     * @param {Object} e The keydown event
     */
    togglePlayback (e) {
      if ($.is(e.target, 'input,textarea,button,select')) {
        return true
      }

      // Whatever play/pause control is there, we blindly click it.
      const play = document.querySelector('#mainFooter .play')
      play ? play.click() : document.querySelector('#mainFooter .pause').click()

      e.preventDefault()
    },

    /**
     * Play the previous song when user presses K.
     *
     * @param {Object} e The keydown event
     */
    playPrev (e) {
      if ($.is(e.target, 'input,textarea')) {
        return true
      }

      playback.playPrev()
      e.preventDefault()
    },

    /**
     * Play the next song when user presses J.
     *
     * @param {Object} e The keydown event
     */
    playNext (e) {
      if ($.is(e.target, 'input,textarea')) {
        return true
      }

      playback.playNext()
      e.preventDefault()
    },

    /**
     * Put focus into the search field when user presses F.
     *
     * @param {Object} e The keydown event
     */
    search (e) {
      if ($.is(e.target, 'input,textarea') || e.metaKey || e.ctrlKey) {
        return true
      }

      const selectBox = document.querySelector('#searchForm input[type="search"]')
      selectBox.focus()
      selectBox.select()
      e.preventDefault()
    },

    /**
     * Request for notification permission if it's not provided and the user is OK with notifs.
     */
    requestNotifPermission () {
      if (window.Notification && preferences.notify && window.Notification.permission !== 'granted') {
        window.Notification.requestPermission(result => {
          if (result === 'denied') {
            preferences.notify = false
          }
        })
      }
    }

  },

  // sockets: {
  //   addsong: function (song) {
  //     //this.users.push(user);
  //     alert("add a song!")
  //   }
  // },

  created () {
    event.on({
      /**
       * When the user logs in, set the whole app to be "authenticated" and initialize it.
       */
      'user:loggedin': () => {
        this.authenticated = true
        this.init()
      },

      /**
       * Shows the "Edit Song" form.
       *
       * @param {Array.<Object>} An array of songs to edit
       */
      'songs:edit': songs => this.$refs.editSongsForm.open(songs),

      /**
       * Log the current user out and reset the application state.
       */
      async logout () {
        await userStore.logout()
        ls.remove('jwt-token')
        forceReloadWindow()
      },

      /**
       * Init our basic, custom router on ready to determine app state.
       */
      'koel:ready': () => {
        router.init()
      }
    })
  }
}

// Register our custom key codes
Vue.config.keyCodes = {
  a: 65,
  j: 74,
  k: 75,
  f: 70,
  mediaNext: 176,
  mediaPrev: 177,
  mediaToggle: 179
}

// …and the global directives
Vue.directive('koel-focus', focusDirective)
Vue.directive('koel-clickaway', clickawayDirective)
</script>

<style lang="scss">
@import "assets/sass/app.scss";
@import "assets/sass/partials/_vars.scss";
@import "assets/sass/partials/_mixins.scss";
@import "assets/sass/partials/_shared.scss";

#dragGhost {
  position: absolute;
  display: inline-block;
  background: $colorGreen;
  padding: .8rem;
  border-radius: .2rem;
  color: #fff;
  font-family: $fontFamily;
  font-size: 1rem;
  font-weight: $fontWeight_Thin;
  top: -100px;
  left: 0px;

  /**
   * We can totally hide this element on touch devices, because there's
   * no drag and drop support there anyway.
   */
  html.touchevents & {
    display: none;
  }
}

#copyArea {
  position: absolute;
  left: -9999px;
  width: 1px;
  height: 1px;
  bottom: 1px;

  html.touchevents & {
    display: none;
  }
}

#main, .login-wrapper {
  display: flex;
  min-height: 100vh;
  flex-direction: column;

  background: $colorMainBgr;
  color: $colorMainText;

  font-family: $fontFamily;
  font-size: 1rem;
  line-height: 1.5rem;
  font-weight: $fontWeight_Thin;

  padding-bottom: $footerHeight;
}

.login-wrapper {
  @include vertical-center();

  padding-bottom: 0;
}
</style>

<template>
  <header id="mainHeader">
    <h1 class="brand" v-once>{{ appTitle }}</h1>
    <span class="hamburger" @click="toggleSidebar">
      <i class="fa fa-bars"></i>
    </span>
    <span class="magnifier" @click="toggleSearchForm">
      <i class="fa fa-search"></i>
    </span>
    <search-form/>
    <user-badge/>
  </header>

</template>

<script>
import config from '../../config'
import { event } from '../../utils'
import searchForm from './search-form.vue'
import userBadge from './user-badge.vue'

export default {
  components: { searchForm, userBadge },

  data () {
    return {
      appTitle: config.appTitle
    }
  },

  methods: {
    /**
     * No I'm not documenting this.
     */
    toggleSidebar () {
      event.emit('sidebar:toggle')
    },

    /**
     * or this.
     */
    toggleSearchForm () {
      event.emit('search:toggle')
    }
  }
}
</script>

<style lang="scss">
@import "../../assets/sass/partials/_vars.scss";
@import "../../assets/sass/partials/_mixins.scss";

#mainHeader {
  height: $headerHeight;
  background: $color2ndBgr;
  display: flex;
  border-bottom: 1px solid $colorMainBgr;

  h1.brand {
    flex: 1;
    color: $colorMainText;
    font-size: 1.7rem;
    font-weight: $fontWeight_UltraThin;
    opacity: 0;
    line-height: $headerHeight;
    text-align: center;
  }

  .hamburger, .magnifier {
    font-size: 1.4rem;
    flex: 0 0 48px;
    order: -1;
    line-height: $headerHeight;
    text-align: center;
    display: none;
  }

  @media only screen and (max-width: 667px) {
    display: flex;
    align-content: stretch;
    justify-content: flext-start;

    .hamburger, .magnifier {
      display: inline-block;
    }

    h1.brand {
      opacity: 1;
    }
  }
}
</style>

<template>
  <div id="mainWrapper">
    <sidebar/>
    <main-content/>
    <extra/>
  </div>
</template>

<script>
import sidebar from './sidebar/index.vue'
import mainContent from './main-content/index.vue'
import extra from './extra/index.vue'

export default {
  components: { sidebar, mainContent, extra }
}
</script>

<style lang="scss">
@import "../../assets/sass/partials/_vars.scss";
@import "../../assets/sass/partials/_mixins.scss";

#mainWrapper {
  display: flex;
  flex: 1;
}
</style>

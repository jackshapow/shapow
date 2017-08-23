// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import './static-loader'
import Vue from 'vue'
import App from './App'
import { event } from './utils'
import { http } from './services'
import { VirtualScroller } from 'vue-virtual-scroller/dist/vue-virtual-scroller'
Vue.component('virtual-scroller', VirtualScroller)


Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  template: '<App/>',
  components: { App },
  created () {
    event.init()
    http.init()
  }

})

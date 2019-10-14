'use strict'

import Vue from 'vue'
import router from '@/router'
import store from '@/store'
import App from '@/App.vue'
import '@/libraries/ant'

// Prevent the tip on Vue startup
Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})

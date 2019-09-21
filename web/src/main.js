'use strict'

import Vue from 'vue'
import router from '@/router'
import App from '@/App.vue'

// Prevent the tip on Vue startup
Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})

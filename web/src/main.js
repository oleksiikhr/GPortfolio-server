'use strict'

import Vue from 'vue'
import router from '@/router'
import App from '@/App.vue'

// Prevent the tip on Vue startup
Vue.config.productionTip = false

// Load ant styles (edit in webpack.config.js)
import "ant-design-vue/dist/antd.less";

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})

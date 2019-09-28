'use strict'

import Vue from 'vue'
import router from '@/router'
import store from '@/store'
import App from '@/App.vue'

import AStep from 'ant-design-vue/lib/vc-steps/Step'
import APopover from 'ant-design-vue/lib/popover'
import AButton from 'ant-design-vue/lib/button'
import AAvatar from 'ant-design-vue/lib/avatar'
import ASteps from 'ant-design-vue/lib/steps'
import AIcon from 'ant-design-vue/lib/icon'

Vue.component('APopover', APopover)
Vue.component('AAvatar', AAvatar)
Vue.component('AButton', AButton)
Vue.component('ASteps', ASteps)
Vue.component('AStep', AStep)
Vue.component('AIcon', AIcon)

// Prevent the tip on Vue startup
Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})

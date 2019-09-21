'use strict'

import VueRouter from 'vue-router'
import Vue from 'vue'
import routes from '@/router/routes'

Vue.use(VueRouter)

export default new VueRouter({
  // eslint-disable-next-line
  mode: process.env.NODE_ENV === 'production' ? 'history' : 'hash',
  scrollBehavior: () => ({ y: 0 }),
  routes
})

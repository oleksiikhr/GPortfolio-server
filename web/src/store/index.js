'use strict'

import modules from '@/store/modules'
import Vuex from 'vuex'
import Vue from 'vue'

Vue.use(Vuex)

export default new Vuex.Store({
  // eslint-disable-next-line
  strict: process.env.NODE_ENV !== 'production',
  modules
})

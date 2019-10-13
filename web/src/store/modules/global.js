'use strict'

import config from '@/data/config'

const state = {
  ...config.global
}

const mutations = {
  UPDATE(state, { key, value }) {
    state[key] = value
  }
}

const actions = {
  //
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}


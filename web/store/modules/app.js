'use strict'

const state = {
  step: 0
}

const mutations = {
  SET_STEP(state, toggle) {
    state.step = toggle
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

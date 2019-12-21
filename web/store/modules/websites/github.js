'use strict'

const state = {
  profile: {},
  repositories: []
}

const mutations = {
  SET_PROFILE(state, obj) {
    state.profile = obj
  },
  SET_REPOSITORIES(state, arr) {
    state.repositories = arr
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

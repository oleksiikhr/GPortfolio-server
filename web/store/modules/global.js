'use strict'

const state = {
  template: '',
  base: '',
  opg: {},
  pwa: {},
  domain: ''
}

const mutations = {
  SET_TEMPLATE(state, str) {
    state.template = str
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

import Vue from 'vue'
import types from './types'

const state = {
  Users: []
}

const actions = {
  [types.actions.GET_USERS] (store, vm) {
    let loading = vm.$loading.open()

    Vue.http.get('api/users').then(response => {
      store.commit(types.mutations.LOAD_USERS, response.body)
      loading.close()
    }, failed => {
      store.commit(types.mutations.LOAD_USERS, []) // reset on failure
      loading.close()
    })
  }
}

const mutations = {
  [types.mutations.LOAD_USERS] (state, users) {
    state.Users = users
  }
}

const getters = {
  [types.getters.USERS] (state) {
    return state.Users
  }
}

export default {
  state,
  actions,
  mutations,
  getters
}

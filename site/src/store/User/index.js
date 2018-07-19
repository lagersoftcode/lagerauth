import Vue from 'vue'
import types from './types'

const state = {
  Users: [],
  User: {}
}

const actions = {
  [types.actions.GET_USERS] (store, vm) {
    let loading = vm.$loading.open()

    Vue.http.get('api/users').then(response => {
      store.commit(types.mutations.LOAD_USERS, response.body)
      loading.close()
    }, failed => {
      vm.$snackbar.open({
        message: `error getting users: ${failed.body}`,
        type: 'is-danger'
      })
      store.commit(types.mutations.LOAD_USERS, []) // reset on failure
      loading.close()
    })
  },
  [types.actions.GET_USER] (store, {vm, id}) {
    return new Promise((resolve, reject) => {
      let loading = vm.$loading.open()

      if (!id) {
        store.commit(types.mutations.LOAD_USER, {})
        loading.close()
        return
      }

      Vue.http.get(`api/users/${id}`).then(response => {
        store.commit(types.mutations.LOAD_USER, response.body)
        resolve()
      }, failed => {
        vm.$snackbar.open({
          message: `error getting user: ${failed.body}`,
          type: 'is-danger'
        })
        store.commit(types.mutations.LOAD_USER, {})
      })

      loading.close()
    })
  },
  [types.actions.DELETE_USER] (store, {vm, item}) {
    let loading = vm.$loading.open()

    Vue.http.delete(`api/users/${item.id}`).then(response => {
      store.dispatch(types.actions.GET_USERS, vm)
    }, failed => {
      vm.$snackbar.open({
        message: `error deleting user: ${failed.body}`,
        type: 'is-danger'
      })
    })

    loading.close()
  },
  [types.actions.CREATE_USER] (store, {vm, item}) {
    let loading = vm.$loading.open()

    Vue.http.post(`api/users`, item).then(response => {
      vm.$router.push('/users')
    }, failed => {
      vm.$snackbar.open({
        message: `error creating user: ${failed.body}`,
        type: 'is-danger'
      })
    })

    loading.close()
  },
  [types.actions.UPDATE_USER] (store, {vm, item}) {
    let loading = vm.$loading.open()

    Vue.http.put(`api/users/${item.id}`, item).then(response => {
      vm.$router.push('/users')
    }, failed => {
      vm.$snackbar.open({
        message: `error updating user: ${failed.body}`,
        type: 'is-danger'
      })
    })

    loading.close()
  }
}

const mutations = {
  [types.mutations.LOAD_USERS] (state, users) {
    state.Users = users
  },
  [types.mutations.LOAD_USER] (state, user) {
    state.User = user
  }
}

const getters = {
  [types.getters.USERS] (state) {
    return state.Users
  },
  [types.getters.USER] (state) {
    return state.User
  }
}

export default {
  state,
  actions,
  mutations,
  getters
}

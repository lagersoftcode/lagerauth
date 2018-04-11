import Vue from 'vue'
import types from './types'

const state = {
  Roles: [],
  Role: {}
}

const actions = {
  [types.actions.GET_ROLES] (store, vm) {
    let loading = vm.$loading.open()

    Vue.http.get('api/roles').then(response => {
      store.commit(types.mutations.LOAD_ROLES, response.body)
      loading.close()
    }, failed => {
      vm.$snackbar.open({
        message: `error getting roles: ${failed.body}`,
        type: 'is-danger'
      })
      store.commit(types.mutations.LOAD_ROLES, []) // reset on failure
      loading.close()
    })
  },
  [types.actions.GET_ROLE] (store, {vm, id}) {
    return new Promise((resolve, reject) => {
      let loading = vm.$loading.open()

      if (!id) {
        store.commit(types.mutations.LOAD_ROLE, {})
        loading.close()
        return
      }

      Vue.http.get(`api/roles/${id}`).then(response => {
        store.commit(types.mutations.LOAD_ROLE, response.body)
        resolve()
      }, failed => {
        vm.$snackbar.open({
          message: `error getting role: ${failed.body}`,
          type: 'is-danger'
        })
        store.commit(types.mutations.LOAD_ROLE, {})
      })

      loading.close()
    })
  },
  [types.actions.DELETE_ROLE] (store, {vm, item}) {
    let loading = vm.$loading.open()

    Vue.http.delete(`api/roles/${item.id}`).then(response => {
      store.dispatch(types.actions.GET_ROLES, vm)
    }, failed => {
      vm.$snackbar.open({
        message: `error deleting role: ${failed.body}`,
        type: 'is-danger'
      })
    })

    loading.close()
  },
  [types.actions.CREATE_ROLE] (store, {vm, item}) {
    let loading = vm.$loading.open()

    Vue.http.post(`api/roles`, item).then(response => {
      vm.$router.push('/roles')
    }, failed => {
      vm.$snackbar.open({
        message: `error creating role: ${failed.body}`,
        type: 'is-danger'
      })
    })

    loading.close()
  },
  [types.actions.UPDATE_ROLE] (store, {vm, item}) {
    let loading = vm.$loading.open()

    Vue.http.put(`api/roles/${item.id}`, item).then(response => {
      vm.$router.push('/roles')
    }, failed => {
      vm.$snackbar.open({
        message: `error updating role: ${failed.body}`,
        type: 'is-danger'
      })
    })

    loading.close()
  }
}

const mutations = {
  [types.mutations.LOAD_ROLES] (state, roles) {
    state.Roles = roles
  },
  [types.mutations.LOAD_ROLE] (state, role) {
    state.Role = role
  }
}

const getters = {
  [types.getters.ROLES] (state) {
    return state.Roles
  },
  [types.getters.ROLE] (state) {
    return state.Role
  }
}

export default {
  state,
  actions,
  mutations,
  getters
}

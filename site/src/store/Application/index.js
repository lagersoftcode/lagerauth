import Vue from 'vue'
import types from './types'

const state = {
  Applications: [],
  Application: {}
}

const actions = {
  [types.actions.GET_APPLICATIONS] (store, vm) {
    let loading = vm.$loading.open()

    Vue.http.get('api/applications').then(response => {
      store.commit(types.mutations.LOAD_APPLICATIONS, response.body)
    }, failed => {
      vm.$snackbar.open({
        message: `error getting applications: ${failed.body}`,
        type: 'is-danger'
      })
      store.commit(types.mutations.LOAD_APPLICATIONS, []) // reset on failure
    })

    loading.close()
  },
  [types.actions.GET_APPLICATION] (store, {vm, id}) {
    return new Promise((resolve, reject) => {
      let loading = vm.$loading.open()

      if (!id) {
        store.commit(types.mutations.LOAD_APPLICATION, {})
        loading.close()
        return
      }

      Vue.http.get(`api/applications/${id}`).then(response => {
        store.commit(types.mutations.LOAD_APPLICATION, response.body)
        resolve()
      }, failed => {
        vm.$snackbar.open({
          message: `error getting application: ${failed.body}`,
          type: 'is-danger'
        })
        store.commit(types.mutations.LOAD_APPLICATION, {})
      })

      loading.close()
    })
  },
  [types.actions.DELETE_APPLICATION] (store, {vm, item}) {
    let loading = vm.$loading.open()

    Vue.http.delete(`api/applications/${item.id}`).then(response => {
      store.dispatch(types.actions.GET_APPLICATIONS, vm)
    }, failed => {
      vm.$snackbar.open({
        message: `error deleting application: ${failed.body}`,
        type: 'is-danger'
      })
    })

    loading.close()
  },
  [types.actions.CREATE_APPLICATION] (store, {vm, item}) {
    let loading = vm.$loading.open()

    Vue.http.post(`api/applications`, item).then(response => {
      vm.$router.push('/applications')
    }, failed => {
      vm.$snackbar.open({
        message: `error creating application: ${failed.body}`,
        type: 'is-danger'
      })
    })

    loading.close()
  },
  [types.actions.UPDATE_APPLICATION] (store, {vm, item}) {
    let loading = vm.$loading.open()

    Vue.http.put(`api/applications/${item.id}`, item).then(response => {
      vm.$router.push('/applications')
    }, failed => {
      vm.$snackbar.open({
        message: `error updating application: ${failed.body}`,
        type: 'is-danger'
      })
    })

    loading.close()
  }
}

const mutations = {
  [types.mutations.LOAD_APPLICATIONS] (state, applications) {
    state.Applications = applications
  },
  [types.mutations.LOAD_APPLICATION] (state, application) {
    state.Application = application
  }
}

const getters = {
  [types.getters.APPLICATIONS] (state) {
    return state.Applications
  },
  [types.getters.APPLICATION] (state) {
    return state.Application
  }
}

export default {
  state,
  actions,
  mutations,
  getters
}

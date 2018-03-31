import types from './types'
import router from '../../router'
import Vue from 'vue'
import jwtDecode from 'jwt-decode'
import {loadAuthInterceptor} from '../../resource'

const state = {
  isLoggedIn: false,
  email: ''
}

const actions = {
  [types.actions.LOG_IN] (store, vm) {
    var data = new FormData()
    data.append('code', vm.$route.query.code)
    data.append('client_id', process.env.OAUTH_CLIENT_ID)

    var loading = vm.$loading.open()
    Vue.http.post('token', data).then(response => {
      loading.close()

      if (response.data.access_token) {
        localStorage.setItem('access_token', response.data.access_token)
        store.commit(types.mutations.LOG_IN, response.data.access_token) // success
        vm.$router.push('/')
      }
    })
  },
  [types.actions.LOG_OUT] (store) {
    localStorage.removeItem('access_token')
    store.commit(types.mutations.LOG_OUT)
    router.push('/login')
  },
  [types.actions.REDIRECT_OATH] (store) {
    var redirectUri = `${location.origin}/Token`
    var oauthUri = `${process.env.OAUTH_LOGIN}?client_id=${process.env.OAUTH_CLIENT_ID}&redirect_uri=${redirectUri}`
    location.href = oauthUri // redirect to oauth
  }
}

const mutations = {
  [types.mutations.LOG_IN] (state, token) {
    var decoded = jwtDecode(token)
    state.isLoggedIn = true
    state.email = decoded.email

    loadAuthInterceptor()
  },
  [types.mutations.LOG_OUT] (state) {
    state.isLoggedIn = false
    state.email = ''
  }
}

const getters = {
  [types.getters.IS_LOGGED_IN] (state) {
    return state.isLoggedIn
  },
  [types.getters.EMAIL] (state) {
    return state.email
  }
}

// Main export
export default {
  state,
  actions,
  mutations,
  getters
}

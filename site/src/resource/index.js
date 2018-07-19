import Vue from 'vue'
import VueResource from 'vue-resource'
import types from '../store/types'

Vue.use(VueResource)

Vue.http.options.root = process.env.API_URL
window.interceptorLoaded = false

export const loadAuthInterceptor = () => {
  var token = localStorage.getItem('access_token')
  if (token != null && window.interceptorLoaded === false) {
    Vue.http.interceptors.push((request, next) => {
      request.headers.set('Authorization', `Bearer ${token}`)
      request.headers.set('Accept', 'application/json')
      window.interceptorLoaded = true
      next(res => {
        if (res.status === 401 || res.status === 403) {
          window.vm.$store.dispatch(types.Login.actions.LOG_OUT)
        }
      })
    })
  }
}

export default VueResource

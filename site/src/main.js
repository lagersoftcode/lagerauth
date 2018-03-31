// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import App from './App'
import {default as router, authorizeRoute} from './router'
import store from './store'
import './resource'
import types from './store/types'

Vue.use(Buefy, {defaultIconPack: 'fa'})
Vue.config.productionTip = false

/* eslint-disable no-new */
window.vm = new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  components: {
    App
  },
  created () {
    var token = localStorage.getItem('access_token')
    if (token === null) {
      if (this.$route.name !== 'Token') {
        this.$router.push({name: 'Login'})
      }
    } else {
      this.$store.commit(types.Login.mutations.LOG_IN, token)
    }
  }
})

router.beforeEach(authorizeRoute(window.vm))

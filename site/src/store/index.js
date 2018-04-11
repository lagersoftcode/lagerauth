import Vue from 'vue'
import Vuex from 'vuex'
import Login from './Login'
import Application from './Application'
import User from './User'
import Role from './Role'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    Login,
    Application,
    User,
    Role
  }
})


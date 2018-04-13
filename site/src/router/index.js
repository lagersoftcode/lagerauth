import Vue from 'vue'
import Router from 'vue-router'
import Home from '../components/Home/Home'
import Login from '../components/Login/Login'
import Token from '../components/Login/Token'
import Applications from '../components/Applications/Applications'
import Application from '../components/Applications/Application'
import Roles from '../components/Roles/Roles'
import Role from '../components/Roles/Role'
import Users from '../components/Users/Users'
import User from '../components/Users/User'

Vue.use(Router)

export default new Router({
  routes: [
    { path: '/', name: 'Home', component: Home },

    { path: '/login', name: 'Login', component: Login },
    { path: '/token', name: 'Token', component: Token },

    { path: '/applications', name: 'Applications', component: Applications },
    { path: '/application/:id?', name: 'Application', component: Application },

    { path: '/roles', name: 'Roles', component: Roles },
    { path: '/role/:id?', name: 'Role', component: Role },

    { path: '/users', name: 'Users', component: Users },
    { path: '/user/:id?', name: 'User', component: User }
  ],
  mode: 'history'
})

export function authorizeRoute (vm) {
  return (to, from, next) => {
    if (to.name !== 'Login' && (!vm.$store.state.Login.isLoggedIn)) {
      vm.$router.push({name: 'Login', query: { return: to.fullPath }})
    }
    next()
  }
}

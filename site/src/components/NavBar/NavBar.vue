<template>
<nav v-if="isLoggedIn" class="navbar">
  <div class="navbar-brand">
    <a class="navbar-item" @click="$router.push('/')">
      <img src="/static/logo-alt.png" width="112" height="28">
    </a>
    <div id="navBurger" class="navbar-burger" @click="toggleMenu">
      <span></span>
      <span></span>
      <span></span>
    </div>
  </div>

  
  <div id="navMenu" class="navbar-menu">
    <div class="navbar-start">
      <a class="navbar-item" @click="goTo('/applications')">applications</a>
      <a class="navbar-item" @click="goTo('/users')">users</a>
      <a class="navbar-item" @click="goTo('/roles')">roles</a>
    </div>

    <div class="navbar-end">
      <div class="navbar-item">
        <div class="field is-grouped">
          <p class="control">
            <a class="button is-primary" @click="logout">
              <span>{{email}}</span>
              <span class="icon"><i class="fa fa-sign-out"></i></span>
              <span>logout</span>
            </a>
          </p>
        </div>
      </div>
    </div>    
  </div>
</nav>
</template>


<script>
import {mapActions, mapGetters} from 'vuex'
import types from '../../store/types'

export default {
  name: 'NavBar',
  data () { return {} },
  computed: {
    ...mapGetters({
      isLoggedIn: types.Login.getters.IS_LOGGED_IN,
      email: types.Login.getters.EMAIL
    })
  },
  methods: {
    toggleMenu (event) {
      let burger = event.target
      let menu = document.querySelector('.navbar-menu')

      burger.classList.toggle('is-active')
      menu.classList.toggle('is-active')
    },
    goTo (where) {
      let burger = document.querySelector('.navbar-burger')
      let menu = document.querySelector('.navbar-menu')

      burger.classList.remove('is-active')
      menu.classList.remove('is-active')
      this.$router.push(where)
    },
    ...mapActions({logout: types.Login.actions.LOG_OUT})
  }
}
</script>
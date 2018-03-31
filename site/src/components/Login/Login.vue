<template>
  <div class="hero is-primary is-fullheight">
    <div class="hero-body">
      <div class="container has-text-centered">
        <h1 class="title">you are not logged in</h1>
        <a class="button" @click="redirectOAuth">
          <span class="icon">
            <i class="fa fa-beer"></i>
          </span>
          <span>log in with your lagersoft account</span>
        </a>
      </div>
    </div>
    
  </div>
  
</template>

<script>
import types from '../../store/types'

export default {
  name: 'Login',
  methods: {
    redirectOAuth () {
      this.$store.dispatch(types.Login.actions.REDIRECT_OATH)
    }
  }
}
</script>
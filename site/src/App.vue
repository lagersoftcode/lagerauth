<template>
  <div id="app">
    <nav-bar />
    <router-view />
  </div>
</template>

<script>
import NavBar from './components/NavBar/NavBar'

export default {
  name: 'app',
  components: {
    NavBar
  }
}
</script>

<style>
@import '~font-awesome/css/font-awesome.min.css';
@import '~bulma/css/bulma.css';

.input,
.dropdown,
.dropdown-toggle {
  width: 100%;
}

.v-select .dropdown-menu > .highlight > a {
  background-color: #00d1b2;
}
</style>
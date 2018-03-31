<template>
</template>

<script>
import types from '../../store/types'

export default {
  name: 'Token',
  created () {
    this.$store.dispatch(types.Login.actions.LOG_IN, this)
  }
}
</script>
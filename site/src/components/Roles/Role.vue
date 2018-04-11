<template>
  <section class="section">
    <div class="container">


        <h1 v-if="this.isNew" class="title" >create role</h1>
        <h1 v-else class="title">edit role</h1>
        <hr />

        <div class="columns">
          <div class="column">
            <form @submit.prevent="save">

              <b-field label="application" >
                <v-select 
                v-model="application" 
                :options="applications"
                value = "id"
                label = "name"
                required
                />
              </b-field>           
              <b-field label="name">
                  <b-input v-model="name" required></b-input>
              </b-field>

              <b-field label="description">
                <b-input type="textarea" v-model="description"></b-input>
              </b-field>

              <button class="button" @click.prevent="$router.push('/roles')">cancel</button>
              <input type="submit" :value="this.isNew ? 'create' : 'update'" class="button is-success">
            </form>
          </div>
          <div class="column">
            <b-field label="permissions">
              <permissions :data.sync="permissions" />
            </b-field>
          </div>
        </div>
    </div>
  </section>
</template>

<script>
    import VSelect from 'vue-select'
    import Permissions from './Permissions'
    import types from '../../store/Role/types'
    import appTypes from '../../store/Application/types'
    import {mapGetters} from 'vuex'

    export default {
      name: 'Role',
      data () {
        return {
          id: null,
          name: '',
          description: '',
          application: null,
          permissions: []
        }
      },
      created () {
        if (this.applications.length === 0) {
          this.$store.dispatch(appTypes.actions.GET_APPLICATIONS, this).then(() => {
            this.load()
          })
        } else {
          this.load()
        }
      },
      computed: {
        ...mapGetters({
          role: types.getters.ROLE,
          applications: appTypes.getters.APPLICATIONS
        }),
        isNew () {
          return this.id === null
        }
      },
      methods: {
        load () {
          this.$store.dispatch(types.actions.GET_ROLE, {vm: this, id: this.$route.params.id}).then(() => {
            this.id = this.role.id || null
            this.name = this.role.name || ''
            this.description = this.role.description || ''
            this.application = this.role.application || {}
            this.permissions = this.role.permissions || []
          })
        },
        save () {
          if (this.application === null) {
            this.$snackbar.open({
              message: 'application cannot be blank',
              type: 'is-danger'
            })
            return
          }
          let item = {id: this.id, name: this.name, description: this.description, application: this.application, permissions: this.permissions}
          if (item.id) {
            this.$store.dispatch(types.actions.UPDATE_ROLE, {vm: this, item})
          } else {
            this.$store.dispatch(types.actions.CREATE_ROLE, {vm: this, item})
          }
        }
      },
      components: {
        Permissions,
        VSelect
      }
    }
</script>

<style>
  .input,
  .dropdown,
  .dropdown-toggle {
    width: 100%;
  }

  .v-select .dropdown-menu > .highlight > a {
    background-color: #00d1b2;
  }
</style>

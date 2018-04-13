<template>
  <section class="section">
    <div class="container">

        <h1 v-if="this.isNew" class="title" >create user</h1>
        <h1 v-else class="title">edit user</h1>
        <hr />

        <div class="columns">
          <div class="column">
            <form @submit.prevent="save">
              <b-field label="email">
                <b-input v-model="email" type="email" required></b-input>
              </b-field>

              <b-field label="name">
                  <b-input v-model="name" required></b-input>
              </b-field>

              <b-field label="department">
                <b-input v-model="department" required></b-input>
              </b-field>

              <b-field label="enabled">
                <input type="checkbox" v-model="enabled" />
              </b-field>                      

              <b-field label="description">
                <b-input type="textarea" v-model="description"></b-input>
              </b-field>

              <button class="button" @click.prevent="$router.push('/users')">cancel</button>
              <input type="submit" :value="this.isNew ? 'create' : 'update'" class="button is-success">
            </form>
          </div>

          <div class="column">
            <b-field label="applications">
              <v-select 
                v-model="applications" 
                :options="allApplications"
                value = "id"
                label = "name"
                multiple
              />      
            </b-field>

            <b-field label="roles">
              <v-select 
                v-model="roles" 
                :options="allRoles"
                value = "id"
                label = "name"
                multiple
              />      
            </b-field>
          </div>
        </div>
    </div>
  </section>
</template>

<script>
    import VSelect from 'vue-select'
    import types from '../../store/User/types'
    import appTypes from '../../store/Application/types'
    import roleTypes from '../../store/Role/types'
    import {mapGetters} from 'vuex'

    export default {
      name: 'User',
      data () {
        return {
          id: null,
          email: '',
          name: '',
          department: '',
          description: '',
          enabled: false,

          applications: [],
          roles: []
        }
      },
      created () {
        this.$store.dispatch(appTypes.actions.GET_APPLICATIONS, this)
        .then(() => { this.$store.dispatch(roleTypes.actions.GET_ROLES, this) })
        .then(() => { this.load() })
      },
      computed: {
        ...mapGetters({
          allApplications: appTypes.getters.APPLICATIONS,
          allRoles: roleTypes.getters.ROLES,
          user: types.getters.USER
        }),
        isNew () {
          return this.id === null
        }
      },
      methods: {
        load () {
          this.$store.dispatch(types.actions.GET_USER, {vm: this, id: this.$route.params.id}).then(() => {
            this.id = this.user.id || null
            this.email = this.user.email || ''
            this.name = this.user.name || ''
            this.department = this.user.department || ''
            this.description = this.user.description || ''
            this.enabled = this.user.enabled || false

            this.applications = this.user.applications || []
            this.roles = this.user.roles || []
          })
        },
        save () {
          let item = {id: this.id, email: this.email, name: this.name, department: this.department, description: this.description, enabled: this.enabled, applications: this.applications, roles: this.roles}
          if (item.id) {
            this.$store.dispatch(types.actions.UPDATE_USER, {vm: this, item})
          } else {
            this.$store.dispatch(types.actions.CREATE_USER, {vm: this, item})
          }
        }
      },
      components: {
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

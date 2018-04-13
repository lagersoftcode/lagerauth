<template>
  <section class="section">
    <div class="container">

       <h1 class="title">users</h1>
      <button class="button is-primary" @click="navigate(null)">create new</button>

      <b-table :data="users" :selected.sync="selected" paginated :per-page='pagination'>
        
        <template scope="props">
          <b-table-column field="id" label="id" width="40" numeric sortable> {{ props.row.id }} </b-table-column>
          <b-table-column field="name" label="name" sortable> {{ props.row.name }} </b-table-column>
          <b-table-column field="email" label="email" sortable> {{ props.row.email }} </b-table-column>
          <b-table-column field="department" label="department" sortable> {{ props.row.department }} </b-table-column>
          <b-table-column label="description"> {{ props.row.description }} </b-table-column>
          <b-table-column label="enabled"> {{ props.row.enabled }} </b-table-column>
         
          <b-table-column label="">
            <button class="button is-info" @click="navigate(props.row)">edit</button>
            <button class="button is-danger" @click="tryDeleteItem(props.row)">delete</button>
          </b-table-column>
        </template>

        <template slot="empty">
          <section class="section">
            <div class="content has-text-grey has-text-centered">
              <p>
                  <b-icon
                      icon="sentiment_very_dissatisfied"
                      size="is-large">
                  </b-icon>
              </p>
              <p>no users found</p>
            </div>
          </section>
        </template>

      </b-table>
    </div>
  </section>
</template>

<script>
    import types from '../../store/User/types'
    import {mapGetters} from 'vuex'
    
    export default {
      name: 'Users',
      data () {
        return {
          selected: null
        }
      },
      created () {
        this.$store.dispatch(types.actions.GET_USERS, this)
      },
      computed: {
        pagination: () => process.env.PAGINATION,
        ...mapGetters({
          users: types.getters.USERS
        })
      },
      methods: {
        navigate (item) {
          if (item === null) {
            this.$router.push('/user')
          } else {
            this.$router.push(`/user/${item.id}`)
          }
        },
        tryDeleteItem (item) {
          this.$dialog.confirm({
            message: `are you sure you want to delete the user: <b>${item.name}</b> with login: ${item.email} ?`,
            type: 'is-danger',
            hasIcon: true,
            confirmText: 'delete user',
            cancelText: 'cancel',
            onConfirm: () => this.$store.dispatch(types.actions.DELETE_USER, {vm: this, item})
          })
        }
      }
    }
</script>
<template>
  <section class="section">
    <div class="container">

      <h1 class="title">roles</h1>

      <div class="columns">
        <div class="column">
          <button class="button is-primary" @click="navigate(null)">create new</button>
        </div>
      
        <div class="column is-4 is-offset-4">
          <div class="control has-icons-right">
            <b-input v-model="search"></b-input>  
            <span class="icon is-small is-right">
              <i class="fa fa-search"></i>
            </span>
          </div>
        </div>
      </div>
      
      <b-table :data="roles" selectable :selected.sync="selected" paginated :per-page='pagination'>        
        <template scope="props">
          <b-table-column field="id" label="id" width="40" numeric sortable> {{ props.row.id }} </b-table-column>
          <b-table-column field="application" label="application" sortable> {{ props.row.application.name }} </b-table-column>
          <b-table-column field="name" label="name" sortable> {{ props.row.name }} </b-table-column>
          <b-table-column label="description"> {{ props.row.description }} </b-table-column>
          <b-table-column label="permissions"> {{ count(props.row.permissions) }} </b-table-column>
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
              <p>no roles found</p>
            </div>
          </section>
        </template>

      </b-table>
    </div>
  </section>
</template>

<script>
    import types from '../../store/Role/types'
    import {arrCount} from '../../helpers/count'

    export default {
      name: 'Roles',
      data () {
        return {
          selected: null,
          search: ''
        }
      },
      created () {
        this.$store.dispatch(types.actions.GET_ROLES, this)
      },
      computed: {
        pagination: () => process.env.PAGINATION,
        roles () {
          return this.$store.state.Role.Roles.filter(this.searchPredicate)
        }
      },
      methods: {
        navigate (item) {
          if (item === null) {
            this.$router.push('/role')
          } else {
            this.$router.push(`/role/${item.id}`)
          }
        },
        tryDeleteItem (item) {
          this.$dialog.confirm({
            message: `are you sure you want to delete the role: <b>${item.name}</b> ?`,
            type: 'is-danger',
            hasIcon: true,
            confirmText: 'delete role',
            cancelText: 'cancel',
            onConfirm: () => this.$store.dispatch(types.actions.DELETE_ROLE, {vm: this, item})
          })
        },
        count (ary) {
          return arrCount(ary)
        },
        searchPredicate (item) {
          if (!this.search) {
            return true
          }

          return item.application.name.toLowerCase().includes(this.search.toLowerCase()) ||
          item.name.toLowerCase().includes(this.search.toLowerCase()) ||
          item.description.toLowerCase().includes(this.search.toLowerCase())
        }
      }
    }
</script>
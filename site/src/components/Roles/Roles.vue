<template>
  <section class="section">
    <div class="container">

      <h1 class="title">roles</h1>

      <button class="button is-primary" @click="navigate(null)">create new</button>
      
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
    import {mapGetters} from 'vuex'

    export default {
      name: 'Roles',
      data () {
        return {
          selected: null
        }
      },
      created () {
        this.$store.dispatch(types.actions.GET_ROLES, this)
      },
      computed: {
        pagination: () => process.env.PAGINATION,
        ...mapGetters({
          roles: types.getters.ROLES
        })
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
          if (!ary) {
            return 0
          }
          return ary.length
        }
      }
    }
</script>
<template>
  <section class="section">
    <div class="container">
      <b-table :data="users" :selected.sync="selected" paginated :per-page='pagination'>
        
        <template scope="props">
          <b-table-column label="id" width="40" numeric> {{ props.row.id }} </b-table-column>
          <b-table-column label="email"> {{ props.row.email }} </b-table-column>
          <b-table-column label="name"> {{ props.row.name }} </b-table-column>
          <b-table-column label="department"> {{ props.row.department }} </b-table-column>
          <b-table-column label="description"> {{ props.row.description }} </b-table-column>
          <b-table-column label="enabled"> {{ !props.row.isLockedout }} </b-table-column>
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
      }
    }
</script>
<template>
  <section class="section">
    <div class="container">

      <h1 class="title">applications</h1>

      <button class="button is-primary" @click="navigate(null)">Create New</button>
      
      <b-table :data="applications" selectable :selected.sync="selected" paginated :per-page='pagination'>        
        <template scope="props">
          <b-table-column label="id" width="40" numeric> {{ props.row.id }} </b-table-column>
          <b-table-column label="client id" > {{ props.row.clientId }} </b-table-column>
          <b-table-column label="secret key" > <button @click="showSecretKey(props.row)" class="button">show</button> </b-table-column>
          <b-table-column label="name"> {{ props.row.name }} </b-table-column>
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
              <p>no applications found</p>
            </div>
          </section>
        </template>

      </b-table>
    </div>
  </section>
</template>

<script>
    import types from '../../store/Application/types'
    import {mapGetters} from 'vuex'

    export default {
      name: 'Applications',
      data () {
        return {
          selected: null
        }
      },
      created () {
        this.$store.dispatch(types.actions.GET_APPLICATIONS, this)
      },
      computed: {
        pagination: () => process.env.PAGINATION,
        ...mapGetters({
          applications: types.getters.APPLICATIONS
        })
      },
      methods: {
        navigate (item) {
          if (item === null) {
            this.$router.push('/application')
          } else {
            this.$router.push(`/application/${item.id}`)
          }
        },
        showSecretKey (item) {
          this.$dialog.alert({
            title: `secret key for ${item.name}`,
            message: item.secretKey,
            confirmText: 'ok',
            size: 'is-large'
          })
        },
        tryDeleteItem (item) {
          this.$dialog.confirm({
            message: `are you sure you want to delete the application: <b>${item.name}</b> ?`,
            type: 'is-danger',
            hasIcon: true,
            confirmText: 'delete application',
            cancelText: 'cancel',
            onConfirm: () => this.confirmDeleteItem(item)
          })
        },
        confirmDeleteItem (item) {
          let messages = [
            'i am afraid i cant let you do that',
            `i'ts not a good idea to delete THIS application`,
            'think hard and long what you are trying to do',
            'are you trying to KILL me?'
          ]
          let message = messages[Math.floor(Math.random() * messages.length)]

          if (item.clientId === process.env.OAUTH_CLIENT_ID) {
            this.$snackbar.open({
              message,
              type: 'is-danger',
              actionText: 'got it'
            })
          } else {
            this.$store.dispatch(types.actions.DELETE_APPLICATION, {vm: this, item})
          }
        }
      }
    }
</script>
<template>
  <section class="section">
    <div class="container">

      <h1 class="title">applications</h1>

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
      
      <b-table :data="applications" selectable :selected.sync="selected" paginated :per-page='pagination'>        
        <template scope="props">
          <b-table-column field="id" label="id" width="40" numeric sortable> {{ props.row.id }} </b-table-column>
          <b-table-column field="name" label="name" sortable> {{ props.row.name }} </b-table-column>
          <b-table-column label="client id" > {{ props.row.clientId }} </b-table-column>
          <b-table-column label="secret key" > <button @click="showSecretKey(props.row)" class="button">show</button> </b-table-column>
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

    export default {
      name: 'Applications',
      data () {
        return {
          selected: null,
          search: ''
        }
      },
      created () {
        this.$store.dispatch(types.actions.GET_APPLICATIONS, this)
      },
      computed: {
        pagination: () => process.env.PAGINATION,
        applications () {
          return this.$store.state.Application.Applications.filter(this.searchPredicate)
        }
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
        },
        searchPredicate (item) {
          if (!this.search) {
            return true
          }

          return item.name.toLowerCase().includes(this.search.toLowerCase()) ||
          item.description.toLowerCase().includes(this.search.toLowerCase()) ||
          item.clientId.toLowerCase().includes(this.search.toLowerCase()) ||
          item.secretKey.toLowerCase().includes(this.search.toLowerCase())
        }
      }
    }
</script>
<template>
  <section class="section">
    <div class="container">


        <h1 v-if="this.isNew" class="title" >create application</h1>
        <h1 v-else class="title">edit application</h1>
        <hr />

        <div class="columns">
          <div class="column is-6">
            
            <form @submit.prevent="save">             
              <b-field label="name">
                  <b-input v-model="name" required></b-input>
              </b-field>

              <div class="field">
                <label class="label">client id</label>
                <div class="field has-addons">
                  <p class="control is-expanded">
                    <b-input v-model="clientId" required></b-input>
                  </p>
                  <button @click="generateClientId" v-if="this.clientId === ''" class="button is-primary">generate</button>
                </div>
              </div>

              <div class="field">
                <label class="label">secret key</label>
                <div class="field has-addons">
                  <p class="control is-expanded">
                    <b-input v-model="secretKey" required></b-input>
                  </p>
                  <button @click="generateSecretKey" v-if="this.secretKey === ''" class="button is-primary">generate</button>
                </div>
              </div>

              <b-field label="enabled">
                <input type="checkbox" v-model="enabled" />
              </b-field>

              <b-field label="description">
                <b-input type="textarea" v-model="description"></b-input>
              </b-field>

              <button class="button" @click.prevent="$router.push('/applications')">cancel</button>
              <input type="submit" :value="this.isNew ? 'create' : 'update'" class="button is-success">
            </form>
            
          </div>
        </div>

    </div>
  </section>
</template>

<script>
    import types from '../../store/Application/types'
    import {mapGetters} from 'vuex'
    import uuid from 'uuid/v4'
    import generatePassword from 'password-generator'

    export default {
      name: 'Application',
      data () {
        return {
          id: null,
          name: '',
          description: '',
          clientId: '',
          secretKey: '',
          enabled: false
        }
      },
      created () {
        this.$store.dispatch(types.actions.GET_APPLICATION, {vm: this, id: this.$route.params.id}).then(() => {
          this.id = this.application.id || null
          this.name = this.application.name || ''
          this.clientId = this.application.clientId || ''
          this.secretKey = this.application.secretKey || ''
          this.enabled = this.application.enabled || false
          this.description = this.application.description || ''
        })
      },
      computed: {
        isNew () {
          return this.id === null
        },
        ...mapGetters({
          application: types.getters.APPLICATION
        })
      },
      methods: {
        generateClientId () {
          this.clientId = uuid()
        },
        generateSecretKey () {
          this.secretKey = generatePassword(48, false, /[\da-z]/i)
        },
        save () {
          let item = {id: this.id, name: this.name, description: this.description, clientId: this.clientId, secretKey: this.secretKey, enabled: this.enabled}

          if (item.id) {
            this.$store.dispatch(types.actions.UPDATE_APPLICATION, {vm: this, item})
          } else {
            this.$store.dispatch(types.actions.CREATE_APPLICATION, {vm: this, item})
          }
        }
      }
    }
</script>
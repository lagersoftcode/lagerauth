<template>
  <div>
      <table>
          <tr>
              <th>method</th>
              <th>controller</th>
              <th>action</th>
              <th></th>
          </tr>

          <tr>
              <td>
                <b-select v-model="current.method" expanded>
                  <option value="*">*</option>
                  <option value="get">get</option>
                  <option value="post">post</option>
                  <option value="delete">delete</option>
                  <option value="put">put</option>
                  <option value="patch">patch</option>
                  <option value="head">head</option>
                  <option value="options">options</option>
                </b-select>
              </td>
              <td><b-input v-model="current.controller" /></td>
              <td><b-input v-model="current.action" @keyup.enter="addItem"/></td>
              <td>
                  <button @click="addItem" class="button is-success">
                      <span class="icon is-small">
                          <i class="fa fa-plus" />
                      </span>
                  </button>
              </td>
          </tr>

          <tr :key="itemKey(item)" v-for="item in data">
              <td>{{item.method}}</td>
              <td>{{item.controller}}</td>
              <td>{{item.action}}</td>
              <td>
                  <button @click="removeItem(item)" class="button is-danger">
                    <span class="icon is-small">
                        <i class="fa fa-times" />
                    </span>                                        
                  </button>
              </td>
          </tr>
      </table>
  </div>
</template>

<script>
export default {
  name: 'Permissions',
  props: {
    data: {
      default: [],
      type: Array
    }
  },
  data () {
    return {
      current: {
        method: '',
        controller: '',
        action: ''
      }
    }
  },
  methods: {
    itemKey (item) {
      return item.method + '_' + item.controller + '_' + item.action
    },
    removeItem (item) {
      let newList = this.data.filter(i => this.itemKey(item) !== this.itemKey(i))
      this.$emit('update:data', newList)
    },
    addItem () {
      if (this.current.method && this.current.action && this.current.controller) {
        let exists = this.data.filter((i) => i.action === this.current.action && i.method === this.current.method && i.controller === this.current.controller)
        if (exists.length === 0) {
          let newList = this.data.slice(0)
          let clone = { method: this.current.method, controller: this.current.controller, action: this.current.action }
          newList.push(clone)
          this.$emit('update:data', newList)
        }
        this.current.method = ''
        this.current.controller = ''
        this.current.action = ''
      }
    }
  }
}
</script>

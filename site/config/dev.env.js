var merge = require('webpack-merge')
var prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',

  API_URL: '"http://localhost:8081"',
  OAUTH_LOGIN: '"http://localhost:8081/auth"'
})

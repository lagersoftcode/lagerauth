var merge = require('webpack-merge')
var prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',

  PAGINATION: 15,
  API_URL: '"http://localhost:8081"',
  OAUTH_LOGIN: '"http://localhost:8081/auth"',
  OAUTH_CLIENT_ID: '"00000000-0000-0000-0000-000000000001"'
})

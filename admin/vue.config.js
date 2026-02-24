const path = require('path');
module.exports = {
  publicPath: process.env.NODE_ENV === 'production' ? '/admin/' : '/',
  devServer: {
      host: '0.0.0.0',
      port: 5000,
      proxy: {
        '/api/admin/': {
          target: 'http://localhost:8080/'
        }
      }
  },
  pluginOptions: {
    'style-resources-loader': {
      preProcessor: 'less',
      patterns: [
        path.resolve(__dirname, './src/style/index.less'),
        path.resolve(__dirname, './src/style/variables.less'),
      ]
    }
  }
}

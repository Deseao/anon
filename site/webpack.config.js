const path = require('path')

module.exports = {
  entry: ['babel-polyfill','./src/app.js'],
  output: {
    path: path.resolve(__dirname, 'public'),
    filename: 'bundle.js'
  },
  module: {
    rules: [
      { use: 'babel-loader', test:/\.js$/ }
    ]
  },
  devServer: {
    contentBase: path.join(__dirname, 'public'),
    historyApiFallback: true,
    overlay: {
      warnings: true,
      errors: true
    },
    watchContentBase: true,
    stats: "minimal"
  }
}

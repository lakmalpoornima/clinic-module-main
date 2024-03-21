module.exports = {
  lintOnSave: false,
  runtimeCompiler: true,
  configureWebpack: {
    devtool: process.env.NODE_ENV === 'production' ? false :'eval-source-map',
    //devtool: 'eval-source-map',
    // devtool: false,
  },
  transpileDependencies: [
    '@coreui/utils'
  ],
  // use this option for production linking
  // publicPath: process.env.NODE_ENV === 'production' ? '/vue/demo/3.1.0' : '/'
}

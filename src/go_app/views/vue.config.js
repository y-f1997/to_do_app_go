module.exports = {
  devServer: {
      port: 5081,
      disableHostCheck: true,
      proxy: {
        '^/': { 
          target: 'http://localhost:5080'
        }
      }
  },
};
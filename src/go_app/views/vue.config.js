module.exports = {
  devServer: {
      port: 5443,
      disableHostCheck: true,
      proxy: {
        '^/': { 
          target: 'http://localhost:5080'
        }
      }
  },
};
module.exports = (options, req) => ({
  entry: 'app.js',
  autoprefixer: {
    browsers: ['ie > 8', 'last 3 versions']
  }
  // Other options
})

// Note that you can directly export an object too:
// module.exports = { port: 5000 }
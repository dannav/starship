const express = require('express')
const next = require('next')

const port = parseInt(process.env.PORT, 10) || 6060
const dev = process.env.NODE_ENV !== 'production'
const app = next({ dev })
const handle = app.getRequestHandler()

// for every request that comes through the server pass it to nextjs
// which will execute our react code and render markup
app.prepare().then(() => {
  const server = express()

  server.get('/static/*', (req, res) => {
    handle(req, res)
  })

  server.get('/posts/:id', (req, res) => {
    return app.render(req, res, '/posts', { id: req.params.id })
  })

  server.get('*', (req, res) => {
    return handle(req, res)
  })

  server.listen(port, err => {
    if (err) throw err
    console.log(`> Ready on http://localhost:${port}`)
  })
})

const express = require('express')
const next = require('next')

const port = parseInt(process.env.PORT, 10) || 6060
const dev = process.env.NODE_ENV !== 'production'
const app = next({ dev })
const handle = app.getRequestHandler()

app.prepare().then(() => {
  const server = express()

  server.get('/static/*', (req, res) => {
    handle(req, res)
  })

  server.get('/not-found', (req, res) => {
    handle(req, res, '/not-found')
  })

  server.get('/:folder/:page', (req, res) => {
    return app.render(req, res, '/doc', { folder: req.params.folder, page: req.params.page })
  })

  server.get('*', (req, res) => {
    return handle(req, res)
  })

  server.listen(port, err => {
    if (err) throw err
    console.log(`> Ready on http://localhost:${port}`)
  })
})

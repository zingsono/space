const path =  require('path')
const Vue = require('vue')
const server = require('express')()
const renderer = require('vue-server-renderer').createRenderer({template: require('fs').readFileSync(path.join(__dirname,'index.template.html'), 'utf-8')})

const createApp = require('./app')

server.get('*', (req, res) => {
    const app = createApp({ url: req.url })

    const context = {
        title: 'hello',
        meta: `
            <meta name="keywords" content="1">
            <meta name="description" content="2">
        `
    }
    renderer.renderToString(app ,context,(err, html) => {
        if (err) {
            res.status(500).end('Internal Server Error')
            return
        }
        res.end(html)
    })
})

server.listen(8080)
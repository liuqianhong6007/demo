module.exports = {
    publicPath: '/product/front-page',
    devServer: {
        proxy: {
            '/auth': {
                target: 'http://47.101.171.154/',
                ws: true,
                changeOrigin: true
            },
            '/etcd': {
                target: 'http://47.101.171.154/',
                changeOrigin: true
            },
            '/mytable': {
                target: 'http://47.101.171.154/',
                changeOrigin: true
            }
        }
    }
}
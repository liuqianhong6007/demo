module.exports = {
    publicPath: process.env.NODE_ENV === 'production' ? '/product/front-page' : '/dev/front-page',
    devServer: {
        proxy: {
            '/auth': {
                target: 'http://127.0.0.1:8081/',
                ws: true,
                changeOrigin: true
            },
            '/etcd': {
                target: 'http://127.0.0.1:8082/',
                changeOrigin: true
            }
        }
    }
}
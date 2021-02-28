module.exports = {
    publicPath: process.env.NODE_ENV === 'production' ? '/product/etcd-front' : '/dev/etcd-front',
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
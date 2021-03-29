

module.exports = {
    dev: {
        assetsSubDirectory: 'static',
        assetsPublicPath: '/',
        proxyTable: {
            'v1': {
                target: 'http://42.193.110.73:8083',
                changeOrigin: true,
                pathRewrite: {
                    '^/v1': ''
                }
            }
        },
        host: 'localhost',

    }
}
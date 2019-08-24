const path = require('path')
let pack = require('./package.json')
let zipPath = './build/' + pack.name + '-' + pack.version + '.zip'
console.log('zipPath=',zipPath)

module.exports = {
    publicPath: './',
    outputDir: './dist',
    indexPath: 'index.html',
    assetsDir: 'assets',
    filenameHashing: true,
    productionSourceMap: false,
    devServer: {
        open: false,
        port: 8000
    },
    configureWebpack: {
        plugins: [
           /* //使用Gzip压缩插件
            new (require('compression-webpack-plugin'))({
                test: /\.js$|\.html$|\.css/,
                threshold: 5*1024, //超过5k的文件压缩
                deleteOriginalAssets:false  //删除原文件
            })*/
        ],
        externals: {
            //vue:'Vue',
            //'vue-router': 'VueRouter',
        },
        optimization: {
            minimizer: [
                /*// 删除发布版本控制台日志console.log
                new (require('terser-webpack-plugin'))({
                    terserOptions: {
                        compress: {
                            //warnings: true,
                            //drop_console: true,
                            //drop_debugger: true,
                            //pure_funcs: ['console.log']
                        }
                    }
                })*/
            ]
        },
        // 警告 webpack 的性能提示
        performance: {
            //hints:'warning',
            hints: false, //关闭 webpack 的性能提示
            //入口起点的最大体积
            maxEntrypointSize: 2000 * 1024,
            //生成文件的最大体积
            maxAssetSize: 3000 * 1024,
            //只给出 js 文件的性能提示
            /*assetFilter: function(assetFilename) {
                return assetFilename.endsWith('.js')
            }*/
        }
    }
}
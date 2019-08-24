const path = require('path')

module.exports = {
    mode: 'production',
    entry: './src/main.js',
    output: {
        publicPath:'/',
        filename: 'main.js',
        path: path.resolve(__dirname, 'dist')
    },
    //devtool: 'inline-source-map',
    devServer: {
        contentBase: './dist',
        hot: true
    },
    module: {
        rules: [
            {test: /\.css$/, use: ['style-loader', 'css-loader']},
            {test: /\.(png|svg|jpg|gif|woff|woff2|eot|ttf|otf)$/, use: ['file-loader']},
            {test: /\.tsx?$/, use: 'ts-loader', exclude: /node_modules/}
        ]
    },
    plugins: [
        /* 每次编译清理输出文件夹 */
        new (require('clean-webpack-plugin'))(),
        /* 输出文件index.html  配置参数 https://github.com/jantimon/html-webpack-plugin  */
        new (require('html-webpack-plugin'))({
            title: '网页Title',
            hash: true,
            template:'./public/index.html',
            filename: 'index.html',
            favicon: './public/favicon.ico',
            meta:{
                'viewport': 'width=device-width, initial-scale=1, shrink-to-fit=no',
            }
        }),
        /* 模块热替换 */
        new (require('webpack').NamedModulesPlugin)(),
        new (require('webpack').HotModuleReplacementPlugin)(),
        /* 定义环境变量 */
        new (require('webpack').DefinePlugin)({
            'process.env.NODE_ENV': JSON.stringify('production')
        }),
    ],
    /* 外部依赖 */
    externals: {

    }
}

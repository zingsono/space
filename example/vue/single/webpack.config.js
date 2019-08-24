const path = require('path');
const VueLoaderPlugin = require('vue-loader/lib/plugin');

const MinifyPlugin = require("babel-minify-webpack-plugin");

module.exports = {
    mode: 'production', //production  development
    entry: {
        app: ['./src/index.js']
    },
    output: {
        path: path.resolve(process.cwd(), './lib'),
        publicPath: '/dist/',
        filename: 'index.js'
    },
    devServer:{
        hot:true
    },
    resolve: {
        extensions: ['.js', '.vue', '.json'],
    },
    externals: {
    },
    optimization: {
        minimizer: [ ]
    },
    module: {
        rules: [
            // 它会应用到普通的 `.js` 文件
            // 以及 `.vue` 文件中的 `<script>` 块
            {
                test: /\.js$/,
                loader: 'babel-loader'
            },
            {
                test: /\.vue$/,
                loader: 'vue-loader',
                options: {
                    compilerOptions: {
                        preserveWhitespace: false
                    }
                }
            },
            // 它会应用到普通的 `.css` 文件
            // 以及 `.vue` 文件中的 `<style>` 块
            {
                test: /\.css$/,
                loaders: ['vue-style-loader','style-loader', 'css-loader']
            },
            {
                test: /\.(svg|otf|ttf|woff2?|eot|gif|png|jpe?g)(\?\S*)?$/,
                loader: 'url-loader' ,
                query: {
                    limit: 10000,
                    name: path.posix.join('static', '[name].[hash:7].[ext]')
                }
            }
        ]
    },
    plugins: [
        // 请确保引入这个插件来施展魔法
        new VueLoaderPlugin(),
        new MinifyPlugin({},{comments:false})
    ]
};

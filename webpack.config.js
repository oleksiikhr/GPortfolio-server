'use strict'

const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const { CleanWebpackPlugin } = require('clean-webpack-plugin')
const { VueLoaderPlugin } = require('vue-loader/lib/index')
const WebpackPwaManifest = require('webpack-pwa-manifest')
const HtmlWebpackPlugin = require('html-webpack-plugin')
const { GenerateSW } = require('workbox-webpack-plugin')
const path = require('path')
const packageJson = require('./package')

require('dotenv').config({
  path: path.resolve(__dirname, './.env')
})

module.exports = (env, argv) => {
  const isProd = argv && argv.mode === 'production'

  const config = {
    mode: isProd ? 'production' : 'development',
    entry: [
      'ant-design-vue/dist/antd.less',
      './web/styles/index.scss',
      './web/main.js'
    ],
    devtool: isProd ? false : 'source-map',
    output: {
      filename: 'static/[name].js',
      chunkFilename: 'static/chunks/[name].[hash].js',
      publicPath: '/',
      path: path.resolve(__dirname, './dist')
    },
    devServer: {
      publicPath: '/',
      contentBase: './dist',
      host: process.env.WEBPACK_DEV_HOST || 'localhost',
      port: +process.env.WEBPACK_DEV_PORT || 3000,
      hot: true,
      clientLogLevel: 'error',
      proxy: {
        '/api/*': {
          target: (process.env.APP_TLS === 'true' ? 'https://' : 'http://') + (process.env.APP_ADDR || 'localhost:8080'),
          changeOrigin: true
        }
      }
    },
    module: {
      rules: [
        {
          test: /\.s?css$/,
          use: [
            {
              loader: MiniCssExtractPlugin.loader,
              options: {
                hmr: !isProd
              }
            },
            'css-loader',
            'resolve-url-loader',
            {
              loader: 'sass-loader',
              options: {
                sourceMap: true
              }
            }
          ]
        },
        {
          test: /\.less$/,
          use: [
            MiniCssExtractPlugin.loader,
            'css-loader',
            {
              loader: 'less-loader',
              options: {
                modifyVars: {
                  'primary-color': '#3f3d56',
                  'link-color': '#3f3d56'
                },
                javascriptEnabled: true
              }
            }
          ]
        },
        {
          test: /\.(png|svg|jpe?g|gif)$/,
          use: [
            {
              loader: 'file-loader',
              options: {
                outputPath: 'static/images'
              }
            }
          ]
        },
        {
          test: /\.(eot|ttf|woff|woff2)$/,
          use: [
            {
              loader: 'file-loader',
              options: {
                outputPath: 'static/files'
              }
            }
          ]
        },
        {
          test: /\.vue$/,
          use: 'vue-loader'
        },
        {
          test: /\.m?js$/,
          exclude: /(node_modules|bower_components)/,
          use: [
            'babel-loader',
            'eslint-loader'
          ]
        }
      ]
    },
    plugins: [
      new VueLoaderPlugin,
      new HtmlWebpackPlugin({
        filename: 'index.html',
        template: './index.ejs',
        inject: true,
        meta: {
          description: 'Create an automatic portfolio based on Github and other various data',
          robots: 'index, follow',
          viewport: 'width=device-width, initial-scale=1, shrink-to-fit=no'
        },
        minify: isProd ? {
          collapseWhitespace: true,
          removeComments: true,
          removeRedundantAttributes: true,
          removeScriptTypeAttributes: true,
          removeStyleLinkTypeAttributes: true
        } : false,
        chunksSortMode: 'none'
      }),
      new MiniCssExtractPlugin({
        filename: 'static/[name].css',
        chunkFilename: 'static/css/[name].[hash].css'
      })
    ],
    resolve: {
      extensions: ['.js', '.vue', '.json'],
      alias: {
        '@': path.resolve(__dirname, './web/'),
        'scss': path.resolve(__dirname, './web/styles/')
      }
    }
  }

  if (isProd) {
    config.plugins.unshift(new CleanWebpackPlugin)
    config.plugins.push(
      new WebpackPwaManifest({
        background_color: '#fff',
        description: packageJson.description,
        filename: 'static/manifest.[hash].json',
        icons: [
          {
            destination: 'static/icons',
            sizes: [96, 128, 192, 256, 384, 512],
            src: require('./web/images/icon.png')
          }
        ],
        name: 'GPortfolio',
        short_name: 'GPortfolio',
        start_url: '/',
        theme_color: '#fff'
      }),
      new GenerateSW({
        swDest: 'sw.js',
        skipWaiting: true,
        clientsClaim: true,
        importWorkboxFrom: 'local',
        importsDirectory: 'static/pwa',
        navigateFallback: '/index.html',
        navigateFallbackWhitelist: [
          /^static/, /^sw\.js$/, /^index\.html$/, /^favicon\.ico$/
        ]
      })
    )
  }

  return config
}

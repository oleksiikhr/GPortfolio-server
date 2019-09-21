'use strict'

const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const { CleanWebpackPlugin } = require('clean-webpack-plugin')
const WebpackPwaManifest = require ('webpack-pwa-manifest')
const { VueLoaderPlugin } = require('vue-loader/lib/index')
const HtmlWebpackPlugin = require('html-webpack-plugin')
const { GenerateSW } = require('workbox-webpack-plugin')
const path = require('path')
const packageJson = require('./package')

require('dotenv').config({
  path: path.resolve(__dirname, '../.env')
})

module.exports = (env, { mode }) => {
  const isProd = mode === 'production'

  const config = {
    mode: mode,
    entry: [
      './src/main.js',
      './src/styles/index.scss'
    ],
    devtool: isProd ? false : 'source-map',
    output: {
      filename: 'static/[name].js',
      chunkFilename: 'static/chunks/[name].[hash].js',
      publicPath: '/',
      path: path.resolve(__dirname, '../dist')
    },
    devServer: {
      publicPath: '/',
      contentBase: './dist',
      host: process.env.WEBPACK_DEV_HOST || 'localhost',
      port: process.env.WEBPACK_DEV_PORT || '3000',
      hot: true,
      clientLogLevel: 'error',
      disableHostCheck: true,
      proxy: {
        '/api/*': {
          target: process.env.APP_ADDR || 'http://localhost:8080/',
          changeOrigin: true
        }
      }
    },
    module: {
      rules: [
        {
          test: /\.scss$/,
          use: [
            MiniCssExtractPlugin.loader,
            'css-loader',
            'sass-loader'
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
      new CleanWebpackPlugin,
      new HtmlWebpackPlugin({
        filename: 'index.html',
        template: './index.ejs',
        inject: true,
        meta: {
          description: 'Create an automatic portfolio based on Github and other various data',
          robots: 'index, follow',
          viewport: 'width=device-width, initial-scale=1, shrink-to-fit=no',
        },
        minify: isProd ? {
          collapseWhitespace: true,
          removeComments: true,
          removeRedundantAttributes: true,
          removeScriptTypeAttributes: true,
          removeStyleLinkTypeAttributes: true,
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
        '@': path.resolve(__dirname, './src/'),
        'scss': path.resolve(__dirname, './src/styles/')
      }
    }
  }

  if (isProd) {
    config.plugins.push(
      new WebpackPwaManifest({
        background_color: '#fff',
        description: packageJson.description,
        filename: 'static/manifest.[hash].json',
        icons: [
          {
            destination: 'static/icons',
            sizes: [96, 128, 192, 256, 384, 512],
            src: path.resolve('public/icon.png')
          }
        ],
        name: `GPortfolio`,
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

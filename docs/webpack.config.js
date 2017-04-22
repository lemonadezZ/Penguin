var path = require('path');
var webpack=require('webpack');
var HtmlWebpackPlugin = require('html-webpack-plugin')

module.exports = {
  entry: {
	main:'./src/app.jsx',
	react:['react','react-dom'],
	},
  output: {
    filename: 'js/[name]-[chunkhash:8].js',
	path: path.resolve(__dirname, 'dist')
	
  },
	 module: {
		rules:[	
			{
				 test: /\.js$/,
				 exclude: /(node_modules)/,
				 use: ["balel-loader"]
			}
		],
		rules:[	
			{
				 test: /\.jsx$/,
				 exclude: /(node_modules)/,
				 use: ["jsx-loader"]
			}
		]
	},
		
	plugins: [
	
	new webpack.optimize.UglifyJsPlugin({
      		compress: {
       		 warnings: false
      		}
    	}),
	new webpack.optimize.CommonsChunkPlugin({
			name:['react'], 
			filename:'vendor-[chunkhash:8].js'
		}),
	new HtmlWebpackPlugin({
		title: 'Penguin',
		template: 'src/index.html',
		excludeChunks:['test']
	//	minify: { collapseWhitespace:true},
	})

	],
	resolve: {
		 extensions: ['.js', '.json', '.coffee','jsx'],
		 modules: [
     			"node_modules",
      			path.resolve(__dirname, "src")
    		],
	}
	
}

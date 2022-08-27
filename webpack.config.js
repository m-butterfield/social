const path = require("path");

module.exports = {
  entry: {
    index: "./app/static/ts/app/index.tsx",
  },
  devtool: "inline-source-map",
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: "ts-loader",
        exclude: /node_modules/,
      },
      {
        test: /\.css$/i,
        use: ["style-loader", "css-loader"],
      },
      {
        test: /\.(png|jpe?g|gif)$/i,
        use: [
          {
            loader: "file-loader",
          },
        ],
      },
    ],
  },
  resolve: {
    extensions: [".tsx", ".ts", ".js"],
    modules: [
      path.resolve(__dirname, "node_modules"),
      path.resolve(__dirname, "app", "static", "ts"),
    ]
  },
  output: {
    filename: "[name].bundle.js",
    path: path.resolve(__dirname, "app", "static", "js", "dist"),
  },
};

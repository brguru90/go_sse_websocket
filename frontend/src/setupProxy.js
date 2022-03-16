// const proxy = require("http-proxy-middleware")
// // var ip = require("ip")
const http = require("http")

const keepAliveAgent = new http.Agent({keepAlive: true})
// module.exports = function (app) {
//     app.use(
//         proxy("/api", {
//             target: "http://localhost:8000/",
//             // target: "http://" + ip.address(),
//             changeOrigin: true,
//             // agent: keepAliveAgent,
//             // pathRewrite: {
//             //     "^/api/": "/", // remove base path
//             // },
//         })
//     )
// }


const { createProxyMiddleware } = require("http-proxy-middleware");

module.exports = function (app) {
  app.use(
    "/api",
    createProxyMiddleware({
      target: "http://localhost:8000",
      changeOrigin: true,
      secure: false,
      agent: keepAliveAgent,
    })
  );
}

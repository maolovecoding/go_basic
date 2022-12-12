const http = require('http')

const server = http.createServer()
server.on('request',(req, res)=>{
  console.log(req)
  res.write("你好")
  res.end()
})

server.listen(8888)
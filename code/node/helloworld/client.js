var PROTO_PATH = __dirname + '../../../protos/helloworld.proto'

var grpc = require('grpc')
var protoLoader = require('@grpc/proto-loader')
var packageDefinition = protoLoader.loadSync(PROTO_PATH)
var hello_proto = grpc.loadPackageDefinition(packageDefinition).helloworld

function main() {
  var client = new hello_proto.Greeter('localhost:50051', grpc.credentials.createInsecure())

  const md = new grpc.Metadata()
  md.add('token', 'xyz')
  md.add('request-id', '123')
  
  client.sayHello({ name: 'world' }, md, (err, response) => {
    console.log('Greeting: ', response.message)
  })
}

main()

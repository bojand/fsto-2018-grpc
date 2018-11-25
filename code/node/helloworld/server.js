var PROTO_PATH = __dirname + '../../../protos/helloworld.proto'
var grpc = require('grpc')
var protoLoader = require('@grpc/proto-loader')
var packageDefinition = protoLoader.loadSync(PROTO_PATH)
var hello_proto = grpc.loadPackageDefinition(packageDefinition).helloworld

// Implements the SayHello RPC method.
function sayHello(call, callback) {
  const metadata = call.metadata.getMap()
  for (const k in metadata) {
    console.log(`${k}: ${metadata[k]}`)
  }

  callback(null, { message: 'Hello ' + call.request.name })
}

// Starts an RPC server that receives requests for the Greeter service
function main() {
  var server = new grpc.Server()
  server.addService(hello_proto.Greeter.service, { sayHello: sayHello })
  server.bind('0.0.0.0:50051', grpc.ServerCredentials.createInsecure())
  server.start()
}

main()

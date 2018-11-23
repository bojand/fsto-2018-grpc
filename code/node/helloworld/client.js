var PROTO_PATH = __dirname + '../../../protos/helloworld.proto';

var grpc = require('grpc');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(PROTO_PATH);
var hello_proto = grpc.loadPackageDefinition(packageDefinition).helloworld;

function main() {
  var client = new hello_proto.Greeter('localhost:50051',
                                       grpc.credentials.createInsecure());
  
  client.sayHello({ name: 'world' }, (err, response) => {
    console.log('Greeting: ', response.message);
  });
}

main();

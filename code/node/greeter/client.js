const PROTO_PATH = __dirname + '../../../protos/greeter.proto'

const async = require('async')
const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')
const packageDefinition = protoLoader.loadSync(PROTO_PATH)
const proto = grpc.loadPackageDefinition(packageDefinition).greeter

const client = new proto.Greeter('localhost:50051', grpc.credentials.createInsecure())

const NAMES = ['Bob', 'Kate', 'Jim', 'Sara']

function sayHello(fn) {
  console.log('client:sayHello')

  client.sayHello({ name: 'world' }, (err, response) => {
    console.log('Greeting:', response.message)
    fn()
  })
}

function sayHellos(fn) {
  console.log('client:sayHellos')

  const call = client.sayHellos({ name: 'world', count: 5 })
  call.on('data', ({ message }) => console.log('Greeting:', message))
  call.on('end', fn)
}

function greetMany(fn) {
  console.log('client:greetMany')

  call = client.greetMany((err, response) => {
    console.log('Greeting:', response.message)
    fn()
  })

  let n = 0
  const timer = setInterval(() => {
    if (n < NAMES.length) {
      call.write({ name: NAMES[n] })
      n++
    } else {
      clearInterval(timer)
      call.end()
    }
  }, 200)
}

function greetChat(fn) {
  console.log('client:greetChat')

  call = client.greetChat()

  let n = 0
  const timer = setInterval(() => {
    if (n < NAMES.length) {
      call.write({ name: NAMES[n] })
      n++
    } else {
      clearInterval(timer)
      call.end()
    }
  }, 200)

  call.on('data', ({ message }) => console.log('Greeting:', message))

  call.on('end', fn)
}

function main() {
  async.series([sayHello, sayHellos, greetMany, greetChat])
}

main()

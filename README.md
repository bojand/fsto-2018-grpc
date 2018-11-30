## Beyond REST - A Guite to gRPC
### [Full Stack Totonto 2018](http://fsto.co/)

#### Toronto, ON

November 29 & 30, 2018

### Presentation

```sh
npm install
npm start
open http://127.0.0.1:8080
```

### Code samples

#### Prerequisites

- [Go 1.11](http://golang.org)
- [Node.js 8 or 10](http://nodejs.org) 
- [protoc](https://developers.google.com/protocol-buffers/)

#### Instructions

All code samples are located in `code` directory.

`protos` contains the Protocol Buffer definition files.

`codegen.sh` will use `protoc` Protocol Buffers compiler to generate code for Go and Node.js.

Go code samples are in `go` directory.

Node code samples are in `node` directory.

Simple helloworld example is in `helloworld` directory for each language.

Example implementing full Greeter service with streaming calls is in `greeter` directory for each language.

For Node.js examples first use `npm install` to install dependencies in each example directory.
To run Node.js example do `node server.js` and `node client.js` respectivey for each example.

To build Go code in `server` and `client` directories run `go build`, then execute each executable. Optionally you should be able to do `go run server.go` and `go run client.go` for each example.

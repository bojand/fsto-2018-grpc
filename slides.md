class: center, middle

<img src="/img/FSTO_400x400.png" alt="FSTO Logo" width="156px">

.title[Beyond REST]

## A guide to gRPC

???

- **P** for presentation mode
- **C** to clone the window

- Hello and thank you for having me!

---

# ABOUT ME

Bojan Djurkovic

Lead Software Engineer @ Cvent

@bojantweets

https://github.com/bojand/fsto-2018-grpc

???

- My name is Bojan Djurkovic and I am from Fredericton New Brunswick.
- I work as a Lead Software Engineer at Cvent
- Cvent a company that makes software solutions for planning and managing conventions and conferences and events similar to this one, but we're not involved with this event, and I am here kind of on my own
- We're headquartered in Washington DC, and we're probably one of the bigger companies you have never heard of. We're around 3500 employees with offices all over the world and we have 2 offices in Canada, in Vancouver and Fredericton 
- I should mention that we are hiring and the cost of living in Fredericton is more favourable than Toronto 
- Currently most of my work at Cvent involves Java and Kafka
- On the side I try to explore different technologies and do some open source
- Over the past couple of years most of that has been focused on gRPC
- Which is what we're going to be talking about today...

---

# A SIMPLER TIME...

<img src="/img/the_internet.jpg" alt="Welcome to the internet" width="100%">

.footnote[Source: Gizmodo]

???

Lets go back to a simpler time. The internet was simple. Servers served data, usually some HTML, and browsers consumed it.

---

# A NEW AGE

<img src="/img/programming.jpg" alt="programming" width="100%">

.footnote[Source: Pexels]

???

- Then our applications and our servers got more complex and dynamic. 
- How we worked with data and how we consumed it has changed.
- Today browsers are not the only clients we have to deal with. We have different mobile and embedded devices and platforms using our applications.

---

# MICROSERVICES!

.center[<img src="/img/microservices2.png" alt="microservices" width="500">]

.footnote[Source: ArcGIC]

???

- As our systems grew and evolved, to overcome the challenges of scale both technically and organizationally we started breaking up the monolith into microservices
- A single PHP application became a number of different services written in Pyton, Java, Node.js, Go and others.
- Our systems' architecture has drastically changed over the past decade or so.

---

# PRESENT DAY

- Complex polyglot distributed computing landscape
- "API-First"

.center[
  <img src="/img/how.png">
]

???

Today we live in a complex polyglot distributed computing landspace where numerous services and applications needs to interact with each other. So how do all these different components communicate in this new API-first reality?

---

# REST

<blockquote><strong>HTTP + JSON is the de facto standard for REST communication</strong></blockquote>

```sh
$ curl https://api.stripe.com/v1/charges \
   -u sk_test_4eC39HqLyjWDarjtT1zdp7dc: \
   -d amount=2000 \
   -d currency=usd \
   -d source=tok_visa \
   -d description="Charge for jenny.rosen@example.com"
```

???

- HTTP/REST is great in many ways and  we're all familiar with it
- Text-based and relatively easy to debug
- Tooling for testing & inspection
- Well-supported in most languages 
- Cacheable
- Scalable
- Easy?
- Standardized?
- Performant?

---

# REST API CONSIDERATIONS

- Schema
- Authentication
- Documentation
- Versioning
- Root endpoint
- Status code & client errors
- Redirects
- HTTP verbs
- Hypermedia
- Pagination
- Conditional Requests
- CORS
- JSON-P
- Callbacks

???

So if we are building a REST API, what are all the things we need to consider? This is probably not an exhaustive or complete list, but just some items that we need to think about.

---

# CLIENT LIBRARIES

.left[![Issue](/img/language_issue.png)]
<!-- <img src="/img/language_issue.png" alt="Issue" width="640px"> -->

???

- We get requests for client libraries in particular language

---

# REST API DEFICIENCIES

- HTTP/1 is not performant
- Text-based protocol is developer-friendly but inefficient
- Streaming is difficult
- No bi-directional streaming
- Not all operations are "resource"-based
- Semantics
  * `POST`/`PUT`/`PATCH`
  * Status codes & Error responses
  * Single vs. plural resource names
  * ID in param or in body?

???

- Example of non-resourceful: 
    * Encrypt some text
    * Restart some application or host
    * Describe an image
    * Classify a sentence for sentiment analysis

---

# EXAMPLES

.left[<img src="/img/twilio-logo-red.svg" alt="Twilio Logo" height="42px">]

- `200 OK` for `GET`
- `201 CREATED` for `POST` and `PUT`
- `204 NO CONTENT` for `DELETE`

.left[<img src="/img/Stripe logo - blue.svg" alt="Stripe Logo" height="56px">]

- `200 OK` for all successful requests

???

- Not picking on any one company here these are just a couple of examples of well-established and successful API products. 

- Even with these popular API's we see a difference in opinion on the very basic and fundamental factors of the how REST API's should be designed and structured.

---

class: center, middle

.title[RPC]

???

- What we really want is the the convenience of local function calls... but executed in distributed manner.

---

.center[![Issue](/img/grpc-logo.svg)]

<blockquote><strong>A high performance, open-source universal RPC framework</strong></blockquote>

http://grpc.io/

https://github.com/grpc-ecosystem/awesome-grpc

???

- Originally a Google project internally called "Stubby"
- Open sourced, mainly developed by Google employees

---

# gRPC ?

- 1.0 'g' stands for ['gRPC'](https://github.com/grpc/grpc/tree/v1.0.x)
- 1.1 'g' stands for ['good'](https://github.com/grpc/grpc/tree/v1.1.x)
- 1.2 'g' stands for ['green'](https://github.com/grpc/grpc/tree/v1.2.x)
- 1.3 'g' stands for ['gentle'](https://github.com/grpc/grpc/tree/v1.3.x)
- 1.4 'g' stands for ['gregarious'](https://github.com/grpc/grpc/tree/v1.4.x)
- 1.6 'g' stands for ['garcia'](https://github.com/grpc/grpc/tree/v1.6.x)
- 1.7 'g' stands for ['gambit'](https://github.com/grpc/grpc/tree/v1.7.x)
- 1.8 'g' stands for ['generous'](https://github.com/grpc/grpc/tree/v1.8.x)
- 1.9 'g' stands for ['glossy'](https://github.com/grpc/grpc/tree/v1.9.x)
- 1.10 'g' stands for ['glamorous'](https://github.com/grpc/grpc/tree/v1.10.x)
- 1.11 'g' stands for ['gorgeous'](https://github.com/grpc/grpc/tree/v1.11.x)
- ... https://github.com/grpc/grpc/blob/master/doc/g_stands_for.md

???

- What does gRPC stand for?
- "g" stands for something different in every version 

---

# SERVICE DEFINITION

```proto
// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```

???

- This is a Protocol Buffer definition file
- It's Interface Description Language used to describe types and services
- It's also an efficient binary serialization format
- This is simple and concise
- Just by reading it we can understand the general idea of this service and the API contract
- `protoc` compiles `.proto` file to generate language-specific code
- `protoc` compiler with plugin support
- Plugins to extend functionality

---

# CODE GENERATION

```sh
$ protoc -I helloworld/ \ 
  helloworld/helloworld.proto \
  --go_out=plugins=grpc:helloworld
```

```sh
$ npm install -g grpc-tools
$ grpc_tools_node_protoc \
  --js_out=import_style=commonjs,binary:../codegen/ \
  --grpc_out=../codegen \
  --plugin=protoc-gen-grpc=grpc_node_plugin \
  helloworld.proto
```

???

- Install `protoc` compiler
- Compile `.proto` file to generate language-specific code
- Generated code is not to be edited
- Generated code is not necessarily idiomatic for the target language

---

# MECHANISM

<br>

.center[![Architecture](/img/grpc-arch.svg)]

???

- Generated code provides client libraries and server stubs
- gRPC libraries provides RPC Mechanisms
- Unary - simple client request & server response
- Streaming request and single server response
- Single client request and streaming response 
- Duplex / bi-directional streaming
- Streaming allows for no / easier pagination mechanisms without need for a cursor or page number

---

# DETAILS

- HTTP/2
- RPC using Protocol Buffers (or JSON)
- Forwards & backwards compatible on the wire
- Mobile: Android and Objective-C, Experimental Swift
- Polyglot: C++, Go, Java, Ruby, Node.js, Python, C#, PHP, Dart

???

- HTTP2 is binary protocol that is fully multiplexed, instead of ordered and blocking
- Multiple requests can be serviced at the same time in one long-lived connection
- RPC using Protocol Buffers (or JSON)
- Forwards & backwards compatible on the wire
- gRPC core implementations in C++, Go and Java. Most others based on C++ core.

---

# SERVER - GO

```go
type server struct{}

func (s *server) SayHello(ctx context.Context, 
    in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	grpcServer := grpc.NewServer()
  pb.RegisterGreeterServer(grpcServer, &server{})
  reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
```

???

- SayHello is the implementation of our service
- The code in main() is a bit of boiler plate
- Reflection is for introspection. 
  - The service can explain what services and methods this grpc server has
  - Client can connect and build the client without knowing what lives on the server

---

# CLIENT - GO

```go
func main() {
	conn, _ := grpc.Dial("localhost:50051",
		grpc.WithInsecure())
	defer conn.Close()
*	c := pb.NewGreeterClient(conn)
*	ctx, cancel := context.WithTimeout(context.Background(),
*		10*time.Second)
	defer cancel()
*	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
```

???

- We create a connection
- We create a client from our generated code
- And we call our function to communicate with the server
- One interesting thing to notice is we get timeout support for free

---

# SERVER - NODE.JS

```js
const PROTO_PATH = __dirname + './protos/helloworld.proto'
const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')
const packageDefinition = protoLoader.loadSync(PROTO_PATH)
const proto = 
  grpc.loadPackageDefinition(packageDefinition).helloworld

function sayHello(call, callback) {
  callback(null, { message: 'Hello ' + call.request.name })
}

function main() {
  const server = new grpc.Server()
  server.addService(proto.Greeter.service, 
    { sayHello: sayHello })
  server.bind('0.0.0.0:50051',
    grpc.ServerCredentials.createInsecure())
  server.start()
}

main()
```

???

- Same idea for Node.js
- In case of Node.js we have the option of using dynamic generation of our client and server code, which is sometimes more convenient

---

# CLIENT - NODE.JS

```js
const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');
const packageDefinition = protoLoader.loadSync(PROTO_PATH);
const proto = 
  grpc.loadPackageDefinition(packageDefinition).helloworld;

function main() {
*  const client = new proto.Greeter(
*   'localhost:50051', grpc.credentials.createInsecure());

* const deadline = 
*   new Date().setSeconds(new Date().getSeconds() + 5)

* client.sayHello({ name: 'world' }, { deadline }, 
*   (err, response) => {
      console.log('Greeting: ', response.message);
*  });
}

main();
```

???

- Similar code for Node.js client side

---

# MORE COMPLEX

```proto
syntax = "proto3";

package greeter;

service Greeter {
    rpc SayHello (HelloReq) returns (HelloRes) {}
    rpc SayHellos (HelloReq) returns (stream HelloRes) {}
    rpc GreetMany (stream HelloReq) returns (HelloRes) {}
    rpc GreetChat (stream HelloReq) returns (stream HelloRes) {}
}

message HelloReq {
    string name = 1;
}

message HelloRes {
    string message = 1;
}
```

???

- Here we see different types in Protocol Buffer definition
- And an example of streaming requests

---

# BIDI STREAMING - SERVER

```go
func (s *server) GreetChat(stream pb.Greeter_GreetChatServer)
error {
	for {
*		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
*		msg := &pb.HelloRes{Message: "Hello " + in.Name}
*		if err := stream.Send(msg); err != nil {
			return err
		}
	}
}
```

???

- Putting both together we essentially have a chat where both client and server our utilizing a stream to send and receive messages

---

# BIDI STREAMING - CLIENT

```js
*call = client.greetChat()
const NAMES = ['Bob', 'Kate', 'Jim', 'Sara']
let n = 0
const timer = setInterval(() => {
  if (n < NAMES.length) {
*    call.write({ name: NAMES[n] })
    n++
  } else {
    clearInterval(timer)
*    call.end()
  }
}, 200)

*call.on('data',
*  ({ message }) => console.log('Greeting:', message))

*call.on('end', () => console.log('done'))
```

???

- Just an example of calling our BiDi method in Node.js
- Note that there are no parameters as we send them via stream operation

---

# METADATA - CLIENT

```go
conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
defer conn.Close()

c := pb.NewGreeterClient(conn)

ctx, cancel := context.WithTimeout(context.Background(),
  time.Second)
defer cancel()

*ctx = metadata.AppendToOutgoingContext(
*  ctx, "token", "xyz", "request-id", "123")

res, _ := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})

log.Printf("Greeting: %s", res.Message)
```

???

- gRPC "Headers"
- We can use metadata to send additional contextual information about calls
- This can be used for authentication, request id, or tracing
- All implementations should support it
- With Node.js just add it as an additional parameter

---

# METADATA - SERVER

```js
function sayHello(call, callback) {
*  const metadata = call.metadata.getMap()
  for (const k in metadata) {
    console.log(`${k}: ${metadata[k]}`)
  }

  callback(null, { message: 'Hello ' + call.request.name })
}
```

```sh
user-agent: grpc-go/1.16.0
token: xyz
request-id: 123
```

???

- Example of getting metadata on server in Node.js
- gRPC automatically adds user-agent metadata field
- Similarly Go provides a utility function to get metadata from context

---

# TOOLING - CLI

```sh
$ grpc_cli ls localhost:50051
helloworld.Greeter
grpc.reflection.v1alpha.ServerReflection

$ grpc_cli ls localhost:50051 helloworld.Greeter -l
filename: helloworld.proto
package: helloworld;
service Greeter {
  rpc SayHello(helloworld.HelloRequest) returns (helloworld.HelloReply) {}
}

$ grpc_cli call localhost:50051 SayHello 'name: "john"'
connecting to localhost:50051
message: "Hello john"

Rpc succeeded with OK status
```

???

- `grpc_cli` is the official command line tool
- There are other options such as grpcurl

---

# WEB SUPPORT ?

.center[![gRPC-Web](/img/grpc-web.png)]

```sh
protoc helloworld.proto \
  --js_out=import_style=commonjs:./codegen \
  --grpc-web_out=import_style=commonjs:./codegen
```

???

- Kinda
- We generate types like normal using `protoc`, this generates types and client to be used in our browser application in React, Vue or Angular and others
- In our web application we use the web client and the generated types to communicate with the server and we can have our Protocol Buffer types in the whole stack
- To enable browser support the protocol implemented here is actually slightly different than what gRPC uses in the rest of the ecosystem. Therefore we need a middle proxy to do translation of requests from the browser to our gRPC services.
- By default the gRPC project recomments using the Envoy proxy to accomplish this, and they do provide example configuration files to set this all up.
- In case you don't know Envoy is an open source proxy server implemented by Lyft.
- Some of you may be following re:Invent announcements this week, and the new AWS service mash product actually uses this Envoy proxy as well.
- It's important to note that right now only Unary and server streaming calls are supported
- Nginx can also work in this situation but the official recommendation is the Envoy proxy

---

# HTTP / JSON + gRPC

```proto
package helloworld;

*import "google/api/annotations.proto";

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {
*    option (google.api.http) = {
*      get: "/say"
*    };
  }
}
```

- Use [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) to build a REST API service
- Use Envoy's gRPC-JSON transcoder filter

???

- Sometimes we want to suport HTTP / JSON along with gRPC
- In that case we can use a `protoc` plugin to annotate our service definition with REST API details
- `grpc-gateway` plugin can be used to generate code for a proxy service that will handle our HTTP/JSON requests
- `grpc-gateway` can be used to generate swagger definition and documentation

- Alternatively Envoy's gRPC-JSON transcoder filter can be used to allow a RESTful JSON API client to send requests to Envoy over HTTP and get proxied to a gRPC service. 

- There are intricacies with how both handle streaming and neither solution provide BiDi streaming support

---

# DEALING WITH CHANGE

- Name of fields are less important than field numbers
- Do not change the type or number of a field
- Adding fields is safe
- Deprecate a field before removal
- Reuse a field number if absolutely sure
- Be aware of the default values for the data types
- If you need a version set it in package name 
  * Ex: `company.service.v1`

???

- Field name can be changed and will not effect serialization
- We do not want to change the type of number of a field
- Adding fields is safe
- When remving a field we should deprecate it first and there is a mechanism for doing that
- One we remove a field we do not want to reuse that field number of field name until we're absolutely sure that the old version of it is not used by any clients
- When deserializing data gRPC always assigns a default value to any fields present in the message definition and not present in the payload.
- boolean will get defaulted to fale, string to empty string, int to 0
- so we need to be aware of that sideeffect when implementing the login in our API's
---

# CHANGE - ADD

```proto
// v2
message HelloRequest {
  string name = 1;
  bool capitalize = 2;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```

???

Generally servers would be released before clients.

**Client v1 <-> Server v2**
- client will not know about capitalize, and it will default to `false`
- client will get just the message

**Client v2 <-> Server v1**
- client may set the flag to `true`
- server will not know about `capitalize` and will return old reply format
- client will get a message

---

# CHANGE - REMOVAL

```proto
// v3
message HelloRequest {
*  string name = 1 [deprecated=true];
  bool capitalize = 2;
*  string first_name = 3;
*  string last_name = 4;
}
```

```proto
// v4
message HelloRequest {
*  reserved 1;
*  reserved "name";

  bool reverse = 2;
  string first_name = 3;
  string last_name = 4;
}
```

???

- Do not remove a field number immediately
- First mark it deprecated and allow for clients to update.
- Deprecating a field may have meaningul result in code generated by protoc in some languages. Foe example for Java it will use the @Deprecated annotion.
- Keep server logic in place
- Once ready to remove, remove it.
- We can reserve a field number and name to prohibit developers from accidentally reusing it in future changes to the message
- The protocol buffer compiler will complain if any future users try to use these field identifiers. 
- Once enough time has passed that you know there will be no binary serialization of original field in the wild, remove reserved

---

# WORKFLOW & DESIGN

<img src="/img/googleapis.png" alt="Issues" height="400">

https://cloud.google.com/apis/design/

???

- Monorepo for all type and service definitions
- Review API changes with normal PR process
- Automatically test compilation, linting, etc...
- Services version control generated code as needed
- Services should be small and concise
- Do one thing and one thing well
- gRPC does not solve the problem of properly architecting and designing API's
- Be concise and consistent
- Have a style for consistent API design
  * Ex: `Update` vs `Save` etc...
  * https://cloud.google.com/apis/design/

---

# CHALLENGES

- Load Balancing
- Browser Support
- Debuggability
- Documentation
- Poor feature parity between languages
  * Ex: Interceptors / middleware
- Inconsistency between languages
  * Ex: timeout vs. deadline

???

- Load balancing is an improving issue, Envoy, Linkerd and Nginx can all support gRPC now
- gRPC-Web was generally available at the end of October
- The fact that we are dealing with binary data means we can't just inspect data across the wire. A new tool called Channelz can be used to gather comprehensive runtime info about connections in gRPC. It is designed to help debug live programs.
- gRPC documentation beyond the basic tutorial is non-existent and / or scattered and is lacking in more detailed reference and guidance on more advanced topics and examples
- There is inconsistent feature set between languages. For example Java and Go both have client and server interceptors, while client side interceptors were only recently added to Node.js and there is no server side middleware in Node.js at all. There are 3rd party modules to address this issue.
- Inconsistency in semantics between languages. 
  * timeout in Go vs. deadline in Node.js

---

<img src="/img/ghmwissues.png" alt="Issues">

???

- Can we have server-side middleware please?
- All github issues for people asking for Node.js server middleware
- There are other open source libraries and frameworks that expose this functionality

---

# BUT WHAT ABOUT?

- SOAP / WSDL
- Swagger & JSON Schema
- Thrift
- MessagePack
- Twirp
- GraphQL

???

**SOAP / WSDL**
- Tied to XML (protobuf is pluggable)
- Unnecessarily complex and inflexible with regards to compatibility
- No Streaming

**Swagger**
- It is machine readable
- Lots of tooling
- Tied to HTTP/JSON, Performance issues and no streaming
- Very verbose and cumbersome, a single definition takes pages of code

**Thrift**
- Started out as a promising serialization format similar to Protocol Buffers
- Failed to build a supported RPC system out of it due to level of effort required.

**MessagePack**
- Pretty flexible and well supported binary serialization format
- There is RPC on top but poor for building well designed and maintainable contracts and APIs

**Twirp**
- A simpler gRPC from Twitch that works with HTTP/1
- Good alternative if you're not comfortable with the hard HTTP2 requirement

**GraphQL**
- Interesting option for clients / frontends to query exactly the data they need
- Human readable and schema-based with types
- Still works over HTTP and no streaming
- Perhaps not ideal for service <-> service communication

**Future**

> A furious bout of language and protocol design takes place and a new distributed computing paradigm is announced that is compliant with the latest programming model. After several years, the percentage of distributed applications is discovered not to have increased significantly, and the cycle begins anew. - Waldo et al

---

# REJOINER

http://rejoiner.io

<img src="/img/rejoiner.png" alt="Rejoiner" width="580">

???

- Interesting project to expose gRPC API's via a uniform GraphQL API

---

# GRPC USERS

- Google - Google Cloud Services APIs and internally other products
- Square - Most internal RPC
- Lyft
- Netflix
- CoreOS - etcd v3 API is entirely gRPC
- Coockroach Labs
- Bugsnag
- VSCO
- Namely
- and others...

???

- Google - PubSub, Speech Rec
- Netflix heavily uses Java and has been active in RFPs for Node.js

---

# SHOULD YOU USE IT?

.center[<img src="/img/itdepends.png" alt="it depends" width="420">]

???

- Like everything technical... it depends. 
- It depends on your needs and requirements and context. It's important to do your own evaluation and research and really determine what compromises you're OK with for what benefits when making any kind of decision on technical choices.
- I believe gRPC is a pretty good option for an RPC mechanism, at least for inter-service communication in microservices architecture.

---

# THANK YOU!

Bojan Djurkovic

Lead Software Engineer

https://github.com/bojand

@bojantweets

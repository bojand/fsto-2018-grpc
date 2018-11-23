# helloworld.proto - Go
protoc -I protos/ protos/helloworld.proto --go_out=plugins=grpc:codegen 

# helloworld.proto - Node
grpc_tools_node_protoc -I protos \
--js_out=import_style=commonjs,binary:codegen \
--grpc_out=codegen \
--plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` \
helloworld.proto

# greeter.proto - Go
protoc -I protos/ protos/greeter.proto --go_out=plugins=grpc:codegen 

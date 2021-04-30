# Useful notes
# 
# https://github.com/grpc/grpc-go/tree/v1.37.0/examples/helloworld

# Setup
brew install protobuf
(cd && go get github.com/gogo/protobuf/protoc-gen-gofast)

curl -sSL https://zipkin.io/quickstart.sh | bash -s
java -jar zipkin.jar
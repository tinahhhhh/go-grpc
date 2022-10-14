# build command : docker build . -t server
# run command : docker run -it server
FROM golang:latest

# Install grpc
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

# Install protoc and zip system library
RUN apt-get update && apt-get install -y zip git && \
    mkdir /opt/protoc && cd /opt/protoc && wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protoc-3.7.0-linux-x86_64.zip && \
    unzip protoc-3.7.0-linux-x86_64.zip

ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin

# Git clone the project and generate stub files
RUN git clone https://github.com/tinahhhhh/go-grpc.git
RUN cd go-grpc && protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    spam/spam.proto

ENTRYPOINT cd go-grpc/server && go run main.go -url 

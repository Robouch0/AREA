FROM golang:1.23.4-bullseye

WORKDIR /area

COPY . .

RUN apt-get update
RUN apt install -y protobuf-compiler

# Install protoc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN protoc --go_out=./protogen \
    --go_opt=paths=source_relative \
    --go-grpc_out=./protogen \
    --go-grpc_opt=paths=source_relative gRPC/proto/*.proto

RUN go mod download

RUN go build -o ./area-server

# Modify the env to the prod one
# Later do a more suitable switch from dev to prod (vice versa)
RUN mv .env.prod .env

EXPOSE 8080

CMD [ "./area-server" ]

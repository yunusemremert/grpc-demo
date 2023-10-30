### Dependencies

1. You will need `protoc` to generate proto buffer code.

```bash
$ brew install protobuf
$ protoc --version
```

2. Install the protocol compiler plugins for Go using the following commands:

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

3. Update your `PATH` so that the protoc compiler can find the plugins:

```bash
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

---

### Create Protoc File Command

1. Create folder for protobuf files

```bash
$ mkdir chat 
```

2. Paste the following code into the terminal screen or run the make command

```bash
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    chat/chat.proto
```

or command line

```bash
$ make generate_grpc_code
```

---

### Run the server and client service

```bash
$ go run server.go
```

and open a new  terminal screen

```bash
$ go run client.go --name=John Wick
```
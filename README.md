
### Dependencies

1. Install the protocol compiler plugins for Go using the following commands:

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

2. Update your `PATH` so that the protoc compiler can find the plugins:

```
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

---

### Create Protoc File Command

1. Create folder

```bash
$ mkdir chat 
```

2. Copy/Paste screen terminal panel

```
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    chat/chat.proto
```

or

```
make generate_grpc_code
```
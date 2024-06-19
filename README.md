# Go GRPC exploration

## Requirements

This section lists the requirements of the project, knowing that I work under linux, the examples will not necessarily be adapted to you.

Here are the links for all kinds of installations:\
[protoc](https://grpc.io/docs/protoc-installation/)\
[golang](https://go.dev/doc/install)

### Golang Installation

I used go 1.22.4 for this exploration.\
Exemple: 

```shell
# fetch pre-compiled golang sdk.
curl -LO https://go.dev/dl/go1.22.4.linux-amd64.tar.gz

# Tar the file under a directory of your choice.
tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz

# Update your environment's path variable.
export PATH="$PATH:/usr/local/go/bin"

# Check installation
golang version
```

### Protocol Buffer Compiler Installation

In my case i chosed pre-compiled binaries for linux as a zip file.\
Exemple: 

```shell
# fetch the file using commands such as the following.
curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v27.1/protoc-27.1-linux-x86_64.zip

# Unzip the file under $HOME/.local/protoc or a directory of your choice.
unzip protoc-27.1-linux-x86_64.zip -d $HOME/.local/protoc

# Update your environment's path variable to include the path to the protoc executable.
export PATH="$PATH:$HOME/.local/protoc/bin"

# Check installation
protoc --version
```

Also install the proto file generator for golang.

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest #1.34.2
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest #1.4

# Update your environment's path variable to include the path to the protoc executable.
export PATH="$PATH:$(go env GOPATH)/bin"

# Check installation
which protoc-gen-go
which protoc-gen-go-grpc
```

## Usage

### gen/proto

Here is the command generates the golang code enabling interaction with protobuf:
```shell
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/exploration.proto
```
This command creates two files.

**exploration_grpc.pb.go**\
which contains all the protocol buffer code to populate, serialize, and retrieve request and response message types.

**exploration.pb.go**\
which contains the following:
- An interface type for clients to call with the methods defined in the explorationService service.
- An interface type for servers to implement, also with the methods defined in the explorationService service.

Then launch the server then the client to see the interactions via protobuf !

```shell
go run server/server.go 
go run client/client.go 
```
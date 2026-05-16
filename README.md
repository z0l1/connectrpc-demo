### preparation steps

1. install go deps

```
go install github.com/bufbuild/buf/cmd/buf@latest
```

```
go get -tool google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

```
go get -tool connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
```

2. install node deps

```
npm install -g @bufbuild/protoc-gen-es @connectrpc/protoc-gen-connect-es @connectrpc/connect @connectrpc/validate
```

3. make `buf.yaml` via `buf config init`

4. make `buf.gen.yaml`

5. `scripts/bufgen.sh`

6. win

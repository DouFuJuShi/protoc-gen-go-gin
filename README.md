# protoc-gen-go-gin
## protoc-gen-go-gin
```shell
# install 
go install https://github.com/DouFuJuShi/protoc-gen-go-gin@latest

# about other proto files

# how to use
protoc \
    -I ./examples/api \
    --go_out ./examples/api --go_opt=paths=source_relative \
    --go-gin_out ./examples/api --go-gin_opt=paths=source_relative \
    ./examples/api/user/v1/user.proto
```

## About protoc-go-inject-tag
[protoc-go-inject-tag](https://github.com/favadi/protoc-go-inject-tag)

```shell
# install
go install https://github.com/favadi/protoc-go-inject-tag@latest

# inject tags
protoc-go-inject-tag -input=user.pb.go
```

```protobuf
// user.proto
message UserRequest {
    // @gotags: form:"title" uri:"id"
    string UserId = 1;
}
```

```go
// user.pb.go
// see UserId's struct tags
type UserRequest struct {
    state         protoimpl.MessageState
    sizeCache     protoimpl.SizeCache
    unknownFields protoimpl.UnknownFields
    
    // @gotags: form:"title" uri:"id"
    UserId string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty" form:"title" uri:"id"`
}
```

## Reference
[proto options](https://protobuf.dev/programming-guides/proto3/#options)   

[protocolbuffers/protobuf-go/](https://pkg.go.dev/google.golang.org/protobuf/compiler/protogen)    

[kratos/protoc-gen-go-http](https://github.com/go-kratos/kratos/tree/main/cmd/protoc-gen-go-http)    

[mohuishou/protoc-gen-go-gin](https://github.com/mohuishou/protoc-gen-go-gin)    

[pbgo: 基于 Protobuf 的框架](https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-07-pbgo.html)

[如何自定义 protoc 插件](https://yusank.github.io/posts/go-protoc-http/)

TKS！！！
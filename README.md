# protoc-gen-go-gin
## 说明
### 主要利用proto 的 Extensions 和 Custom Options
1. [Extensions](https://protobuf.dev/programming-guides/proto2/#extensions)
```protobuff
// example
// file kittens/video_ext.proto

import "kittens/video.proto";
import "media/user_content.proto";

package kittens;

// This extension allows kitten videos in a media.UserContent message.
extend media.UserContent {
  // Video is a message imported from kittens/video.proto
  repeated Video kitten_videos = 126;
}

// file media/user_content.proto

package media;

// A container message to hold stuff that a user has created.
message UserContent {
  extensions 100 to 199;
}
```

2. [Custom Options](https://protobuf.dev/programming-guides/proto2/#customoptions)      

Protocol Buffers允许定义和使用自己的选项。请注意，这是大多数人不需要的高级功能。
由于选项是由 google/protobuf/descriptor.proto 中定义的消息定义的（如 FileOptions 或 FieldOptions），    
定义您自己的选项只是扩展这些消息的问题。例如：

```protobuf
import "google/protobuf/descriptor.proto";

extend google.protobuf.FileOptions {
  optional string my_file_option = 50000;
}
extend google.protobuf.MessageOptions {
  optional int32 my_message_option = 50001;
}
extend google.protobuf.FieldOptions {
  optional float my_field_option = 50002;
}
extend google.protobuf.OneofOptions {
  optional int64 my_oneof_option = 50003;
}
extend google.protobuf.EnumOptions {
  optional bool my_enum_option = 50004;
}
extend google.protobuf.EnumValueOptions {
  optional uint32 my_enum_value_option = 50005;
}
extend google.protobuf.ServiceOptions {
  optional MyEnum my_service_option = 50006;
}
extend google.protobuf.MethodOptions {
  optional MyMessage my_method_option = 50007;
}

option (my_file_option) = "Hello world!";

message MyMessage {
  option (my_message_option) = 1234;

  optional int32 foo = 1 [(my_field_option) = 4.5];
  optional string bar = 2;
  oneof qux {
    option (my_oneof_option) = 42;

    string quux = 3;
  }
}

enum MyEnum {
  option (my_enum_option) = true;

  FOO = 1 [(my_enum_value_option) = 321];
  BAR = 2;
}

message RequestType {}
message ResponseType {}

service MyService {
  option (my_service_option) = FOO;

  rpc MyMethod(RequestType) returns(ResponseType) {
    // Note:  my_method_option has type MyMessage.  We can set each field
    //   within it using a separate "option" line.
    option (my_method_option).foo = 567;
    option (my_method_option).bar = "Some string";
  }

  rpc OtherMethod(RequestType) returns(ResponseType) {
    option (my_method_option) = {
      foo: 567,
      bar: "Some string"
    };
  }
}

// ---------------------------------------------------------
// 请注意，如果您想在定义它的包之外的包中使用自定义选项，则必须在选项名称前加上包名称前缀，就像您在类型名称中所做的那样。例如
// foo.proto
import "google/protobuf/descriptor.proto";
package foo;
extend google.protobuf.MessageOptions {
  optional string my_option = 51234;
}

// bar.proto
import "foo.proto";
package bar;
message OtherMessage {
  option (foo.my_option) = "Hello world!";
}
```

3. 自定义http Option   
3.1 利用google预定义的Option   https://github.com/googleapis/googleapis   
下载以下两个文件到一个目录下，在执行protoc时候通过-I指定目录，能够使其他proto文件import   
[google/api/annotations.proto](https://github.com/googleapis/googleapis/blob/master/google/api/annotations.proto)   
[google/api/http.proto](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto)   

两个文件proto通过protoc的go代码 在 google.golang.org/genproto/googleapis/api/annotations    

3.2 完全自定义Option [参考custom/http.proto]

```shell
// 运行以下 生成http.pb.go 这个文件需要能够后期被引用到
// 如果没安装protoc-gen-go插件 用go install安装
// go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
protoc --go_out .  --go_opt=paths=source_relative ./custom/http.proto
```

4. 编写插件
插件库 https://pkg.go.dev/google.golang.org/protobuf/compiler/protogen

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
protoc-go-inject-tag -input=hello.pb.go
```

```protobuf
// user.proto
message HelloRequest {
    // @gotags: form:"title" uri:"id"
    string UserId = 1;
}
```

```go
// user.pb.go
// see UserId's struct tags
type HelloRequest struct {
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
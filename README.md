# GO

## 基础语法

### 并发

#### 1. goroutine

```go
package main

import (
	"fmt"
	"sync"
)

//go 语言协程是一种轻量级线程，GO运行时调度

func Test(n int) {
	defer wg.Done() //相当于wg.Add(-1)
	fmt.Println(n)
}

//WaitGroup 会阻塞主线程 等所有的goroutine 执行完成
var wg sync.WaitGroup

func main() {
	fmt.Println("并发编程")
	for i := 0; i < 1000; i++ {
		wg.Add(1) //添加或减少goroutine 的数量
		go Test(i)
	}
	wg.Wait() //执行阻塞 等待所有协程执行完毕

}

```



## gin 框架

## gPRC

### protobuf

#### 1. protobuf 基本格式

```protobuf
syntax = "proto3"; //表示使用proto3语法 如果没有指定，编译器会使用proto2
option go_package = ".;proto"; //生成的go文件路径以及包名 .表示当前目录 proto表示包名

//service 关键字定义服务
service Greeter {
  rpc GetStream(StreamReqData) returns (stream StreamResData); //服务端流模式
  rpc PostStream(stream StreamReqData) returns (StreamResData); //客户端流模式
  rpc AllStream(stream StreamReqData)  returns (stream StreamResData); //双向流模式
}

message StreamReqData{
  string  data = 1; // 1是编号
}

message StreamResData {
  string data = 1;
}
```

[参考资料](https://lixiangyun.gitbook.io/protobuf3/)

#### 2.protobuf编译成go文件

 1. 下载安装`Protobuf Buffers`编译器 `https://github.com/protocolbuffers/protobuf`

 2. 解压缩设置，设置环境变量(path)指向解压后 的`bin`目录

 3. goland 中安装`proto`插件

 4. 执行编译命令(可以把命令做成脚本直接运行即可)

    ```shell
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./helloworld.proto // ./helloworld.proto 表示当前目录下的helloworld.proto文件
    ```

    

#### 3.proto文件中import另一个proto文件

将`com.proto`中的部分`message` 导入到`hi.proto`

```protobuf
//com.proto
syntax = "proto3";

message Empty{

}

message Pong{
  string id = 1;
}
```



```protobuf
//hi.proto
syntax  = "proto3";
import "com.proto";  //导入com.proto
import "google/protobuf/empty.proto"; //导入自带的empty
option go_package = ".;proto";

service SendHi{
  rpc Ping(Empty) returns (Pong); //使用com.proto中的message
  rpc PingGoogle(google.protobuf.Empty) returns (Pong);//使用自带的empty
  
}
```



#### 4.嵌套的message 对象


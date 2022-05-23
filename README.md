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



#### 2. channel

Go语言中的通道(channel)是一种特殊的类型。通道像一个传送带或者队列，总是循环先进先出的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。



#####  channel类型

`channel`是一种引用类型，声明通道类型的格式如下

```go
var 变量 chan 元素类型
```

打个栗子：

```go
var ch1 chan int //表示声明一个传递整型的通道
var ch2 chan bool //表示声明一个传递布尔类型的通道
```



##### 创建channel

```go
var ch chan int //channel 是引用类型，空值是nil
fmt.Println(ch) //<nil> 声明的channel 需要使用make初始化后才能使用

//使用make创建channel
make(chan 元素类型,[缓冲大小])
//打个栗子
ch1 :=make(chan int)
ch2 :=make(chan bool)

```

##### channel类型

1. 有缓冲通道
2. 无缓冲通道(阻塞通道，同步通道) 必须要有接收方

```go
func main() {
	//无缓冲通道栗子
	ch := make(chan int)
	go rev(ch) //启用goroutine 从通道接受值
	ch <- 10
	close(ch)

	//有缓冲通道
	ch2 := make(chan int, 1) //创建一个缓冲区容量为1的缓冲通道
	ch2 <- 10
	fmt.Println("发送成功")

	//	for range 从通道中循环取值
	rangechannel()

}

func rev(c chan int) {
	ret := <-c
	fmt.Println("接受成功", ret)
}

//for range 从通道循环取值

func rangechannel() {
	ch := make(chan int)
	ch1 := make(chan int)

	//向通道中发送值
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	//	接受值

	go func() {
		for {
			i, ok := <-ch //通道关闭后再取值 ok = false
			if !ok {
				break
			}

			ch1 <- i * i

		}
		close(ch1)
	}()

	//主goroutine中从ch1中接受打印值
	for v := range ch1 { //通道关闭后会推出for range循环
		fmt.Println(v)
	}
}

```



##### 单向通道

`chan<- int`是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行接收操作。
`<-chan in`t是一个只读单向通道（只能从其读取int类型值），可以对其执行接收操作但是不能执行发送操作。

打个栗子：

```go
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go counter(ch1)      //数据写入ch1
	go squarer(ch2, ch1) //把ch1中的数据 写入ch2
	printer(ch2)         //输出 ch2中的数据

}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)

}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}

}
```



## gin 框架

### gin路由组

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/hh") //路由组
	{
		v1.POST("/hi", Login)
		v1.GET("/ha", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"HELLO": "WORLD",
			})
		})
	}

	r.Run(":9090")

}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
```



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





## 常用库

###  Viper

Viper是适用于Go应用程序（包括`Twelve-Factor App`）的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式。它支持以下特性：

- 设置默认值
- 从`JSON`、`TOML`、`YAML`、`HCL`、`envfile`和`Java properties`格式的配置文件读取配置信息
- 实时监控和重新读取配置文件（可选）
- 从环境变量中读取
- 从远程配置系统（etcd或Consul）读取并监控配置变化
- 从命令行参数读取配置
- 从buffer读取配置
- 显式配置值



#### 1. Viper 的安装

```go
go get github.com/spf13/viper
```



#### 2.Viper 读取配置文件



```go
viper.SetConfigFile("./config.yaml") // 指定配置文件路径
viper.SetConfigName("config") // 配置文件名称(无扩展名)
viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
viper.AddConfigPath("/etc/appname/")   // 查找配置文件所在的路径
viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
viper.AddConfigPath(".")               // 还可以在工作目录中查找配置
err := viper.ReadInConfig() // 查找并读取配置文件
if err != nil { // 处理读取配置文件的错误
	panic(fmt.Errorf("Fatal error config file: %s \n", err))
}
```



打个栗子：

```go
// 读取yaml配置文件 并将配置文件映射到单层struct
type ServerConfig struct {
	ServiceName string `mapstructure:"name"`
	Port        int    `mapstructure:"port"`
}

// 读取yaml 文件配置
func readYaml() {
	v := viper.New()

	// 指定文件路径
	v.SetConfigFile(".\\config\\config.yaml")

	// 查找阅读配置文件并处理错误信息
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	// 将配置文件映射到struct
	sc := ServerConfig{}
	if err := v.Unmarshal(&sc); err != nil {
		panic(err)
	}
	fmt.Println(sc)

	// 直接输出配置文件信息
	name := v.Get("name")
	port := v.Get("port")
	fmt.Println(name)
	fmt.Println(port)
}

```



在打个栗子:

```go
// 读取yaml 多层配置信息并映射到嵌套结构体

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type allServerConfig struct {
	ServiceName string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	MysqlInfo   MysqlConfig `mapstructure:"mysql"`
}

func readMoreYaml() {
	v := viper.New()

	// 指定文件路径
	v.SetConfigFile(".\\config\\config.yaml")

	// 查找阅读配置文件并处理错误信息
	if err := v.ReadInConfig(); err != nil {
		panic(err)

	}

	// 将配置文件映射到struct
	allsc := allServerConfig{}
	if err := v.Unmarshal(&allsc); err != nil {
		panic(err)
	}
	fmt.Println(allsc)

}
```



#### 3.Viper读取环境变量

```go
// 将线上线下配置文件进行隔离
// 不用改任何代码 而且线上线下的配置文件能隔离开
// 通过环境变量实现此功能

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func chooseConfig() {
	debug := GetEnvInfo("CONFIG_DEBUG")
	configFileName := ".\\config\\config.yaml"

	if !debug {
		configFileName = ".\\config\\config_local.yaml"
	}

	v := viper.New()

	// 指定文件路径
	v.SetConfigFile(configFileName)

	// 查找阅读配置文件并处理错误信息
	if err := v.ReadInConfig(); err != nil {
		panic(err)

	}

	// 将配置文件映射到struct
	allsc := allServerConfig{}
	if err := v.Unmarshal(&allsc); err != nil {
		panic(err)
	}
	fmt.Println(allsc)

	// viper - 动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config uploadfile changed:", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&allsc)

		fmt.Println(allsc)

	})

	time.Sleep(time.Second * 200)
}

```





### Zap日志库

[Zap](https://github.com/uber-go/zap)是非常快的、结构化的，分日志级别的Go日志库。



#### 1. Zap安装

```go
go get -u go.uber.org/zap
```



#### 2.配置Logger

定制化Logger,将日志写入文件而不是终端。

我们将使用`zap.New(…)`方法来手动传递所有配置，而不是使用像`zap.NewProduction()`这样的预置方法来创建logger。

```go
func New(core zapcore.Core, options ...Option) *Logger
```

`zapcore.Core`需要三个配置——`Encoder`，`WriteSyncer`，`LogLevel`。

1.**Encoder**:编码器(如何写入日志)。我们将使用开箱即用的`NewJSONEncoder()`，并使用预先设置的`ProductionEncoderConfig()`。

```go
   zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
```

2.**WriterSyncer** ：指定日志将写到哪里去。我们使用`zapcore.AddSync()`函数并且将打开的文件句柄传进去。

```go
   file, _ := os.Create("./test.log")
   writeSyncer := zapcore.AddSync(file)
```

3.**Log Level**：哪种级别的日志将被写入。

我们将修改上述部分中的Logger代码，并重写`InitLogger()`方法。其余方法—`main()` /`SimpleHttpGet()`保持不变。



打个栗子：

```go
var logger *zap.Logger

func main() {
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("www.google.com")
	simpleHttpGet("https://baidu.com")

}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller()) //AddCaller() 函数调用方信息

}

func getEncoder() zapcore.Encoder {
	//返回人能看懂的时间格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewJSONEncoder(encoderConfig)//设置日志格式为Json格式 
    //若不想使用Json格式的日志，return zapcore.NewConsoleEncoder(encoderConfig)
}


//这里使用Lumberjack对日志进行切割，日志写入文件

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    100,   //M
		MaxBackups: 5,     //备份数量
		MaxAge:     30,    //保留旧文件最大天数
		Compress:   false, //是否压缩归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("success...", zap.String(
			"StatusCode", resp.Status), zap.String("url", url))
		resp.Body.Close()
	}
}

```



#### 3.Gin中使用Zap日志库

1. 首先使用根据Zap封装好的Logger中间件，替换原有的中间件。

```go
// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
```



打个栗子：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var logger *zap.Logger

func main() {
	InitLogger()

	r := gin.New()

	r.Use(GinLogger(logger), GinRecovery(logger, true))

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	r.Run()
}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller()) //AddCaller() 函数调用方信息

}

func getEncoder() zapcore.Encoder {
	//返回人能看懂的时间格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "ginlogger/test.log",
		MaxSize:    100,   //M
		MaxBackups: 5,     //备份数量
		MaxAge:     30,    //保留旧文件最大天数
		Compress:   false, //是否压缩归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
```



 

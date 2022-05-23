package main

//定制logger
// 将logger 写入文件

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

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

	return zapcore.NewJSONEncoder(encoderConfig)
}

//func getLogWriter() zapcore.WriteSyncer {
//	//file, _ := os.Create("./test.log")
//	file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744) //不覆盖原来的日志
//	return zapcore.AddSync(file)
//}

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

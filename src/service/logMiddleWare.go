package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

func LogMiddleWare() gin.HandlerFunc {
	var (
		logFilePath = "./log" //文件存储路径
		logFileName = "frontAPI.log"
	)
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开/写入文件失败", err)
		return nil
	}

	logger := logrus.New()

	logger.SetLevel(logrus.ErrorLevel)

	logger.SetOutput(file)
	logger.SetOutput(os.Stdout)

	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y-%m-%d-%H-%M.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		//rotatelogs.WithMaxAge(7*24*time.Hour), //以hour为单位的整数

		rotatelogs.WithRotationCount(10),
		rotatelogs.WithRotationSize(200*1024*1024),

		// 设置日志切割时间间隔(1天)
		//rotatelogs.WithRotationTime(24*time.Hour),
	)

	// hook机制的设置
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(writerMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return func(c *gin.Context) {
		c.Next()
		//请求方式
		method := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()
		// 打印日志
		logger.WithFields(logrus.Fields{
			"status_code": statusCode,
			"client_ip":   clientIP,
			"req_method":  method,
			"req_uri":     reqUrl,
		}).Info()
	}

}

package main

import (
	"fmt"
	"jaemon-go/protos/common"

	"jaemon-go/utils"

	logger "github.com/sirupsen/logrus"
	"os"
)

// 初始化信息
func init() {
	fmt.Println("初始化参数")
	// logrus 默认的日志输出有 time、level 和 msg 3个 Field，其中 time 可以不显示，方法如下：
	//logger.SetFormatter(&logger.TextFormatter{DisableTimestamp: true})

	// 设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logger.SetFormatter(&logger.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint: false,
	})
	// 设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	logger.SetOutput(os.Stdout)
	// 设置日志级别为 info 以上
	logger.SetLevel(logger.InfoLevel)
	// 显示行号等信息
	logger.SetReportCaller(true)
}

func main() {
	var name string = "answer"
	var data []byte = []byte(name)

	var response = &common.Response{
		Code:    10000,
		Message: "success",
		Result:  data,
	}

	fmt.Println(response.Message)

	fmt.Println(string(response.Result))

	fmt.Println()
	logger.WithFields(
		logger.Fields{
			"key": "jaemon",
		}).Info("Start Appliaction...")
	fmt.Println()

	utils.Execute()

	utils.ExecuteEcdsa()
}

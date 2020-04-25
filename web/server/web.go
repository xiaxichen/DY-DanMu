package main

import (
	"DY-DanMu/web/server/route"
	"fmt"
	Log "github.com/sirupsen/logrus"
	"os"
	"runtime"
)

func init() {
	// 设置日志格式为json格式
	//Log.SetFormatter(&Log.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	Log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	Log.SetLevel(Log.DebugLevel)
	// 设置行号
	Log.SetReportCaller(true)

	Log.SetFormatter(&Log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			//filename := path.Base(f.File)
			return fmt.Sprintf("%s", f.Function), ""
		},
	})
}

func main() {
	router := route.Router()
	Log.Fatal(router.Run(":8081"))
}

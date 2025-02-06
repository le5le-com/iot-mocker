package main

import (
	"fmt"
	"le5le/iot-mocker/config"
	"le5le/iot-mocker/mqttServer"
	"le5le/iot-mocker/router"
	"le5le/iot-mocker/utils"
	"os"
	"runtime"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// 处理panic
	defer func() {
		if r := recover(); r != nil {
			// 创建足够大的缓冲区存储堆栈信息
			buf := make([]byte, 1024)
			// Stack函数将堆栈信息写入缓冲区
			n := runtime.Stack(buf, false)
			// 打印堆栈信息
			fmt.Printf("Recover from panic: %v\n%s\n", r, buf[:n])
		}
	}()

	config.Init()

	// 设置日志
	zerolog.SetGlobalLevel(zerolog.Level(config.App.Log.Level))
	if config.App.Log.Filename == "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	} else {
		filepath := config.App.Log.Filename
		if filepath[0] != '/' {
			filepath = utils.CurrentPath + filepath
		}
		fmt.Printf("日志位置: %s \n", filepath)
		log.Logger = log.Output(&lumberjack.Logger{
			Filename:   filepath,
			MaxSize:    config.App.Log.MaxSize, // mb
			MaxBackups: config.App.Log.MaxBackups,
			MaxAge:     config.App.Log.MaxAge, // days
		})
	}

	// 最大cpu使用核心数
	if config.App.CPU > 0 {
		runtime.GOMAXPROCS(config.App.CPU)
	}

	mqttServer.Init()

	router.Listen()
}

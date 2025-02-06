package router

import (
	"le5le/iot-mocker/apis"
	"le5le/iot-mocker/config"
	"le5le/iot-mocker/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// Listen 监听路由
func Listen() {
	h := server.Default(server.WithHostPorts("0.0.0.0:" + config.App.Port))

	h.StaticFS("/", &app.FS{Root: utils.WebDir, IndexNames: []string{"index.html"}})

	apis.Route(h)

	h.Spin()

	utils.CtxCancel()
}

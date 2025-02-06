package apis

import (
	"le5le/iot-mocker/mqttServer"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Route(h *server.Hertz) {

	h.POST("/api/iot/devices", getDevices)
	h.POST("/api/iot/device/properties", getDeviceProperties)

	// 可视化应用连接mqtt的连接信息
	h.GET("/api/iot/app/mqtt", GetAppMqtt)

	// 订阅数据点。传设备和数据点或订阅Token，返回订阅Token
	h.POST("/api/iot/subscribe/properties", SubscribeProperty)
	// 取消订阅数据点。传订阅Token
	h.POST("/api/iot/unsubscribe/properties", UnsubscribeProperty)

	mqttServer.MakeDatas()
}

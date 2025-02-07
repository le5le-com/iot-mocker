package apis

import (
	"context"
	"fmt"
	"le5le/iot-mocker/config"
	"le5le/iot-mocker/mqttServer"
	"le5le/iot-mocker/utils"
	"sync/atomic"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func getDevices(ctx context.Context, c *app.RequestContext) {
	devices := make([]*Device, 5)
	for i := 1; i <= len(devices); i++ {
		strI := fmt.Sprintf("%0*d", 5, i)
		devices[i] = &Device{
			Id:    "00000001-0000-0000-0000-0000000" + strI,
			Name:  productNames[i],
			Sn:    fmt.Sprintf("le5le-iot-test-%d", i+1),
			Token: fmt.Sprintf("iot-test-%d", i+1),
		}
	}

	c.JSON(consts.StatusOK, utils.Json{"list": devices, "total": 5})
}

func getDeviceProperties(ctx context.Context, c *app.RequestContext) {
	params := &struct {
		DeviceId string `json:"deviceId"`
	}{}

	err := sonic.Unmarshal(c.Request.Body(), params)
	if err != nil {
		c.JSON(consts.StatusOK, utils.Json{"error": err.Error()})
		return
	}

	properties := CreateMqttProperties(params.DeviceId)
	c.JSON(consts.StatusOK, utils.Json{"list": properties, "total": len(properties)})
}

func GetAppMqtt(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.Json{
		"host":          config.App.MqttBroker.Host,
		"tlsServerName": config.App.MqttBroker.TlsServerName,
		"tcpPort":       config.App.MqttBroker.TcpPort,
		"tlsPort":       config.App.MqttBroker.TlsPort,
		"wsPort":        config.App.MqttBroker.WsPort,
		"wssPort":       config.App.MqttBroker.WssPort,
		"username":      config.App.MqttBroker.Username,
		"password":      config.App.MqttBroker.Password,
	})
}

func SubscribeProperty(ctx context.Context, c *app.RequestContext) {
	atomic.AddInt64(&mqttServer.ConnectedCount, 1)
	mqttServer.LastPing = time.Now().Unix()

	c.JSON(consts.StatusOK, utils.Json{"token": "test"})
}

func UnsubscribeProperty(ctx context.Context, c *app.RequestContext) {
	atomic.AddInt64(&mqttServer.ConnectedCount, -1)
}

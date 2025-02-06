package config

import (
	"os"

	"le5le/iot-mocker/utils"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

// App 全局配置文件实例
var App Config

// Init 读取全局配置文件
func Init() {
	configPath := utils.CurrentPath + "config.yaml"
	_, err := os.Stat(configPath)
	if err == nil || os.IsExist(err) {
		data, err := os.ReadFile(configPath)
		if err == nil {
			yaml.Unmarshal(data, &App)
		} else {
			App.Name = "乐吾乐数据模拟服务"
			App.CPU = 1
		}
	}

	if App.Port == "" {
		App.Port = "7100"
	}

	if App.Host == "" {
		App.Host = utils.GetLocalIP()
	}

	if App.MqttBroker.Host == "" {
		App.MqttBroker.Host = utils.GetLocalIP()
	}

	if App.MqttBroker.TcpPort == "" {
		App.MqttBroker.TcpPort = "5883"
	}
	if App.MqttBroker.TlsPort == "" {
		App.MqttBroker.TlsPort = "5884"
	}
	if App.MqttBroker.WsPort == "" {
		App.MqttBroker.WsPort = "5083"
	}
	if App.MqttBroker.WssPort == "" {
		App.MqttBroker.WssPort = "5084"
	}

	if App.Interval < 1 {
		App.Interval = 1
	}

	log.Info().Msgf("App config: %+v", App)
}

func Save() error {
	configPath := utils.CurrentPath + "config.yaml"

	dataFile, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer dataFile.Close()

	bytes, err := yaml.Marshal(App)
	if err != nil {
		return err
	}
	_, err = dataFile.Write(bytes)
	return err
}

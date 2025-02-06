package apis

import (
	"fmt"
)

var productNames []string = []string{"家居网关", "电网管理", "暖通设备", "仓储系统", "泵站控制", "机场设备", "停车管理"}

type Device struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Sn    string `json:"sn"`
	Token string `json:"token"`
}

type DeviceProperty struct {
	EntityId string `json:"entityId"` // 设备ID
	Name     string `json:"name,omitempty"`
	Key      string `json:"key"`
	Type     string `json:"type,omitempty"`
}

type SubscribePropertyParams struct {
	Token   string `json:"token"`
	Devices []struct {
		DeviceId   string   `json:"deviceId"`
		Token      string   `json:"token"`
		Properties []string `json:"properties"`
	} `json:"devices"`
}

func CreateMqttProperties(deviceId string) []*DeviceProperty {
	data := make([]*DeviceProperty, 0)

	data = append(data, &DeviceProperty{
		EntityId: deviceId,
		Key:      "temperature",
		Name:     "温度",
		Type:     "int",
	})
	data = append(data, &DeviceProperty{
		EntityId: deviceId,
		Key:      "humidity",
		Name:     "湿度",
		Type:     "int",
	})
	data = append(data, &DeviceProperty{
		EntityId: deviceId,
		Key:      "light",
		Name:     "亮度",
		Type:     "int",
	})
	data = append(data, &DeviceProperty{
		EntityId: deviceId,
		Key:      "co2",
		Name:     "二氧化碳",
		Type:     "int",
	})
	data = append(data, &DeviceProperty{
		EntityId: deviceId,
		Key:      "pm25",
		Name:     "PM25",
		Type:     "int",
	})
	data = append(data, &DeviceProperty{
		EntityId: deviceId,
		Key:      "aqi",
		Name:     "空气质量指数",
		Type:     "int",
	})
	data = append(data, &DeviceProperty{
		EntityId: deviceId,
		Key:      "AcVoltage",
		Name:     "交流电压",
		Type:     "double",
	})
	data = append(data, &DeviceProperty{
		EntityId: deviceId,
		Key:      "AcCurrent",
		Name:     "交流电流",
		Type:     "double",
	})
	data = append(data, &DeviceProperty{
		EntityId: deviceId,
		Key:      "DcVoltage",
		Name:     "直流电压",
		Type:     "double",
	})
	data = append(data, &DeviceProperty{
		EntityId: deviceId,
		Key:      "DcCurrent",
		Name:     "直流电流",
		Type:     "double",
	})
	data = append(data, &DeviceProperty{
		EntityId: deviceId,
		Key:      "longitude",
		Name:     "经度",
		Type:     "double",
	})
	data = append(data, &DeviceProperty{
		EntityId: deviceId,
		Key:      "latitude",
		Name:     "纬度",
		Type:     "double",
	})
	data = append(data, &DeviceProperty{
		EntityId: deviceId,
		Key:      "on",
		Name:     "开关",
		Type:     "bool",
	})
	data = append(data, &DeviceProperty{
		EntityId: deviceId,
		Key:      "state",
		Name:     "状态",
		Type:     "bool",
	})

	for i := 1; i <= 5; i++ {
		strI := fmt.Sprintf("%0*d", 3, i)
		data = append(data, &DeviceProperty{
			EntityId: deviceId,
			Key:      "int" + strI,
			Name:     "整数" + strI,
			Type:     "int",
		})
	}

	for i := 1; i <= 5; i++ {
		strI := fmt.Sprintf("%0*d", 3, i)
		data = append(data, &DeviceProperty{
			EntityId: deviceId,
			Key:      "float" + strI,
			Name:     "浮点数" + strI,
			Type:     "double",
		})
	}

	for i := 1; i <= 5; i++ {
		strI := fmt.Sprintf("%0*d", 3, i)
		data = append(data, &DeviceProperty{
			EntityId: deviceId,
			Key:      "bool" + strI,
			Name:     "布尔" + strI,
			Type:     "bool",
		})
	}

	for i := 1; i <= 5; i++ {
		strI := fmt.Sprintf("%0*d", 3, i)
		data = append(data, &DeviceProperty{
			EntityId: deviceId,
			Key:      "text" + strI,
			Name:     "文字" + strI,
			Type:     "text",
		})
	}

	return data
}

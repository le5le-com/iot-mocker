package mqttServer

import (
	"encoding/json"
	"fmt"
	"le5le/iot-mocker/utils"
)

var datas = []utils.Json{}
var dataJsons = []string{}

func MakeDatas() {
	datas = datas[:0]
	dataJsons = dataJsons[:0]

	for i := 0; i < 100; i++ {
		data := MakeData()
		datas = append(datas, data)
		jsonText, err := json.Marshal(data)
		if err == nil {
			dataJsons = append(dataJsons, string(jsonText))
		}
	}
}

func MakeData() utils.Json {
	data := utils.Json{}

	data["temperature"] = utils.GetRandInt(20, 30)
	data["humidity"] = utils.GetRandInt(40, 60)
	data["light"] = utils.GetRandInt(0, 100)
	data["co2"] = utils.GetRandInt(300, 500)
	data["pm25"] = utils.GetRandInt(0, 100)
	data["aqi"] = utils.GetRandInt(0, 100)

	data["AcVoltage"] = utils.GetRandFloat(210, 230)
	data["AcCurrent"] = utils.GetRandFloat(0, 1)

	data["DcVoltage"] = utils.GetRandFloat(3, 12)
	data["DcCurrent"] = utils.GetRandFloat(0, 1)

	data["longitude"] = utils.GetRandFloat(114, 128)
	data["latitude"] = utils.GetRandFloat(30, 32)

	data["on"] = utils.GetRandBool()
	data["state"] = utils.GetRandBool()

	for i := 0; i < 5; i++ {
		strI := fmt.Sprintf("%0*d", 3, i)

		data["int"+strI] = utils.GetRandInt(10*i, 10*i+10)
		data["float"+strI] = utils.GetRandFloat(10*float64(i), 10*float64(i)+10)
		data["bool"+strI] = utils.GetRandBool()
		data["text"+strI] = utils.GetRandString(5)
	}

	return data
}

var dataIndex = 0

func GetData() utils.Json {
	dataIndex++
	if dataIndex >= 100 {
		dataIndex = 0
	}
	return datas[dataIndex]
}

func GetDataJson() string {
	dataIndex++
	if dataIndex >= 100 {
		dataIndex = 0
	}
	return dataJsons[dataIndex]
}

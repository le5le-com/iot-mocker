package mqttServer

import (
	"encoding/json"
	"fmt"
	"le5le/iot-mocker/utils"
	"math/rand"
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

func getRandId() string {
	return fmt.Sprintf("00000001-0000-0000-0000-0000000%0*d", 5, rand.Intn(5))
}

func MakeData() utils.Json {
	data := utils.Json{}

	data[getRandId()+"#temperature"] = utils.GetRandInt(20, 30)
	data[getRandId()+"#humidity"] = utils.GetRandInt(40, 60)
	data[getRandId()+"#light"] = utils.GetRandInt(0, 100)
	data[getRandId()+"#co2"] = utils.GetRandInt(300, 500)
	data[getRandId()+"#pm25"] = utils.GetRandInt(0, 100)
	data[getRandId()+"#aqi"] = utils.GetRandInt(0, 100)

	data[getRandId()+"#AcVoltage"] = utils.GetRandFloat(210, 230)
	data[getRandId()+"#AcCurrent"] = utils.GetRandFloat(0, 1)

	data[getRandId()+"#DcVoltage"] = utils.GetRandFloat(3, 12)
	data[getRandId()+"#DcCurrent"] = utils.GetRandFloat(0, 1)

	data[getRandId()+"#longitude"] = utils.GetRandFloat(114, 128)
	data[getRandId()+"#latitude"] = utils.GetRandFloat(30, 32)

	data[getRandId()+"#on"] = utils.GetRandBool()
	data[getRandId()+"#state"] = utils.GetRandBool()

	for deviceIndex := 1; deviceIndex <= 5; deviceIndex++ {
		id := fmt.Sprintf("00000001-0000-0000-0000-0000000%0*d", 5, deviceIndex)

		for i := 0; i < 5; i++ {
			strI := fmt.Sprintf("%0*d", 3, i)
			data[id+"#int"+strI] = utils.GetRandInt(10*i, 10*i+10)
			data[id+"#float"+strI] = utils.GetRandFloat(10*float64(i), 10*float64(i)+10)
			data[id+"#bool"+strI] = utils.GetRandBool()
			data[id+"#text"+strI] = utils.GetRandString(5)
		}
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

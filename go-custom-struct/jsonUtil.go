package mkilic

import (
	"encoding/json"
	"log"
)

func ObjectTojson(device *Device) string {

	var jsonData []byte
	jsonData,err := json.Marshal(device)
	if err != nil {
		log.Println(err)
	}

	return string(jsonData)
}

func JsonToObject(jsonData []byte) *Device {

	var unDevice Device
	_ = json.Unmarshal(jsonData, &unDevice)

	return unDevice
}

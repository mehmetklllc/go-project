package mkilic

import (
	"encoding/json"
	"github.com/mehmetklllc/go-project/go-custom-struct"
	"log"
)

func ObjectTojson(device *mkilic.Device) string {

	var jsonData []byte
	jsonData,err := json.Marshal(device)
	if err != nil {
		log.Println(err)
	}

	return string(jsonData)
}

func JsonToObject(jsonData []byte) *mkilic.Device {

	var unDevice mkilic.Device
	_ = json.Unmarshal(jsonData, &unDevice)

	return unDevice
}

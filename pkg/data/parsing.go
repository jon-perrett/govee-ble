package data

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type GoveeAdvertisementData struct {
	Temperature float32
	Humidity    float32
	Battery     int32
}

func ParseData(data map[int][]byte) GoveeAdvertisementData {
	tempData := data[60552]

	tempHumidity := getTempHumidity(tempData)
	battery := getBatteryLevel(tempData)

	temp := getTemp(tempHumidity)
	humidity := getHumidity(tempHumidity)
	return GoveeAdvertisementData{Temperature: temp, Humidity: humidity, Battery: int32(battery)}
}

func getHumidity(tempHumidity int32) float32 {
	humidity := float32(tempHumidity%1000) / 10
	return humidity
}

func getTemp(tempHumidity int32) float32 {
	temp := float32(tempHumidity/1000) / 10
	return temp
}

func getBatteryLevel(tempData []byte) int16 {
	// looks like battery is little endian for some reason ??
	var battery int16
	err := binary.Read(bytes.NewReader(tempData[len(tempData)-2:]), binary.LittleEndian, &battery)
	if err != nil {
		fmt.Println(tempData[len(tempData)-2:])
		fmt.Println(err)
	}
	return battery
}

func getTempHumidity(tempData []byte) int32 {
	var tempHumidity int32
	err := binary.Read(bytes.NewReader(tempData[0:len(tempData)-2]), binary.BigEndian, &tempHumidity)
	if err != nil {
		fmt.Println(err)
	}
	return tempHumidity
}

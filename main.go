package main

import (
	"github.com/jonperrett/govee-ble/bluetooth"
	"github.com/jonperrett/govee-ble/data"
)

func main() {
	advertisementData := bluetooth.GetAdvertisementManufacturerData()
	goveeData := data.ParseData(advertisementData)
	println(goveeData.Battery)
	println(goveeData.Humidity)
	println(goveeData.Temperature)
}

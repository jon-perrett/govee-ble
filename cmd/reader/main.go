package main

import (
	"time"

	"github.com/jonperrett/govee-ble/pkg/reading"
	"github.com/jonperrett/govee-ble/pkg/zero"
)

func main() {
	writer := zero.ZeroWriter{Address: "localhost:5555"}
	for {
		reading := reading.GetReading()
		go writer.Write(reading)
		time.Sleep(time.Second * 30)
	}

}

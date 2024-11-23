package main

import (
	"log"

	pb "github.com/jonperrett/govee-ble/generated/github.com/jonperrett/govee-ble/proto"
	"github.com/jonperrett/govee-ble/pkg/zero"
	"google.golang.org/protobuf/proto"
)

func main() {
	reader := zero.ZeroReader{Address: "127.0.0.1:5556"}
	for {
		data, err := reader.Poll()
		if err != nil {
			log.Println("Unable to read data")
		}
		reading := pb.Reading{}
		proto.Unmarshal(data, &reading)
		print(reading.GetBattery())

	}
}

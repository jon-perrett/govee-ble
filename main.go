package main

import (
	"io/ioutil"
	"log"

	"github.com/jonperrett/govee-ble/bluetooth"
	"github.com/jonperrett/govee-ble/data"
	pb "github.com/jonperrett/govee-ble/generated/github.com/jonperrett/govee-ble/proto"
	"google.golang.org/protobuf/proto"
	"gopkg.in/zeromq/goczmq.v4"
)

func main() {
	advertisementData := bluetooth.GetAdvertisementManufacturerData()
	goveeData := data.ParseData(advertisementData)
	println(goveeData.Battery)
	println(goveeData.Humidity)
	println(goveeData.Temperature)
	reading := &pb.Reading{Temperature: goveeData.Temperature, Humidity: goveeData.Humidity, Battery: goveeData.Battery}
	out, err := proto.Marshal(reading)
	if err != nil {
		log.Fatalln("Failed to encode reading:", err)
	}
	if err := ioutil.WriteFile("./outfile.pb", out, 0644); err != nil {
		log.Fatalln("Failed to write reading:", err)
	}
	sock, err := goczmq.NewReq("tcp://127.0.0.1:5555")
	if err != nil {
		log.Fatalln("Unable to connect to broker")
	}
	sock.Write(out)
	buf := make([]byte, 32)
	sock.Read(buf)
	println(string(buf[:]))
}

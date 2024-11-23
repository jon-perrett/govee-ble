package reading

import (
	"log"
	"time"

	pb "github.com/jonperrett/govee-ble/generated/github.com/jonperrett/govee-ble/proto"
	"github.com/jonperrett/govee-ble/pkg/bluetooth"
	"github.com/jonperrett/govee-ble/pkg/data"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetReading() []byte {
	advertisementData := bluetooth.GetAdvertisementManufacturerData()
	goveeData := data.ParseData(advertisementData)
	log.Printf("Reading data: %+v\n", goveeData)
	reading := &pb.Reading{Temperature: goveeData.Temperature, Humidity: goveeData.Humidity, Battery: goveeData.Battery, Timestamp: timestamppb.New(time.Now())}
	out, err := proto.Marshal(reading)
	if err != nil {
		log.Println("Could not write reading to protobuf format")
	}
	return out
}

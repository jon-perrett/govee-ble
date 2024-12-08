package store

import (
	"context"
	"fmt"
	"log"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	pb "github.com/jonperrett/govee-ble/generated/github.com/jonperrett/govee-ble/proto"
)

type InfluxStore struct {
	Address string
	Bucket  string
	Org     string
	Token   string
}

func (s *InfluxStore) WritePoint(reading *pb.Reading) {

	// Store the URL of your InfluxDB instance
	url := fmt.Sprintf("http://%s", s.Address)
	client := influxdb2.NewClient(url, s.Token)

	writeAPI := client.WriteAPIBlocking(s.Org, s.Bucket)

	p := influxdb2.NewPoint("temp-humidity",
		map[string]string{"type": "reading"},
		map[string]interface{}{"temp": reading.GetTemperature(), "humidity": reading.GetHumidity(), "battery": reading.GetBattery()},
		reading.GetTimestamp().AsTime())
	writeAPI.WritePoint(context.Background(), p)
	client.Close()

}

func NewInfluxFromEnvironment(address string) InfluxStore {
	token, org, bucket := parseEnvironment()
	log.Printf("Token: %s, Org: %s, Bucket %s", token, org, bucket)
	return InfluxStore{Address: address, Token: token, Org: org, Bucket: bucket}
}

func parseEnvironment() (string, string, string) {
	var influxToken string
	if token, ok := os.LookupEnv("INFLUX_TOKEN"); !ok {
		log.Fatalf("No InfluxDB token set.")
	} else {
		influxToken = token
	}
	var organisation string
	if org, ok := os.LookupEnv("INFLUX_ORG"); !ok {
		log.Fatalf("No InfluxDB org set.")
	} else {
		organisation = org
	}
	var bucket string
	if b, ok := os.LookupEnv("INFLUX_ORG"); !ok {
		log.Fatalf("No InfluxDB org set.")
	} else {
		bucket = b
	}
	return influxToken, organisation, bucket
}

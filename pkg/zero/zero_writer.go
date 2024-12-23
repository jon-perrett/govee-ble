package zero

import (
	"fmt"
	"log"

	"gopkg.in/zeromq/goczmq.v4"
)

type ZeroWriter struct {
	Address string
}

func (writer *ZeroWriter) Write(data []byte) {
	sock, err := goczmq.NewReq(fmt.Sprintf("tcp://%s", writer.Address))
	if err != nil {
		log.Fatalln("Unable to connect to broker")
		return
	}
	sock.Write(data)
	buf := make([]byte, 32)
	n, err := sock.Read(buf)
	if (string(buf[:n]) != "OK") || (err != nil) {
		log.Println("Failed to write to zeromq")
		log.Println(err)
		return
	}
	log.Printf("Wrote %d bytes to zero MQ\n", len(data))
}

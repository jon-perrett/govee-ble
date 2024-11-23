package zero

import (
	"fmt"
	"log"

	"gopkg.in/zeromq/goczmq.v4"
)

type ZeroReader struct {
	Address string
}

func (writer *ZeroReader) Poll() ([]byte, error) {
	sock, err := goczmq.NewSub(fmt.Sprintf("tcp://%s", writer.Address), "")
	if err != nil {
		log.Fatalln("Unable to connect to broker")
		return []byte{}, err
	}

	buf := make([]byte, 32)
	_, err = sock.Read(buf)
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}
	return buf, nil
}

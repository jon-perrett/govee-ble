package zero

import (
	"fmt"
	"log"

	"gopkg.in/zeromq/goczmq.v4"
)

type ZeroReader struct {
	Address string
	sock    *goczmq.Sock
	rw      *goczmq.ReadWriter
}

func (z *ZeroReader) Initialise() {
	sock, err := goczmq.NewSub(fmt.Sprintf("tcp://%s", z.Address), "")

	if err != nil {
		log.Println(err)
		log.Fatalln("Unable to connect to broker")
	}
	rw, err := goczmq.NewReadWriter(sock)
	if err != nil {
		log.Println(err)
		log.Fatalln("Unable to create ReadWriter")
	}
	z.rw = rw
	z.sock = sock
	log.Println("Initialised subscriber.")
}

func (z *ZeroReader) Poll() ([]byte, error) {
	buf := make([]byte, 32)
	_, err := z.sock.Read(buf)
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}
	log.Printf("Received data %s", string(buf))
	return buf, nil
}

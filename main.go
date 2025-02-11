package main

import (
	"fmt"
	"log"

	"github.com/lastcoyotes/distributedstorage/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		// can extend this with our own decoder
		// for the sake of learning its just GobDecoeder
		Decoder: p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("YAP!")
	select {}
}

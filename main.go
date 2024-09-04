package main

import (
	"log"

	"github.com/evgeny-s/go-distributed-storage/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOps{
		ListenAddr:    ":3001",
		Decoder:       p2p.DefaultDecoder{},
		HandshakeFunc: p2p.NOPHandshakeFunc,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}

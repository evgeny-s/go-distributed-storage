package main

import (
	"log"

	"github.com/evgeny-s/go-distributed-storage/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3001")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}

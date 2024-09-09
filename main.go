package main

import (
	"log"
	"time"

	"github.com/evgeny-s/go-distributed-storage/p2p"
)

func OnPeer(peer p2p.Peer) error {
	peer.Close()
	// fmt.Println("doing some logic with the peer outside of TCPTransport")
	return nil
}

func main() {
	tcpTransportOpts := p2p.TCPTransportOps{
		ListenAddr:    ":3001",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{

		StorageRoot:       "3001_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
	}

	s := NewFileServer(fileServerOpts)

	go func() {
		time.Sleep(time.Second * 3)
		s.Stop()
	}()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

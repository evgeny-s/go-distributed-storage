package main

import (
	"log"

	"github.com/evgeny-s/go-distributed-storage/p2p"
)

func OnPeer(peer p2p.Peer) error {
	peer.Close()
	// fmt.Println("doing some logic with the peer outside of TCPTransport")
	return nil
}

func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcpTransportOpts := p2p.TCPTransportOps{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{

		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}

	s := NewFileServer(fileServerOpts)

	tcpTransport.OnPeer = s.OnPeer

	return s
}

func main() {
	s1 := makeServer(":3001", "")
	s2 := makeServer(":4000", ":3001")

	go func() { log.Fatal(s1.Start()) }()

	s2.Start()
}

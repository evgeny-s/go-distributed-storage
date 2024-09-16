package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

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

	time.Sleep(2 * time.Second)
	go s2.Start()
	time.Sleep(2 * time.Second)

	// data := bytes.NewReader([]byte("my big data file here!"))
	// s2.Store("coolpicture.jpg", data)
	// time.Sleep(time.Millisecond * 5)

	r, err := s2.Get("coolpicture.jpg")
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}

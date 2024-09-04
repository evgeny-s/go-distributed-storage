package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(TCPTransportOps{
		ListenAddr:    listenAddr,
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       DefaultDecoder{} ,
	})

	assert.Equal(t, tr.TCPTransportOps.ListenAddr, listenAddr)

	assert.Nil(t, tr.ListenAndAccept())
}

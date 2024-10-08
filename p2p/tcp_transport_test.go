package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	// listenAddr :=
	opts := TCPTransportOpts{
		ListenAddr:    ":4000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}
	tr := NewTCPTransport(opts)

	assert.Equal(t, tr.ListenAddr, ":4000")

	assert.Nil(t, tr.ListenAndAccept())
}

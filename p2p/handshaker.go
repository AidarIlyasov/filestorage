package p2p

import "errors"

var ErrInvalidHandshke = errors.New("invalid handshake")

// HandshakeFunc
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error {
	return nil
}

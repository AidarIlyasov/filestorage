package p2p

import "net"

// RPC holds any arbitarary data that is being set over the
// eacg transport between two nodes in the network.
type RPC struct {
	From    net.Addr
	Payload []byte
}

package p2p

// Peer is an interface that representation of remote node. (remote connecting)
type Peer interface {
}

// Transport is anything that handles the communication
// between the nodes in the network. This can be of the
// form (TCP, UDP, Websockets)
type Transport interface {
	ListenAndAccept() error
}

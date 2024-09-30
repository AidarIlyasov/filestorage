package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP estabilished connection
type TPCPeer struct {
	// conn is the inderliying connection of the peer
	conn net.Conn

	// if we dial and retrive a conn => outbound == true
	// if we accept and retrive a conn => outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TPCPeer {
	return &TPCPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddres string
	listenner    net.Listener
	shakeHands   HandshakeFunc
	decoder      Decoder

	mu    sync.RWMutex
	peers map[net.Addr]*Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		shakeHands:   func(Peer) error { return nil },
		listenAddres: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listenner, err = net.Listen("tcp", t.listenAddres)

	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return err
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listenner.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
			continue
		}

		go t.hanndleConn(conn)
	}
}

func (t *TCPTransport) hanndleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.shakeHands(peer); err != nil {

	}

	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Println("TCP error %s", err)
			continue
		}
	}
}

type Temp struct {
}

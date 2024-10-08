package p2p

import (
	"fmt"
	"net"
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

func (p *TPCPeer) Close() error {
	return p.conn.Close()
}

type TCPTransport struct {
	TCPTransportOpts
	listenner net.Listener
	rpcch     chan RPC
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
	OnPeer        func(Peer) error // if it retruns error we will drop the peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch:            make(chan RPC),
	}
}

// consume implements the Transport interface, wich will return read-only channel
// fro reading the incomning message received from another peer in the network
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listenner, err = net.Listen("tcp", t.ListenAddr)

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

		fmt.Printf("new incoming connection %+v\n", conn)

		go t.hanndleConn(conn)
	}
}

func (t *TCPTransport) hanndleConn(conn net.Conn) {
	var err error

	defer func() {
		fmt.Printf("Droping the connection %s\n", err)
		conn.Close()
	}()

	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake error: %s \n", err)
		return
	}

	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			return
		}
	}

	// Read loop
	rpc := RPC{}
	for {
		err := t.Decoder.Decode(conn, &rpc)
		if err == net.ErrClosed {
			return
		}

		if err != nil {
			fmt.Printf("TCP read error %s", err)
			continue
		}

		rpc.From = conn.RemoteAddr()
		t.rpcch <- rpc

		// fmt.Printf("message %+v, from %s\n", buf[:n], msg.From.String())
	}
}

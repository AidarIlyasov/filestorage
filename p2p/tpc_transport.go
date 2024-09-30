package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddres string
	listenner    net.Listener
	mu           sync.RWMutex
	peers        map[net.Addr]*Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
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
	fmt.Printf("new incoming conntection %+v\n", conn)
}

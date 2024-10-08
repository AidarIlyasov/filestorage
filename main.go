package main

import (
	"filestorage/p2p"
	"fmt"
	"log"
)

func OnPeer(p p2p.Peer) error {
	fmt.Println("doing some login with the peer outside of TCPTransport")
	return nil

	// return fmt.Errorf("failed the onpeer connection")
}

func main() {
	tr := p2p.NewTCPTransport(p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	})

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()

	select {}
}

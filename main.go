package main

import (
	"fmt"
	"log"

	"github.com/anmolrajsoni15/dfswithgo/p2p"
)

func OnPeer(p2p.Peer) error {
	fmt.Println("doing some logic with the peer outside of TCPTarnsport")
	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Println("Received message from main: ", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Failed to listen and accept: %v", err)
	}

	select {}
}

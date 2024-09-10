package main

import (
	"log"

	"github.com/anmolrajsoni15/dfswithgo/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Failed to listen and accept: %v", err)
	}

	select {}
}

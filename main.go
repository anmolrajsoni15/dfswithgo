package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/anmolrajsoni15/dfswithgo/p2p"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcptransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(tcptransportOpts)

	fileServerOpts := FileServerOpts{
		EncKey:            newEncryptionKey(),
		StorageRoot:       "network_" + strings.Split(listenAddr, ":")[1],
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}

	s := NewFileServer(fileServerOpts)

	tcpTransport.OnPeer = s.OnPeer

	return s
}

func main() {
	s1 := makeServer(":3007", "")
	s2 := makeServer(":7007", "")
	s3 := makeServer(":5007", ":3007", ":7007")

	go func() { log.Fatal(s1.Start()) }()
	time.Sleep(1000 * time.Millisecond)
	go func() { log.Fatal(s2.Start()) }()

	time.Sleep(2 * time.Second)

	go s3.Start()
	time.Sleep(2 * time.Second)

	for i := 0; i < 2; i++ {
		key := fmt.Sprintf("picture_%d.png", i)
		data := bytes.NewReader([]byte("my big data file here!"))
		s3.Store(key, data)

		time.Sleep(3000 * time.Millisecond)
		if err := s3.store.Delete(s3.ID, key); err != nil {
			log.Fatal("error0:- ", err)
		}

		r, err := s3.Get(key)
		if err != nil {
			log.Fatal("error1:- ", err)
		}

		b, err := io.ReadAll(r)
		if err != nil {
			log.Fatal("error2:- ", err)
		}

		fmt.Println(string(b))

		// if closer, ok := r.(io.Closer); ok {
		// 	closer.Close()
		// }
	}
}



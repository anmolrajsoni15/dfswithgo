package p2p

// Peer is an interface that represents a peer in the network.
type Peer interface {
}

type Transport interface {
	ListenAndAccept() error
}

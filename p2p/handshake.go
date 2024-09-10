package p2p

//handshake func is a function that is called after a connection is established
type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error { return nil }

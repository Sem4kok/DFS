package p2p

// Peer is an interface that represents remote node
type Peer interface {
}

// Transport is handler that connects two or more
// nodes in the network. Could be:
// (TCP, UDP, websocket, gRPC, etc...)
type Transport interface {
	ListenAndAccept() error
	Shutdown()
}

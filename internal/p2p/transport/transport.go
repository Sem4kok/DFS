package transport

import (
	"github.com/Sem4kok/DFS/internal/p2p/message"
)

// Peer is an interface that represents remote node
type Peer interface {
	Close() error
}

// Transport is handler that connects two or more
// nodes in the network. Could be:
// (TCP, UDP, websocket, gRPC, etc...)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan message.RPC
	Shutdown()
}

package p2p

import (
	"fmt"
	"github.com/Sem4kok/DFS/internal/logger"
	"github.com/Sem4kok/DFS/internal/p2p"
	"net"
	"sync"
)

type TCPTransport struct {
	addr string
	lg   logger.Logger

	mu       *sync.Mutex
	listener net.Listener

	peersLock *sync.RWMutex
	peers     map[net.Addr]p2p.Peer

	isRunning bool
}

// NewTCPTransport returns TCPTransport structure
func NewTCPTransport(addr string) *TCPTransport {
	return &TCPTransport{
		addr:      addr,
		mu:        &sync.Mutex{},
		peersLock: &sync.RWMutex{},
		peers:     make(map[net.Addr]p2p.Peer),
		isRunning: false,
	}
}

// ListenAndAccept starting new listening on specified addr
// specially for tcp connections
// starts accepting loop
func (t *TCPTransport) ListenAndAccept() error {
	lr, err := net.Listen("tcp", t.addr)
	if err != nil {
		return err
	}
	t.listener = lr

	go t.startAcceptLoop()

	t.isRunning = true
	return nil
}

// startAcceptLoop waiting for requests
// when received then process
func (t *TCPTransport) startAcceptLoop() {
	const op = "TCP: startAcceptLoop:"
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			t.lg.Error(op, err)
		}

		go t.handleRequest(conn)
	}
}

func (t *TCPTransport) handleRequest(conn net.Conn) {
	t.lg.Info(fmt.Sprintf("requst handled: %+v", conn))
}

// Shutdown closes TCPTransport
func (t *TCPTransport) Shutdown() {
	const op = "TCP: Shutdown:"

	t.mu.Lock()
	defer t.mu.Unlock()

	if t.isRunning == false {
		return
	}

	if err := t.listener.Close(); err != nil {
		t.lg.Error(op, err)
	}
}

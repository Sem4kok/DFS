package transport

import (
	"fmt"
	"github.com/Sem4kok/DFS/internal/logger"
	"github.com/Sem4kok/DFS/internal/p2p/decoder"
	"github.com/Sem4kok/DFS/internal/p2p/handshake"
	"github.com/Sem4kok/DFS/internal/p2p/transport"
	"io"
	"net"
	"sync"
)

// TCPPeer represents the remote node over TCP
// if dial -> outbound = true
// if accept -> outbound = false
type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

type TCPTransportOpts struct {
	Addr          string
	Lg            logger.Logger
	HandshakeFunc handshake.HandshakeFunc
	Decoder       decoder.Decoder
}

type TCPTransport struct {
	*TCPTransportOpts

	mu       *sync.Mutex
	listener net.Listener

	peersLock *sync.RWMutex
	peers     map[net.Addr]transport.Peer

	isRunning bool
}

// NewTCPTransport returns TCPTransport structure
func NewTCPTransport(opts *TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		mu:               &sync.Mutex{},
		peersLock:        &sync.RWMutex{},
		peers:            make(map[net.Addr]transport.Peer),
		isRunning:        false,
	}
}

// ListenAndAccept starting new listening on specified addr
// specially for tcp connections
// starts accepting loop
func (t *TCPTransport) ListenAndAccept() error {
	lr, err := net.Listen("tcp", t.Addr)
	if err != nil {
		return err
	}
	t.listener = lr

	go t.startAcceptLoop()

	t.Lg.Info("TCPTransport listening on:", t.Addr)
	t.isRunning = true
	return nil
}

// startAcceptLoop waiting for requests
// when received then process
func (t *TCPTransport) startAcceptLoop() {
	const op = "TCP: startAcceptLoop:"
	t.Lg.Info("TCPTransport starting accept loop...")
	for {
		if !t.isRunning {
			return
		}

		conn, err := t.listener.Accept()
		if err != nil {
			t.Lg.Error(op, err)
		}

		go t.handleRequest(conn)
	}
}

func (t *TCPTransport) handleRequest(conn net.Conn) {
	peer := TCPPeer{conn: conn, outbound: true}
	t.Lg.Info(fmt.Sprintf("requst handled: %+v", peer))

	if err := t.HandshakeFunc(conn); err != nil {
		conn.Close()
		t.Lg.Error(fmt.Sprintf("handshake error: %v", err))
		return
	}

	msg := &decoder.Message{}
	var errorCounter int
	for {
		err := t.Decoder.Decode(conn, msg)
		switch err {
		case nil:
		case io.EOF:
			fallthrough
		default:
			errorCounter++
			return
		}

		t.Lg.Info(fmt.Sprintf("message: %+v", msg))
	}

}

// Shutdown closes TCPTransport
func (t *TCPTransport) Shutdown() {
	const op = "TCP: Shutdown:"

	t.mu.Lock()
	defer t.mu.Unlock()

	if t.isRunning == false {
		return
	}

	t.Lg.Info("TCPTransport shutting down...")
	t.isRunning = false
	if err := t.listener.Close(); err != nil {
		t.Lg.Error(op, err)
	}
}
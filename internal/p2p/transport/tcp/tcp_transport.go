package transport

import (
	"fmt"
	"github.com/Sem4kok/DFS/internal/logger"
	"github.com/Sem4kok/DFS/internal/p2p/decoder"
	"github.com/Sem4kok/DFS/internal/p2p/handshake"
	"github.com/Sem4kok/DFS/internal/p2p/message"
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

func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

type TCPTransportOpts struct {
	Addr          string
	Lg            logger.Logger
	HandshakeFunc handshake.HandshakeFunc
	Decoder       decoder.Decoder
}

type TCPTransport struct {
	*TCPTransportOpts
	rpcch chan message.RPC

	mu       *sync.Mutex
	listener net.Listener

	isRunning bool
}

func (t *TCPTransport) Consume() <-chan message.RPC {
	return t.rpcch
}

// NewTCPTransport returns TCPTransport structure
func NewTCPTransport(opts *TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		mu:               &sync.Mutex{},
		rpcch:            make(chan message.RPC),
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
		peer.Close()
		t.Lg.Error(fmt.Sprintf("handshake error: %v", err))
		return
	}

	rpc := message.RPC{}
	var errorCounter int
	for {
		err := t.Decoder.Decode(conn, &rpc)
		switch err {
		case nil:
		case io.EOF:
			fallthrough
		default:
			errorCounter++
			return
		}

		t.rpcch <- rpc
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

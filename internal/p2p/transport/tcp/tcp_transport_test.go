package transport

import (
	"github.com/Sem4kok/DFS/internal/logger"
	"github.com/Sem4kok/DFS/internal/p2p/handshake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTCPTransport(t *testing.T) {
	opts := &TCPTransportOpts{
		Addr:          ":8981",
		Lg:            &logger.NOPLogger{},
		HandshakeFunc: handshake.NOPHandshakeFunc,
	}

	tr := NewTCPTransport(opts)
	assert.Equal(t, tr.Addr, opts.Addr)

	assert.Nil(t, tr.ListenAndAccept())
	assert.Equal(t, tr.isRunning, true)

}

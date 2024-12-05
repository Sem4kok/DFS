package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTCPTransport(t *testing.T) {
	addr := ":8981"

	tr := NewTCPTransport(addr, nil)
	assert.Equal(t, tr.addr, addr)

	assert.Nil(t, tr.ListenAndAccept())
	assert.Equal(t, tr.isRunning, true)

}

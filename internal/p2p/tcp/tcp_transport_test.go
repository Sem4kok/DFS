package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTCPTransport(t *testing.T) {
	addr := ":8981"
	tr := NewTCPTransport(addr)
	assert.Equal(t, tr.addr, addr)

}

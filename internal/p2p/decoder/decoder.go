package decoder

import (
	"github.com/Sem4kok/DFS/internal/p2p/message"
	"io"
)

type Message struct {
	Payload []byte
}

type Decoder interface {
	Decode(r io.Reader, rpc *message.RPC) error
}

type NopDecoder struct{}

func (d *NopDecoder) Decode(any) error {
	return nil
}

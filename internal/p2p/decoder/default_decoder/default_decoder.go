package decoder

import (
	"github.com/Sem4kok/DFS/internal/p2p/message"
	"io"
)

type DefaultDecoder struct{}

func (d *DefaultDecoder) Decode(r io.Reader, rpc *message.RPC) error {
	buf := make([]byte, 1028)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	rpc.Payload = buf[:n]
	return nil
}

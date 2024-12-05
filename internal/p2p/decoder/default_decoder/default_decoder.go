package decoder

import (
	"github.com/Sem4kok/DFS/internal/p2p/decoder"
	"io"
)

type DefaultDecoder struct{}

func (d *DefaultDecoder) Decode(r io.Reader, msg *decoder.Message) error {
	buf := make([]byte, 1028)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	msg.Payload = buf[:n]
	return nil
}

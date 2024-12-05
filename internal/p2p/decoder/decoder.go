package decoder

import "io"

type Message struct {
	Payload []byte
}

type Decoder interface {
	Decode(r io.Reader, msg *Message) error
}

type NopDecoder struct{}

func (d *NopDecoder) Decode(any) error {
	return nil
}

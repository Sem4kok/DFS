package decoder

import (
	"encoding/gob"
	"io"
)

type GOBDecoder struct{}

func (d *GOBDecoder) Decode(r io.Reader, v any) error {
	return gob.NewDecoder(r).Decode(v)
}

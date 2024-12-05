package handshake

type HandshakeFunc func(any) error

func NOPHandshakeFunc(v any) error {
	return nil
}

package lcache

type PeerPicker interface {
	PeerPicker(key string) (PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}

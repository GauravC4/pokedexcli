package pokecache

type Cache interface {
	Get(key string) ([]byte, bool)
	Add(key string, val []byte)
}

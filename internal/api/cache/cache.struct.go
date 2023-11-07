package cache

// Value интерфейс, определяющий, какие структуры можно хранить в cache
type Value interface {
	Get(string) bool
	Set(string)
	Remove(string)
}

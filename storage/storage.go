package storage

type Storage interface {
	Get(name string) (string, error)
	Set(name string, value string) error
}

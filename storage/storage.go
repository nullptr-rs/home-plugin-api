package storage

type Storage interface {
	SetDatabase(database string) error
	SetCollection(collection string) error
	CreateCollection(collection string) error

	Database
	Collection
}

type Filter string
type Pipeline string

type Cursor interface {
	Current() []byte
	All(results interface{}) error
	Close() error
	Decode(result interface{}) error
	Err() error
	ID() int64
	Next() bool
	TryNext() bool
}

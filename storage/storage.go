package storage

type Storage interface {
	SetDatabase(database string) error
	SetCollection(collection string) error

	ListDatabases() ([]string, error)
	NumberSessionsInProgress() int
	Ping() error

	Database
	Collection
}

type Filter string
type Pipeline string

type Cursor interface {
	Current() []byte
	All(results interface{}) error
	Decode(result interface{}) error

	Err() error
	ID() int64
	Next() bool
	TryNext() bool

	Close() error
}

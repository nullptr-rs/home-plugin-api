package storage

type Storage interface {
	SetDatabase(database string)
	SetCollection(collection string)

	ListDatabases(filter Filter) ([]string, error)
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

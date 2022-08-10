package storage

type Database interface {
	Aggregate(pipeline Pipeline) (*Cursor, error)

	CreateCollection(collection string) error
	CreateView(view string, viewOn string, pipeline Pipeline) error

	ListCollectionNames(filter Filter) ([]string, error)
	ListCollections(filter Filter) (*Cursor, error)

	Name() string
	Drop() error
}

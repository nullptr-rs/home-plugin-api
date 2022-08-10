package storage

type Collection interface {
	AggregateCollection(pipeline Pipeline) (*Cursor, error)

	CountDocuments(filter Filter) (int64, error)
	EstimatedDocumentCount() (int64, error)

	DeleteOne(filter Filter) (int64, error)
	DeleteMany(filter Filter) (int64, error)

	Distinct(field string, filter Filter) ([]interface{}, error)

	Find(filter Filter) (*Cursor, error)
	FindOne(filter Filter) (*Cursor, []byte, error)
	FindOneAndDelete(filter Filter) (*Cursor, []byte, error)
	FindOneAndReplace(filter Filter, replacement interface{}) (*Cursor, []byte, error)
	FindOneAndUpdate(filter Filter, update interface{}) (*Cursor, []byte, error)

	InsertMany(documents []interface{}) ([]interface{}, error)
	InsertOne(document interface{}) (interface{}, error)

	ReplaceOne(filter Filter, replacement interface{}) (int64, int64, int64, error)

	UpdateByID(id interface{}, update interface{}) (int64, int64, int64, error)
	UpdateMany(filter Filter, update interface{}) (int64, int64, int64, error)
	UpdateOne(filter Filter, update interface{}) (int64, int64, int64, error)

	CollectionName() string
	DropCollection() error
}

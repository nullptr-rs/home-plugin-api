package storage

type Collection interface {
	Aggregate(pipeline Pipeline) (*Cursor, error)
	CountDocuments(filter Filter) (int64, error)
	DeleteOne(filter Filter) error
	DeleteMany(filter Filter) error
	Distinct(field string, filter Filter) ([]interface{}, error)
	Drop() error
	EstimatedDocumentCount() (int64, error)
	Find(filter Filter) (*Cursor, error)
	FindOne(filter Filter) (*Cursor, []byte, error)
	FindOneAndDelete(filter Filter) (*Cursor, []byte, error)
	FindOneAndReplace(filter Filter, replacement interface{}) (*Cursor, []byte, error)
	FindOneAndUpdate(filter Filter, update interface{}) (*Cursor, []byte, error)
	InsertMany(documents []interface{}) ([]interface{}, error)
	InsertOne(document interface{}) ([]interface{}, error)
	Name() string
	ReplaceOne(filter Filter, replacement interface{}) (int64, int64, int64, error)
	UpdateByID(id interface{}, update interface{}) (int64, int64, int64, error)
	UpdateMany(filter Filter, update interface{}) (int64, int64, int64, error)
	UpdateOne(filter Filter, update interface{}) (int64, int64, int64, error)
}
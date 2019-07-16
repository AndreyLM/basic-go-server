package db

// DB - interface for db
type DB interface {
	Connect() error
	Find(findBy, objects interface{}, limit, offset int32) error
	Get(model interface{}) error
	Exists(model interface{}) (bool, error)
	Update(id int64, model interface{}) error
	Store(model interface{}) error
	Delete(id int64, model interface{}) error
}

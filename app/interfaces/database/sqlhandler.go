package database

type SqlHandler interface {
	AutoMigration() error
	Find(interface{}) (interface{}, error)
	FindById(interface{}, int) (interface{}, error)
	Create(interface{}) (interface{}, error)
	Save(interface{}) (interface{}, error)
	Delete(interface{}) (interface{}, error)
}
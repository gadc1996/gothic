package domain

type Repository interface {
	ListAll(data map[string]interface{}) ([]*Entity, error)
	FindById(id int64) (*Entity, error)
}
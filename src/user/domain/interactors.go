package domain

type DatabaseInteractor interface {
	Save(entity *Entity) (*Entity, error)
}
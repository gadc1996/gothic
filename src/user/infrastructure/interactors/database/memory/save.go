package memory

import d "github.com/gadc1996/reservation-manager/src/user/domain"

func (interactor Interactor) Save(entity *d.Entity) (*d.Entity, error) {
	return entity, nil
}
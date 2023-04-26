package application

import (
	d "github.com/gadc1996/reservation-manager/src/user/domain"
)

func Save(entity *d.Entity, database d.DatabaseInteractor) (*d.Entity, error) {
	if entity, error := database.Save(entity); error != nil {
		return nil, error
	} else {
		return entity, nil
	}
}

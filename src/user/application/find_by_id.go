package application

import (
	d "github.com/gadc1996/reservation-manager/src/user/domain"
)

func FindById(id int64, repository d.Repository) (*d.Entity, error) {
	if entity, error := repository.FindById(id); error != nil {
		return nil, error
	} else {
		return entity, nil
	}
}

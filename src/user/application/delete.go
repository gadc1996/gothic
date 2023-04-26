package application

import (
	d "github.com/gadc1996/reservation-manager/src/user/domain"
)

func Delete(entity *d.Entity, database d.DatabaseInteractor) error {
	// return database.Delete(entity.Id)
	return nil
}

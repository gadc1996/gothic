package mysql

import d "github.com/gadc1996/reservation-manager/src/user/domain"

func (interactor Interactor) Save(entity *d.Entity) (*d.Entity, error) {
	result, error := database.Exec("INSERT INTO users(name) VALUES(?)", entity.Name)
	if error != nil {
		return nil, error
	}

	id, error := result.LastInsertId()
	if error != nil {
		return nil, error
	}

	entity, error = interactor.FindById(id)
	if error != nil {
		return nil, error
	}

	return entity, nil
}
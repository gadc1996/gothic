package memory

import (
	d "github.com/gadc1996/reservation-manager/src/user/domain"
)

func (repository *Repository) ListAll(data map[string]interface{}) ([]*d.Entity, error) {
	var entities []*d.Entity
	count := 5
	for i := 0; i < count; i++ {
		entities = append(entities, d.Mock())
	}

	return entities, nil

}

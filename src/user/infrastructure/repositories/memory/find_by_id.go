package memory

import (
	domain "github.com/gadc1996/reservation-manager/src/user/domain"
)

func (repository *Repository) FindById(id int64) (*domain.Entity, error) {
	mock := domain.Mock()

	return &domain.Entity{
		Id:       id,
		Name:     mock.Name,
		Password: mock.Password,
		Email:    mock.Email,
	}, nil
}

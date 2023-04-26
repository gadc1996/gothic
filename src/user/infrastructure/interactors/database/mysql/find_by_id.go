package mysql

import d "github.com/gadc1996/reservation-manager/src/user/domain"

func (interactor Interactor) FindById(id int64) (*d.Entity, error) {
	mock := d.Mock()
	mock.Id = id
	return mock, nil
}

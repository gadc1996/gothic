package mysql

import (
	"testing"

	d "github.com/gadc1996/reservation-manager/src/user/domain"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	mock := d.Mock()
	entity, error := interactor.Save(mock)
	assert.NotNil(t, entity)
	assert.NoError(t, error)
}
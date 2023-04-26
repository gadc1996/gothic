package application

import (
	"testing"

	d "github.com/gadc1996/reservation-manager/src/user/domain"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	entity := d.Mock()
	
	entity, error := Save(entity, database)

	assert.NotNil(t, entity)
	assert.NoError(t, error)
}


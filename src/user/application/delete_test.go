package application

import (
	"testing"

	d "github.com/gadc1996/reservation-manager/src/user/domain"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	entity := d.Mock()
	
	error := Delete(entity, database)

	assert.NoError(t, error)
}


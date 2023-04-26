package application

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestFindById(t *testing.T) {
	id := gofakeit.Int64()
	
	entity, error := FindById(id, repository)

	assert.NotNil(t, entity)
	assert.NoError(t, error)
}

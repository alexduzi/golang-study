package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)

	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)
}

func TestCalculateTaxWithError(t *testing.T) {
	tax, err := CalculateTax(0)

	assert.NotNil(t, err)
	assert.Equal(t, "amount must be greater than 0", err.Error())
	assert.Equal(t, float64(0), tax)
}

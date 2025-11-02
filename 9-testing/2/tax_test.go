package tax2

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)
}

func TestCalculateTaxWhenAmountIsNegative(t *testing.T) {
	tax, err := CalculateTax(-1)
	assert.Error(t, err, "negative amount")
	assert.Equal(t, 0.0, tax)
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &MockTaxRepository{}
	repository.On("SaveTax", 10.0).Return(nil)
	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)
	repository.AssertExpectations(t)
}

func TestCalculateTaxAndNotSave(t *testing.T) {
	repository := &MockTaxRepository{}
	repository.On("SaveTax", 0.0).Return(errors.New("error save tax"))
	err := CalculateTaxAndSave(0.0, repository)
	assert.NotNil(t, err)
	assert.Equal(t, "error save tax", err.Error())
	assert.Error(t, err, "error save tax")
	repository.AssertExpectations(t)
}

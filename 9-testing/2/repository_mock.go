package tax2

import (
	"github.com/stretchr/testify/mock"
)

type MockTaxRepository struct {
	mock.Mock
}

func (m *MockTaxRepository) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}

package tax2

import "errors"

type Repository interface {
	SaveTax(tax float64) error
}

func CalculateTax(amount float64) (float64, error) {
	if amount < 0 {
		return 0.0, errors.New("negative amount")
	}
	if amount == 0 {
		return 0, nil
	}

	if amount >= 1000 {
		return 10, nil
	}

	return 5, nil
}

func CalculateTax2(amount float64) float64 {
	if amount <= 0 {
		return 0.0
	}

	if amount >= 1000 {
		return 10
	}

	return 5
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax2(amount)
	return repository.SaveTax(tax)
}

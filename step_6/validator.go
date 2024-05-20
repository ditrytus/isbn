package step_6

import "fmt"

func Validate(isbn ISBN) error {
	if isbn.CalculateCheckDigit() != isbn.CheckDigit() {
		return fmt.Errorf(
			"check digit mismatch, expected %d got %d",
			isbn.CalculateCheckDigit(),
			isbn.CheckDigit(),
		)
	}
	return nil
}

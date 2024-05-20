package step_6

import "fmt"

func ParseISBN(s string) (ISBN, error) {
	var digits []int
	for _, r := range s {
		if r >= '0' && r <= '9' {
			digits = append(digits, int(r-'0'))
		}
	}
	return NewISBN(digits...)
}

func NewISBN(digits ...int) (ISBN, error) {
	switch len(digits) {
	case 10:
		return NewISBN10(digits...), nil
	case 13:
		return NewISBN13(digits...), nil
	default:
		return nil, fmt.Errorf("invalid isbn length, expected 10 or 13 digits, got %d", len(digits))
	}
}

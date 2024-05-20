package step_6

import (
	"strconv"
	"strings"
)

type ISBN interface {
	CalculateCheckDigit() int
	CheckDigit() int
	String() string
}

func ParseISBN(s string) ISBN {
	var digits []int
	for _, r := range s {
		if r >= '0' && r <= '9' {
			digits = append(digits, int(r-'0'))
		}
	}
	return NewISBN(digits...)
}

func NewISBN(digits ...int) ISBN {
	switch len(digits) {
	case 10:
		return NewISBN10(digits...)
	case 13:
		return NewISBN13(digits...)
	default:
		panic("invalid isbn length")
	}
}

type isbn []int

func (n isbn) String() string {
	var b strings.Builder
	for _, digit := range n {
		b.WriteString(strconv.Itoa(digit))
	}
	return b.String()
}

type isbn10 isbn

func NewISBN10(digits ...int) isbn10 {
	return digits
}

func (n isbn10) CheckDigit() int {
	return n[len(n)-1]
}

func (n isbn10) CalculateCheckDigit() int {
	var sum int
	for i, digit := range n[:len(n)-1] {
		sum += (10 - i) * digit
	}
	return 11 - sum%11
}

func (n isbn10) String() string {
	return isbn(n).String()
}

type isbn13 isbn

func NewISBN13(digits ...int) isbn13 {
	return digits
}

func (n isbn13) CheckDigit() int {
	return n[len(n)-1]
}

func (n isbn13) CalculateCheckDigit() int {
	var sum int
	for i, digit := range n[:(len(isbn(n)) - 1)] {
		if i%2 == 0 {
			sum += digit
		} else {
			sum += 3 * digit
		}
	}
	mod := 10 - sum%10
	if mod == 10 {
		return 0
	}
	return mod
}

func (n isbn13) String() string {
	return isbn(n).String()
}

func Validate(isbn ISBN) bool {
	return isbn.CalculateCheckDigit() == isbn.CheckDigit()
}

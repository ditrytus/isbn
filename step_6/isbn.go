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
	var b strings.Builder
	b.WriteString(isbn(n[:len(n)-1]).String())
	if n.CheckDigit() == 10 {
		b.WriteString("X")
	} else {
		b.WriteString(strconv.Itoa(n.CheckDigit()))
	}
	return b.String()
}

func (n isbn10) ISBN13() isbn13 {
	return append([]int{9, 7, 8}, n...)
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

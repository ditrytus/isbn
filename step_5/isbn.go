package step_5

type ISBN []int

func (isbn ISBN) lastDigitIndex() int {
	return len(isbn) - 1
}

func (isbn ISBN) checkDigit() int {
	return isbn[isbn.lastDigitIndex()]
}

type ISBN10 ISBN

func (isbn ISBN10) CalculateCheckDigit() int {
	var sum int
	for i, digit := range isbn[:ISBN(isbn).lastDigitIndex()] {
		sum += (10 - i) * digit
	}
	return 11 - sum%11
}

func (isbn ISBN10) Validate() bool {
	return isbn.CalculateCheckDigit() == ISBN(isbn).checkDigit()
}

type ISBN13 ISBN

func (isbn ISBN13) CalculateCheckDigit() int {
	var sum int
	for i, digit := range isbn[:12] {
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

func (isbn ISBN13) Validate() bool {
	return isbn.CalculateCheckDigit() == ISBN(isbn).checkDigit()
}

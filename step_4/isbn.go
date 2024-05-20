package step_4

type ISBN []int

type ISBN10 ISBN

func (isbn ISBN10) CalculateCheckDigit() int {
	var sum int
	for i, digit := range isbn[:9] {
		sum += (10 - i) * digit
	}
	return 11 - sum%11
}

func (isbn ISBN10) Validate() bool {
	var sum int
	for i, digit := range isbn {
		sum += (10 - i) * digit
	}
	return sum%11 == 0
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
	var sum int
	for i, digit := range isbn {
		if i%2 == 0 {
			sum += digit
		} else {
			sum += 3 * digit
		}
	}
	return sum%10 == 0
}

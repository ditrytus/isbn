package step_3

type ISBNVersion int

const (
	ISBN10 ISBNVersion = 10
	ISBN13             = 13
)

type ISBN struct {
	digits  []int
	version ISBNVersion
}

func NewISBN(digits ...int) ISBN {
	switch ISBNVersion(len(digits)) {
	case ISBN10:
		return ISBN{digits, ISBN10}
	case ISBN13:
		return ISBN{digits, ISBN13}
	default:
		panic("invalid isbn length")
	}
}

func (isbn ISBN) Validate() bool {
	switch isbn.version {
	case ISBN10:
		var sum int
		for i, digit := range isbn.digits {
			sum += (10 - i) * digit
		}
		return sum%11 == 0
	case ISBN13:
		var sum int
		for i, digit := range isbn.digits {
			if i%2 == 0 {
				sum += digit
			} else {
				sum += 3 * digit
			}
		}
		return sum%10 == 0
	default:
		panic("invalid isbn version")
	}
}

func (isbn ISBN) CalculateCheckDigit() int {
	switch isbn.version {
	case ISBN10:
		var sum int
		for i, digit := range isbn.digits[:9] {
			sum += (10 - i) * digit
		}
		return 11 - sum%11
	case ISBN13:
		var sum int
		for i, digit := range isbn.digits[:12] {
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
	default:
		panic("invalid isbn version")
	}
}

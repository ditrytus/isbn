package step_1

func ValidateISBN(isbn []int) bool {
	switch len(isbn) {
	case 10:
		var sum int
		for i, digit := range isbn {
			sum += (10 - i) * digit
		}
		return sum%11 == 0
	case 13:
		var sum int
		for i, digit := range isbn {
			if i%2 == 0 {
				sum += digit
			} else {
				sum += 3 * digit
			}
		}
		return sum%10 == 0
	default:
		panic("ivalid isbn length")
	}
}

func CalculateISBNCheckDigit(isbn []int) int {
	switch len(isbn) {
	case 9, 10:
		var sum int
		for i, digit := range isbn[:9] {
			sum += (10 - i) * digit
		}
		return 11 - sum%11
	case 12, 13:
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
	default:
		panic("invalid isbn length")
	}
}

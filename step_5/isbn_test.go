package step_5

import "testing"

func Test_ISBN10_Validate(t *testing.T) {
	tests := []struct {
		name string
		isbn ISBN10
		want bool
	}{
		// ISBN 83-7054-061-9
		// Sapkowski, Andrzej. Opowieści o Wiedźminie.  T.1.  Ostatnie życzenie /  Andrzej Sapkowski. Warszawa: Super Nowa, 2004. Print."
		{"valid isbn 10", ISBN10{8, 3, 7, 0, 5, 4, 0, 6, 1, 9}, true},
		{"invalid isbn 10", ISBN10{8, 3, 7, 0, 5, 4, 0, 6, 1, 7}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.isbn.Validate(); got != tt.want {
				t.Errorf("isbn.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ISBN10_CalculateCheckDigit(t *testing.T) {
	tests := []struct {
		name string
		isbn ISBN10
		want int
	}{
		// ISBN 83-7054-061-9
		// Sapkowski, Andrzej. Opowieści o Wiedźminie.  T.1.  Ostatnie życzenie /  Andrzej Sapkowski. Warszawa: Super Nowa, 2004. Print."
		{"calculate isbn 10 check digit", ISBN10{8, 3, 7, 0, 5, 4, 0, 6, 1, 0}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.isbn.CalculateCheckDigit(); got != tt.want {
				t.Errorf("isbn.CalculateCheckDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ISBN13_Validate(t *testing.T) {
	tests := []struct {
		name string
		isbn ISBN13
		want bool
	}{
		// ISBN 978-83-7578-063-5
		// Sapkowski, Andrzej. Ostatnie życzenie / Andrzej Sapkowski. Wydanie 1 w tej edycji. Warszawa: SuperNowa, 2014. Print.
		{"valid isbn 13", ISBN13{9, 7, 8, 8, 3, 7, 5, 7, 8, 0, 6, 3, 5}, true},
		{"invalid isbn 13", ISBN13{9, 7, 8, 8, 3, 7, 5, 7, 8, 0, 6, 3, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.isbn.Validate(); got != tt.want {
				t.Errorf("isbn.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ISBN13_CalculateCheckDigit(t *testing.T) {
	tests := []struct {
		name string
		isbn ISBN13
		want int
	}{
		// ISBN 978-83-7578-063-5
		// Sapkowski, Andrzej. Ostatnie życzenie / Andrzej Sapkowski. Wydanie 1 w tej edycji. Warszawa: SuperNowa, 2014. Print.
		{"calculate isbn 13 check digit", ISBN13{9, 7, 8, 8, 3, 7, 5, 7, 8, 0, 6, 3, 0}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.isbn.CalculateCheckDigit(); got != tt.want {
				t.Errorf("isbn.CalculateCheckDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

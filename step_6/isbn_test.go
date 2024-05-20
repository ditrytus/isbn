package step_6

import "testing"

func Test_Validate(t *testing.T) {
	tests := []struct {
		name string
		isbn ISBN
		want bool
	}{
		// isbn 83-7054-061-9
		// Sapkowski, Andrzej. Opowieści o Wiedźminie.  T.1.  Ostatnie życzenie /  Andrzej Sapkowski. Warszawa: Super Nowa, 2004. Print."
		{"valid isbn 10", NewISBN10(8, 3, 7, 0, 5, 4, 0, 6, 1, 9), true},
		{"invalid isbn 10", NewISBN10(8, 3, 7, 0, 5, 4, 0, 6, 1, 7), false},

		// isbn 978-83-7578-063-5
		// Sapkowski, Andrzej. Ostatnie życzenie / Andrzej Sapkowski. Wydanie 1 w tej edycji. Warszawa: SuperNowa, 2014. Print.
		{"valid isbn 13", NewISBN13(9, 7, 8, 8, 3, 7, 5, 7, 8, 0, 6, 3, 5), true},
		{"invalid isbn 13", NewISBN13(9, 7, 8, 8, 3, 7, 5, 7, 8, 0, 6, 3, 3), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.isbn); got != nil == tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ISBN_CalculateCheckDigit(t *testing.T) {
	tests := []struct {
		name string
		isbn ISBN
		want int
	}{
		// isbn 83-7054-061-9
		// Sapkowski, Andrzej. Opowieści o Wiedźminie.  T.1.  Ostatnie życzenie /  Andrzej Sapkowski. Warszawa: Super Nowa, 2004. Print."
		{"calculate isbn 10 check digit", NewISBN10(8, 3, 7, 0, 5, 4, 0, 6, 1, 0), 9},

		// isbn 978-83-7578-063-5
		// Sapkowski, Andrzej. Ostatnie życzenie / Andrzej Sapkowski. Wydanie 1 w tej edycji. Warszawa: SuperNowa, 2014. Print.
		{"calculate isbn 13 check digit", NewISBN13(9, 7, 8, 8, 3, 7, 5, 7, 8, 0, 6, 3, 0), 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.isbn.CalculateCheckDigit(); got != tt.want {
				t.Errorf("isbn.CalculateCheckDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ISBN_CheckDigit(t *testing.T) {
	tests := []struct {
		name string
		isbn ISBN
		want int
	}{
		// isbn 83-7054-061-9
		// Sapkowski, Andrzej. Opowieści o Wiedźminie.  T.1.  Ostatnie życzenie /  Andrzej Sapkowski. Warszawa: Super Nowa, 2004. Print."
		{"isbn 10 check digit", NewISBN10(8, 3, 7, 0, 5, 4, 0, 6, 1, 9), 9},

		// isbn 978-83-7578-063-5
		// Sapkowski, Andrzej. Ostatnie życzenie / Andrzej Sapkowski. Wydanie 1 w tej edycji. Warszawa: SuperNowa, 2014. Print.
		{"isbn 13 check digit", NewISBN13(9, 7, 8, 8, 3, 7, 5, 7, 8, 0, 6, 3, 5), 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.isbn.CheckDigit(); got != tt.want {
				t.Errorf("isbn.CheckDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

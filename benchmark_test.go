package searchvietnamese

import "testing"

var (
	boolResult   bool
	stringResult string
)

func BenchmarkContain(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = Contains("Nguyễn Hoàng Nam", "nguyen")
	}

	boolResult = r
}

func BenchmarkToAlphabet(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = ToAlphabetSensitive([]rune{'N', 'g', 'u', 'y', 'ễ', 'n', ' ', 'H', 'o', 'à', 'n', 'g', ' ', 'N', 'a', 'm'})
	}

	stringResult = r
}

func BenchmarkStrictContain(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = StrictContains("Nguyễn Hoàng Nam", "nguyen")
	}

	boolResult = r
}

func BenchmarkStrictToAlphabet(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = StrictToAlphabetSensitive([]rune{'N', 'g', 'u', 'y', 'ễ', 'n', ' ', 'H', 'o', 'à', 'n', 'g', ' ', 'N', 'a', 'm'})
	}

	stringResult = r
}

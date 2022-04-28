package searchvietnamese

import "testing"

func TestToAlphabetSensitive1(t *testing.T) {
	got := ToAlphabetSensitive([]rune{'N', 'g', 'u', 'y', 'ễ', 'n', ' ', 'H', 'o', 'à', 'n', 'g', ' ', 'N', 'a', 'm'})
	if got != "Nguyen Hoang Nam" {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyen")
	}
}

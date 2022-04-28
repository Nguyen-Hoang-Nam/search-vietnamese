package searchvietnamese

import "testing"

func TestFuzzyMatch1(t *testing.T) {
	got := FuzzyMatch("NguyỄn", "NuyÊn")
	if !got {
		t.Errorf("NguyỄn does fuzzy match NuyÊn")
	}
}

func TestFuzzyMatch2(t *testing.T) {
	got := FuzzyMatch("NguyỄn", "NguE")
	if !got {
		t.Errorf("NguyỄn does fuzzy match NguE")
	}
}

func TestFuzzyMatch3(t *testing.T) {
	got := FuzzyMatch("NguyỄn", "Nguyeen")
	if got {
		t.Errorf("NguyỄn does not fuzzy match Nguyeen")
	}
}

func TestEqualSensitive1(t *testing.T) {
	got := EqualSensitive("NguyỄn", "NguyÊn")
	if !got {
		t.Errorf("NguyỄn does equal NguyÊn")
	}
}

func TestEqualSensitive2(t *testing.T) {
	got := EqualSensitive("NguyỄn", "Nguyen")
	if got {
		t.Errorf("NguyỄn does equal Nguyen")
	}
}

func TestEqualSensitive3(t *testing.T) {
	got := EqualSensitive("NguyỄn", "NguyAn")
	if got {
		t.Errorf("NguyỄn does not equal NguyAn")
	}
}

func TestFuzzyMatchSensitive1(t *testing.T) {
	got := FuzzyMatchSensitive("NguyỄn", "NuyÊn")
	if !got {
		t.Errorf("NguyỄn does fuzzy match NuyÊn")
	}
}

func TestFuzzyMatchSensitive2(t *testing.T) {
	got := FuzzyMatchSensitive("NguyỄn", "Ngue")
	if got {
		t.Errorf("NguyỄn does fuzzy match NguE")
	}
}

func TestFuzzyMatchSensitive3(t *testing.T) {
	got := FuzzyMatchSensitive("NguyỄn", "Nguyeen")
	if got {
		t.Errorf("NguyỄn does not fuzzy match Nguyeen")
	}
}

func TestToAlphabet(t *testing.T) {
	got := ToAlphabetSensitive([]rune{'N', 'g', 'u', 'y', 'ễ', 'n'})
	if got != "Nguyen" {
		t.Errorf("Nguyễn does not to Nguyen")
	}
}

func TestIsMatchVietnamese1(t *testing.T) {
	got := isMatchVietnamese('ễ', 'e')
	if !got {
		t.Errorf("ễ does not match e")
	}
}

func TestIsMatchVietnamese2(t *testing.T) {
	got := isMatchVietnamese('ễ', 'ê')
	if !got {
		t.Errorf("ễ does not match ê")
	}
}

func TestIsMatchVietnamese3(t *testing.T) {
	got := isMatchVietnamese('ễ', 'â')
	if got {
		t.Errorf("ễ match â")
	}
}

func TestIndexAt(t *testing.T) {
	got := indexAt("NguyỄn Hoàng Nam", "NguyỄn", 0)
	if got != 0 {
		t.Errorf("NguyÊn does not in position 0 of NguyỄn Hoàng Nam")
	}
}

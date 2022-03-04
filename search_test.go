package searchvietnamese

import (
	"testing"
)

var boolResult bool

func BenchmarkContain(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = Contains("Nguyễn Hoàng Nam", "nguyen")
	}

	boolResult = r
}

func TestContain1(t *testing.T) {
	got := Contains("Nguyễn Hoàng Nam", "nguyen")
	if !got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyen")
	}
}

func TestContain2(t *testing.T) {
	got := Contains("Nguyễn Hoàng Nam", "nguyên")
	if !got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestContain3(t *testing.T) {
	got := Contains("Nguyễn Hoàng Nam", "nguyn")
	if got {
		t.Errorf("Nguyễn Hoàng Nam does contain nguyn")
	}
}

func TestIndex1(t *testing.T) {
	got := Index("Nguyễn Hoàng Nam", "hoang")
	if got != 7 {
		t.Errorf("Nguyễn Hoàng Nam does not contain hoang")
	}
}

func TestIndex2(t *testing.T) {
	got := Index("Nguyễn Hoàng Nam", "nguyên")
	if got != 0 {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestContainSensitive1(t *testing.T) {
	got := ContainsSensitive("Nguyễn Hoàng Nam", "nguyen")
	if got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyen")
	}
}

func TestContainSensitive2(t *testing.T) {
	got := ContainsSensitive("Nguyễn Hoàng Nam", "nguyên")
	if got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestContainSensitive3(t *testing.T) {
	got := ContainsSensitive("Nguyễn Hoàng Nam", "Nguyen")
	if !got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyen")
	}
}

func TestContainSensitive4(t *testing.T) {
	got := ContainsSensitive("Nguyễn Hoàng Nam", "Nguyên")
	if !got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestContainSensitive5(t *testing.T) {
	got := ContainsSensitive("NguyỄn Hoàng Nam", "Nguyen")
	if got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyen")
	}
}

func TestContainSensitive6(t *testing.T) {
	got := ContainsSensitive("NguyỄn Hoàng Nam", "Nguyên")
	if got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestContainSensitive7(t *testing.T) {
	got := ContainsSensitive("NguyỄn Hoàng Nam", "NguyEn")
	if !got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyen")
	}
}

func TestContainSensitive8(t *testing.T) {
	got := ContainsSensitive("NguyỄn Hoàng Nam", "NguyÊn")
	if !got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestIndexSensitive1(t *testing.T) {
	got := IndexSensitive("Nguyễn Hoàng Nam", "hoang")
	if got != -1 {
		t.Errorf("Nguyễn Hoàng Nam does not contain hoang")
	}
}

func TestIndexSensitive2(t *testing.T) {
	got := IndexSensitive("Nguyễn Hoàng Nam", "nguyên")
	if got != -1 {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestIndexSensitive3(t *testing.T) {
	got := IndexSensitive("Nguyễn Hoàng Nam", "Hoang")
	if got != 7 {
		t.Errorf("Nguyễn Hoàng Nam does not contain hoang")
	}
}

func TestIndexSensitive4(t *testing.T) {
	got := IndexSensitive("Nguyễn Hoàng Nam", "Nguyên")
	if got != 0 {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestIndexSensitive5(t *testing.T) {
	got := IndexSensitive("NguyỄn Hoàng Nam", "Nguyên")
	if got != -1 {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestIndexSensitive6(t *testing.T) {
	got := IndexSensitive("NguyỄn Hoàng Nam", "NguyÊn")
	if got != 0 {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestEqual1(t *testing.T) {
	got := Equal("NguyỄn", "NguyÊn")
	if !got {
		t.Errorf("NguyỄn does equal NguyÊn")
	}
}

func TestEqual2(t *testing.T) {
	got := Equal("NguyỄn", "NguyEn")
	if !got {
		t.Errorf("NguyỄn does equal NguyEn")
	}
}

func TestEqual3(t *testing.T) {
	got := Equal("NguyỄn", "NguyAn")
	if got {
		t.Errorf("NguyỄn does not equal NguyAn")
	}
}

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
	got := ToAlphabet("Nguyễn")
	if got != "Nguyen" {
		t.Errorf("Nguyễn does not to Nguyen")
	}
}

func TestIsMatchVietnamese1(t *testing.T) {
	got := isMatchVietnamese("ễ", "e")
	if !got {
		t.Errorf("ễ does not match e")
	}
}

func TestIsMatchVietnamese2(t *testing.T) {
	got := isMatchVietnamese("ễ", "ê")
	if !got {
		t.Errorf("ễ does not match ê")
	}
}

func TestIsMatchVietnamese3(t *testing.T) {
	got := isMatchVietnamese("ễ", "â")
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

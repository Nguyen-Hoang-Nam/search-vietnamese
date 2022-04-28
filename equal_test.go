package searchvietnamese

import "testing"

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

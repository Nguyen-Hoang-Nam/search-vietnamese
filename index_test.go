package searchvietnamese

import "testing"

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

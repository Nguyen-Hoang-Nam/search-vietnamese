package searchvietnamese_test

import (
	"testing"

	searchvietnamese "github.com/Nguyen-Hoang-Nam/search-vietnamese"
)

var boolResult bool

func BenchmarkContain(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = searchvietnamese.Contains("Nguyễn Hoàng Nam", "nguyen")
	}

	boolResult = r
}

func TestContain1(t *testing.T) {
	got := searchvietnamese.Contains("Nguyễn Hoàng Nam", "nguyen")
	if !got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyen")
	}
}

func TestContain2(t *testing.T) {
	got := searchvietnamese.Contains("Nguyễn Hoàng Nam", "nguyên")
	if !got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestIndex1(t *testing.T) {
	got := searchvietnamese.Index("Nguyễn Hoàng Nam", "hoang")
	if got != 7 {
		t.Errorf("Nguyễn Hoàng Nam does not contain hoang")
	}
}

func TestIndex2(t *testing.T) {
	got := searchvietnamese.Index("Nguyễn Hoàng Nam", "nguyên")
	if got != 0 {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestContainSensitive1(t *testing.T) {
	got := searchvietnamese.ContainsSensitive("Nguyễn Hoàng Nam", "nguyen")
	if got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyen")
	}
}

func TestContainSensitive2(t *testing.T) {
	got := searchvietnamese.ContainsSensitive("Nguyễn Hoàng Nam", "nguyên")
	if got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestContainSensitive3(t *testing.T) {
	got := searchvietnamese.ContainsSensitive("Nguyễn Hoàng Nam", "Nguyen")
	if !got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyen")
	}
}

func TestContainSensitive4(t *testing.T) {
	got := searchvietnamese.ContainsSensitive("Nguyễn Hoàng Nam", "Nguyên")
	if !got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestContainSensitive5(t *testing.T) {
	got := searchvietnamese.ContainsSensitive("NguyỄn Hoàng Nam", "Nguyen")
	if got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyen")
	}
}

func TestContainSensitive6(t *testing.T) {
	got := searchvietnamese.ContainsSensitive("NguyỄn Hoàng Nam", "Nguyên")
	if got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestContainSensitive7(t *testing.T) {
	got := searchvietnamese.ContainsSensitive("NguyỄn Hoàng Nam", "NguyEn")
	if !got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyen")
	}
}

func TestContainSensitive8(t *testing.T) {
	got := searchvietnamese.ContainsSensitive("NguyỄn Hoàng Nam", "NguyÊn")
	if !got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestIndexSensitive1(t *testing.T) {
	got := searchvietnamese.IndexSensitive("Nguyễn Hoàng Nam", "hoang")
	if got != -1 {
		t.Errorf("Nguyễn Hoàng Nam does not contain hoang")
	}
}

func TestIndexSensitive2(t *testing.T) {
	got := searchvietnamese.IndexSensitive("Nguyễn Hoàng Nam", "nguyên")
	if got != -1 {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestIndexSensitive3(t *testing.T) {
	got := searchvietnamese.IndexSensitive("Nguyễn Hoàng Nam", "Hoang")
	if got != 7 {
		t.Errorf("Nguyễn Hoàng Nam does not contain hoang")
	}
}

func TestIndexSensitive4(t *testing.T) {
	got := searchvietnamese.IndexSensitive("Nguyễn Hoàng Nam", "Nguyên")
	if got != 0 {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestIndexSensitive5(t *testing.T) {
	got := searchvietnamese.IndexSensitive("NguyỄn Hoàng Nam", "Nguyên")
	if got != -1 {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestIndexSensitive6(t *testing.T) {
	got := searchvietnamese.IndexSensitive("NguyỄn Hoàng Nam", "NguyÊn")
	if got != 0 {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

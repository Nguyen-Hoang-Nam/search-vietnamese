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

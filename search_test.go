package searchvietnamese_test

import (
	"testing"

	searchvietnamese "github.com/Nguyen-Hoang-Nam/search-vietnamese"
)

func TestContain(t *testing.T) {
	got := searchvietnamese.Contain("Nguyễn Hoàng Nam", "nguyen")
	if !got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyen")
	}
}

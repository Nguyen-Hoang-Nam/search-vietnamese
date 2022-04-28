package searchvietnamese

import "testing"

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

func TestStrictContain1(t *testing.T) {
	got := StrictContains("Nguyễn Hoàng Nam", "nguyen")
	if !got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyen")
	}
}

func TestStrictContain2(t *testing.T) {
	got := StrictContains("Nguyễn Hoàng Nam", "nguyên")
	if !got {
		t.Errorf("Nguyễn Hoàng Nam does not contain nguyên")
	}
}

func TestStrictContain3(t *testing.T) {
	got := StrictContains("Nguyễn Hoàng Nam", "nguyn")
	if got {
		t.Errorf("Nguyễn Hoàng Nam does contain nguyn")
	}
}

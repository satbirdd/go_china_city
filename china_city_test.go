package go_china_city

import (
	"testing"
)

func test_province_code(t *testing.T) {
	p := Province("110000")
	if p != "110000" {
		t.Errorf("province extract error, should be %v, got %v", "110000", p)
	}
}

package go_china_city

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestProvinceCode(t *testing.T) {
	p := Province("110000")
	if p != "110000" {
		t.Errorf("province extract error, should be %v, got %v", "110000", p)
	}

	p1 := Province("110100")
	if p1 != "110000" {
		t.Errorf("province extract error, should be %v, got %v", "110000", p1)
	}

	p2 := Province("110101")
	if p2 != "110000" {
		t.Errorf("province extract error, should be %v, got %v", "110000", p2)
	}
}

func TestCityCode(t *testing.T) {
	p2 := City("110101")
	if p2 != "110100" {
		t.Errorf("province extract error, should be %v, got %v", "110100", p2)
	}
}

func TestGet(t *testing.T) {
	p := Get("110000", true)
	if p != "北京市" {
		t.Errorf("province extract error, should be %v, got %v", "北京市", p)
	}

	ps := Get("110000", false)
	if ps != "北京市" {
		t.Errorf("province extract error, should be %v, got %v", "北京市", ps)
	}

	p1 := Get("110100", true)
	if p1 != "北京市市辖区" {
		t.Errorf("city extract error, should be %v, got %v", "北京市市辖区", p1)
	}

	p1s := Get("110100", false)
	if p1s != "市辖区" {
		t.Errorf("city extract error, should be %v, got %v", "市辖区", p1s)
	}

	p2 := Get("110101", true)
	if p2 != "北京市市辖区东城区" {
		t.Errorf("district extract error, should be %v, got %v", "北京市市辖区东城区", p2)
	}

	p2s := Get("110101", false)
	if p2s != "东城区" {
		t.Errorf("district extract error, should be %v, got %v", "东城区", p2s)
	}
}

func TestList(t *testing.T) {
	code := "110100"
	list := List(&code, true)
	data, _ := json.Marshal(list)
	fmt.Printf("%v", string(data))
}

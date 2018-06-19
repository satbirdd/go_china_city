package go_china_city

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regxp"
	"strings"
)

const (
	CHINA   = "000000" //全国
	PATTERN = regxp.MustCompile(`([0-9]{2})([0-9]{2})([0-9]{2})`)
)

type node struct {
	Id             string `json:"id"`
	Text           string `json:"text"`
	SensitiveAreas bool   `json:"sensitive_areas"`
}

var cities map[string]node

func Province(code string) string {
	// match(code)[1].ljust(6, '0')
	return fmt.Sprintf("%v0000", code[0:2])
}

func City(code string) string {
	return fmt.Sprintf("%v00", code[0:5])
}

func District(code string) string {
	return code
}

// func match(code string) string {
// 	PATTERN.FindStringSubmatch(code)
// }

func Get(code string, prependParent bool) string {
	area := getArea(code)

	if prependParent {
		if strings.HasSuffix(code, "0000") {
			return area
		} else if strings.HasSuffix(code, "00") {
			parentCode := Province(code)
			area = getProvince(parentCode) + area
		} else {
			cityCode := City(code)
			provinceCode := Province(code)
			area = getProvince(provinceCode) + getCity(cityCode) + area
		}
	}

	return area
}

func getArea(code string) string {
	if strings.HasSuffix(code, "0000") {
		return getProvince(code)
	} else if strings.HasSuffix(code, "00") {
		return getCity(code)
	} else {
		return getDistrict(code)
	}

	return ""
}

func getProvince(code string) string {
	for _, province := range cities["province"] {
		if province.Id == code {
			return province.Text
		}
	}
}

func getCity(code string) string {
	for _, province := range cities["city"] {
		if province.Id == code {
			return province.Text
		}
	}
}

func getDistrict(code string) string {
	for _, province := range cities["district"] {
		if province.Id == code {
			return province.Text
		}
	}
}

func init() {
	data, err := ioutil.ReadFile("./db/area.json")
	if err != nil {
		log.Fatalf("china city init failed, %v", err)
	}

	err = json.Unmarshal(data, &cities)
	if err != nil {
		log.Fatalf("china city init failed, %v", err)
	}
}

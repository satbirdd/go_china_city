package go_china_city

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "regexp"
	"strings"
)

const (
	CHINA = "000000" //全国
)

// var PATTERN = regexp.MustCompile(`([0-9]{2})([0-9]{2})([0-9]{2})`)

type node struct {
	Id             string `json:"id"`
	Text           string `json:"text"`
	SensitiveAreas bool   `json:"sensitive_areas"`
	Children       []*node
}

func (n node) WithoutChildren() node {
	return node{
		Id:             n.Id,
		Text:           n.Text,
		SensitiveAreas: n.SensitiveAreas,
	}
}

var cities map[string][]node

func Province(code string) string {
	// match(code)[1].ljust(6, '0')
	return fmt.Sprintf("%v0000", code[0:2])
}

func City(code string) string {
	return fmt.Sprintf("%v00", code[0:4])
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
			provinceCode := Province(code)
			cityCode := City(code)
			area = getProvince(provinceCode) + getCity(cityCode) + area
		}
	}

	return area
}

func List(parentId *string, showAll bool) []*node {
	var (
		parentCode string
		list       []*node
	)

	if parentId == nil {
		parentCode = "000000"
	} else {
		parentCode = *parentId
	}

	if parentCode == "" {
		return list
	}

	provinceId := Province(parentCode)
	cityId := City(parentCode)
	districtId := District(parentCode)

	if parentCode == "000000" {
		for _, prov := range cities["province"] {
			if showAll {
				newP := prov
				list = append(list, &newP)
			} else {
				singleProv := prov.WithoutChildren()
				list = append(list, &singleProv)
			}
		}
	} else if parentCode == provinceId {
		for _, prov := range cities["province"] {
			if prov.Id == parentCode {
				if showAll {
					list = prov.Children
				} else {
					for _, child := range prov.Children {
						singleProv := child.WithoutChildren()
						list = append(list, &singleProv)
					}
				}
			}
		}
	} else if parentCode == cityId {
		for _, city := range cities["city"] {
			if city.Id == parentCode {
				if showAll {
					list = city.Children
				} else {
					for _, child := range city.Children {
						singleDistrict := child.WithoutChildren()
						list = append(list, &singleDistrict)
					}
				}
			}
		}
	} else if parentCode == districtId {
		for _, district := range cities["district"] {
			if district.Id == parentCode {
				list = district.Children
			}
		}
	}

	return list
}

// func Code(area string) (string, error) {
// 	for _, province := range cities["province"] {
// 		if province.Text == area {
// 			return province.Id, nil
// 		}
// 	}

// 	for _, city := range cities["city"] {
// 		if city.Text == area {
// 			return city.Id, nil
// 		}
// 	}

// 	for _, district := range cities["district"] {
// 		if district.Text == area {
// 			return district.Id, nil
// 		}
// 	}

// 	return "", fmt.Errorf("not found")
// }

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

	return ""
}

func getCity(code string) string {
	for _, province := range cities["city"] {
		if province.Id == code {
			return province.Text
		}
	}

	return ""
}

func getDistrict(code string) string {
	for _, province := range cities["district"] {
		if province.Id == code {
			return province.Text
		}
	}

	return ""
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

	for i, street := range cities["street"] {
		dCode := street.Id[0:6]
		for j, district := range cities["district"] {
			if district.Id == dCode {
				cities["district"][j].Children = append(district.Children, &cities["street"][i])
			}
		}
	}

	for i, district := range cities["district"] {
		cCode := fmt.Sprintf("%v00", district.Id[0:4])
		for j, city := range cities["city"] {
			if city.Id == cCode {
				cities["city"][j].Children = append(city.Children, &cities["district"][i])
			}
		}
	}

	for i, city := range cities["city"] {
		pCode := fmt.Sprintf("%v0000", city.Id[0:2])
		for j, province := range cities["province"] {
			if province.Id == pCode {
				cities["province"][j].Children = append(province.Children, &cities["city"][i])
			}
		}
	}
}

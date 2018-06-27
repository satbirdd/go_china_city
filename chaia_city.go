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
	IndexChildren  map[string]*node
	Children       []*node
}

func (n node) WithoutChildren() node {
	return node{
		Id:             n.Id,
		Text:           n.Text,
		SensitiveAreas: n.SensitiveAreas,
	}
}

var cities = map[string]*node{}

func Province(code string) string {
	// match(code)[1].ljust(6, '0')
	return fmt.Sprintf("%v0000", code[0:2])
}

func City(code string) string {
	return fmt.Sprintf("%v00", code[0:4])
}

func District(code string) string {
	return code[0:6]
}

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
		for _, prov := range cities {
			if showAll {
				// newP := *prov
				list = append(list, prov)
			} else {
				singleProv := prov.WithoutChildren()
				list = append(list, &singleProv)
			}
		}
	} else if parentCode == provinceId {
		for _, city := range cities[provinceId].IndexChildren {
			if showAll {
				// newP := *city
				list = append(list, city)
			} else {
				singleProv := city.WithoutChildren()
				list = append(list, &singleProv)
			}
		}
	} else if parentCode == cityId {
		for _, district := range cities[provinceId].IndexChildren[cityId].IndexChildren {
			if showAll {
				// newP := *prov
				list = append(list, district)
			} else {
				singleProv := district.WithoutChildren()
				list = append(list, &singleProv)
			}
		}
	} else if parentCode == districtId {
		for _, street := range cities[provinceId].IndexChildren[cityId].IndexChildren[districtId].IndexChildren {
			list = append(list, street)
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
	province := cities[code]
	return province.Text
}

func getCity(code string) string {
	pCode := Province(code)
	city := cities[pCode].IndexChildren[code]
	return city.Text
}

func getDistrict(code string) string {
	pCode := Province(code)
	cCode := City(code)
	district := cities[pCode].IndexChildren[cCode].IndexChildren[code]
	return district.Text
}

func init() {
	data, err := ioutil.ReadFile("./db/area.json")
	if err != nil {
		log.Fatalf("china city init failed, %v", err)
	}

	var flatCities map[string][]node
	err = json.Unmarshal(data, &flatCities)
	if err != nil {
		log.Fatalf("china city init failed, %v", err)
	}

	for i, province := range flatCities["province"] {
		cities[province.Id] = &flatCities["province"][i]
		cities[province.Id].IndexChildren = map[string]*node{}
	}

	for i, city := range flatCities["city"] {
		pCode := Province(city.Id)
		cities[pCode].IndexChildren[city.Id] = &flatCities["city"][i]
		cities[pCode].IndexChildren[city.Id].IndexChildren = map[string]*node{}
	}

	for i, district := range flatCities["district"] {
		pCode := Province(district.Id)
		cCode := City(district.Id)
		cities[pCode].IndexChildren[cCode].IndexChildren[district.Id] = &flatCities["district"][i]
		cities[pCode].IndexChildren[cCode].IndexChildren[district.Id].IndexChildren = map[string]*node{}
	}

	for i, street := range flatCities["street"] {
		pCode := Province(street.Id)
		cCode := City(street.Id)
		dCode := District(street.Id)
		cities[pCode].IndexChildren[cCode].IndexChildren[dCode].IndexChildren[street.Id] = &flatCities["street"][i]
	}
}

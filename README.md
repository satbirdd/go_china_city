The go implementation of saberma/china_city, which is wrote by ruby, https://github.com/saberma/china_city.

=================================
package main

import (
	"fmt"
	chinaCity "github.com/satbirdd/go_china_city"
)

func main() {
	fmt.Print(chinaCity.List(nil, false))

	// [{"id":"520000","text":"贵州省"},{"id":"620000","text":"甘肃省"},....]

	code := "520000"
	fmt.Print(chinaCity.List(&code, false))
	// [{"id":"520300","text":"遵义市"},{"id":"520400","text":"安顺市"},{"id":"520500","text":"毕节市"},{"id":"520600","text":"铜仁市"},{"id":"522700","text":"黔南布依族苗族自治州"},{"id":"520100","text":"贵阳市"},{"id":"520200","text":"六盘水市"},{"id":"522300","text":"黔西南布依族苗族自治州"},{"id":"522600","text":"黔东南苗族侗族自治州"}]

	fmt.Print(chinaCity.Get("520300", false))
	// "遵义市"

	fmt.Print(chinaCity.Get("520300", true))
	// "贵州省遵义市"
}

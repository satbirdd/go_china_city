The go implementation of saberma/china_city, which is wrote by ruby, https://github.com/saberma/china_city.

=================================
package main

import (
	"fmt"
	chinaCity "github.com/satbirdd/go_china_city"
)

func main() {
	fmt.Print(chinaCity.List(nil, false))
	// [{"id":"520000","text":"贵州省"},{"id":"620000","text":"甘肃省"},{"id":"710000","text":"台湾省"},{"id":"310000","text":"上海市"},{"id":"340000","text":"安徽省"},{"id":"420000","text":"湖北省"},{"id":"440000","text":"广东省"},{"id":"230000","text":"黑龙江省"},{"id":"330000","text":"浙江省"},{"id":"530000","text":"云南省"},{"id":"640000","text":"宁夏回族自治区"},{"id":"120000","text":"天津市"},{"id":"410000","text":"河南省"},{"id":"350000","text":"福建省"},{"id":"360000","text":"江西省"},{"id":"430000","text":"湖南省"},{"id":"510000","text":"四川省"},{"id":"370000","text":"山东省"},{"id":"460000","text":"海南省"},{"id":"820000","text":"澳门特别行政区"},{"id":"210000","text":"辽宁省"},{"id":"450000","text":"广西壮族自治区"},{"id":"630000","text":"青海省"},{"id":"810000","text":"香港特别行政区"},{"id":"110000","text":"北京市"},{"id":"150000","text":"内蒙古自治区"},{"id":"320000","text":"江苏省"},{"id":"540000","text":"西藏自治区"},{"id":"610000","text":"陕西省"},{"id":"650000","text":"新疆维吾尔自治区"},{"id":"130000","text":"河北省"},{"id":"140000","text":"山西省"},{"id":"220000","text":"吉林省"},{"id":"500000","text":"重庆市"}]

	code := "520000"
	fmt.Print(chinaCity.List(&code, false))
	// [{"id":"520300","text":"遵义市"},{"id":"520400","text":"安顺市"},{"id":"520500","text":"毕节市"},{"id":"520600","text":"铜仁市"},{"id":"522700","text":"黔南布依族苗族自治州"},{"id":"520100","text":"贵阳市"},{"id":"520200","text":"六盘水市"},{"id":"522300","text":"黔西南布依族苗族自治州"},{"id":"522600","text":"黔东南苗族侗族自治州"}]

	fmt.Print(chinaCity.Get("520300", false))
	// "遵义市"

	fmt.Print(chinaCity.Get("520300", true))
	// "贵州省遵义市"
}

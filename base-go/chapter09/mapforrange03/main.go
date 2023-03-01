package main

import (
	"fmt"
)

func main() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "南京"
	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}
	// map的长度 内置函数len
	fmt.Println("cities 有", len(cities), "对 key-value")

	// for range 复杂遍历map
	// 案例演示 key 是string, value 是 map
	// studentMap := make(map[string]map[string]string, 3)
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京长安街！"

	// 不能少 还是要 make
	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "上海南京路"

	// fmt.Println(studentMap)
	// fmt.Println(studentMap["stu02"])

	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v \n", k2, v2)
		}
		fmt.Println()
	}
}

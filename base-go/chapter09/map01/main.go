package main

import (
	"fmt"
)

func main() {
	// map声明
	// 必须 make 后使用
	var a = make(map[string]string, 10)

	a["no1"] = "宋江"
	a["no2"] = "吴用"
	// key 是否可以重复？？？ "宋江" 会被覆盖
	a["no1"] = "武松"
	// value 是否可以重复？？？ 可以再次打印出 "吴用"
	a["no3"] = "吴用"
	// map 的key value 是无序的
	fmt.Println(a)

	// map 的三种方式
	// 上面是第一种

	// 第二种 (建议)
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "南京"
	fmt.Println(cities)

	// 第三种方式  声明时直接赋值
	// var heroes map[string]string = map[string]string{}
	heroes := map[string]string{ // 类型推导
		"hero1": "宋江",
		"hero2": "吴用",
		"hero3": "卢俊义",
	}
	heroes["hero4"] = "林冲"
	fmt.Println(heroes)

	// 案例演示 key 是string, value 是 map
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

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu02"])
}

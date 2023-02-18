package main

import (
	"fmt"
)

// 编写一个函数
func modifyUser(users map[string]map[string]string, name string) {
	// users["ueserName1"] = make(map[string]string)
	// users["ueserName1"]["name"] = "张良"
	// users["ueserName1"]["psw"] = "112233"

	// 判断
	// v , ok := users[name] //也可以使用
	if users[name] != nil {
		// 有这个用户
		users[name]["psw"] = "888888"
	} else {
		// 没有这个用户
		users[name] = make(map[string]string)
		users[name]["psw"] = "888888"
		users[name]["nickname"] = "昵称=" + name // 示意
	}

}

func main() {

	// map声明 赋值
	users := make(map[string]map[string]string)
	users["张良"] = make(map[string]string)
	users["张良"]["psw"] = "112233"
	users["张良"]["nickname"] = "smith"

	// 调用
	modifyUser(users, "jack")
	modifyUser(users, "mary")
	modifyUser(users, "张良")

	fmt.Println(users)
	for i, v := range users {
		fmt.Printf("%v %v \n", i, v)
	}

}

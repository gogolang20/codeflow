package main

import (
	"fmt"
)

func main() {
	/*
		查找
		输入名字 查找是否在顺序
	*/
	names := [4]string{"白眉鹰王", "金毛狮王", "青翼蝠王", "紫衫龙王"}
	var heroName = ""
	fmt.Println("请输入哟查找的人名...")
	fmt.Scanln(&heroName)

	// 顺序查找：方式一
	// for i := 0; i < len(names); i++ {
	// 	if heroName  == names[i] {
	// 		//找到值 就跳出循环
	// 		fmt.Printf("找到%v, 下标%v\n", heroName, i)
	// 		break
	// 	} else if i == (len(names) - 1) {//找到数组最后一个值
	// 		fmt.Printf("没有找到%v\n", heroName)
	// 	}
	// }

	// 顺序查找：方式二 (推荐使用)
	index := -1
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			index = i // 将找到的值对应的下标赋给 index
			break
		}
	}
	if index != -1 { // 说明找到了
		fmt.Printf("找到%v, 下标%v\n", heroName, index)
	} else {
		fmt.Println("没有找到", heroName)
	}

}

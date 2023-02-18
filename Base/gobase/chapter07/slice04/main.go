package main

import (
	"fmt"
)

func main() {

	// 演示slice使用
	var intArr [5]int = [...]int{1, 23, 45, 435, 54}
	// 定义一个切片
	// slice 是切片的名
	// intArr[1:3] 表示slice 引用到intArr 这个数组
	// 引用intArr数组的起始下标为1 最后下标为3 但是不包含3
	slice := intArr[1:4] // [23, 45, 435]
	sliceTest := intArr[:]
	fmt.Println("sliceTest=", sliceTest)
	// 切片可以继续切片
	sliceTest2 := slice[:2] // [23, 45]
	fmt.Println("sliceTest2=", sliceTest2)

	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是=", slice)
	fmt.Println("slice 的元素个数=", len(slice))
	fmt.Println("slice 的容量=", cap(slice)) // 容量是目前可以存放最多个数的元素 是可以动态变化的

	fmt.Printf("intArr[1]的地址=%p\n", &intArr[1])
	fmt.Printf("slice[0]的地址=%p slice[0]=%v\n", &slice[0], slice[0])
	fmt.Printf("slice的地址=%p \n", &slice) // slice 的地址

	fmt.Println()
	slice[1] = 476
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是=", slice)

	// 切片的使用方法 上面是第一种
	// 第二种
	var slice2 []float64 = make([]float64, 5, 10)
	slice2[1] = 5
	slice2[3] = 9
	fmt.Println("slice2=", slice2)

	fmt.Println()
	// 第三种
	var slice3 []string = []string{"tom", "mary", "jack"}
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	// 切片遍历
	// 传统方式
	for i := 0; i < len(slice3); i++ {
		fmt.Printf("slice3[%v] = %v\n", i, slice3[i])
	}

	// for index, value range 遍历
	for index, value := range slice3 {
		fmt.Printf("index = %v value = %v \n", index, value)
	}

	// append 内置函数扩容
	var slice4 []int = []int{100, 200, 300}
	// 通过append 直接给slice4追加具体的元素
	slice4 = append(slice4, 400, 500, 600)
	// 通过append 直接将slice4追加给slice4
	slice4 = append(slice4, slice4...)

	fmt.Println("slice4", slice4)

	// copy 只能使用切片copy
	var slice5 []int = []int{1, 3, 4, 5, 6, 9}
	var slice6 = make([]int, 10)
	fmt.Println("slice6=", slice6)
	// 将slice5的值拷贝到slice6
	copy(slice6, slice5)
	fmt.Println("slice6=", slice6)
	fmt.Println("slice5=", slice5)

}

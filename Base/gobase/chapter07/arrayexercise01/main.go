package main

import (
	"fmt"
	"math/rand" // 生成一个随机数应用包
	"time"
)

func main() {

	// 字符可以进行运算的特点来赋值
	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i)
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", myChars[i])
	}

	fmt.Println()
	// 寻找数组最大值及下标
	var intArr [5]int = [...]int{3, 5, 90, -43, 9}
	var maxVal int = intArr[0]
	var maxValIndex int = 0

	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal=%v maxValIndex=%v\n", maxVal, maxValIndex)

	fmt.Println()
	// 求一个数组的和
	// 求平均值 （小数）
	var intArr2 [5]int = [...]int{3, 5, 90, -43, 9}
	sum := 0
	for _, val := range intArr2 {
		sum += val
	}

	// 转成小数
	fmt.Printf("sum=%v 平均值=%v\n", sum, float64(sum)/float64(len(intArr2)))

	fmt.Println()
	// 随机生成五个数 反转打印
	// rand,Intn
	// 得到随机数后 放入数组
	// 反转  交换的次数是 数组/2
	var intArr3 [5]int
	// 为了每次生成的随机数不一样 需要给一个seed值

	len := len(intArr3) // 跟简便
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr3[i] = rand.Intn(100) // [0<=n<100)
	}

	fmt.Printf("！交换前= %d \n", intArr3)
	// 交换
	// temp := 0 //临时变量 用于交换
	for i := 0; i < len/2; i++ {
		// temp = intArr3[len-1-i]
		// intArr3[len-1-i] = intArr3[i]
		// intArr3[i] = temp
		intArr3[i], intArr3[len-1-i] = intArr3[len-1-i], intArr3[i]
	}

	fmt.Printf("！交换后= %d \n", intArr3)

}

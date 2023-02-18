package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
桶排序思想下的排序：计数排序 & 基数排序
1)桶排序思想下的排序都是不基于比较的排序
2)时间复杂度为O(N)，额外空间负载度O(M)
3)应用范围有限，需要样本的数据状况满足桶的划分


基数排序
适用于 非负的十进制数
（如果有负数可以先加上一个最小负数，基数排序后再还原）


桶排序 （基础班 1h20min）
时间复杂度 O(N)

计数排序
    准备一个数组最大最小值范围的新数组，根据数组内元素出现次数依次记录

基数排序
    准备十个队列桶，先从个位数依次放入桶中
    倒出后，再按照十位数重复上面的操作
*/

// 计数排序
func CountSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	max := math.MinInt
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	bucket := make([]int, max+1) //辅助数组
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]]++
	}
	i := 0
	for j := 0; j < len(bucket); j++ {
		for ; bucket[j] > 0; bucket[j]-- {
			arr[i] = j
			i++
		}
	}
}

// 基数排序 #先根据最低位数字开始
// 适用于非负的十进制数
// 如果改写成适用于负数，可以整个数组加某个数，变成全部整数在排序
// digit 是最大值的十进制位数

func RadixSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	radixSort(arr, 0, len(arr)-1, maxbits(arr))
}

// arr[L..R]排序  ,  最大值的十进制位数digit
func radixSort(arr []int, l, r int, digit int) {
	const radix = 10
	help := make([]int, r-l+1)
	for d := 1; d <= digit; d++ { // 有多少位就进出几次
		count := make([]int, radix) // count[0..9]
		for i := l; i <= r; i++ {   //每个元素所在 d位 数出现次数
			j := getDigit(arr[i], d)
			count[j]++
		}
		for i := 1; i < radix; i++ { //所在位数字 的前缀和
			count[i] = count[i] + count[i-1]
		}
		for i := r; i >= l; i-- { //出桶：从数组的右边向左边依次弹出数
			j := getDigit(arr[i], d)
			help[count[j]-1] = arr[i]
			count[j]--
		}
		copy(arr[l:r+1], help[:]) //将桶内数拷贝给 原数组
	}
}

// 求数组的最大位数
func maxbits(arr []int) int {
	res, max := 0, 0
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	for max != 0 {
		res++
		max /= 10
	}
	return res
}

// 求一个数某一位的数字
func getDigit(num int, d int) int {
	return num / int(math.Pow(10.0, float64(d-1))) % 10
}

func main() {
	now := time.Now()
	maxSize := 10000
	rand.Seed(time.Now().UnixNano())
	arr01 := make([]int, maxSize)
	arr02 := make([]int, maxSize)
	for i := 0; i < len(arr01); i++ {
		arr01[i] = rand.Intn(100000)
	}
	copy(arr02, arr01)

	CountSort(arr01)
	//Insertion(arr02)
	RadixSort(arr02)
	for i := 0; i < len(arr01); i++ {
		if arr01[i] != arr02[i] {
			fmt.Println("Oops")
			break
		}
	}
	fmt.Println("Success")

	fmt.Println(time.Since(now))
}

func Insertion(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

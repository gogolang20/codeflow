package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
异或运算：相同为0，不同为1 （无进位相加！！！）
N ^ N == 0
N ^ 0 == N
a ^ b == b ^ a
(a ^ b) ^ c == a ^ (b ^ c)
*/

/*
如何不用额外变量交换两个数？
a = a ^ b
b = a ^ b
a = a ^ b
注意：a 和 b 内存地址不相同
*/

func Swap(a, b int) {
	a = a ^ b
	b = a ^ b // a ^ b ^ b
	a = a ^ b // a ^ b ^ a
}

/*
一个数组中有一种数出现了奇数次，其他数都出现了偶数次，怎么找到并打印这种数？？？
*/

func EvenTimesOddTimes(arr []int) int {
	eor := 0
	for _, value := range arr {
		eor ^= value
	}
	return eor
}

/*
一个数组中有两种数出现了奇数次，其他数都出现了偶数次，怎么找到并打印这两种数？？？
eor = a ^ b #eor != 0, 相当于两个出现奇数次的数做异或运算

前提技巧！
把一个int 类型的数，提取出二进制最右侧的1来？？？
res := a & (-a) #无论正负数都适用  #负数的补码是符号位不变 #取反符号位也变化
(-a) == (^a + 1)
*/

func OddTimesNum(arr []int) {
	eor := 0
	for i := 0; i < len(arr); i++ {
		eor ^= arr[i] // 最后结果相当于 eor = a ^ b
	}
	rightOne := eor & (-eor) // 找到最右侧为 1 的数，要找的两个数二进制在这一位不相同

	One := 0 // 其中一个数
	for _, value := range arr {
		// 只有 二进制在 这一位为 1的数，才进行异或运算
		if (value & rightOne) != 0 { // != 或者 == 都可以
			One ^= value // 出现偶数次的数字抵消
		}
	}
	Other := One ^ eor // 相当于 a ^ b ^ a
	fmt.Printf("One=%v Other=%v\n", One, Other)
}

/*
一个数组中有一种数出现了K次，其他数都出现了M次
M > 1 && K < M
找到，出现了K次的数
要求 额外空间复杂度O(1)，时间复杂度O(N)
负数没有影响
*/
func onlyKTimes(arr []int, k int, m int) int {
	var t [32]int // 申请的数组用于记录传入切片的每个 二进制位出现 1 的次数
	for _, value := range arr {
		for j := 0; j < 32; j++ { // 遍历每个元素的二进制位
			t[j] += (value >> j) & 1
		}
	}
	ans := 0
	for i := 0; i < 32; i++ {
		if (t[i] % m) == 0 { // 要找数二进制在该位置 不为 1
			continue
		}
		if (t[i] % m) == k { // 要找数二进制在该位置 为 1
			ans |= (1 << i)
		} else {
			return -1
		}
	}
	// 如果传入切片 中k 的值是0 ，需要再次特殊判断
	if ans == 0 {
		count := 0
		for i := 0; i < len(arr); i++ {
			if arr[i] == 0 {
				count++
			}
		}
		if count != k {
			return -1
		}
	}
	return ans
}

func main() {
	Swap(3, 5)
	arr := []int{1, 2}
	Swap(arr[0], arr[1])

	var arr2 = []int{6, 6, 10, 12, 12, 4, 4, 3, 3}
	fmt.Println(EvenTimesOddTimes(arr2))

	// a := 255
	// res1 := a ^ (-a)
	// res2 := a ^ (^a + 1)
	// fmt.Println(res1, res2)
	var arr3 = []int{6, 6, 6, 10, 12, 12, 4, 4, 3, 3, 9, 9}
	OddTimesNum(arr3)

	// 0 出现了2次，其他数都出现了3次
	var arr4 = []int{6, 6, 6, 12, 12, 12, 4, 4, 4, 1, 1}
	fmt.Println(onlyKTimes(arr4, 2, 3))
	fmt.Println(findOne(arr4, 2, 3))

	// 对数器测试
	testTime := 100000
	kind := 8
	success := true
	for i := 0; i < testTime; i++ {
		// 有一种数出现了K次，其他数都出现了M次
		rand.Seed(time.Now().UnixNano())
		k := rand.Intn(10) + 1
		m := rand.Intn(10) + 1
		if k > m {
			k, m = m, k
		}
		if k == m {
			m++
		}
		arrNew := generateRandomArray(kind, k, m)
		arrTest := make([]int, len(arrNew))
		copy(arrTest, arrNew)

		res1 := onlyKTimes(arrNew, k, m)
		// res1 := onlyKTimes2(arrNew, k, m)
		res2 := findOne(arrTest, k, m)

		if res1 != res2 {
			success = false
			break
		}
	}
	if success == false {
		fmt.Println("Oops")
	} else {
		fmt.Println("Success")
	}
}

// 对数器
func findOne(arr []int, k, m int) int {
	set := make(map[int]int)
	for _, value := range arr {
		set[value]++
	}
	for key, value := range set {
		if value == k {
			return key
		}
		if value == m {
			continue
		}
	}
	return -1
}

func generateRandomArray(kind int, k, m int) []int {
	// arr := make([]int, (m*kind + k))
	arr := make([]int, 0)
	var set = make(map[int]struct{})
	for i := 0; i < kind; i++ {
		res := rand.Intn(200)
		if _, ok := set[res]; !ok {
			for j := 0; j < m; j++ {
				arr = append(arr, res)
			}
			set[res] = struct{}{}
		} else {
			i--
		}
	}

	for {
		res := rand.Intn(200)
		if _, ok := set[res]; !ok {
			for j := 0; j < k; j++ {
				arr = append(arr, res)
			}
		}
		break
	}
	return arr
}

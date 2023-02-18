package main

import "fmt"

/*
对数器找规律

1）某个面试题，输入参数类型简单，并且只有一个实际参数

2）要求的返回值类型也简单，并且只有一个

3）用暴力方法，把输入参数对应的返回值，打印出来看看，进而优化code
*/

/*
小虎去买苹果，商店只提供两种类型的塑料袋，每种类型都有任意数量。
1）能装下6个苹果的袋子
2）能装下8个苹果的袋子
小虎可以自由使用两种袋子来装苹果，但是小虎有强迫症，他要求自己使用的袋子数量必须最少，且使用的每个袋子必须装满。
给定一个正整数N，返回至少使用多少袋子。如果N无法让使用的每个袋子必须装满，返回-1
*/

func minBags(apple int) int {
	if apple < 0 {
		return -1
	}
	bag8 := (apple >> 3)
	rest := apple - (bag8 << 3)
	for bag8 >= 0 {
		// rest 个
		if rest%6 == 0 {
			return bag8 + (rest / 6)
		} else {
			bag8--
			rest += 8
		}
	}
	return -1
}

func minBagAwesome(apple int) int {
	if (apple & 1) != 0 { // 如果是奇数，返回-1
		return -1
	}
	if apple < 18 {
		if apple == 12 || apple == 14 || apple == 16 {
			return 2
		} else if apple == 6 || apple == 8 {
			return 1
		} else if apple == 0 {
			return 0
		}
		return -1
	}
	return (apple-18)/8 + 3
}

func main() {
	num := 40
	fmt.Println(minBags(num))
	fmt.Println(minBagAwesome(num))
	for apple := 1; apple < num; apple++ {
		fmt.Println("bags: ", minBags(apple))
	}
}

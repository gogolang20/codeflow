package main

import (
	"fmt"
)

func main() {
	var key byte
	fmt.Println("请输入一个字符 a, b, c, d, e, f, g")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("今天星期一")
	case 'b':
		fmt.Println("今天星期二")
	case 'c':
		fmt.Println("今天星期三")
	case 'd':
		fmt.Println("今天星期四")
	case 'e':
		fmt.Println("今天星期五")
	case 'f':
		fmt.Println("今天星期六")
	case 'g':
		fmt.Println("今天星期日")
	default:
		fmt.Println("今天")
	}

	// 课堂练习1
	// var char byte
	// fmt.Println("请输入一个字符：")
	// fmt.Scanf("%c", &char)

	// switch char {
	// 	case 'a' :
	// 		fmt.Println("A")
	// 	case 'b' :
	// 		fmt.Println("B")
	// 	case 'c' :
	// 		fmt.Println("C")
	// 	case 'd' :
	// 		fmt.Println("D")
	// 	case 'e' :
	// 		fmt.Println("E")
	// 	case 'f' :
	// 		fmt.Println("F")
	// 	default :
	// 		fmt.Println("other")
	// }

	// 课堂练习2
	// var score float64
	// fmt.Println("请输入成绩")
	// fmt.Scanln(&score)

	// switch int(score / 60) {
	// 	case 1 :
	// 		fmt.Println("ok")
	// 	case 0 :
	// 		fmt.Println("wrong")
	// 	default :
	// 		fmt.Println("bad")
	// }

	// 课堂练习3

	// 课堂练习4
}

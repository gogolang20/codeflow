package main

/*
资料地址
https://blog.csdn.net/hjxisking/article/details/107469429
*/

type sumable interface {
	sum(int, int) int
}

type myFunc func(int) int

func (f myFunc) sum(a, b int) int {
	res := a + b
	return f(res)
}

func sum10(num int) int {
	return num * 10
}

func sum100(num int) int {
	return num * 100
}

type icansum struct {
	name string
	res  int
}

func (ics *icansum) sum(a, b int) int {
	ics.res = a + b
	return ics.res
}

func handlerSum(handler sumable, a, b int) int {
	res := handler.sum(a, b)
	println("handlerSum res:", res)
	return res
}

func main() {
	newFunc1 := myFunc(sum10)
	newFunc2 := myFunc(sum100)

	handlerSum(newFunc1, 1, 1)
	handlerSum(newFunc2, 1, 1)

	ics := &icansum{"I can sum", 0}
	handlerSum(ics, 1, 1)
}

package main

func aaa() (done func(), err error) {
	return func() {
		print("aaa: done")
	}, nil
}

// 本质上在函数 bbb 执行完毕后， 变量 done 已经变成了一个递归函数。
// 递归的过程是：函数 bbb 调用变量 done 后，会输出 bbb: surprise! 字符串，
// 然后又调用变量 done。而变量 done 又是这个闭包（匿名函数），从而实现不断递归调用和输出。
// func bbb() (done func(), _ error) {
//	done, err := aaa()
//	return func() {
//		print("bbb: surprise!")
//		done()
//	}, err
// }

func bbb() (func(), error) {
	done, err := aaa()
	return func() {
		print("bbb: surprise!")
		done()
	}, err
}

func main() {
	done, _ := bbb()
	done()
}

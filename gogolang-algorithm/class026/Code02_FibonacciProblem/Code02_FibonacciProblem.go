package main

import "fmt"

/*
求斐波那契数列矩阵乘法的方法

1）斐波那契数列的线性求解（O(N)）的方式非常好理解

2）同时利用线性代数，也可以改写出另一种表示

 | F(N) , F(N-1) | = | F(2), F(1) |  *  某个二阶矩阵的N-2次方  a^(N-2)

3）求出这个二阶矩阵，进而最快求出这个二阶矩阵的N-2次方
*/

/*
类似斐波那契数列的递归优化

如果某个递归，除了初始项之外，具有如下的形式

F(N) = C1 * F(N) + C2 * F(N-1) + … + Ck * F(N-k) ( C1…Ck 和k都是常数)

并且这个递归的表达式是严格的、不随条件转移的

那么都存在类似斐波那契数列的优化，时间复杂度都能优化成O(logN)
*/

/*
斐波那契数列矩阵乘法方式的实现

f3,f2 = f2,f1 * |a,b   c,d|  a,b 在上  c,d 在下
f2 * a + f1 * c = f3
f1 * b + f1 * d = f2
*/

func f1(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return f1(n-1) + f1(n-2)
}

func f2(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	res := 1
	pre := 1
	tmp := 0
	for i := 3; i <= n; i++ {
		tmp = res
		res = res + pre
		pre = tmp
	}
	return res
}

// O(logN)
func f3(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	// [ 1 ,1 ]
	// [ 1, 0 ]
	base := [][]int{
		{1, 1},
		{1, 0},
	}
	res := matrixPower(base, n-2) // 矩阵^(N-2)
	return res[0][0] + res[1][0]
}

func matrixPower(m [][]int, p int) [][]int {
	res := make([][]int, len(m))
	for i := range res {
		res[i] = make([]int, len(m[0]))
	}
	for i := 0; i < len(res); i++ {
		res[i][i] = 1
	}
	// res = 矩阵中的1
	t := m // 矩阵1次方
	for ; p != 0; p >>= 1 {
		if (p & 1) != 0 {
			res = muliMatrix(res, t)
		}
		t = muliMatrix(t, t)
	}
	return res
}

// 两个矩阵乘完之后的结果返回
func muliMatrix(m1, m2 [][]int) [][]int {
	res := make([][]int, len(m1))
	for i := range res {
		res[i] = make([]int, len(m2[0]))
	}
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m2); k++ {
				res[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return res
}

/*
第一年农场有1只成熟的母牛A，往后的每年：

1）每一只成熟的母牛都会生一只母牛

2）每一只新出生的母牛都在出生的第三年成熟

3）每一只母牛永远不会死

返回N年后牛的数量
f(N) = f(N-1) + f(N-3)
*/
func c1(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 || n == 3 {
		return n
	}
	return c1(n-1) + c1(n-3)
}

func c3(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 || n == 3 {
		return n
	}
	base := [][]int{
		{1, 1, 0},
		{0, 0, 1},
		{1, 0, 0}}
	res := matrixPower(base, n-3)
	return 3*res[0][0] + 2*res[1][0] + res[2][0]
}
func main() {
	test := 6
	fmt.Println(c1(test))
	fmt.Println(c3(test))
}

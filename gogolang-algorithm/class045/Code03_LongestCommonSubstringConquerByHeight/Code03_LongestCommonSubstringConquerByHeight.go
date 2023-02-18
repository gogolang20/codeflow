package main

// 最长公共子串问题是面试常见题目之一
// 假设str1长度N，str2长度M
// 因为最优解的难度所限，一般在面试场上回答出O(N*M)的解法已经是比较优秀了
// 因为得到O(N*M)的解法，就已经需要用到动态规划了
// 但其实这个问题的最优解是O(N+M)，为了达到这个复杂度可是不容易
// 首先需要用到DC3算法得到后缀数组(sa)
// 进而用sa数组去生成height数组
// 而且在生成的时候，还有一个不回退的优化，都非常不容易理解
// 这就是后缀数组在面试算法中的地位 : 德高望重的噩梦
func lcs1(s1, s2 string) int {
	if s1 == "" || s2 == "" {
		return 0
	}
	str1 := []byte(s1)
	str2 := []byte(s2)
	row := 0
	col := len(str2) - 1
	max := 0
	for row < len(str1) {
		i := row
		j := col
		length := 0
		for i < len(str1) && j < len(str2) {
			if str1[i] != str2[j] {
				length = 0
			} else {
				length++
			}
			if length > max {
				max = length
			}
			i++
			j++
		}
		if col > 0 {
			col--
		} else {
			row++
		}
	}
	return max
}

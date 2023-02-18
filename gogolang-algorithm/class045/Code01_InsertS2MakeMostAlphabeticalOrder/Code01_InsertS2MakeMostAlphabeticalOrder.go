package main

/*
给定两个字符串str1和str2，
想把str2整体插入到str1中的某个位置
形成最大的字典序
返回字典序最大的结果
*/

// 暴力方法
func right(s1, s2 string) string {
	if s1 == "" {
		return s2
	}
	if s2 == "" {
		return s1
	}
	p1 := s1 + s2
	p2 := s2 + s1
	ans := p2
	if p1 > p2 {
		ans = p1
	}
	for end := 1; end < len(s1); end++ {
		cur := s1[0:end] + s2 + s1[end:]
		if cur > ans {
			ans = cur
		}
	}
	return ans
}

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/last-substring-in-lexicographical-order/

// 等待测试
func lastSubstring(s string) string {
	if s == "" {
		return s
	}
	N := len(s)
	str := []byte(s)
	min := 0
	max := 0
	for _, cha := range str {
		min = Min(min, int(cha))
		max = Max(max, int(cha))
	}
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = int(str[i]) - min + 1
	}
	dc3 := sa(arr, max-min+1)
	res := make([]byte, 0)
	for _, val := range dc3[N-1:] {
		res = append(res, byte(val))
	}
	return string(res)
}

// sa 数组
func sa(nums []int, max int) []int {
	n := len(nums)
	arr := make([]int, n+3)
	for i := 0; i < n; i++ {
		arr[i] = nums[i]
	}
	return skew(arr, n, max)
}

func skew(nums []int, n, K int) []int {
	n0 := (n + 2) / 3
	n1 := (n + 1) / 3
	n2 := n / 3
	n02 := n0 + n2
	s12 := make([]int, n02+3)
	sa12 := make([]int, n02+3)
	for i, j := 0, 0; i < n+(n0-n1); i++ {
		if 0 != i%3 {
			s12[j] = i
			j++
		}
	}
	radixPass(nums, s12, sa12, 2, n02, K)
	radixPass(nums, sa12, s12, 1, n02, K)
	radixPass(nums, s12, sa12, 0, n02, K)
	name := 0
	c0 := -1
	c1 := -1
	c2 := -1
	for i := 0; i < n02; i++ {
		if c0 != nums[sa12[i]] || c1 != nums[sa12[i]+1] || c2 != nums[sa12[i]+2] {
			name++
			c0 = nums[sa12[i]]
			c1 = nums[sa12[i]+1]
			c2 = nums[sa12[i]+2]
		}
		if 1 == sa12[i]%3 {
			s12[sa12[i]/3] = name
		} else {
			s12[sa12[i]/3+n0] = name
		}
	}
	if name < n02 {
		sa12 = skew(s12, n02, name)
		for i := 0; i < n02; i++ {
			s12[sa12[i]] = i + 1
		}
	} else {
		for i := 0; i < n02; i++ {
			sa12[s12[i]-1] = i
		}
	}
	s0 := make([]int, n0)
	sa0 := make([]int, n0)
	for i, j := 0, 0; i < n02; i++ {
		if sa12[i] < n0 {
			s0[j] = 3 * sa12[i]
			j++
		}
	}
	radixPass(nums, s0, sa0, 0, n0, K)
	sa := make([]int, n)
	for p, t, k := 0, n0-n1, 0; k < n; k++ {
		i := 0
		if sa12[t] < n0 {
			i = sa12[t]*3 + 1
		} else {
			i = (sa12[t]-n0)*3 + 2
		}
		j := sa0[p]

		var tempInsert bool
		if sa12[t] < n0 {
			tempInsert = leq(nums[i], s12[sa12[t]+n0], nums[j], s12[j/3])
		} else {
			tempInsert = leq2(nums[i], nums[i+1], s12[sa12[t]-n0+1], nums[j], nums[j+1], s12[j/3+n0])
		}
		if tempInsert {
			sa[k] = i
			t++
			if t == n02 {
				for k++; p < n0; p++ {
					sa[k] = sa0[p]
					k++
				}
			}
		} else {
			sa[k] = j
			p++
			if p == n0 {
				for k++; t < n02; t++ {
					if sa12[t] < n0 {
						sa[k] = sa12[t]*3 + 1
					} else {
						sa[k] = (sa12[t]-n0)*3 + 2
					}
					k++
				}
			}
		}
	}
	return sa
}

func radixPass(nums, input, output []int, offset, n, k int) {
	cnt := make([]int, k+1)
	for i := 0; i < n; i++ {
		cnt[nums[input[i]+offset]]++
	}
	for i, sum := 0, 0; i < len(cnt); i++ {
		t := cnt[i]
		cnt[i] = sum
		sum += t
	}
	for i := 0; i < n; i++ {
		output[cnt[nums[input[i]+offset]]] = input[i]
		cnt[nums[input[i]+offset]]++
	}
}

func main() {
	s := "abab"
	fmt.Println(lastSubstring(s))
}

//func rank(sa []int) []int {
//	n := len(sa)
//	ans := make([]int, n)
//	for i := 0; i < n; i++ {
//		ans[sa[i]] = i
//	}
//	return ans
//}
//
//func height(s []int, sa, rank []int) []int {
//	n := len(s)
//	ans := make([]int, n)
//	for i, k := 0, 0; i < n; i++ {
//		if rank[i] != 0 {
//			if k > 0 {
//				k--
//			}
//			j := sa[rank[i]-1]
//			for i+k < n && j+k < n && s[i+k] == s[j+k] {
//				k++
//			}
//			ans[rank[i]] = k
//		}
//	}
//	return ans
//}

// lexicographic order
func leq(a1, a2, b1, b2 int) bool {
	return a1 < b1 || (a1 == b1 && a2 <= b2) // for pairs
}

func leq2(a1, a2, a3, b1, b2, b3 int) bool {
	return a1 < b1 || (a1 == b1 && leq(a2, a3, b2, b3)) // and triples
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a > b {
		return a
	}
	return b
}

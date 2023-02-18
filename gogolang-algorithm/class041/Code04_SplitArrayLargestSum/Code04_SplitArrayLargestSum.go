package main

import "math"

// leetcode原题
// 测试链接：https://leetcode.com/problems/split-array-largest-sum/

// 测试通过
// 课上现场写的方法，用了枚举优化，O(N * K)
func splitArray2(nums []int, K int) int {
	N := len(nums)
	sum := make([]int, N+1)
	for i := 0; i < N; i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, K+1)
	}
	best := make([][]int, N)
	for i := range dp {
		best[i] = make([]int, K+1)
	}
	for j := 1; j <= K; j++ {
		dp[0][j] = nums[0]
		best[0][j] = -1
	}
	for i := 1; i < N; i++ {
		dp[i][1] = sumFunc(sum, 0, i)
		best[i][1] = -1
	}
	// 从第2列开始，从左往右
	// 每一列，从下往上
	// 为什么这样的顺序？因为要去凑（左，下）优化位置对儿！
	for j := 2; j <= K; j++ {
		for i := N - 1; i >= 1; i-- {
			down := best[i][j-1]
			// 如果i==N-1，则不优化上限
			up := 0 // index out of range
			if i == N-1 {
				up = N - 1
			} else {
				up = best[i+1][j]
			}
			ans := math.MaxInt
			bestChoose := -1
			for leftEnd := down; leftEnd <= up; leftEnd++ {
				leftCost := 0
				if leftEnd == -1 {
					leftCost = 0
				} else {
					leftCost = dp[leftEnd][j-1]
				}
				rightCost := 0
				if leftEnd == i {
					rightCost = 0
				} else {
					rightCost = sumFunc(sum, leftEnd+1, i)
				}
				cur := Max(leftCost, rightCost)
				// 注意下面的if一定是 < 课上的错误就是此处！当时写的 <= ！
				// 也就是说，只有取得明显的好处才移动！
				// 举个例子来说明，比如[2,6,4,4]，3个画匠时候，如下两种方案都是最优:
				// (2,6) (4) 两个画匠负责 | (4) 最后一个画匠负责
				// (2,6) (4,4)两个画匠负责 | 最后一个画匠什么也不负责
				// 第一种方案划分为，[0~2] [3~3]
				// 第二种方案划分为，[0~3] [无]
				// 两种方案的答案都是8，但是划分点位置一定不要移动!
				// 只有明显取得好处时(<)，划分点位置才移动!
				// 也就是说后面的方案如果==前面的最优，不要移动！只有优于前面的最优，才移动
				// 比如上面的两个方案，如果你移动到了方案二，你会得到:
				// [2,6,4,4] 三个画匠时，最优为[0~3](前两个画家) [无](最后一个画家)，
				// 最优划分点为3位置(best[3][3])
				// 那么当4个画匠时，也就是求解dp[3][4]时
				// 因为best[3][3] = 3，这个值提供了dp[3][4]的下限
				// 而事实上dp[3][4]的最优划分为:
				// [0~2]（三个画家处理） [3~3] (一个画家处理)，此时最优解为6
				// 所以，你就得不到dp[3][4]的最优解了，因为划分点已经越过2了
				// 提供了对数器验证，你可以改成<=，对数器和leetcode都过不了
				// 这里是<，对数器和leetcode都能通过
				// 这里面会让同学们感到困惑的点：
				// 为啥==的时候，不移动，只有<的时候，才移动呢？例子懂了，但是道理何在？
				// 哈哈哈哈哈，看了邮局选址问题，你更懵，请看42节！
				if cur < ans {
					ans = cur
					bestChoose = leftEnd
				}
			}
			dp[i][j] = ans
			best[i][j] = bestChoose
		}
	}
	return dp[N-1][K]
}

// 测试最优  不是四边形不等式技巧
func splitArray3(nums []int, M int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	l := 0
	r := sum
	ans := 0
	for l <= r {
		mid := (l + r) / 2
		cur := getNeedParts(nums, mid)
		if cur <= M {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func getNeedParts(arr []int, aim int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] > aim {
			return math.MaxInt
		}
	}
	parts := 1
	all := arr[0]
	for i := 1; i < len(arr); i++ {
		if all+arr[i] > aim {
			parts++
			all = arr[i]
		} else {
			all += arr[i]
		}
	}
	return parts
}

// 求原数组arr[L...R]的累加和
func sumFunc(sum []int, L, R int) int {
	return sum[R+1] - sum[L]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

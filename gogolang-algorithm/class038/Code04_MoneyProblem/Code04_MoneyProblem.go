package main

/*
int[] d，d[i]：i号怪兽的能力
int[] p，p[i]：i号怪兽要求的钱
开始时你的能力是0，你的目标是从0号怪兽开始，通过所有的怪兽。
如果你当前的能力，小于i号怪兽的能力，你必须付出p[i]的钱，贿赂这个怪兽，然后怪兽就会加入你，他的能力直接累加到你的能力上；
	如果你当前的能力，大于等于i号怪兽的能力，你可以选择直接通过，你的能力并不会下降，你也可以选择贿赂这个怪兽，然后怪兽就会加入你，他的能力直接累加到你的能力上。
返回通过所有的怪兽，需要花的最小钱数。
*/

// int[] d d[i]：i号怪兽的武力
// int[] p p[i]：i号怪兽要求的钱
// ability 当前你所具有的能力
// index 来到了第index个怪兽的面前

// 目前，你的能力是ability，你来到了index号怪兽的面前，如果要通过后续所有的怪兽，
// 请返回需要花的最少钱数
func process1(d, p []int, ability, index int) int {
	if index == len(d) {
		return 0
	}
	if ability < d[index] {
		return p[index] + process1(d, p, ability+d[index], index+1)
	} else { // ability >= d[index] 可以贿赂，也可以不贿赂
		return Min(p[index]+process1(d, p, ability+d[index], index+1), process1(d, p, ability, index+1))
	}
}

func func1(d, p []int) int {
	return process1(d, p, 0, 0)
}


func minMoney2(d, p []int) int {
	allMoney := 0
	for i := 0; i < len(p); i++ {
		allMoney += p[i]
	}
	N := len(d)
	for money := 0; money < allMoney; money++ {
		if process2(d, p, N-1, money) != -1 {
			return money
		}
	}
	return allMoney
}

// 从0....index号怪兽，花的钱，必须严格==money
// 如果通过不了，返回-1
// 如果可以通过，返回能通过情况下的最大能力值
func process2(d, p []int, index, money int) int {
	if index == -1 { // 一个怪兽也没遇到呢
		if money == 0 {
			return 0
		} else {
			return -1
		}
	}
	// index >= 0
	// 1) 不贿赂当前index号怪兽
	preMaxAbility := process2(d, p, index-1, money)
	p1 := -1
	if preMaxAbility != -1 && preMaxAbility >= d[index] {
		p1 = preMaxAbility
	}
	// 2) 贿赂当前的怪兽 当前的钱 p[index]
	preMaxAbility2 := process2(d, p, index-1, money-p[index])
	p2 := -1
	if preMaxAbility2 != -1 {
		p2 = d[index] + preMaxAbility2
	}
	return Max(p1, p2)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

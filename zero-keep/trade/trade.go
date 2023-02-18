package main

import (
	"fmt"
)

type kind struct {
	name        string  // 品种名
	price       float64 // 价格
	unit        uint    // 一手单位
	MarginRatio float64 // 保证金比例
}

func (kd *kind) Cal() float64 {
	res := kd.price * float64(kd.unit) * kd.MarginRatio
	fmt.Printf("%s: %.0f\n", kd.name, res)

	return res
}

// prise: 每手占用资金
func Hands(prise float64) int {
	var total float64 = 30000 // 总资金
	var risk float64 = 0.2    // 承担风险比例
	var hand int              // 应持仓手数

	hand = int(total * risk / prise)
	fmt.Printf("Should buy %v hands\n", hand)

	return hand
}

func main() {
	var kinds []*kind
	kinds = append(kinds, &kind{
		name:        "C", // 玉米
		price:       2674,
		unit:        10,
		MarginRatio: 0.16,
	})
	kinds = append(kinds, &kind{
		name:        "SR", // 糖
		price:       5822,
		unit:        10,
		MarginRatio: 0.13,
	})
	kinds = append(kinds, &kind{
		name:        "MA", // 甲醇
		price:       2338,
		unit:        10,
		MarginRatio: 0.16,
	})
	kinds = append(kinds, &kind{
		name:        "FU", // 燃油
		price:       2951,
		unit:        10,
		MarginRatio: 0.24,
	})
	kinds = append(kinds, &kind{
		name:        "RB", // 螺纹
		price:       3728,
		unit:        10,
		MarginRatio: 0.2,
	})
	kinds = append(kinds, &kind{
		name:        "CF", // 棉花
		price:       14690,
		unit:        5,
		MarginRatio: 0.12,
	})

	forHands := make([]float64, 0)
	for i := range kinds {
		forHands = append(forHands, kinds[i].Cal())
	}

	for i := range forHands {
		Hands(forHands[i])
	}
}

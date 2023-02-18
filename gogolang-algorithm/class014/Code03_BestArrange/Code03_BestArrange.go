package main

import (
	"fmt"
	"sort"
)

/*
一些项目要占用一个会议室宣讲，会议室不能同时容纳两个项目的宣讲。
给你每一个项目开始的时间和结束的时间
你来安排宣讲的日程，要求会议室进行的宣讲的场次最多。
返回最多的宣讲场次。
*/

type Program struct {
	start int
	end   int
}
type Programs []*Program

func (pg Programs) Len() int {
	return len(pg)
}
func (pg Programs) Less(i, j int) bool {
	return pg[i].end < pg[j].end
}
func (pg Programs) Swap(i, j int) {
	pg[i], pg[j] = pg[j], pg[i]
}

func bestArrange2(pg Programs) int {
	sort.Sort(pg)
	timeLine, result := 0, 0
	for index := range pg {
		if timeLine <= pg[index].start {
			result++
			timeLine = pg[index].end
		}
	}
	return result
}

func main() {
	var pg Programs = []*Program{
		{
			start: 6,
			end:   7,
		}, {
			start: 1,
			end:   3,
		},
		{
			start: 3,
			end:   7,
		},
		{
			start: 6,
			end:   9,
		},
	}
	for i := range pg {
		fmt.Println(*pg[i])
	}
	fmt.Println(bestArrange2(pg))
}

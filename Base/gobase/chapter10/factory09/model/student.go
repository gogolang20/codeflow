package model

// Student 的S必须大写 才可以在其他包被应用
type student struct {
	Name  string
	score float64
}

// 工厂模式
func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		score: s,
	}
}

// 结构体内部的字段是小写 解决方式
func (s *student) GetScore() float64 {
	return s.score
}

package Abstract_Factory

import "fmt"

//抽象工厂模式
type Lunch interface {
	Cook()
}

type rise struct{}

func (c *rise) Cook() {
	fmt.Println("it is a rise.")
}

type Tomato struct{}

func (c *Tomato) Cook() {
	fmt.Println("it is Tomato.")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type simpleLunchFactory struct{}

func NewSimpleShapeFactory() LunchFactory {
	return &simpleLunchFactory{}
}

func (s *simpleLunchFactory) CreateFood() Lunch {
	return &rise{}
}

func (s *simpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}

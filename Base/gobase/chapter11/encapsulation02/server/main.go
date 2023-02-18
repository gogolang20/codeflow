package main

import (
	"fmt"

	"gobase/chapter11/encapsulation02/model"
)

func main() {
	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(8000)

	fmt.Println(p)
	fmt.Println(p.Name, "age=", p.GetAge(), "sal=", p.GetSal())
}

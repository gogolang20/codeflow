package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery1(q interface{}) {
	tpe := reflect.TypeOf(q)
	if tpe.Kind() != reflect.Struct {
		fmt.Println("unsupported type")
		return
	}

	val := reflect.ValueOf(q)
	num := val.NumField()
	sql := fmt.Sprintf("insert into %s values(", tpe.Name())

	for i := 0; i < num; i++ {
		// fmt.Println(val.Field(i))
		if val.Field(i).Kind() == reflect.Int {
			sql += fmt.Sprintf("%d", val.Field(i))
		} else {
			sql += fmt.Sprintf("\"%s\"", val.Field(i))
		}
		if i != num-1 {
			sql += ","
		}
	}
	sql += ")"
	fmt.Println(sql)
}

func createQuery(q interface{}) {
	tpe := reflect.TypeOf(q)

	switch tpe.Kind() {
	case reflect.Struct:
		val := reflect.ValueOf(q)
		sql := fmt.Sprintf("insert into %s values( ", tpe.Name())
		num := val.NumField()

		for i := 0; i < num; i++ {
			if val.Field(i).Kind() == reflect.Int {
				sql += fmt.Sprintf("%v", val.Field(i))
			} else {
				sql += fmt.Sprintf("\"%v\"", val.Field(i))
			}
			if i != num-1 {
				sql += " , "
			}
		}
		sql += " )"
		fmt.Println(sql)
	default:
		fmt.Println("unsupported type")
	}
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)

	createQuery1(o)
	createQuery1(e)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

type UserList struct {
	Users []User `json:"users"`
}

func (ulist *UserList) FormatPrint() {
	for _, u := range ulist.Users {
		fmt.Println("User Type:", u.Type)
		fmt.Println("User Age:", u.Age)
		fmt.Println("User Name:", u.Name)
		fmt.Println("Facebook Url:", u.Social.Facebook)
	}
}

func main() {
	filePath := "homework001/users.json"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("users.json error:", err)
	}

	fmt.Println("Successfully opened users.json")

	var userList UserList
	if err = json.Unmarshal(file, &userList); err != nil {
		panic(fmt.Sprintf("json unmarshal error: %s", err))
	}

	userList.FormatPrint()
	fmt.Printf("%v\n", userList)
}

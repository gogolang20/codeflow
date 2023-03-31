package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
作业一：解析JSON文件到struct并输出

把下面json保存为“users.json”，然后写一个main.go 从文件读取并解析成user。

	{
	  "users": [
	    {
	      "name": "Elliot",
	      "type": "Reader",
	      "age": 23,
	      "social": {
	        "facebook": "https://facebook.com",
	        "twitter": "https://twitter.com"
	      }
	    },
	    {
	      "name": "Fraser",
	      "type": "Author",
	      "age": 17,
	      "social": {
	        "facebook": "https://facebook.com",
	        "twitter": "https://twitter.com"
	      }
	    }
	  ]
	}

执行结果如下
$ go run main.go
Successfully opened users.json
User Type: Reader
User Age: 23
User Name: Elliot
Facebook Url: https://facebook.com
User Type: Author
User Age: 17
User Name: Fraser
Facebook Url: https://facebook.com
{[{Elliot Reader 23 {https://facebook.com https://twitter.com}} {Fraser Author 17 {https://facebook.com https://twitter.com}}]}
*/
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
	file, err := os.ReadFile(filePath)
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

package main

import (
	"flag"
	"fmt"
	"os"
)

func isValide() {
	if len(os.Args) < 2 {
		fmt.Println("less args")
		os.Exit(1)
	}
}

func main() {
	isValide()
	StrCmd := flag.NewFlagSet("get", flag.ExitOnError)

	str := StrCmd.String("data", "默认", "获取数据")
	switch os.Args[1] {
	case "get":
		err := StrCmd.Parse(os.Args[2:])
		if err != nil {
			os.Exit(1)
		}
	default:
		fmt.Println("unknow")
	}

	if StrCmd.Parsed() {
		if *str == "" {
			os.Exit(1)
		}
		fmt.Println(*str)
	}
}

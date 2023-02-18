package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	// 通过 go 向 redis 写入输入和读取数据
	// 1 链接到 redis
	conn, err := redis.Dial("tcp", "192.168.0.101:6379") // 192.168.0.101:6379
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()
	// fmt.Println("conn=", conn)

	// 2 通过 go 向 redis 写入数据 string key:value
	_, err = conn.Do("hmset", "user02", "name", "jack", "age", 32)
	if err != nil {
		fmt.Println("hmset err=", err)
		return
	}

	// 3 通过 go 向 redis 读取数据 string key:value
	r, err := redis.Strings(conn.Do("hmget", "user02", "name", "age"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}

	// fmt.Println("r=", r) //r= [jack 32]
	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}

	fmt.Println("success")

}

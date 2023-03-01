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
	_, err = conn.Do("hset", "user01", "name", "tom")
	if err != nil {
		fmt.Println("hset err=", err)
		return
	}
	_, err = conn.Do("hset", "user01", "age", 18)
	if err != nil {
		fmt.Println("hset err=", err)
		return
	}

	// 3 通过 go 向 redis 读取数据 string key:value
	r1, err := redis.String(conn.Do("hget", "user01", "name"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}
	r2, err := redis.Int(conn.Do("hget", "user01", "age"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}

	// 不需要使用类型断言 会报错 k := r.(string)
	fmt.Printf("r1=%v r2=%v\n", r1, r2)

	fmt.Println("success")

}

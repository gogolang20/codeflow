package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// 定义一个全局的pool
var pool *redis.Pool

// 当启动程序时 就初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8, // 最大空闲连接池
		MaxActive:   0,
		IdleTimeout: 100,
		/*
			func Dial(network, address string, options ...DialOption) (Conn, error) {
				return DialContext(context.Background(), network, address, options...)
			}
		*/
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	// 先从pool 取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("set", "name", "汤姆猫")
	if err != nil {
		fmt.Println("conn.Do set err: ", err)
		return
	}

	// 取出
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("conn.Do get err: ", err)
		return
	}
	fmt.Println("r: ", r)

	// 如果要从pool 取出链接 一定要确保 pool 没有关闭

}

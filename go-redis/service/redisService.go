package service

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"golang-leaning/go-redis/init"
)

var conn = init.InitPool().Get()

// 添加键值
func Put(key, value string) {
	//写入值{"test-Key":"test-Value"}
	_, err := conn.Do("SET", key, value)
	if err != nil {
		fmt.Println("redis set value failed >>>", err)
	}
}

//检查是否存在key值
func Exists(key string) bool {
	//检查是否存在key值
	exists, err := redis.Bool(conn.Do("EXISTS", "test-Key"))
	if err != nil {
		fmt.Println("illegal exception")
	}
	fmt.Printf("exists or not: %v \n", exists)
	return exists
}

// 获取值
func GetValueByKey(key string) string {
	//read value
	v, err := redis.String(conn.Do("GET", "test-Key"))
	if err != nil {
		fmt.Println("redis get value failed >>>", err)
	}
	return v
}

// 删除key
func Delete(key string) {
	//del kv
	_, err := conn.Do("DEL", "test-Key")
	if err != nil {
		fmt.Println("redis delelte value failed >>>", err)
	}
}

// 给定一个kv的过期时间
func PutTime(key, value, time string) {
	//EX,5秒
	_, err := conn.Do("SET", key, value, "EX", time)
	if err != nil {
		fmt.Println("redis set value failed >>>", err)
	}
}

package init

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

//声明一些全局变量
var (
	pool     *redis.Pool
	server   = "www.gochanghai.com:6379"
	password = "lch2199."
)

//初始化一个pool
func InitPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3, /*最大的空闲连接数*/
		MaxActive:   5, /*最大的激活连接数*/
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

package initialize

import (
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"process_data/common"
	"process_data/config"
	"time"
)

func InitRedis() error {
	common.RedisPool = &redis.Pool{
		MaxIdle:     config.GConfig.Redis.MaxConn,
		MaxActive:   config.GConfig.Redis.MaxConn,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.GConfig.Redis.Host+":"+config.GConfig.Redis.Port)
			if err != nil {
				log.Println("cache Dial failed: ", err.Error())
				return nil, err
			}
			if config.GConfig.Redis.Password != "" {
				if _, err := c.Do("AUTH", config.GConfig.Redis.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return err
			}
			return err
		},
	}
	log.Printf("InitRedis Success.")
	return nil
}

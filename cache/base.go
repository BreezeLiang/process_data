package cache

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"log"
	"process_data/common"
	"process_data/config"
)

var (
	ErrInvalidParam           = errors.New("Invalid Parameters")
	ErrZRevRangeByScoreFailed = errors.New("ZREVRANGEBYSCORE failed")
	ErrZRangeByScoreFailed    = errors.New("ZRANGEBYSCORE failed")
)

type RedisClient struct {
	pool   *redis.Pool
	prefix string
}

func Cache() *RedisClient {
	return &RedisClient{
		pool:   common.RedisPool,
		prefix: config.GConfig.Redis.Prefix,
	}
}

func (c *RedisClient) SET(key, value string) error {
	conn := c.pool.Get()
	defer conn.Close()

	_, err := conn.Do(common.CommandSET, c.prefix+key, value)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

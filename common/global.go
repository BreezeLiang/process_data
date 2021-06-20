package common

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
)

var (
	Mysqlx    *sqlx.DB
	RedisPool *redis.Pool
)

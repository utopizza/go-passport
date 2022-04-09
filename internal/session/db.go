package session

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)

var db *redis.Client

func init() {
	db = createMiniRedis()
}

func GetRDB() *redis.Client {
	return db
}

func createMiniRedis() *redis.Client {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	return rdb
}

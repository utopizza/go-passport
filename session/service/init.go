package service

import "github.com/go-redis/redis/v8"

var defaultRDB *redis.Client

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defaultRDB = rdb
}

func GetRDB() *redis.Client {
	return defaultRDB
}

package app

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/Unknwon/goconfig"
)

var	REDIS redis.Conn

func initRedis() {
	REDIS = redis_connect()
}
 
func closeRedis() {
	REDIS.Close()
}
 
func redis_connect() redis.Conn {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		panic("get configuration failed")
	}
	host, err := cfg.GetValue("redis", "host")
	if err != nil {
		fmt.Println("redis host error")
	}
	port, err := cfg.GetValue("redis", "port")
	if err != nil {
		fmt.Println("redis port error")
	}	
	server := fmt.Sprintf("%s:%s", host, port)
	c, err := redis.Dial("tcp", server)
	if err != nil {
		fmt.Println("connect to redis error"+err.Error())
	}
	return c
}

func RedisGet(key string) string {
	initRedis()
	defer closeRedis()
	value, err := redis.String(REDIS.Do("TYPE", key))
    if err != nil {
        fmt.Println("redis get failed:", err)
	}
	if value == "string" {
		v, _ := redis.String(REDIS.Do("GET", key))
		return v
	}
	return "目前只支持string类型"
}

func RedisKeys() []string {
	initRedis()
	val, err := redis.Strings(REDIS.Do("KEYS", "*"))
	if err != nil {
        fmt.Println("redis get failed:", err)
	}
	closeRedis()
	return val
}

func RedisKeysFilter(filter string) []string {
	initRedis()
	val, err := redis.Strings(REDIS.Do("KEYS", filter+"*"))
	if err != nil {
        fmt.Println("redis get failed:", err)
	}
	closeRedis()
	return val
}
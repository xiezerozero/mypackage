package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"strings"
)

func main() {
	//optObject()
	Exp()
}

func Exp() {
	s := []rune("世界世界世界")
	first3 := string(s[0:3])
	last3 := string(s[len(s)-3:])
	fmt.Println(first3, last3)
}

// 10万个数据拆分到1000个hash中,每个hash存100个
func optObject() {
	client := connectLocalRedis()
	for i := 0; i < 100000; i++ {
		key, field := keyAndField(fmt.Sprintf("object:%d", i))
		client.HSet(key, field, "val")
	}
}

func keyAndField(key string) (string, string) {
	keyAndField := strings.Split(key, ":")
	if len(keyAndField[1]) > 2 {
		return keyAndField[0] + ":" + keyAndField[1][:len(keyAndField[1])-2], keyAndField[1][len(keyAndField[1])-2:]
	}
	return keyAndField[0], keyAndField[1]
}

// 10万个单key
func noOptObject() {
	client := connectLocalRedis()
	for i := 0; i < 100000; i++ {
		key := fmt.Sprintf("object:%d", i)
		client.Set(key, "val", 0)
	}
}

func connectLocalRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)
	return rdb
}

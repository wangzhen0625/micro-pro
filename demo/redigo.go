package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var redisConn redis.Conn

// This example implements ZPOP as described at
// http://redis.io/topics/transactions using WATCH/MULTI/EXEC and scripting.
func main() {
	var err error
	redisConn, err = redis.Dial("tcp", "192.168.74.84:6379")

	if err != nil {
		log.Fatal(err)
	}
	count := 10000
	defer redisConn.Close()
	for i := 0; i < count; i++ {
		go doRedis(i)
	}

	time.Sleep(10 * time.Second)

}

func doRedis(i int) {
	redisConn.Do("PUBLISH", "wz", i)
}

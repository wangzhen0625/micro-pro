package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// This example implements ZPOP as described at
// http://redis.io/topics/transactions using WATCH/MULTI/EXEC and scripting.
func main() {
	c, err := redis.Dial("tcp", "192.168.74.50:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	v, err := c.Do("SET", "name", "red")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	v, err = redis.String(c.Do("GET", "a"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)

}

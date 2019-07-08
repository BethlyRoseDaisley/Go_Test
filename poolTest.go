package main

import(
	"fmt"
	"time"
	"./redigo/redis"
)

func newRedisPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 1,
		MaxActive: 3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func () (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

var (
	redisPool *redis.Pool
	redisServer = "127.0.0.1:6379"
)

func TestPing(){
	conn := redisPool.Get()
	res, _ := conn.Do("ping")
	fmt.Println(res.(string))
}

func TestPoolMaxIdle() {

	go func(){
		fmt.Println("go---1---")
		conn := redisPool.Get()
		fmt.Println("go---1---: Get Conn")
		defer conn.Close()

		time.Sleep(5*time.Second)
		res, err := conn.Do("ping")
		if err != nil {
			fmt.Println("go---1---error:", err)
			return
		}
		fmt.Println("go---1---:", res.(string))
	}()

	go func(){
		fmt.Println("go---2---")
		conn := redisPool.Get()
		fmt.Println("go---2---: Get Conn")
		defer conn.Close()

		time.Sleep(5*time.Second)
		res, err := conn.Do("ping")
		if err != nil {
			fmt.Println("go---1---error:", err)
			return
		}
		fmt.Println("go---2---:", res.(string))
	}()

	go func(){
		fmt.Println("go---3---")
		conn := redisPool.Get()
		fmt.Println("go---3---: Get Conn")
		defer conn.Close()

		time.Sleep(5*time.Second)
		res, err := conn.Do("ping")
		if err != nil {
			fmt.Println("go---1---error:", err)
			return
		}
		fmt.Println("go---3---:", res.(string))
	}()
}

func main()  {
	redisPool = newRedisPool(redisServer)
	fmt.Println("-------------TestPing------------")
	TestPing()
	fmt.Println("-------------TestPoolMaxIdle------------")
	TestPoolMaxIdle()

	time.Sleep(20*time.Second)
}
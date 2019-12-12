package main

import(
	"github.com/gomodule/redigo/redis"
	"errors"
	"fmt"
	"time"
)


func newRedisPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func () (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

func Lock(conn redis.Conn, expire int) {
	for {
		err := TryLock(conn, expire)
		fmt.Println(err)
		if err == nil {
			return
		}
		time.Sleep(10*time.Millisecond)
	}
}

func TryLock(conn redis.Conn, expire int) error {
	err := conn.Send("select", 0) 
	if err != nil {
		return err
	}

	err = conn.Send("setnx", "room_lock", "1")
	if err != nil {
		return err
	}

	err = conn.Flush()
	if err != nil {
		return err
	}

	_, err = redis.ReceiveWithTimeout(conn, 3*time.Second)
	if err != nil {
		return err
	}

	ok, err := redis.ReceiveWithTimeout(conn, 3*time.Second)
	if err != nil {
		return err
	}

	if ok.(int64) != 1 {
		return errors.New("key is locked")
	}

	if expire > 0 {
		_, err = conn.Do("expire", "room_lock", expire)
		if err != nil {
			return err
		}
	}

	return nil
}

func Unlock(conn redis.Conn) error {
	err := conn.Send("select", 0)
	if err != nil {
		return err
	}

	err = conn.Send("del", "room_lock")
	if err != nil {
		return err 
	}

	err = conn.Flush()
	if err != nil {
		return err
	}

	_, err = redis.ReceiveWithTimeout(conn, 3*time.Second)
	if err != nil {
		return err
	}

	ok, err := redis.ReceiveWithTimeout(conn, 3*time.Second)
	if err != nil {
		return err
	}

	if ok.(int64) != 1 {
		return errors.New("key del failed unknow error.")
	}

	return nil
}

/*
func main() {

	fmt.Println("-----start--------")

	rp := newRedisPool("127.0.0.1:6379")

	conn := rp.Get()
	defer conn.Close()

	//lock
	fmt.Println(TryLock(conn, 5))

	//lock an lock key
	time.Sleep(1*time.Second)
	fmt.Println(TryLock(conn, 5))

	//expire lock
	time.Sleep(5*time.Second)
	fmt.Println(TryLock(conn, 5))

	//unlock
	time.Sleep(1*time.Second)
	fmt.Println(UnLock(conn))

	//lock
	time.Sleep(1*time.Second)
	fmt.Println(TryLock(conn, 5))

	time.Sleep(time.Hour)
}*/

func main() {

	fmt.Println("-----start--------")

	rp := newRedisPool("127.0.0.1:6379")

	conn := rp.Get()
	defer conn.Close()

	//lock
	Lock(conn, 3)

	//lock an lock key
	Lock(conn, 3)

	//unlock
	fmt.Println(Unlock(conn))

	//lock
	Lock(conn, 3)

	time.Sleep(time.Hour)
}
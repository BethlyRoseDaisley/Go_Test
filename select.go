package main

import (
	"fmt"
	"time"
)

func wChan1(ch chan string) {
	for {
		time.Sleep(3000 * time.Millisecond)
		ch <- "msg from chan1"
	}
}

func wChan2(ch chan string) {
	for {
		time.Sleep(5000 * time.Millisecond)
		ch <- "msg from chan2"
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go wChan1(c1)
	go wChan2(c2)

	for {
		select {
		case r1 := <-c1:
			fmt.Println(r1)
		case r2 := <-c2:
			fmt.Println(r2)
		}
	}
}
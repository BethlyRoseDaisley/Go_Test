package main

import(
	"time"
	"fmt"
)

func main(){
	fmt.Printf("%02d%02d%02d %02d:%02d:%02d\n", time.Now().Year(),
        time.Now().Month(),
        time.Now().Day(),
        time.Now().Hour(),
        time.Now().Minute(),
		time.Now().Second())
		

	fmt.Println(time.Now().Format("2006/1/02 15:04:05"))
	fmt.Println(time.Now().Format(time.UnixDate))



	go func() {
		tick1 := time.Tick(1000 * time.Millisecond)
		tick2 := time.After(5000 * time.Millisecond)
		cnt := 0
		for {
			select {
			case <-tick1:
				fmt.Println("tick happen / 1s")
				cnt = cnt + 1
				fmt.Println("cnt : ", cnt)
				if cnt > 10 {
					return
				}
			case <-tick2:
				fmt.Println("tick happen once 5s")
			}
		}
	}()

	for {//必须 x time.Second 默认单位是ns， 习惯性写1，会让cpu%100 哈。
		time.Sleep(1*time.Second)
	}
}
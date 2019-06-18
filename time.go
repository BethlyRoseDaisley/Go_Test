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



	go func()  {
		tick := time.Tick(1000 * time.Millisecond)
		boom := time.After(5000 * time.Millisecond)
		for {
			select {
			case <-tick:
				fmt.Println("tick.")
			case <-boom:
				fmt.Println("BOOM!")
			default:
				fmt.Println(" .")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	for {
		time.Sleep(1*time.Second)
	}
}
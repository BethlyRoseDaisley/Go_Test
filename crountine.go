/*package main 

import(
	"fmt"
	"time"
)

var share_cnt uint64 = 0

func incrShareCnt() {
	for i:=0; i < 1000000; i++ {
		share_cnt++
	}
	
	fmt.Println(share_cnt)
}

func main()  {
	
	for i:=0; i < 2; i++ {
		go incrShareCnt()
	}


	time.Sleep(1000*time.Second)

}*/

/*package main 

import(
	"fmt"
	"time"
	"strconv"
)

func main() {

	msg_chan := make(chan string)
	done 	 := make(chan bool)


	i := 0

	go func() {
		for  {
			i++
			time.Sleep(1*time.Second)
			msg_chan <- "on message"
			<- done
		}
	}()

	go func() {
		for {
			select {
			case msg := <- msg_chan :
				i++
				fmt.Println(msg + " " + strconv.Itoa(i))
				time.Sleep(2*time.Second)
				done <- true
			}
		}

	}()


	time.Sleep(20*time.Second)
}*/

/*
package main

import(
	"fmt"
	"time"
	"sync"
)

var share_cnt uint64 = 0

var lck sync.Mutex

func incrShareCnt() {
	for i:=0; i < 1000000; i++ {
		lck.Lock()
		share_cnt++
		lck.Unlock()
	}
	
	fmt.Println(share_cnt)
}

func main()  {
	
	for i:=0; i < 2; i++ {
		go incrShareCnt()
	}


	time.Sleep(1000*time.Second)

}*/

/*package main

import(
	"sync"
	"net/http"
)

var wg sync.WaitGroup
var urls = []string{
    "http://www.baidu.com/",
    "http://www.taobao.com/",
    "http://www.tianmao.com/",
}
func main() {

for _, url := range urls {
    // Increment the WaitGroup counter.
    wg.Add(1)
    // Launch a goroutine to fetch the URL.
    go func(url string) {
        // Decrement the counter when the goroutine completes.
        defer wg.Done()
        // Fetch the URL.
        http.Get(url)
    }(url)
}
// Wait for all HTTP fetches to complete.
wg.Wait()

}*/
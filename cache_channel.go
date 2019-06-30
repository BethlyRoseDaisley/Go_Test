/*package main 

import(
	//"fmt"
	"time"
	"strconv"
	log "github.com/astaxie/beego/logs"
)

func main() {
	log.Debug("---main--- start")

	msgs := make(chan string, 10)

	i := 0

	go func() {
		time.Sleep(3*time.Second)

		for {
			time.Sleep(1*time.Second)
			i++
			msg := "msg " + strconv.Itoa(i)
			msgs <- msg
			log.Debug("------ put msg : ", msg)
		}
	}()

	go func() {

		for {

			log.Debug("---------------- to  get : ")

			get := <- msgs

			log.Debug("---------------- pop msg : ", get)

			time.Sleep(3*time.Second)
		}

	}()

	time.Sleep(100*time.Second)

}*/



package main 

import(
	//"fmt"
	"time"
	"strconv"
	log "github.com/astaxie/beego/logs"
)

func main() {
	log.Debug("---main--- start")

	msgs := make(chan string, 3)

	i := 0

	go func() {
		time.Sleep(3*time.Second)

		for {

			time.Sleep(1*time.Second)
			i++
			msg := "msg " + strconv.Itoa(i)

			select {
				case msgs <- msg:
					log.Debug("------ put msg : ", msg)
				default :
					log.Debug("-----msgs chan cache full sleep 1s-----")
					log.Debug("-----ignore this msg-----> : ", msg)
			}
		}
	}()

	go func() {

		for {

			get := <- msgs

			log.Debug("---------------- pop msg : ", get)

			time.Sleep(3*time.Second)
		}

	}()

	time.Sleep(100*time.Second)

}
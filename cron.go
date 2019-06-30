package main 

import(
	"time"

	log "github.com/astaxie/beego/logs"
)

/*
type Cron struct {
	entries []*Entry
	stop 	chan struct{}
	add 	chan *Entry
	running bool
}*/


func main()  {
	
	log.Debug("---------main---------")

	timer := time.NewTicker(time.Second)

	for {
		select {
		case <- timer.C :
				log.Debug("Tick happen")
		}
	}

}
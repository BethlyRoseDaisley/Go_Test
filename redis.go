package main

import(
	//"github.com/go-redis/redis"
	"github.com/gomodule/redigo/redis"
	log "github.com/astaxie/beego/logs"
)

func main() {
  client, err := redis.Dial("tcp", "127.0.0.1:6379")
  if err != nil {
      log.Critical("redis dial failed.")
  }
  defer client.Close()

  _, err = client.Do("Select", "3")
  if err != nil {
    log.Critical("redis select failed.")
  }

  res, err := redis.String(client.Do("Get", "sip::accounts_info:accounts:ailumiyana"))
  if err != nil {
  	log.Critical("redis set failed.")
  }

  log.Info(res)

}
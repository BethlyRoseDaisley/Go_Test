package main

import (
	"net"
	log "github.com/astaxie/beego/logs"
)

func handleConnection(conn *net.UDPConn) {
	log.Debug("on message")

	buffer := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Error("failed to read UDP msg because of ", err.Error())
		return
	}

	log.Debug("recv ", n, " message from ", /*remoteAddr.String(),*/ ": ", buffer)
}

func main(){
	log.Info("==========system init==========")
	
	localAddress := net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8008};
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8008")
	if err != nil{
		log.Critical("net ResolveUDPAddr Error.")
	}

	log.Debug("local addresses : ", addr.IP, ":", addr.Port)

	log.Debug("net start.")
	conn, err := net.ListenUDP("udp", &localAddress)
	if err != nil {
    log.Critical("net ListenUDP.")
	}
 	log.Debug("net end.")

	defer conn.Close()
 
	for {
		handleConnection(conn)
	}

}
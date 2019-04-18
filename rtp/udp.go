package main

import (
	"net"
	"time"
	"encoding/hex"
	"./rtp"
	log "github.com/astaxie/beego/logs"
)

func handleAudioConnection(conn *net.UDPConn) {

	for {
		log.Debug("on message")

		buffer := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Error("failed to read UDP msg because of ", err.Error())
			return
		}

		parser := rtp.NewParser(buffer, n)

		log.Debug("recv ", n, " message from ", remoteAddr, ": ", hex.EncodeToString(buffer))
	}
}

func handleVedioConnection(conn *net.UDPConn) {
	for {
		log.Debug("on message")

		buffer := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Error("failed to read UDP msg because of ", err.Error())
			return
		}

		log.Debug("recv ", n, " message from ", remoteAddr, ": ", buffer)
	}
}

func main(){
	log.Info("==========system init==========")
	
	//localAddress := net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8008};
	audioAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:8008")
	if err != nil{
		log.Critical("net ResolveUDPAddr Error.")
	}

	vedioAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:8009")
	if err != nil{
		log.Critical("net ResolveUDPAddr Error.")
	}

	log.Debug("local audio addresses : ", audioAddr.IP, ":", audioAddr.Port)
	log.Debug("local vedio addresses : ", vedioAddr.IP, ":", vedioAddr.Port)

	audioConn, err := net.ListenUDP("udp", audioAddr)
	if err != nil {
    log.Critical("net ListenUDP.")
	}

	vedioConn, err := net.ListenUDP("udp", vedioAddr)
	if err != nil {
    log.Critical("net ListenUDP.")
	}

	defer audioConn.Close()
	defer vedioConn.Close()

	log.Debug("rtp serve started.")
 
 	go handleAudioConnection(audioConn)
 	go handleVedioConnection(vedioConn)

	for {
		time.Sleep(10 * time.Second)
	}

}
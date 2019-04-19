package main

import (
	"net"
	"time"
	//"encoding/hex"
	"./rtp"
	log "github.com/astaxie/beego/logs"
)

func handleAudioConnection(conn *net.UDPConn) {
	rtpParser := rtp.NewParser(4096)
	
	for {
		log.Debug("on message")

		n, remoteAddr, err := conn.ReadFromUDP(rtpParser.Buffer())
		if err != nil {
			log.Error("failed to read UDP msg because of ", err.Error())
			return
		}
		rtpParser.SetPacketLength(n);

		log.Debug("recv ", n, " message from ", remoteAddr)//, ": ", hex.EncodeToString(rtpParser.Buffer()))
		rtpParser.Print("rtp audio");
	}
}

func handleVedioConnection(conn *net.UDPConn) {
	rtpParser := rtp.NewParser(4096)

	for {
		log.Debug("on message")

		n, remoteAddr, err := conn.ReadFromUDP(rtpParser.Buffer())
		if err != nil {
			log.Error("failed to read UDP msg because of ", err.Error())
			return
		}
		rtpParser.SetPacketLength(n);

		log.Debug("recv ", n, " message from ", remoteAddr)//, ": ", hex.EncodeToString(rtpParser.Buffer()))
		rtpParser.Print("rtp vedio");

		conn.WriteToUDP(rtpParser.Buffer()[0:n], &net.UDPAddr{IP: net.ParseIP("192.168.0.78"), Port: 1234})
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
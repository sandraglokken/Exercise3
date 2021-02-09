package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var serverPort, pcPort, ip, message string
	serverPort = ":30000"
	pcPort = ":20011"
	ip = "255.255.255.255" + pcPort
	message = "hei server"

	for {
		UDPsend(pcPort, ip, message)
		UDPlisten(serverPort) //returns ip 10.100.23.147:51832
		time.Sleep(time.Second * 2)
	}
}

func UDPlisten(port string) string {
	channel, err := net.ListenPacket("udp", port)
	if err != nil {
		panic(err)
	}
	defer channel.Close()
	buff := make([]byte, 1024)
	n, addr, _ := channel.ReadFrom(buff)
	fmt.Printf("%s sent this : %s\n", addr, buff[:n])
	return addr.String()
}

func UDPsend(port string, ip string, message string) {
	listenAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		panic(err)
	}
	list, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		panic(err)
	}
	defer list.Close()
	addr, err := net.ResolveUDPAddr("udp", ip)
	if err != nil {
		panic(err)
	}
	_, error := list.WriteTo([]byte(message), addr)
	if error != nil {
		panic(error)
	}

}

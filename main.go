package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var serverPortFixed, serverPortMarked, messageTCP, serverPort, serverIP, pcPort, pcIP, messageUDP string
	serverPort = ":30000"
	serverPortFixed = ":34933"
	serverPortMarked = ":33546"
	serverIP = "10.100.23.147"
	pcPort = ":20011"
	pcIP = "255.255.255.255" + pcPort
	messageUDP = "hey server"
	messageTCP = "Connect to: 10.100.23.255:34933"
	fmt.Printf("%s %s%s\n", serverIP, serverPortFixed, serverPortMarked)

	for {
		UDPsend(pcPort, pcIP, messageUDP)
		UDPlisten(serverPort) //returns ip 10.100.23.147:51832
		TCP_connect(serverIP+serverPortFixed, messageTCP)
		time.Sleep(time.Second)
	}
}

func UDPlisten(port string) string {
	channel, err := net.ListenPacket("udp", port)
	checkError(err)
	defer channel.Close()
	buff := make([]byte, 1024)
	n, addr, _ := channel.ReadFrom(buff)
	fmt.Printf("%s sent this : %s\n", addr, buff[:n])
	return addr.String()
}

func UDPsend(port string, ip string, message string) {
	listenAddr, err := net.ResolveUDPAddr("udp", port)
	checkError(err)
	list, err := net.ListenUDP("udp", listenAddr)
	checkError(err)
	defer list.Close()
	addr, err := net.ResolveUDPAddr("udp", ip)
	checkError(err)
	_, err = list.WriteTo([]byte(message), addr)
	checkError(err)

}

func TCP_connect(port string, message string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", port)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	p := make([]byte, 1024)
	n, err := conn.Read(p)
	checkError(err)
	fmt.Printf("Message is: %s\n", p[:n])
	time.Sleep(time.Second)
	conn.Write(append([]byte(message+"\n\r"), 0))
	time.Sleep(time.Second)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

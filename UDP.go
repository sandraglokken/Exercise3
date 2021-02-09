package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	UDPlisten(":30000")

}

func UDPlisten(port string) {
	channel, err := net.ListenPacket("udp", port)
	if err != nil {
		log.Fatal(err)
	}
	buff := make([]byte, 1024)
	_, addr, _ := channel.ReadFrom(buff)
	fmt.Printf(addr.String())
}

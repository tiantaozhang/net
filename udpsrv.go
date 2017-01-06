package main

import (
	"fmt"
	"net"
	"time"
	"os"
)

func main() {
	server := ":1200"
	udpaddr, err := net.ResolveUDPAddr("udp", server)
	checkError(err)
	conn, err := net.ListenUDP("udp", udpaddr)
	checkError(err)
	for {
		handle(conn)
	}

}

func handle(conn *net.UDPConn) {
	var b [512]byte
	n, udpaddr, err := conn.ReadFromUDP(b[:])
	checkError(err)
	fmt.Printf("read from ip:%v ,msg:%v\n", udpaddr, string(b[:n]))
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), udpaddr)
}

func checkError(err error) { if err != nil {
	fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	os.Exit(1)
} }
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	tcpadrr, err := net.ResolveTCPAddr("tcp", ":8888")
	checkError(err)
	tcplistener, err := net.ListenTCP("tcp", tcpadrr)
	checkError(err)
	for {
		conn, err := tcplistener.Accept()
		if err != nil {
			continue
		}
		go handleTcp(conn)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func handleTcp(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(2*time.Minute))
	for {
		var b [512]byte
		n, err := conn.Read(b[:])
		if err != nil {
			fmt.Println(err)
			break
		}
		if n == 0 {
			break // connection already closed by client
		}
		fmt.Printf("read from %v ,msg:%v\n", conn.RemoteAddr(), string(b[:n]))
		conn.Write(b[:n])
	}
}

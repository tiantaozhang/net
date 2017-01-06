package main

import (
	"os"
	"net"
	"fmt"
)

func main() {
	if len(os.Args)!=2{
		panic("args!=2")
	}
	udpaddr,err:=net.ResolveUDPAddr("udp",os.Args[1])
	CheckError(err)
	conn,err:=net.DialUDP("udp",nil,udpaddr)
	CheckError(err)
	_, err = conn.Write([]byte("anything"))
	CheckError(err)
	var buf [512]byte
	n, err := conn.Read(buf[:])
	CheckError(err)
	fmt.Println(string(buf[:n]))
	os.Exit(0)
}
func CheckError(err error) { if err != nil {
	fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	os.Exit(1)
} }
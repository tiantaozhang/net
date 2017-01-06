package main

import (
	"bytes"
	"fmt"
	"io"
	//"io/ioutil"
	"net"
	"os"
	"io/ioutil"
)

func main() {
	addr := ":8888"
	if len(os.Args) == 2 {
		addr = os.Args[1]
	}
	tcpaddr, err := net.ResolveTCPAddr("tcp", addr)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpaddr)
	checkError(err)
	//defer conn.Close()
	_, err = conn.Write([]byte("hello world"))
	checkError(err)
	defer conn.Close()
	conn.CloseWrite()
	b,err:=ioutil.ReadAll(conn)
	checkError(err)
	fmt.Printf("read from server:%v\n", string(b))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}



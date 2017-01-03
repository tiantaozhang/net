package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	defer catchPanic()
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "params err, %s\n", strings.Join(os.Args, ","))
		os.Exit(1)
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkErr(err)
	//HEAD 请求，只返回报文头
	n, err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErr(err)
	fmt.Println("write ", n, "byte")

	result, err := readFully(conn)
	checkErr(err)
	fmt.Println(string(result))

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func catchPanic() {
	if err := recover(); err != nil {
		fmt.Println("Panic info is: ", err)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil

}

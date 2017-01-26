package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

func main() {
	buffer := make([]byte, 1024)
	ServerIP := "129.241.187.43"
	ServerPort := ":33546"
	addr, _ := net.ResolveTCPAddr("tcp", ServerIP+ServerPort)
	conn, _ := net.DialTCP("tcp", nil, addr)
	addr1, _ := net.ResolveTCPAddr("tcp", "129.241.187.161:30008")
	_, _ = conn.Write([]byte("Connect to: 129.241.187.161:30008\x00"))
	listener, _ := net.ListenTCP("tcp", addr1)
	conn1, _ := listener.Accept()

	i := 0
	for {
		i++
		time.Sleep(2 * time.Second)
		_, _ = conn1.Write([]byte("JALLAJALLA\x00"))
		n, _ := conn1.Read(buffer)
		fmt.Println(string(buffer[0:n]))
		if i == 5 {
			conn.Close()
			conn1.Close()
			os.Exit(0)
		}
	}
	os.Exit(0)
}

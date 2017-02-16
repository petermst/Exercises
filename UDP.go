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
	ListenerAddress := "127.0.0.1:30008"
	SenderAddress := ":30008"
	//129.241.187.43
	address, _ := net.ResolveUDPAddr("udp", ListenerAddress)
	listener, _ := net.ListenUDP("udp", address)

	addr, _ := net.ResolveUDPAddr("udp", SenderAddress)
	conn, err := net.DialUDP("udp", nil, addr)

	fmt.Printf("%s\n", listener.LocalAddr())
	fmt.Printf("%s\n", conn.RemoteAddr())

	checkError(err)
	//sendMSG := make([]byte, 1024)

	for {

		time.Sleep(time.Second)

		_, err := conn.Write([]byte("Trololol"))

		checkError(err)

		n, addr, _ := listener.ReadFromUDP(buffer)

		fmt.Println(string(buffer[0:n]))
		fmt.Printf("%s\n", addr)
	}
	os.Exit(0)
}

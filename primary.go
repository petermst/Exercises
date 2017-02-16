package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

func main() {

	fmt.Printf("I am Primary now. \n")

	//cmd := exec.Command("gnome-terminal", "-e", "go run backup.go")
	//err := cmd.Run()

	SenderAddress := "129.241.187.255:30008"
	addr, _ := net.ResolveUDPAddr("udp", SenderAddress)
	conn, err := net.DialUDP("udp", nil, addr)
	checkError(err)

	value := 0

	for {
		value++
		fmt.Println(value)

		_, err := conn.Write([]byte("AAAA"))
		checkError(err)

		time.Sleep(time.Second)
	}
}

package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

func main() {

	last_value := backup()

	primary(last_value)

	os.Exit(0)
}

func primary(value int) {

	fmt.Printf("I am Primary now. \n")

	cmd := exec.Command("gnome-terminal", "-e", "go run processpairs.go")
	err := cmd.Run()

	SenderAddress := "129.241.187.255:30008"

	addr, _ := net.ResolveUDPAddr("udp", SenderAddress)
	conn, err := net.DialUDP("udp", nil, addr)

	checkError(err)

	for {
		value++

		fmt.Println(value)

		_, err := conn.Write([]byte(strconv.Itoa(value)))

		checkError(err)

		time.Sleep(time.Second)
	}
}

func backup() int {
	value := 0
	buffer := make([]byte, 1024)
	ListenerAddress := ":30008"

	address, _ := net.ResolveUDPAddr("udp", ListenerAddress)
	listener, err := net.ListenUDP("udp", address)
	checkError(err)

	fmt.Printf("I am backup\n")

	for {
		listener.SetReadDeadline(time.Now().Add(3 * time.Second))
		n, _, err := listener.ReadFromUDP(buffer)

		if err != nil {
			listener.Close()
			return value
		}
		value, _ = strconv.Atoi(string(buffer[0:n]))
	}
}

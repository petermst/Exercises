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

	buffer := make([]byte, 1024)
	ListenerAddress := ":30008"

	address, _ := net.ResolveUDPAddr("udp", ListenerAddress)
	listener, err := net.ListenUDP("udp", address)
	checkError(err)

	fmt.Printf("I am backup\n")
	//1 listen
	//2 if stop, return last value

	listener.SetReadDeadline(time.Now().Add(3 * time.Second))

	for {

		n, _, err := listener.ReadFromUDP(buffer)

		if err != nil {
			cmd := exec.Command("gnome-terminal", "-e", "go run primary.go")
			cmd.Run()
			os.Exit(1)
		}
		fmt.Println(string(buffer[0:n]))
	}

}

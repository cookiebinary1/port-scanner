package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	done := make(chan bool)
	proc_num := 0

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go example.com")
		os.Exit(0)
	}
	host := os.Args[1]

	fmt.Printf("Port Scanning on host %s (waiting delay max 1s)..\n", host)

	for port := 1; port < 65530; port++ {
		proc_num++
		go func(port int) {
			open := scanPort("tcp", host, port)
			if open {
				fmt.Printf("Port %d Open: %t\n", port, open)
			}
			done <- true
		}(port)
	}

	for i := 0; i < proc_num; i++ {
		<-done
	}

	fmt.Print("Finished\n")
}

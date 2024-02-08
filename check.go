package main

import (
	"fmt"
	"net"
	"time"
)

func healthCheck(destination string, port string) string {
	address := destination + ":" + port
	timeout := time.Second * 5

	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable \nError: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable \nFrom: %v \nTo: %v", destination,
			conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}

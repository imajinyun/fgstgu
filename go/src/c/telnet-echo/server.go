package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func server(address string, exitChan chan int) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}
	fmt.Println("Listen:", address)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go handleSession(conn, exitChan)
	}
}

func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started:")
	reader := bufio.NewReader(conn)
	for {
		str, err := reader.ReadString('\n')
		if err == nil {
			str = strings.TrimSpace(str)
			if !processTelnetCommand(str, exitChan) {
				conn.Close()
				break
			}
			conn.Write([]byte(str + "\r\n"))
		} else {
			fmt.Println("Session closed")
			conn.Close()
			break
		}
	}
}

func processTelnetCommand(str string, exitChan chan int) bool {
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")
		return false
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")
		exitChan <- 0
		return false
	}
	fmt.Println(str)
	return true
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	conn, err := net.Listen("tcp", ":3333")
	if err != nil {
		fmt.Println("Error 1: ", err)
	}
	c, err := conn.Accept()
	if err != nil {
		fmt.Println("Error 2:", err)
		return
	}
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("Error 3: ", err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("TCP Server Exiting ...")
			return
		}
		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		_, _ = c.Write([]byte(myTime))
	}
}

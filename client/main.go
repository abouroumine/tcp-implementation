package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	ADDRESS = "localhost:3333"
)

func main() {

	c, err := net.Dial("tcp", ADDRESS)

	if err != nil {
		fmt.Println("Error 1 is: ", err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error 2 is: ", err)
		}
		_, _ = fmt.Fprintf(c, text+"\n")
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP Client Existing ...")
			return
		}
	}

}

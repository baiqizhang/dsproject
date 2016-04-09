package main

import (
	"bufio"
	"dsproject/util"
	"fmt"
	"io"
	"strings"
)

func handleNode(c client) {
	fmt.Println(c.conn.RemoteAddr().String())
	reader := bufio.NewReader(c.conn)

	// Read handler
	for {
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		util.CheckError(err)

		fmt.Println("[Node Message]:" + message)

		words := strings.Split(message, " ")

		// if connection comes from CarNode
		if words[0] == "NAME" {
			fmt.Println("[Node] Register Name:" + words[1])
			c.name = words[1]
		}
		fmt.Println("error: message not recognized")
	}
}

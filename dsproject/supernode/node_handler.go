package main

import (
	"bufio"
	"dsproject/util"
	"fmt"
	"io"
	"strings"
)

func handleNode(client util.Client) {
	fmt.Println(client.Conn.RemoteAddr().String())
	reader := bufio.NewReader(client.Conn)

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
			client.Name = words[1]
		}
		fmt.Println("error: message not recognized")
	}
}

package main

import (
	"bufio"
	"dsproject/util"
	"fmt"
	"io"
	"strings"
)

func handleClient(client util.Client) {
	conn := client.Conn
	fmt.Println("[New Client]:" + conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)

	// Read handler
	for {
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		util.CheckError(err)

		fmt.Print("[Message Received]:" + message)

		words := strings.Split(message, " ")

		// if connection comes from CarNode
		if words[0] == "COMPUTERESULT" {
			fmt.Println("[COMPUTERESULT]:" + words[1])
			break
		}
		fmt.Println("error: message not recognized")
	}
}

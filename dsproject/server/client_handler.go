package main

import (
	"bufio"
	"dsproject/util"
	"fmt"
	"io"
	"strings"
)

func handleClient(client *util.Client) {
	conn := client.Conn
	fmt.Println("[New Client]:" + conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
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
		if words[0] == "SUPERNODE" {
			client.Type = "SUPERNODE"
			if words[1] == "REGISTER" {
				client.Name = words[2]
				continue
			}
		}
		if words[0] == "NODE" {
			client.Type = "NODE"
			if words[1] == "REGISTER" {
				fmt.Println("[Node Register] send hardcoded supernode addr")
				writer.WriteString("127.0.0.1:6060\n")
				writer.Flush()
				client.Conn.Close()
				return
			}
		}
		fmt.Println("error: message not recognized")
	}
}

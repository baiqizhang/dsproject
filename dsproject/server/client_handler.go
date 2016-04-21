package main

import (
	"bufio"
	"dsproject/util"
	"fmt"
	"io"
	"net"
	"strings"
)

var lastClient *util.Client // = nil
var superNodeCount int

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
		client.Type = words[0]
		if words[0] == "SUPERNODE" {
			if words[1] == "REGISTER" {
				client.Name = words[2]
				if lastClient != nil {
					//tell the newcomer the last supernode's address
					addr := lastClient.Conn.RemoteAddr()
					writer.WriteString("PEERADDR " + addr.(*net.TCPAddr).IP.String() + ":" + lastClient.Name)
					writer.Flush()

					if superNodeCount == 1 {
						tempWriter := bufio.NewWriter(lastClient.Conn)
						addr := client.Conn.RemoteAddr()
						tempWriter.WriteString("PEERADDR " + addr.(*net.TCPAddr).IP.String() + ":" + words[2])
						tempWriter.Flush()
					}
				}
				lastClient = client
				superNodeCount++
				continue
			}
		}
		if words[0] == "NODE" {
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

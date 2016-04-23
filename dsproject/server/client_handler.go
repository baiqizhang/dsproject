package main

import (
	"bufio"
	"dsproject/util"
	"fmt"
	"io"
	"net"
	"strings"
)

var lastSNClient *util.Client // = nil
var superNodeCount int

// var superNodeAliveCounter = make(map[*util.Client]int)

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

		if util.Verbose == 1 {
			fmt.Print("[Message Received]:" + message)
		}

		words := strings.Split(message, " ")

		// if connection comes from CarNode
		client.Type = words[0]
		if words[0] == "SUPERNODE" {
			//mark the client as alive
			// superNodeAliveCounter[&client] = 5

			// check the 2nd arg
			if words[1] == "REGISTER" {
				client.Name = words[2]
				if lastSNClient != nil {
					//tell the newcomer the last supernode's address
					addr := lastSNClient.Conn.RemoteAddr()
					writer.WriteString("PEERADDR " + addr.(*net.TCPAddr).IP.String() + ":" + lastSNClient.Name)
					writer.Flush()

					if superNodeCount == 1 {
						tempWriter := bufio.NewWriter(lastSNClient.Conn)
						addr := client.Conn.RemoteAddr()
						tempWriter.WriteString("PEERADDR " + addr.(*net.TCPAddr).IP.String() + ":" + words[2])
						tempWriter.Flush()
					}
				}
				lastSNClient = client
				superNodeCount++
				continue
			}
			if words[1] == "HEARTBEAT" {
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

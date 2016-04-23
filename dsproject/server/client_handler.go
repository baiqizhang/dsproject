package main

import (
	"bufio"
	"dsproject/util"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

var lastSNClient *util.Client // = nil
var superNodeCount int

var superNodeAliveCounter = make(map[string]int)

//check if connection from SN is broken
func checkConnection() {
	for {
		for key, value := range superNodeAliveCounter {
			superNodeAliveCounter[key] = value - 1
			if util.Verbose == 1 {
				fmt.Printf("Heartbeat count %s = %d\r\n", key, value)
			}
			if value == 0 {
				delete(superNodeAliveCounter, key)
				fmt.Println("[Supernode] Lost connection: " + key)
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}

//handle message from a SN or CarNode
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

		words := strings.Split(strings.Trim(message, "\r\n"), " ")

		// if connection comes from CarNode
		client.Type = words[0]
		if words[0] == "SUPERNODE" {
			newSNAddr := client.Conn.RemoteAddr()
			newSNIPStr := newSNAddr.(*net.TCPAddr).IP.String()
			newSNPortStr := words[2]
			//mark the client as alive
			superNodeAliveCounter[newSNIPStr+":"+newSNPortStr] = 5

			// check the 2nd arg
			if words[1] == "REGISTER" {
				client.Name = newSNPortStr
				if lastSNClient != nil {
					//tell the newcomer the last supernode's ip:port
					lastSNAddr := lastSNClient.Conn.RemoteAddr()
					writer.WriteString("PEERADDR " + lastSNAddr.(*net.TCPAddr).IP.String() + ":" + lastSNClient.Name)
					writer.Flush()

					if superNodeCount == 1 {
						tempWriter := bufio.NewWriter(lastSNClient.Conn)
						tempWriter.WriteString("PEERADDR " + newSNIPStr + ":" + newSNPortStr)
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

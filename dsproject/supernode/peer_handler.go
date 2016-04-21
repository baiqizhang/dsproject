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

func listenPeer(port string) {
	// listen to node connection requests? (not sure if is required)
	listener, err := net.Listen("tcp", ":"+port)
	util.CheckError(err)
	fmt.Println("Supernode Listening at " + port + " for SuperNode connection")
	for {
		conn, err := listener.Accept()
		util.CheckError(err)

		newClient := util.Client{Conn: conn, Name: "none"}

		lastClient = &newClient
		// clients = append(clients, newClient)
		go handlePeer(newClient)

		//TODO handle newly joined peer
	}

}

func handlePeer(client util.Client) {
	fmt.Println("[Peer Listener] new connection from" + client.Conn.RemoteAddr().String())
	reader := bufio.NewReader(client.Conn)

	// Read handler
	for {
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		util.CheckError(err)

		fmt.Println("[Previous Node Message]:" + message)
		words := strings.Split(strings.Trim(message, "\r\n"), " ")

		//TODO processPeerCommand(message)
		if words[0] == "REDIRECT" {

		}

	}
}

func dialPeer(peerAddr string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", peerAddr)
	util.CheckError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	util.CheckError(err)

	reader := bufio.NewReader(conn)
	// Read handler
	for {
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		util.CheckError(err)

		fmt.Println("[Next Peer Message]:" + message)
		//TODO processPeerCommand(message)
	}

}

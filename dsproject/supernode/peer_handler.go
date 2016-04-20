package main

import (
	"bufio"
	"dsproject/util"
	"fmt"
	"io"
	"net"
	"strings"
)

func listenPeer() {
	// listen to node connection requests? (not sure if is required)
	listener, err := net.Listen("tcp", ":5050")
	util.CheckError(err)
	fmt.Println("Supernode Listening at 5050 for CarNode connection")
	for {
		conn, err := listener.Accept()
		util.CheckError(err)

		newClient := util.Client{Conn: conn, Name: "none"}
		clients = append(clients, newClient)
		go handlePeer(newClient)

		//TODO handle newly joined peer
	}

}

func handlePeer(client util.Client) {
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

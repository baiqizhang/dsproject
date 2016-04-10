package main

import (
	"bufio"
	"dsproject/util"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

// Default HTTP Request Handler for UI
func defaultHTTPHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Team 9 %s!</h1>", r.URL.Path[1:])
}

var clients []*util.Client

func main() {
	//start HTTP UI server at 8080
	go func() {
		http.HandleFunc("/", defaultHTTPHandler)
		http.ListenAndServe(":8080", nil)
	}()

	fmt.Print("web server running on 8080\n")

	// start TCP, listening at 7070
	go func() {
		listener, err := net.Listen("tcp", ":7070")
		util.CheckError(err)

		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			newClient := util.Client{Conn: conn, Name: "none"}
			clients = append(clients, &newClient)
			go handleClient(&newClient)
		}
	}()

	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		cmd, _ := stdin.ReadString('\n')
		processCommand(cmd)
	}
}

// process command from Server Terminal
func processCommand(cmd string) {
	args := strings.Split(strings.Trim(cmd, "\r\n"), " ")

	//Compute distance to the customer
	if args[0] == "PICKUP" {
		//Pickup the customer
		source := util.ParseFloatCoordinates(args[1], args[2])
		dest := util.ParseFloatCoordinates(args[3], args[4])
		if source == nil || dest == nil {
			fmt.Println("Error: incorrect PICKUP format:" + cmd)
			return
		}

		var zero []byte
		for _, client := range clients {
			if client.Type == "NODE" {
				continue
			}
			fmt.Println("client " + client.Type + " " + client.Name)
			conn := client.Conn
			fmt.Println(conn.RemoteAddr().String())
			reader := bufio.NewReader(conn)
			_, err := reader.Read(zero)
			if err != nil {
				continue
			}

			fmt.Println("[PICKUP] send to SN:" + client.Conn.RemoteAddr().String())
			writer := bufio.NewWriter(conn)
			writer.WriteString("PICKUP " + args[1] + " " + args[2] + " " + args[3] + " " + args[4] + "\n")
			writer.Flush()
		}
	}
}

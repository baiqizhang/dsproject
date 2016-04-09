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
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Team 9 %s!</h1>", r.URL.Path[1:])
}

var clients []util.Client

func main() {
	//start HTTP UI server at 8080
	go func() {
		http.HandleFunc("/", defaultHandler)
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
			clients = append(clients, newClient)
			go handleClient(newClient)
		}
	}()

	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		cmd, _ := stdin.ReadString('\n')
		processCommand(cmd)
	}
}

func processCommand(cmd string) {
	args := strings.Split(cmd, " ")

	//Compute distance to the customer
	if args[0] == "PICKUP" {
		//Pickup the customer
		source := util.ParseFloatCoordinates(args[1], args[2])
		dest := util.ParseFloatCoordinates(args[3], args[4])
		if source == nil || dest == nil {
			fmt.Println("Error: incorrect PICKUP format")
			os.Exit(0)
		}
		for _, client := range clients {
			conn := client.Conn
			writer := bufio.NewWriter(conn)
			writer.WriteString("PICKUP " + args[1] + " " + args[2] + " " + args[3] + " " + args[4])
		}
	}
}

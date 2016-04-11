package main

import (
	"bufio"
	"dsproject/util"
	"fmt"
	"io"
	"strings"
    "math"
)

//counter for carnodes which are ordinary nodes and counter for supernodes
var countcar int = 0
var countsuper int = 0
var finalResult float = MaxFloat64
var finalClient util.Client

func handleNode(client util.Client) {
	fmt.Println(client.Conn.RemoteAddr().String())
	reader := bufio.NewReader(client.Conn)
    
    message, err := reader.ReadString('\n')
		
		// if connection comes from CarNode
		if message == "CARNODE\n" {
			fmt.Println("this is a " + message)
			countcar++
		}
		
        
       // if connection comes from SuperNode
		if message == "SUPERNODE\n" {
			fmt.Println("this is a " + message)
			countsuper++
		} 

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
		if words[0] == "COMPUTERESULT" {
			fmt.Println("[Node] Register Name:" + words[1])
			client.Name = words[1]
            result, _ := strconv.ParseFloat(words[2], 64)
            
            if(result < finalResult){
                finalResult = result
                finalClient = client
            }
            
		}
       
       
        
		fmt.Println("error: message not recognized")
	}
}

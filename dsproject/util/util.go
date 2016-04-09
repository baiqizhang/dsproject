package util

import (
	"bufio"
	"fmt"
	"os"
)

//Message base message type, not used yet
type Message struct {
	src  string
	kind string
	data string
}

//Node communication unit
type Node struct {
	incoming chan string
	outgoing chan string
	reader   *bufio.Reader
	writer   *bufio.Writer
}

//CheckError just check and print error
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		//os.Exit(1)
	}
}

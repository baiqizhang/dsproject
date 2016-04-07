package util

import (
	"fmt"
	"os"
)

//Message base message type, not used yet
type Message struct {
	src  string
	kind string
	data string
}

//CheckError just check and print error
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		//os.Exit(1)
	}
}

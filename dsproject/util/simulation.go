package util

import (
	"dsproject/util"
	"fmt"
	"time"
)

//DriveCustomer simulate a ride from currentLoc to customerLoc then to destLoc
func DriveCustomer(customerLoc *util.Point, dest *util.Point) {
	idle = false

	// simulate picking up customer
	time.Sleep(1500 * time.Millisecond)

	// update current location
	curLoc = customerLoc
	fmt.Println("Customer picked up")

	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Drop customer")
	curLoc = dest
}

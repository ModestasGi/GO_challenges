package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Print("Enter IP address to check if it's valid: ")
	var ipd string
	fmt.Scanln(&ipd)

	fmt.Println(checkIP(ipd))

}

func checkIP(a string) bool {

	if net.ParseIP(a) == nil {
		return false
	} else {
		return true
	}

}

package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Print("Enter number to determin if it's PRIME number: ")
	var sn string
	fmt.Scanln(&sn)
	var n, _ = new(big.Int).SetString(sn, 0)
	if n.ProbablyPrime(20) {
		fmt.Println(n, " is probably a prime number")
	} else {
		fmt.Println(n, " is probably not a prime number")
	}
}

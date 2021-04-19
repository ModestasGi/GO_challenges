package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first word or phrase: ")
	fw, _ := reader.ReadString('\n')
	fmt.Print("Enter second word or phrase: ")
	sw, _ := reader.ReadString('\n')

	fwc := strings.Split(string(fw), " ")
	swc := strings.Split(string(sw), " ")

	var fwcj string
	for i := 0; i < len(fwc); i++ {
		fwcj = fwcj + fwc[i]
	}

	var swcj string
	for x := 0; x < len(swc); x++ {
		swcj = swcj + swc[x]
	}

	if isAnagram(fwcj, swcj) {
		fmt.Println("Here are anagrams")
	} else {
		fmt.Println("Here are not anagrams")
	}
}

func isAnagram(fs string, ss string) bool {

	a := strings.ToLower(fs)
	b := strings.ToLower(ss)

	ab := []byte(a)
	sort.Slice(ab, func(i int, j int) bool {
		return ab[i] < ab[j]
	})
	bb := []byte(b)
	sort.Slice(bb, func(i int, j int) bool {
		return bb[i] < bb[j]
	})

	c := bytes.Compare(ab, bb)

	if c == 0 {
		return true
	} else {
		return false
	}
}

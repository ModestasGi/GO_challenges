package main

import (
	"bufio"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var logincounter int

func main() {

	fmt.Println("to Log in press 1; to sign in press 2")
	reader := bufio.NewReader(os.Stdin)
	a, _ := reader.ReadString('\n')
	if a == "1\n" {
		logIN()
	}
	if a == "2\n" {
		signIN()
	} else {
		fmt.Println("Selected unknown option.")
		os.Exit(1)
	}
}

func logIN() {

	fmt.Print("your login name: ")
	var ln string
	fmt.Scanln(&ln)

	if !noValidChar(ln) {
		logIN()
	}

	fmt.Print("your login password: ")
	var lp string
	fmt.Scanln(&lp)
	database, _ := sql.Open("sqlite3", "./users.db")
	q := "SELECT password FROM users_list WHERE name = '" + ln + "';"
	rows, _ := database.Query(q)
	var sp string
	for rows.Next() {
		rows.Scan(&sp)
	}

	sh := hashig(lp)

	result := sh == sp

	if result {
		fmt.Println("\nYou are logged in")
		os.Exit(0)
	} else {
		logincounter++
		fmt.Println("\nYour username or login password is incorect \n ")

		if logincounter > 2 {
			fmt.Println("Sorry. You've tried login incorectly too many times.")
			os.Exit(2)
		} else {
			logIN()
		}
	}
}

func signIN() {
	fmt.Print("choose your login name (longer than 5 characters): ")
	var sln string
	fmt.Scanln(&sln)

	database, _ := sql.Open("sqlite3", "./users.db")

	if !noValidChar(sln) {
		signIN()
	}

	q := "SELECT name FROM users_list WHERE name = '" + sln + "';"

	rows, _ := database.Query(q)
	var n string
	for rows.Next() {
		rows.Scan(&n)
	}

	if n != "" {
		fmt.Println("Sorry this login name already exist")
		signIN()
	} else {
		if len(sln) < 5 {
			fmt.Println("Your login name should be longer than 5 characters")
			signIN()
		} else {
			hashedPassword := passwordValidation()
			statement, _ := database.Prepare("INSERT INTO users_list VALUES (?, ?);")
			statement.Exec(sln, hashedPassword)
			fmt.Println("\nYou registered\n ")
		}
	}
	logIN()
}

func noValidChar(s string) bool {
	// add more characters as needed
	c := []string{"'", "&", ";"}
	for i := 0; i < len(c); i++ {
		if strings.Contains(s, c[i]) {
			fmt.Println("found restricted character ->", c[i])
			return false
		}
	}
	return true
}

func passwordValidation() string {
	fmt.Println("\nchoose your login password:\n ")
	fmt.Println("Instructions for password creation:\nIt should be longer then 8 characters\nIt should contain number" +
		"\nIt should contain lower and upper case letters\nIt should contain special characters such as * @ # $ % & ! \n")

	var slp string
	fmt.Scanln(&slp)

	if len(slp) < 8 {
		fmt.Println("Your password should be longer then 8 characters")
		passwordValidation()
	}
	if !containsCharacters(slp) {
		fmt.Println("Your password should contain special characters such as * @ # $ % & !")
		passwordValidation()

	}
	if !containsUpperLower(slp) {
		fmt.Println("Your password should contain lower AND upper case letters")
		passwordValidation()
	}
	if !containNumbers(slp) {
		fmt.Println("Your password should contain numbers")
		passwordValidation()
	}
	return hashig(slp)
}

func hashig(s string) string {
	x := sha256.New()
	x.Write([]byte(s))
	hash256 := x.Sum(nil)

	backs := hex.EncodeToString(hash256[:])

	return backs
}

func containNumbers(s string) bool {
	nu := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	index := 0
	for i := 0; i < len(nu); i++ {
		if strings.Contains(s, nu[i]) {
			index++
		}
	}
	if index == 0 {
		return false
	} else {
		return true
	}
}

func containsCharacters(s string) bool {
	ch := []string{"*", "@", "#", "$", "%", "&", "!"}
	index := 0

	for i := 0; i < len(ch); i++ {
		if strings.Contains(s, ch[i]) {
			index++
		}
	}

	if index > 0 {
		return true
	} else {
		return false
	}
}

func containsUpperLower(s string) bool {
	up := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	lo := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	index_up := 0
	index_lo := 0

	for i := 0; i < len(up); i++ {
		if strings.Contains(s, up[i]) {
			index_up++
		}
	}
	if index_up == 0 {
		return false
	} else {
		for x := 0; x < len(lo); x++ {
			if strings.Contains(s, lo[x]) {
				index_lo++
			}
		}
		if index_lo == 0 {
			return false
		} else {
			return true
		}
	}
}

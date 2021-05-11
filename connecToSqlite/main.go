package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "./bogo.db")

	statement, _ := database.Prepare("CREATE TABLE 'operationaldata' ('period' NUMERIC UNIQUE, 'lowestP'	NUMERIC, " +
		"'openingP'	NUMERIC, 	'closingP'	NUMERIC, 'highestP'	NUMERIC, 'volumesold'	NUMERIC, 'beforeclosingP'	NUMERIC, " +
		"'volumeweightedaverageP'	NUMERIC, PRIMARY KEY(period));")
	statement.Exec()

	statement, _ = database.Prepare("INSERT INTO operationaldata VALUES (?, ?, ?, ?, ?, ?, ?, ?);")
	statement.Exec()	

	resp, err := http.Get("https://bitex.la/api-v1/rest/btc_usd/market/last_24_hours")

	// https://bitex.la/api-v1/rest/btc_usd/market/last_24_hours
	// https://bitex.la/api-v1/rest/btc_usd/market/last_30_days Last 30 Days Candles
	// https://bitex.la/api-v1/rest/btc_usd/market/last_7_days  Last 7 Days Candles

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	s := strings.Split(string(body), "[")
	var sa []string
	var sb []string

	for i := 0; i < len(s); i++ {
		t := strings.Split(s[i], "],")
		sa = append(sa, t[0])
	}

	for x := 0; x < len(sa); x++ {
		u := strings.Split(sa[x], "]]")
		sb = append(sb, u[0])
	}

	for _, n := range sb {
		fmt.Println(n)
	}

	rows, _ := database.Query("SELECT * FROM operationaldata;")
	var period int
	var lowestP int
	var openingP int
	var closingP int
	var highestP int
	var volumesold int
	var beforeclosingP int
	var volumeweightedaverageP int

	for rows.Next() {
		rows.Scan(&period, &lowestP, &openingP, &closingP, &highestP, &volumesold, &beforeclosingP, &volumeweightedaverageP)
		fmt.Println(period, " ", lowestP, " ", openingP, " ", closingP, " ", highestP, " ", volumesold, " ", beforeclosingP, " ", volumeweightedaverageP)
	}

}

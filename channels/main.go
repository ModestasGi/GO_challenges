package main

import (
	"fmt"
	"net/http"
)

// not the best approach, because we need to wait response and only then this code use other link
// need to use go routines

// ending processes with CTRL + C

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) // creating channel, communicating with strings

	for _, link := range links {
		go checkLink(link, c) // creating go routine and passing channel as an argument
	}

	// communicating via channels, receiving messages. loop to create channel for each go routine = part of slice = link

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	/*
		infinit loop should be
		 for {
			 go checkLink(<- c, c)
		 }

		 OR

		 for l := range c {
			go func(link string) {
				time.Sleep(5 * time.Second) // pause for 5 seconds
				checkLink(link, c) // grabing from slice dedicated to channel, here the case of slice "links"
			}
		 } (l) // string copied to memory to be reachable for go routine
	*/

}

// just to check if link is up or down
func checkLink(link string, c chan string) { // adding channel into arguments field
	_, err := http.Get(link) // code frozen at this point while executing -> solution, create go routine and channel (it has some rules)
	if err != nil {
		c <- link + " might be down! \n" // sending message via channel
		return
	}

	c <- link + " is up! \n"

}

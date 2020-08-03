package main

import (
	"flag"

	. "github.com/komisan19/twgo/tweet"
	. "github.com/komisan19/twgo/search"
)

func main() {
	var num string
	get := flag.Bool("g", false, "Get by tweet")
	post := flag.Bool("p", false, "Post by tweet")
	search := flag.Bool("s", false, "Search by tweet")
	flag.Parse()

	if *get {
		Get(num)
		return
	}
	if *post {
		Post()
		return
	}
	if *search {
		Search()
		return
	}
}

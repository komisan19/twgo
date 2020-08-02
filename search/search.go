package search

import(
	"fmt"
	"flag"
	. "github.com/komisan19/twgo/secret"
)

func Search() {
	api := GetTwitterApi()

	text := flag.Arg(0)
	flag.Parse()

	searchResult, _ := api.GetSearch(text, nil)
	for _, tweet := range searchResult.Statuses {
			fmt.Println(tweet.Text)
	}
}
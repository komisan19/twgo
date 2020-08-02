package tweet

import (
	. "github.com/komisan19/twgo/secret"
	. "fmt"
	"flag"
)

func Post() {
	api := GetTwitterApi()

	text := flag.Arg(0)
	flag.Parse()

	tweet, err := api.PostTweet(text, nil)
	if err != nil{
		panic(err)
	}

	Println("ツイートしました:", tweet.Text)
}

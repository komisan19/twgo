package tweet

import (
	. "github.com/komisan19/twgo/secret"
	. "fmt"
)

func Post() {
	api := GetTwitterApi()

	text := "コマンドラインからツイート"
	tweet, err := api.PostTweet(text, nil)
	if err != nil{
		panic(err)
	}

	Println(tweet.Text)
}

package tweet

import (
	. "github.com/komisan19/twgo/secret"
	. "fmt"
	"net/url"
)

func Get() {
  api := GetTwitterApi()

  v := url.Values{}
  v.Set("count","10")

  tweets, err := api.GetHomeTimeline(v)
  if err != nil {
    panic(err)
  }

  for _, tweet := range tweets {
    Println("tweet: ", tweet.Text)
  }

}

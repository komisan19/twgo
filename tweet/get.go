package tweet

import (
  "net/url"
  . "github.com/komisan19/twgo/secret"
	. "fmt"
)

func Get(num string) {
  api := GetTwitterApi()

  v := url.Values{}
  v.Set("count",num)

  tweets, err := api.GetHomeTimeline(v)
  if err != nil {
    panic(err)
  }

  for _, tweet := range tweets {
    Println("tweet: ", tweet.Text)
  }

}

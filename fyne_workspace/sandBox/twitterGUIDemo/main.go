package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	var accessToken = "accessToken"
	var accessTokenSecret = "accessTokenSecret"
	var consumerKey = "consumerKey"
	var consumerSecret = "consumerSecret"

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	v := url.Values{}
	v.Set("count", "3")
	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
		fmt.Printf("Error to getHomeTimeline. err:%v\n", err)
		os.Exit(1)
	}
	for i, tweet := range tweets {
		fmt.Printf("%d, TweetName:%v\nTweet:%v\n\n", i, tweet.User.Name, tweet.FullText)
	}
	for i, tweet := range tweets {
		json, err := json.MarshalIndent(tweet, "", "  ")
		if err != nil {
			fmt.Printf("Error to json Marshal Error. err:%s\n", err)
			os.Exit(1)
		}
		file, err := os.Create(fmt.Sprintf("timeline%d.json", i))
		if err != nil {
			fmt.Printf("Error to fileCreate. err:%v\n", err)
		}
		file.Write(json)
	}

}

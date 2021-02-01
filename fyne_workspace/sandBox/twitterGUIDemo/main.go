package main

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	var accessToken = os.Getenv("twitterAPIAccessToken")
	var accessTokenSecret = os.Getenv("twitterAPIAccessTokenSecret")
	var consumerKey = os.Getenv("twitterAPIConsumerKey")
	var consumerSecret = os.Getenv("twitterAPIConsumerSecret")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	v := url.Values{}
	v.Set("count", "1")

	tweets, err := api.GetHomeTimeline(v)

	if err != nil {
		fmt.Printf("Error to getHomeTimeline. err:%v\n", err)
		os.Exit(1)
	}

	for i, tweet := range tweets {
		createdTime, err := tweet.CreatedAtTime()
		if err == nil {
			jst := time.FixedZone("Asia/Tokyo", 9*60*60)
			createdTimeJST := createdTime.In(jst)
			fmt.Println(createdTimeJST)
		}
		fmt.Printf("%d, TweetName:%v\nTweet:%v\n\n", i, tweet.User.Name, tweet.FullText)
	}
}

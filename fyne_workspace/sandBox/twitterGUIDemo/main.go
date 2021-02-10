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

	// GetHomeTimeline
	// v := url.Values{}
	// v.Set("count", "1")

	// tweets, err := api.GetHomeTimeline(v)

	// if err != nil {
	// 	fmt.Printf("Error to getHomeTimeline. err:%v\n", err)
	// 	os.Exit(1)
	// }

	// for i, tweet := range tweets {
	// 	createdTime, err := tweet.CreatedAtTime()
	// 	if err == nil {
	// 		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	// 		createdTimeJST := createdTime.In(jst)
	// 		fmt.Println(createdTimeJST)
	// 	}
	// 	fmt.Printf("%d, TweetName:%v\nTweet:%v\n\n", i, tweet.User.Name, tweet.FullText)
	// }

	// GetSearch Sample (5秒ごとに指定した単語のツイートを取得して表示する)
	for {
		v := url.Values{}
		v.Set("count", "1")
		searchWord := "グラブル"
		searchResult, err := api.GetSearch(searchWord, v)
		if err != nil {
			panic(err)
		}

		fmt.Println("---")
		if len(searchResult.Statuses) == 0 {
			fmt.Println("No Result.")
		} else {
			for _, tweet := range searchResult.Statuses {
				createdAtTime, err := tweet.CreatedAtTime()
				if err == nil {
					jst := time.FixedZone("Asia/Tokyo", 9*60*60)
					createdAtTimeJST := createdAtTime.In(jst)
					fmt.Print("[CreatedAtTime]:")
					fmt.Println(createdAtTimeJST)
				}

				username := tweet.User.Name
				fmt.Printf("[Username]:%s\n", username)

				fmt.Print(tweet.FullText)
			}
			fmt.Println()
		}
		fmt.Println("---")

		time.Sleep(time.Second * 5)
	}
}

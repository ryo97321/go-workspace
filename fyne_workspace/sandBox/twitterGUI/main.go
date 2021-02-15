package main

import (
	"image/color"
	"net/url"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/ChimeraCoder/anaconda"
)

type MyTweetStruct struct {
	username     string
	createAtTime time.Time
	fullText     string
}

func getTweet() MyTweetStruct {

	var accessToken = os.Getenv("twitterAPIAccessToken")
	var accessTokenSecret = os.Getenv("twitterAPIAccessTokenSecret")
	var consumerKey = os.Getenv("twitterAPIConsumerKey")
	var consumerSecret = os.Getenv("twitterAPIConsumerSecret")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	v := url.Values{}
	v.Set("count", "1")

	myTweetStruct := MyTweetStruct{"", time.Now(), ""}

	searchWord := "グラブル"
	searchResult, err := api.GetSearch(searchWord, v)
	if err != nil {
		return myTweetStruct
	}

	tweetUserName := ""
	tweetCreatedAtTime := time.Now()
	tweetFullText := ""

	if len(searchResult.Statuses) == 0 {
		tweetFullText = "No Result"
	} else {
		for _, tweet := range searchResult.Statuses {
			tweetUserName = tweet.User.Name
			tweetCreatedAtTime, _ = tweet.CreatedAtTime()
			tweetFullText = tweet.FullText
		}
	}

	myTweetStruct.username = tweetUserName
	myTweetStruct.createAtTime = tweetCreatedAtTime
	myTweetStruct.fullText = tweetFullText

	return myTweetStruct
}

func setTweetPer10Seconds(w fyne.Window) {
	for {
		time.Sleep(time.Second * 10)

		myTweetStruct := getTweet()

		fullText := myTweetStruct.fullText
		username := myTweetStruct.username
		createdAtTime := myTweetStruct.createAtTime

		fullTextObject := canvas.NewText(fullText, color.Black)
		usernameTextObject := canvas.NewText(username, color.Black)
		createdAtTimeTextObject := canvas.NewText(createdAtTime.String(), color.Black)

		content := container.New(layout.NewVBoxLayout(), createdAtTimeTextObject, usernameTextObject, fullTextObject)

		w.SetContent(content)
		w.Show()
	}
}

func main() {
	app := app.New()
	w := app.NewWindow("TwitterGUI")

	text := canvas.NewText("No Result", color.Black)
	text.Alignment = fyne.TextAlignCenter
	w.SetContent(text)
	w.Resize(fyne.NewSize(300, 300))

	go setTweetPer10Seconds(w)

	w.ShowAndRun()

	// Run command
	// FYNE_FONT=C:\\Windows\\Fonts\\Meiryo.ttc go run main.go
}

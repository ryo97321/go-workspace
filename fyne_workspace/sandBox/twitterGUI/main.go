package main

import (
	"image/color"
	"net/url"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/ChimeraCoder/anaconda"
)

type MyTweetStruct struct {
	username     string
	createAtTime time.Time
	fullText     string
}

// *** Now Developing ***
// func getTweet() MyTweetStruct {

// 	var accessToken = os.Getenv("twitterAPIAccessToken")
// 	var accessTokenSecret = os.Getenv("twitterAPIAccessTokenSecret")
// 	var consumerKey = os.Getenv("twitterAPIConsumerKey")
// 	var consumerSecret = os.Getenv("twitterAPIConsumerSecret")

// 	anaconda.SetConsumerKey(consumerKey)
// 	anaconda.SetConsumerSecret(consumerSecret)
// 	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

// 	v := url.Values{}
// 	v.Set("count", "1")

// 	myTweetStruct := MyTweetStruct{"", time.Now(), ""}

// 	searchWord := "グラブル"
// 	searchResult, err := api.GetSearch(searchWord, v)
// 	if err != nil {
// 		return myTweetStruct
// 	}

// 	tweetUserName := ""
// 	tweetCreatedAtTime := time.Now()
// 	tweetFullText := ""

// 	if len(searchResult.Statuses) == 0 {
// 		tweetFullText = "No Result"
// 	} else {
// 		for _, tweet := range searchResult.Statuses {
// 			tweetUserName = tweet.User.Name
// 			tweetCreatedAtTime, _ = tweet.CreatedAtTime()
// 			tweetFullText = tweet.FullText
// 		}
// 	}

// 	myTweetStruct.username = tweetUserName
// 	myTweetStruct.createAtTime = tweetCreatedAtTime
// 	myTweetStruct.fullText = tweetFullText

// 	return myTweetStruct
// }

func getTweetFullText() string {

	var accessToken = os.Getenv("twitterAPIAccessToken")
	var accessTokenSecret = os.Getenv("twitterAPIAccessTokenSecret")
	var consumerKey = os.Getenv("twitterAPIConsumerKey")
	var consumerSecret = os.Getenv("twitterAPIConsumerSecret")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	v := url.Values{}
	v.Set("count", "1")

	searchWord := "グラブル"
	searchResult, err := api.GetSearch(searchWord, v)
	if err != nil {
		return "GetSearch Error"
	}

	tweetFullText := ""
	if len(searchResult.Statuses) == 0 {
		tweetFullText = "No Result"
	} else {
		for _, tweet := range searchResult.Statuses {
			tweetFullText = tweet.FullText
		}
	}

	return tweetFullText
}

func setTweetPer10Seconds(text *canvas.Text, w fyne.Window) {
	for {
		text.Refresh()

		time.Sleep(time.Second * 10)
		text.Text = getTweetFullText()

		w.SetContent(text)
		w.Show()
	}
}

// *** Now Developing ***
// func setTweetPer10Seconds(text *canvas.Text, w fyne.Window) {
// 	for {
// 		text.Refresh()

// 		time.Sleep(time.Second * 10)
// 		// text.Text = getTweetFullText()

// 		myTweetStruct := getTweet()

// 		text.Text = myTweetStruct.fullText
// 		username := myTweetStruct.username
// 		createdAtTime := myTweetStruct.createAtTime

// 		usernameText := canvas.NewText(username, color.Black)
// 		usernameText.Alignment = fyne.TextAlignCenter
// 		createdAtTimeText := canvas.NewText(createdAtTime.String(), color.Black)
// 		createdAtTimeText.Alignment = fyne.TextAlignCenter

// 		content := container.New(layout.NewVBoxLayout(), createdAtTimeText, usernameText, text)

// 		w.SetContent(content)
// 		w.Show()
// 	}
// }

func main() {
	app := app.New()
	w := app.NewWindow("TwitterGUI")

	text := canvas.NewText("No Result", color.Black)
	text.Alignment = fyne.TextAlignCenter
	w.SetContent(text)
	w.Resize(fyne.NewSize(300, 300))

	go setTweetPer10Seconds(text, w)

	w.ShowAndRun()

	// Run command
	// FYNE_FONT=C:\\Windows\\Fonts\\Meiryo.ttc go run main.go
}

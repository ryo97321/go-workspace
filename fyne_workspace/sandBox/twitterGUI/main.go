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

func main() {
	app := app.New()
	w := app.NewWindow("TwitterGUI")

	text := canvas.NewText("No Result", color.Black)
	text.Alignment = fyne.TextAlignCenter
	w.SetContent(text)
	w.Resize(fyne.NewSize(300, 300))

	go setTweetPer10Seconds(text, w)

	w.ShowAndRun()
}

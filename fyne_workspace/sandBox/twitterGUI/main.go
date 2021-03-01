package main

import (
	"image/color"
	"net/url"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/ChimeraCoder/anaconda"
)

// MyTweetStruct 3つのフィールドを持つTweetの構造体
type MyTweetStruct struct {
	username      string    // ユーザー名
	createdAtTime time.Time // Tweetの生成時刻
	fullTextLines []string  // 本文を改行で区切ったもの
}

// Tweetを取得する関数
func getMyTweetStruct() MyTweetStruct {

	var accessToken = os.Getenv("twitterAPIAccessToken")
	var accessTokenSecret = os.Getenv("twitterAPIAccessTokenSecret")
	var consumerKey = os.Getenv("twitterAPIConsumerKey")
	var consumerSecret = os.Getenv("twitterAPIConsumerSecret")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	v := url.Values{}
	v.Set("count", "1")            // 1件のTweetを取得
	v.Set("result_type", "recent") // 直近のTweetを取得

	myTweetStruct := MyTweetStruct{"", time.Now(), nil}

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
	myTweetStruct.createdAtTime = tweetCreatedAtTime

	fullTextLines := strings.Split(tweetFullText, "\n")
	myTweetStruct.fullTextLines = fullTextLines

	return myTweetStruct
}

// 時間をJSTに変換する関数
func timeToJST(createdAtTime time.Time) time.Time {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	createdAtTimeJST := createdAtTime.In(jst)

	return createdAtTimeJST
}

// 10秒おきにTweetを表示する関数
func setTweetPer10Seconds(w fyne.Window) {
	for {
		time.Sleep(time.Second * 10)

		myTweetStruct := getMyTweetStruct()

		fullTextLines := myTweetStruct.fullTextLines
		username := myTweetStruct.username
		createdAtTime := myTweetStruct.createdAtTime

		usernameTextObject := canvas.NewText(username, color.Black)
		createdAtTimeTextObject := canvas.NewText(timeToJST(createdAtTime).String(), color.Black)

		content := container.New(layout.NewVBoxLayout(), createdAtTimeTextObject, usernameTextObject)

		for _, line := range fullTextLines {
			fullTextLineObject := canvas.NewText(line, color.Black)
			content.Add(fullTextLineObject)
		}

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

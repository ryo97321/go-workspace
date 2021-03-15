package main

import (
	"bufio"
	"fmt"
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
	fullTextLines []string  // 本文を「改行文字」もしくは「40文字」で分割したもの
}

// fullTextを分割する関数（改行区切り かつ 1行の最大文字数は40文字）
func splitFullText(fullText string) []string {
	fullTextLines := strings.Split(fullText, "\n")

	nMaxString := 40 // 1行の最大文字数
	var splitedFullText []string
	for _, fullTextLine := range fullTextLines {
		var splitedLines []string
		runes := []rune(fullTextLine)
		if len(runes) < nMaxString { // 40文字未満なら、そのままsplitedFullTextに加える
			splitedLines = append(splitedLines, string(runes))
		} else {
			s := ""
			for i, r := range runes {
				s += string(r)

				if (i+1)%nMaxString == 0 {
					splitedLines = append(splitedLines, s)
					s = ""
				}
			}
			if s != "" { // s に文字列が残っていたらsplitedLinesに加える
				splitedLines = append(splitedLines, s)
			}
		}

		// 40文字ずつに分割した文字列スライスをsplitedFullTextに加える
		for _, splitedLine := range splitedLines {
			splitedFullText = append(splitedFullText, splitedLine)
		}
	}

	return splitedFullText
}

// Tweetを取得する関数
func getMyTweetStruct(searchWord string) MyTweetStruct {

	var accessToken = os.Getenv("twitterAPIAccessToken")
	var accessTokenSecret = os.Getenv("twitterAPIAccessTokenSecret")
	var consumerKey = os.Getenv("twitterAPIConsumerKey")
	var consumerSecret = os.Getenv("twitterAPIConsumerSecret")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	v := url.Values{}
	v.Set("count", "1")            // 1回の検索で1件のTweetを取得するように設定
	v.Set("result_type", "recent") // 直近のTweetを取得するように設定

	myTweetStruct := MyTweetStruct{"", time.Now(), nil}

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

	fullTextLines := splitFullText(tweetFullText)
	myTweetStruct.fullTextLines = fullTextLines

	return myTweetStruct
}

// 時間をJSTに変換する関数
func timeToJST(createdAtTime time.Time) time.Time {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	createdAtTimeJST := createdAtTime.In(jst)

	return createdAtTimeJST
}

// 時間を「YYYY/MM/DD HH:mm:ss」の形式に変換する関数
func formatCreatedAtTime(jstTime time.Time) string {
	year := jstTime.Year()
	month := fmt.Sprintf("%d", jstTime.Month())
	day := jstTime.Day()
	hour := jstTime.Hour()
	minute := jstTime.Minute()
	second := jstTime.Second()

	formatedJstTime := fmt.Sprintf("%d/%02s/%02d %02d:%02d:%02d", year, month, day, hour, minute, second)

	return formatedJstTime
}

// 10秒おきにTweetを表示する関数
func setTweetPer10Seconds(w fyne.Window, searchWord string) {
	for {
		time.Sleep(time.Second * 10)

		myTweetStruct := getMyTweetStruct(searchWord)

		fullTextLines := myTweetStruct.fullTextLines
		username := myTweetStruct.username
		createdAtTime := myTweetStruct.createdAtTime

		usernameTextObject := canvas.NewText(username, color.White)

		formatedCreatedAtTime := formatCreatedAtTime(timeToJST(createdAtTime))
		createdAtTimeTextObject := canvas.NewText(formatedCreatedAtTime, color.White)

		content := container.New(layout.NewVBoxLayout(), createdAtTimeTextObject, usernameTextObject)

		content.Add(canvas.NewText("----------", color.White))
		for _, line := range fullTextLines {
			fullTextLineObject := canvas.NewText(line, color.White)
			content.Add(fullTextLineObject)
		}
		content.Add(canvas.NewText("----------", color.White))

		w.SetContent(content)
		w.Show()
	}
}

func main() {
	app := app.New()
	w := app.NewWindow("TwitterGUI")

	text := canvas.NewText("No Result", color.White)
	text.Alignment = fyne.TextAlignCenter
	w.SetContent(text)
	w.Resize(fyne.NewSize(500, 300))

	scanner := bufio.NewScanner(os.Stdin)
	var searchWord string

	// searchWord が空のときループ
	for {
		fmt.Print("検索ワード：")
		scanner.Scan()
		searchWord = scanner.Text()

		if searchWord != "" {
			break
		}
	}

	title := fmt.Sprintf("TwitterAPI [検索ワード：%s]", searchWord)
	w.SetTitle(title)

	go setTweetPer10Seconds(w, searchWord)

	w.ShowAndRun()

	// Run command
	// FYNE_FONT=C:\\Windows\\Fonts\\Meiryo.ttc go run main.go
}

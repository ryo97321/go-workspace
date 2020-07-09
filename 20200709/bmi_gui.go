package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func calcBMI(weight, height float64) (string, string) {
	var status string
	bmi := weight / (height * height)

	if bmi < 18.5 {
		status = "低体重（やせ）"
	} else if bmi >= 18.5 {
		status = "普通体重"
	} else {
		status = "肥満"
	}

	bmiStr := strconv.FormatFloat(bmi, 'f', -1, 64)

	return bmiStr, status
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	// /
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	// /bmi
	router.POST("/bmi", func(ctx *gin.Context) {
		// Postで渡された値を受け取る
		weightStr := ctx.PostForm("weight")
		heightStr := ctx.PostForm("height")

		weight, _ := strconv.ParseFloat(weightStr, 64)
		height, _ := strconv.ParseFloat(heightStr, 64)

		bmi, status := calcBMI(weight, height)

		ctx.HTML(200, "bmi.html", gin.H{"bmi": bmi, "status": status})
	})

	router.Run()
}

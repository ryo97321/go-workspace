package main

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f, _ := excelize.OpenFile("hello.xlsx")
	c, _ := f.GetCellValue("Sheet2", "C2")

	println(c)
}


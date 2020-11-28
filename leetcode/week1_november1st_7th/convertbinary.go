package main

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	head := []int{0, 0}
	linkedList := list.New()

	for i := 0; i < len(head); i++ {
		item := head[i]
		linkedList.PushBack(item)
	}

	linkedListLen := linkedList.Len()
	exp := linkedListLen - 1
	var result int = 0
	for e := linkedList.Front(); e != nil; e = e.Next() {
		value := e.Value.(int) * int(math.Pow(2.0, float64(exp)))
		result += value
		exp--
	}

	fmt.Println(result)
}

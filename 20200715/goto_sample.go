package main

func myFunc() {
	i := 0
Here:
	println(i)
	i++
	if i == 10 {
		goto End
	}
	goto Here
End:
	println("Finish")
}

func main() {
	myFunc()
}

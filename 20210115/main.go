package main

import "bytes"

func main() {
	var buf bytes.Buffer
	b := byte('a')
	buf.WriteByte(b)
	b = byte('b')
	buf.WriteByte(b)
	b = byte('c')
	buf.WriteByte(b)

	println(buf.String())
}

package main

import (
	"fmt"
	"os"
)

type Stringer interface {
	String() string
}

func ToStringer(v interface{}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s, nil
	}
	return nil, MyError("CastError")
}

type MyError string

func (e MyError) Error() string {
	return string(e)
}

type S string

func (s S) String() string {
	return string(s)
}

func main() {
	v := S("bool")
	if s, err := ToStringer(v); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	} else {
		fmt.Println(s.String())
	}
}

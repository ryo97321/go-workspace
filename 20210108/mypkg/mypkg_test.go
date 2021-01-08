package mypkg_test

import (
	"projects/go-workspace/20210108/mypkg"
	"testing"
)

func TestHex_String(t *testing.T) {
	expect := "a"
	actual := mypkg.Hex(10).String()
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

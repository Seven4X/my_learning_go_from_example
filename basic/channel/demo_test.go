package main

import "testing"

func TestCh(t *testing.T) {
	var x chan string
	var y chan string
	println(x == y)
}

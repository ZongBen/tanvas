package main

import (
	"time"
	"unicode/utf8"

	"github.com/ZongBen/tanvas"
	"github.com/ZongBen/tanvas/tanminal"
)

func main() {
	word := "Hello"
	s_len := utf8.RuneCountInString(word)
	startupPosition := 1 - s_len
	t := new(tanminal.Tanminal)
	c := tanvas.CreateCanvas(10, 1, 1)
	s := c.CreateSection(startupPosition, 0, s_len, 1, 1)

	s.SetRow(0, 0, word)

	x := startupPosition
	width, _, _ := c.GetDimensions()
	for {
		if x > width {
			x = startupPosition
		}
		c.MoveSection(&s, x, 0)
		t.Flush(&c)
		x++
		<-time.After(1000 * time.Millisecond)
	}
}

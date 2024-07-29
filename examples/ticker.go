package main

import (
	"time"
	"unicode/utf8"

	"github.com/ZongBen/tanvas"
	"github.com/ZongBen/tanvas/tanminal"
)

func main() {
	t := new(tanminal.Tanminal)
	c := tanvas.CreateCanvas(20, 2, 1)

	go setTicker(&c, 1, "Hello, World!", 100)
	go setTicker(&c, 2, "This is a ticker!", 200)

	for {
		t.Flush(&c)
		<-time.After(33 * time.Millisecond)
	}
}

func setTicker(c *tanvas.Canvas, line int, word string, speed time.Duration) {
	width, _, _ := c.GetDimensions()
	s_len := utf8.RuneCountInString(word)
	startupPosition := width
	s := c.CreateSection(startupPosition, line-1, s_len, 1, 1)

	s.SetRow(0, 0, word)

	x := startupPosition
	for {
		if x < -s_len {
			x = startupPosition
		}
		c.MoveSection(&s, x, line-1)
		x--
		<-time.After(speed * time.Millisecond)
	}
}

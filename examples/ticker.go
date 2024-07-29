package main

import (
	"time"

	"github.com/ZongBen/tanvas"
	"github.com/ZongBen/tanvas/tanminal"
)

func main() {
	t := new(tanminal.Tanminal)
	c := tanvas.CreateCanvas(10, 1, 1)
	s := c.CreateSection(-4, 0, 10, 1, 1)

	s.SetRow(0, 0, "Hello")

	x := -4
	width, _, _ := c.GetDimensions()
	for {
		t.Flush(&c)
		if x > width {
			x = -4
		}
		c.MoveSection(&s, x, 0)
		x++
		<-time.After(100 * time.Millisecond)
	}
}

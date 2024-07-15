package main

import (
	"github.com/ZongBen/tanvas/utiles"
)

func main() {
	c := utiles.CreateCanvas(10, 10)
	s := c.CreateSection(0, 0, 5, 5)

	ta := s.CreateTextArea(true)

	ta.SetContent("Hello, World!")

	c.Render()
}

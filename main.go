package main

import (
	"github.com/ZongBen/tanvas/utiles"
)

func main() {
	c := utiles.CreateCanvas(10, 10)
	s := c.CreateSection(0, 0, 5, 5)

	s.SetChar(0, 0, 'H')
	s.SetChar(1, 0, 'e')
	s.SetChar(2, 0, 'l')
	s.SetChar(3, 0, 'l')
	s.SetChar(4, 0, 'o')

	c.Render()
}

package main

import (
	"github.com/ZongBen/tanvas/pkg/utils"
)

func main() {
	c := utils.CreateCanvas(10, 3, 3)

	s0 := c.CreateSection(0, 0, 10, 3, 0)
	s1 := c.CreateSection(0, 0, 10, 3, 1)
	s2 := c.CreateSection(0, 0, 10, 3, 2)

	s0.SetRow(0, "0123456789")
	s0.SetRow(1, "abcdefghij")
	s0.SetRow(2, "ABCDEFGHIJ")

	s1.SetRow(0, "takebylay1")

	s2.SetRowOffset(1, 3, "test")
	c.Render()
}

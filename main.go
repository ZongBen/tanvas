package main

import (
	"github.com/ZongBen/tanvas/pkg/utils"
)

func main() {
	c1 := utils.CreateCanvas(10, 3, 3)
	c2 := utils.CreateCanvas(10, 3, 3)

	s0 := c1.CreateSection(0, 0, 10, 3, 0)
	s1 := c1.CreateSection(0, 0, 10, 3, 1)
	s2 := c1.CreateSection(0, 0, 10, 3, 2)

	c2_s0 := c2.CreateSection(0, 0, 10, 3, 0)

	c2_s0.SetRow(0, 0, "0123456789")

	s0.SetRow(0, 0, "0123456789")
	s0.SetRow(0, 1, "abcdefghij")
	s0.SetRow(0, 2, "ABCDEFGHIJ")

	s1.SetRow(0, 0, "takebylay1")

	s2.SetRow(3, 1, "test               ")
	c1.Render()
	c2.Render()
}

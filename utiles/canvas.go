package utiles

import "strings"

type canvas struct {
	width  int
	height int
	char   [][]rune
}

func CreateCanvas(width, height int) canvas {
	c := canvas{width: width, height: height}
	c.char = make([][]rune, height)
	for i := range c.char {
		c.char[i] = make([]rune, width)
	}
	return c
}

func (c *canvas) CreateSection(x, y, width, height int) section {
	s := section{width: width, height: height}
	s.char = make([][]*rune, height)
	for i := range s.char {
		s.char[i] = make([]*rune, width)
		for j := range s.char[i] {
			s.char[i][j] = &c.char[y+i][x+j]
		}
	}
	return s
}

func (c *canvas) Render() {
	sb := new(strings.Builder)

	for _, row := range c.char {
		for _, char := range row {
			if char == 0 {
				sb.WriteString(" ")
				continue
			}
			sb.WriteString(string(char))
		}
		sb.WriteString("\n")
	}

	print(sb.String())
}

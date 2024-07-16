package utiles

import "strings"

type canvas struct {
	width  int
	height int
	depth  int
	single [][][]single
}

type single struct {
	char    rune
	display bool
}

func CreateCanvas(width, height, depth int) canvas {
	c := canvas{width: width, height: height, depth: depth}
	c.single = make([][][]single, height)
	for i := range c.single {
		c.single[i] = make([][]single, width)
		for j := range c.single[i] {
			c.single[i][j] = make([]single, depth)
		}
	}
	return c
}

func (c *canvas) CreateSection(x, y, width, height, layer int) section {
	s := section{width: width, height: height, display: true}
	s.single = make([][]*single, height)
	for i := range s.single {
		s.single[i] = make([]*single, width)
		for j := range s.single[i] {
			s.single[i][j] = &c.single[y+i][x+j][layer]
		}
	}
	return s
}

func (c *canvas) Render() {
	sb := new(strings.Builder)

	for _, row := range c.single {
		for _, char := range row {
			for depth := c.depth - 1; depth >= 0; depth-- {
				if char[depth].char == 0 {
					if depth == 0 {
						sb.WriteString(" ")
					}
					continue
				}

				if char[depth].display {
					sb.WriteString(string(char[depth].char))
					break
				}
			}
		}
		sb.WriteString("\n")
	}
	print(sb.String())
}

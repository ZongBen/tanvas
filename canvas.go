package tanvas

import "strings"

type Canvas interface {
	CreateSection(x, y, width, height, layer int) section
	Render() string
}

type canvas struct {
	width     int
	height    int
	depth     int
	container [][][]single
}

var sb = new(strings.Builder)

func CreateCanvas(width, height, depth int) canvas {
	c := canvas{width: width, height: height, depth: depth}
	c.container = make([][][]single, height)
	for i := range c.container {
		c.container[i] = make([][]single, width)
		for j := range c.container[i] {
			c.container[i][j] = make([]single, depth)
		}
	}
	return c
}

func (c *canvas) CreateSection(x, y, width, height, layer int) section {
	s := section{width: width, height: height, display: true}
	s.plate = make([][]*single, height)
	for i := range s.plate {
		s.plate[i] = make([]*single, width)
		for j := range s.plate[i] {
			s.plate[i][j] = &c.container[y+i][x+j][layer]
		}
	}
	return s
}

func (c *canvas) Render() string {
	sb.Reset()
	for _, row := range c.container {
		for _, single := range row {
			for depth := c.depth - 1; depth >= 0; depth-- {
				if single[depth].char == 0 {
					if depth == 0 {
						sb.WriteString(" ")
					}
					continue
				}

				if single[depth].display {
					sb.WriteString(string(single[depth].char))
					break
				}
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

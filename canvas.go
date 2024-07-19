package tanvas

import (
	"sync"
)

type Canvas interface {
	CreateSection(x, y, width, height, layer int) section
	Render() string
	Clear()
	SetOffset(x, y int)
	GetDimensions() (int, int, int)
}

type canvas struct {
	width     int
	height    int
	depth     int
	offset_x  int
	offset_y  int
	container [][][]single
}

func (c *canvas) GetDimensions() (int, int, int) {
	return c.width, c.height, c.depth
}

func (c *canvas) SetOffset(x, y int) {
	c.offset_x = x
	c.offset_y = y
}

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

func (c *canvas) Project() string {
	wg := new(sync.WaitGroup)
	width := c.width + c.offset_x + 1 // +1 for newline
	height := c.height
	result := make([]rune, width*height)

	wg.Add(1)
	offset_y := ""
	go func() {
		for i := 0; i < c.offset_y; i++ {
			offset_y += "\n"
		}
		wg.Done()
	}()

	for y, row := range c.container {
		wg.Add(1)
		go func(y int, asyncRow [][]single) {
			for x, cell := range asyncRow {
				wg.Add(1)
				go func(x, y int, asynCell []single) {
					for depth := c.depth - 1; depth >= 0; depth-- {
						if asynCell[depth].char == 0 {
							if depth == 0 {
								result[y*width+x] = ' '
							}
							continue
						}

						if asynCell[depth].display {
							result[y*width+x] = asynCell[depth].char
							break
						}
					}
					wg.Done()
				}(x+c.offset_x, y, cell)
			}
			for i := 0; i < c.offset_x; i++ {
				result[y*width+i] = ' '
			}
			result[y*width+c.width+c.offset_x] = '\n'
			wg.Done()
		}(y, row)
	}
	wg.Wait()
	return offset_y + string(result)
}

func (c *canvas) Clear() {
	for _, row := range c.container {
		for _, single := range row {
			for i := range single {
				single[i].char = 0
			}
		}
	}
}

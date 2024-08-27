package tanvas

import (
	"errors"
	"sync"
)

type Canvas struct {
	width     int
	height    int
	layer     int
	offset_x  int
	offset_y  int
	container [][][]single
}

// Get the dimensions of the canvas.
func (c *Canvas) GetDimensions() (int, int, int) {
	return c.width, c.height, c.layer
}

// Set the offset of the canvas.
func (c *Canvas) SetOffset(offsetX, offsetY int) {
	c.offset_x = offsetX
	c.offset_y = offsetY
}

// Create a new canvas with the given width, height, and layer.
// The layer is the number of layers that can be displayed on the canvas.
// All parameters are 1 index based.
func CreateCanvas(width, height, layer int) (Canvas, error) {
	if width < 1 || height < 1 || layer < 1 {
		return Canvas{}, errors.New("Width, height, and layer must be greater than 0.")
	}
	wg := new(sync.WaitGroup)
	c := Canvas{width: width, height: height, layer: layer}
	c.container = make([][][]single, height)
	for j := range c.container {
		wg.Add(1)
		c.container[j] = make([][]single, width)
		go func(j int) {
			for i := range c.container[j] {
				c.container[j][i] = make([]single, layer)
			}
			wg.Done()
		}(j)
	}
	wg.Wait()
	return c, nil
}

// Create a new section on the canvas with the given offset, width, height, and layer.
// It is fine to create a section that is outside the bounds of the canvas.
// The section will be clipped to the canvas.
func (c *Canvas) CreateSection(offsetX, offsetY, width, height, layer int) Section {
	wg := new(sync.WaitGroup)
	s := Section{width: width, height: height, layer: layer, display: true}
	s.shadow = make([][]*single, height)
	s.content = make([][]single, height)
	for j := range s.shadow {
		wg.Add(1)
		s.shadow[j] = make([]*single, width)
		s.content[j] = make([]single, width)
		go func(j int) {
			for i := range s.shadow[j] {
				x := offsetX + i
				y := offsetY + j
				if y >= c.height || x >= c.width || x < 0 || y < 0 {
					continue
				}
				s.shadow[j][i] = &c.container[offsetY+j][offsetX+i][layer-1]
			}
			wg.Done()
		}(j)
	}
	wg.Wait()
	return s
}

// Move the section to the given offset.
func (c *Canvas) MoveSection(s *Section, offsetX, offsetY int) {
	s.Clear()
	wg := new(sync.WaitGroup)
	for j := range s.shadow {
		wg.Add(1)
		go func(j int) {
			for i := range s.shadow[j] {
				x := offsetX + i
				y := offsetY + j
				if y >= c.height || x >= c.width || x < 0 || y < 0 {
					continue
				}
				s.shadow[j][i] = &c.container[offsetY+j][offsetX+i][s.layer-1]
				*s.shadow[j][i] = s.content[j][i]
			}
			wg.Done()
		}(j)
	}
	wg.Wait()
}

// Project 3D canvas to string.
func (c *Canvas) Project() string {
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
					for layer := c.layer - 1; layer >= 0; layer-- {
						if asynCell[layer].char == 0 {
							if layer == 0 {
								result[y*width+x] = ' '
							}
							continue
						}

						if asynCell[layer].display {
							result[y*width+x] = asynCell[layer].char
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

// Clear the canvas.
func (c *Canvas) Clear() {
	wg := new(sync.WaitGroup)
	for _, row := range c.container {
		wg.Add(1)
		go func(row [][]single) {
			for _, cell := range row {
				wg.Add(1)
				go func(cell []single) {
					for layer := range cell {
						cell[layer].char = 0
					}
					wg.Done()
				}(cell)
			}
			wg.Done()
		}(row)
	}
	wg.Wait()
}

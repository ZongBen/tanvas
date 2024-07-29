package tanminal

import (
	"strings"
	"sync"

	"github.com/ZongBen/tanvas"
	"github.com/nsf/termbox-go"
)

func init() {
	if !termbox.IsInit {
		termbox.Init()
	}
}

type Tanminal struct {
	offset_x              int
	offset_y              int
	horizontalAlignCenter bool
	verticalAlignCenter   bool
}

func CreateTanminal() Tanminal {
	return Tanminal{}
}

func (t *Tanminal) Flush(c *tanvas.Canvas) {
	offsetX, offsetY := t.getOffset()
	Clear()
	wg := new(sync.WaitGroup)
	lines := strings.Split(c.Project(), "\n")
	for y, line := range lines {
		wg.Add(1)
		go func(y int, line string) {
			for x, char := range line {
				termbox.SetChar(x+offsetX, y+offsetY, char)
			}
			wg.Done()
		}(y, line)
	}
	wg.Wait()
	termbox.Flush()
}

func Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func Close() {
	termbox.Close()
}

func (t *Tanminal) AlignCenter(offsetX, offsetY int) {
	t.AlignCenterX(offsetX)
	t.AlignCenterY(offsetY)
}

func (t *Tanminal) AlignCenterX(offsetX int) {
	t.horizontalAlignCenter = true
	t.offset_x = offsetX
}

func (t *Tanminal) AlignCenterY(offsetY int) {
	t.verticalAlignCenter = true
	t.offset_y = offsetY
}

func getTerminalCenter() (int, int) {
	width, height := termbox.Size()
	return width / 2, height / 2
}

func (t *Tanminal) getOffset() (int, int) {
	width, height := getTerminalCenter()
	offsetX, offsetY := 0, 0
	if t.horizontalAlignCenter {
		offsetX = width + t.offset_x
	}
	if t.verticalAlignCenter {
		offsetY = height + t.offset_y
	}
	return offsetX, offsetY
}

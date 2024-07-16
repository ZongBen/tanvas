package utiles

import (
	"strings"
)

type textarea struct {
	width    int
	height   int
	char     [][]rune
	autoWrap bool
}

func CreateTextArea(width, height int, autoWrap bool) textarea {
	area := textarea{width: width, height: height, autoWrap: autoWrap}
	area.char = make([][]rune, height)
	for i := range area.char {
		area.char[i] = make([]rune, width)
	}
	return area
}

func (ta *textarea) SetContent(content string) {
	lineList := strings.Split(content, "\n")
	if !ta.autoWrap {
		for y, line := range lineList[:min(ta.height, len(lineList))] {
			copy(ta.char[y], []rune(line))
		}
	} else {
		rowOffset := 0
		for i, line := range lineList {
			times := len(line) / ta.width
			for j := 0; j <= times; j++ {
				rowOffset += j
				if rowOffset+i >= ta.height {
					break
				}
				copy(ta.char[rowOffset+i], []rune(line))
			}
		}
	}
}

func (ta *textarea) ClearTextArea() {
	ta.char = make([][]rune, ta.height)
	for i := range ta.char {
		ta.char[i] = make([]rune, ta.width)
	}
}

func (ta *textarea) GetContent() [][]rune {
	return ta.char
}

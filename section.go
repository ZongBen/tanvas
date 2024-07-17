package tanvas

import ()

type Section interface {
	SetChar(row, col int, char rune)
	SetRow(offset, row int, content string)
	SetCol(offset, col int, content string)
	SetDisplay(display bool)
	ToggleDisplay()
}

type section struct {
	width   int
	height  int
	plate   [][]*single
	display bool
}

func (s *section) SetChar(row, col int, char rune) {
	*s.plate[row][col] = single{char: char, display: s.display}
}

func (s *section) SetRow(offset, row int, content string) {
	max_len := min(len(content), s.width-offset)
	for i, char := range content[:max_len] {
		s.SetChar(row, offset+i, char)
	}
}

func (s *section) SetCol(offset, col int, content string) {
	max_len := min(len(content), s.height-offset)
	for i, char := range content[:max_len] {
		s.SetChar(offset+i, col, char)
	}
}

func (s *section) SetDisplay(display bool) {
	s.display = display
	s.setSectionDisplay(display)
}

func (s *section) ToggleDisplay() {
	s.display = !s.display
	s.setSectionDisplay(s.display)
}

func (s *section) setSectionDisplay(display bool) {
	for _, row := range s.plate {
		for _, cell := range row {
			cell.display = display
		}
	}
}

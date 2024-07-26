package tanvas

import ()

type Section interface {
	SetChar(row, col int, char rune)
	SetRow(offset, row int, content string)
	SetCol(offset, col int, content string)
	SetDisplay(display bool)
	Clear()
}

type section struct {
	width   int
	height  int
	layer   int
	plate   [][]single
	shadow  [][]*single
	display bool
}

func (s *section) SetChar(row, col int, char rune) {
	if row < 0 || row >= s.height || col < 0 || col >= s.width {
		return
	}
	single := single{char: char, display: s.display}
	s.plate[row][col] = single
	if s.shadow[row][col] != nil {
		*s.shadow[row][col] = single
	}
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

func (s *section) setSectionDisplay(display bool) {
	for _, row := range s.shadow {
		for _, cell := range row {
			cell.display = display
		}
	}
}

func (s *section) Clear() {
	for _, row := range s.shadow {
		for _, cell := range row {
			if cell == nil {
				continue
			}
			cell.char = 0
		}
	}
}

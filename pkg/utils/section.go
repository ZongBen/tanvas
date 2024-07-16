package utils

type section struct {
	width   int
	height  int
	single  [][]*single
	display bool
}

func (s *section) SetChar(row, col int, char rune) {
	*s.single[row][col] = single{char: char, display: s.display}
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
	for _, row := range s.single {
		for _, cell := range row {
			cell.display = display
		}
	}
}

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

func (s *section) SetRow(row int, content string) {
	for i, char := range content[:min(len(content), s.width)] {
		s.SetChar(row, i, char)
	}
}

func (s *section) SetRowOffset(offset, row int, content string) {
	for i, char := range content[:min(len(content), s.width)] {
		if offset+i >= s.width {
			break
		}
		s.SetChar(row, offset+i, char)
	}
}

func (s *section) SetCol(col int, content string) {
	for i, char := range content[:min(len(content), s.height)] {
		s.SetChar(i, col, char)
	}
}

func (s *section) SetColOffset(offset, col int, content string) {
	for i, char := range content[:min(len(content), s.height)] {
		if offset+i >= s.height {
			break
		}
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

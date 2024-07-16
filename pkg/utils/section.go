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

func (s *section) SetRowOffset(row, col int, content string) {
	for i, char := range content[:min(len(content), s.width)] {
		if col+i >= s.width {
			break
		}
		s.SetChar(row, col+i, char)
	}
}

func (s *section) SetCol(col int, content string) {
	for i, char := range content[:min(len(content), s.height)] {
		s.SetChar(i, col, char)
	}
}

func (s *section) SetColOffset(row, col int, content string) {
	for i, char := range content[:min(len(content), s.height)] {
		if row+i >= s.height {
			break
		}
		s.SetChar(row+i, col, char)
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

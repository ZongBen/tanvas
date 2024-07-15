package utiles

type section struct {
	width  int
	height int
	char   [][]*rune
}

func (s *section) CreateTextArea(autoWrap bool) textarea {
	area := CreateTextArea(s.width, s.height, autoWrap)
	area.char = s.char
	return area
}

func (s *section) SetChar(row, col int, char rune) {
	*s.char[row][col] = char
}

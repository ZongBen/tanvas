package utiles

type section struct {
	width  int
	height int
	char   [][]*rune
}

func (s *section) SetChar(row, col int, char rune) {
	*s.char[row][col] = char
}

func (s *section) CreateTextArea(autoWrap bool) textarea {
	return CreateTextArea(s.width, s.height, autoWrap)
}

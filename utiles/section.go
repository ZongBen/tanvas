package utiles

type section struct {
	width  int
	height int
	char   [][]*rune
}

func (s *section) SetChar(row, col int, char rune) {
	*s.char[row][col] = char
}

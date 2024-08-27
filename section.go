package tanvas

import "sync"

type Section struct {
	width   int
	height  int
	layer   int
	content [][]single
	shadow  [][]*single
	display bool
}

// Set a character at the given row and column.
func (s *Section) SetChar(row, col int, char rune) {
	if row < 0 || row >= s.height || col < 0 || col >= s.width {
		return
	}
	single := single{char: char, display: s.display}
	s.content[row][col] = single
	if s.shadow[row][col] != nil {
		*s.shadow[row][col] = single
	}
}

// Set a row of characters at the given offset.
func (s *Section) SetRow(offset, row int, content string) {
	max_len := min(len(content), s.width-offset)
	for i, char := range content[:max_len] {
		s.SetChar(row, offset+i, char)
	}
}

// Set a column of characters at the given offset.
func (s *Section) SetCol(offset, col int, content string) {
	max_len := min(len(content), s.height-offset)
	for i, char := range content[:max_len] {
		s.SetChar(offset+i, col, char)
	}
}

// Set section visibility.
func (s *Section) SetDisplay(display bool) {
	s.display = display
	s.setSectionDisplay(display)
}

func (s *Section) setSectionDisplay(display bool) {
	wg := new(sync.WaitGroup)
	for j := range s.shadow {
		wg.Add(1)
		go func(j int) {
			for i := range s.shadow[j] {
				s.content[j][i].display = display
				if s.shadow[j][i] != nil {
					s.shadow[j][i].display = display
				}
			}
			wg.Done()
		}(j)
	}
	wg.Wait()
}

// Clear the section.
func (s *Section) Clear() {
	wg := new(sync.WaitGroup)
	for _, row := range s.shadow {
		wg.Add(1)
		go func(row []*single) {
			for _, cell := range row {
				if cell == nil {
					continue
				}
				cell.char = 0
			}
			wg.Done()
		}(row)
	}
	wg.Wait()
}

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

func (s *Section) SetRow(offset, row int, content string) {
	max_len := min(len(content), s.width-offset)
	for i, char := range content[:max_len] {
		s.SetChar(row, offset+i, char)
	}
}

func (s *Section) SetCol(offset, col int, content string) {
	max_len := min(len(content), s.height-offset)
	for i, char := range content[:max_len] {
		s.SetChar(offset+i, col, char)
	}
}

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

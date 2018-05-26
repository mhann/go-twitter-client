package main

import (
	"github.com/gdamore/tcell"
)

type StatusBar struct {
	Message string
}

func (bar *StatusBar) Render(s tcell.Screen) {
	screenWidth, screenHeight := s.Size()

	y := screenHeight - 1 // Screen height is zero-indexed
	x := 0

	// Clear the status bar.
	for ; x < screenWidth; x++ {
		s.SetContent(x, y, ' ', []rune{}, 0)
	}

	// Write the text with two char offset from left
	x = 2
	for _, char := range bar.Message {
		if x > screenWidth {
			continue
		}

		s.SetContent(x, y, char, []rune{}, 0)
		x++
	}

	s.Show()
}

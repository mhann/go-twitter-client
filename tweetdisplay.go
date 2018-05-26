package main

import (
	"github.com/gdamore/tcell"
)

type TweetDisplay struct {
	Tweet Tweet
}

func (td *TweetDisplay) Render(s tcell.Screen) {
	screenWidth, screenHeight := s.Size()

	y := 0 // Screen height is zero-indexed
	x := 0

	tweetDisplayHeight := screenHeight - 1 // Don't overwrite status bar

	for xToClear := 0; xToClear < screenWidth; xToClear++ {
		for yToClear := 0; yToClear < tweetDisplayHeight; yToClear++ {
			s.SetContent(xToClear, yToClear, ' ', []rune{}, 0)
		}
	}

	tweetContent := td.Tweet.Contents

	x = 3
	y = 1
	for _, char := range tweetContent {
		if x >= screenWidth {
			continue
		}

		if char == '\n' {
			y++
			x = 3
		}

		s.SetContent(x, y, char, []rune{}, 0)
		x++
	}

	s.Show()
}

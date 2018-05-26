package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gdamore/tcell"
)

type screenSize struct {
	width  int
	height int
}

var screen tcell.Screen
var currentSize screenSize

func main() {
	// FIXME: Would be nicer if this was made in StartTweetListener and then passed back to us
	tweetChannel := make(chan Tweet)
	go StartTweetListener(tweetChannel)

	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Println("Could not start tcell for gomatrix. View ~/.gomatrix-log for error messages.")
		log.Printf("Cannot alloc screen, tcell.NewScreen() gave an error:\n%s", err)
		os.Exit(1)
	}

	err = screen.Init()
	if err != nil {
		fmt.Println("Could not start tcell for gomatrix. View ~/.gomatrix-log for error messages.")
		log.Printf("Cannot start gomatrix, screen.Init() gave an error:\n%s", err)
		os.Exit(1)
	}

	screen.HideCursor()
	screen.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorBlack))
	screen.Clear()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// make chan for tembox events and run poller to send events on chan
	eventChan := make(chan tcell.Event)
	go func() {
		for {
			event := screen.PollEvent()
			eventChan <- event
		}
	}()

	done := false

	for !done {
		select {
		case event := <-eventChan:
			switch ev := event.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyCtrlZ, tcell.KeyCtrlC:
					done = true
					continue
				case tcell.KeyRune:
					switch ev.Rune() {
					case 'w':
						writeTextAtBottom("Up")
					case 'a':
						writeTextAtBottom("Left")
					case 'd':
						writeTextAtBottom("Right")
					case 's':
						writeTextAtBottom("Down")
					}
				}
			}
		case tweet := <-tweetChannel:
			processNewTweet(tweet)
		case <-sigChan:
			done = true
			continue
		}
	}

	screen.Fini()
}

func writeTextAtBottom(text string) {
	//offsetX := 1
	//offsetY := 1

	//for _, char := range text {
	//	fmt.Println(char)
	emitStr(screen, 1, 1, tcell.StyleDefault, "test")
	//screen.Show()
	//offsetX++
	//}
}

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := 1

		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}

		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func processNewTweet(tweet Tweet) {
	//fmt.Println(tweet.Contents)
	//fmt.Println(" ------ ")
}
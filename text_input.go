package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var (
	inText = ""
	ctrl   = 0
)

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)

	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}

	return false
}

func inputUpdate(screen *ebiten.Image) (string, int) {
	// Add a string from InputChars, that returns string input by users.
	// Note that InputChars result changes every frame, so you need to call this
	// every frame.
	txt := string(ebiten.InputChars())
	ctrl = isNull

	// If the enter key is pressed, add a line break.
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyKPEnter) {
		//		inText += "\n"
		ctrl = hasCR
	}

	// If the backspace key is pressed, remove one character.
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(inText) >= 1 {
			inText = inText[:len(inText)-1]
		}
		ctrl = isNew
	}

	switch ctrl {
	case isNull:
		if len(txt) > 0 {
			inText += txt
			fmt.Println("isNull(txt) = ", inText)
			return inText, isNew
		}
		return inText, isNull
	case isNew:
		fmt.Println("isNew = ", inText)
		return inText, isNew
	case hasCR:
		fmt.Println("hasCR = ", inText)
		return inText, hasCR
	default:
		return "", isNull
	}
}

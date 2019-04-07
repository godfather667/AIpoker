// text_input.go performs the character input functions for AIpoker.
//
package main

import (
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
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

// inputUpdate reads characters entered on the default screen.
func inputUpdate(screen *ebiten.Image) (string, int) {
	// Add a string from InputChars, that returns string input by users.
	// Note that InputChars result changes every frame, so you need to call this
	// every frame.
	txt := string(ebiten.InputChars())
	ctrl = isNull

	// If the enter key is pressed, add a line break.
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyKPEnter) {
		ctrl = hasCR
	}

	// If the backspace key is pressed, remove one character.
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(inText) >= 1 {
			inText = inText[:len(inText)-1]
		}
		ctrl = isNew
	}

	// Insure that input characters are digits
	if len(txt) > 0 {
		if _, err := strconv.Atoi(txt); err != nil { // is txt a number?
			setError("Text Error - Bad Number", screen)
			inText = ""
			return inText, isNew
		}
	}
	//
	// Handle Resulting Control Mode Bit Actions
	//
	switch ctrl {
	case isNull:
		if len(txt) > 0 {
			inText += txt
			return inText, isNew
		}
		return inText, isNull
	case isNew:
		return inText, isNew
	case hasCR:
		return inText, hasCR
	default:
		return "", isNull
	}
}

package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	delay    = 30
	interval = 3
)

var tx = ""

// repeatingKeyPressed return true when key is pressed considering the repeatx state.
func repeatingKeyPressed(key ebiten.Key) bool {
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

func TextInput(screen *ebiten.Image) (tx string, e bool) {
	// Add a string from InputChars, that returns string input by users.
	// Note that InputChars result changes every frame, so you need to call this
	// every frame.
	tx += string(ebiten.InputChars())

	// If the enter key is pressed, add a line break.
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyKPEnter) {
		tx += "\n"
		return tx, true
	}

	// If the backspace key is pressed, remove one character.
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(tx) >= 1 {
			tx = tx[:len(tx)-1]
		}
	}
	return tx, false
}
